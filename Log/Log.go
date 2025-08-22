package Log

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

//func NewLogger(opts ...Option) (*RotateLog, error) {
//	r := &RotateLog{
//		mutex: &sync.Mutex{},
//		close: make(chan struct{}, 1),
//	}
//	for _, opt := range opts {
//		opt(r)
//	}
//
//	if err := os.Mkdir(filepath.Dir(r.FilePath), 0755); err != nil && !os.IsExist(err) {
//		return nil, err
//	}
//
//	if err := r.rotateFile(time.Now()); err != nil {
//		return nil, err
//	}
//
//	if r.RotateTime != 0 {
//		go r.handleEvent()
//	}
//
//	return r, nil
//}

//func (r *RotateLog) rotateFile(now time.Time) error {
//	if r.RotateTime != 0 {
//		nextRotateTime := r.nextRotateTime(now, r.RotateTime)
//		r.rotate = time.After(nextRotateTime)
//	}
//
//	latestLogPath := r.FilePath + r.FileName + r.getLatestLogPath(now)
//	r.mutex.Lock()
//	defer r.mutex.Unlock()
//	file, err := os.OpenFile(latestLogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
//	if err != nil {
//		return err
//	}
//	if r.file != nil {
//		r.file.Close()
//	}
//	r.file = file
//
//	if len(r.curLogLinkpath) > 0 {
//		os.Remove(r.curLogLinkpath)
//		os.Link(latestLogPath, r.curLogLinkpath)
//	}
//
//	if r.maxAge > 0 && len(r.deleteFileWildcard) > 0 { // at present
//		go r.deleteExpiredFile(now)
//	}
//
//	return nil
//}

func (l *Logger) Write(b []byte) (n int, err error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	length := int64(len(b))

	writeLen := int64(len(b))
	if writeLen > l.max() {
		return 0, fmt.Errorf(
			"write length %d exceeds maximum file size %d", writeLen, l.max(),
		)
	}

	if l.file == nil {
		if err = l.openExistingOrNew(len(p)); err != nil {
			return 0, err
		}
	}

	if l.size+writeLen > l.max() {
		if err := l.rotate(); err != nil {
			return 0, err
		}
	}

	n, err = l.file.Write(p)
	l.size += int64(n)

	return n, err
}

func (r *RotateLog) Close() error {
	r.close <- struct{}{}
	return r.file.Close()
}

//func (r *RotateLog) handleEvent() {
//	for {
//		select {
//		case <-r.close:
//			return
//		case now := <-r.rotate:
//			r.rotateFile(now)
//		}
//	}
//}

//	func (r *RotateLog) nextRotateTime(now time.Time, duration time.Duration) time.Duration {
//	   nowUnixNao := now.UnixNano()
//	   NanoSecond := duration.Nanoseconds()
//	   nextRotateTime := NanoSecond - (nowUnixNao % NanoSecond)
//	   return time.Duration(nextRotateTime)
//	}
func (r *RotateLog) getLatestLogPath(t time.Time) string {
	return t.Format(r.FilePath)
}

func (l *Logger) openExistingOrNew(writeLen int) error {
	l.mill()

	filename := l.filename()
	info, err := osStat(filename)
	if os.IsNotExist(err) {
		return l.createFile()
	}
	if err != nil {
		return fmt.Errorf("error getting log file info: %s", err)
	}

	if info.Size()+int64(writeLen) >= l.max() {
		return l.rotate()
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		// if we fail to open the old log file for some reason, just ignore
		// it and open a new log file.
		return l.openNew()
	}
	l.file = file
	l.size = info.Size()
	return nil
}

// 创建新文件
func (r *RotateLog) createFile() error {
	err := os.MkdirAll(r.dir(), 0755)
	if err != nil {
		return fmt.Errorf("can't make directories for new logfile: %s", err)
	}

	name := r.filename()
	mode := os.FileMode(0600)
	info, err := osStat(name)
	if err == nil {
		mode = info.Mode()
		newName := r.backupName(name, r.LocalTime)
		if err := os.Rename(name, newName); err != nil {
			return fmt.Errorf("can't rename log file: %s", err)
		}
	}

	return r.openFile(name, mode)
}

//func (r *RotateLog) openFile(name string, mode os.FileMode) error {
//	f, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, mode)
//	if err != nil {
//		return fmt.Errorf("can't open new logfile: %s", err)
//	}
//	r.file = f
//	r.size = 0
//	return nil
//}

// 获取文件名
func (r *RotateLog) filename() string {
	if r.FileName != "" {
		return r.FileName
	}
	name := filepath.Base(os.Args[0]) + defaultSuffix
	return filepath.Join(os.TempDir(), name)
}
func (r *RotateLog) dir() string {
	return filepath.Dir(r.filename())
}
