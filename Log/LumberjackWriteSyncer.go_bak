package Log

//
//import (
//	"bytes"
//	"github.com/natefinch/lumberjack"
//	"time"
//)
//
//type lumberjackWriteSyncer struct {
//	*lumberjack.Logger
//	buf       *bytes.Buffer
//	logChan   chan []byte
//	closeChan chan interface{}
//	maxSize   int
//}
//
//func newLumberjackWriteSyncer(l *lumberjack.Logger) *lumberjackWriteSyncer {
//	ws := &lumberjackWriteSyncer{
//		Logger:    l,
//		buf:       bytes.NewBuffer([]byte{}),
//		logChan:   make(chan []byte, 5000),
//		closeChan: make(chan interface{}),
//		maxSize:   1024,
//	}
//	go ws.run()
//	return ws
//}
//
//func (l *lumberjackWriteSyncer) run() {
//	ticker := time.NewTicker(1 * time.Second)
//	for {
//		select {
//		case <-ticker.C:
//			if l.buf.Len() > 0 {
//				l.sync()
//			}
//		case bs := <-l.logChan:
//			_, err := l.buf.Write(bs)
//			if err != nil {
//				continue
//			}
//			if l.buf.Len() > l.maxSize {
//				l.sync()
//			}
//		case <-l.closeChan:
//			l.sync()
//			return
//		}
//	}
//}
//
//func (l *lumberjackWriteSyncer) Stop() {
//	close(l.closeChan)
//}
//
//func (l *lumberjackWriteSyncer) Write(bs []byte) (int, error) {
//	b := make([]byte, len(bs))
//	for i, c := range bs {
//		b[i] = c
//	}
//	l.logChan <- b
//	return 0, nil
//}
//
////func (l *lumberjackWriteSyncer) Sync() error {
////	return nil
////}
//
//func (l *lumberjackWriteSyncer) sync() error {
//	defer l.buf.Reset()
//	_, err := l.Logger.Write(l.buf.Bytes())
//	if err != nil {
//		return err
//	}
//	return nil
//}
