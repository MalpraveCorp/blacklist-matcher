// Code generated by go-bindata. (@generated) DO NOT EDIT.

// Package main generated by go-bindata.// sources:
// bindata.go
// blacklist.txt
// main.go
// main_test.go
package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func bindataGoBytes() ([]byte, error) {
	return bindataRead(
		_bindataGo,
		"bindata.go",
	)
}

func bindataGo() (*asset, error) {
	bytes, err := bindataGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bindata.go", size: 0, mode: os.FileMode(436), modTime: time.Unix(1601792725, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _blacklistTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x57\x3d\x92\xe4\xba\x0d\xce\x71\x0a\x46\xce\x1c\xbc\x57\xe5\x03\x41\x14\x5a\xc2\x88\x24\xf8\x00\xb0\x35\xda\xd8\x57\xb0\x53\x57\x39\x71\xee\x68\x8f\xb3\x47\xf0\x11\x5c\xa0\xba\x67\xb7\xec\x60\x80\x8f\x7f\x4d\x8a\xfc\xf0\x01\xf3\xfb\xf6\x5b\x86\xdf\xd3\xc6\x5a\x2c\xfd\x96\xf2\xe8\x80\x59\xc5\xa5\x4a\xdf\xb9\x30\x02\x16\x5c\xb0\x62\xda\xc5\x53\x97\x7c\x90\x47\x97\x1d\xd8\x52\xe7\x4e\x85\x1b\x01\x36\x2c\x80\x8d\x0b\xb7\x6d\x18\x60\x0b\xd3\xc9\x76\x76\x40\x35\xda\xa5\x10\xa0\x59\xfc\xbd\x71\x1d\x2d\xef\x80\xc3\x25\x91\x8a\x73\x9e\xf8\x05\x17\x5c\xa8\x60\x5b\x03\x5c\x69\x41\x77\xd2\x1b\x7f\x0c\xce\x04\x0b\x96\x92\x36\xdc\x5e\x40\xf1\x79\xdd\xf0\xe0\x7c\x70\x7b\xf5\x97\x5f\x1b\x86\xf9\x78\xa1\xf1\xee\x6e\xdb\xa2\x62\xb0\xa0\xd2\x72\x0f\x2b\x95\x2b\x15\xda\xb0\xcc\x46\xc3\x83\xe2\x10\xe6\xa8\x5f\x5e\x26\xe0\x86\x81\x96\x13\x96\xd5\x2a\x2c\x84\x2d\xce\x38\x9d\x85\x7f\x92\xa6\x5c\xa6\x7f\x37\x0b\xf7\x18\x32\x67\x2c\xec\x17\x2c\xbc\xa5\xa5\xcc\x9d\x03\x29\xa1\xb9\x4d\x7c\xb4\xb8\x6a\xbd\x1b\xce\xb3\xb7\x2e\x71\x56\xd6\xb5\xc8\x5c\xe1\x79\xbf\x2d\x19\xcc\x5f\x49\x79\x0e\x14\x69\x2b\x25\xcc\xce\xd2\xde\x2d\x69\xe9\xff\xfa\xcf\x0f\x59\xa6\x4f\x5f\xe0\x92\xa1\xa9\x08\xae\xb0\x94\x41\xe9\xc4\xc7\xa3\x50\xe0\xda\x0f\x6e\xb0\x48\x89\xbd\x0d\x16\x69\x2b\x6e\x14\x3e\xbe\x4e\x62\xbd\xc8\x12\x03\xe2\x57\xca\x58\x0a\x2c\x2a\x67\x4b\xb6\xcb\x39\xbf\x43\x47\x23\xf7\x9f\xfb\x8f\xe3\xc0\x83\x60\x19\xa5\xac\xd7\x0b\x90\xa7\x27\x2f\x37\x9e\xe4\x59\x46\xdb\xd2\x64\x4c\xa0\x17\xb0\xb8\xb9\xe1\x3e\x4d\xde\x89\xe2\x40\xc3\x7d\x0e\x67\xac\x54\x92\xcb\x44\x41\xeb\xf0\x56\x86\x87\x3f\x77\xd1\x18\xd0\x4e\x9e\x26\x01\x49\x5f\xcd\xaf\xd6\x2e\x59\x0a\x3a\x25\x15\xa3\x65\xac\x06\x99\x35\x17\xfa\x20\x3d\x20\x17\x7a\x4e\x62\x26\x73\xc2\x1a\xf3\x0b\xfb\x34\xa2\x6c\x90\x8b\xdc\xcf\x8e\xb5\x47\x6b\x98\x93\x3e\x46\x3e\x60\x3e\x4d\x9e\x97\x97\xa5\xab\x14\xdc\x1a\xe3\x8d\x5f\x91\x96\x45\xdb\xfd\x0d\x22\x6d\x1a\x83\xac\x84\xb5\x33\x41\x1e\x35\xfe\x6a\x10\x37\x8f\xf6\x15\x6b\x79\x34\x87\x15\xf5\x60\x82\x75\x1e\x1b\xfb\x8d\x6e\x40\xd4\x93\xef\x2a\xe8\x13\x7f\xc1\xb6\x7e\xed\xbb\x72\x3e\x60\xe5\xb2\x0a\xac\xdc\xb6\x42\x0b\xa9\x5e\xbf\x60\x26\x83\x95\xd5\xaf\xd4\xb9\x14\x39\xdf\x2d\xc3\xb8\xb3\x6f\xb0\xca\xb6\x31\x25\xf3\xab\xd0\xab\xf1\x13\x5f\xbf\xf4\x5f\x5f\xf0\xab\xb3\x64\x72\x87\x55\x2a\x37\x9c\xc4\x78\x41\xe5\xcf\x80\x35\xb6\x96\x76\xd0\x95\xfa\x14\x8c\x55\xc6\x52\x28\xad\xd2\xb6\x37\xee\xd4\xc8\xf5\xb5\xba\xbf\x09\xb6\xea\x95\xf6\x51\x3b\xac\xcf\x15\x81\xd0\x53\xbd\x52\x68\x10\xe5\xbc\x33\xd0\x07\xe6\x51\xee\x45\x2f\xd1\x99\xce\x2a\x90\x65\x51\x07\x1a\x6d\xe4\x1d\x1e\xb8\xc5\xdf\x26\x0e\x0f\xca\x58\xe0\x41\x25\xba\xa9\xcc\xd5\x01\x7c\xb6\x2b\x16\x4a\xf6\xc7\x60\xf5\x78\xa4\x07\xd5\x55\x2a\x3c\x78\xdb\x66\x93\xdb\x46\x1a\x92\xf3\x82\x77\xa7\xdd\x73\x45\x3c\x3d\xc8\xd9\xf6\x89\x23\x1e\x1f\x2a\x7e\x0f\x06\x7d\xc2\xa4\x20\x79\x90\xe2\x31\x05\xec\xe5\xee\x09\xa1\x4b\x31\xb0\x6e\x94\x3a\x86\x74\xdc\x8d\x2f\xec\xd8\x50\x19\x36\x6c\x5b\x9a\xa7\xd8\xf0\x4a\x46\x9f\xb0\x51\x63\xc7\x62\xb0\x31\x36\xbf\x25\x24\xe2\x26\x49\x7b\xfb\xe4\xd2\xe1\x4e\x11\x9b\x34\x4a\x27\x97\x15\x36\x41\xcf\x9f\xd3\x19\xc1\x26\x6b\x5a\xb1\x36\xd8\xe4\x38\x46\xb8\xb2\xd2\x3b\xfa\x61\x13\x59\xbb\xc4\x8f\x88\xcc\x5c\x03\x9b\x28\x6d\x68\x15\x36\x95\x4e\x61\x47\xbf\x8f\xf3\x67\xeb\xe2\xb0\x0d\x15\xd8\x23\xd0\xe2\x2e\x02\xdc\x5e\xd7\x94\x23\x86\x03\xdd\x80\x9a\x23\xc3\x2e\xf5\x9d\x3b\xf6\x49\x17\xd8\x45\xe2\xcb\x23\x6d\x65\xd4\x72\x83\x3d\xa8\xbe\xcb\x99\x5c\xd2\xc1\xa5\xbc\x71\x1d\xba\xc6\xe4\xb1\x51\x7a\xa0\x43\x10\x27\xae\x96\x5b\x26\x73\xe0\xe6\xa4\x59\x86\x1a\xc1\x47\x08\xad\x3c\x1e\xf0\x81\x5c\xd2\x82\xec\x13\xdd\x80\x4a\xb9\x82\x9b\x23\xb0\xbe\xe6\xf1\x86\x8b\x48\xf8\x9f\x80\xf4\x46\xdf\xbe\xc1\xc7\xd8\x36\x83\x83\x0f\x82\x83\xdb\x82\xc7\x08\x7f\x84\x72\x4c\x70\xc1\xd1\x64\x59\xe2\x38\x85\xd0\x77\xd2\xa4\x64\xae\xc8\xcd\xbf\x7a\x66\x7b\xdb\x3d\xc5\xf1\x28\xfa\xab\xb4\xd4\x51\xfd\x82\x22\x85\x1d\x21\xa4\xa9\xe2\x64\x4c\xc5\x83\x52\xa5\x94\xa5\x12\xfc\x0f\x6f\x2b\x9a\x0f\x5d\xd0\x09\x2a\x35\xdc\x28\x61\x72\x15\x36\xa8\x5c\x1e\x50\xd9\x8c\xa5\xa1\x5e\xa9\x8b\xf1\x8c\x9f\x2a\x71\x86\xa0\x21\x29\x54\x19\x6d\x4d\xf2\x48\x4f\x8a\x32\xa0\x6a\x8a\xd7\x33\xa8\xe3\xf1\x48\x2b\x47\x36\x0c\xb8\xf2\x33\x76\x6b\x58\x97\x82\xd0\xf0\x44\xdb\x19\x1a\x6d\x2a\xd0\x48\x1a\x7e\x63\x68\x71\x63\xd3\x92\x86\x4b\x4d\x36\x68\x5c\xfb\x2e\x15\x43\x3d\x1b\xf7\x5e\xe8\xe5\x0c\x9a\x3d\xce\xc4\x15\xb7\xc0\x63\xa5\x30\x91\x66\xdb\x15\x4b\x5e\xee\x5e\x29\xd9\xa5\x0f\xb3\x0b\xa4\x8a\xce\xbd\x83\xd9\x79\xf4\xe4\xe7\xcd\x51\x9b\x3d\xdb\xb8\x52\xf8\x0f\x54\x10\x9d\x94\x15\xdd\x2e\xe8\x48\xeb\x14\x50\x82\x8e\x07\x43\xc7\xe6\xa1\x93\xe1\x2f\xe8\xb4\xca\x42\xa8\x13\xbc\x66\xd1\x2d\x04\x9d\x1a\x1b\xf4\x3d\x7e\x34\x08\xdf\x99\x32\xc5\x7d\xcd\x84\xd7\xe3\x7a\x63\x16\x9b\xa5\xce\x37\x98\xbe\xe0\xb5\xc8\x05\xbd\x10\xda\x50\x4a\x91\xf6\x1d\xba\xc4\xe3\xd5\x49\xf5\x2e\xed\x8a\x69\xd0\x45\x1e\x61\xda\x34\x1e\xd1\xde\x47\xc3\x36\x47\x7a\xca\xfb\x70\x9a\xf0\x8d\xb4\x4d\x23\xb7\xdd\x14\xfb\x7e\x41\xd7\x60\x7f\xc2\xb2\x90\x7a\xea\x4c\x9a\xe7\xc9\x7c\xcf\xd0\xc7\x12\xdf\x3a\x2f\xf0\x8f\x41\xf8\x08\x4b\x61\xb9\x82\xe2\xb6\x13\xae\xe1\x39\xd4\x66\xd6\x08\x33\x1b\x29\xce\xa8\x0a\x67\x0e\x4a\xd9\x47\x05\xa5\x27\xa9\x05\x19\xcf\x29\x0d\xca\x35\x62\x5d\xf9\x4e\x76\x2a\x76\xa5\x8e\xa5\xfe\x44\x29\x74\x21\x58\xff\x97\x64\x1c\x71\x62\xa0\x51\x15\x04\x4f\x6b\x6c\x07\x86\x6b\x28\xb9\x61\x73\xd1\x51\xc1\x32\x3a\x58\xde\x4b\xa4\x0d\xcb\x6c\x26\x53\x80\x8d\x2a\x35\x88\x67\x30\xfa\x94\x30\x17\xd8\x8e\x4f\x5a\xd3\x5d\xb2\xbd\x5b\xf7\xa7\xda\x3e\x55\x1e\x6c\xe7\x25\xd4\x74\x3e\x59\x98\xa5\x70\xed\x13\x79\xcc\x12\x47\xb0\x5d\xf9\x16\x11\x3b\x88\x1c\xac\x60\x73\xba\x08\x66\x31\x62\x7f\xaa\x60\x35\x40\xc3\x48\x1f\xd6\xe4\x8c\xe2\x74\xce\x97\xc8\x81\xdf\xe8\x06\x17\x58\xe7\x0c\xd6\x8b\xc8\x46\x6f\x9f\xaa\x88\x45\xeb\xd5\xa9\x84\x6b\xd4\xad\x06\xd6\x47\x3b\x20\xf4\xa0\x87\x84\x4f\xf0\xf6\x3d\x4a\x56\x73\xe5\x9e\x72\x19\x0b\xcc\x24\x9c\x66\x66\x86\xa8\x8a\xa7\x31\xb0\xc1\x99\x57\x7a\x45\x81\x8d\xe2\x7a\xa5\x53\xe6\x55\x9d\x51\xfa\x1e\x08\x76\xce\x2c\x06\x1e\x32\x44\x6b\x0a\x75\x01\x47\x73\x4a\xf5\x02\x27\x4c\x0b\xde\x84\xf7\x5d\x89\x2c\x94\xe6\x2e\x3e\x66\x1f\xd3\x9a\x46\x07\x9f\x9a\x75\xee\xec\x04\xce\x0e\xb3\xd2\x75\xf6\x19\x4a\x3e\x6f\xd3\xa5\x6d\x83\x12\xb7\x84\xe0\x12\x71\x6e\xe0\x62\x16\x7b\xcb\x49\x65\x52\xcd\x15\x5b\xbb\xc0\x95\x97\xfb\xe5\x7d\x2c\x77\xa2\xf1\xb1\xbc\xbc\xed\x17\xf8\x89\x0e\x7e\x72\x3b\x6e\xcb\x04\x5f\xe1\x9e\x5e\x02\x00\xa3\xad\x4a\x77\x18\x8e\x6e\x07\xab\xc3\x50\xf2\x5d\x31\xcd\xf0\x1a\x5f\x95\xd3\x33\x18\x8e\x30\x95\x2e\x4d\xe1\x83\x27\x2f\x8a\x2e\x0a\x4f\x96\xa8\x67\xcf\xf8\x2f\xe6\x29\x8a\x4a\xef\x45\x72\xd1\x50\x78\x8e\xf2\x44\x38\xb1\x1d\x70\x92\xcf\x7f\x3f\x4e\xf2\xb4\x46\xc1\x07\xf3\x46\x52\x9f\xc9\xf3\x8c\x87\x8b\x40\x8a\xfb\x3f\x95\xdb\x51\x28\xea\x4f\xd4\x47\xd4\x0b\x9f\x9f\xf0\xf9\xf9\x09\x17\x0a\xc3\x45\x51\x9e\x7d\xd5\xdc\x17\x3f\x1e\x17\x7c\x93\xf7\x71\xff\xf3\x8f\xbf\xff\x0d\x7e\x7c\xff\xeb\x8f\xef\xff\xfa\xf1\xfd\xdf\x3f\xbe\xff\xf3\xbf\x01\x00\x00\xff\xff\x0c\x09\xc1\x1b\xfe\x0d\x00\x00")

func blacklistTxtBytes() ([]byte, error) {
	return bindataRead(
		_blacklistTxt,
		"blacklist.txt",
	)
}

func blacklistTxt() (*asset, error) {
	bytes, err := blacklistTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "blacklist.txt", size: 3582, mode: os.FileMode(436), modTime: time.Unix(1601792581, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mainGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x55\x51\x6b\xe4\x36\x10\x7e\xb6\x7e\xc5\xd4\x50\xb0\xc1\x67\x6f\x4b\xfb\x12\x48\x21\x09\xb4\x0d\xb4\x25\x74\x0b\x7d\xb8\x1e\x87\x6c\x8d\xbd\x22\xb2\xe4\x93\xc6\x9b\x5d\x8e\xfd\xef\x65\x24\x3b\xd9\xeb\xa6\xa4\x34\x90\xb5\x3d\x33\x9a\xf9\xbe\xd1\xe8\xd3\x24\xbb\x47\x39\x20\x8c\x52\x5b\x21\xf4\x38\x39\x4f\x50\x88\x2c\xef\x9c\x25\x3c\x50\x2e\xb2\xbc\x1f\xe3\x23\x90\xd7\x76\x08\xb9\x10\x59\x3e\x38\x23\xed\x50\x3b\x3f\x34\x87\x86\xe3\x1a\xfe\x9e\xe5\x80\xf9\xab\xde\x80\xd2\x77\xbb\xb4\x54\xd3\x6e\x6e\xeb\xce\x8d\x8d\x7c\x0a\xfc\xff\xce\xc8\xb1\x55\xf2\xdd\xe0\x1a\xdc\xa3\xa5\x90\xbf\x15\x96\xde\x72\x91\x35\x0d\x9c\x47\x76\xc6\xcd\xaa\x37\xd2\x63\x23\x77\xae\x73\x5e\x06\xdd\x3d\xe6\xa2\x14\x82\x8e\x13\xc2\xef\x18\x26\x67\x03\x42\xaa\x53\xdf\x3c\xdc\xff\x24\x09\x9f\xe4\xf1\xc1\xbb\xc3\x71\x75\xaf\xc1\x9f\x66\x0c\xf4\xaf\xb1\xd1\x2b\x44\x3f\xdb\x0e\x7e\x96\x56\x19\xf4\x45\x47\x07\x58\x3a\x57\xdf\xa5\x67\x05\x7e\x49\xb4\x2c\x29\xa1\x58\x0b\x55\x80\xde\x3b\x5f\xc2\x67\x21\xb2\xd6\xc8\xee\xd1\xe8\x40\xd1\x0a\x57\xd7\x60\x9c\x54\x7f\x3a\xaf\x6e\x57\x4f\x51\x8a\x4c\xf7\xd1\xfd\xd5\x35\x58\x6d\xe0\xb3\xc8\x32\x8f\x34\x7b\xfb\x4c\x8e\x4d\xd9\x96\x24\xcd\xe1\xce\x29\xbc\x82\xef\x37\x9b\x8a\x6d\xb7\x4e\x1d\xaf\x20\xfd\xf5\x23\xd5\xdb\xc9\x6b\x4b\x7d\x91\xdf\x49\x6b\x1d\xc5\x72\xf0\x8c\x02\x94\x24\xd9\xca\x80\x79\xc9\xab\x4f\x15\xd7\x13\xd9\x49\x88\xec\xc9\x79\x55\x41\xef\x66\xab\x18\x66\xaf\xed\x0b\x44\xc6\x5b\x9c\x51\x59\xc8\xd7\x5c\xbc\x14\x71\xc7\x16\xbc\xce\x42\xaf\x7d\x20\x18\x25\x75\xbb\x48\x2c\xa5\xbc\xbe\x06\xf2\x33\xfe\x27\x6e\xdf\xbd\xc5\x6d\x3b\xb7\xa3\x0e\x41\x3b\x0b\x3a\x80\xb6\x7b\x69\xb4\xaa\xe1\xc7\x58\x29\xe2\x04\x06\x8a\x0a\x1e\xf1\xc8\xc4\xae\xe0\xeb\x90\x57\xc0\xaf\x17\xc4\x5f\x41\x73\x0e\xe6\xdb\x04\xe6\x1c\xcb\x3f\xea\xc7\xea\x79\x25\xd6\xac\xe2\x24\x44\xd3\xc0\x8e\x68\x0a\x57\x4d\xf3\xe6\x28\x73\x70\x98\xa7\xc9\x05\x54\xe6\x08\x2d\x12\xa1\x07\xb9\xf4\x50\xdb\x01\x8c\xf4\x03\x82\x9d\xc7\x16\x3d\xb8\x1e\x26\xc9\x21\x36\x70\x90\xb3\x1d\x72\x86\x76\x26\x98\xd0\xf7\xce\x8f\xd2\x76\x08\x01\x71\x04\x72\xd0\x22\xe0\xa7\x59\xef\xd1\xa0\x25\x36\x04\x92\x56\x49\xaf\xc0\xe8\xd6\x4b\x7f\x84\xde\x79\x4e\x40\x18\x28\xc0\x93\xa6\x1d\x7c\xb3\xd9\x04\xae\xb3\x74\x2f\x80\xb4\x0a\x24\xb4\x4e\x1d\xd9\xcc\x27\x00\x66\xab\xd0\x73\xe4\x26\xb6\x35\xd4\xe9\xd0\x5c\xce\x4d\xca\xf0\xfe\x43\x92\x9a\x0a\x82\xef\x20\xbd\x97\x10\x9d\xb0\x7a\xd2\xa4\xb4\xce\x19\x3e\x3c\xd9\xc8\x73\x98\x44\xa6\xfe\x0d\x9f\x8a\x55\x8d\xea\x3f\x76\x52\x57\xab\xe7\x7e\xb0\xce\xe3\x9d\x0c\x58\x8a\xac\x77\x1e\x3e\x56\xb0\xe7\x95\x5e\xda\x01\x13\xb6\x38\x75\xba\x82\x8f\x6c\x1f\xeb\x7b\xab\xf0\xb0\x8d\x45\x8b\xe0\xbb\x0a\xf6\x25\xfb\x7b\xd0\xf0\x03\x6c\x62\xf0\x3a\x15\xfb\x2a\x8e\x2d\x8f\x0c\x4f\xcb\x6a\xce\xf3\x0a\x7a\x69\x02\x2e\x7b\xfd\x25\xf5\x5f\x16\xea\x37\x2f\x7b\xfc\xbf\xbb\xd0\x34\x90\x1a\x71\x36\x30\xdc\x8d\x04\xff\x57\x9e\x10\xf4\x29\x7b\x99\x82\xa3\x29\x24\xa6\xd1\x5f\xbc\xff\xd0\x1e\x09\x99\x6a\x99\x62\xf8\x30\x3d\xf0\x59\x32\xb6\x88\x54\xb8\x64\x72\xe9\xbe\x30\x68\x8b\x25\x4b\xc9\x0d\x59\x61\xac\xe4\x17\xdf\xd2\x19\xf6\x9c\xe2\xef\x45\x6f\x9a\x06\x4e\x8b\x96\xbe\xa2\x7a\x70\xd1\x13\x96\xc0\x67\xed\xcc\x58\xaa\x9e\x55\xf3\x26\x04\xa4\x22\x7f\xd6\xa0\x9a\x0e\x94\xbf\x2a\x9c\x5f\x70\xbb\xbd\x10\x3e\xe8\xdc\x6c\x14\xb0\x34\xb6\x18\x61\xa1\xe2\x44\x2f\x7b\xbb\x5c\x89\xf5\x76\x32\x9a\x8a\xf4\x55\xf0\xf2\xb2\x82\xfc\x2f\x9b\x97\xe9\x94\xaf\xcc\xf8\x92\x2d\x22\xe0\x74\x81\xd5\x5b\x92\x9e\x8a\xe5\xf2\x28\xc5\x49\xfc\x1d\x00\x00\xff\xff\x8f\xf0\xbc\xbd\x8f\x07\x00\x00")

func mainGoBytes() ([]byte, error) {
	return bindataRead(
		_mainGo,
		"main.go",
	)
}

func mainGo() (*asset, error) {
	bytes, err := mainGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "main.go", size: 1935, mode: os.FileMode(420), modTime: time.Unix(1601792200, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _main_testGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x57\x5b\x6f\xdb\xc8\x15\x7e\x96\x7e\xc5\xa9\x9f\xe2\x42\x91\x17\xdb\x16\x58\x18\xc8\xc3\x6e\x9b\x76\x5b\xf4\x61\xd1\xa4\x4f\x45\xb1\x18\x93\x63\x79\x10\x8a\x54\xc8\xa1\x2f\x08\x02\x98\x64\xb2\x51\xe2\xa4\x6b\xc7\x56\x6c\xd9\xb2\x2d\xcb\xb6\x2c\xc9\xb1\x11\xdf\x36\x54\x2e\x9b\x9f\x31\xaf\x04\xc6\x94\xe5\x7f\x51\x9c\xa1\xe4\x75\x2e\xd8\xa2\x45\xb1\x40\x1c\x91\x73\xbe\xf3\x9d\xef\x7c\x87\x97\x61\x81\x68\xb7\x48\x8e\x42\x9e\x30\xf3\x5b\x4e\x1d\x9e\x4e\xb3\x7c\xc1\xb2\x39\x5c\x49\xa7\x06\x34\xcb\xe4\x74\x92\x0f\xa4\x53\x03\x39\xc6\xc7\xdc\x91\xac\x66\xe5\x87\xc8\x84\x83\x7f\x57\x0d\x92\x1f\xd1\xc9\xd5\x9c\x35\x44\xc7\xa9\xc9\x9d\x0f\x60\x0e\xb7\x29\xd7\xc6\xec\x21\xa4\x65\xa3\x53\x43\xc4\x71\xa8\xad\xc8\x9c\x09\x4a\xec\xab\x8e\x46\xcc\xa1\x11\x83\x68\xb7\x0c\xe6\xf0\xab\x79\xc2\xb5\x31\x6a\x63\x5c\x65\x98\xb9\x81\xf4\x60\x3a\xcd\xa7\x0a\x14\xfe\x46\x6f\xbb\xd4\xe1\x90\x14\xca\x7e\xf9\xcd\x9f\xff\x44\x38\x9d\x20\x53\xdf\xd8\xd6\xe4\x54\x2f\x9a\x4e\x8f\x13\x1b\x1c\x32\x4a\x6f\xd2\x49\x7e\xdd\xcc\x19\xcc\x19\x83\x6b\x30\xf0\x57\xcb\xa6\x79\x60\x05\xc7\xcd\x83\x6e\x19\x96\x0d\x0e\xe3\x40\xf2\x94\x67\x40\xb3\x4c\x87\x6a\x9c\x72\xd7\x06\xa2\xb3\x02\x73\x34\x66\xe6\x80\x1a\x8c\x67\xc0\xa1\x3a\xe8\x16\x50\xe6\x3a\x79\x4b\x07\x4e\xd1\x1a\x60\xa6\xc6\x74\xa6\xbb\x26\x07\x97\x83\x41\x46\x2c\x9b\x02\xe5\x09\x35\x5a\x99\x33\x09\x10\x83\xdd\x76\x49\x16\xfe\xce\x81\x9a\x2c\x0f\x44\x87\x3c\xc3\x83\x71\x6a\x32\x92\xcf\xc0\x6d\x97\x39\x60\x5a\x0e\xb7\x5d\x1d\xe8\x24\xb5\x35\xc6\x09\x67\x96\x09\xae\x61\x90\xbc\x66\x25\xcc\x08\x62\x0e\xc3\x4a\x8a\x92\x15\x80\x4e\x02\x25\xa0\x59\xf9\xbc\xa5\x5b\x49\x03\xb7\x5d\xc2\xb3\xf0\x07\xa4\x24\x2e\xa7\xc0\x6c\xd7\xa6\xbd\x5e\x99\x09\x36\x2d\xd8\x74\x8c\x9a\x3a\xb5\x19\xc7\x85\x71\xcb\x70\x0b\x9c\x70\x0a\xe3\xd8\x29\x50\xc7\xa1\xa0\x31\xc3\xe8\x3b\x44\x81\xba\x30\xea\xe6\x18\xe1\x60\xa2\x20\x28\x10\x9b\x11\xee\xda\x59\xb8\x3e\xa9\xd1\x02\xa7\x2e\xda\x68\x72\xb0\x34\x8d\x50\x8d\x70\xd0\xdc\x02\xd3\x09\xc7\x0c\xcb\x84\x82\x6d\x31\x9d\x9a\xe8\x22\x3a\xc5\x4c\xd0\x5c\xa3\x40\xb0\x6f\xb0\x46\x47\x99\xc6\x08\xe8\xd4\xa1\x36\x46\xf3\x96\x81\x32\x08\x1a\xc4\x74\xc0\x51\xab\xee\xdd\x7c\x76\xe0\xfd\xb1\xfe\x85\x14\x2c\x93\x3a\x14\xe7\x1a\xef\x2f\x4b\xff\xdd\xd9\x7a\x2d\xbe\x7f\x20\xbd\x7d\xe9\x07\x32\x28\xc9\xe0\x79\x77\xef\x45\x5c\x3d\x96\xde\xf6\xe9\xab\x99\xb3\xdd\x6d\x19\x34\x64\x50\x8d\xdb\xa5\xf3\xe9\x75\xe9\x55\x3a\xc5\x7a\xbc\xf5\x4c\x7a\xff\x92\x7e\x51\x7a\xf5\xce\xf4\xea\xd9\xca\xb1\xf4\xdf\x4a\xff\x48\xfa\xad\xee\xbb\xb9\x6e\xf3\x40\x06\x6b\x32\xa8\xc9\xa0\x25\x83\xe6\xd9\xee\x4c\x67\xab\xd2\xad\x1f\x7d\xfe\xd9\xd9\xd6\x2b\xe9\x1d\xc4\x8f\x8e\x3b\x8b\xdb\xbf\x8b\x9f\x54\x4f\xc3\x50\xfa\x2d\xe9\xb7\xe3\xfd\xe5\xb3\x93\x3a\x1e\x07\xcf\xe3\xfb\xc5\xf8\x5d\x55\x7a\x0b\xd2\x9f\x96\xc1\x73\xe9\x1d\x75\x67\xea\xdd\xfa\x91\xf4\xb6\xa5\x37\x2b\xbd\x0d\xe9\x35\xe3\xfd\x07\x9d\xf5\xf9\x78\xee\xed\xb9\x17\x4a\xaf\x29\xbd\x2d\xe9\x55\xe5\xb4\x1f\x87\xe1\xd9\xc9\x63\xe9\x85\x71\xa5\xd1\x79\x34\x2d\xfd\xd7\x32\xb8\x27\x83\x25\xe9\x1f\x75\x0e\x5f\x77\x1e\x7b\xd2\x7f\x27\xfd\x46\xbc\x57\xef\xbc\x59\x90\xde\x8f\xdd\xfa\xb3\x4e\xe5\xa1\x0c\x8a\x32\xf8\x3e\x7e\xfc\xea\xf4\xf5\x92\xf4\xd6\xe3\xad\xe5\xb8\x7a\x20\xbd\x79\xd4\xef\xed\x4a\x7f\xa6\xb3\x51\x8c\x9f\x1c\xff\xe6\xbc\xbc\x25\xfd\xe2\xf9\xee\x52\x77\xf3\x6d\x27\x2c\x9d\x86\xd3\xd2\xfb\x11\xc3\x65\xff\xf4\xdd\x66\x5c\x0f\xa4\x7f\x4f\x7a\x33\x32\x08\xe4\xb4\xdf\xd9\x9c\x91\x5e\xe3\x7c\xe7\x7e\x3c\x5b\x47\x22\x74\xe7\xe4\x6c\x65\xff\xbc\xba\x24\xfd\x97\xd2\xdf\x95\xc1\x53\x19\x14\xcf\x57\x6b\xf1\x5c\x43\x7a\x73\xd8\xb0\xf7\xa6\x53\x29\x7e\xd1\x09\x9f\xa3\xe4\x60\x5d\x06\xf5\xee\xda\x93\xf3\x67\x0f\x65\xb0\xd4\x3d\x7a\x73\xbe\xbb\xa4\x72\x9e\x9c\x86\x4b\xe7\xc1\x89\x0c\x1e\xc9\xe0\xbb\xee\xf4\xfd\xcf\xcf\x56\xf6\xa5\x57\x96\xde\xba\xf4\xbf\x93\xfe\xcc\x59\x65\xbd\xdb\xda\x93\x5e\x49\x7a\xeb\xdd\xbd\x97\x9d\x4a\x51\x7a\x2b\xd2\xf7\x3a\x33\x1b\x71\xb1\xd8\x99\x79\x2a\xbd\x13\xe9\x7d\xff\xc1\x85\xf1\x35\x1d\xb1\xe9\x04\x5e\x16\x62\x56\xec\x88\x86\x98\x15\x65\x10\x4f\x45\x59\xcc\x8a\x92\x78\x2a\x1a\x62\x4b\x94\xc5\x02\x88\x39\x10\x73\x08\xc8\xe0\xef\x82\x68\x8a\x45\xd1\xc2\xf5\x9a\xca\x2a\x89\xaa\x28\x89\x16\x88\xa6\xa8\x20\x60\x5e\x6c\x66\x15\x10\x44\x59\xac\x88\x12\xae\x22\x6a\x43\x94\x11\x85\x29\xeb\x59\x10\xcf\xc4\x8a\x68\x80\x58\x14\x25\xb1\x25\x9a\x49\xac\x22\x4a\xa2\x2c\x76\x44\x59\x6c\x89\xf9\xa4\xf6\x9a\x58\xc8\x82\x98\x15\x55\x31\xab\x84\xac\x89\x59\xb1\x86\x89\x15\xb1\xa6\x0a\xd7\x40\x65\xaf\x89\x92\xfa\x2d\x8b\xd5\x0c\xa2\xe6\x31\xd0\x12\x25\xd1\x40\xe2\x05\xb1\x83\x0c\x18\xc5\x93\x59\xb1\xa4\x12\x36\x54\x87\x0d\xac\x33\x2f\x5a\xaa\x4e\x09\xff\x6b\x88\x6d\x84\x95\x45\x09\xf1\xd8\x7b\x25\x69\x71\x41\xd4\xc4\x9c\x68\xa0\xda\x4c\x0f\x3d\x27\x9a\x62\x0b\x19\x90\xaa\xa2\x48\x4b\x7d\xe3\xd6\x44\x4b\xac\x67\x40\xac\x88\x56\x62\xe2\x02\x26\x2d\x2a\x82\x55\xe8\x49\xc0\x86\x93\x1e\x57\x11\x59\x15\x65\xb1\xa1\x5c\x57\x56\xcd\x61\xd9\x92\x68\x89\xa6\x28\x89\xb9\x44\xc5\x22\x3a\x96\xc1\xbe\xab\xa2\x85\x33\x7b\x8f\x73\x5e\x1d\xae\x24\xa7\xbd\xca\x4d\x51\xc9\xa2\x1d\x55\xb1\x0c\xca\x06\x1c\x48\x2d\x81\xcc\xa9\x3e\x2b\x6a\xe4\xe5\x44\x61\x29\x93\xb4\x3f\x8b\x7c\x5b\xa2\x25\x16\x41\x55\x57\xfd\x60\xc6\x62\x8f\xa2\xa5\x64\xf7\x9f\x38\xae\xf9\x89\x57\xc9\x97\xe0\x70\x9b\x15\x40\x33\xdc\x11\xc0\xa7\x2d\x3e\xcf\x5d\x0a\x13\x63\xd4\xa6\x49\xac\x40\x6d\x07\x9f\x7e\xe3\x4c\xa7\x40\x74\xd7\xc0\x37\x00\xa7\x36\x27\xcc\xcc\xab\x07\x62\xc1\xa6\xba\x95\x67\x26\x31\xb9\x31\x85\x4f\x46\x3e\x46\x61\xd4\xb2\xf3\x60\x8d\x26\x1c\x9c\x12\x87\x82\x65\x83\xc5\xc7\xa8\x0d\xd4\xb6\x38\xd3\xf0\x9c\x4e\xaa\x23\x9d\x98\x1a\x75\xb2\x70\xe3\x42\x8c\x03\x7c\xaa\xc0\x34\x62\x18\x53\x40\x74\xab\xc0\x81\x80\xc9\x72\x63\x5c\x29\xb5\x6c\x18\xc1\x7b\x85\x4f\x19\x34\x03\xc4\xd4\x41\x23\x26\x10\xc3\xb1\x2e\xc0\x7c\x8c\x12\x6e\xab\xa2\x1a\x19\x21\x36\xe5\x57\x15\xfc\x63\x3b\x6e\x8e\x11\x86\x5e\x44\x61\x25\x0a\xdb\x51\xfb\x61\x14\xee\x45\xe1\x51\x14\x96\xa3\x70\x27\x0a\x6b\x51\xe8\x45\xe1\x61\x14\x6e\x46\x61\xab\xbf\x7e\xa0\x16\x97\xd5\xe2\x71\x14\xee\xf6\x01\x6b\x51\xf8\x52\x31\x94\xa3\x70\x31\x0a\x4f\xa2\x76\x51\xb1\x95\xa3\x30\x54\x80\x24\xb1\x19\xb5\x1f\x5f\x00\xe0\x0b\xb8\x12\x85\x7e\x0f\xd7\xf6\x10\x81\x69\xe5\x41\x50\xe0\x6d\x75\x76\xa8\xc4\xec\x61\x0a\x2e\xee\x24\x8b\x10\x85\x49\xcc\xbb\xa4\x72\x59\x69\x2a\x2b\xce\x1f\xfa\x6a\x12\xe9\x9b\x8a\xa5\xd8\x2f\xf7\x40\x29\x4e\xb2\x0e\xa3\x70\x03\x17\xdb\xd3\x51\x58\x52\xd1\xcd\xa8\xfd\x18\x14\x62\xf9\xa2\xa0\x02\x9d\x44\xe1\x2a\x5a\x85\xf9\x3b\xaa\xa7\x9d\x4b\x8c\x8f\x12\xb1\xa0\x56\x8f\x55\xda\x03\x65\x51\xf9\xbd\x4c\x6c\xb4\xd4\x17\xbe\x19\x85\x2f\x54\xe5\xa2\xd2\xfe\x00\xa2\x76\xd0\x6f\x61\xbe\x1f\x5b\x50\xf9\x35\xc5\xeb\xa9\xb2\xe5\x5e\x0e\xf6\x78\xd8\xab\xd1\x3b\xee\xbb\xdf\xbe\x87\x89\xed\x87\x0a\xe9\xf7\xa3\x9f\x9c\x53\xe9\x92\x3f\x2d\xd5\xcd\xf2\x45\xfb\x9b\xea\xbc\xa6\x58\x16\x7b\x2c\x6d\x35\x9c\xf0\x85\x72\xaa\xd8\x6f\xbf\x8f\xbc\x98\x47\x4f\xfa\xc5\x48\xf6\xfa\x7e\x24\x4e\xde\x53\x25\x0f\xa3\xb0\x96\x0c\xf3\x07\x15\xfe\x20\xb6\xa1\xfe\x1d\x24\xd7\x85\xf2\xa6\x78\xf9\x1a\xd8\xe9\xb9\x91\x58\xdc\x7e\x08\x0a\x95\xcc\xab\xa8\x14\xb6\xfa\x34\xc9\xb0\x9f\x29\x85\xc7\xaa\x54\x4d\x2d\xb6\x54\x2f\xcb\x9f\x14\xa6\xac\x49\xa6\xdc\x2b\x70\x61\x2b\xde\x4e\xa3\xae\xa9\xc1\x4d\xea\xf0\xaf\x89\xa9\x1b\xd4\xbe\xc2\xe1\xd7\xbd\x2d\x6f\xf6\xe6\x20\xdc\x49\xa7\x53\x1a\x9f\x84\xe1\x6b\xd0\xdb\x80\x67\xbf\x22\xda\xad\x9c\x6d\xb9\xa6\x7e\x65\x30\x9d\x42\xac\x83\xe1\x7f\xfc\x13\xb7\x8f\x1a\x87\x3b\xe9\x54\xca\xee\xed\x91\x01\xd4\x6e\x3e\xdb\xdf\x15\xa7\x52\x74\xb2\x40\x35\xfe\x95\xa5\x4f\xa9\xa7\x8b\x99\xbb\x58\xfb\xbd\xa5\x53\x60\x26\x4f\xa7\xee\x22\x05\xfe\xf5\x79\x86\x3f\x24\xba\x83\x04\xc3\x1f\xee\xaf\xef\x66\x30\xe7\xa7\x0a\xc3\x30\x70\xc3\x1d\xc9\x33\xc7\xc1\xad\x2c\x73\x60\x9c\x18\x4c\x1f\xb8\x84\xc2\x9a\xc3\xf0\xf9\x67\x9f\xe1\x9a\x4a\xff\x6f\xca\xf6\xf7\x7f\xbf\x74\xdd\x64\x7b\xf1\x4b\x55\xfd\xe8\xdd\xf3\x1f\x0b\x33\x53\x95\xce\xc2\x1f\xf1\x2a\x01\xf5\x59\x05\xf8\x5d\x45\x75\xb8\x45\xa7\x26\x2c\x5b\x1f\xee\xbd\x47\x3e\x96\xf7\xdb\xff\x59\x1e\xbe\x0b\xfe\x3f\xda\x1c\xfa\xb3\xba\xee\xa6\xd3\xa9\x51\xcb\x86\x6f\x33\x80\x97\x3f\x5e\xfd\x36\x31\x73\x14\x92\x9b\x21\xb9\x01\x9c\x02\x7e\x0c\x65\x80\xda\x36\x02\x94\xf2\xfe\x3d\xa6\xf1\xc9\x24\x35\xdb\x6b\x6e\x30\x9d\x4a\xb1\x51\x85\xfd\xd5\x35\x30\x99\xa1\x48\x52\x43\x43\x58\x31\x9d\x4a\x25\x9f\xab\xd9\xeb\xb7\x5d\x62\x5c\xe1\xbd\xdc\x9f\xd4\x65\xa0\x5f\x2f\x7b\x83\x13\xee\x3a\xb8\x38\xf8\xf3\x79\xe8\xce\xa5\x3c\x3c\x1d\x54\xad\xdd\x4d\xff\x3b\x00\x00\xff\xff\x2c\xf7\x7f\x8f\x89\x0f\x00\x00")

func main_testGoBytes() ([]byte, error) {
	return bindataRead(
		_main_testGo,
		"main_test.go",
	)
}

func main_testGo() (*asset, error) {
	bytes, err := main_testGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "main_test.go", size: 3977, mode: os.FileMode(436), modTime: time.Unix(1601792696, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"bindata.go":    bindataGo,
	"blacklist.txt": blacklistTxt,
	"main.go":       mainGo,
	"main_test.go":  main_testGo,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("nonexistent") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"bindata.go":    &bintree{bindataGo, map[string]*bintree{}},
	"blacklist.txt": &bintree{blacklistTxt, map[string]*bintree{}},
	"main.go":       &bintree{mainGo, map[string]*bintree{}},
	"main_test.go":  &bintree{main_testGo, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
