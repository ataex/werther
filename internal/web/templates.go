// Code generated by go-bindata. DO NOT EDIT.
// sources:
// templates/login.tmpl (1.216kB)
// templates/static/script.js (1.24kB)
// templates/static/style.css (4.316kB)

package web

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _loginTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x93\x4f\x6f\xdb\x3c\x0c\xc6\xcf\xce\xa7\x20\x74\x78\x6f\x8d\x3f\xc0\x6b\xe7\x52\x6c\x40\x81\x02\x2b\xda\x0e\x3b\x16\xb2\x44\xc7\x4a\x64\xc9\x20\x15\xaf\x81\xe1\xef\x3e\xc8\xb6\xf2\x67\x4d\xb0\x93\xa5\x87\x0f\x7f\x22\x25\x7a\x18\x40\x63\x6d\x1c\x82\x08\x26\x58\x14\x30\x8e\xab\xec\xd9\x6f\x8d\x83\x17\xf2\xbd\xd1\x48\xf0\x0b\x29\x34\x48\xab\x61\x00\x74\x3a\x3a\x56\x17\x79\x1c\x8e\x29\xaf\xb0\xc6\xed\x81\xd0\x96\xb3\xca\x0d\x62\x10\xd0\x10\xd6\x51\x91\xc1\xa8\x7c\x0a\xac\x15\xb3\xd8\xdc\x01\xee\x78\xa1\xb1\x22\xd3\x05\x08\xc7\x0e\x4b\x11\xf0\x33\xe4\x3b\xd9\xcb\x59\x15\xc0\xa4\xce\xd0\x49\x5b\xef\x58\x6c\x8a\x65\x73\x8f\xae\xbc\x0b\xe8\xc2\x72\x84\x36\x3d\x28\x2b\x99\x4b\x61\x63\xd3\x0f\x9d\xdc\xa2\xd8\xac\xb2\xab\x50\xed\xa9\x9d\xc4\xac\xe8\x92\xd6\x22\x73\xf2\x66\xc3\x00\xa6\x86\xf5\x13\x3f\xb9\x5e\x5a\xa3\x1f\x09\x35\xba\x60\xa4\xe5\xe9\x9c\x2c\xcb\x96\x08\x1c\x18\xc9\xc9\x16\xc1\x13\x74\x92\xf9\xb7\x27\xbd\x20\xd0\x32\x9e\x38\x21\xda\xec\x37\x22\x4f\x67\xc4\x2c\x02\x23\xf5\x48\x80\x31\x78\x99\xbb\xf8\xfe\x73\x15\x77\xff\xa7\xc0\x7c\x05\xb1\xf6\xbc\x9b\x7b\x88\xed\x5c\x77\x3d\x35\x08\x52\x05\xe3\x5d\x29\x86\x01\xd6\xd3\x08\xfc\x7c\x7d\x86\x71\x14\xd0\x62\x68\xbc\x2e\xc5\xcb\x8f\xb7\xf7\xb9\xe1\xac\x30\xae\x3b\xa4\xb7\x69\x8c\xd6\xe8\x04\xc4\xbe\x4a\xa1\x98\xea\x8f\xe0\xf7\x51\xe9\xa5\x3d\x60\x19\x81\x8f\x6f\xaf\xdf\xdf\xa3\x08\xe3\xf8\x4f\xc4\x54\xd5\x87\x6a\xa4\xb5\xe8\xb6\x78\xc5\x49\xe2\xc4\xf9\x0a\x8a\x73\x22\xa0\xb3\x52\x61\xe3\xad\x46\x2a\x45\xba\xf2\x44\x3f\xed\xf3\x1b\x85\xa4\x37\xf9\x8b\x71\x96\x67\xc6\x69\x9f\xa7\x22\x2e\xe6\x45\x35\xa8\xf6\x95\xff\x04\xc2\x16\xdb\x0a\xe9\x21\x4e\x9d\x34\x0e\x69\xb9\xbe\x9b\xf6\x07\xdf\x23\x59\x79\x4c\x9e\xeb\xc2\x92\x2b\x55\x90\xd8\x02\xf2\x93\xff\x16\xf4\xcb\xd1\xf7\x7c\x71\xd1\x4a\xda\xc7\x7f\x48\x9b\xfe\x04\xbd\xda\x58\x59\xa1\x85\xda\xd3\x45\x01\x9b\xd7\x65\x05\x2d\x16\xf9\xe4\x48\x5d\x9e\x73\x2f\x97\xd5\x21\x04\xef\x96\xbe\xf8\x50\xb5\x26\x88\xcd\xf4\xe6\x45\x3e\xc7\xe6\x41\xcd\xe3\x5c\x4e\x3f\xe3\x92\xbc\x7c\xcf\x63\xfd\x27\x00\x00\xff\xff\x7b\xe3\xda\xaf\xc0\x04\x00\x00")

func loginTmplBytes() ([]byte, error) {
	return bindataRead(
		_loginTmpl,
		"login.tmpl",
	)
}

func loginTmpl() (*asset, error) {
	bytes, err := loginTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "login.tmpl", size: 1216, mode: os.FileMode(0644), modTime: time.Unix(1557837078, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x57, 0xfc, 0x84, 0x28, 0x53, 0xbc, 0x7e, 0xe8, 0xf7, 0x63, 0x72, 0x3, 0x51, 0x21, 0xda, 0x48, 0x1f, 0x45, 0x63, 0xe8, 0x32, 0x66, 0x1b, 0xfd, 0x47, 0xca, 0x33, 0x3c, 0x36, 0x79, 0xe4, 0xf9}}
	return a, nil
}

var _staticScriptJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x52\xcb\x6e\xdb\x30\x10\xbc\xeb\x2b\x16\xba\x44\x0e\x54\xe6\x92\x53\x0d\xdd\x9a\xa2\x05\xd2\x53\xda\x53\xd1\x03\x23\x8e\x54\xa2\x7c\x38\x7c\xc8\x28\x6a\xff\x7b\x41\xdb\x7a\xf7\x11\x98\x07\x81\x10\x77\x66\x67\x67\xe7\xee\x96\x3a\xa9\xdf\x92\x47\x50\x84\x40\xc1\x57\xf7\xe4\xcf\xdf\x7d\x75\x4f\xb7\x77\xd9\x5e\x1a\x61\xf7\xcc\x1a\x65\xb9\xa0\x8a\x9a\x68\xea\x20\xad\x29\x36\xf4\x2b\x23\x22\xea\xb8\xa3\xe8\xe1\x1e\x14\x34\x55\x24\x6c\x1d\x35\x4c\x60\x2f\x11\xee\xe7\x13\x14\xea\x60\x5d\x91\x33\x65\x5b\x69\xde\x34\xd6\x69\x92\x66\x17\xc3\x57\xc3\x35\xaa\x9b\x04\x4d\xb7\x9b\x6f\xf9\xa6\x3c\x11\xa6\xb3\xe3\xde\x5f\x49\x98\xa0\x7b\xeb\xc4\x9c\xd0\x41\xe3\x4a\xc2\x04\xd5\xcf\x70\x73\xc2\x53\xf1\xfb\x54\xfb\x2a\xc6\x7c\xb3\xcd\x4e\xd8\xde\x2a\xd6\x71\x15\x41\x15\x79\x78\x2f\xad\x79\x0a\xd6\xf1\x16\xac\x45\xf8\x18\xa0\x8b\xd1\x98\xcd\x36\x9b\x4e\xc0\xea\xef\xa8\x7f\x40\xfc\x03\x3a\x28\xee\x9b\xca\x86\x8a\x65\xe3\x8a\x4c\x54\x8a\x0e\x87\x95\xa4\x8a\xf2\xbc\xdf\xee\x4c\x72\x63\xeb\xe8\x8b\x8b\x9e\x23\x41\x79\x4c\xca\xfa\x9d\x2d\xca\xce\x0a\x06\xbb\x18\x17\xe2\xa1\x83\x09\x8f\xd2\x07\x18\xb8\x22\xf7\xf1\x59\xcb\x90\x97\x63\xb6\x30\x6d\x9f\x02\xa6\x7d\xfb\x9f\xed\xed\x98\x86\xf7\xbc\xc5\xe0\xf4\x75\x83\xd3\xe1\x30\xa0\x67\x53\xad\xd0\xeb\x97\xb9\x6d\xe9\x5c\x74\x33\x69\x0c\xdc\x87\xcf\x9f\x1e\xa9\xa2\xfc\xcb\x65\xb3\xc4\x8d\xa0\x3e\xae\xc4\x1d\xc8\xe1\x25\x4a\x07\x91\x6f\x67\x2c\x60\x3b\x87\xe4\xd9\x3b\x34\x3c\xaa\xd0\x7b\x3b\x86\x3b\x44\x67\xc6\x7f\xc7\xe1\xb6\x48\x88\x5f\x85\xab\x5c\x98\x30\x61\x4e\xde\x2d\x43\xb7\x9c\xef\x6f\xfc\x43\x02\xcb\x55\x6e\x27\x1d\x56\x11\xfa\x03\xa5\x83\xb6\x1d\xd6\xb9\x1e\x67\x3d\xc7\xac\xa4\x86\x2b\x9f\xf4\x1f\xb7\xd9\xef\x00\x00\x00\xff\xff\xf0\x3d\xdb\x34\xd8\x04\x00\x00")

func staticScriptJsBytes() ([]byte, error) {
	return bindataRead(
		_staticScriptJs,
		"static/script.js",
	)
}

func staticScriptJs() (*asset, error) {
	bytes, err := staticScriptJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/script.js", size: 1240, mode: os.FileMode(0644), modTime: time.Unix(1557837078, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x21, 0x83, 0x40, 0xc4, 0xb1, 0x4e, 0x2c, 0xf8, 0x84, 0x11, 0x9b, 0x80, 0xc2, 0xe6, 0xab, 0xb5, 0xf8, 0xd5, 0x3b, 0xc9, 0x2e, 0x5b, 0x12, 0x7, 0x29, 0x2f, 0x21, 0x5f, 0x59, 0x35, 0xf7, 0xad}}
	return a, nil
}

var _staticStyleCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x57\xfb\x6e\xa3\xb8\x17\xfe\x9f\xa7\x38\xea\x4f\x95\x9a\x9f\x02\x63\x20\xe9\x24\x44\xbb\x3b\xcf\xb0\x6f\x60\xc0\x80\x55\xb0\x91\xed\x4c\x49\x57\xdd\x67\x5f\xf9\xc2\xb5\x24\xcd\xac\xaa\x4d\x66\x2a\x6c\x8e\xcf\xe5\x3b\x9f\x3f\x3b\x3f\x68\xd3\x72\xa1\xe0\x2c\xea\xa7\x4a\xa9\x56\x26\xdf\xbe\x15\x9c\x29\x19\x94\x9c\x97\x35\xc1\x2d\x95\x41\xc6\x9b\x6f\x99\x94\x7f\x14\xb8\xa1\xf5\xe5\xb7\x3f\x79\xca\x15\x4f\x62\x84\x36\x27\xcf\xab\x54\x53\xc3\x5f\x1e\x40\xca\x3b\x5f\xd2\x37\xca\xca\x04\x52\x2e\x72\x22\xfc\x94\x77\x27\xef\xdd\xf3\xfe\xbf\xf5\x92\x04\x17\x8a\x08\xfd\x90\x92\x82\x0b\xf2\x61\x0d\x65\x15\x11\x54\x99\x05\x41\xcd\x4b\xca\xfc\x16\x97\xd6\xee\x95\xe6\xaa\x4a\x20\x7e\x46\x6d\x77\xf2\x00\x5a\x9c\xe7\x66\xd1\xe1\x11\x10\x20\x3d\xd5\x60\x51\x52\x96\x00\x3e\x2b\x6e\x7c\xfc\x68\x48\x4e\x31\x3c\x35\xb8\xf3\xdd\xfa\xdd\x01\xb5\xdd\xc6\x78\x5c\x46\x18\x62\x84\x08\x3d\x9e\xcc\xc4\x34\xc8\xfe\xd1\x46\x79\x37\xd9\x15\x5c\x34\x66\x55\xcb\x25\x55\x94\xb3\x04\x04\xa9\xb1\xa2\x3f\x89\x36\x7a\xf3\x29\xcb\x49\x97\x40\xa8\x47\x29\xce\x5e\x4a\xc1\xcf\x2c\x4f\xe0\x7f\x85\xf9\xd8\x7c\x87\xb4\x86\xb2\xfa\x1a\x90\xa9\x42\xa7\xb2\x28\x37\x0a\x0e\xe1\x9e\x34\x7a\x4e\x91\x4e\xf9\xb8\xa6\x25\x4b\x20\x23\x4c\x11\x71\xea\x01\xad\x70\xce\x5f\xb5\x17\x04\x11\x6a\x3b\x40\x20\xca\x14\x3f\xa1\x2d\xb8\x7f\x41\xb4\xd9\x02\x82\x7d\xdb\x99\xff\x2b\xef\x77\x9b\x7b\x30\x1c\x70\x18\x53\x0f\xf7\x8f\x36\x79\x74\x72\xf3\xc3\xca\x15\x64\x9d\xd1\x34\x6b\xc6\x19\x59\x20\x4d\x59\x7b\x56\x26\x8e\xa6\xa6\x6f\x69\x98\xc0\x83\x25\xe2\xc3\x16\x24\x66\xd2\x97\x44\x50\x83\x2c\x3f\xab\x9a\x32\xe2\xbc\xcf\xe1\x8f\xf4\xf7\xe4\x7d\xe8\xb6\x25\x6c\x32\xa7\x92\x06\x30\xdc\x2f\x5a\xd0\x4f\x5c\xa3\xbb\x4b\x52\xd2\x37\x92\x40\xb8\x6b\xbb\xbb\x81\x1c\xcb\x9c\xfb\x08\x5c\xc7\x27\x88\xa4\x67\xa5\x38\xbb\x1f\x12\x43\x16\x25\x30\x93\x7a\x79\x02\xe7\xb6\x25\x22\xc3\x92\x7c\x82\x57\x94\xe2\xe7\x2c\xfd\x0c\xaf\x11\x9b\x00\x1d\x6c\xae\x19\xaf\xb9\x98\xf1\xfd\x03\x2a\x00\xfe\x2b\x49\x5f\xa8\x4b\xcc\xed\x23\x5c\xd7\x80\x82\x18\x88\x4b\xee\xd6\xbb\xec\x2c\xa4\x8e\xd2\x72\x6a\xf9\x7f\x27\xd0\x13\xf4\xa6\x5c\x0c\x5c\xee\x0b\xf4\xe3\x2b\xe8\x27\x15\xff\x49\x84\x95\xb1\x19\x66\xdf\x8f\xfb\xf4\xfb\xe9\x83\x39\xce\xb4\x3e\xac\xd8\xc7\x87\x1d\x8e\x4f\x30\x2e\x08\x1a\x22\x65\x2f\x4c\x6b\x64\x1c\xd1\x7d\x3e\x3e\x87\x4b\x74\xf7\xbf\xc0\xb9\x59\x28\x80\x86\x32\xbf\x22\xb4\xac\x54\x02\xf1\x2a\x1a\x33\x2e\x66\x9c\x29\x4c\x99\x43\xe1\x1e\x31\x9c\xaa\x1e\x5a\x55\xbd\xd3\xdc\xb3\x3b\x2d\xb6\xd3\x29\x73\x92\x98\x90\x7a\x8e\x30\x95\xc0\xc3\x83\xf6\x94\x53\xd9\xd6\xf8\x92\x40\x5a\xf3\xec\xc5\x20\x55\x13\x2c\xf4\x06\x55\xd5\xc2\x31\x04\x94\x15\x7c\x06\xf1\x5e\x8b\xa5\xcd\x61\x5d\x5f\xd7\x1c\x54\xe1\xd5\x36\xcd\x65\x6e\x02\x63\xfc\xdc\x8e\x4a\xf1\xda\xe3\x8d\xd0\xb4\xb7\x21\xd6\xdf\xf5\x98\xb2\xc5\xcc\xd5\x6f\x8d\x77\xb9\xfe\x2e\x89\x10\x39\x22\xac\xaf\xc7\x33\x0f\xc8\x7c\x86\xc2\x73\x92\x71\x81\x6d\x37\xad\x22\x5f\xf5\x13\x14\x73\x4f\xa4\x88\xd3\xd8\x26\x9e\xf2\xfc\x72\xbf\x48\xf5\x72\x60\x4b\x68\x38\x57\x95\x41\x0f\x33\x45\x71\x4d\xb1\x24\xa6\x42\xbf\xe1\x6f\x3e\x97\xdd\x07\xbb\x52\xe0\x8b\xcc\x70\xed\x92\x15\xa4\x21\x4d\x4a\x84\xbf\x60\xa9\x6d\x8a\x9f\x72\xa5\x78\x93\x40\x64\x09\x3d\xc1\xed\xa8\x25\xce\x94\x5b\x91\xec\x25\xe5\x9d\x2d\x6f\x45\x6c\x46\x8b\x9c\xfe\x1c\x06\xbe\xd6\x85\x1a\x5f\x6e\x6c\x0a\x6d\x51\xd4\xfa\xb8\xab\x68\x9e\x13\xb6\xf0\x36\x9e\x04\xbc\xc5\x19\x55\x97\x5e\x68\x07\x67\x38\x95\xbc\x3e\xab\x75\x19\x04\xa8\x49\xa1\x96\x87\x99\xe9\x2d\x6f\xdd\xd3\x7c\x5b\x2e\xd4\xbd\x97\x00\x3b\x9e\xa5\x56\xe3\x94\xd8\xfb\x5e\x51\x73\xac\x12\x13\xca\x84\xa4\x8c\x0c\xda\x41\x19\xd5\x3d\x5b\x00\x8b\x82\xa3\x05\x7b\xa2\xb9\xe1\x8e\x34\x10\xea\xe9\xcf\x15\x6b\x25\x89\xa9\xff\x30\x88\xf6\xbd\x66\xf5\x11\x50\x10\x1d\x48\x33\x6a\xfb\xfb\x8d\xb6\xcd\x79\x32\x2b\xef\x9e\x45\x93\x39\xfd\xd0\x60\xf1\x72\x83\x01\xe3\x11\xe0\x0f\x1b\x87\x90\x49\x2f\x22\xa7\x8e\x3d\xa4\x76\x7c\x23\x91\x3e\x68\x32\xd1\xc7\x41\x10\xfb\x5b\xd5\x15\x06\xcd\x65\xb4\x3f\xdc\x25\xaf\x69\x0e\xaf\x15\xb5\x56\xb3\xe3\xda\xde\x23\x04\x57\x58\x91\xa7\xdd\x3e\x27\xe5\xc6\x6e\x4e\x79\xf3\xfd\xad\x77\x03\xeb\x22\x5b\xb9\x03\xc2\x69\xa5\xa5\xf4\x77\x3b\x30\x3c\x8e\xfb\x8b\x98\xb9\x7c\x39\x6b\x04\x51\xdb\xe9\x57\x9a\xe6\xbf\x40\xa9\x5f\x6c\xea\xf8\x93\xc4\xdd\x65\xc6\xf4\xfb\x99\x77\xef\x5f\x06\x98\xf4\x6f\xda\xfe\x60\x88\x34\xec\xd6\x20\xea\xa7\x2c\x3a\x61\x18\x1c\xfb\x19\x03\xd1\xb0\x64\x09\x52\xdc\x5f\xfd\x57\x36\x85\x51\x9f\xc4\x0c\x49\x0e\x7f\xc3\x5a\xda\xbf\x5f\x45\x66\x85\xd8\xe3\x2d\x72\x72\x6d\xfc\x9a\x90\x6b\x5c\x1f\xc4\xe7\x7a\x08\x77\x77\xfb\x2f\x6a\xf3\x60\x2d\xb4\x07\xeb\xc1\x3d\xf8\xb4\x4e\xd7\x4a\x97\x80\xdb\x9d\x2b\xa5\x7e\x59\x89\x79\x9e\x7f\x4d\x80\x9b\xcd\x5a\x16\x66\x7e\x92\x86\x47\xb4\x85\xc9\x9f\xcd\xc4\xce\x9e\x65\x33\x0d\x70\xc7\x9e\x13\xca\x7f\x02\x00\x00\xff\xff\x23\x57\xc6\x0d\xdc\x10\x00\x00")

func staticStyleCssBytes() ([]byte, error) {
	return bindataRead(
		_staticStyleCss,
		"static/style.css",
	)
}

func staticStyleCss() (*asset, error) {
	bytes, err := staticStyleCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/style.css", size: 4316, mode: os.FileMode(0644), modTime: time.Unix(1557837078, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf8, 0x53, 0xba, 0x33, 0x44, 0x16, 0x28, 0xdc, 0x9c, 0x4f, 0x69, 0xd7, 0x30, 0x5, 0x56, 0x8f, 0x1f, 0x78, 0xe3, 0x53, 0x41, 0xe6, 0x42, 0x95, 0x4, 0xaa, 0x5b, 0x40, 0xc, 0x30, 0x4d, 0x68}}
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

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
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

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"login.tmpl": loginTmpl,

	"static/script.js": staticScriptJs,

	"static/style.css": staticStyleCss,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
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
	"login.tmpl": &bintree{loginTmpl, map[string]*bintree{}},
	"static": &bintree{nil, map[string]*bintree{
		"script.js": &bintree{staticScriptJs, map[string]*bintree{}},
		"style.css": &bintree{staticStyleCss, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
