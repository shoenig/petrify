// Code generated by petrify. DO NOT EDIT.
// sources:
// in/a/test.asset
// in/b/test.asset
// in/c/test.asset
// in/test.asset

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"
	"unsafe"
	"path"
	"html/template"
)

func bindataRead(data, name string) ([]byte, error) {
	var empty [0]byte
	sx := (*reflect.StringHeader)(unsafe.Pointer(&data))
	b := empty[:]
	bx := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bx.Data = sx.Data
	bx.Len = len(data)
	bx.Cap = bx.Len
	return b, nil
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

var _inATestAsset = "\x2f\x2f\x20\x73\x61\x6d\x70\x6c\x65\x20\x66\x69\x6c\x65\x0a"

func inATestAssetBytes() ([]byte, error) {
	return bindataRead(
		_inATestAsset,
		"in/a/test.asset",
	)
}

func inATestAsset() (*asset, error) {
	bytes, err := inATestAssetBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "in/a/test.asset", size: 15, mode: os.FileMode(420), modTime: time.Unix(1558857921, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _inBTestAsset = "\x2f\x2f\x20\x73\x61\x6d\x70\x6c\x65\x20\x66\x69\x6c\x65\x0a"

func inBTestAssetBytes() ([]byte, error) {
	return bindataRead(
		_inBTestAsset,
		"in/b/test.asset",
	)
}

func inBTestAsset() (*asset, error) {
	bytes, err := inBTestAssetBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "in/b/test.asset", size: 15, mode: os.FileMode(420), modTime: time.Unix(1558857921, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _inCTestAsset = "\x2f\x2f\x20\x73\x61\x6d\x70\x6c\x65\x20\x66\x69\x6c\x65\x0a"

func inCTestAssetBytes() ([]byte, error) {
	return bindataRead(
		_inCTestAsset,
		"in/c/test.asset",
	)
}

func inCTestAsset() (*asset, error) {
	bytes, err := inCTestAssetBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "in/c/test.asset", size: 15, mode: os.FileMode(420), modTime: time.Unix(1558857921, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _inTestAsset = "\x2f\x2f\x20\x73\x61\x6d\x70\x6c\x65\x20\x66\x69\x6c\x65\x0a"

func inTestAssetBytes() ([]byte, error) {
	return bindataRead(
		_inTestAsset,
		"in/test.asset",
	)
}

func inTestAsset() (*asset, error) {
	bytes, err := inTestAssetBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "in/test.asset", size: 15, mode: os.FileMode(420), modTime: time.Unix(1558857921, 0)}
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

var defaultPageFuncMap = template.FuncMap{
	"eq": func(a, b interface{}) bool {
		return a == b
	},
}

// MustParseTemplates is like MustAsset but assumes the input files
// are html templates.
func MustParseTemplates(files ...string) *template.Template {
	if len(files) == 0 {
		panic("MustParseTemplates requires at least one filename")
	}

	// the first file is the root template
	root := template.New(path.Base(files[0]))
	if b, e := Asset(string(files[0])); e != nil {
		panic("failed to load " + files[0] + ": " + e.Error())
	} else if _, e = root.Parse(string(b)); e != nil {
		panic("failed to parse " + files[0] + ": " + e.Error())
	}

	// the remaining files are children templates
	for _, file := range files[1:] {
		child := root.New(path.Base(file))
		if b, e := Asset(string(file)); e != nil {
			panic("failed to load " + string(file) + ": " + e.Error())
		} else if _, e = child.Parse(string(b)); e != nil {
			panic("failed to parse " + string(file) + ": " + e.Error())
		}
	}

	root.Funcs(defaultPageFuncMap)

	return root
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"in/a/test.asset": inATestAsset,
	"in/b/test.asset": inBTestAsset,
	"in/c/test.asset": inCTestAsset,
	"in/test.asset": inTestAsset,
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
	"in": &bintree{nil, map[string]*bintree{
		"a": &bintree{nil, map[string]*bintree{
			"test.asset": &bintree{inATestAsset, map[string]*bintree{}},
		}},
		"b": &bintree{nil, map[string]*bintree{
			"test.asset": &bintree{inBTestAsset, map[string]*bintree{}},
		}},
		"c": &bintree{nil, map[string]*bintree{
			"test.asset": &bintree{inCTestAsset, map[string]*bintree{}},
		}},
		"test.asset": &bintree{inTestAsset, map[string]*bintree{}},
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

