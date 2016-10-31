// Code generated by go-bindata.
// sources:
// templates/Caddyfile
// templates/docker-compose.yml
// templates/setup.sh
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

var _templatesCaddyfile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x91\x41\x6f\xdb\x30\x0c\x85\xef\xfa\x15\x0f\xd0\x75\x8e\x0f\xcd\x29\xb7\x02\xf1\xb6\x02\x43\x03\x34\xe9\xbd\x9a\xcd\x44\x42\x55\x29\x10\xe9\xba\x85\xe1\xff\x3e\x44\x76\xb3\x35\xa9\xb1\x93\xc0\x47\xf2\x3d\x7e\x90\xc6\x7a\x83\xfb\xcd\x0e\xd5\xfa\x6e\x87\xef\x77\xbf\xaa\x6f\xb8\x7d\xdc\x6d\x7e\x54\xf7\xd5\xc3\xed\xae\x5a\x2f\x94\x56\x1a\x0f\xd4\x32\x41\x2c\xc1\xf0\x73\xe1\x02\x8b\xf1\x1e\x66\x2f\x94\x40\x8d\x13\x17\x0e\xb9\x5b\xc7\xb0\x77\x07\xec\x9d\x27\x98\xd0\x20\x51\x91\xda\x70\xb9\xa8\x34\x24\x46\x8f\xce\x89\xc5\x53\xc1\x4f\xab\x1c\xf2\xc9\xbb\x60\xa5\x95\xea\x7b\x2c\x7e\x46\x96\x60\x5e\x08\xc3\x80\x5e\x29\x40\x9f\x06\x1b\x1c\x53\x7c\x7b\x57\x18\x5f\x94\x27\xad\x84\x15\x39\xae\xca\xa9\xe8\x15\x80\x9c\x12\x5b\x19\x27\xb2\x22\xc9\x04\x3e\x9a\x44\x41\x14\x30\x8c\x9e\xe4\x9f\x6d\x4c\xe1\xd2\x76\x92\xcf\xce\x53\xbd\x5a\x2e\x97\xcb\xab\x84\xa9\x39\x1f\x62\x5a\xb1\x57\x87\xb7\x62\xff\x1e\x9e\x8b\xcb\xc3\x5b\xb1\x73\x9e\x7d\x0f\xb7\x47\x88\x82\xc5\x23\xd3\xf6\x06\xc5\x30\xe4\x28\x16\x23\xae\xce\x5f\xc1\x0a\x48\x31\x0a\xca\x57\x93\xca\xae\xeb\xf2\x5e\x01\xf2\x4c\xe7\xf9\xb5\x63\xf3\xdb\x53\x03\xc3\xd8\xde\xa0\x33\x0c\x0a\x59\xc9\xed\x2b\x43\xfd\xb5\x65\x68\x30\x4c\xb0\xb5\xf9\x87\xf2\x03\xb0\x36\xef\x1f\x7c\x33\xa7\x9f\x56\xb7\x94\x5e\x89\xd1\xb9\xe6\x40\x02\x2b\x2f\x3e\x37\xe8\xad\xa6\xa3\xa0\x1c\x75\x9e\x5c\x3e\x81\xfc\x17\x65\x84\x99\x09\xd0\xf3\x11\x67\xb0\x2f\xbe\x61\x50\x7f\x02\x00\x00\xff\xff\x3e\x32\xb0\x75\x48\x03\x00\x00")

func templatesCaddyfileBytes() ([]byte, error) {
	return bindataRead(
		_templatesCaddyfile,
		"templates/Caddyfile",
	)
}

func templatesCaddyfile() (*asset, error) {
	bytes, err := templatesCaddyfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/Caddyfile", size: 840, mode: os.FileMode(420), modTime: time.Unix(1476824429, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesDockerComposeYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x56\xdf\x6f\xe2\x38\x10\x7e\xcf\x5f\x61\x95\x87\x7d\x39\xc8\xee\xb2\x3a\x9d\x22\xf1\x90\x06\x1f\x8d\xc8\x12\x14\x87\x76\xfb\x94\x35\x89\x21\x39\x52\x1b\xd9\x0e\x1c\x42\xfc\xef\x27\x9b\x1f\x35\x21\x50\x74\x6a\xab\x56\xb1\xfd\xcd\xd8\x33\xf3\xcd\x67\xb7\x40\x3f\x04\xa3\x30\x06\xb0\xef\xc7\xe0\x6f\x3f\x80\x7f\x00\x77\x12\x87\x03\x38\x82\x91\x1b\xc3\x7e\xc7\x6a\x59\x2d\x10\x91\x4a\x10\x20\x73\x02\xb0\x58\xb4\x0b\x2a\x24\x2e\x4b\x80\x67\x92\x70\x40\xb2\x42\x16\x74\xae\x57\x53\x46\x67\xc5\x1c\xcc\x8a\x92\x00\x4c\x33\xc0\x49\x9b\x57\xb4\x6e\x68\xb5\x80\x64\xac\x04\xeb\x42\xe6\xe0\x77\x5b\xfc\x76\xf4\x26\x67\xbe\xdb\xc2\x6a\x59\xd6\x8a\x70\x51\x30\xea\x80\x2f\xdf\xbf\x58\x56\xeb\xea\x8f\xd5\x02\x88\xf0\x55\x91\x12\x71\x0b\x65\x89\x03\xc8\xb1\x2c\x00\x52\xbc\x71\x2c\x00\x00\x28\xde\xf0\x9c\x38\x20\x65\x1c\x97\x4b\xce\xfe\x21\xa9\xb4\xd5\xe2\x76\x0b\x3a\x5e\x8e\x29\x25\x25\xd8\xed\x34\x94\x13\x21\x31\x97\x0e\x60\xb4\x3d\xc3\x45\x59\x71\xe2\x7c\xfb\xaa\x97\x08\x5d\x15\x9c\xd1\x37\x42\xe5\xde\x2d\x00\x6d\xf0\x30\x70\xfd\x7e\x4f\x39\x1a\x30\x36\x2f\x89\x4b\x71\xb9\x91\x45\x2a\xfc\x3e\xd8\xed\x1e\xde\x71\xee\x24\x7e\x4a\xbc\xc0\x87\xa3\x38\xf1\xfb\xbd\x14\x6f\xea\x8b\xea\x5f\x18\xf9\xf1\xab\x76\x17\x31\x26\x27\x51\x00\x76\x3b\x1b\x57\x32\xb7\x53\x46\x29\x49\xa5\x61\x04\x83\xe1\x53\x18\x8d\x92\x49\x14\xd4\x2d\x48\xb9\xc8\x19\xa7\xe6\x0e\x68\xd8\x04\xc4\x62\x91\xd5\xcf\x01\x47\xee\x63\x00\xfb\x3d\xc9\x2b\x62\xac\xc5\xd1\x04\xc5\xbd\x19\x2e\xc5\x61\x96\x12\xb9\x66\x7c\x21\xde\x93\x31\xc5\xe9\xa2\x2d\x72\x52\xce\x54\xfa\x95\xef\xeb\xf9\xd7\xab\x9f\x51\x00\x15\xd9\x53\x88\xe2\xde\xd7\x8e\xfe\x75\xfe\xfa\x5a\x8b\x3b\x08\x07\x03\x7f\x34\x48\x02\xf8\x0c\x83\xde\xb7\xda\xea\xcf\x70\x34\x08\x93\x49\xe4\xf7\xde\x18\x9d\xb3\x6c\xea\xd8\xb6\x8e\xa1\xad\xc7\xb6\x24\x42\xd6\x4c\xf6\x19\x4a\xbc\x30\x42\xbd\x38\x9a\xc0\xda\xb2\xce\xe2\x78\xf2\x18\xf8\x5e\x32\x84\xfb\x72\xba\x95\xcc\xc7\xd5\xb4\x2c\xd2\x21\xd9\xd4\x98\x81\x86\x49\x04\x3d\x77\x1c\x7b\x4f\x6e\x82\xa0\x17\xc1\x78\x5f\x27\x92\xe2\xa5\x4c\x73\x8c\x48\xca\x89\x3c\x59\xdd\x4c\xbc\x9a\xc8\xc8\x92\xd0\x4c\x24\x8c\xbe\x43\x8c\x90\x54\x71\x0e\x0c\xb9\x5e\x9f\x23\xe0\xb3\x4a\xa4\xc8\x97\x4b\xb9\x74\xec\x3a\xe9\x50\x37\x79\x9c\x78\xc3\x43\xd0\xa8\xfb\x58\xa5\x0b\x23\xda\x13\xc6\x45\xf0\x44\x61\xd4\x85\x34\x5b\xb2\x82\xd6\x70\xee\x0b\x4a\x22\x38\xf0\xc3\xd1\x3e\xeb\x2f\x28\x22\xf3\x82\xd1\x4b\x94\xeb\x79\x10\x21\x55\x9f\xe4\xd0\xc0\xee\x0b\x72\xd3\x94\x08\x31\x24\x9b\x8b\xee\x3d\xb3\xb8\x80\x9f\x83\x4f\xc5\xbc\x59\x45\x0d\xd5\x2c\x41\x4f\xb0\x7f\x16\x5e\x31\x03\x94\x49\xd0\x99\x08\x82\xba\x60\xb7\xab\x35\xed\xba\xc8\xe6\x44\x0a\x7b\xbb\x05\x84\x66\x27\x87\x0d\x86\x7a\x7e\xc5\xca\xea\x8d\x18\x6c\xe9\x9c\x8a\x6b\x57\x82\xdb\x82\xa7\x36\x5e\x2e\x6d\xa5\x9d\x84\x1f\xbd\x1f\x5d\xee\x77\xf8\xff\xbc\x53\xc5\xd6\x6a\x50\xc9\xfc\x86\x1a\xab\x41\x5b\x63\x3e\x83\x70\x5e\x18\xb9\xc1\xa1\x0f\xc3\x28\xee\x9d\x49\x82\xb1\xa8\x35\x2d\x19\x47\xe1\xaf\xd7\x7a\x23\x9b\xa8\x70\x08\x47\x09\xfc\x35\xf6\xa3\xd7\x24\xf6\x7f\xc2\x5e\xf7\x4f\x90\xf3\x66\xf0\x51\x4e\x02\x43\x4e\x54\x5c\xcd\x6a\x62\x1e\x34\xf2\x9f\xdd\x18\x9e\x2b\x06\x2f\x56\x58\x92\x0b\x86\x99\x66\x77\xeb\x8c\x61\x84\x20\x42\x7e\x38\x32\xc5\x06\x11\xa1\x6e\xe2\x26\x92\x1a\x86\x6e\x10\x84\x2f\xb0\x7f\xb8\xc8\x90\xba\xc6\x40\x8d\x9d\x29\x2e\x4b\xc5\x8b\x66\x07\x51\x18\xc6\x8d\x17\x51\x25\xf3\x9b\xea\xa6\x00\xd7\x49\x76\x4a\xb0\xa2\xda\x8d\x07\x02\x00\x2d\xd0\xc7\x12\x4f\xb1\x20\xe2\x36\xd4\x02\xa6\x66\x9e\x51\x77\x3f\xd3\xed\x7c\xff\x88\x9d\x17\xad\x67\x7a\xb4\x33\x2c\xb1\x9d\x4d\xef\xbb\x4c\x4f\x11\x7e\xd2\x49\x0c\x87\xf7\x1c\x44\xa7\xff\xe3\xdc\xbe\x90\xa9\x7e\xa8\x11\xfe\x71\x76\x53\x9c\x65\xe7\x0f\x34\x3c\x2d\x98\x60\x33\xf5\x38\x53\x4b\x25\x56\xdd\x72\x4d\xc3\x3c\x85\x51\x0f\x51\xc7\x26\x32\x7d\x1f\x1e\x20\xdb\x6d\xbb\x59\x0e\x6b\x0a\xb8\xc2\xdc\x5e\xaf\xd7\x67\xb2\x57\x17\x3e\x65\xa2\x8f\xd4\x4e\x09\x97\xc2\xb1\x39\x63\xd2\xee\xe8\x29\x0d\x58\x32\x35\xfd\x4e\xf8\xe3\x53\x44\x91\x7c\xcc\xb8\xea\x27\xf3\xfb\xe1\xfc\x8c\xfa\x7c\x28\x30\x77\x3b\x79\xf8\xf1\xa3\xab\xfe\x1e\x1a\xcf\xd5\x54\x2b\x21\x0b\x46\xaf\xb4\xce\xbd\x92\x6d\xe2\xb5\x7e\x1f\x07\x87\xac\x19\x49\xd9\xdc\xf7\x6a\x07\xa3\xc3\x51\x6f\xbe\xde\xcd\x78\x0e\x91\xec\x3f\x8f\x47\x3e\x8d\xf6\xb7\xc9\x7d\x5b\x3f\xef\xb9\x73\x73\x67\x83\x5f\x66\xa1\xf7\x57\xcc\xbf\x92\x70\x8a\x4b\x07\xe8\xc7\x6f\x43\x33\x5e\x22\x2e\x94\xa3\x0e\xf9\x2f\x00\x00\xff\xff\x7e\x88\x3e\xf3\x98\x0d\x00\x00")

func templatesDockerComposeYmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesDockerComposeYml,
		"templates/docker-compose.yml",
	)
}

func templatesDockerComposeYml() (*asset, error) {
	bytes, err := templatesDockerComposeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/docker-compose.yml", size: 3480, mode: os.FileMode(420), modTime: time.Unix(1477946301, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSetupSh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x90\xc1\x6e\xc3\x20\x10\x44\xef\x7c\xc5\x94\x9c\x71\xee\xb9\xb6\xbd\x55\x51\xa4\xa8\x1f\x40\xf1\xa6\x46\xc5\x60\xb1\x38\x6e\x64\xf9\xdf\x2b\x70\xdc\x3a\xee\x09\xc4\xcc\x32\x6f\x76\xf7\xb4\xff\xb0\x7e\xcf\x8d\x10\x4c\x09\x8a\x84\x20\xd3\x04\xc8\x53\xef\x9c\xf5\x9f\x60\x8a\x57\x6b\x08\x4e\xdf\x28\xf2\x41\xde\x65\x29\x44\x1d\xcc\x17\x45\x65\x42\xdb\x05\x26\x74\xbd\x73\xcb\xf0\x5b\x31\x97\x27\xaa\xab\xd5\xcc\x7c\x79\x8e\xa4\x53\xfe\x3c\x35\xb4\x04\x30\x7c\x18\x0e\xbf\x16\x29\xc4\x0e\xe7\xa4\x63\x7a\x30\x55\xdb\xd4\xbe\x83\xaa\xff\x86\xe6\xf3\xbc\x7c\x69\x72\x50\x21\x18\x47\xd8\x0b\xaa\x93\x66\x1e\x42\xac\x31\x4d\x9b\x99\x63\x18\x66\x7b\xe6\xea\x99\xe2\xba\xea\x0e\x05\x99\x0a\x4b\x16\xff\x71\xc4\xde\x43\xa9\xd8\x42\xf7\xa9\x81\x71\xf6\x9e\x0d\x75\x81\x52\xd4\x6a\xeb\x20\xc7\x11\xd5\x6b\xb9\x4e\x93\x84\x52\xdd\x42\x53\x94\x15\x5b\x16\xbd\x6e\x69\x16\x5e\x2c\x77\x4e\xdf\x8e\xf9\x61\x9a\xe4\xb6\xed\x3b\x53\xc4\xa0\xd7\x6d\x1f\x0d\x27\x47\x3a\x23\x52\x1b\xae\xb9\x82\x65\x5c\xac\x23\x68\x86\x4d\x30\xc1\x27\x6d\x3d\x97\x6e\x9d\xd3\xd6\x27\xfa\x4e\x58\xd8\xe6\xdd\x91\x2f\x2b\xfb\x09\x00\x00\xff\xff\x1b\x30\xfc\xf6\x2e\x02\x00\x00")

func templatesSetupShBytes() ([]byte, error) {
	return bindataRead(
		_templatesSetupSh,
		"templates/setup.sh",
	)
}

func templatesSetupSh() (*asset, error) {
	bytes, err := templatesSetupShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/setup.sh", size: 558, mode: os.FileMode(420), modTime: time.Unix(1476995253, 0)}
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
	"templates/Caddyfile": templatesCaddyfile,
	"templates/docker-compose.yml": templatesDockerComposeYml,
	"templates/setup.sh": templatesSetupSh,
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
		"Caddyfile": &bintree{templatesCaddyfile, map[string]*bintree{}},
		"docker-compose.yml": &bintree{templatesDockerComposeYml, map[string]*bintree{}},
		"setup.sh": &bintree{templatesSetupSh, map[string]*bintree{}},
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

