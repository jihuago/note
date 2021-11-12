package main

import (
	"bytes"
	"errors"
	"io/fs"
	"time"
)

// 实现fs.File 和 fs.FileInfo
// 实现文件系统先实现这两个接口
type file struct {
	name string
	context *bytes.Buffer
	modTime time.Time
	closed bool
}

func (f *file) Read(p []byte) (int, error) {
	if f.closed {
		return 0, errors.New("file closed")
	}

	return f.context.Read(p)
}

func (f *file) Stat() (fs.FileInfo, error) {
	if f.closed {
		return nil, errors.New("file closed")
	}

	return f, nil
}

func (f *file) Close() error {
	f.closed = true
	return nil
}

// 实现fs.FileInfo
func (f *file) Name() string {
	return f.name
}

func (f *file) Size() int64 {
	return int64(f.context.Len())
}

func (f *file) Mode() fs.FileMode {
	// 固定为 0444
	return 0444
}

func (f *file) ModTime() time.Time {
	return f.modTime
}

// IsDir 目前未实现目录功能
func (f *file) IsDir() bool {
	return false
}

func (f *file) Sys() interface{} {
	return nil
}
