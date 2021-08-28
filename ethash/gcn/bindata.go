// Code generated by go-bindata. DO NOT EDIT.
// -- Common file --

package gcn


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
	"gcn/bin/ethash_baffin_lws128.bin":         bindataGcnBinEthashbaffinlws128Bin,
	"gcn/bin/ethash_baffin_lws128_exit.bin":    bindataGcnBinEthashbaffinlws128exitBin,
	"gcn/bin/ethash_baffin_lws256.bin":         bindataGcnBinEthashbaffinlws256Bin,
	"gcn/bin/ethash_baffin_lws256_exit.bin":    bindataGcnBinEthashbaffinlws256exitBin,
	"gcn/bin/ethash_baffin_lws64.bin":          bindataGcnBinEthashbaffinlws64Bin,
	"gcn/bin/ethash_baffin_lws64_exit.bin":     bindataGcnBinEthashbaffinlws64exitBin,
	"gcn/bin/ethash_ellesmere_lws128.bin":      bindataGcnBinEthashellesmerelws128Bin,
	"gcn/bin/ethash_ellesmere_lws128_exit.bin": bindataGcnBinEthashellesmerelws128exitBin,
	"gcn/bin/ethash_ellesmere_lws256.bin":      bindataGcnBinEthashellesmerelws256Bin,
	"gcn/bin/ethash_ellesmere_lws256_exit.bin": bindataGcnBinEthashellesmerelws256exitBin,
	"gcn/bin/ethash_ellesmere_lws64.bin":       bindataGcnBinEthashellesmerelws64Bin,
	"gcn/bin/ethash_ellesmere_lws64_exit.bin":  bindataGcnBinEthashellesmerelws64exitBin,
	"gcn/bin/ethash_gfx900_lws128.bin":         bindataGcnBinEthashgfx900lws128Bin,
	"gcn/bin/ethash_gfx900_lws128_exit.bin":    bindataGcnBinEthashgfx900lws128exitBin,
	"gcn/bin/ethash_gfx900_lws256.bin":         bindataGcnBinEthashgfx900lws256Bin,
	"gcn/bin/ethash_gfx900_lws256_exit.bin":    bindataGcnBinEthashgfx900lws256exitBin,
	"gcn/bin/ethash_gfx900_lws64.bin":          bindataGcnBinEthashgfx900lws64Bin,
	"gcn/bin/ethash_gfx900_lws64_exit.bin":     bindataGcnBinEthashgfx900lws64exitBin,
	"gcn/bin/ethash_gfx901_lws128.bin":         bindataGcnBinEthashgfx901lws128Bin,
	"gcn/bin/ethash_gfx901_lws128_exit.bin":    bindataGcnBinEthashgfx901lws128exitBin,
	"gcn/bin/ethash_gfx901_lws256.bin":         bindataGcnBinEthashgfx901lws256Bin,
	"gcn/bin/ethash_gfx901_lws256_exit.bin":    bindataGcnBinEthashgfx901lws256exitBin,
	"gcn/bin/ethash_gfx901_lws64.bin":          bindataGcnBinEthashgfx901lws64Bin,
	"gcn/bin/ethash_gfx901_lws64_exit.bin":     bindataGcnBinEthashgfx901lws64exitBin,
	"gcn/bin/ethash_gfx906_lws128.bin":         bindataGcnBinEthashgfx906lws128Bin,
	"gcn/bin/ethash_gfx906_lws128_exit.bin":    bindataGcnBinEthashgfx906lws128exitBin,
	"gcn/bin/ethash_gfx906_lws256.bin":         bindataGcnBinEthashgfx906lws256Bin,
	"gcn/bin/ethash_gfx906_lws256_exit.bin":    bindataGcnBinEthashgfx906lws256exitBin,
	"gcn/bin/ethash_gfx906_lws64.bin":          bindataGcnBinEthashgfx906lws64Bin,
	"gcn/bin/ethash_gfx906_lws64_exit.bin":     bindataGcnBinEthashgfx906lws64exitBin,
	"gcn/bin/ethash_tonga_lws128.bin":          bindataGcnBinEthashtongalws128Bin,
	"gcn/bin/ethash_tonga_lws128_exit.bin":     bindataGcnBinEthashtongalws128exitBin,
	"gcn/bin/ethash_tonga_lws256.bin":          bindataGcnBinEthashtongalws256Bin,
	"gcn/bin/ethash_tonga_lws256_exit.bin":     bindataGcnBinEthashtongalws256exitBin,
	"gcn/bin/ethash_tonga_lws64.bin":           bindataGcnBinEthashtongalws64Bin,
	"gcn/bin/ethash_tonga_lws64_exit.bin":      bindataGcnBinEthashtongalws64exitBin,
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
	"gcn": {Func: nil, Children: map[string]*bintree{
		"bin": {Func: nil, Children: map[string]*bintree{
			"ethash_baffin_lws128.bin": {Func: bindataGcnBinEthashbaffinlws128Bin, Children: map[string]*bintree{}},
			"ethash_baffin_lws128_exit.bin": {Func: bindataGcnBinEthashbaffinlws128exitBin, Children: map[string]*bintree{}},
			"ethash_baffin_lws256.bin": {Func: bindataGcnBinEthashbaffinlws256Bin, Children: map[string]*bintree{}},
			"ethash_baffin_lws256_exit.bin": {Func: bindataGcnBinEthashbaffinlws256exitBin, Children: map[string]*bintree{}},
			"ethash_baffin_lws64.bin": {Func: bindataGcnBinEthashbaffinlws64Bin, Children: map[string]*bintree{}},
			"ethash_baffin_lws64_exit.bin": {Func: bindataGcnBinEthashbaffinlws64exitBin, Children: map[string]*bintree{}},
			"ethash_ellesmere_lws128.bin": {Func: bindataGcnBinEthashellesmerelws128Bin, Children: map[string]*bintree{}},
			"ethash_ellesmere_lws128_exit.bin": {Func: bindataGcnBinEthashellesmerelws128exitBin, Children: map[string]*bintree{}},
			"ethash_ellesmere_lws256.bin": {Func: bindataGcnBinEthashellesmerelws256Bin, Children: map[string]*bintree{}},
			"ethash_ellesmere_lws256_exit.bin": {Func: bindataGcnBinEthashellesmerelws256exitBin, Children: map[string]*bintree{}},
			"ethash_ellesmere_lws64.bin": {Func: bindataGcnBinEthashellesmerelws64Bin, Children: map[string]*bintree{}},
			"ethash_ellesmere_lws64_exit.bin": {Func: bindataGcnBinEthashellesmerelws64exitBin, Children: map[string]*bintree{}},
			"ethash_gfx900_lws128.bin": {Func: bindataGcnBinEthashgfx900lws128Bin, Children: map[string]*bintree{}},
			"ethash_gfx900_lws128_exit.bin": {Func: bindataGcnBinEthashgfx900lws128exitBin, Children: map[string]*bintree{}},
			"ethash_gfx900_lws256.bin": {Func: bindataGcnBinEthashgfx900lws256Bin, Children: map[string]*bintree{}},
			"ethash_gfx900_lws256_exit.bin": {Func: bindataGcnBinEthashgfx900lws256exitBin, Children: map[string]*bintree{}},
			"ethash_gfx900_lws64.bin": {Func: bindataGcnBinEthashgfx900lws64Bin, Children: map[string]*bintree{}},
			"ethash_gfx900_lws64_exit.bin": {Func: bindataGcnBinEthashgfx900lws64exitBin, Children: map[string]*bintree{}},
			"ethash_gfx901_lws128.bin": {Func: bindataGcnBinEthashgfx901lws128Bin, Children: map[string]*bintree{}},
			"ethash_gfx901_lws128_exit.bin": {Func: bindataGcnBinEthashgfx901lws128exitBin, Children: map[string]*bintree{}},
			"ethash_gfx901_lws256.bin": {Func: bindataGcnBinEthashgfx901lws256Bin, Children: map[string]*bintree{}},
			"ethash_gfx901_lws256_exit.bin": {Func: bindataGcnBinEthashgfx901lws256exitBin, Children: map[string]*bintree{}},
			"ethash_gfx901_lws64.bin": {Func: bindataGcnBinEthashgfx901lws64Bin, Children: map[string]*bintree{}},
			"ethash_gfx901_lws64_exit.bin": {Func: bindataGcnBinEthashgfx901lws64exitBin, Children: map[string]*bintree{}},
			"ethash_gfx906_lws128.bin": {Func: bindataGcnBinEthashgfx906lws128Bin, Children: map[string]*bintree{}},
			"ethash_gfx906_lws128_exit.bin": {Func: bindataGcnBinEthashgfx906lws128exitBin, Children: map[string]*bintree{}},
			"ethash_gfx906_lws256.bin": {Func: bindataGcnBinEthashgfx906lws256Bin, Children: map[string]*bintree{}},
			"ethash_gfx906_lws256_exit.bin": {Func: bindataGcnBinEthashgfx906lws256exitBin, Children: map[string]*bintree{}},
			"ethash_gfx906_lws64.bin": {Func: bindataGcnBinEthashgfx906lws64Bin, Children: map[string]*bintree{}},
			"ethash_gfx906_lws64_exit.bin": {Func: bindataGcnBinEthashgfx906lws64exitBin, Children: map[string]*bintree{}},
			"ethash_tonga_lws128.bin": {Func: bindataGcnBinEthashtongalws128Bin, Children: map[string]*bintree{}},
			"ethash_tonga_lws128_exit.bin": {Func: bindataGcnBinEthashtongalws128exitBin, Children: map[string]*bintree{}},
			"ethash_tonga_lws256.bin": {Func: bindataGcnBinEthashtongalws256Bin, Children: map[string]*bintree{}},
			"ethash_tonga_lws256_exit.bin": {Func: bindataGcnBinEthashtongalws256exitBin, Children: map[string]*bintree{}},
			"ethash_tonga_lws64.bin": {Func: bindataGcnBinEthashtongalws64Bin, Children: map[string]*bintree{}},
			"ethash_tonga_lws64_exit.bin": {Func: bindataGcnBinEthashtongalws64exitBin, Children: map[string]*bintree{}},
		}},
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