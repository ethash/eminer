// Code generated by go-bindata. DO NOT EDIT.
// sources:
// cl/kernel1.cl
// cl/kernel2.cl
// cl/kernel3.cl

package ethash


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
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
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
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataClKernel1Cl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x59\x6d\x6f\x9b\x4a\x16\xfe\x0c\xbf\xe2\x5c\x55\xad\x98\x98\x9b\x32\x03\x26\x58\xae\x23\x79\x13\xdf\x36\xda\x34\x89\x92\xb4\xbb\x5a\xcb\x41\xc4\x26\x0e\xaa\x0d\x59\xc0\xbd\x69\xbb\xd9\xdf\x7e\x35\x67\x66\x80\x01\x9a\xf4\x56\x0a\xf5\xcc\x79\xce\xeb\x3c\xf3\xc2\xf0\xf6\x2d\x9c\x3f\x94\xc9\x36\xf9\x1e\xaf\xc0\x83\xf2\x3e\x8f\xa3\x55\x61\x9a\xaf\x56\xf1\x5d\x92\xc6\x70\x7e\x31\x3b\x3b\x3a\x0d\x8f\x67\x9f\x4f\x8e\x66\xe1\xf4\xe3\xb1\x61\x38\x3f\x11\x9e\x7d\x3e\x39\x3e\x99\x1a\xf4\x27\xca\x20\x61\x9f\x67\x67\xc7\xe7\x97\xa6\xf9\x2a\xb9\x4b\x57\xf1\x1d\x1c\x4f\xdf\x87\x57\x27\xff\x99\x55\x5a\xaa\x03\x02\x37\x08\x86\x23\xd7\x7c\x15\xa7\xab\xe4\xae\xd6\x38\x3d\x79\xff\xe1\x5a\xd7\xa9\xbb\x80\xf9\x8c\xba\xa3\x5a\x49\x22\x66\xd7\x1f\xa6\x57\x1f\xc2\xe3\xe9\xf5\xf4\x6a\x76\x1d\x5e\x4c\x2f\x67\x67\xd7\x57\xc0\x86\x7e\x8d\x39\x3b\x3f\x9e\x09\x23\x96\xef\xbd\xf5\x48\x2d\xf9\xe3\xec\x73\x78\x71\x79\xf2\x71\x66\x38\x8f\x0e\x75\x1c\x87\x8e\x5c\x4d\x6a\x3d\xda\xf0\x8d\x80\xfc\x67\x59\x8f\x04\xf6\x6a\x2d\xb8\x01\xeb\x1b\x21\x9a\xb9\xcb\xd9\xf1\xa7\xa3\x99\xf5\x15\x95\xb8\x05\xf5\xf7\x75\xff\xd1\x86\xaf\xfb\xdf\x08\x7f\x7e\xc7\xe7\x9f\x8d\x50\x2e\xcf\xaf\x4f\x7d\x2f\xa4\xd2\x63\x54\x84\xbb\x4d\x96\xae\xad\x3c\x2b\xa3\x32\xb6\xaa\xf6\x23\xb1\xc1\xc2\x9f\x84\xfb\x26\x6d\x03\x4c\x1a\x68\xda\xb3\xbe\x11\x18\x80\xcb\x1a\xee\xae\xfe\x35\xbd\xf0\x3d\x9e\x0f\x40\x6d\x9c\xff\x58\xde\x47\x79\x60\x3d\x92\xfd\xe2\xc0\x1f\x7a\x2e\xa3\x0e\x69\xea\x7d\x98\xba\xe1\x90\x32\xab\x20\x70\x97\xe5\x60\xed\x92\xb4\x84\x04\x26\x10\x8c\x21\x81\xdf\x26\xc0\x86\x63\x18\x0c\x12\x02\x3f\xa0\x98\x27\x0b\x98\x80\xf3\xe9\x74\x0c\xf0\x04\xc5\x3c\xc0\xe6\x63\xe0\x68\xff\x28\x97\x7f\x89\x97\xcb\xe8\x4b\x78\x47\x7d\xc7\xb1\x0a\x1b\x82\x86\x53\x4d\x16\xd9\x90\xed\xca\xe2\x7b\xd3\x7f\xce\xcd\x8e\x21\x87\x77\xc0\xdc\x31\x77\xdd\x54\x09\xf3\x6c\x97\xae\xb8\x62\x3e\x18\xd8\xc0\x86\x64\x0c\x4f\x3f\x41\x30\x57\x99\x37\xcd\x30\x5c\x66\x69\x51\x46\x69\x09\x58\x1e\xc0\x26\x5c\x1e\xcd\x99\xc7\x13\xf9\x61\x72\xda\x74\x72\xb1\xdb\xdd\x81\x13\x30\xd9\x1d\xe8\xdd\xd3\x56\x77\x20\xff\xfa\x8d\xfc\xa3\xd5\x1d\xe8\x2e\x75\x23\x41\xbb\x5b\x8a\x46\x3d\xb6\x9d\x3a\x92\x56\x77\xd0\xe3\xb2\xc7\x88\xf4\xd1\x36\x12\xe8\x71\xeb\xc3\xde\xdb\x1d\x38\xc1\xa8\x3f\x6e\xb7\xbf\xbb\xaf\xb0\x5c\xd0\x5b\x41\xa7\xaf\xde\x4e\x7f\xf7\x4f\x2b\xd8\xb1\xfd\xec\x30\x38\x58\xc1\xa7\xb1\x69\x16\x65\x54\x26\x4b\xf8\x9a\x25\xab\x3e\xee\x21\xc3\xf6\x20\xb2\x25\xcb\x04\xaf\xb5\x96\xa4\xfd\x0f\xd3\x90\x9d\x48\xca\xad\x03\x13\x88\xe6\xce\x02\x6e\x20\x9a\x0f\xc5\x7f\x54\x36\xa9\x6c\x33\x6c\x57\x6b\x42\x34\x67\xa2\xff\x40\xc2\x64\x93\xca\x36\x63\x0b\x1b\x28\x19\xb7\x3c\x51\xf4\x44\x05\xc6\x97\x2a\xb2\x49\x65\x9b\xd1\x96\x27\x57\xf4\x07\x12\x26\x9b\x54\xb6\x99\xdb\xeb\x89\xa1\xa7\x17\x83\xd4\x3d\x79\xa2\x7f\x24\x61\xb2\x49\x65\x9b\x79\xbd\x9e\x5c\xf4\xf4\x62\x90\xba\xa7\x17\x8b\xdd\xe7\xc9\x43\x4f\x2f\x06\xa9\x7b\x7a\xb1\xd8\xc2\x93\xee\xaa\xdc\x3e\xc8\x91\xba\xd9\x3a\x5c\x28\x22\x9e\xc0\xd6\x1b\xf3\xc6\xb0\xd9\xa0\x9a\x88\x6a\x32\x56\xcb\x78\xd3\x17\x2d\x47\x20\xa9\xd6\xd2\x64\xac\x96\x61\x53\xb4\x28\xca\x0e\x9a\x0d\xaa\x89\xa8\x26\x63\xb5\x8c\x37\x5d\xd1\x62\x28\x0b\x9a\x0d\xaa\x89\xa8\x26\x63\xb5\x8c\x37\x3d\xd1\x72\x51\x36\x6a\x36\xa8\x26\xa2\x9a\x8c\xd5\x32\x14\xf2\xc5\xbf\xda\x6e\x79\x59\x6c\xa0\x8c\x8c\x65\x89\x26\xcd\xf1\x1b\x2d\x6c\x60\x0e\x51\xfe\x34\x3d\x9c\x65\x6c\x44\xaa\x6c\x35\x29\xe5\x7c\x3d\x20\x55\x70\x9a\x59\x41\xb1\x80\x54\xa3\xa4\x1b\x5e\xd8\xe0\x4a\xa7\x1d\xb3\x38\xb5\x29\xa9\xca\xaf\xd9\xa5\x7c\x3a\xf2\x0d\x52\xd5\x55\x97\xf2\x6c\x40\x7a\xa5\x9d\x74\x50\xd7\x23\x55\xdd\x75\xbf\xc3\x85\x0d\x32\x57\x4e\xb2\x89\x3e\x73\x6d\x60\x32\xd7\x4e\xaa\x38\x6f\x95\xd9\x8e\x94\xcf\x00\x90\xc5\x67\xed\x91\x09\xb8\x5d\x97\x28\xca\xe8\x01\xe1\xb0\x49\x21\xf5\x5b\x52\x1e\xad\xf4\xd9\x0e\x16\xd3\x94\x45\xe8\x54\x08\x3d\xaa\xea\x06\x6d\xe9\x01\x77\xa9\xaa\x7b\xd0\x96\xe2\x64\x96\xc3\x46\x69\x4b\xca\x55\xc1\x27\x6a\x0a\xe9\x9a\x9c\x0d\xa0\x52\x71\x34\x69\xb9\x7d\xb0\x41\xae\x11\x72\x21\x1a\xca\x0d\x63\x5c\xf5\xf8\x72\xb9\xa8\xd6\x8a\x09\xdc\x26\x65\x11\x6f\xe2\x65\x89\xcb\xdd\x0d\x92\x8a\xff\xb2\x39\x8e\x34\x17\x95\xcb\xa3\x79\x8e\x9a\xc9\x1d\x58\xb8\x47\xc1\x21\x50\xbe\x4f\x99\x86\x9a\x2d\x4d\x6b\x94\x5b\x73\x85\x21\x9b\xdb\xe5\x71\x2b\xa6\x36\x81\x8c\x03\x3d\x81\xb1\xb9\x8a\x04\xba\x6d\xa0\xbb\xb8\xd9\x0e\x6d\x69\xd4\x53\x30\xaf\x0d\xf3\x16\x37\x5b\x1f\x11\xf6\x76\x88\x29\x68\x21\x7b\x32\x64\x43\x16\x68\xc8\x0b\x64\x18\xb2\x38\xbe\x68\x49\x3a\x34\xcd\x0e\x79\x98\x07\xdc\xf5\x90\x3f\x7c\xe1\x5f\xad\x05\x4d\xa4\xcf\x91\x81\x00\xd9\x5c\x47\x21\x0f\xda\xc8\x03\x8e\x1c\x09\x90\xcd\x75\x14\x32\x68\x23\x03\x99\x3b\x9a\x1d\x55\xb8\x51\x1b\x37\x92\xc9\x8f\xea\xe4\xb5\xec\x03\x95\xbd\x4a\x9f\x3a\x22\x63\x55\x00\x4a\x65\x5b\x51\x4c\x1b\x53\xa4\x08\xc5\x71\xa2\x82\x24\x54\xc6\x52\x91\x59\xc3\x23\x09\xa8\x2b\x91\x36\xea\xd6\xf8\x0e\x15\x28\x72\x81\x7a\x12\x69\xa3\x6e\x8d\xef\x30\x82\x2a\x4a\x48\x17\x5e\x03\xdc\xe1\x05\x55\xc4\xa0\x0d\x66\xd4\x75\x18\xb6\xea\xe0\xd7\x75\xe8\x70\x81\x22\x19\x28\x0e\x1a\x45\x3a\x50\xbf\xe1\xba\xc3\x08\x8a\x94\xa0\x81\x44\xda\xa8\x5b\xe3\x3b\xbc\xa0\x48\x0c\x3a\x92\x48\x1b\x75\x6b\x7c\x87\x1d\x54\xd1\x43\xba\x18\x35\xc0\x1d\x8a\x50\xc5\x11\x3a\xea\xa9\x03\x6b\xf1\x81\xd5\x7c\x60\x1d\x3e\x30\xb1\x64\xe0\x48\x31\xe4\x03\x6b\xf0\x81\x75\xf8\xc0\x90\x0f\xcc\x95\x48\x1b\x75\x6b\x7c\x77\x69\x40\x3e\x30\x4f\x22\x6d\xd4\xad\xf1\x1d\x3e\x30\xc5\x07\xe9\xa2\xc1\x07\xd6\xe1\x03\x53\x7c\x60\x8a\x0f\x86\x61\x3c\x99\xf8\xf7\x64\x3e\x99\x66\xf9\xed\x21\x5e\xc5\x77\x50\x94\xf9\x6e\x59\xf2\x03\x39\xae\xa3\xe2\x59\xcc\x5d\x06\x6f\xa1\x48\xbe\xc7\xd9\x9d\x7c\x53\x5f\x8c\xcd\x27\xb8\x8f\x8a\x7b\x97\x85\xe5\xb8\x36\xb0\x4b\x93\x2c\xad\xf4\x3d\x10\xe7\x37\xaf\x98\xfb\x5e\xcb\x84\x47\x70\x89\xe5\xef\x01\xf8\x68\x41\x92\xb4\x24\xb8\x9e\x27\x69\xc9\xc4\xb3\x8b\x60\x15\xc4\x13\xcf\x2e\xc4\xab\x20\x81\x78\x76\x21\x41\x05\xa1\xbe\xb1\x8a\xca\x48\xe5\xe6\x7b\xdd\xdc\x4c\x95\x1b\xc8\xdc\x8a\x39\x65\x41\x4f\x79\x3a\x15\xe8\xc2\xbc\x1a\x17\x48\x5c\xd0\x87\x0b\xea\x52\x01\x77\x8b\xd5\xd2\x61\x5a\xb9\x24\x86\xf5\x80\x1a\x15\x93\xa8\x4e\x64\x7a\xd1\x24\xaa\x13\x57\xab\x6e\x02\x45\xfd\x1e\x18\xf5\x6b\xb6\x50\x16\x3c\x43\x17\x45\x37\xe6\x38\x3d\x05\xd5\xb8\xa2\x43\x7a\xc9\xd2\x81\xb0\x2e\x15\x3a\x18\x8d\x0b\x2a\xa5\x0e\xaa\x99\x12\x73\x1c\x4c\xe9\x55\x72\xd7\xba\x47\xfc\x6d\xd2\x7b\xf1\x68\x86\x61\x54\x96\x79\x72\xbb\x2b\xe3\x30\xb4\xac\x3c\xfe\xef\x2a\xfc\x33\xcb\xbf\x84\xeb\x3c\xdb\x3d\x84\xdc\x8f\xf5\xfe\xf2\xfc\xd3\x05\x5e\xf1\xd9\x40\xf9\x1b\x11\x21\xea\xa6\x30\x0c\xbf\xc4\x79\x1a\x6f\xc4\x8b\x77\x11\x47\x1b\xcb\x34\xc2\x70\xbd\xc9\x6e\x23\xde\xb9\x89\xca\x64\x13\xe3\x78\xec\x41\x1e\x17\x65\x9e\x2c\x4b\x58\x87\xd9\xae\x7c\xd8\x95\x36\xc7\x56\xb7\x40\x6a\x06\x8b\x97\xf2\x3d\x58\x87\xf7\x71\xb4\x8a\x73\xbb\x61\xb1\x1a\xb7\x1a\xb4\x8a\xd6\xf4\x65\x08\xb3\xd5\x54\x29\xca\x28\x2f\xc3\x34\x4b\x97\x71\xd5\x57\x46\xf9\x3a\x2e\x4d\x43\xec\xd2\xc8\x6d\xf1\xc6\xb7\x4e\x56\x30\x81\x75\x5c\x4a\xeb\x61\xb2\xb2\xf0\x0c\xd9\xc0\x88\xeb\xdf\x50\x20\x93\x15\xbc\x86\xeb\x0f\x97\xb3\xe9\xf1\x95\x0e\xe3\x71\x09\x90\x25\x50\x75\x59\x09\x1c\x1e\x02\xbe\x43\x85\xe1\x26\x5b\xca\x24\xf8\x8c\x87\xe2\x3e\xca\xe3\xdb\xdd\xdd\x1c\x2f\x62\x39\x98\x73\xa2\x03\xdb\x93\x4e\x10\x0e\x93\x4a\x0d\x06\xca\x6f\x7d\x42\x2d\xca\xa8\x8c\xe7\x6c\x88\x34\xb6\xe4\xec\xdf\x23\xd8\x4d\xc4\xf1\xd4\xb2\xda\xb7\x73\xde\x1e\x51\xe3\x41\xf0\x70\x6b\x1a\xc2\x0e\xae\xef\x8d\x9a\xc2\x80\x17\x81\xcb\xf5\x0b\x4c\xbf\x73\x81\x69\x1a\xd2\x84\xb8\xc6\x1c\xf3\xe5\x5f\x59\x1d\x8a\xab\x4c\xfd\xbe\x0c\xaf\x32\x15\xa2\xff\xb2\xd3\xe1\x08\xd3\xd0\xaf\x3b\x39\xde\xc6\x17\x2b\xd3\x78\xf5\x90\x47\xeb\x6d\x04\xbb\x34\xcf\x36\x1b\xa0\xcd\x30\x4b\x1c\x1d\x67\x8c\x3f\xde\x55\xa3\xc8\x9b\x83\x81\x08\x98\x9f\xe9\x10\x36\xa9\x87\x5d\x48\x0c\x2c\xf9\xef\x87\x7c\xcd\xc6\x0a\x8a\xb9\xd9\xa8\xeb\x58\xec\x71\xc6\x6d\x94\xe7\x49\x9c\x5b\x47\xa7\xff\x0c\x4f\xcf\x8f\xa6\xa7\xe1\xc7\xd9\xc7\xf0\x8f\xd9\xd9\xd1\x4c\x1c\x0a\xc4\x5a\xb7\x4d\x1e\xd5\x40\xfe\x7e\x28\x97\x88\x9a\x6a\x6f\x40\x1c\x10\x7e\xc9\x18\x24\x69\x52\x3a\xba\xb5\x42\x46\xf4\xa2\x81\x6e\xc5\x1a\x25\x8b\x44\xc1\x22\x78\x07\xd3\xa3\xa3\xd9\xd5\xd5\xec\x8a\xb7\x06\x13\x71\xe2\x35\x0c\xe3\x36\xcb\x36\xb0\x7b\x58\x45\x65\x1c\x2a\x76\x36\x66\x0c\xaf\x54\xc4\xe9\xef\x92\x7a\xe2\xc8\xb3\x91\xee\x99\xf7\xe8\x94\x72\x24\xa5\x82\x06\xa3\xc4\x10\x35\xfd\xa9\x7e\xa3\x95\x3c\x4c\xf0\x03\x82\x28\xcd\x0d\x58\x11\x0c\x20\x21\xb6\x1c\x39\xd8\x23\x6f\xb6\xc9\x23\x99\x27\x0b\x1e\x98\xfa\xbc\x22\x8e\x35\x4f\xf8\x7c\xb1\x72\xc6\x4b\x8b\x12\x5f\x30\x70\xfd\x92\x70\x1e\x7a\x3b\xca\x37\xe2\x65\x0f\x33\xd0\x94\x98\x8a\x45\x1c\x1b\x91\x2d\x3c\xa1\x6d\xf2\x68\x0b\xc4\xbc\x6d\xeb\xf0\x10\xe8\x62\xbf\xcd\x25\x75\x58\x7b\x3e\x9f\xea\x90\x66\x6a\x95\x64\x4d\x43\x9c\xf7\x62\x73\xb3\x1a\x1f\x6b\xb6\xc9\xe3\xfe\x26\x23\x36\xb4\xfa\xee\x13\x42\x7e\x89\x81\xcf\x4c\xba\xce\xf2\x85\x67\x59\x15\x9f\x3c\xec\xfc\xea\xd4\xc3\xe4\x74\x8a\x51\xf7\x6f\x2d\x5b\xe2\xd5\xea\xd9\x75\x4b\xbc\xa5\xfc\x9d\x85\x4b\xdc\x2e\xf0\x22\xc8\xcf\x4a\xc2\x90\xb3\x20\xf0\x4e\xee\x5c\x22\x2c\x8c\xba\xd8\x64\x25\x4c\x60\x9b\xa4\x82\xc8\xe4\xe3\xf4\xdf\xe1\xf9\xa7\xeb\x8b\x4f\xd7\x57\x36\x44\x65\xb6\x4d\x96\x61\x92\x2e\xad\x37\x6a\x17\x46\x4b\x03\x71\xa5\x6a\x54\x9d\xdc\xce\x42\xec\x68\x63\x79\x30\xd7\x37\xfb\x75\x9c\xc6\x39\x9f\x65\xab\x68\x1d\x26\x65\xbc\x15\x55\xc3\xad\xc0\x06\x8d\xf9\xb8\x45\x55\xc4\xdf\x24\xeb\xfb\xb2\xbd\x63\xcb\x5d\x4c\x6e\xe7\x3d\xea\x6a\x1f\xc7\x54\xd1\x53\x9a\xad\x62\xb1\xa1\xa2\x4f\xbe\xf1\x74\xf6\xe9\x67\xbc\xa8\xa2\x2a\x33\x87\xd5\x24\xdf\x63\x04\xf2\xb8\xdc\xe5\x29\xc7\x54\xe7\x2a\xe0\x89\x72\x34\xef\x55\xbf\xf7\xd5\xb9\x0c\x17\x14\x99\xdc\x5c\xd9\x7c\xdd\xf8\xc6\xba\xd8\x17\xc7\xf9\x96\x6e\x21\xef\x79\xa4\x0a\x97\x57\x5f\x01\x6b\x20\x9e\x45\x49\x77\x57\x55\x4b\x60\xff\x87\xda\x06\x65\x51\xe3\x21\xca\xe3\xb4\x0c\x93\x74\x15\xab\xd5\x42\x45\x7a\x03\x89\x0d\xad\xc0\x12\x78\x5d\x7f\xdd\xc5\x65\xb0\xce\x86\x73\xa5\xbf\x06\xdc\x6a\x8f\xc4\xae\x8a\xd3\x8c\x42\x14\x45\xcd\xbd\x67\x12\x37\x4c\xa3\xbb\x64\x36\x87\xef\x0d\x08\x6a\x74\x17\x49\x6e\xb9\x42\x4d\xe0\xff\x78\xd0\x12\x0b\x64\xa5\xfd\x96\x6b\xff\xaf\x69\x8d\x12\x11\x1b\x4c\xa0\x27\x9b\xb1\xf9\x64\xfe\x15\x00\x00\xff\xff\x58\x7f\xce\xa4\x06\x20\x00\x00")

func bindataClKernel1ClBytes() ([]byte, error) {
	return bindataRead(
		_bindataClKernel1Cl,
		"cl/kernel1.cl",
	)
}



func bindataClKernel1Cl() (*asset, error) {
	bytes, err := bindataClKernel1ClBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "cl/kernel1.cl",
		size: 8198,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1629582740, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataClKernel2Cl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x59\xff\x6f\x9b\xba\x16\xff\x19\xfe\x8a\x73\x35\x6d\x82\x86\xdb\x61\x43\x28\x51\x96\x4a\x79\x6d\xee\x56\xbd\xae\xad\xda\x6e\xef\xe9\x45\x29\xa2\x09\x4d\xd1\x12\xe8\x03\x67\xb7\xdb\x5e\xdf\xdf\x7e\xe5\x63\x1b\x30\xb0\x76\x77\x52\x59\xec\xf3\xf1\xf9\xe6\x8f\x8f\x8d\x79\xfb\x16\xce\x1f\x58\xba\x4d\xbf\x27\x2b\x08\x81\xdd\x17\x49\xbc\x2a\x4d\xf3\xd5\x2a\xb9\x4b\xb3\x04\xce\x2f\x66\x67\x47\xa7\xd1\xf1\xec\xf3\xc9\xd1\x2c\x9a\x7e\x3c\x36\x0c\xf7\x27\xc2\xb3\xcf\x27\xc7\x27\x53\x83\xfc\x64\x30\x48\xd8\xe7\xd9\xd9\xf1\xf9\xa5\x69\xbe\x4a\xef\xb2\x55\x72\x07\xc7\xd3\xf7\xd1\xd5\xc9\x7f\x66\xd5\x28\xd5\x01\xa1\x17\x86\xc3\x91\x67\xbe\x4a\xb2\x55\x7a\x57\x8f\x38\x3d\x79\xff\xe1\x5a\x1f\x53\x77\x01\x0d\x28\xf1\x46\xf5\x20\x89\x98\x5d\x7f\x98\x5e\x7d\x88\x8e\xa7\xd7\xd3\xab\xd9\x75\x74\x31\xbd\x9c\x9d\x5d\x5f\x01\x1d\x06\x35\xe6\xec\xfc\x78\x26\x94\x58\x81\xff\xd6\xb7\x6b\xc9\x1f\x67\x9f\xa3\x8b\xcb\x93\x8f\x33\xc3\x7d\x74\x89\xeb\xba\x64\xe4\x69\x52\xeb\xd1\x81\x6f\x36\xc8\x7f\x96\xf5\x68\xc3\x5e\x3d\x0a\x6e\xc0\xfa\x66\xdb\x9a\xba\xcb\xd9\xf1\xa7\xa3\x99\xf5\x15\x07\x71\x0d\xea\xef\xeb\xfe\xa3\x03\x5f\xf7\xbf\xd9\xfc\xf9\x1d\x9f\x7f\x36\x5c\xb9\x3c\xbf\x3e\x0d\xfc\x88\x48\x8b\x71\x19\xed\x36\x79\xb6\xb6\x8a\x9c\xc5\x2c\xb1\xaa\xf6\xa3\xed\x80\x85\x3f\x6d\x6e\xdb\x6e\x2b\xa0\x52\x41\x53\x9f\xf5\xcd\x86\x01\x78\xb4\x61\xee\xea\x5f\xd3\x8b\xc0\xe7\xf1\x00\xd4\xca\xf9\x8f\xe5\x7d\x5c\x84\xd6\xa3\xbd\x5f\x1e\x04\x43\xdf\xa3\xc4\xb5\x9b\xe3\x3e\x4c\xbd\x68\x48\xa8\x55\xda\x70\x97\x17\x60\xed\xd2\x8c\x41\x0a\x13\x08\xc7\x90\xc2\x6f\x13\xa0\xc3\x31\x0c\x06\xa9\x0d\x3f\xa0\x9c\xa7\x0b\x98\x80\xfb\xe9\x74\x0c\xf0\x04\xe5\x3c\xc4\xe6\x63\xe8\x6a\xff\x08\x97\x7f\x49\x96\xcb\xf8\x4b\x74\x47\x02\xd7\xb5\x4a\x07\xc2\x86\x51\x4d\x16\x3b\x90\xef\x58\xf9\xbd\x69\xbf\xe0\x6a\xc7\x50\xc0\x3b\xa0\xde\x98\x9b\x6e\x0e\x89\x8a\x7c\x97\xad\xf8\xc0\x62\x30\x70\x80\x0e\xed\x31\x3c\xfd\x04\x41\x3d\xa5\xde\x34\xa3\x68\x99\x67\x25\x8b\x33\x06\x98\x1e\xc0\x26\x5c\x1e\xcd\xa9\xcf\x03\xf9\x61\x72\xda\x74\x62\x71\xda\xdd\xa1\x1b\x52\xd9\x1d\xea\xdd\xd3\x56\x77\x28\xff\xfa\x95\xfc\xa3\xd5\x1d\xea\x26\x75\x25\x61\xbb\x5b\x8a\x46\x3d\xba\xdd\xda\x93\x56\x77\xd8\x63\xb2\x47\x89\xb4\xd1\x56\x12\xea\x7e\xeb\xd3\xde\xdb\x1d\xba\xe1\xa8\xdf\x6f\xaf\xbf\xbb\x2f\xb1\x5c\xd0\x9b\x41\xb7\x2f\xdf\x6e\x7f\xf7\x4f\x33\xd8\xd1\xfd\xec\x34\xb8\x98\xc1\xa7\xb1\x69\x96\x2c\x66\xe9\x12\xbe\xe6\xe9\xaa\x8f\x7b\xc8\xb0\x3d\x88\x1d\xc9\x32\xc1\x6b\xad\x25\x69\xff\xc3\x34\x64\x27\x92\x72\xeb\xc2\x04\xe2\xb9\xbb\x80\x1b\x88\xe7\x43\xf1\x1f\x91\x4d\x22\xdb\x14\xdb\x55\x4d\x88\xe7\x54\xf4\x1f\x48\x98\x6c\x12\xd9\xa6\x74\xe1\x00\xb1\xc7\x2d\x4b\x04\x2d\x11\x81\x09\xe4\x10\xd9\x24\xb2\x4d\x49\xcb\x92\x27\xfa\x43\x09\x93\x4d\x22\xdb\xd4\xeb\xb5\x44\xd1\xd2\x8b\x4e\xea\x96\x7c\xd1\x3f\x92\x30\xd9\x24\xb2\x4d\xfd\x5e\x4b\x1e\x5a\x7a\xd1\x49\xdd\xd2\x8b\xc9\xee\xb3\xe4\xa3\xa5\x17\x9d\xd4\x2d\xbd\x98\x6c\x61\x49\x37\xc5\xb6\x0f\x72\xa6\x6e\xb6\x2e\x17\x0a\x8f\x27\xb0\xf5\xc7\xbc\x31\x6c\x36\x88\x26\x22\x9a\x8c\xd6\x32\xde\x0c\x44\xcb\x15\x48\xa2\xb5\x34\x19\xad\x65\xd8\x14\x2d\x82\xb2\x83\x66\x83\x68\x22\xa2\xc9\x68\x2d\xe3\x4d\x4f\xb4\x28\xca\xc2\x66\x83\x68\x22\xa2\xc9\x68\x2d\xe3\x4d\x5f\xb4\x3c\x94\x8d\x9a\x0d\xa2\x89\x88\x26\xa3\xb5\x0c\x85\xbc\xf8\x57\xdb\x2d\x4f\x8b\x03\x84\xda\x63\x99\xa2\x49\x73\xfe\x46\x0b\x07\xa8\x6b\x2b\x7b\xda\x38\x5c\x65\x74\x64\x57\xd1\x6a\x52\xc2\xf9\x7a\x60\x57\xce\x69\x6a\x05\xc5\x42\xbb\x9a\x25\x5d\xf1\xc2\x01\x4f\x1a\xed\xa8\xc5\xa5\x4d\xec\x2a\xfd\x9a\x5e\xc2\x97\x23\xdf\x20\x55\x5e\x75\x29\x8f\x06\xa4\x55\xd2\x09\x07\xc7\xfa\x76\x95\x77\xdd\xee\x70\xe1\x80\x8c\x95\x93\x6c\xa2\xaf\x5c\x07\xa8\x8c\xb5\x13\x2a\xae\x5b\xa5\xb6\x23\xe5\x2b\x00\x64\xf2\x69\x7b\x66\x42\xae\xd7\xb3\x15\x65\x74\x87\x70\xda\xa4\x90\x04\x2d\x29\xf7\x56\xda\x6c\x3b\x8b\x61\xca\x24\x74\x32\x84\x16\x55\x76\xc3\xb6\xf4\x80\x9b\x54\xd9\x3d\x68\x4b\x71\x31\xcb\x69\x23\xa4\x25\xe5\x43\x21\xb0\xd5\x12\xd2\x47\x72\x36\x80\x0a\xc5\xd5\xa4\x6c\xfb\xe0\x80\xac\x11\xb2\x10\x0d\xe5\x86\x31\xae\x7a\x02\x59\x2e\xaa\x5a\x31\x81\xdb\x94\x95\xc9\x26\x59\x32\x2c\x77\x37\x48\x2a\xfe\xcb\xe1\x38\xbb\x59\x54\x2e\x8f\xe6\x05\x8e\x4c\xef\xc0\xc2\x3d\x0a\x0e\x81\xf0\x7d\xca\x34\xd4\x6a\x69\x6a\x23\x5c\x9b\x27\x14\x39\x5c\x2f\xf7\x5b\x31\xb5\x09\xa4\x1c\xe8\x0b\x8c\xc3\x87\x48\xa0\xd7\x06\x7a\x8b\x9b\xed\xd0\x91\x4a\x7d\x05\xf3\xdb\x30\x7f\x71\xb3\x0d\x10\xe1\x6c\x87\x18\x82\xe6\xb2\x2f\x5d\x36\x64\x82\x86\x3c\x41\x86\x21\x93\x13\x88\x96\xa4\x43\x53\xed\x90\xbb\x79\xc0\x4d\x0f\xf9\x23\x10\xf6\x55\x2d\x68\x22\x03\x8e\x0c\x05\xc8\xe1\x63\x14\xf2\xa0\x8d\x3c\xe0\xc8\x91\x00\x39\x7c\x8c\x42\x86\x6d\x64\x28\x63\x47\xb5\xa3\x0a\x37\x6a\xe3\x46\x32\xf8\x51\x1d\xbc\x16\x7d\xa8\xa2\x57\xe1\x13\x57\x44\xac\x12\x40\x88\x6c\x2b\x8a\x69\x73\x8a\x14\x21\x38\x4f\x44\x90\x84\x48\x5f\x2a\x32\x6b\x78\x24\x01\xf1\x24\xd2\xc1\xb1\x35\xbe\x43\x05\x82\x5c\x20\xbe\x44\x3a\x38\xb6\xc6\x77\x18\x41\x14\x25\xa4\x09\xbf\x01\xee\xf0\x82\x28\x62\x90\x06\x33\xea\x3c\x0c\x5b\x79\x08\xea\x3c\x74\xb8\x40\x90\x0c\x04\x27\x8d\x20\x1d\x48\xd0\x30\xdd\x61\x04\x41\x4a\x90\x50\x22\x1d\x1c\x5b\xe3\x3b\xbc\x20\x48\x0c\x32\x92\x48\x07\xc7\xd6\xf8\x0e\x3b\x88\xa2\x87\x34\x31\x6a\x80\x3b\x14\x21\x8a\x23\x64\xd4\x93\x07\xda\xe2\x03\xad\xf9\x40\x3b\x7c\xa0\xa2\x64\xe0\x4c\x51\xe4\x03\x6d\xf0\x81\x76\xf8\x40\x91\x0f\xd4\x93\x48\x07\xc7\xd6\xf8\x6e\x69\x40\x3e\x50\x5f\x22\x1d\x1c\x5b\xe3\x3b\x7c\xa0\x8a\x0f\xd2\x44\x83\x0f\xb4\xc3\x07\xaa\xf8\x40\x15\x1f\x0c\xc3\x78\x32\xf1\xef\xc9\x7c\x32\x4d\xf6\xed\x21\x59\x25\x77\x50\xb2\x62\xb7\x64\xfc\x40\x8e\x75\x54\x3c\xcb\xb9\x47\xe1\x2d\x94\xe9\xf7\x24\xbf\x93\x6f\xea\x8b\xb1\xf9\x04\xf7\x71\x79\xef\xd1\x88\x8d\x6b\x05\xbb\x2c\xcd\xb3\x6a\xbc\x0f\xe2\xfc\xe6\x97\xf3\xc0\x6f\xa9\xf0\x6d\x2c\xb1\xfc\x3d\x00\x1f\x2d\x48\x9a\x31\x1b\xeb\x79\x9a\x31\x2a\x9e\x5d\x04\xad\x20\xbe\x78\x76\x21\x7e\x05\x09\xc5\xb3\x0b\x09\x2b\x08\x09\x8c\x55\xcc\x62\x15\x5b\xe0\x77\x63\x33\x55\x6c\x20\x63\x2b\xe7\x84\x86\x3d\xe9\xe9\x64\xa0\x0b\xf3\x6b\x5c\x28\x71\x61\x1f\x2e\xac\x53\x05\xdc\x2c\x66\x4b\x87\x69\xe9\x92\x18\xda\x03\x6a\x64\x4c\xa2\x3a\x9e\xe9\x49\x93\xa8\x8e\x5f\xad\xbc\x09\x14\x09\x7a\x60\x24\xa8\xd9\x42\x68\xf8\x0c\x5d\x14\xdd\xa8\xeb\xf6\x24\x54\xe3\x8a\x0e\xe9\x25\x4b\x07\x42\xbb\x54\xe8\x60\x34\x2e\xa8\x90\x3a\xa8\x66\x48\xd4\x75\x31\xa4\x57\xe9\x5d\xeb\x1e\xf1\xb7\x49\xef\xc5\xa3\x19\x45\x31\x63\x45\x7a\xbb\x63\x49\x14\x59\x56\x91\xfc\x77\x15\xfd\x99\x17\x5f\xa2\x75\x91\xef\x1e\x22\x6e\xc7\x7a\x7f\x79\xfe\xe9\x02\xaf\xf8\x1c\x20\xfc\x8d\xc8\xb6\xd5\x4d\x61\x14\x7d\x49\x8a\x2c\xd9\x88\x17\xef\x32\x89\x37\x96\x69\x44\xd1\x7a\x93\xdf\xc6\xbc\x73\x13\xb3\x74\x93\xe0\x7c\xec\x41\x91\x94\xac\x48\x97\x0c\xd6\x51\xbe\x63\x0f\x3b\xe6\x70\x6c\x75\x0b\xa4\x56\xb0\x78\x29\xdf\x83\x75\x74\x9f\xc4\xab\xa4\x70\x1a\x1a\xab\x79\xab\x41\xab\x78\x4d\x5e\x86\x50\x47\x2d\x95\x92\xc5\x05\x8b\xb2\x3c\x5b\x26\x55\x1f\x8b\x8b\x75\xc2\x4c\x43\xec\xd2\xc8\x6d\xf1\xc6\xb7\x4e\x57\x30\x81\x75\xc2\xa4\xf6\x28\x5d\x59\x78\x86\x6c\x60\xc4\xf5\x6f\x24\x90\xe9\x0a\x5e\xc3\xf5\x87\xcb\xd9\xf4\xf8\x4a\x87\x71\xbf\x04\xc8\x12\xa8\x3a\xad\x36\x1c\x1e\x02\xbe\xfe\x44\xd1\x26\x5f\xca\x20\xf8\x8a\x87\xf2\x3e\x2e\x92\xdb\xdd\xdd\x1c\x2f\x62\x39\x98\x73\xa2\x03\xdb\x93\x46\x10\x0e\x93\x6a\x18\x0c\x94\xdd\xfa\x84\x5a\xb2\x98\x25\x73\x3a\x44\x1a\x5b\x72\xf5\xef\xd9\xd8\x6d\x8b\xe3\xa9\x65\xb5\x6f\xe7\xfc\x3d\x5b\xcd\x87\x8d\x87\x5b\xd3\x10\x7a\xb0\xbe\x37\x72\x0a\x03\x9e\x04\x2e\xd7\x2f\x30\x83\xce\x05\xa6\x69\x48\x15\xe2\x1a\x73\xcc\xcb\xbf\xd2\x3a\x14\x57\x99\xfa\x7d\x19\x5e\x65\x2a\x44\xff\x65\xa7\xcb\x11\xa6\xa1\x5f\x77\x72\xbc\x83\x2f\x56\xa6\xf1\xea\xa1\x88\xd7\xdb\x18\x76\x59\x91\x6f\x36\x40\x9a\x6e\x32\x9c\x1d\x77\x8c\x3f\xde\x55\xb3\xc8\x9b\x83\x81\x70\x98\x9f\xe9\x10\x36\xa9\xa7\x5d\x48\x0c\x4c\xf9\xef\x87\xbc\x66\x63\x06\xc5\xda\x6c\xe4\x75\x2c\xf6\x38\xe3\x36\x2e\x8a\x34\x29\xac\xa3\xd3\x7f\x46\xa7\xe7\x47\xd3\xd3\xe8\xe3\xec\x63\xf4\xc7\xec\xec\x68\x26\x0e\x05\xa2\x22\x6e\xd3\x47\x35\x91\xbf\x1f\xca\xe2\x58\x53\xed\x0d\x78\xa8\xf0\x97\x94\x41\x9a\xa5\xcc\xd5\xb5\x95\xd2\xa3\x17\x15\x74\x33\xd6\x48\x59\x2c\x12\x16\xc3\x3b\x98\x1e\x1d\xcd\xae\xae\x66\x57\xbc\x35\x98\x88\xf3\xbe\x61\x18\xb7\x79\xbe\x81\xdd\xc3\x2a\x66\x49\xa4\xd8\xd9\x58\x31\x3c\x53\x31\xa7\x3f\xb5\xeb\x85\x23\xcf\x46\xba\x65\xde\xa3\x53\xca\x95\x94\xf2\x1b\x8c\x12\x53\xd4\xb4\xa7\xfa\x8d\x56\xf0\x30\xc1\x0f\x08\x22\x35\x37\x60\xc5\x30\x80\xd4\x76\xe4\xcc\xc1\x9e\xfd\x66\x9b\x3e\xda\xf3\x74\xc1\x1d\x53\x9f\x57\xc4\xb1\xe6\x09\x9f\x2f\x66\xce\x78\xa9\x28\xf1\x82\x81\xf5\x4b\xc2\xb9\xeb\x6d\x2f\xdf\x88\x97\x3d\x8c\x40\x1b\x44\x95\x2f\xe2\xd8\x88\x6c\xe1\x01\x6d\xd3\x47\x47\x20\xe6\x6d\x5d\x87\x87\x40\x16\xfb\x6d\x2e\xa9\xc3\xda\xf3\xf1\x54\x87\x34\xb3\x95\xc9\x5a\x8f\x70\x40\x7d\xa3\xe1\xd9\xfb\x25\x82\x3d\xb3\xa6\x3a\xd5\x09\x8f\xaa\xca\xbc\x3c\xcb\xfc\xea\xca\x42\xdf\x75\x06\x11\xef\x6f\x55\x25\xf1\xe6\xf4\x6c\x59\x12\x2f\x21\x7f\xa7\x2e\x89\xcb\x03\x9e\x04\xf9\xd5\x48\x28\x72\x17\x36\xbc\x93\x1b\x93\x70\x0b\xbd\x2e\x37\x39\x83\x09\x6c\xd3\x4c\xf0\xd4\xfe\x38\xfd\x77\x74\xfe\xe9\xfa\xe2\xd3\xf5\x95\x03\x31\xcb\xb7\xe9\x32\x4a\xb3\xa5\xf5\x46\x6d\xb2\xa8\x69\x20\x6e\x4c\x8d\xaa\x93\xeb\x59\x88\x0d\x6b\x2c\xcf\xdd\xfa\x5e\xbe\x4e\xb2\xa4\xe0\x8b\x68\x15\xaf\xa3\x94\x25\x5b\x91\x35\xac\xf4\x0e\x68\xc4\xc6\x1d\xa8\xe2\xf5\x26\x5d\xdf\xb3\xf6\x86\x2c\x37\x29\xb9\x5b\xf7\x0c\x57\xdb\x34\x86\x8a\x96\xb2\x7c\x95\x88\xfd\x12\x6d\xf2\x7d\xa5\xb3\x0d\x3f\x63\x45\x25\x55\xa9\x39\xac\xd6\xf0\x1e\xb5\xa1\x48\xd8\xae\xc8\x38\xa6\x3a\x36\x01\x0f\x94\xa3\x79\xaf\xfa\xbd\xaf\x8e\x5d\x58\x2f\x64\x70\x73\xa5\xf3\x75\xe3\x13\xea\x62\x5f\x9c\xd6\x5b\x63\x4b\x79\x8d\x23\x87\x70\x79\xf5\x91\xaf\x06\xe2\x51\xd3\xee\x6e\x9a\xaa\xc2\xf5\x7f\x87\x6d\x50\x16\x47\x3c\xc4\x45\x92\xb1\x28\xcd\x56\x89\x2a\x06\xca\xd3\x1b\x48\x1d\x68\x39\x96\xc2\xeb\xfa\xe3\x2d\x56\xb9\x3a\x1a\xce\x95\xfe\x1c\x70\xad\x3d\x12\xa7\x4a\x4e\xd3\x0b\x91\x14\xb5\xf6\x9e\x09\xdc\x30\x8d\x6e\x45\x6c\x4e\xdf\x1b\x10\xd4\xe8\xd6\x40\xae\xb9\x42\x4d\xe0\xff\x78\x17\x2d\xea\x5f\x35\xfa\x2d\x1f\xfd\xbf\xa6\x36\x62\x0b\xdf\x60\x02\x3d\xd1\x8c\xcd\x27\xf3\xaf\x00\x00\x00\xff\xff\xd8\x69\xf3\xe5\xe5\x1f\x00\x00")

func bindataClKernel2ClBytes() ([]byte, error) {
	return bindataRead(
		_bindataClKernel2Cl,
		"cl/kernel2.cl",
	)
}



func bindataClKernel2Cl() (*asset, error) {
	bytes, err := bindataClKernel2ClBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "cl/kernel2.cl",
		size: 8165,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1629581283, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataClKernel3Cl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x59\x71\x73\x9b\x3a\x12\xff\x1b\x3e\xc5\xbe\xe9\xbc\x0e\x8a\xb9\x14\x09\x4c\xf0\xb8\xee\x8c\x2f\xf1\x6b\x33\x97\x26\x99\x38\xed\xdd\x9c\xc7\x61\x88\x8d\x1d\xa6\x36\xce\x01\xee\x73\xdb\xcb\x7d\xf6\x1b\xad\x24\x40\x40\x93\xbe\xce\x84\x5a\xda\xdf\xee\x6a\x57\x3f\xad\x84\x78\xf3\x06\xae\x1e\x8b\x64\x9b\x7c\x8f\x97\xc0\xa0\x78\xc8\xe2\x68\x99\x9b\xe6\xab\x65\xbc\x4a\xd2\x18\xae\xae\x27\x97\xa7\x17\xe1\xd9\xe4\xf3\xf9\xe9\x24\x1c\x7f\x3c\x33\x0c\xe7\x27\xc2\xcb\xcf\xe7\x67\xe7\x63\x83\xfe\x44\x19\x24\xec\xf3\xe4\xf2\xec\xea\xc6\x34\x5f\x25\xab\x74\x19\xaf\xe0\x6c\xfc\x3e\x9c\x9e\xff\x7b\x52\x6a\xa9\x0e\x08\xdc\x20\xe8\x0f\x5c\xf3\x55\x9c\x2e\x93\x55\xa5\x71\x71\xfe\xfe\xc3\xad\xae\x53\x75\x01\xf3\x19\x75\x07\x95\x92\x44\x4c\x6e\x3f\x8c\xa7\x1f\xc2\xb3\xf1\xed\x78\x3a\xb9\x0d\xaf\xc7\x37\x93\xcb\xdb\x29\xb0\xbe\x5f\x61\x2e\xaf\xce\x26\xc2\x88\xe5\x7b\x6f\x3c\x52\x49\xfe\xb8\xfc\x1c\x5e\xdf\x9c\x7f\x9c\x18\xce\xc1\xa1\x8e\xe3\xd0\x81\xab\x49\xad\x83\x0d\xdf\x08\xc8\x7f\x96\x75\x20\x70\x54\x69\xc1\x1d\x58\xdf\x08\xd1\xcc\xdd\x4c\xce\x3e\x9d\x4e\xac\xaf\xa8\xc4\x2d\xa8\xbf\xaf\xc7\x07\x1b\xbe\x1e\x7f\x23\xfc\xf9\x1d\x9f\x7f\xd6\x86\x72\x73\x75\x7b\xe1\x7b\x21\x95\x1e\xa3\x3c\xdc\x6f\x76\xe9\xda\xca\x76\x45\x54\xc4\x56\xd9\x3e\x10\x1b\x2c\xfc\x49\xb8\x6f\xd2\x34\xc0\xa4\x81\xba\x3d\xeb\x1b\x81\x1e\xb8\xac\xe6\x6e\xfa\xcf\xf1\xb5\xef\xf1\x78\x00\x2a\xe3\xfc\xc7\xe2\x21\xca\x02\xeb\x40\x8e\xf3\x13\xbf\xef\xb9\x8c\x3a\xa4\xae\xf7\x61\xec\x86\x7d\xca\xac\x9c\xc0\x6a\x97\x81\xb5\x4f\xd2\x02\x12\x18\x41\x30\x84\x04\x7e\x1b\x01\xeb\x0f\xa1\xd7\x4b\x08\xfc\x80\x7c\x96\xcc\x61\x04\xce\xa7\x8b\x21\xc0\x13\xe4\xb3\x00\x9b\x87\xc0\xd1\xfe\x51\x2e\xff\x12\x2f\x16\xd1\x97\x70\x45\x7d\xc7\xb1\x72\x1b\x82\x9a\x53\x4d\x16\xd9\xb0\xdb\x17\xf9\xf7\xba\xff\x8c\x9b\x1d\x42\x06\x6f\x81\xb9\x43\xee\xba\xae\x12\x66\xbb\x7d\xba\xe4\x8a\x59\xaf\x67\x03\xeb\x93\x21\x3c\xfd\x04\xc1\x5c\x65\xde\x34\xc3\x70\xb1\x4b\xf3\x22\x4a\x0b\xc0\xf4\x00\x36\xe1\xe6\x74\xc6\x3c\x1e\xc8\x0f\x93\xd3\xa6\x15\x8b\xdd\xec\x0e\x9c\x80\xc9\xee\x40\xef\x1e\x37\xba\x03\xf9\xd7\x6d\xe4\xef\x8d\xee\x40\x77\xa9\x1b\x09\x9a\xdd\x52\x34\xe8\xb0\xed\x54\x23\x69\x74\x07\x1d\x2e\x3b\x8c\x48\x1f\x4d\x23\x81\x3e\x6e\x7d\xda\x3b\xbb\x03\x27\x18\x74\x8f\xdb\xed\xee\xee\x4a\x2c\x17\x74\x66\xd0\xe9\xca\xb7\xd3\xdd\xfd\xd3\x0c\xb6\x6c\x3f\x3b\x0d\x0e\x66\xf0\x69\x68\x9a\x79\x11\x15\xc9\x02\xbe\xee\x92\x65\x17\xf7\x90\x61\x47\x10\xd9\x92\x65\x82\xd7\x5a\x4b\xd2\xfe\x87\x69\xc8\x4e\x24\xe5\xd6\x81\x11\x44\x33\x67\x0e\x77\x10\xcd\xfa\xe2\x3f\x2a\x9b\x54\xb6\x19\xb6\xcb\x9a\x10\xcd\x98\xe8\x3f\x91\x30\xd9\xa4\xb2\xcd\xd8\xdc\x06\x4a\x86\x0d\x4f\x14\x3d\x51\x81\xf1\xa5\x8a\x6c\x52\xd9\x66\xb4\xe1\xc9\x15\xfd\x81\x84\xc9\x26\x95\x6d\xe6\x76\x7a\x62\xe8\xe9\xc5\x41\xea\x9e\x3c\xd1\x3f\x90\x30\xd9\xa4\xb2\xcd\xbc\x4e\x4f\x2e\x7a\x7a\x71\x90\xba\xa7\x17\x93\xdd\xe5\xc9\x43\x4f\x2f\x0e\x52\xf7\xf4\x62\xb2\x85\x27\xdd\x55\xb1\x7d\x94\x33\x75\xb7\x75\xb8\x50\x8c\x78\x04\x5b\x6f\xc8\x1b\xfd\x7a\x83\x6a\x22\xaa\xc9\x58\x25\xe3\x4d\x5f\xb4\x1c\x81\xa4\x5a\x4b\x93\xb1\x4a\x86\x4d\xd1\xa2\x28\x3b\xa9\x37\xa8\x26\xa2\x9a\x8c\x55\x32\xde\x74\x45\x8b\xa1\x2c\xa8\x37\xa8\x26\xa2\x9a\x8c\x55\x32\xde\xf4\x44\xcb\x45\xd9\xa0\xde\xa0\x9a\x88\x6a\x32\x56\xc9\x50\xc8\x8b\x7f\xb9\xdd\xf2\xb4\xd8\x40\x19\x19\xca\x14\x8d\xea\xf3\x37\x98\xdb\xc0\x1c\xa2\xfc\x69\x7a\xb8\xca\xd8\x80\x94\xd1\x6a\x52\xca\xf9\x7a\x42\xca\xc1\x69\x66\x05\xc5\x02\x52\xce\x92\x6e\x78\x6e\x83\x2b\x9d\xb6\xcc\xe2\xd2\xa6\xa4\x4c\xbf\x66\x97\xf2\xe5\xc8\x37\x48\x95\x57\x5d\xca\xa3\x01\xe9\x95\xb6\xc2\x41\x5d\x8f\x94\x79\xd7\xfd\xf6\xe7\x36\xc8\x58\x39\xc9\x46\xfa\xca\xb5\x81\xc9\x58\x5b\xa1\xe2\xba\x55\x66\x5b\x52\xbe\x02\x40\x26\x9f\x35\x67\x26\xe0\x76\x5d\xa2\x28\xa3\x0f\x08\xa7\x4d\x0a\xa9\xdf\x90\xf2\xd1\x4a\x9f\xcd\xc1\x62\x98\x32\x09\xad\x0c\xa1\x47\x95\xdd\xa0\x29\x3d\xe1\x2e\x55\x76\x4f\x9a\x52\x5c\xcc\x72\xda\x28\x6d\x48\xb9\x2a\xf8\x44\x2d\x21\x5d\x93\xb3\x01\x54\x28\x8e\x26\x2d\xb6\x8f\x36\xc8\x1a\x21\x0b\x51\x5f\x6e\x18\xc3\xb2\xc7\x97\xe5\xa2\xac\x15\x23\xb8\x4f\x8a\x3c\xde\xc4\x8b\x02\xcb\xdd\x1d\x92\x8a\xff\xb2\x39\x8e\xd4\x8b\xca\xcd\xe9\x2c\x43\xcd\x64\x05\x16\xee\x51\xf0\x0e\x28\xdf\xa7\x4c\x43\xad\x96\xba\x35\xca\xad\xb9\xc2\x90\xcd\xed\xf2\x71\x2b\xa6\xd6\x81\x8c\x03\x3d\x81\xb1\xb9\x8a\x04\xba\x4d\xa0\x3b\xbf\xdb\xf6\x6d\x69\xd4\x53\x30\xaf\x09\xf3\xe6\x77\x5b\x1f\x11\xf6\xb6\x8f\x21\x68\x43\xf6\xe4\x90\x0d\x99\xa0\x3e\x4f\x90\x61\xc8\xe4\xf8\xa2\x25\xe9\x50\x37\xdb\xe7\xc3\x3c\xe1\xae\xfb\xfc\xe1\x0b\xff\xaa\x16\xd4\x91\x3e\x47\x06\x02\x64\x73\x1d\x85\x3c\x69\x22\x4f\x38\x72\x20\x40\x36\xd7\x51\xc8\xa0\x89\x0c\x64\xec\x68\x76\x50\xe2\x06\x4d\xdc\x40\x06\x3f\xa8\x82\xd7\xa2\x0f\x54\xf4\x2a\x7c\xea\x88\x88\x55\x02\x28\x95\x6d\x45\x31\x6d\x4e\x91\x22\x14\xe7\x89\x0a\x92\x50\x39\x96\x92\xcc\x1a\x1e\x49\x40\x5d\x89\xb4\x51\xb7\xc2\xb7\xa8\x40\x91\x0b\xd4\x93\x48\x1b\x75\x2b\x7c\x8b\x11\x54\x51\x42\xba\xf0\x6a\xe0\x16\x2f\xa8\x22\x06\xad\x31\xa3\xca\x43\xbf\x91\x07\xbf\xca\x43\x8b\x0b\x14\xc9\x40\x71\xd2\x28\xd2\x81\xfa\x35\xd7\x2d\x46\x50\xa4\x04\x0d\x24\xd2\x46\xdd\x0a\xdf\xe2\x05\x45\x62\xd0\x81\x44\xda\xa8\x5b\xe1\x5b\xec\xa0\x8a\x1e\xd2\xc5\xa0\x06\x6e\x51\x84\x2a\x8e\xd0\x41\x47\x1e\x58\x83\x0f\xac\xe2\x03\x6b\xf1\x81\x89\x92\x81\x33\xc5\x90\x0f\xac\xc6\x07\xd6\xe2\x03\x43\x3e\x30\x57\x22\x6d\xd4\xad\xf0\xed\xd2\x80\x7c\x60\x9e\x44\xda\xa8\x5b\xe1\x5b\x7c\x60\x8a\x0f\xd2\x45\x8d\x0f\xac\xc5\x07\xa6\xf8\xc0\x14\x1f\x0c\xc3\x78\x32\xf1\xef\xc9\x7c\x32\xcd\xe2\xdb\x63\xbc\x8c\x57\x90\x17\xd9\x7e\x51\xf0\x03\x39\xd6\x51\xf1\xcc\x67\x2e\x83\x37\x90\x27\xdf\xe3\xdd\x4a\xbe\xa9\xcf\x87\xe6\x13\x3c\x44\xf9\x83\xcb\xc2\x62\x58\x19\xd8\xa7\xc9\x2e\x2d\xf5\x3d\x10\xe7\x37\x2f\x9f\xf9\x5e\xc3\x84\x47\xb0\xc4\xf2\xf7\x00\x7c\x34\x20\x49\x5a\x10\xac\xe7\x49\x5a\x30\xf1\x6c\x23\x58\x09\xf1\xc4\xb3\x0d\xf1\x4a\x48\x20\x9e\x6d\x48\x50\x42\xa8\x6f\x2c\xa3\x22\x52\xb1\xf9\x5e\x3b\x36\x53\xc5\x06\x32\xb6\x7c\x46\x59\xd0\x91\x9e\x56\x06\xda\x30\xaf\xc2\x05\x12\x17\x74\xe1\x82\x2a\x55\xc0\xdd\x62\xb6\x74\x98\x96\x2e\x89\x61\x1d\xa0\x5a\xc6\x24\xaa\x35\x32\x3d\x69\x12\xd5\x1a\x57\x23\x6f\x02\x45\xfd\x0e\x18\xf5\x2b\xb6\x50\x16\x3c\x43\x17\x45\x37\xe6\x38\x1d\x09\xd5\xb8\xa2\x43\x3a\xc9\xd2\x82\xb0\x36\x15\x5a\x18\x8d\x0b\x2a\xa4\x16\xaa\x1e\x12\x73\x1c\x0c\xe9\x55\xb2\x6a\xdc\x23\xfe\x36\xea\xbc\x78\x34\xc3\x30\x2a\x8a\x2c\xb9\xdf\x17\x71\x18\x5a\x56\x16\xff\x67\x19\xfe\xb9\xcb\xbe\x84\xeb\x6c\xb7\x7f\x0c\xb9\x1f\xeb\xfd\xcd\xd5\xa7\x6b\xbc\xe2\xb3\x81\xf2\x37\x22\x42\xd4\x4d\x61\x18\x7e\x89\xb3\x34\xde\x88\x17\xef\x3c\x8e\x36\x96\x69\x84\xe1\x7a\xb3\xbb\x8f\x78\xe7\x26\x2a\x92\x4d\x8c\xf3\x71\x04\x59\x9c\x17\x59\xb2\x28\x60\x1d\xee\xf6\xc5\xe3\xbe\xb0\x39\xb6\xbc\x05\x52\x2b\x58\xbc\x94\x1f\xc1\x3a\x7c\x88\xa3\x65\x9c\xd9\x35\x8b\xe5\xbc\x55\xa0\x65\xb4\xa6\x2f\x43\x98\xad\x96\x4a\x5e\x44\x59\x11\xa6\xbb\x74\x11\x97\x7d\x45\x94\xad\xe3\xc2\x34\xc4\x2e\x8d\xdc\x16\x6f\x7c\xeb\x64\x09\x23\x58\xc7\x85\xb4\x1e\x26\x4b\x0b\xcf\x90\x35\x8c\xb8\xfe\x0d\x05\x32\x59\xc2\xef\x70\xfb\xe1\x66\x32\x3e\x9b\xea\x30\x3e\x2e\x01\xb2\x04\xaa\x4a\x2b\x81\x77\xef\x00\x5f\xc4\xc2\x70\xb3\x5b\xc8\x20\xf8\x8a\x87\xfc\x21\xca\xe2\xfb\xfd\x6a\x86\x17\xb1\x1c\xcc\x39\xd1\x82\x1d\x49\x27\x08\x87\x51\xa9\x06\x3d\xe5\xb7\x3a\xa1\xe6\x45\x54\xc4\x33\xd6\x47\x1a\x5b\x72\xf5\x1f\x11\xec\x26\xe2\x78\x6a\x59\xcd\xdb\x39\xef\x88\xa8\xf9\x20\x78\xb8\x35\x0d\x61\x07\xeb\x7b\x2d\xa7\xd0\xe3\x49\xe0\x72\xfd\x02\xd3\x6f\x5d\x60\x9a\x86\x34\x21\xae\x31\x87\xbc\xfc\x2b\xab\x7d\x71\x95\xa9\xdf\x97\xe1\x55\xa6\x42\x74\x5f\x76\x3a\x1c\x61\x1a\xfa\x75\x27\xc7\xdb\xf8\x62\x65\x1a\xaf\x1e\xb3\x68\xbd\x8d\x60\x9f\x66\xbb\xcd\x06\x68\x7d\x98\x05\xce\x8e\x33\xc4\x1f\x6f\xcb\x59\xe4\xcd\x5e\x4f\x0c\x98\x9f\xe9\x10\x36\xaa\xa6\x5d\x48\x0c\x4c\xf9\xdf\xde\xf1\x9a\x8d\x19\x14\x6b\xb3\x96\xd7\xa1\xd8\xe3\x8c\xfb\x28\xcb\x92\x38\xb3\x4e\x2f\xfe\x11\x5e\x5c\x9d\x8e\x2f\xc2\x8f\x93\x8f\xe1\x1f\x93\xcb\xd3\x89\x38\x14\xc8\x2a\xb6\x4d\x0e\x6a\x26\x85\xd9\xe1\xaf\x2a\x43\x92\x26\x85\x53\x29\x8b\x32\x25\x46\xf0\xa2\x81\x76\x86\x6a\x29\x8a\x44\x82\x22\x78\x0b\xe3\xd3\xd3\xc9\x74\x3a\x99\xf2\x56\x6f\x04\xd4\x97\x79\xb8\xdf\xed\x36\xb0\x7f\x5c\x46\x45\x1c\x2a\x3a\xd6\x96\x08\x4f\x4d\xc4\xf9\xee\x91\x6a\xa5\xc8\xc3\x90\xee\x9a\xf7\xe8\x1c\x72\x24\x87\xa8\x5f\xe3\x90\x98\x94\xba\x43\xd5\x6f\x34\xc2\x87\x11\x7e\x32\x10\xc9\xb9\x03\x2b\x82\x1e\x24\xc4\x96\x73\x05\x47\xe4\xf5\x36\x39\x90\x59\x32\xe7\x23\x53\x1f\x54\xc4\x41\xe6\x09\x9f\x2f\xe6\xce\x78\xa9\x0c\xf1\x12\x81\x15\x4b\xc2\xf9\xd0\x9b\xa3\x7c\x2d\x5e\xef\x30\x02\x4d\x89\xa9\xb1\xe0\x7f\x55\x6a\x0e\x22\x35\x07\x78\x8b\x99\x39\x28\xb2\xf2\xd3\x64\x72\x98\x1d\x54\xe0\xa2\x61\x0b\x63\xb3\xa6\x5b\x5e\x82\xe6\xc7\x6a\x9b\x29\x67\x6c\x3e\x3b\xa8\xd3\xdc\xaf\x24\xa1\x3c\xcb\x99\x5a\xfa\xbd\xba\x45\xbe\x3c\xc4\xce\x6e\x71\x7c\xed\xbb\xce\x36\x39\x1c\xe7\x0e\x65\x2e\xb1\xbb\x24\x5e\xdf\x3f\xe9\x96\x04\x83\xe8\xbe\x5b\xb2\x58\xc6\x2b\x62\x1a\x06\xf9\x25\xf2\x3f\xb3\xbe\x5b\x95\x12\x8f\xcd\x2a\x46\x79\xae\xfa\xd5\x55\x8e\x09\xd2\xc9\x4d\xdd\xbf\x54\x21\xc5\x5b\xdc\xb3\x25\x52\xbc\x10\xfd\x95\x1a\x29\x2e\x32\x78\x12\xe4\x17\x2c\x61\xc8\x99\x13\x78\x2b\x37\x49\x31\x2c\x1c\x75\xbe\xd9\x15\x30\x82\x6d\x92\x8a\x15\x44\x3e\x8e\xff\x15\x5e\x7d\xba\xbd\xfe\x74\x3b\xb5\x21\x2a\x76\xdb\x64\x11\x26\xe9\xc2\x7a\xad\x36\x7c\xb4\xd4\x13\xb7\xb7\x46\xd9\xc9\xed\xcc\xc5\xe6\x39\x94\xef\x00\xfa\xb9\x62\x1d\xa7\x71\xc6\x97\xf7\x32\x5a\x87\x49\x11\x6f\x45\xd6\x70\xd7\xb1\x41\x5b\x72\xb8\x1b\x96\x2b\x6e\x93\xac\x1f\x8a\xe6\xe1\x40\x6e\x98\xf2\xe4\xd0\xa1\xae\x8e\x0c\x18\x2a\x7a\x4a\x77\xcb\x58\xec\xdd\xe8\x93\xef\x71\xad\x23\xc1\x33\x5e\x54\x52\x95\x99\x77\x65\x75\x39\x62\x04\xb2\xb8\xd8\x67\x29\xc7\x94\x47\x38\xe0\x81\x72\x34\xef\x55\xbf\xcb\xb5\x89\x95\x4c\x06\x37\x53\x36\x7f\xaf\x7d\xce\x9d\x1f\x8b\xed\xa2\xa1\x9b\xcb\x2b\x25\xa9\xc2\xe5\xe5\x07\xc7\x0a\x88\xc7\x5e\xd2\xde\xc0\x55\xf1\xed\xfe\x26\x5c\xa3\x2c\x6a\x3c\x46\x59\x9c\x16\x61\x92\x2e\xe3\x83\x2c\x3f\x6a\xa4\x77\x90\xd8\xd0\x18\x58\x02\xbf\x57\x1f\x92\xb1\xfe\x56\xd1\x70\xae\x74\xe7\x80\x5b\xed\x90\xd8\x65\x72\xea\xa3\x10\x49\x51\x6b\xef\x99\xc0\x0d\xd3\x68\xd7\xea\xfa\xf4\xbd\x06\x41\x8d\x76\x75\xe6\x96\x4b\xd4\x08\xfe\x87\xf7\xe2\xa2\xdc\x96\xda\x6f\xb8\xf6\x7f\xeb\xd6\x28\x11\x63\x83\x11\x74\x44\x33\x34\x9f\xcc\xff\x07\x00\x00\xff\xff\x6e\x60\xcd\x5f\x71\x20\x00\x00")

func bindataClKernel3ClBytes() ([]byte, error) {
	return bindataRead(
		_bindataClKernel3Cl,
		"cl/kernel3.cl",
	)
}



func bindataClKernel3Cl() (*asset, error) {
	bytes, err := bindataClKernel3ClBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "cl/kernel3.cl",
		size: 8305,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1629581470, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}


//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"cl/kernel1.cl": bindataClKernel1Cl,
	"cl/kernel2.cl": bindataClKernel2Cl,
	"cl/kernel3.cl": bindataClKernel3Cl,
}

//
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
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"cl": {Func: nil, Children: map[string]*bintree{
		"kernel1.cl": {Func: bindataClKernel1Cl, Children: map[string]*bintree{}},
		"kernel2.cl": {Func: bindataClKernel2Cl, Children: map[string]*bintree{}},
		"kernel3.cl": {Func: bindataClKernel3Cl, Children: map[string]*bintree{}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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
