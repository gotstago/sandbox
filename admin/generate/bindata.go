// Code generated by go-bindata.
// sources:
// painkiller/painkiller.go
// painkiller/pill_string.go
// DO NOT EDIT!

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

var _painkillerPainkillerGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x8c\x31\xb2\xc2\x30\x0c\x44\x7b\x9d\x42\xe5\xff\x05\x93\x9e\x99\x14\xa1\xa3\xcb\x15\x94\x20\x8c\x26\x46\xf2\x28\xa2\xc8\xed\x71\xec\x06\x57\xef\x79\x57\x5b\x68\xdd\x28\x31\x16\x12\xdd\x24\x67\x76\x80\x61\x48\x76\x4d\xac\xec\x14\x8c\x7b\xb8\x68\x62\xc7\x4b\x1c\x85\xc7\xb9\x96\x00\x4e\xc4\x13\x51\x34\x00\x56\xd3\x3d\xf0\x0f\xb0\xbe\x39\xd3\xca\x8b\xf5\x74\x44\xb1\xa0\xf6\x3f\xed\x45\xea\x52\xe3\xfb\xf2\x29\x6e\x4f\xee\x36\x93\xd7\x93\xa0\xb7\xe5\xe6\x37\x56\x7a\xf8\xd1\x65\x6a\x89\xa8\x95\x17\x6b\xdd\xfb\x2d\xff\x7f\x03\x00\x00\xff\xff\x87\x6f\xed\xe3\xbe\x00\x00\x00")

func painkillerPainkillerGoBytes() ([]byte, error) {
	return bindataRead(
		_painkillerPainkillerGo,
		"painkiller/painkiller.go",
	)
}

func painkillerPainkillerGo() (*asset, error) {
	bytes, err := painkillerPainkillerGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "painkiller/painkiller.go", size: 190, mode: os.FileMode(420), modTime: time.Unix(1462037454, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _painkillerPill_stringGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x8f\x4d\x6b\xf2\x40\x14\x85\xd7\xde\x5f\x71\x18\x78\x21\xe1\x8d\xf1\x13\x5a\xda\x5a\x68\x6b\x17\x6e\xaa\xa0\x3b\x09\x32\x26\x37\x32\x98\xcc\x0c\xe3\xa4\x54\xd4\xff\xde\xd1\xb6\xd4\xae\xe6\xe3\x9e\xe7\x5c\x9e\x4e\x07\x2f\xa6\x60\x6c\x58\xb3\x93\x9e\x0b\xac\xf7\x10\x3b\xef\x94\xde\xb0\x43\xdb\xef\x2d\x8f\x66\xaa\xaa\xc4\x3d\xc6\x53\xbc\x4d\x17\x78\x1d\x4f\x16\x44\x56\xe6\x5b\xb9\x61\x58\xa9\xf4\x36\xcc\xd9\x11\xa9\xda\x1a\xe7\x21\xca\xda\x0b\xa2\xdc\xe8\x9d\xc7\xea\x0c\xaf\xb4\xac\x19\x23\x88\x59\x25\x73\x5e\x9b\xa7\x9d\x55\x61\xc3\x64\xdd\x58\x67\x4a\xd6\x33\xe9\xc2\xbf\x97\xb5\xa9\x9e\x59\xcb\xc2\xed\xab\x50\xf0\x2e\xdd\x37\xae\x74\xc1\x1f\x81\x5f\xa6\x69\x9a\x35\x4a\xfb\xdb\x43\x37\xc1\x4d\x82\xde\x30\x41\x7f\x90\x60\x10\xce\x61\xff\x44\x54\x36\x3a\x47\xa4\x70\xc6\x62\xcc\x2f\x1e\x51\x8c\x2f\x21\x1c\xa8\xa5\x4a\x28\x3c\xa0\x8b\xe3\x31\x5c\x1e\x47\x97\x64\x54\xb1\x8e\xae\x56\xc5\xed\x5e\x7c\x0e\xb7\x1c\xfb\xc6\x69\x04\xa1\x74\x6e\x43\x85\x2f\x23\x71\x01\xfe\x15\xb1\x48\xa0\x62\x6a\x9d\xe8\x27\xf5\xab\xba\xbc\xea\x5a\xaa\xec\xee\xcf\xf3\x7f\x2f\xcb\xe8\x44\x9f\x01\x00\x00\xff\xff\xfa\x51\xb7\xf6\x7a\x01\x00\x00")

func painkillerPill_stringGoBytes() ([]byte, error) {
	return bindataRead(
		_painkillerPill_stringGo,
		"painkiller/pill_string.go",
	)
}

func painkillerPill_stringGo() (*asset, error) {
	bytes, err := painkillerPill_stringGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "painkiller/pill_string.go", size: 378, mode: os.FileMode(420), modTime: time.Unix(1462037458, 0)}
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
	"painkiller/painkiller.go": painkillerPainkillerGo,
	"painkiller/pill_string.go": painkillerPill_stringGo,
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
	"painkiller": &bintree{nil, map[string]*bintree{
		"painkiller.go": &bintree{painkillerPainkillerGo, map[string]*bintree{}},
		"pill_string.go": &bintree{painkillerPill_stringGo, map[string]*bintree{}},
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
