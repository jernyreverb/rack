// Code generated by go-bindata.
// sources:
// models/templates/app.tmpl
// DO NOT EDIT!

package models

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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

var _templatesAppTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x3c\x6b\x73\xdb\x46\x92\xdf\xf3\x2b\xa6\x50\xb9\xa2\x12\x53\x90\xec\xec\xed\xed\x31\xe7\xad\x92\x29\xda\x56\xa2\x07\x8f\x94\xbd\x75\xb1\x5d\x29\x08\x18\x8a\x88\x40\x80\xc1\x43\x8f\xa8\xf8\xdf\xaf\x7b\x66\x00\xcc\x93\x80\x5e\x77\xdc\x8d\x45\x62\x7a\x7a\x7a\x7a\x7a\xfa\x35\x8d\xb9\xbf\x27\x11\x5d\xc4\x29\x25\x5e\xb0\x5e\x7b\x64\xb3\xf9\x8e\x10\x78\xf8\x3d\xfc\x22\xa3\xb7\xc4\x3f\x80\xbf\xcd\xc3\x55\x90\xc6\x0b\x5a\x94\xac\xe5\xa4\xfe\xc1\x9b\xe1\x3f\x42\xbc\x83\x7f\xcd\xcf\xe9\x6a\x9d\x04\x25\x7d\x9f\xe5\xab\xa0\xfc\x4c\xf3\x22\xce\x52\x8f\x8c\x88\xf7\x66\xff\xf5\xfe\xee\xfe\x7f\xc2\xff\xbd\x21\x07\x1f\x67\x69\x14\x97\xd0\x5e\x78\x23\x81\x82\x8d\x54\x0a\x1c\xc4\xbb\x08\x92\x20\x0d\x69\xbe\x1b\xb6\xa0\xfa\xd8\x46\xa7\x75\x9e\x85\xb4\x28\x1e\xd4\x27\xa7\x97\x71\x51\xe6\x77\x5d\x9d\xbc\xa3\xb4\xa4\x79\x1a\x24\x48\x31\xf1\xde\xa7\xa3\xd1\xe4\xcf\x2a\x48\x70\x06\x5f\xf0\xc9\x8c\x2e\xe0\x6b\x0b\x46\x36\x43\xe2\xfd\x0f\x05\x6c\xdf\xe0\x6b\x8d\x65\x9a\xc7\xd7\x30\x6e\x07\x92\x1a\xca\x8e\xe3\x1d\xb0\xe6\x6a\x4e\xc3\x2a\x8f\xcb\xbb\x0f\x79\x56\xad\x91\xcd\xf7\x32\x3a\xf8\xfd\xe5\x9e\x61\xc3\x05\x50\x61\x11\xa7\xf7\x8d\xcf\x4b\x20\xf5\xa6\x41\x1e\xac\x28\x50\xce\xba\x6e\x5f\x91\x35\xc2\x3e\x60\x35\xac\xf0\xf5\x5c\xc6\x49\x55\xc0\xb0\x92\x18\xc0\xc3\xf3\xbb\x35\xe5\x84\x97\x79\x9c\x5e\x7a\xc3\xb6\xe9\x90\x2e\x82\x2a\x29\x59\xab\xfa\xbc\x08\xf3\x78\x5d\xd6\x32\xe7\x89\xa6\x96\x6b\x87\x74\x9d\x64\x77\x2b\x9a\x96\x27\xc1\x6d\xbc\xaa\x56\x96\x31\xa1\xe3\x69\xb5\xba\x00\x7a\x2c\x43\x32\x49\xde\x77\x0d\x0a\xad\x02\x2f\x59\xd3\x3c\x84\x61\x82\x4b\x4a\xb2\x05\x11\x6c\xa0\x05\x29\x33\x72\x45\xe9\x9a\xe4\x55\x9a\xc2\xb4\xc8\xcd\x32\x4e\x28\xec\x43\xa4\x0b\xa7\xb9\x8d\xe4\x38\x7d\x24\xc9\xaf\xb7\x93\xcc\xf1\x3e\x1f\xc9\x93\xf4\x3a\xce\xb3\x14\x69\xb6\x13\xeb\x5e\xd2\x2d\x2b\x6a\x5d\x50\x79\x43\xf6\x1b\x47\x41\x78\x96\x26\x77\x24\x48\x92\xec\x86\x04\x21\x4e\x17\x27\x5b\x2e\xe3\x82\xa0\x0e\x5c\xe4\xd9\x8a\xc4\x69\x11\x47\x14\x1e\x52\xf2\x79\x3a\x76\xd0\x7c\x9a\xc9\x0d\x07\x88\x90\x46\x9f\x83\xa4\xa2\x7c\x57\xb3\xfd\x3b\x64\x70\xe4\x9b\x31\x89\x5f\xe9\xdd\x4b\xf3\x49\x52\x39\x8f\x60\xd3\xa7\x82\x92\x79\x75\x91\xd2\xb2\x10\x88\x90\x4f\xc5\x9a\x86\xf1\xe2\x0e\xd9\xb2\xcb\x78\x94\x64\x41\x44\x6a\x15\x41\x68\x1a\xad\xb3\x38\x2d\x8b\x17\xe1\xd9\x8c\x26\x34\x28\x6c\x13\x7a\x6e\x9d\x31\xa3\xeb\xac\x88\xcb\x2c\xb7\x2d\xd2\xd3\x06\x9b\x67\x15\x6c\x39\x12\x66\xc0\xbc\xbc\x1d\xc6\x20\x41\xd5\xdd\xcf\x4d\xc5\x39\x88\xf6\xb1\xb2\x74\x85\x18\x8f\x5c\xe2\x80\x64\x91\xe5\xcd\xa6\xb0\x10\xc7\x05\xc3\x41\xd6\x31\x18\xd6\xff\x02\xcf\x00\xac\xd2\xf8\xcd\x68\xc4\x81\x47\xa3\xa3\xe8\x9f\x8f\x21\x15\x44\x8d\x14\x7c\xbc\x7e\x54\xb9\xe5\xfe\x65\x88\x5b\x8b\xed\xd1\x8f\xc8\xda\x41\x52\xa8\xd3\xf6\xde\xce\x6c\xf2\xdf\x9f\x8e\x66\x93\xc3\x1f\xc8\x71\xb0\xba\x88\x02\x32\x06\x6b\x99\xad\xce\xb3\x75\x1c\x92\x8f\x41\x1a\x25\xb0\x62\x62\x3b\x90\x1a\xa3\x44\x26\xa8\xf7\x63\x9a\x5e\x96\x4b\x46\xe4\x6b\xb9\x49\x53\x00\x26\x7d\xa0\xf0\xec\x9c\x6b\x99\x06\x30\xc8\xb1\xc7\x32\xac\x83\x41\xd3\xf1\xf8\xe8\x70\xf6\xec\x22\x8f\x23\x23\x62\xfb\xf0\x8a\x57\x74\x02\x2d\x30\x8a\x2c\xdf\xde\x34\xcb\xcb\x69\x9e\x95\x59\x98\x69\x96\x67\x59\x96\x6b\xee\xd7\xa1\x6c\xd1\x94\xe6\x12\x9c\xf7\xf1\xfc\x7c\x8a\x2a\xed\x28\x2d\x4a\xdc\x69\xb6\x36\xb6\xd7\xa9\x0b\x62\xee\xb5\xdc\x11\xc3\x15\xdb\xc7\x9b\x3f\x79\x40\x65\xc4\x32\xdc\x32\xbf\xf3\xb1\x73\x7a\xa2\xc9\x3d\xd8\x7c\x7e\xac\x0f\x95\x6c\x99\x1a\x82\x3f\x6d\x28\xb2\xb1\xae\xf7\x8c\x16\x4c\x2b\x2b\x0b\x2e\x6d\xb9\x59\x96\x38\xcc\x28\xdb\x13\x47\x07\x27\xa3\x11\x83\x91\x66\x02\x83\x83\x73\x55\xc6\x54\xd5\x92\x68\xf6\x8a\xa2\x5a\x51\x84\x9f\x66\x49\x1c\xde\x1d\x66\x61\x65\xf8\x4d\x7c\x2b\x34\xba\x02\x63\xa9\x37\xbb\x10\x4e\xbd\xfe\x0f\x69\x10\x06\x34\x2f\x41\xf9\x88\xfe\x5f\x94\x26\xa2\xe1\x63\xe0\x93\xc5\x82\x86\xcc\x18\x33\xf3\xab\x61\x13\xa4\xc7\x69\x18\xaf\xeb\x90\x67\x4e\xf3\xeb\x38\xa4\xdc\x40\x27\x4c\x1f\xf9\xc1\x2a\xf8\x2b\x4b\x83\x9b\xc2\x0f\xb3\x95\x12\xa5\xc8\x13\x0d\x85\x42\x83\x7e\x45\x59\x8c\xda\x89\xb7\xd6\xbd\xfe\x6c\x94\xdf\x72\xab\x82\x19\x02\x16\x50\x6a\x40\xfc\x9e\xa7\x3e\x46\x4e\x72\x5e\xab\x3c\xd0\x39\xc0\x21\xef\x4e\x21\xea\x61\x3c\x88\x56\xe0\x09\x43\x1c\x18\x80\x15\x36\x78\xe1\x75\x2c\x10\x83\xe9\xb3\x48\x0c\x50\x59\x28\x64\xac\xb1\x14\x12\xcb\xbc\x1f\xf1\x67\x2d\x98\xfc\x01\xd9\x74\xb0\x4d\xfe\xd5\x42\x6e\x0c\x15\x2b\x89\xf6\x16\xb1\xe6\xa6\x67\x34\x7a\x5f\xa5\x9c\xaa\x5e\xd2\x3d\x06\xc7\xc6\x94\xe4\xf9\x4f\xef\xaa\xf0\x8a\x96\x6d\xfc\xfb\x0b\xf8\x89\x5c\x34\x76\x61\xa6\xf0\x07\xe2\xf0\xeb\xec\x16\xbe\xb7\xe1\x30\x23\x63\x06\x81\x3a\xaa\x70\x98\xbc\x29\x67\x80\x58\x78\xd2\x3a\x56\x8e\x34\xe7\x36\x72\x4f\x41\xdb\xe4\x28\x30\x22\xde\xe3\x12\xbd\xb7\x60\xe9\x0b\x78\xec\xff\x15\xaf\x3d\x3e\x96\x53\x0a\x85\x09\x46\x64\x71\x1a\xd1\x5b\x9f\xde\x8a\x98\x44\x01\x3b\xa1\x2b\xf0\xed\xe6\xf1\x5f\x8c\xa9\xaf\xdf\xfc\x43\x6d\xae\xd5\x0a\x27\xfd\x03\x2d\x0f\x4a\x2e\x1b\x86\xee\x41\xc9\xc8\x53\x63\x9f\x79\xb3\x2a\x2d\x63\x2e\xc9\x29\xf0\xfd\x8f\x42\x1d\xe0\x1c\xda\xb2\x8a\x49\xd8\x4f\xfb\x9e\x5b\x20\xec\xf1\x7e\xde\x68\x45\xe2\x3b\x42\xfd\x10\x42\xbe\x3f\xb2\x8b\x3e\xa0\x75\x56\x40\x06\xed\x99\x48\x28\xb8\x02\xda\x82\xbc\x49\xe6\xb8\xb0\xdb\x3a\xd5\x1e\xaf\xe7\x40\x5a\x94\x3c\x15\xa3\xda\x8a\xb3\xaa\x5c\x57\x65\x77\xfe\x2a\x13\x70\xc4\xdf\x3e\xb9\x16\xae\x6f\xc2\xca\xde\xa3\x8d\x1b\xca\x52\xf3\x5d\x50\x49\x61\x8c\xc5\x65\x4d\x6c\x82\x06\x4e\xb7\x89\xdf\xe1\x7f\x30\x32\xc4\x72\x0c\xaf\x94\x32\xb4\xe5\xd9\xea\x64\x61\x1e\xa4\x97\x94\x7c\x7f\xc5\x72\x85\x93\x14\x08\x45\x1d\x5b\xd4\x93\xe1\xc9\x2b\x80\xab\xd6\xa0\x38\x10\x6e\xb3\x69\x0d\x8b\x99\x13\xc3\xfd\xeb\x49\x02\xef\x4d\xd2\xe0\x22\xa1\x91\x8a\xa1\xed\x7a\x9a\xb1\xad\x63\x4d\xae\xe1\x93\x39\x28\x82\x90\xef\xae\x7d\x59\x17\xa8\xf8\xde\xd7\x4a\x80\xab\x1b\xd4\x0f\xbb\xaf\x19\x15\x82\x90\x96\x2f\xdb\x39\x54\xe7\xbe\x34\xee\x50\x17\x77\x5a\x32\xa8\x42\x86\xe4\x8f\xd4\xba\x79\x9c\xad\x56\xc1\x21\x4d\xe2\x55\x5c\xd2\x08\xfd\x24\x4f\x4a\x1c\xb5\xf9\x9f\xe1\xfe\xf0\xcd\xbf\xff\x5d\x6e\x53\x62\x0c\x9e\x3c\x32\xb2\x3e\x79\x95\x0e\xc9\x78\xfa\x89\x54\x69\x5c\xf2\x27\x14\xf7\x1f\x1d\x12\xd0\x79\xe4\xe4\x1d\xf6\x98\x1d\x9c\x48\x2d\x5e\xbb\x3f\xfa\xb2\xa7\x11\x61\x36\x7f\xef\x38\xbb\x54\xc3\x5c\x8b\xbc\x36\x30\x5c\x42\x87\x1d\x23\x48\x8a\xc0\x35\x86\x6a\xec\xb2\xcb\x82\xfd\xcb\x81\xfa\x0c\xd1\xaa\xa5\x5e\x09\x73\x47\x92\x3d\x5e\xb4\xdd\xfc\x8f\x01\x44\xaf\xf5\x6a\x08\xd9\xd0\xa4\xa7\x05\x16\xdb\xa7\x90\x72\xd5\x92\x18\xf9\x28\x60\xd0\x34\x19\xcf\xcf\x83\xe2\xea\x10\x89\x8f\x4b\x4b\xe4\xb9\x86\x29\x16\x67\xcc\x6a\x2a\x8e\xc1\xb0\xf1\xfc\x98\x09\xfa\x66\x89\x21\x39\x38\x06\x85\xfa\x18\x12\xb0\xe4\x1f\xbd\xf6\xf7\xfb\x39\x11\x62\xe0\xf3\xec\x8a\xa6\x9d\x16\xd2\x69\x1d\x85\x93\xe7\x70\x38\x34\x37\x03\xbc\xb3\xf0\x8a\xf5\x60\xdb\x1e\x97\xab\xe1\xa1\x67\xba\x1e\x72\x32\xaa\x41\x54\x3f\xd3\x40\xb5\xdc\x68\x03\x2e\x3f\xd7\xba\x34\x4e\x8d\x00\xc5\xdf\x1a\x08\x72\x9c\xab\x38\xd5\xe1\xe5\xb3\x3d\x5a\xf0\xb9\x6a\x6a\xb7\x99\x52\xad\x7b\x75\x3f\xd5\xe2\xe0\xd6\xbe\xb2\xca\x12\x8b\x83\x7b\xb4\x0a\x2e\x25\x48\xf6\xd3\x0e\x7a\x7f\x8f\x62\x4f\x7d\xa6\xcb\xd2\xc8\x3f\xc8\xf3\xe0\x6e\xb3\x31\xe0\x88\x50\x77\x69\x64\x09\x6e\x6a\x54\xf5\xf6\x60\x9e\xd8\x10\xd0\x26\xcc\xcb\x66\x9b\xa5\xcf\x10\x32\x49\x0c\xc7\x66\x33\xbc\xbf\xa7\x49\x41\x37\x1b\xf8\x9b\x46\x5b\x7a\xc1\x54\xeb\xf1\x60\xa2\x0e\x02\x5d\x28\xbe\xd9\x18\x83\xe3\xa2\x52\x48\xa9\x4c\x3d\x4f\x7b\x80\x45\xec\x62\x12\xb0\xe3\x1a\x75\xa7\xa5\xf3\xc6\x12\xa3\xb9\x88\xf3\xc6\xeb\xaa\xdd\x3a\x92\xf1\x7c\x6d\x37\x9e\x8d\x5c\x18\x16\xd4\x44\xcd\xbd\x62\x2b\xf6\x37\x4f\xc7\xee\x3a\x88\x90\x40\x0e\xa6\xd3\x5a\x4a\x51\x15\x6f\x11\x69\xdc\xe7\x07\xe3\x5f\x05\x34\x4d\xaf\xc5\x6f\x27\x34\xa8\x92\xdf\x67\x93\x0f\x47\x67\xa7\x72\x1f\xe9\xa9\xab\xa7\xe4\x41\xd1\x3b\x10\x62\xbe\x88\x5c\x84\xa5\x29\x11\xeb\xfa\x33\xd9\x45\x81\xe1\xbd\x3c\xcf\x0e\x46\xb8\x85\xc0\x11\x84\xef\xd4\x08\x0b\xff\x63\x93\x90\x6d\x02\xdc\xda\xc7\x07\x4c\xc8\x3f\x8e\xd3\xab\xcf\x41\x5e\xb8\x88\x34\x68\xec\xa0\xce\x4d\x83\x77\x7c\xf6\xe1\xf7\x0f\xb3\xb3\x4f\x53\x97\x2b\x61\x5d\xc4\xe9\xec\x6c\x3c\x99\xcf\x4d\x9d\x67\x00\xdb\x04\xf0\x73\x96\x54\x2b\x4b\x26\x42\x67\x0b\xf5\x4f\x32\x08\xdf\xd0\xab\x15\x5d\x5c\x0c\xe1\x5e\x02\xfd\x93\xf8\x1f\x33\x70\x27\xbc\xbd\xeb\x20\xdf\x03\x47\x6d\x2f\xca\x20\xa6\xce\xfd\x02\xfe\xb8\x97\x1c\x27\xc1\x3a\x6e\x36\x23\xf8\x36\xce\x60\x4c\xf0\x63\x72\x87\x20\x72\x8e\xa2\x12\x72\x22\x74\xc4\xda\x7b\xd7\x7c\x1a\x7b\x66\x0c\xaf\x19\xd7\x3d\xd4\x9a\x8c\xab\xa8\x61\x1d\xc4\xd9\xc2\x7d\x99\xc4\x2d\x82\xe7\x6e\x23\xcd\xd9\x38\xa3\xeb\x34\xe3\x8e\x26\x31\x81\x2d\xca\xd9\x9b\xdc\x96\x79\x80\xd4\x76\xaf\xae\x65\x17\x37\x9d\x4f\x82\xb5\x73\xa9\x5d\x6b\x88\x1d\x65\x43\x2c\xf6\x87\x9d\x3d\x68\x8b\xd7\x07\x51\x04\x4e\x70\x51\x77\xa8\xf7\x90\xdd\x48\x3d\x78\x6b\x3d\x91\x93\xb5\xef\xea\xe2\xe3\xd3\xb0\x63\xae\x5e\xca\xe1\x6f\x5d\x29\x1f\x81\xdd\x5b\x4f\x17\xf5\x11\xca\xba\x7b\x5f\xb8\x4d\x17\x0e\x03\x8f\xfd\x77\xf5\x81\xdb\x66\x83\x6b\xea\xd0\x41\x6c\x1a\x08\xde\xec\x08\xe7\xd2\x39\xb7\xc9\x0b\x2e\x1f\x1e\xb1\xc5\x09\xbd\xa4\x51\xab\x24\xdb\x67\x16\x52\x0d\x02\xdd\x04\x68\x39\xce\xbe\x39\x4d\x7b\xe8\x23\xa5\x18\xda\xb8\xa3\xa9\x42\xe2\x4e\xb8\x96\x57\xb0\xb9\xb8\x6a\x98\xf4\x9d\xc2\x63\xee\xc9\x42\xe4\x26\xad\xac\x32\x03\x4f\x8e\x0c\x9b\xf5\xaf\x53\xbb\x6c\x30\x87\x5f\x6d\x5b\x41\x35\xf2\xb1\x04\x4d\x2c\x5a\x6b\x9f\x77\xc4\x6d\x66\x1c\xf0\xff\x18\xb0\x49\x45\x41\xcd\x56\xaa\x9f\x69\xa0\x87\xb4\x88\x73\x1a\x8d\xd1\x84\x5a\x5d\x49\x47\x96\xa7\x97\x2b\xf9\xb2\x91\xa3\x19\x8e\x6f\xa3\xd3\x8c\xad\xf5\x40\xb4\xc9\x22\x37\x99\xbd\x56\x0e\x0c\xae\x69\x85\x45\x53\x5e\xff\xd3\xee\x62\x03\x42\x97\x49\xb3\x9e\x6a\x1b\x0e\x51\x19\xa5\xe3\xc0\xd2\x83\x7a\x23\x58\xce\x6c\x24\x1f\xc9\xa1\x9d\xed\x7a\x59\xe5\x64\xff\x3d\x67\x59\xb7\xa1\x45\xfd\xda\x54\xad\x5d\xc9\x6a\x12\x62\xd5\xae\x4e\x65\x26\xa3\xec\x4a\xa8\x59\xcb\x25\xd5\xa4\x63\x33\x7b\x39\xa3\xf4\xbd\x48\x62\xb1\xe9\x82\x7f\xc2\xa5\xce\x9f\x4a\x4f\x25\xe0\x7a\x94\x69\x0e\xc3\xde\x22\xfc\x1a\x22\xca\x72\x41\xbc\x1a\xf7\xbf\xc1\xb0\x0a\x4e\x3d\x79\xe5\xcb\x36\x59\xca\x58\xd5\xf9\x09\x7d\x0c\xab\xc1\x1c\xa3\xf2\x59\xc4\xa1\x51\xdd\xe1\x2c\xa8\xd4\xa7\xda\x89\x96\x89\x82\x51\x7c\xf4\xa8\x25\xb1\xe7\x80\xed\xcb\xd1\x94\xe1\x60\x80\xd4\x9b\x79\xad\xba\xa8\xfb\x6b\x2b\xf8\x10\x1e\xbe\x48\x21\xd5\x63\x28\x64\x6e\xd1\x63\x48\x43\x63\x8c\xf9\x1b\x69\xb0\x59\x90\x46\xd9\xaa\x20\x3b\x71\x99\x05\xed\x28\x3f\x18\xe6\x7d\xeb\x44\x1e\xb5\xfc\x6a\x8e\xdb\x95\xfe\x15\x0b\x7c\xa2\xab\xaa\x6e\xe9\x68\xf6\x5e\xc3\x63\x8d\xb5\x1a\x1f\xb7\xbb\x3d\x5a\xdf\xf6\xd8\x40\xca\xc4\xeb\x26\x1d\xd7\x4d\x51\xa9\x5c\x6f\x1e\x9e\xce\xb9\x35\xfc\xa6\x96\x59\xbc\x88\x38\xd7\x5f\x1f\xe2\xe1\x39\xb0\x2b\x3e\x90\x98\xb5\xa7\x0d\xf7\x3c\x12\x5e\x7f\xad\xdd\x8c\x17\x20\x5c\x16\x1b\x5f\x36\xb7\xac\x4f\x99\x63\xd2\x27\x68\x73\x5f\xcf\x20\xef\xfa\x89\xcb\x0b\x48\xbc\x45\xe0\x5c\x65\x92\x4f\xe4\xa4\x7a\x38\xc4\x2b\x05\x95\x91\xa4\x2a\x5b\xab\x4f\xec\x31\x30\xf5\xcc\xcd\x70\x5a\x48\x0f\x77\x72\xb7\x26\xd5\x70\x3f\xd4\x12\xd1\xa3\xf4\x52\x84\xfb\x5a\x7c\xb2\x75\xcf\x09\x28\xdd\xbb\x62\x51\x0d\x66\x2c\x58\x51\x83\xe9\x7e\x79\xe3\x38\xca\x8f\x90\xdf\xde\xbe\xcf\xfe\xb7\xb7\x6f\x1e\x25\xb8\xf2\x48\x6d\x6f\xa9\x26\x43\x54\xfd\x59\x22\x45\x47\x18\xeb\x1d\xad\xe5\x02\x2f\x2c\x52\x33\xea\x77\xde\xe7\xd9\x0a\x27\x6e\xdb\xc9\x06\xf0\x79\xe6\x02\x55\x03\xd2\xae\x08\xad\xdb\xf9\x93\x63\xb2\xcf\xeb\xf0\x28\xd2\x59\x61\x1c\xca\x0f\x9d\x1b\xc0\x76\x44\xcc\x85\x36\x09\x8a\x32\x0e\xdb\xbd\x0f\x2b\x8f\x27\x9c\xad\x2a\x68\x85\xf8\x71\xa6\x41\x09\x8a\x7b\xec\xce\x76\xde\xae\x5d\xd3\x66\x3a\xe7\xe1\x92\x02\x0a\x2f\x6e\x5f\xab\x51\xa2\x60\xde\xee\x8d\x24\x08\x35\x18\x6e\x2b\x94\xd5\x03\xb1\xba\x3a\x78\xe8\x70\xd8\xb5\x22\x62\xc3\xdf\xd7\x01\x35\xa7\x5e\x86\xb7\x6e\x80\x96\x72\x8d\xb0\xe6\xb5\x86\xa1\x3c\x27\xb7\x34\x19\x71\xa5\x73\xca\x47\x36\x6c\xe6\x3c\xad\x73\x33\x67\xa4\x8a\x3b\x8a\x4e\x4a\x59\x7d\xd8\x61\x0e\x91\x12\x08\x19\x2f\x9a\xe3\x64\x08\x59\x82\x1f\x68\x72\x86\x72\x49\xd2\xdf\xf7\x15\x65\xd6\xe2\x91\x2b\x58\x80\xf8\x28\xa1\x52\x1d\x13\x0a\x99\xf4\x48\x44\x6f\x12\x9a\x3c\x2b\x8a\xdf\xb2\x94\xd6\x43\xb6\x4d\x1f\x69\x90\x94\xcb\xf1\x92\x86\x57\x7a\xfa\x82\x37\xdd\x9d\x2f\x41\x85\x2e\xb3\x84\x65\xb5\xde\xa8\x02\xc5\x98\x78\xcd\x6a\x22\x19\x11\xbc\x4b\xfd\xd4\x88\x70\xcf\x83\xfc\xd2\x5e\xee\x66\x24\x10\x25\x74\xb5\x42\x03\x74\x23\xa7\x84\xba\x36\x66\xed\x68\x08\x54\xf0\xcb\x95\x63\x94\x47\x0c\xca\xa5\xa6\xe2\x2c\x19\x0b\x95\xff\xbc\xa7\xb4\x02\x0a\xf0\xa7\x74\x69\xe5\x66\x1b\xee\x4a\x6b\x52\x97\xf9\x3e\xa7\xdd\x52\x8c\x3b\x67\xa7\x6f\xcb\x82\x7a\xb2\xf9\x50\xfd\x25\xad\xf8\x98\xf5\xef\x6f\xdf\x54\xd4\xb6\x03\x79\xc3\x75\x7f\x64\x00\x37\x6c\x2b\x9f\xb1\xbe\xd9\x92\x08\x76\x5a\x4f\xd9\x10\xf4\x36\x91\xb6\xca\x6a\x85\x73\x3a\x80\x9d\x73\x2d\x1e\x3e\xb0\x2d\x01\xf6\xc0\x78\xd1\x9c\x39\xb2\x44\xe2\x55\x6d\x64\x5f\x6e\x2d\x0c\x29\x70\xaa\xee\x6d\xa0\x4f\x25\xc3\xc8\xe0\x7f\xd3\x8b\x84\x9f\xd7\x7f\x71\x54\x5e\xf7\xdc\xc0\xe6\x86\xbd\xbd\xdb\xb6\x6b\x2d\xe5\x29\x6a\x41\x37\x37\x38\x0a\x1e\x6b\x89\x3b\xeb\x54\xfb\x4b\x0a\xb8\xd4\x64\xab\x69\x2f\xcb\x3c\xbe\xa8\x4a\x3e\x61\xfb\x71\x5d\x53\x31\xd3\x45\x06\x03\x6e\x42\x4d\x34\x57\x96\x23\x66\xdb\xf9\x8b\xbc\x7f\x38\x21\xcf\xb1\x83\x8c\xe2\xf2\xee\x54\xe7\xd3\xe5\xe7\xf8\xdd\x38\xcb\xae\x62\x3a\x07\x5f\xf5\x0a\x22\xca\xa2\x68\xfc\x07\x9c\x95\xba\xba\xc1\x82\x65\xc1\xb1\x16\x4a\xc1\xa1\xc5\xb7\x3c\xfe\xed\x11\xf6\xba\x82\x29\xf1\x22\x75\xa3\x2d\x48\x2b\xdc\xb6\xb7\xb0\x9b\x3a\xea\xe6\x4c\xab\xd3\x17\xde\x98\x7d\x34\x80\x96\x5b\xcd\xc2\x48\xc1\x40\x57\x58\x6e\xa9\xcc\x96\x2a\x0e\x59\xd9\x0b\xb8\x4a\xe9\x2f\xd9\x45\x61\x56\x1e\xa3\x17\x95\x36\x65\xff\xae\x00\xc3\xf9\x7e\x80\x33\x26\xae\x21\xa5\x83\x6c\xa5\xa6\x66\x17\xd9\x26\x55\xe4\xec\xe2\x2c\xe4\xec\x86\x51\x77\xbf\x14\x0f\x24\x98\x27\x56\xd5\x6f\xad\xa9\x97\x3c\xa1\xd7\xfb\x8a\xf3\x6a\xbc\xf4\xe0\xfd\x16\xaf\xdf\xc7\x8c\x14\xdd\x07\xf4\xbe\xa6\xa6\x13\x38\xa8\xc0\x95\x28\x40\xb9\x84\xe5\xe0\x67\xfd\x1d\x1f\xfd\xf7\x75\x90\x13\xf6\x1e\x18\x79\x4b\x72\xfa\x67\x15\xe7\x74\x67\xc0\x1e\x0c\x7e\x30\x3a\x23\x70\x70\xa3\x80\xc2\xcf\xdd\x22\xba\x72\x00\x87\x49\x56\x45\xcd\xfb\x10\xd0\x2f\xa5\x37\x88\xc1\x1f\x63\x43\x73\x70\xb6\x63\xef\xfd\x67\x45\xf3\xbb\x82\x97\x9c\x49\x43\x4a\x8f\x2d\xc3\xda\x10\xf1\xfc\x39\xe0\xb8\xd7\x5b\xf1\x9f\x26\x4d\x32\x22\x03\x5d\x8e\x06\x7a\x87\x4d\xe7\x80\xe2\x90\xc0\x07\x3c\xfe\xe9\xd9\xe1\xe4\xf7\xf3\x63\xac\xd9\xfa\x65\x32\x3e\xff\xfd\xd3\xe9\xc1\xa7\xf3\x8f\x67\xb3\xa3\xdf\x26\x87\x40\xce\x60\xbf\x7b\x81\xe8\xed\x1a\x55\x73\x2d\x9f\xd0\x6b\x21\x64\x7f\x87\x5e\xd3\xb4\x1c\x92\x30\x83\xd0\xe1\xb6\xfc\xc1\x3e\x3b\x68\x2d\x40\x44\xfd\x24\xbb\xdc\x19\xe0\xdb\xa0\x93\xf9\x39\x99\x4d\xc6\x93\xa3\xcf\x93\x43\x98\x31\x79\x45\x7e\x99\x9f\x9d\xfa\x9c\xa1\xf1\xe2\x8e\xa3\xfd\xa1\x9b\xb3\x0c\xbb\xb2\xc0\x7e\xc4\xb2\x54\x17\x94\xb1\xb4\xd8\xe1\x7c\x1f\x4a\x14\xe7\xf9\x90\x44\x41\x19\x38\x88\xc5\x0f\x28\x15\x84\xdb\x02\xa1\x4f\x0b\xa1\x0d\x6a\xeb\x2f\x1b\xee\x5a\x6f\x45\x86\x32\x12\x05\xc5\xf2\x22\x0b\xf2\xc8\x89\xa9\x86\x5c\x07\x45\x71\x93\xf5\x00\x14\x79\x7b\x58\x32\x9c\xb2\xcf\x99\xf2\x65\xff\x9b\x2f\xde\xfc\xe8\x31\x52\x7d\xa7\x86\x89\xa3\xbd\x6f\xa3\xd7\x42\xd5\x1f\x58\xab\x1d\x44\x1d\xbf\xdd\xff\xb9\x26\xd0\x4f\xd8\x1b\xbe\xff\x8c\x7f\x26\xf1\xab\x57\x1d\x8c\xc7\x0f\x2e\x91\xe8\xfb\x25\xae\xa7\xf3\x2b\xbd\x23\x6f\x41\xa6\x0f\x6b\x46\x0e\x7a\x60\xc2\x4f\xc3\x79\x98\xa4\x81\x95\x59\xfc\xad\x8c\xc2\x8f\x1e\x18\x78\x5b\x1b\xbb\xf8\xc3\x78\xd3\xf2\xfe\x51\xec\x69\xbb\xe3\x5c\x9a\xc5\xaa\x99\x34\x15\x32\xd4\x97\x47\xb5\xcc\x01\x8b\x1c\x98\xff\xaf\x19\xc5\x04\x9c\xe5\xa3\x9d\x8a\x55\xfe\x2c\xc1\x31\x1c\xb5\x4b\x3d\xec\x82\x47\xa5\x37\x22\x7f\xfb\xdb\x4f\xdd\x90\x41\xb9\x04\x3d\xb6\x07\xce\x58\xb1\x87\xea\x8c\xe9\x2f\x1f\x0f\x0c\x5e\xc1\xe3\xe6\x9d\x14\xa9\x4d\x3c\x63\xed\x79\x95\x0e\x3a\xc7\x00\x0e\x2f\xb3\x08\x46\x99\x9e\xcd\xcf\xbb\xc1\x97\x34\x88\x60\x7d\x46\xfd\xd6\x76\x70\x10\x86\x74\x5d\x0e\x00\x3d\x10\x9d\x60\xf0\x05\x5c\xdd\xfb\xa3\xc8\x7a\x50\xc6\x10\x60\x6d\x01\xcc\x6b\x17\x9d\x2a\x1d\xcd\xed\xee\xcd\xcd\xcd\x2e\xaa\xe8\xdd\x2a\x07\x49\xc6\xeb\x22\xa2\x9e\x78\x3f\x15\x34\xdf\x3d\xb8\x04\xd4\x88\x15\x7c\xca\x64\xcf\x30\x88\x7a\xa7\x4d\x27\xea\xa0\x62\x0b\xc6\xdf\xd2\x1c\xe1\xaa\xd4\xd2\xbd\x55\x38\x1f\xa4\xe6\x50\x3a\xc1\x61\x00\xc9\x64\xbe\x8c\x8f\xce\x03\xb8\xa5\x3b\x42\x62\x25\x93\x04\x6e\x6d\x8f\x3d\x88\xf8\x2e\xb2\xe8\x0e\x8d\xb6\x69\xb3\x75\x68\xc5\xe6\xe2\xcb\xba\x55\x31\x1a\x0c\xf1\x2d\x27\xbf\x60\xbf\xd0\xcb\x73\x9b\x2c\x2b\x9a\x8f\x42\xa6\x00\x8f\x66\xaf\x11\xad\x90\x38\x8b\xd5\xd6\x91\x32\x22\x68\x39\x41\x41\x00\x04\x3b\x83\xaa\x5c\xfc\xc3\xe2\x47\xd9\xfa\x01\xbb\x06\x68\x86\x06\x12\xff\xc2\x65\x95\x5e\x01\x07\x39\x7b\x5e\xbd\x25\xec\x01\xd9\xf4\xc7\x08\x31\x87\x8c\xb0\xa7\x46\x94\x99\x83\x43\x77\x8f\x27\x7a\xa1\xaf\xe4\x47\x59\x4a\x2d\x5e\xa7\x0e\xde\x31\x0b\x4b\xf3\x36\x70\x10\x42\x3e\xe1\x3c\xcf\xf2\x81\xea\x16\x65\x5d\x0e\x8f\x3e\xe7\xc1\x04\x3b\xa1\x38\xf0\xde\x7d\xa4\x89\x4d\x7d\x11\xc4\xc9\x4e\x8f\x3e\x0f\x9c\x1c\xf3\x57\x40\xc9\x1f\x82\x7c\xc0\x2e\x91\x9c\x74\x49\x54\x3b\x66\x08\x3a\x81\xbd\x67\x32\x80\x20\x96\x6b\x6a\xf1\xe0\x99\x17\xe1\x06\x82\x63\xba\x53\x53\xbb\x9d\x0d\x08\x0f\x12\xba\x45\x58\xac\x36\xd5\x42\x15\x68\x30\xe9\x81\xfa\x5a\xb8\xe5\x90\x4e\x3b\x3a\x17\x39\xb0\x8e\x70\x1b\xc3\xde\xf9\x12\x6c\xa7\x88\x7b\x67\x95\x72\xdf\x84\x76\xae\x87\x2c\x2e\x46\x23\x06\xd4\x19\x72\x4b\xa1\xb6\x7f\x9c\xa5\x97\x75\x64\x5d\x84\x4b\x1a\x55\xea\x95\x15\x73\xf1\x6c\x72\xbb\xc6\xa3\x64\x71\x1c\xc8\x88\x13\x2d\x5a\xb1\x01\x3f\xe4\x30\x32\x61\x2c\xaa\xb6\x47\xe0\x72\x5e\xc1\x55\x05\x7a\x14\x59\x08\x16\xe7\x29\xda\x91\xcc\x5a\x1c\x4a\x7c\xc5\xdb\x0c\xbf\xc2\xd7\xaf\x9e\x48\x21\x88\x6e\x5f\x61\x94\xaf\x75\x70\xd7\x02\x88\x8c\x58\x03\x20\x04\xb6\x05\x10\xaf\x4e\x31\x00\xe9\x70\xc4\x96\x89\xb1\xac\x1f\xcf\x8c\x4c\x69\xbe\x8a\x8b\xc2\x96\x42\x21\x7a\x0e\x45\x82\xb5\x2d\x29\x51\xd6\x54\x5c\x12\x81\x48\xf8\xe5\x05\xa3\x23\x30\xca\x57\xd4\x76\x51\x83\x92\x72\x21\x8f\x5c\x14\xe9\x46\x10\x1c\x94\x6d\xf2\x42\xbb\x03\x44\x96\x23\x96\x71\x62\x68\x9c\x75\x43\x86\xb8\x4b\x03\x3f\x3c\xdb\x65\xbd\xcc\x91\xbf\xf7\xcb\xaf\x8f\xf8\x18\x14\x93\xb1\x7c\xa7\x10\x23\xea\x2c\x57\xb2\xd6\xdb\x6e\x66\x54\xaf\xa2\x00\x62\xab\x62\x97\x06\x45\xc9\x5e\x0f\x97\x4f\x48\x1f\x88\xe3\x06\xf4\xc0\xee\x9b\x27\xe0\xa0\x15\xc7\xc1\xe8\x10\x28\x50\x4a\x9d\xaf\x2f\xdb\xee\x4a\x68\x59\x85\x2d\x96\xdb\xcf\x2c\xa5\xe3\x33\x09\x4c\xec\x03\xa5\x8c\x40\x61\xbc\x00\xb0\x96\x96\xdb\x75\xd7\xd3\x4b\xca\xa5\x5b\xdc\xda\x7a\x6e\x57\xa5\xcd\x77\xb5\xc4\x75\xf3\x4d\x7b\x79\xbd\xe6\x1a\x3f\x60\xea\xcb\x08\xe3\xfd\x76\x46\x11\xc4\x14\x58\xd0\x0e\xa8\x6a\x5a\x3a\x96\xe5\x21\x03\xe9\xec\xb3\xe0\x1d\x1a\x4c\x13\xfb\xd1\xc9\x15\xed\x12\x0d\xa5\xf8\xdc\xfe\xa6\xbd\x7e\x5b\x92\x63\xf9\x7b\xdd\x92\xe4\xbc\xfc\x48\xbb\x7e\xa6\xbd\x89\x48\x79\x4e\x8c\x7b\x89\x94\x66\xed\xe4\xa5\xe3\xea\x24\xf5\xda\x24\x7d\x1c\xe9\x12\x25\xad\x09\xb3\x85\xa1\xae\x4d\x35\x18\xf7\x2b\x31\x8e\xa2\x22\xeb\x9d\x44\xb5\x5e\x6d\x2d\x8c\x71\x99\x92\xfd\x38\x4f\x31\xee\xea\x89\x8c\xbc\xde\x43\x13\xcc\x7d\xb7\xd5\xf3\x5f\x5b\xe5\x5c\x64\xd6\x4a\x79\x39\x14\xde\x57\x79\xd1\x94\x43\x1d\x8a\xcc\xe7\x8f\xd6\xc3\x39\x57\x1f\xae\x09\x68\x5e\x9f\xc2\x15\x58\x66\x66\x2d\xac\xea\xc6\x36\xd3\x71\xfd\x2b\x2e\x97\x3d\x70\x85\x6f\x3a\x89\x07\x90\x03\x08\xda\xb3\x3c\xfe\x8b\x5a\x4b\x05\x8d\x5e\xb6\xa3\x46\xe9\x0a\x2a\x2b\x5f\x7f\xb4\xa0\xd1\x9e\xa8\x87\xb8\xdf\xac\xd2\x2b\x9b\xae\x0e\x45\x2c\xdf\xcb\x63\x5e\x77\xa3\x2a\x9b\xf9\x4f\xa3\x91\xb8\x79\x4a\x68\x9b\x43\x9a\x50\x94\x93\xe6\x94\x11\x66\x88\xaf\x78\x74\x68\x23\x76\x53\x2c\x66\x6c\x72\x5e\xf5\xa0\x97\x91\x81\x2f\xac\xbd\x6b\x78\x5f\xdf\xe2\xe0\x15\x10\x52\xd1\x15\x2a\xd7\xe6\xac\x57\x5c\x76\x45\x54\xab\x2f\xe0\xf1\x5e\xc2\xa1\xd3\x46\xc8\x56\xcb\xc6\x36\x89\x6b\xff\x1b\x00\x00\xff\xff\xc1\x23\x1b\xa6\xee\x5b\x00\x00")

func templatesAppTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesAppTmpl,
		"templates/app.tmpl",
	)
}

func templatesAppTmpl() (*asset, error) {
	bytes, err := templatesAppTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/app.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
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
	"templates/app.tmpl": templatesAppTmpl,
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
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
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
	"templates": &bintree{nil, map[string]*bintree{
		"app.tmpl": &bintree{templatesAppTmpl, map[string]*bintree{}},
	}},
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

