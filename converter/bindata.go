// Code generated by go-bindata.
// sources:
// data/cx_to_cyjs_style.csv
// DO NOT EDIT!

package converter

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

var _dataCx_to_cyjs_styleCsv = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x92\xd1\x8e\xb2\x30\x10\x85\xef\x7d\x96\xf2\x10\x15\x46\x21\x3f\x69\x4d\xa9\x31\xfe\x37\xa4\x62\x85\x66\xb1\x35\xa5\xc6\x75\x9f\x7e\x0d\x68\xa4\xd2\xd5\x3b\xc2\x7c\x73\xce\xe4\x9c\x12\x9a\x40\x39\xa7\x2c\x01\x56\xae\x70\x46\x38\xda\x89\xea\xab\xb6\xe6\xac\xf7\x51\x65\x5a\x63\x51\xe7\xac\xd2\xf5\x6c\x4c\x72\x86\x49\xb1\xc2\x0c\x48\xbc\x45\x3b\x63\xf7\xd2\x46\xe6\x24\x2a\xe5\xae\x48\x9f\x8f\x3b\x69\x3d\x7c\x93\x25\x3c\x7d\x70\x17\xb5\x77\x8d\x47\x2d\xb2\x3c\x2f\x63\x9a\x53\xf6\xde\xdc\x77\x7d\x92\x21\xe7\x1c\xcf\x21\x47\x95\xd1\x4e\x6a\xe7\xa9\xf4\x93\xbb\xdd\xd4\x63\x98\x2e\x28\xe1\x65\x91\xfd\x07\x74\xb8\x29\x44\x9d\xfa\x91\x53\xf5\x81\x5a\xe0\xf8\x4e\x1d\xc4\x51\xb5\xd7\xfe\x7b\x4c\x79\x57\x3b\xf9\xed\x82\xf7\x16\x29\x5e\x01\xea\x1a\x71\x92\xde\x3d\x29\x64\xcb\x94\xa3\x46\xaa\xba\x71\xde\xc6\x10\xea\x34\xcd\x02\x72\x88\x39\x24\x1f\xea\x9c\x41\xb2\x0c\x8a\xf4\xff\x83\xf1\x3d\x27\xef\x02\x7a\xa5\xc2\x01\x8d\xa8\x40\x15\xa3\xe9\xc7\xf8\x06\x36\x23\xb7\xf7\xb1\xbd\x45\xd8\x2a\x2d\xa3\xce\x5d\x5b\xe9\x89\x15\x9c\xd1\x7f\x50\xae\xc9\x4b\x3c\x3d\x3e\xf5\xf6\x5c\x43\x86\x05\x5d\xb3\x18\x4a\xcc\x18\xdd\x3c\xca\x33\x67\x5b\xc9\x48\x58\x6b\x2e\x91\xd7\xe4\x20\x89\xd9\x12\xb8\xb7\xe1\x84\xad\xa5\xfb\x73\xe3\xd3\xa9\xbf\x01\x00\x00\xff\xff\xca\x45\xbb\xea\xbc\x03\x00\x00")

func dataCx_to_cyjs_styleCsvBytes() ([]byte, error) {
	return bindataRead(
		_dataCx_to_cyjs_styleCsv,
		"data/cx_to_cyjs_style.csv",
	)
}

func dataCx_to_cyjs_styleCsv() (*asset, error) {
	bytes, err := dataCx_to_cyjs_styleCsvBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/cx_to_cyjs_style.csv", size: 956, mode: os.FileMode(420), modTime: time.Unix(1447890745, 0)}
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
	"data/cx_to_cyjs_style.csv": dataCx_to_cyjs_styleCsv,
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
	"data": &bintree{nil, map[string]*bintree{
		"cx_to_cyjs_style.csv": &bintree{dataCx_to_cyjs_styleCsv, map[string]*bintree{}},
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
