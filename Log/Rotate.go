package Log

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

func NewRotate(opts ...Option) (*RotateLog, error) {
	r := &RotateLog{
		mutex: &sync.Mutex{},
		close: make(chan struct{}, 1),
		//logPath: logPath,
	}
	for _, opt := range opts {
		opt(r)
	}

	if err := r.rotateFile(); err != nil {
		return nil, err
	}
	if r.RotateTime != 0 || r.RotateSize != 0 {
		go r.handleEvent()
	}

	return r, nil
}

func (r *RotateLog) Write(p []byte) (n int, err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	writeLen := int64(len(p))
	if writeLen > r.max() {
		return 0, fmt.Errorf(
			"write length %d exceeds maximum file size %d", writeLen, r.max(),
		)
	}

	if r.file == nil || r.size+writeLen > r.max() {
		if err := r.rotateFile(); err != nil {
			return 0, err
		}
	}

	n, err = r.file.Write(p)
	r.size += int64(n)

	return n, err
}

func (r *RotateLog) Close() error {
	r.close <- struct{}{}
	return r.file.Close()
}

//func (r *RotateLog) write() int64 {
//	if r.RotateSize == 0 {
//		return int64(defaultMaxSize * megabyte)
//	}
//	return int64(r.RotateSize) * int64(megabyte)
//}

func (r *RotateLog) max() int64 {
	if r.RotateSize == 0 {
		return int64(defaultMaxSize * megabyte)
	}
	return int64(r.RotateSize) * int64(megabyte)
}

func (r *RotateLog) hasRole(role int) bool {
	if r.role == 0 {
		r.role = defaultRole
	}
	return (r.role & role) == role
}

func (r *RotateLog) handleEvent() {
	for {
		select {
		case <-r.close:
			return
		case timeChan := <-r.rotateTimeChan:
			_ = r.timeRotateFile(timeChan)
		case sizeChan := <-r.rotateSizeChan:
			_ = r.sizeRotateFile(sizeChan)
		default:
			return
		}
	}
}

func (r *RotateLog) timeRotateFile(timeChan time.Time) error {
	if r.RotateTime != 0 {
		nextTime := r.nextRotateTime(timeChan, r.RotateTime)
		r.rotateTimeChan = time.After(nextTime)
	}
	r.mutex.Lock()
	defer r.mutex.Unlock()

	//bakName := r.backupName(r.LocalTime)
	err := r.rotateFile()
	if err != nil {
		return err
	}

	go func() {
		_ = r.deleteExpiredFile()
	}()

	return nil
}

func (r *RotateLog) nextRotateTime(now time.Time, duration time.Duration) time.Duration {
	nowUnixNao := now.UnixNano()
	NanoSecond := duration.Nanoseconds()
	nextRotateTime := NanoSecond - (nowUnixNao % NanoSecond)
	return time.Duration(nextRotateTime)
}

func (r *RotateLog) sizeRotateFile(sizeChan bool) error {
	//select {
	//case r.rotateSizeChan <- true:
	//default:
	//}
	r.mutex.Lock()
	defer r.mutex.Unlock()

	//bakName := r.backupName(r.LocalTime)
	err := r.rotateFile()
	if err != nil {
		return err
	}

	go func() {
		_ = r.deleteExpiredFile()
	}()

	return nil
}

func (r *RotateLog) rotateFile() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if r.file == nil {
		if dirErr := os.MkdirAll(r.FilePath, 0755); dirErr != nil {
			return dirErr
		}

		file, err := os.OpenFile(r.FilePath+"/"+r.FileName+defaultSuffix, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		r.file = file
		r.size = 0
		return nil
	} else {
		closeErr := r.file.Close()
		if closeErr != nil {
			return closeErr
		}

		if dirErr := os.MkdirAll(r.FilePath, 0755); dirErr != nil {
			return dirErr
		}

		info, err := osStat(r.FilePath + "/" + r.FileName + defaultSuffix)
		if err == nil {
			bakName := r.backupName(r.LocalTime)
			if renameErr := os.Rename(info.Name(), bakName); renameErr != nil {
				return fmt.Errorf("can't rename log file: %s", err)
			}
		}

		file, err := os.OpenFile(r.FilePath+"/"+r.FileName+defaultSuffix, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		r.file = file
		r.size = 0
	}

	go func() {
		_ = r.deleteExpiredFile()
	}()

	return nil
}
func (r *RotateLog) deleteExpiredFile() error {
	if r.MaxBackups == 0 && r.MaxAge == 0 && !r.Compress {
		return nil
	}
	// 文件列表
	matches, err := filepath.Glob(filepath.Join(r.FilePath, r.FileName, "-*"))
	if err != nil {
		return err
	}
	var files []sortFile
	for _, path := range matches {
		fileInfo, fileErr := os.Stat(path)
		if fileErr != nil {
			return fileErr
		} else {
			files = append(files, sortFile{fileInfo.ModTime(), fileInfo})
		}
	}
	sort.Sort(byFormatTime(files))

	// 删除超过备份数量的文件
	if r.MaxBackups > 0 && r.MaxBackups < len(files) {
		for fileIndex, file := range files {
			removeErr := os.Remove(file.Name())
			if removeErr != nil {
				return removeErr
			} else {
				files[fileIndex].timestamp = time.Time{}
				files[fileIndex].FileInfo = nil
			}
		}
	}

	// 删除超过最大保存天数的文件
	if r.MaxAge > 0 {
		expiredTime := time.Now().Add(-r.MaxAge)
		for fileIndex, file := range files {
			if file.timestamp.IsZero() {
				continue
			}
			if file.ModTime().After(expiredTime) {
				continue
			}
			removeErr := os.Remove(file.Name())
			if removeErr != nil {
				return removeErr
			} else {
				files[fileIndex].timestamp = time.Time{}
				files[fileIndex].FileInfo = nil
			}
		}
	}

	if r.Compress {
		for _, file := range files {
			if file.timestamp.IsZero() {
				continue
			}
			if !strings.HasSuffix(file.Name(), compressSuffix) {
				bakName := filepath.Join(file.Name(), compressSuffix)
				compressErr := r.compressFile(file, bakName)
				if compressErr != nil {
					return compressErr
				}
			}
		}

	}
	return nil
}

// 压缩文件
func (r *RotateLog) compressFile(src os.FileInfo, dst string) (err error) {
	f, err := os.Open(src.Name())
	if err != nil {
		return fmt.Errorf("failed to open log file: %v", err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	dstFile, err := os.OpenFile(dst, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, src.Mode())
	if err != nil {
		return fmt.Errorf("failed to open compressed log file: %v", err)
	}
	defer func(dstFile *os.File) {
		_ = dstFile.Close()
	}(dstFile)

	gz := gzip.NewWriter(dstFile)

	defer func() {
		if err != nil {
			_ = os.Remove(dst)
			err = fmt.Errorf("failed to compress log file: %v", err)
		}
	}()

	if _, copyErr := io.Copy(gz, f); copyErr != nil {
		return copyErr
	}
	if gzErr := gz.Close(); gzErr != nil {
		return gzErr
	}
	if dstErr := dstFile.Close(); dstErr != nil {
		return dstErr
	}

	if removeErr := os.Remove(src.Name()); removeErr != nil {
		return removeErr
	}

	return nil
}

func (r *RotateLog) backupName(local bool) string {
	t := currentTime()
	if !local {
		t = t.UTC()
	}

	timestamp := t.Format(backupTimeFormat)
	return filepath.Join(r.FilePath, fmt.Sprintf("%s-%s%s", r.FileName, timestamp, defaultSuffix))
}

type sortFile struct {
	timestamp time.Time
	os.FileInfo
}
type byFormatTime []sortFile

func (b byFormatTime) Less(i, j int) bool {
	return b[i].timestamp.After(b[j].timestamp)
}

func (b byFormatTime) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b byFormatTime) Len() int {
	return len(b)
}

//func removeElementByIndex[T comparable](s []T, index int) []T {
//	if index < 0 || index >= len(s) {
//		// 索引越界处理
//		return s
//	}
//	// 将索引前的部分与索引后的部分拼接起来
//	return append(s[:index], s[index+1:]...)
//}
