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

var _bindataClKernel1Cl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x59\x71\x73\x9b\x3a\x12\xff\x1b\x3e\xc5\xbe\xe9\xb4\x83\x62\x5e\x8a\x04\x26\x78\x5c\x67\xc6\x97\xf8\xb5\x99\x4b\x93\x4c\x92\xf6\x6e\xce\xe3\x30\xc4\x26\x0e\x53\x1b\x72\x80\xfb\xd2\xf6\x72\x9f\xfd\x8d\x56\x12\x20\xa0\x49\x5f\x67\x42\x2d\xed\x6f\x77\xb5\xab\x9f\x56\x42\xbc\x7d\x0b\xe7\x0f\x65\xb2\x4d\xbe\xc7\x2b\xf0\xa0\xbc\xcf\xe3\x68\x55\x98\xe6\xab\x55\x7c\x97\xa4\x31\x9c\x5f\xcc\xce\x8e\x4e\xc3\xe3\xd9\xe7\x93\xa3\x59\x38\xfd\x78\x6c\x18\xce\x4f\x84\x67\x9f\x4f\x8e\x4f\xa6\x06\xfd\x89\x32\x48\xd8\xe7\xd9\xd9\xf1\xf9\xa5\x69\xbe\x4a\xee\xd2\x55\x7c\x07\xc7\xd3\xf7\xe1\xd5\xc9\x7f\x66\x95\x96\xea\x80\xc0\x0d\x82\xe1\xc8\x35\x5f\xc5\xe9\x2a\xb9\xab\x35\x4e\x4f\xde\x7f\xb8\xd6\x75\xea\x2e\x60\x3e\xa3\xee\xa8\x56\x92\x88\xd9\xf5\x87\xe9\xd5\x87\xf0\x78\x7a\x3d\xbd\x9a\x5d\x87\x17\xd3\xcb\xd9\xd9\xf5\x15\xb0\xa1\x5f\x63\xce\xce\x8f\x67\xc2\x88\xe5\x7b\x6f\x3d\x52\x4b\xfe\x38\xfb\x1c\x5e\x5c\x9e\x7c\x9c\x19\xce\xa3\x43\x1d\xc7\xa1\x23\x57\x93\x5a\x8f\x36\x7c\x23\x20\xff\x59\xd6\x23\x81\xbd\x5a\x0b\x6e\xc0\xfa\x46\x88\x66\xee\x72\x76\xfc\xe9\x68\x66\x7d\x45\x25\x6e\x41\xfd\x7d\xdd\x7f\xb4\xe1\xeb\xfe\x37\xc2\x9f\xdf\xf1\xf9\x67\x63\x28\x97\xe7\xd7\xa7\xbe\x17\x52\xe9\x31\x2a\xc2\xdd\x26\x4b\xd7\x56\x9e\x95\x51\x19\x5b\x55\xfb\x91\xd8\x60\xe1\x4f\xc2\x7d\x93\xb6\x01\x26\x0d\x34\xed\x59\xdf\x08\x0c\xc0\x65\x0d\x77\x57\xff\x9a\x5e\xf8\x1e\x8f\x07\xa0\x36\xce\x7f\x2c\xef\xa3\x3c\xb0\x1e\xc9\x7e\x71\xe0\x0f\x3d\x97\x51\x87\x34\xf5\x3e\x4c\xdd\x70\x48\x99\x55\x10\xb8\xcb\x72\xb0\x76\x49\x5a\x42\x02\x13\x08\xc6\x90\xc0\x6f\x13\x60\xc3\x31\x0c\x06\x09\x81\x1f\x50\xcc\x93\x05\x4c\xc0\xf9\x74\x3a\x06\x78\x82\x62\x1e\x60\xf3\x31\x70\xb4\x7f\x94\xcb\xbf\xc4\xcb\x65\xf4\x25\xbc\xa3\xbe\xe3\x58\x85\x0d\x41\xc3\xa9\x26\x8b\x6c\xc8\x76\x65\xf1\xbd\xe9\x3f\xe7\x66\xc7\x90\xc3\x3b\x60\xee\x98\xbb\x6e\xaa\x84\x79\xb6\x4b\x57\x5c\x31\x1f\x0c\x6c\x60\x43\x32\x86\xa7\x9f\x20\x98\xab\xcc\x9b\x66\x18\x2e\xb3\xb4\x28\xa3\xb4\x04\x4c\x0f\x60\x13\x2e\x8f\xe6\xcc\xe3\x81\xfc\x30\x39\x6d\x3a\xb1\xd8\xed\xee\xc0\x09\x98\xec\x0e\xf4\xee\x69\xab\x3b\x90\x7f\xfd\x46\xfe\xd1\xea\x0e\x74\x97\xba\x91\xa0\xdd\x2d\x45\xa3\x1e\xdb\x4e\x3d\x92\x56\x77\xd0\xe3\xb2\xc7\x88\xf4\xd1\x36\x12\xe8\xe3\xd6\xa7\xbd\xb7\x3b\x70\x82\x51\xff\xb8\xdd\xfe\xee\xbe\xc4\x72\x41\x6f\x06\x9d\xbe\x7c\x3b\xfd\xdd\x3f\xcd\x60\xc7\xf6\xb3\xd3\xe0\x60\x06\x9f\xc6\xa6\x59\x94\x51\x99\x2c\xe1\x6b\x96\xac\xfa\xb8\x87\x0c\xdb\x83\xc8\x96\x2c\x13\xbc\xd6\x5a\x92\xf6\x3f\x4c\x43\x76\x22\x29\xb7\x0e\x4c\x20\x9a\x3b\x0b\xb8\x81\x68\x3e\x14\xff\x51\xd9\xa4\xb2\xcd\xb0\x5d\xd5\x84\x68\xce\x44\xff\x81\x84\xc9\x26\x95\x6d\xc6\x16\x36\x50\x32\x6e\x79\xa2\xe8\x89\x0a\x8c\x2f\x55\x64\x93\xca\x36\xa3\x2d\x4f\xae\xe8\x0f\x24\x4c\x36\xa9\x6c\x33\xb7\xd7\x13\x43\x4f\x2f\x0e\x52\xf7\xe4\x89\xfe\x91\x84\xc9\x26\x95\x6d\xe6\xf5\x7a\x72\xd1\xd3\x8b\x83\xd4\x3d\xbd\x98\xec\x3e\x4f\x1e\x7a\x7a\x71\x90\xba\xa7\x17\x93\x2d\x3c\xe9\xae\xca\xed\x83\x9c\xa9\x9b\xad\xc3\x85\x62\xc4\x13\xd8\x7a\x63\xde\x18\x36\x1b\x54\x13\x51\x4d\xc6\x6a\x19\x6f\xfa\xa2\xe5\x08\x24\xd5\x5a\x9a\x8c\xd5\x32\x6c\x8a\x16\x45\xd9\x41\xb3\x41\x35\x11\xd5\x64\xac\x96\xf1\xa6\x2b\x5a\x0c\x65\x41\xb3\x41\x35\x11\xd5\x64\xac\x96\xf1\xa6\x27\x5a\x2e\xca\x46\xcd\x06\xd5\x44\x54\x93\xb1\x5a\x86\x42\x5e\xfc\xab\xed\x96\xa7\xc5\x06\xca\xc8\x58\xa6\x68\xd2\x9c\xbf\xd1\xc2\x06\xe6\x10\xe5\x4f\xd3\xc3\x55\xc6\x46\xa4\x8a\x56\x93\x52\xce\xd7\x03\x52\x0d\x4e\x33\x2b\x28\x16\x90\x6a\x96\x74\xc3\x0b\x1b\x5c\xe9\xb4\x63\x16\x97\x36\x25\x55\xfa\x35\xbb\x94\x2f\x47\xbe\x41\xaa\xbc\xea\x52\x1e\x0d\x48\xaf\xb4\x13\x0e\xea\x7a\xa4\xca\xbb\xee\x77\xb8\xb0\x41\xc6\xca\x49\x36\xd1\x57\xae\x0d\x4c\xc6\xda\x09\x15\xd7\xad\x32\xdb\x91\xf2\x15\x00\x32\xf9\xac\x3d\x33\x01\xb7\xeb\x12\x45\x19\x7d\x40\x38\x6d\x52\x48\xfd\x96\x94\x8f\x56\xfa\x6c\x0f\x16\xc3\x94\x49\xe8\x64\x08\x3d\xaa\xec\x06\x6d\xe9\x01\x77\xa9\xb2\x7b\xd0\x96\xe2\x62\x96\xd3\x46\x69\x4b\xca\x55\xc1\x27\x6a\x09\xe9\x9a\x9c\x0d\xa0\x42\x71\x34\x69\xb9\x7d\xb0\x41\xd6\x08\x59\x88\x86\x72\xc3\x18\x57\x3d\xbe\x2c\x17\x55\xad\x98\xc0\x6d\x52\x16\xf1\x26\x5e\x96\x58\xee\x6e\x90\x54\xfc\x97\xcd\x71\xa4\x59\x54\x2e\x8f\xe6\x39\x6a\x26\x77\x60\xe1\x1e\x05\x87\x40\xf9\x3e\x65\x1a\x6a\xb5\x34\xad\x51\x6e\xcd\x15\x86\x6c\x6e\x97\x8f\x5b\x31\xb5\x09\x64\x1c\xe8\x09\x8c\xcd\x55\x24\xd0\x6d\x03\xdd\xc5\xcd\x76\x68\x4b\xa3\x9e\x82\x79\x6d\x98\xb7\xb8\xd9\xfa\x88\xb0\xb7\x43\x0c\x41\x1b\xb2\x27\x87\x6c\xc8\x04\x0d\x79\x82\x0c\x43\x26\xc7\x17\x2d\x49\x87\xa6\xd9\x21\x1f\xe6\x01\x77\x3d\xe4\x0f\x5f\xf8\x57\xb5\xa0\x89\xf4\x39\x32\x10\x20\x9b\xeb\x28\xe4\x41\x1b\x79\xc0\x91\x23\x01\xb2\xb9\x8e\x42\x06\x6d\x64\x20\x63\x47\xb3\xa3\x0a\x37\x6a\xe3\x46\x32\xf8\x51\x1d\xbc\x16\x7d\xa0\xa2\x57\xe1\x53\x47\x44\xac\x12\x40\xa9\x6c\x2b\x8a\x69\x73\x8a\x14\xa1\x38\x4f\x54\x90\x84\xca\xb1\x54\x64\xd6\xf0\x48\x02\xea\x4a\xa4\x8d\xba\x35\xbe\x43\x05\x8a\x5c\xa0\x9e\x44\xda\xa8\x5b\xe3\x3b\x8c\xa0\x8a\x12\xd2\x85\xd7\x00\x77\x78\x41\x15\x31\x68\x83\x19\x75\x1e\x86\xad\x3c\xf8\x75\x1e\x3a\x5c\xa0\x48\x06\x8a\x93\x46\x91\x0e\xd4\x6f\xb8\xee\x30\x82\x22\x25\x68\x20\x91\x36\xea\xd6\xf8\x0e\x2f\x28\x12\x83\x8e\x24\xd2\x46\xdd\x1a\xdf\x61\x07\x55\xf4\x90\x2e\x46\x0d\x70\x87\x22\x54\x71\x84\x8e\x7a\xf2\xc0\x5a\x7c\x60\x35\x1f\x58\x87\x0f\x4c\x94\x0c\x9c\x29\x86\x7c\x60\x0d\x3e\xb0\x0e\x1f\x18\xf2\x81\xb9\x12\x69\xa3\x6e\x8d\xef\x96\x06\xe4\x03\xf3\x24\xd2\x46\xdd\x1a\xdf\xe1\x03\x53\x7c\x90\x2e\x1a\x7c\x60\x1d\x3e\x30\xc5\x07\xa6\xf8\x60\x18\xc6\x93\x89\x7f\x4f\xe6\x93\x69\x96\xdf\x1e\xe2\x55\x7c\x07\x45\x99\xef\x96\x25\x3f\x90\x63\x1d\x15\xcf\x62\xee\x32\x78\x0b\x45\xf2\x3d\xce\xee\xe4\x9b\xfa\x62\x6c\x3e\xc1\x7d\x54\xdc\xbb\x2c\x2c\xc7\xb5\x81\x5d\x9a\x64\x69\xa5\xef\x81\x38\xbf\x79\xc5\xdc\xf7\x5a\x26\x3c\x82\x25\x96\xbf\x07\xe0\xa3\x05\x49\xd2\x92\x60\x3d\x4f\xd2\x92\x89\x67\x17\xc1\x2a\x88\x27\x9e\x5d\x88\x57\x41\x02\xf1\xec\x42\x82\x0a\x42\x7d\x63\x15\x95\x91\x8a\xcd\xf7\xba\xb1\x99\x2a\x36\x90\xb1\x15\x73\xca\x82\x9e\xf4\x74\x32\xd0\x85\x79\x35\x2e\x90\xb8\xa0\x0f\x17\xd4\xa9\x02\xee\x16\xb3\xa5\xc3\xb4\x74\x49\x0c\xeb\x01\x35\x32\x26\x51\x9d\x91\xe9\x49\x93\xa8\xce\xb8\x5a\x79\x13\x28\xea\xf7\xc0\xa8\x5f\xb3\x85\xb2\xe0\x19\xba\x28\xba\x31\xc7\xe9\x49\xa8\xc6\x15\x1d\xd2\x4b\x96\x0e\x84\x75\xa9\xd0\xc1\x68\x5c\x50\x21\x75\x50\xcd\x90\x98\xe3\x60\x48\xaf\x92\xbb\xd6\x3d\xe2\x6f\x93\xde\x8b\x47\x33\x0c\xa3\xb2\xcc\x93\xdb\x5d\x19\x87\xa1\x65\xe5\xf1\x7f\x57\xe1\x9f\x59\xfe\x25\x5c\xe7\xd9\xee\x21\xe4\x7e\xac\xf7\x97\xe7\x9f\x2e\xf0\x8a\xcf\x06\xca\xdf\x88\x08\x51\x37\x85\x61\xf8\x25\xce\xd3\x78\x23\x5e\xbc\x8b\x38\xda\x58\xa6\x11\x86\xeb\x4d\x76\x1b\xf1\xce\x4d\x54\x26\x9b\x18\xe7\x63\x0f\xf2\xb8\x28\xf3\x64\x59\xc2\x3a\xcc\x76\xe5\xc3\xae\xb4\x39\xb6\xba\x05\x52\x2b\x58\xbc\x94\xef\xc1\x3a\xbc\x8f\xa3\x55\x9c\xdb\x0d\x8b\xd5\xbc\xd5\xa0\x55\xb4\xa6\x2f\x43\x98\xad\x96\x4a\x51\x46\x79\x19\xa6\x59\xba\x8c\xab\xbe\x32\xca\xd7\x31\x1f\x0e\xd2\x3a\xba\xcd\xf2\xd2\x34\xc4\x96\xcd\x77\x71\xec\xc0\xab\x01\x23\x8f\xcb\x5d\x9e\x8e\x79\xa9\x92\x68\xf1\x6a\xb8\x4e\x56\x30\x81\x75\x5c\xca\x61\x84\xc9\xca\xc2\xc3\x66\x03\x23\xee\x89\x43\x81\x4c\x56\xf0\x1a\xae\x3f\x5c\xce\xa6\xc7\x57\x3a\x8c\x07\x20\x40\x96\x40\xd5\xf9\x27\x70\x78\x08\xf8\xb2\x15\x86\x9b\x6c\x29\xa3\xe5\xa5\x01\x8a\xfb\x28\x8f\x6f\x77\x77\x73\xbc\xb1\xe5\x60\x4e\x9e\x0e\x6c\x4f\x3a\x41\x38\x4c\x2a\x35\x18\x28\xbf\xf5\x51\xb6\x28\xa3\x32\x9e\xb3\x21\xf2\xdd\x92\x65\x62\x8f\x60\x37\x11\xe7\x58\xcb\x6a\x5f\xe3\x79\x7b\x44\x4d\x1c\xc1\x53\xb0\x69\x08\x3b\xb8\x11\x34\x92\x0f\x03\x9e\x04\x2e\xd7\x6f\x3a\xfd\xce\x4d\xa7\x69\x48\x13\xe2\xbe\x53\x26\x5f\x74\x0d\xc5\x9d\xa7\x7e\xb1\x86\x77\x9e\x0a\xd1\x7f\x2b\xea\x70\x84\x69\xe8\xf7\xa2\x1c\x6f\xe3\x1b\x98\x69\xbc\x7a\xc8\xa3\xf5\x36\x82\x5d\x9a\x67\x9b\x0d\xd0\xe6\x30\x4b\x9c\x1d\x67\x8c\x3f\xde\x55\xb3\xc8\x9b\x83\x81\x18\x30\xa7\x0d\xc2\x26\xf5\xb4\x0b\x89\x81\x29\xff\xfd\x90\x17\x77\xcc\xa0\x58\xc4\x8d\xbc\x8e\xc5\x66\x68\xdc\x46\x79\x9e\xc4\xb9\x75\x74\xfa\xcf\xf0\xf4\xfc\x68\x7a\x1a\x7e\x9c\x7d\x0c\xff\x98\x9d\x1d\xcd\xc4\xe9\x41\x14\xc5\x6d\xf2\xa8\x26\xf2\xf7\x43\x59\x4b\x6a\xaa\xbd\x01\x71\x92\xf8\x25\x63\x90\xa4\x49\xe9\xe8\xd6\x0a\x39\xa2\x17\x0d\x74\x33\xd6\x48\x59\x24\x12\x16\xc1\x3b\x98\x1e\x1d\xcd\xae\xae\x66\x57\xbc\x35\x98\x88\xa3\xb1\x61\x18\xb7\x59\xb6\x81\xdd\xc3\x2a\x2a\xe3\x50\xb1\xb3\xb1\x62\x78\xa6\x22\x4e\x7f\x97\xd4\x0b\x47\x1e\xa2\x74\xcf\xbc\x47\xa7\x94\x23\x29\x15\x34\x18\x25\xa6\xa8\xe9\x4f\xf5\x1b\xad\xe0\x61\x82\x5f\x1a\x44\x6a\x6e\xc0\x8a\x60\x00\x09\xb1\xe5\xcc\xc1\x1e\x79\xb3\x4d\x1e\xc9\x3c\x59\xf0\x81\xa9\xef\x30\xe2\xfc\xf3\x84\xcf\x17\x33\x67\xbc\x54\xbd\x78\xc1\xc0\x42\x27\xe1\x7c\xe8\xed\x51\xbe\x11\x6f\x85\x18\x81\xa6\xc4\xd4\x58\xc4\xf9\x12\xd9\xc2\x03\xda\x26\x8f\xb6\x40\xcc\xdb\xb6\x0e\x0f\x81\x2e\xf6\xdb\x5c\x52\xa7\xba\xe7\xe3\xa9\x4e\x73\xa6\x96\x49\xd6\x34\xc4\x79\x2f\x76\x41\xab\xf1\x55\x67\x9b\x3c\xee\x6f\x32\x62\x43\xab\xef\x3e\x21\xe4\x97\x18\xf8\xcc\xa2\xeb\x94\x2f\x3c\xf4\xaa\xf1\xc9\x53\xd1\xaf\x2e\x3d\x0c\x4e\xa7\x18\x75\xff\x56\xd9\x12\xef\x60\xcf\xd6\x2d\xf1\x3a\xf3\x77\x0a\x97\xb8\x86\xe0\x49\x90\xdf\x9f\x84\x21\x67\x41\xe0\x9d\xdc\xe2\xea\xe2\xd4\xd8\xd3\xea\x4d\x0d\x63\xc7\x98\x8a\x4d\x56\xc2\x04\xb6\x49\x2a\x68\x4e\x3e\x4e\xff\x1d\x9e\x7f\xba\xbe\xf8\x74\x7d\x65\x43\x54\x66\xdb\x64\x19\x26\xe9\xd2\x7a\xa3\x36\x73\xf4\x33\x10\x37\xb3\x46\xd5\xc9\xed\x2c\xc4\x7e\x37\x96\xe7\x7b\xfd\xcc\xb0\x8e\xd3\x38\xe7\x6b\x70\x15\xad\xc3\xa4\x8c\xb7\x22\xa7\xb8\x51\xd8\xa0\xad\x0b\xdc\xc0\xaa\x65\xb1\x49\xd6\xf7\x65\x7b\xe3\x97\x7b\x9c\x3c\x15\xf4\xa8\xab\xe3\x00\x46\x8e\x9e\xd2\x6c\x15\x8b\xed\x16\x7d\xf2\x6d\xa9\xb3\x8b\x3f\xe3\x45\xa5\x5c\x99\x39\xac\x4a\xc0\x1e\x23\xa0\x32\x6b\x1a\xd5\xf1\x0c\x78\xa0\x1c\xcd\x7b\xd5\xef\x7d\x75\xbc\xc3\x72\x23\x83\x9b\x2b\x9b\xaf\x1b\x9f\x6a\x17\xfb\xe2\xad\xa0\xa5\x5b\xc8\xeb\x22\xa9\xc2\xe5\xd5\xc7\xc4\x1a\x88\x47\x5a\xd2\xdd\x73\x55\x81\xec\xff\xde\xdb\x20\x34\x6a\x3c\x44\x79\x9c\x96\x61\x92\xae\x62\x55\x4b\xd4\x48\x6f\x20\xb1\xa1\x35\xb0\x04\x5e\xd7\x1f\x89\xb1\x48\xd6\xd1\x70\xae\xf4\xe7\x80\x5b\xed\x91\xd8\x55\x72\x9a\xa3\x10\x49\x51\x2b\xf3\x99\xc0\x0d\xd3\xe8\x16\xd4\xe6\xf4\xbd\x01\x41\x8d\x6e\x09\xe5\x96\x2b\xd4\x04\xfe\x8f\xc7\x30\x51\x3e\x2b\xed\xb7\x5c\xfb\x7f\x4d\x6b\x94\x88\xb1\xc1\x04\x7a\xa2\x19\x9b\x4f\xe6\x5f\x01\x00\x00\xff\xff\xc3\x2f\x00\x37\x4d\x20\x00\x00")

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
		size: 8269,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1629589672, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataClKernel2Cl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x59\xff\x6f\x9b\xba\x16\xff\x19\xfe\x8a\x73\x35\x6d\x82\x86\xdb\x61\x43\x28\x51\x96\x4a\x79\x6d\xee\x56\xbd\xae\xad\xda\x6e\xef\xe9\x45\x29\xa2\x09\x4d\xd1\x12\xe8\x03\x67\xb7\xdb\x5e\xdf\xdf\x7e\xe5\x63\x1b\x30\xb0\x76\x77\x52\x59\xec\xf3\xf1\xf9\xe6\x8f\x8f\x8d\x79\xfb\x16\xce\x1f\x58\xba\x4d\xbf\x27\x2b\x08\x81\xdd\x17\x49\xbc\x2a\x4d\xf3\xd5\x2a\xb9\x4b\xb3\x04\xce\x2f\x66\x67\x47\xa7\xd1\xf1\xec\xf3\xc9\xd1\x2c\x9a\x7e\x3c\x36\x0c\xf7\x27\xc2\xb3\xcf\x27\xc7\x27\x53\x83\xfc\x64\x30\x48\xd8\xe7\xd9\xd9\xf1\xf9\xa5\x69\xbe\x4a\xef\xb2\x55\x72\x07\xc7\xd3\xf7\xd1\xd5\xc9\x7f\x66\xd5\x28\xd5\x01\xa1\x17\x86\xc3\x91\x67\xbe\x4a\xb2\x55\x7a\x57\x8f\x38\x3d\x79\xff\xe1\x5a\x1f\x53\x77\x01\x0d\x28\xf1\x46\xf5\x20\x89\x98\x5d\x7f\x98\x5e\x7d\x88\x8e\xa7\xd7\xd3\xab\xd9\x75\x74\x31\xbd\x9c\x9d\x5d\x5f\x01\x1d\x06\x35\xe6\xec\xfc\x78\x26\x94\x58\x81\xff\xd6\xb7\x6b\xc9\x1f\x67\x9f\xa3\x8b\xcb\x93\x8f\x33\xc3\x7d\x74\x89\xeb\xba\x64\xe4\x69\x52\xeb\xd1\x81\x6f\x36\xc8\x7f\x96\xf5\x68\xc3\x5e\x3d\x0a\x6e\xc0\xfa\x66\xdb\x9a\xba\xcb\xd9\xf1\xa7\xa3\x99\xf5\x15\x07\x71\x0d\xea\xef\xeb\xfe\xa3\x03\x5f\xf7\xbf\xd9\xfc\xf9\x1d\x9f\x7f\x36\x5c\xb9\x3c\xbf\x3e\x0d\xfc\x88\x48\x8b\x71\x19\xed\x36\x79\xb6\xb6\x8a\x9c\xc5\x2c\xb1\xaa\xf6\xa3\xed\x80\x85\x3f\x6d\x6e\xdb\x6e\x2b\xa0\x52\x41\x53\x9f\xf5\xcd\x86\x01\x78\xb4\x61\xee\xea\x5f\xd3\x8b\xc0\xe7\xf1\x00\xd4\xca\xf9\x8f\xe5\x7d\x5c\x84\xd6\xa3\xbd\x5f\x1e\x04\x43\xdf\xa3\xc4\xb5\x9b\xe3\x3e\x4c\xbd\x68\x48\xa8\x55\xda\x70\x97\x17\x60\xed\xd2\x8c\x41\x0a\x13\x08\xc7\x90\xc2\x6f\x13\xa0\xc3\x31\x0c\x06\xa9\x0d\x3f\xa0\x9c\xa7\x0b\x98\x80\xfb\xe9\x74\x0c\xf0\x04\xe5\x3c\xc4\xe6\x63\xe8\x6a\xff\x08\x97\x7f\x49\x96\xcb\xf8\x4b\x74\x47\x02\xd7\xb5\x4a\x07\xc2\x86\x51\x4d\x16\x3b\x90\xef\x58\xf9\xbd\x69\xbf\xe0\x6a\xc7\x50\xc0\x3b\xa0\xde\x98\x9b\x6e\x0e\x89\x8a\x7c\x97\xad\xf8\xc0\x62\x30\x70\x80\x0e\xed\x31\x3c\xfd\x04\x41\x3d\xa5\xde\x34\xa3\x68\x99\x67\x25\x8b\x33\x06\x98\x1e\xc0\x26\x5c\x1e\xcd\xa9\xcf\x03\xf9\x61\x72\xda\x74\x62\x71\xda\xdd\xa1\x1b\x52\xd9\x1d\xea\xdd\xd3\x56\x77\x28\xff\xfa\x95\xfc\xa3\xd5\x1d\xea\x26\x75\x25\x61\xbb\x5b\x8a\x46\x3d\xba\xdd\xda\x93\x56\x77\xd8\x63\xb2\x47\x89\xb4\xd1\x56\x12\xea\x7e\xeb\xd3\xde\xdb\x1d\xba\xe1\xa8\xdf\x6f\xaf\xbf\xbb\x2f\xb1\x5c\xd0\x9b\x41\xb7\x2f\xdf\x6e\x7f\xf7\x4f\x33\xd8\xd1\xfd\xec\x34\xb8\x98\xc1\xa7\xb1\x69\x96\x2c\x66\xe9\x12\xbe\xe6\xe9\xaa\x8f\x7b\xc8\xb0\x3d\x88\x1d\xc9\x32\xc1\x6b\xad\x25\x69\xff\xc3\x34\x64\x27\x92\x72\xeb\xc2\x04\xe2\xb9\xbb\x80\x1b\x88\xe7\x43\xf1\x1f\x91\x4d\x22\xdb\x14\xdb\x55\x4d\x88\xe7\x54\xf4\x1f\x48\x98\x6c\x12\xd9\xa6\x74\xe1\x00\xb1\xc7\x2d\x4b\x04\x2d\x11\x81\x09\xe4\x10\xd9\x24\xb2\x4d\x49\xcb\x92\x27\xfa\x43\x09\x93\x4d\x22\xdb\xd4\xeb\xb5\x44\xd1\xd2\x8b\x4e\xea\x96\x7c\xd1\x3f\x92\x30\xd9\x24\xb2\x4d\xfd\x5e\x4b\x1e\x5a\x7a\xd1\x49\xdd\xd2\x8b\xc9\xee\xb3\xe4\xa3\xa5\x17\x9d\xd4\x2d\xbd\x98\x6c\x61\x49\x37\xc5\xb6\x0f\x72\xa6\x6e\xb6\x2e\x17\x0a\x8f\x27\xb0\xf5\xc7\xbc\x31\x6c\x36\x88\x26\x22\x9a\x8c\xd6\x32\xde\x0c\x44\xcb\x15\x48\xa2\xb5\x34\x19\xad\x65\xd8\x14\x2d\x82\xb2\x83\x66\x83\x68\x22\xa2\xc9\x68\x2d\xe3\x4d\x4f\xb4\x28\xca\xc2\x66\x83\x68\x22\xa2\xc9\x68\x2d\xe3\x4d\x5f\xb4\x3c\x94\x8d\x9a\x0d\xa2\x89\x88\x26\xa3\xb5\x0c\x85\xbc\xf8\x57\xdb\x2d\x4f\x8b\x03\x84\xda\x63\x99\xa2\x49\x73\xfe\x46\x0b\x07\xa8\x6b\x2b\x7b\xda\x38\x5c\x65\x74\x64\x57\xd1\x6a\x52\xc2\xf9\x7a\x60\x57\xce\x69\x6a\x05\xc5\x42\xbb\x9a\x25\x5d\xf1\xc2\x01\x4f\x1a\xed\xa8\xc5\xa5\x4d\xec\x2a\xfd\x9a\x5e\xc2\x97\x23\xdf\x20\x55\x5e\x75\x29\x8f\x06\xa4\x55\xd2\x09\x07\xc7\xfa\x76\x95\x77\xdd\xee\x70\xe1\x80\x8c\x95\x93\x6c\xa2\xaf\x5c\x07\xa8\x8c\xb5\x13\x2a\xae\x5b\xa5\xb6\x23\xe5\x2b\x00\x64\xf2\x69\x7b\x66\x42\xae\xd7\xb3\x15\x65\x74\x87\x70\xda\xa4\x90\x04\x2d\x29\xf7\x56\xda\x6c\x3b\x8b\x61\xca\x24\x74\x32\x84\x16\x55\x76\xc3\xb6\xf4\x80\x9b\x54\xd9\x3d\x68\x4b\x71\x31\xcb\x69\x23\xa4\x25\xe5\x43\x21\xb0\xd5\x12\xd2\x47\x72\x36\x80\x0a\xc5\xd5\xa4\x6c\xfb\xe0\x80\xac\x11\xb2\x10\x0d\xe5\x86\x31\xae\x7a\x02\x59\x2e\xaa\x5a\x31\x81\xdb\x94\x95\xc9\x26\x59\x32\x2c\x77\x37\x48\x2a\xfe\xcb\xe1\x38\xbb\x59\x54\x2e\x8f\xe6\x05\x8e\x4c\xef\xc0\xc2\x3d\x0a\x0e\x81\xf0\x7d\xca\x34\xd4\x6a\x69\x6a\x23\x5c\x9b\x27\x14\x39\x5c\x2f\xf7\x5b\x31\xb5\x09\xa4\x1c\xe8\x0b\x8c\xc3\x87\x48\xa0\xd7\x06\x7a\x8b\x9b\xed\xd0\x91\x4a\x7d\x05\xf3\xdb\x30\x7f\x71\xb3\x0d\x10\xe1\x6c\x87\x18\x82\xe6\xb2\x2f\x5d\x36\x64\x82\x86\x3c\x41\x86\x21\x93\x13\x88\x96\xa4\x43\x53\xed\x90\xbb\x79\xc0\x4d\x0f\xf9\x23\x10\xf6\x55\x2d\x68\x22\x03\x8e\x0c\x05\xc8\xe1\x63\x14\xf2\xa0\x8d\x3c\xe0\xc8\x91\x00\x39\x7c\x8c\x42\x86\x6d\x64\x28\x63\x47\xb5\xa3\x0a\x37\x6a\xe3\x46\x32\xf8\x51\x1d\xbc\x16\x7d\xa8\xa2\x57\xe1\x13\x57\x44\xac\x12\x40\x88\x6c\x2b\x8a\x69\x73\x8a\x14\x21\x38\x4f\x44\x90\x84\x48\x5f\x2a\x32\x6b\x78\x24\x01\xf1\x24\xd2\xc1\xb1\x35\xbe\x43\x05\x82\x5c\x20\xbe\x44\x3a\x38\xb6\xc6\x77\x18\x41\x14\x25\xa4\x09\xbf\x01\xee\xf0\x82\x28\x62\x90\x06\x33\xea\x3c\x0c\x5b\x79\x08\xea\x3c\x74\xb8\x40\x90\x0c\x04\x27\x8d\x20\x1d\x48\xd0\x30\xdd\x61\x04\x41\x4a\x90\x50\x22\x1d\x1c\x5b\xe3\x3b\xbc\x20\x48\x0c\x32\x92\x48\x07\xc7\xd6\xf8\x0e\x3b\x88\xa2\x87\x34\x31\x6a\x80\x3b\x14\x21\x8a\x23\x64\xd4\x93\x07\xda\xe2\x03\xad\xf9\x40\x3b\x7c\xa0\xa2\x64\xe0\x4c\x51\xe4\x03\x6d\xf0\x81\x76\xf8\x40\x91\x0f\xd4\x93\x48\x07\xc7\xd6\xf8\x6e\x69\x40\x3e\x50\x5f\x22\x1d\x1c\x5b\xe3\x3b\x7c\xa0\x8a\x0f\xd2\x44\x83\x0f\xb4\xc3\x07\xaa\xf8\x40\x15\x1f\x0c\xc3\x78\x32\xf1\xef\xc9\x7c\x32\x4d\xf6\xed\x21\x59\x25\x77\x50\xb2\x62\xb7\x64\xfc\x40\x8e\x75\x54\x3c\xcb\xb9\x47\xe1\x2d\x94\xe9\xf7\x24\xbf\x93\x6f\xea\x8b\xb1\xf9\x04\xf7\x71\x79\xef\xd1\x88\x8d\x6b\x05\xbb\x2c\xcd\xb3\x6a\xbc\x0f\xe2\xfc\xe6\x97\xf3\xc0\x6f\xa9\xf0\x6d\x2c\xb1\xfc\x3d\x00\x1f\x2d\x48\x9a\x31\x1b\xeb\x79\x9a\x31\x2a\x9e\x5d\x04\xad\x20\xbe\x78\x76\x21\x7e\x05\x09\xc5\xb3\x0b\x09\x2b\x08\x09\x8c\x55\xcc\x62\x15\x5b\xe0\x77\x63\x33\x55\x6c\x20\x63\x2b\xe7\x84\x86\x3d\xe9\xe9\x64\xa0\x0b\xf3\x6b\x5c\x28\x71\x61\x1f\x2e\xac\x53\x05\xdc\x2c\x66\x4b\x87\x69\xe9\x92\x18\xda\x03\x6a\x64\x4c\xa2\x3a\x9e\xe9\x49\x93\xa8\x8e\x5f\xad\xbc\x09\x14\x09\x7a\x60\x24\xa8\xd9\x42\x68\xf8\x0c\x5d\x14\xdd\xa8\xeb\xf6\x24\x54\xe3\x8a\x0e\xe9\x25\x4b\x07\x42\xbb\x54\xe8\x60\x34\x2e\xa8\x90\x3a\xa8\x66\x48\xd4\x75\x31\xa4\x57\xe9\x5d\xeb\x1e\xf1\xb7\x49\xef\xc5\xa3\x19\x45\x31\x63\x45\x7a\xbb\x63\x49\x14\x59\x56\x91\xfc\x77\x15\xfd\x99\x17\x5f\xa2\x75\x91\xef\x1e\x22\x6e\xc7\x7a\x7f\x79\xfe\xe9\x02\xaf\xf8\x1c\x20\xfc\x8d\xc8\xb6\xd5\x4d\x61\x14\x7d\x49\x8a\x2c\xd9\x88\x17\xef\x32\x89\x37\x96\x69\x44\xd1\x7a\x93\xdf\xc6\xbc\x73\x13\xb3\x74\x93\xe0\x7c\xec\x41\x91\x94\xac\x48\x97\x0c\xd6\x51\xbe\x63\x0f\x3b\xe6\x70\x6c\x75\x0b\xa4\x56\xb0\x78\x29\xdf\x83\x75\x74\x9f\xc4\xab\xa4\x70\x1a\x1a\xab\x79\xab\x41\xab\x78\x4d\x5e\x86\x50\x47\x2d\x95\x92\xc5\x05\x8b\xb2\x3c\x5b\x26\x55\x1f\x8b\x8b\x75\xc2\xdd\x41\x5a\xc7\xb7\x79\xc1\x4c\x43\x6c\xd9\x7c\x17\xc7\x0e\xbc\x1a\x30\x8a\x84\xed\x8a\x6c\xcc\x4b\x95\x44\x8b\x57\xc3\x75\xba\x82\x09\xac\x13\x26\xdd\x88\xd2\x95\x85\x87\xcd\x06\x46\xdc\x13\x47\x02\x99\xae\xe0\x35\x5c\x7f\xb8\x9c\x4d\x8f\xaf\x74\x18\x0f\x40\x80\x2c\x81\xaa\xf3\x6f\xc3\xe1\x21\xe0\x7b\x52\x14\x6d\xf2\xa5\x8c\x96\x97\x06\x28\xef\xe3\x22\xb9\xdd\xdd\xcd\xf1\xc6\x96\x83\x39\x79\x3a\xb0\x3d\x69\x04\xe1\x30\xa9\x86\xc1\x40\xd9\xad\x8f\xb2\x25\x8b\x59\x32\xa7\x43\xe4\xbb\x25\xcb\xc4\x9e\x8d\xdd\xb6\x38\xc7\x5a\x56\xfb\x1a\xcf\xdf\xb3\xd5\xc4\xd9\x78\x0a\x36\x0d\xa1\x07\x37\x82\x46\xf2\x61\xc0\x93\xc0\xe5\xfa\x4d\x67\xd0\xb9\xe9\x34\x0d\xa9\x42\xdc\x77\xca\xe4\x8b\xae\xa1\xb8\xf3\xd4\x2f\xd6\xf0\xce\x53\x21\xfa\x6f\x45\x5d\x8e\x30\x0d\xfd\x5e\x94\xe3\x1d\x7c\x03\x33\x8d\x57\x0f\x45\xbc\xde\xc6\xb0\xcb\x8a\x7c\xb3\x01\xd2\x74\x93\xe1\xec\xb8\x63\xfc\xf1\xae\x9a\x45\xde\x1c\x0c\x84\xc3\x9c\x36\x08\x9b\xd4\xd3\x2e\x24\x06\xa6\xfc\xf7\x43\x5e\xdc\x31\x83\x62\x11\x37\xf2\x3a\x16\x9b\xa1\x71\x1b\x17\x45\x9a\x14\xd6\xd1\xe9\x3f\xa3\xd3\xf3\xa3\xe9\x69\xf4\x71\xf6\x31\xfa\x63\x76\x76\x34\x13\xa7\x07\x51\x3a\xb7\xe9\xa3\x9a\xc8\xdf\x0f\x65\x15\xad\xa9\xf6\x06\x3c\x54\xf8\x4b\xca\x20\xcd\x52\xe6\xea\xda\x4a\xe9\xd1\x8b\x0a\xba\x19\x6b\xa4\x2c\x16\x09\x8b\xe1\x1d\x4c\x8f\x8e\x66\x57\x57\xb3\x2b\xde\x1a\x4c\xc4\x8b\x81\x61\x18\xb7\x79\xbe\x81\xdd\xc3\x2a\x66\x49\xa4\xd8\xd9\x58\x31\x3c\x53\x31\xa7\x3f\xb5\xeb\x85\x23\x0f\x51\xba\x65\xde\xa3\x53\xca\x95\x94\xf2\x1b\x8c\x12\x53\xd4\xb4\xa7\xfa\x8d\x56\xf0\x30\xc1\x2f\x0d\x22\x35\x37\x60\xc5\x30\x80\xd4\x76\xe4\xcc\xc1\x9e\xfd\x66\x9b\x3e\xda\xf3\x74\xc1\x1d\x53\xdf\x61\xc4\xf9\xe7\x09\x9f\x2f\x66\xce\x78\xa9\x7a\xf1\x82\x81\x85\x4e\xc2\xb9\xeb\x6d\x2f\xdf\x88\xb7\x42\x8c\x40\x1b\x44\x95\x2f\xe2\x7c\x89\x6c\xe1\x01\x6d\xd3\x47\x47\x20\xe6\x6d\x5d\x87\x87\x40\x16\xfb\x6d\x2e\xa9\x53\xdd\xf3\xf1\x54\xa7\x39\xb3\x95\xc9\x5a\x8f\x70\x40\x7d\xcc\xe1\xd9\xfb\x25\x82\x3d\xb3\xa6\x3a\xd5\x09\xcf\xb4\xca\xbc\x3c\xf4\xfc\xea\xca\x42\xdf\x75\x06\x11\xef\x6f\x55\x25\xf1\x8a\xf5\x6c\x59\x12\x6f\x2b\x7f\xa7\x2e\x89\x5b\x06\x9e\x04\xf9\x79\x49\x28\x72\x17\x36\xbc\x93\x3b\x58\x5d\x7b\x1a\x5b\x56\xbd\x67\x19\x4f\x86\xa1\x56\x7a\xb9\xc9\x19\x4c\x60\x9b\x66\x82\xc7\xf6\xc7\xe9\xbf\xa3\xf3\x4f\xd7\x17\x9f\xae\xaf\x1c\x88\x59\xbe\x4d\x97\x51\x9a\x2d\xad\x37\x6a\xb7\x46\x4b\x03\x71\xf5\x6a\x54\x9d\x5c\xcf\x42\x6c\x68\x63\x79\x80\xd7\x0f\x05\xeb\x24\x4b\x0a\xbe\xc8\x56\xf1\x3a\x4a\x59\xb2\x15\x59\xc5\x9d\xc0\x01\x8d\xf8\xb8\x43\x55\xbc\xdf\xa4\xeb\x7b\xd6\xde\xd9\xe5\x26\x26\xb7\xfd\x9e\xe1\x6a\xbf\xc7\xd8\xd1\x52\x96\xaf\x12\xb1\x9f\xa2\x4d\xbe\xef\x74\xb6\xe9\x67\xac\xa8\xa4\x2b\x35\x87\xd5\x1a\xdf\xa3\x36\xa8\xdc\x9a\x46\x75\xfe\x02\x1e\x28\x47\xf3\x5e\xf5\x7b\x5f\x9d\xdf\xb0\x9e\xc8\xe0\xe6\x4a\xe7\xeb\xc6\xb7\xd8\xc5\xbe\x38\xf6\xb7\xc6\x96\xf2\x3e\x48\x0e\xe1\xf2\xea\x6b\x61\x0d\xc4\x33\xab\xdd\xdd\x54\x55\x05\xec\xff\xa0\xdb\xa0\x34\x8e\x78\x88\x8b\x24\x63\x51\x9a\xad\x12\x55\x2c\x94\xa7\x37\x90\x3a\xd0\x72\x2c\x85\xd7\xf5\x57\x60\xac\x82\x75\x34\x9c\x2b\xfd\x39\xe0\x5a\x7b\x24\x4e\x95\x9c\xa6\x17\x22\x29\x6a\x6d\x3e\x13\xb8\x61\x1a\xdd\x8a\xd9\x9c\xbe\x37\x20\xa8\xd1\xad\x91\x5c\x73\x85\x9a\xc0\xff\xf1\x52\x5b\xd4\xc7\x6a\xf4\x5b\x3e\xfa\x7f\x4d\x6d\xc4\x16\xbe\xc1\x04\x7a\xa2\x19\x9b\x4f\xe6\x5f\x01\x00\x00\xff\xff\x66\x46\xc7\xb8\x2e\x20\x00\x00")

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
		size: 8238,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1629589690, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataClKernel3Cl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x59\x71\x73\x9b\x3a\x12\xff\x1b\x3e\xc5\xbe\xe9\xbc\x0e\x8a\xb9\x16\x09\x4c\xf0\xb8\xee\x8c\x2f\xf1\x7b\xcd\x5c\x9a\x64\x92\xb4\x77\x73\x1e\x87\x21\x36\x76\x98\xda\x90\x03\xdc\xe7\xf6\x5d\xee\xb3\xdf\x68\x25\x01\x02\x9a\xf4\x75\x26\xd4\xd2\xfe\x76\x57\xbb\xfa\x69\x25\xc4\xdb\xb7\x70\xf9\x58\x26\xbb\xe4\x7b\xbc\x02\x06\xe5\x43\x1e\x47\xab\xc2\x34\x5f\xad\xe2\x75\x92\xc6\x70\x79\x35\xbb\x38\x39\x0f\x4f\x67\x9f\xcf\x4e\x66\xe1\xf4\xe3\xa9\x61\x38\x3f\x10\x5e\x7c\x3e\x3b\x3d\x9b\x1a\xf4\x07\xca\x20\x61\x9f\x67\x17\xa7\x97\xd7\xa6\xf9\x2a\x59\xa7\xab\x78\x0d\xa7\xd3\xdf\xc3\x9b\xb3\x7f\xcf\x2a\x2d\xd5\x01\x81\x1b\x04\xc3\x91\x6b\xbe\x8a\xd3\x55\xb2\xae\x35\xce\xcf\x7e\xff\x70\xab\xeb\xd4\x5d\xc0\x7c\x46\xdd\x51\xad\x24\x11\xb3\xdb\x0f\xd3\x9b\x0f\xe1\xe9\xf4\x76\x7a\x33\xbb\x0d\xaf\xa6\xd7\xb3\x8b\xdb\x1b\x60\x43\xbf\xc6\x5c\x5c\x9e\xce\x84\x11\xcb\xf7\xde\x7a\xa4\x96\xfc\x76\xf1\x39\xbc\xba\x3e\xfb\x38\x33\x9c\x83\x43\x1d\xc7\xa1\x23\x57\x93\x5a\x07\x1b\xbe\x11\x90\xff\x2c\xeb\x40\xe0\xa8\xd6\x82\x3b\xb0\xbe\x11\xa2\x99\xbb\x9e\x9d\x7e\x3a\x99\x59\x5f\x51\x89\x5b\x50\x7f\x5f\xdf\x1c\x6c\xf8\xfa\xe6\x1b\xe1\xcf\xef\xf8\xfc\xa3\x31\x94\xeb\xcb\xdb\x73\xdf\x0b\xa9\xf4\x18\x15\xe1\x7e\x9b\xa5\x1b\x2b\xcf\xca\xa8\x8c\xad\xaa\x7d\x20\x36\x58\xf8\x93\x70\xdf\xa4\x6d\x80\x49\x03\x4d\x7b\xd6\x37\x02\x03\x70\x59\xc3\xdd\xcd\x3f\xa7\x57\xbe\xc7\xe3\x01\xa8\x8d\xf3\x1f\xcb\x87\x28\x0f\xac\x03\x79\x53\x1c\xfb\x43\xcf\x65\xd4\x21\x4d\xbd\x0f\x53\x37\x1c\x52\x66\x15\x04\xd6\x59\x0e\xd6\x3e\x49\x4b\x48\x60\x02\xc1\x18\x12\xf8\x65\x02\x6c\x38\x86\xc1\x20\x21\xf0\x27\x14\xf3\x64\x01\x13\x70\x3e\x9d\x8f\x01\x9e\xa0\x98\x07\xd8\x3c\x04\x8e\xf6\x8f\x72\xf9\x97\x78\xb9\x8c\xbe\x84\x6b\xea\x3b\x8e\x55\xd8\x10\x34\x9c\x6a\xb2\xc8\x86\x6c\x5f\x16\xdf\x9b\xfe\x73\x6e\x76\x0c\x39\xbc\x03\xe6\x8e\xb9\xeb\xa6\x4a\x98\x67\xfb\x74\xc5\x15\xf3\xc1\xc0\x06\x36\x24\x63\x78\xfa\x01\x82\xb9\xca\xbc\x69\x86\xe1\x32\x4b\x8b\x32\x4a\x4b\xc0\xf4\x00\x36\xe1\xfa\x64\xce\x3c\x1e\xc8\x9f\x26\xa7\x4d\x27\x16\xbb\xdd\x1d\x38\x01\x93\xdd\x81\xde\x3d\x6d\x75\x07\xf2\xaf\xdf\xc8\xdf\x5b\xdd\x81\xee\x52\x37\x12\xb4\xbb\xa5\x68\xd4\x63\xdb\xa9\x47\xd2\xea\x0e\x7a\x5c\xf6\x18\x91\x3e\xda\x46\x02\x7d\xdc\xfa\xb4\xf7\x76\x07\x4e\x30\xea\x1f\xb7\xdb\xdf\xdd\x97\x58\x2e\xe8\xcd\xa0\xd3\x97\x6f\xa7\xbf\xfb\x87\x19\xec\xd8\x7e\x76\x1a\x1c\xcc\xe0\xd3\xd8\x34\x8b\x32\x2a\x93\x25\x7c\xcd\x92\x55\x1f\xf7\x90\x61\x47\x10\xd9\x92\x65\x82\xd7\x5a\x4b\xd2\xfe\x4f\xd3\x90\x9d\x48\xca\x9d\x03\x13\x88\xe6\xce\x02\xee\x20\x9a\x0f\xc5\x7f\x54\x36\xa9\x6c\x33\x6c\x57\x35\x21\x9a\x33\xd1\x7f\x2c\x61\xb2\x49\x65\x9b\xb1\x85\x0d\x94\x8c\x5b\x9e\x28\x7a\xa2\x02\xe3\x4b\x15\xd9\xa4\xb2\xcd\x68\xcb\x93\x2b\xfa\x03\x09\x93\x4d\x2a\xdb\xcc\xed\xf5\xc4\xd0\xd3\x8b\x83\xd4\x3d\x79\xa2\x7f\x24\x61\xb2\x49\x65\x9b\x79\xbd\x9e\x5c\xf4\xf4\xe2\x20\x75\x4f\x2f\x26\xbb\xcf\x93\x87\x9e\x5e\x1c\xa4\xee\xe9\xc5\x64\x0b\x4f\xba\xab\x72\xf7\x28\x67\xea\x6e\xe7\x70\xa1\x18\xf1\x04\x76\xde\x98\x37\x86\xcd\x06\xd5\x44\x54\x93\xb1\x5a\xc6\x9b\xbe\x68\x39\x02\x49\xb5\x96\x26\x63\xb5\x0c\x9b\xa2\x45\x51\x76\xdc\x6c\x50\x4d\x44\x35\x19\xab\x65\xbc\xe9\x8a\x16\x43\x59\xd0\x6c\x50\x4d\x44\x35\x19\xab\x65\xbc\xe9\x89\x96\x8b\xb2\x51\xb3\x41\x35\x11\xd5\x64\xac\x96\xa1\x90\x17\xff\x6a\xbb\xe5\x69\xb1\x81\x32\x32\x96\x29\x9a\x34\xe7\x6f\xb4\xb0\x81\x39\x44\xf9\xd3\xf4\x70\x95\xb1\x11\xa9\xa2\xd5\xa4\x94\xf3\xf5\x98\x54\x83\xd3\xcc\x0a\x8a\x05\xa4\x9a\x25\xdd\xf0\xc2\x06\x57\x3a\xed\x98\xc5\xa5\x4d\x49\x95\x7e\xcd\x2e\xe5\xcb\x91\x6f\x90\x2a\xaf\xba\x94\x47\x03\xd2\x2b\xed\x84\x83\xba\x1e\xa9\xf2\xae\xfb\x1d\x2e\x6c\x90\xb1\x72\x92\x4d\xf4\x95\x6b\x03\x93\xb1\x76\x42\xc5\x75\xab\xcc\x76\xa4\x7c\x05\x80\x4c\x3e\x6b\xcf\x4c\xc0\xed\xba\x44\x51\x46\x1f\x10\x4e\x9b\x14\x52\xbf\x25\xe5\xa3\x95\x3e\xdb\x83\xc5\x30\x65\x12\x3a\x19\x42\x8f\x2a\xbb\x41\x5b\x7a\xcc\x5d\xaa\xec\x1e\xb7\xa5\xb8\x98\xe5\xb4\x51\xda\x92\x72\x55\xf0\x89\x5a\x42\xba\x26\x67\x03\xa8\x50\x1c\x4d\x5a\xee\x1e\x6d\x90\x35\x42\x16\xa2\xa1\xdc\x30\xc6\x55\x8f\x2f\xcb\x45\x55\x2b\x26\x70\x9f\x94\x45\xbc\x8d\x97\x25\x96\xbb\x3b\x24\x15\xff\x65\x73\x1c\x69\x16\x95\xeb\x93\x79\x8e\x9a\xc9\x1a\x2c\xdc\xa3\xe0\x3d\x50\xbe\x4f\x99\x86\x5a\x2d\x4d\x6b\x94\x5b\x73\x85\x21\x9b\xdb\xe5\xe3\x56\x4c\x6d\x02\x19\x07\x7a\x02\x63\x73\x15\x09\x74\xdb\x40\x77\x71\xb7\x1b\xda\xd2\xa8\xa7\x60\x5e\x1b\xe6\x2d\xee\x76\x3e\x22\xec\xdd\x10\x43\xd0\x86\xec\xc9\x21\x1b\x32\x41\x43\x9e\x20\xc3\x90\xc9\xf1\x45\x4b\xd2\xa1\x69\x76\xc8\x87\x79\xcc\x5d\x0f\xf9\xc3\x17\xfe\x55\x2d\x68\x22\x7d\x8e\x0c\x04\xc8\xe6\x3a\x0a\x79\xdc\x46\x1e\x73\xe4\x48\x80\x6c\xae\xa3\x90\x41\x1b\x19\xc8\xd8\xd1\xec\xa8\xc2\x8d\xda\xb8\x91\x0c\x7e\x54\x07\xaf\x45\x1f\xa8\xe8\x55\xf8\xd4\x11\x11\xab\x04\x50\x2a\xdb\x8a\x62\xda\x9c\x22\x45\x28\xce\x13\x15\x24\xa1\x72\x2c\x15\x99\x35\x3c\x92\x80\xba\x12\x69\xa3\x6e\x8d\xef\x50\x81\x22\x17\xa8\x27\x91\x36\xea\xd6\xf8\x0e\x23\xa8\xa2\x84\x74\xe1\x35\xc0\x1d\x5e\x50\x45\x0c\xda\x60\x46\x9d\x87\x61\x2b\x0f\x7e\x9d\x87\x0e\x17\x28\x92\x81\xe2\xa4\x51\xa4\x03\xf5\x1b\xae\x3b\x8c\xa0\x48\x09\x1a\x48\xa4\x8d\xba\x35\xbe\xc3\x0b\x8a\xc4\xa0\x23\x89\xb4\x51\xb7\xc6\x77\xd8\x41\x15\x3d\xa4\x8b\x51\x03\xdc\xa1\x08\x55\x1c\xa1\xa3\x9e\x3c\xb0\x16\x1f\x58\xcd\x07\xd6\xe1\x03\x13\x25\x03\x67\x8a\x21\x1f\x58\x83\x0f\xac\xc3\x07\x86\x7c\x60\xae\x44\xda\xa8\x5b\xe3\xbb\xa5\x01\xf9\xc0\x3c\x89\xb4\x51\xb7\xc6\x77\xf8\xc0\x14\x1f\xa4\x8b\x06\x1f\x58\x87\x0f\x4c\xf1\x81\x29\x3e\x18\x86\xf1\x64\xe2\xdf\x93\xf9\x64\x9a\xe5\xb7\xc7\x78\x15\xaf\xa1\x28\xf3\xfd\xb2\xe4\x07\x72\xac\xa3\xe2\x59\xcc\x5d\x06\x6f\xa1\x48\xbe\xc7\xd9\x5a\xbe\xa9\x2f\xc6\xe6\x13\x3c\x44\xc5\x83\xcb\xc2\x72\x5c\x1b\xd8\xa7\x49\x96\x56\xfa\x1e\x88\xf3\x9b\x57\xcc\x7d\xaf\x65\xc2\x23\x58\x62\xf9\x7b\x00\x3e\x5a\x90\x24\x2d\x09\xd6\xf3\x24\x2d\x99\x78\x76\x11\xac\x82\x78\xe2\xd9\x85\x78\x15\x24\x10\xcf\x2e\x24\xa8\x20\xd4\x37\x56\x51\x19\xa9\xd8\x7c\xaf\x1b\x9b\xa9\x62\x03\x19\x5b\x31\xa7\x2c\xe8\x49\x4f\x27\x03\x5d\x98\x57\xe3\x02\x89\x0b\xfa\x70\x41\x9d\x2a\xe0\x6e\x31\x5b\x3a\x4c\x4b\x97\xc4\xb0\x1e\x50\x23\x63\x12\xd5\x19\x99\x9e\x34\x89\xea\x8c\xab\x95\x37\x81\xa2\x7e\x0f\x8c\xfa\x35\x5b\x28\x0b\x9e\xa1\x8b\xa2\x1b\x73\x9c\x9e\x84\x6a\x5c\xd1\x21\xbd\x64\xe9\x40\x58\x97\x0a\x1d\x8c\xc6\x05\x15\x52\x07\xd5\x0c\x89\x39\x0e\x86\xf4\x2a\x59\xb7\xee\x11\x7f\x99\xf4\x5e\x3c\x9a\x61\x18\x95\x65\x9e\xdc\xef\xcb\x38\x0c\x2d\x2b\x8f\xff\xb3\x0a\xff\xc8\xf2\x2f\xe1\x26\xcf\xf6\x8f\x21\xf7\x63\xfd\x7e\x7d\xf9\xe9\x0a\xaf\xf8\x6c\xa0\xfc\x8d\x88\x10\x75\x53\x18\x86\x5f\xe2\x3c\x8d\xb7\xe2\xc5\xbb\x88\xa3\xad\x65\x1a\x61\xb8\xd9\x66\xf7\x11\xef\xdc\x46\x65\xb2\x8d\x71\x3e\x8e\x20\x8f\x8b\x32\x4f\x96\x25\x6c\xc2\x6c\x5f\x3e\xee\x4b\x9b\x63\xab\x5b\x20\xb5\x82\xc5\x4b\xf9\x11\x6c\xc2\x87\x38\x5a\xc5\xb9\xdd\xb0\x58\xcd\x5b\x0d\x5a\x45\x1b\xfa\x32\x84\xd9\x6a\xa9\x14\x65\x94\x97\x61\x9a\xa5\xcb\xb8\xea\x2b\xa3\x7c\x13\xf3\xe1\x20\xad\xa3\xfb\x2c\x2f\x4d\x43\x6c\xd9\x7c\x17\xc7\x0e\xbc\x1a\x30\xf2\xb8\xdc\xe7\xe9\x98\x97\x2a\x89\x16\xaf\x86\x9b\x64\x05\x13\xd8\xc4\xa5\x1c\x46\x98\xac\x2c\x3c\x6c\x36\x30\xe2\x9e\x38\x14\xc8\x64\x05\xbf\xc2\xed\x87\xeb\xd9\xf4\xf4\x46\x87\xf1\x00\x04\xc8\x12\xa8\x3a\xff\x04\xde\xbf\x07\x7c\x63\x0b\xc3\x6d\xb6\x94\xd1\xf2\xd2\x00\xc5\x43\x94\xc7\xf7\xfb\xf5\x1c\x6f\x6c\x39\x98\x93\xa7\x03\x3b\x92\x4e\x10\x0e\x93\x4a\x0d\x06\xca\x6f\x7d\x94\x2d\xca\xa8\x8c\xe7\x6c\x88\x7c\xb7\x64\x99\x38\x22\xd8\x4d\xc4\x39\xd6\xb2\xda\xd7\x78\xde\x11\x51\x13\x47\xf0\x14\x6c\x1a\xc2\x0e\x6e\x04\x8d\xe4\xc3\x80\x27\x81\xcb\xf5\x9b\x4e\xbf\x73\xd3\x69\x1a\xd2\x84\xb8\xef\x94\xc9\x17\x5d\x43\x71\xe7\xa9\x5f\xac\xe1\x9d\xa7\x42\xf4\xdf\x8a\x3a\x1c\x61\x1a\xfa\xbd\x28\xc7\xdb\xf8\x06\x66\x1a\xaf\x1e\xf3\x68\xb3\x8b\x60\x9f\xe6\xd9\x76\x0b\xb4\x39\xcc\x12\x67\xc7\x19\xe3\x8f\x77\xd5\x2c\xf2\xe6\x60\x20\x06\xcc\x69\x83\xb0\x49\x3d\xed\x42\x62\x60\xca\xff\xf6\x9e\x17\x77\xcc\xa0\x58\xc4\x8d\xbc\x8e\xc5\x66\x68\xdc\x47\x79\x9e\xc4\xb9\x75\x72\xfe\x8f\xf0\xfc\xf2\x64\x7a\x1e\x7e\x9c\x7d\x0c\x7f\x9b\x5d\x9c\xcc\xc4\xe9\x41\x96\xbb\x5d\x72\x50\x33\x29\xcc\x8e\x7f\x56\x19\x92\x34\x29\x9d\x5a\x59\xd4\x33\x31\x82\x17\x0d\x74\x33\xd4\x48\x51\x24\x12\x14\xc1\x3b\x98\x9e\x9c\xcc\x6e\x6e\x66\x37\xbc\x35\x98\x00\xf5\x65\x1e\xee\xb3\x6c\x0b\xfb\xc7\x55\x54\xc6\xa1\xa2\x63\x63\x89\xf0\xd4\x44\x9c\xef\x1e\xa9\x57\x8a\x3c\x35\xe9\xae\x79\x8f\xce\x21\x47\x72\x88\xfa\x0d\x0e\x89\x49\x69\x3a\x54\xfd\x46\x2b\x7c\x98\xe0\xb7\x05\x91\x9c\x3b\xb0\x22\x18\x40\x42\x6c\x39\x57\x70\x44\x5e\xef\x92\x03\x99\x27\x0b\x3e\x32\xf5\xe5\x45\x9c\x78\x9e\xf0\xf9\x62\xee\x8c\x97\xea\x15\x2f\x11\x58\xda\x24\x9c\x0f\xbd\x3d\xca\xd7\xe2\x3d\x10\x23\xd0\x94\x98\x1a\x0b\xfe\x57\xa7\xe6\x20\x52\x73\x80\x77\x98\x99\x83\x22\x2b\x3f\x76\x26\x87\xf9\x41\x05\x2e\x1a\xb6\x30\x36\x6f\xbb\xe5\x25\x68\xf1\x46\xed\x47\xd5\x8c\x2d\xe6\x07\x75\xec\xfb\x99\x24\x54\x87\x3e\x53\x4b\xbf\xd7\xb4\xc8\x97\x87\x38\x02\x58\x1c\xdf\xf8\x00\xb4\x4b\x0e\x6f\x0a\x87\x32\x97\xd8\x7d\x12\x6f\xe8\x1f\xf7\x4b\x82\x51\x74\xdf\x2f\x59\xae\xe2\x35\x31\x0d\x83\xfc\x14\xf9\x9f\x59\xdf\x9d\x4a\x89\xe7\x6b\x15\xa3\x3c\x80\xfd\xec\x2a\xc7\x04\xe9\xe4\xa6\xee\x5f\xaa\x90\xe2\x75\xef\xd9\x12\x29\xde\x9c\xfe\x4a\x8d\x14\x37\x1e\x3c\x09\xf2\x53\x97\x30\xe4\x2c\x08\xbc\x93\xbb\x69\x5d\x07\x1b\xdb\x67\xbd\x7f\x1a\x4f\x86\xa1\xaa\x50\xb1\xcd\x4a\x98\xc0\x2e\x49\xc5\x0a\x23\x1f\xa7\xff\x0a\x2f\x3f\xdd\x5e\x7d\xba\xbd\xb1\x21\x2a\xb3\x5d\xb2\x0c\x93\x74\x69\xbd\x56\x27\x07\xf4\x34\x10\xd7\xc0\x46\xd5\xc9\xed\x2c\xc4\xe6\x3a\x96\x2f\x13\xfa\x01\x65\x13\xa7\x71\xce\x97\xff\x2a\xda\x84\x49\x19\xef\x44\x56\x71\x57\xb2\x41\x5b\x92\xb8\x5b\x56\x2b\x72\x9b\x6c\x1e\xca\xf6\x29\x43\x6e\xa8\xf2\x08\xd2\xa3\xae\xce\x1e\x18\x3b\x7a\x4a\xb3\x55\x2c\xf6\x76\xf4\xc9\xf7\xc0\xce\x91\xe1\x19\x2f\x2a\xe9\xca\xcc\xfb\xaa\xfa\x1c\x31\x02\x2a\xb7\xa6\x51\x9d\x05\x81\x07\xca\xd1\xbc\x57\xfd\xae\xd6\x2e\x56\x3a\x19\xdc\x5c\xd9\xfc\xb5\xf1\x5d\x78\xf1\x46\x6c\x27\x2d\xdd\x42\xde\x4d\x49\x15\x2e\xaf\xbe\x5c\xd6\x40\x3c\x3f\x93\xee\x06\xaf\x8a\x73\xff\xc7\xe5\x06\xa5\x51\xe3\x31\xca\xe3\xb4\x0c\x93\x74\x15\x1f\x64\x79\x52\x23\xbd\x83\xc4\x86\xd6\xc0\x12\xf8\xb5\xfe\x22\x8d\xf5\xb9\x8e\x86\x73\xa5\x3f\x07\xdc\x6a\x8f\xc4\xae\x92\xd3\x1c\x85\x48\x8a\x5a\x9b\xcf\x04\x6e\x98\x46\xb7\x96\x37\xa7\xef\x35\x08\x6a\x74\xab\x37\xb7\x5c\xa1\x26\xf0\x3f\xbc\x60\x17\xe5\xb8\xd2\x7e\xcb\xb5\xff\xdb\xb4\x46\x89\x18\x1b\x4c\xa0\x27\x9a\xb1\xf9\x64\xfe\x3f\x00\x00\xff\xff\x10\x41\xaf\xc5\xba\x20\x00\x00")

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
		size: 8378,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1629589681, 0),
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
