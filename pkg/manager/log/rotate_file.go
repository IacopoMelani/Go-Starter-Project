package log

import (
	"os"
	"path/filepath"
	"sync"
	"time"
)

const rotateFileDateForma = "2006-01-02"

// Now - time.Now
var Now = time.Now

// RotateFile - Define a wrapper for log file
type RotateFile struct {
	mu          sync.Mutex
	autoRotate  bool
	filename    string
	flag        int
	perm        os.FileMode
	file        *os.File
	currentDate string // 2006-01-02 format
}

// NewRotateFile - Return a new Logging file instance
func NewRotateFile(filename string, flag int, perm os.FileMode, autoRotate bool) (*RotateFile, error) {

	r := &RotateFile{
		filename:   filename,
		flag:       flag,
		perm:       perm,
		autoRotate: autoRotate,
	}

	if err := r.openFile(); err != nil {
		return nil, err
	}

	fileInfo, err := r.file.Stat()
	if err != nil {
		return nil, err
	}

	r.currentDate = fileInfo.ModTime().Format(rotateFileDateForma)

	return r, nil
}

// Rotate - Rotates the file
func (r *RotateFile) Rotate() error {

	r.mu.Lock()
	defer r.mu.Unlock()

	if r.file != nil {
		r.closeFile()
		defer func() {
			if r.file == nil {
				r.openFile()
			}
		}()
	}

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	fullFilename := wd + "/" + r.filename
	newFilename := wd + "/" + filepath.Dir(r.filename) + "/" + nowDefaultFormat() + "-" + filepath.Base(r.filename)

	_, err = os.Stat(wd + "/" + r.filename)
	if err == nil {
		err = os.Rename(fullFilename, newFilename)
		if err != nil {
			return err
		}
	}

	r.currentDate = nowDefaultFormat()

	return nil
}

// Write - Implements io.Writer
func (r *RotateFile) Write(p []byte) (n int, err error) {

	if r.autoRotate && r.isToRotate() {
		if err := r.Rotate(); err != nil {
			return 0, err
		}
	}

	return r.writeRaw(p)
}

// closeFile - Closes the current file
func (r *RotateFile) closeFile() error {

	err := r.file.Close()
	if err != nil {
		return err
	}

	r.file = nil

	return nil
}

// nowDefaultFormat - Returns formatted now time(Y-m-d)
func nowDefaultFormat() string {
	return Now().Format(rotateFileDateForma)
}

// isToRotate - Returns if it's time to rotate the file :)
func (r *RotateFile) isToRotate() bool {
	return time.Now().Format(rotateFileDateForma) != r.currentDate
}

// openFile - Opens the current file
func (r *RotateFile) openFile() error {

	file, err := os.OpenFile(r.filename, r.flag, r.perm)
	if err != nil {
		return err
	}

	r.file = file

	return nil
}

// writeRaw - Writes bytes to file, opens if not
func (r *RotateFile) writeRaw(p []byte) (n int, err error) {

	r.mu.Lock()
	defer r.mu.Unlock()

	if r.file == nil {
		if err := r.openFile(); err != nil {
			return 0, err
		}
	}

	return r.file.Write(p)
}
