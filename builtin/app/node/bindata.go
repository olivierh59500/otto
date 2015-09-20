// Code generated by go-bindata.
// sources:
// data/aws-simple/build/build-node.sh
// data/aws-simple/build/template.json.tpl
// data/aws-simple/deploy/main.tf.tpl
// data/common/dev/Vagrantfile.tpl
// DO NOT EDIT!

package nodeapp

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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
	name string
	size int64
	mode os.FileMode
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

var _dataAwsSimpleBuildBuildNodeSh = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x56\x7f\x4f\x1b\x3d\x12\xfe\x7f\x3f\xc5\x34\xa0\x72\x27\xb1\xbb\x80\xae\x77\x6d\x68\xab\x03\x0a\x2d\x52\x15\x2a\xda\xd3\xbd\x52\xdb\x17\x6d\xd6\x93\xc4\x65\x63\xbb\xb6\x37\x10\x28\xdf\xfd\x7d\xc6\x1b\x12\xf2\xf6\x87\xfa\x0f\x38\xb6\x67\xe6\x99\x99\x67\x1e\xef\xc6\xa3\x72\xa8\x4d\x39\xac\xc2\x24\xcb\x02\x47\xca\x2d\x19\xdb\x9a\xc5\x92\xbd\xe7\x6b\x9d\x96\x4e\x3b\x1e\x55\xba\x59\x6c\x47\x5f\xd5\x9c\x65\x58\x59\xff\x8f\x7f\xd2\x6d\x46\x44\x8d\xad\xab\x86\x82\x6d\x7d\xcd\x23\xdd\xf0\x8b\xcd\xdd\xd5\x76\xa3\x0d\x1b\xfb\x62\x73\x4f\xb6\xb8\x9e\x58\xea\x1d\x9f\x9f\x9f\x9d\x53\x15\x69\xf3\x76\x65\x74\xd7\xdf\xbc\xed\xee\xde\xed\xd3\xdb\x2a\x44\xd8\x8f\x43\xbf\x27\x66\x63\xcf\x8e\x6c\x8c\x96\xca\x59\xe5\x4b\x1c\x94\x61\x1e\xf0\x8f\xbe\x51\x4c\xd8\x0c\xed\xed\x64\x77\x19\xd0\x39\xda\x4a\xe0\xa8\xb7\x79\x7b\x78\xf0\xfe\xcd\xc5\xfb\xb3\xff\x9d\x1f\x1d\xdf\xf5\x64\xe3\xed\xe9\xe0\x78\x70\x76\xd7\xdb\x22\x60\xc8\x32\xcb\x92\x02\x0e\xfe\xdb\xa3\xbd\x97\x8f\x77\xe1\x0e\x4e\xc7\xec\x29\x8f\x5d\xbc\x97\x54\x2a\x9e\x95\xa6\x6d\x9a\x7d\xba\xcb\x6c\x93\x0c\xba\x34\x3e\xca\x8d\xcf\x04\x63\x39\xca\x36\xa8\x6e\x6c\xab\xf2\xda\x9a\x91\x1e\x53\x5d\x19\xd2\x26\xb2\x1f\xb1\x67\xba\xd2\x71\x42\x95\x8b\x54\xdb\xe9\xb4\x32\x2a\x90\x1e\x91\x8e\x5b\x81\x42\xd4\x4d\x83\x9b\xe4\xbc\x45\x9e\x21\x20\x08\xf5\xfe\x5f\xe9\xa8\xcd\x98\x46\x48\x64\xcd\x2d\x30\xc1\x85\x6b\x38\x72\x51\x14\xbd\xac\x35\xb0\xa7\x8f\x1f\x29\x1f\x2d\x8a\xa3\x87\x65\xb2\x28\xb5\x09\xb1\x32\x35\x97\x43\x6b\x63\x3e\xd2\x46\x87\x09\x2b\xfa\xfc\x79\x9f\x94\x45\x59\x43\xc3\x28\xeb\x4e\xf1\x24\x53\xd6\xa0\xa7\x12\xf7\x40\x29\x09\x2b\x48\x81\xc5\xd9\xa0\xa3\xf5\x9a\x03\x01\x33\xb5\x4e\x55\x82\x2a\x05\xb6\x4c\xa1\x55\x56\xae\xe6\x63\xb0\x26\x1d\x32\xe5\xf3\xef\x4e\x12\x0e\xe4\x98\xcf\xc1\x91\x51\xbc\xaa\x3c\xe7\x48\xd6\xb1\x8f\xf0\x9c\x4b\x45\xac\x59\x59\x29\x95\x8b\x25\x5a\xde\x45\x9f\x8b\xa1\x73\x55\xbf\x9e\x78\x1d\xf2\x86\xab\xd2\x58\xc5\xc5\x97\xb0\x16\xe9\x92\xe7\xb0\x9d\x51\x2e\xab\xc0\x7e\x86\x2e\x4e\x2e\x5d\xbf\x2c\x97\xbf\x8b\x76\x88\x72\xb5\x05\x22\xf6\x9f\xee\xe0\xa6\xe7\x7a\x96\xae\xd3\x93\x7f\xef\x9e\x3c\x3b\x7c\x76\x74\x70\xf4\xaf\x9d\xc3\xbd\x93\xff\x64\xa9\xc5\x5b\x8a\x87\x34\x89\xd1\x05\xb8\xb1\x21\xe4\x98\x99\x4a\xca\x51\xb8\x49\x1b\xb4\x35\xae\x0a\x81\x0d\x08\x23\x3e\x4b\xc0\x28\x97\x3b\x14\x7d\x1b\xe2\x9c\xa6\x95\x36\x5b\x20\x56\x02\x1a\x99\xa9\xe4\x58\xa7\xab\x1d\xf7\x43\xd1\xe8\x10\x0b\xb5\xb2\x4c\x1b\x0f\x99\xf7\x93\x5a\x77\x1d\x3b\xed\xaa\x2b\x5d\x1b\xa0\x2c\xdb\xf4\xee\xde\xcf\x36\x0d\xc6\xda\x5c\x6f\xa7\xde\xd9\x38\x01\x26\x57\xd5\x97\xd5\x18\x31\xa5\x83\x7c\xed\xac\x8f\xf4\xea\xf8\xf0\xf4\x60\x70\x71\x72\x7e\x36\xf8\x70\x3c\x78\xf5\xc2\x58\x93\x78\x5b\xd5\x51\xcf\xf8\x57\xdd\x1c\xde\x78\x1a\x43\x25\xa6\xec\xeb\xd6\x6b\x0c\xfb\xb0\xd5\x8d\xca\x59\x00\x44\xf9\xfd\x09\x34\x03\x21\xdd\xd7\x1c\xb9\xd0\x0d\x96\xbb\xe3\xb4\xfc\x39\x11\x92\x8d\x04\xc3\x2c\x9b\x20\x08\xf3\xd4\x81\xb4\x2f\x8d\xff\xb2\x58\x4a\x6e\x39\x5f\xe3\x5a\xa0\x65\xed\xba\x9a\x1c\xcb\x6e\x1d\x3b\x26\xbb\x94\x6c\xca\x61\x7a\xa9\x34\x66\xdb\x51\x19\xfc\xac\x94\xf1\x05\xd5\x5c\x77\x16\x2b\x4f\x37\xd7\x18\xa2\x38\x75\xcb\xa3\x22\x8e\x6f\x28\x3f\xfa\xdb\xfd\xf5\x49\x71\x8d\xae\x31\x16\x40\xde\x82\x66\xeb\xb3\xa1\x94\xec\x81\x69\x4a\x87\x6a\xd8\xb0\xca\x05\xe9\x95\xf5\x0a\x7b\x63\xae\x6d\xa0\x5e\x8f\xd6\x1d\xbf\xe7\x98\x90\xa3\x2c\x53\x1d\x84\x65\x61\xcd\x29\x88\x79\x65\x28\x3f\x5f\x9a\xf5\x7f\x04\xef\x28\xc9\x05\xba\x02\x4f\xa9\x54\xc9\x07\x44\xea\xc3\x44\x43\x7c\x02\xc6\xfb\x6b\xab\x3d\xf4\x40\x24\xe6\x01\x6b\x2d\x49\x97\x2b\x9c\x57\xc1\x1a\x01\x4d\x6c\x66\xda\x5b\x33\x45\x53\xe9\x6a\x22\x72\x86\xa6\x43\xdf\xe0\x0d\xaa\xa2\x88\xaf\xb9\x6e\xa3\x5c\x0d\xe8\xf6\x25\x28\xde\x06\x9f\xde\x17\x58\x6e\xaf\x7e\x81\x24\xcd\x36\x81\xfe\x05\x9d\x22\x44\x13\x84\x55\x0e\x2c\x30\xb1\x99\xc3\x99\x61\x86\x30\x02\x81\xad\x71\x95\x26\x7a\x3c\x11\x61\x04\x73\xa9\x53\xbf\x82\x0e\x9c\x63\x93\x0a\x0f\x08\x92\x88\x09\xed\x68\xa4\x6b\x0d\x1f\x05\xf5\xf3\x6f\x5d\x33\x03\xf2\xca\x35\x6d\xed\x86\xf2\x4f\x01\x41\xef\x0e\x3e\xbc\xd9\xff\x64\xca\xad\x6e\xfc\x52\x45\xba\xbf\x85\xb8\xfe\x81\xd5\x06\x9d\xa1\xa0\x7d\x92\x17\x51\xac\x41\xd9\x07\x65\x12\x75\x0f\x18\xe6\x7b\x41\xf8\x85\x6b\x24\x36\x40\x62\x92\x97\xe7\xa9\x9d\x31\x12\x02\x72\xb8\x4b\x97\xa4\xd0\xc8\x1a\x5a\x40\x90\x3b\xee\x90\xf8\xe9\x43\x67\xb2\x1f\x72\x4e\xcd\x50\x50\x85\x51\xd5\x36\xb1\xeb\x25\x07\x4e\x2f\x2c\x04\x1a\x6d\x71\xd0\x6c\x69\x12\x26\x42\x86\x09\xcb\x70\x5f\xc0\x25\xf4\x7c\xa1\x01\x8a\x56\x18\xb7\x31\xb8\x11\xfe\xd2\x83\x84\xbe\xeb\x8e\x08\x0a\xb3\x09\x26\x04\x46\x8f\x44\xdb\xef\x9f\xa0\x09\x92\x8f\x5d\xb9\x6c\x8b\x68\x60\xbe\x59\xc4\x2b\x32\xcc\x02\x3d\x7f\x3e\x78\x7d\x3a\xf8\xe3\xe8\x6c\x70\xf2\x9d\xf2\x75\x29\x89\xab\x35\xcd\x93\x8d\x35\xcd\xdb\xa0\xd7\x6c\x58\xe2\x2a\x1a\xce\x53\x33\xb2\xe5\xf5\x0b\x8f\x27\xad\x23\x96\xbc\x77\xbe\x1d\xce\xcb\x19\x88\x61\x71\x22\xeb\x85\x3c\x5f\x2c\x0d\x4a\xf9\x18\x89\x69\x96\xf0\x0e\xee\x67\x4b\x7c\xd9\x6f\x02\x5e\xef\xc1\x52\x1f\x7e\x0b\xf6\xe2\x35\x4a\xdf\x4a\x24\xe2\xce\x86\x9e\xee\xec\xa7\x9f\x5d\x22\x0f\xe7\xb7\x74\xed\x10\x92\xd2\x1d\xaf\x32\x5e\x84\x26\x6b\xf6\xf1\xa1\xf3\x00\xbf\x4c\xfb\x79\x6b\x4c\x9a\x74\x37\x5d\xa9\x5d\xde\x2e\x25\x42\x68\xbd\xfc\xe0\xa3\xbc\xa9\xa9\x57\xab\xf5\xb0\xf4\xf8\xb1\x98\xaf\xf4\x5d\xa4\x59\xb5\xb5\x14\xad\xd7\x45\x81\x6b\xf9\x52\x78\xd4\xcb\xfe\x0a\x00\x00\xff\xff\x88\x43\xa9\x98\x3f\x0a\x00\x00"

func dataAwsSimpleBuildBuildNodeShBytes() ([]byte, error) {
	return bindataRead(
		_dataAwsSimpleBuildBuildNodeSh,
		"data/aws-simple/build/build-node.sh",
	)
}

func dataAwsSimpleBuildBuildNodeSh() (*asset, error) {
	bytes, err := dataAwsSimpleBuildBuildNodeShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/aws-simple/build/build-node.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataAwsSimpleBuildTemplateJsonTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x52\xed\x6e\xa3\x30\x10\xfc\xcf\x53\x58\x96\xf2\x0f\x48\xee\x4e\xba\x54\x7d\x95\x2a\x22\x06\xb6\x60\xc5\x36\x96\xd7\x4e\x95\x20\xbf\x7b\xcd\x37\x54\x4d\xc2\x1f\xa3\x9d\xf1\xec\x78\x76\xdb\x88\x84\x8f\x4a\xae\x32\xcd\x8a\x0b\x98\xec\x0a\x06\x79\xa3\xe8\x3b\xa1\x87\xf4\x2d\x3d\xd0\x38\x1a\x38\x57\x66\x38\xcb\x05\x60\x80\x86\x6b\xa1\xc8\xbe\x30\x63\x45\x01\x88\xd9\x05\x6e\x01\x51\x4e\x88\x78\x8d\x22\x14\x06\xec\x23\xd4\x40\x35\x34\xdb\x20\x28\x5c\x15\xfc\xd8\x7a\x04\xfa\xba\x9f\x8c\x68\xd3\x5c\x79\xe7\x31\x38\x0d\x84\x8f\xf1\xd6\xe4\x29\x30\xec\x4d\x43\xf7\x80\x4f\x2e\x80\xc6\x4b\x1d\x1b\x67\x8a\x1e\x69\x77\x24\x3c\x34\x67\x96\x4b\xb2\xf3\x6d\x4b\x1c\x82\x21\xe7\xb9\xf1\x99\x78\x1f\x38\xa0\xca\x15\x6d\x2d\x55\x02\x5a\xae\x02\x30\x44\xb5\xb7\x52\xef\x1b\x6b\x9b\x84\x69\x9d\xda\xea\x4e\x47\xaa\x8f\x1f\xdb\xc3\x1a\x84\xd8\xf8\x2b\x0c\xd7\xb6\x83\x72\xc7\x45\x99\xa8\xa6\x84\x14\xeb\x59\xab\x3f\x4f\x53\x0e\x3d\x67\xcc\x60\x1e\x88\x62\xb2\xd7\xee\xbc\xcc\xd2\x73\x47\x26\xd9\xbd\x51\x09\xe4\xb8\x60\x9b\xf1\x3d\x0a\x66\x3b\xe7\xe7\xe9\xd0\xcd\xc8\x9f\x29\x2e\xc4\x17\x8a\xf3\x9a\x3c\x53\x1b\x48\xaf\xbc\xf5\x2b\x90\x31\xc9\x87\x3c\x78\x72\xfc\x9f\xff\x65\xc7\x3f\xcb\xa2\x50\xae\xd0\x32\x15\x58\x53\x6c\xc5\xbf\x54\x30\x53\xad\x28\x88\x75\xd6\x75\x9e\xe2\x76\xb9\x53\xd6\xad\x42\x95\x3c\x9b\xb0\xb6\xed\xfe\xbc\x27\x3f\xbd\x87\x33\x6c\x11\x93\xfa\x37\xc7\xc3\xce\x9f\xa2\xc8\x47\xdf\x01\x00\x00\xff\xff\x2a\xa6\xb3\x8a\xa5\x03\x00\x00"

func dataAwsSimpleBuildTemplateJsonTplBytes() ([]byte, error) {
	return bindataRead(
		_dataAwsSimpleBuildTemplateJsonTpl,
		"data/aws-simple/build/template.json.tpl",
	)
}

func dataAwsSimpleBuildTemplateJsonTpl() (*asset, error) {
	bytes, err := dataAwsSimpleBuildTemplateJsonTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/aws-simple/build/template.json.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataAwsSimpleDeployMainTfTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x55\x4d\x6f\x1b\x21\x10\xbd\xfb\x57\x8c\x48\x4f\x55\xb3\x4e\x73\x8a\x22\xf5\x16\xa9\xb7\xe6\xd2\x5b\x15\xad\x58\x16\xbb\x28\x18\x10\x1f\xae\x56\xd6\xfe\xf7\x0e\xe0\x35\xcb\xda\x49\x5a\x45\xaa\xba\xb9\xc4\x6f\x66\x78\xf0\xde\x30\x5c\xc1\x57\xae\xb8\xa5\x9e\xf7\xd0\x0d\xf0\xe8\xbd\xfe\x04\xbd\x06\xa5\x3d\xf0\x5e\x78\xd8\x51\x15\xa8\x94\xc3\x6a\xb5\xa7\x56\xd0\x4e\x72\x20\x42\x6d\x2c\x6d\x45\x4f\xe0\x30\xce\x60\xfa\xcb\xb5\x94\x31\xee\x5c\xfb\xcc\x87\x0b\x41\xc7\x99\xe5\xfe\x85\xa0\xe5\x5b\xa1\xd5\x22\x80\xa9\xad\xa2\x3b\x9e\xe0\x79\xc1\x4e\x2c\x32\x85\x72\x9e\x2a\xc6\x5b\x3f\x98\x98\x0e\x3d\xdf\xd0\x20\x3d\x7c\x01\xe2\x6f\x1b\xb7\xc3\x43\x10\x98\x57\xb8\xd0\x29\xdc\x8d\x09\x9d\x14\x6c\xb1\xda\xde\xb0\x96\x89\xde\x5e\x80\x8f\xc7\x5e\x19\xab\xf7\xa2\xe7\x36\xed\x1e\xa1\x15\x40\x39\x7c\x64\xfd\x70\xc0\xc2\xa6\x16\x65\x24\x98\x56\x64\xa8\xd3\x0a\x9e\xd2\xb2\x20\x10\xbf\x2a\x2d\xe3\x98\x82\x9b\xb0\xdc\xe9\x60\x59\xd1\x37\x58\xe1\x87\x76\x6b\x75\x30\x04\x08\x97\x5d\xde\x59\xd4\x30\xae\x72\x38\xe4\x7f\xc7\xf1\x1a\x63\xd7\x79\xd1\xc9\xce\xc4\x9a\x8f\x58\x18\xf3\x6f\x0c\x61\x8c\x6f\x91\xcf\xa5\x05\x01\xf0\xfc\x5e\x33\x2d\xf3\xfe\xae\x3f\x27\x70\x63\xf5\xae\x35\xda\xfa\x04\xde\x24\xcc\xeb\x09\x29\x58\xd4\xb6\xed\xa4\x66\xcf\x0e\xb1\x1f\xe4\xa6\x49\x7f\xeb\x1b\xf2\x84\xf1\x31\x92\x09\xf5\x32\x1b\xf1\xcc\x90\x0b\x84\x77\x97\x18\xef\xfe\x8c\xf2\x6d\x35\xa9\x31\x33\x35\x61\xa1\xe7\x5f\x6a\xf9\xda\xf1\xde\x29\x66\x21\x8b\x91\xf1\x7d\x92\xde\xde\x5e\xe0\x3d\x82\xff\xa1\x8b\xff\xb4\x4b\xcf\x5a\x26\xdd\xb7\xb3\x3e\x39\x7d\x6f\x36\x4c\x1e\x49\x6e\x56\x30\xb9\x59\x0d\xab\x6c\x69\xdd\x9f\x93\xf5\xe7\x9d\xdb\xe0\xae\x9a\x48\xf0\x94\xfc\xc8\x63\xd2\x55\x0c\xb1\x68\x8a\x34\xb8\xfd\xe6\xe3\xb1\x00\x2b\xae\xe0\xfb\xe3\xc3\xe3\x3d\x3e\x04\xcf\x1c\xa4\x70\x9e\x2b\x34\x15\xa2\x58\x0e\x98\x56\x1b\xb1\x0d\x36\x8e\x47\xcc\xcd\x61\x9c\x89\x59\x7c\xd9\x15\x4d\xa1\xf6\x31\x86\x66\xd6\x2c\xfa\xe1\x34\xcc\xcf\x1b\xa0\x84\xa6\xf2\x52\x98\x1c\xb9\x82\x07\x6e\xa4\x1e\x80\xa2\x42\x1e\xf4\xa6\x9c\x79\xe1\xd6\x84\xcf\x2d\xc3\x87\xa5\x36\xec\x38\x77\x77\x22\x19\x54\xbd\x32\x25\x5c\xc1\x33\x27\xe3\xed\xaf\xd6\x59\xd8\x88\x89\xd3\x0b\xb7\x20\x9c\xe0\x3c\x2c\xe2\x5d\xae\x4d\xc5\x95\x5f\x71\x3c\x5a\x78\x32\xd0\xd3\xed\x74\x1b\xbe\x9d\x3d\x03\x27\xd9\x74\xf0\x26\x78\x20\xc1\xca\xac\xc4\x9e\xca\x90\x92\x7f\x7a\x6f\xee\xd7\xeb\x4c\x14\x7b\x29\xae\xde\x2b\x97\xf7\xb7\x8e\xef\xd0\xef\x00\x00\x00\xff\xff\x35\x00\x0b\x42\x46\x08\x00\x00"

func dataAwsSimpleDeployMainTfTplBytes() ([]byte, error) {
	return bindataRead(
		_dataAwsSimpleDeployMainTfTpl,
		"data/aws-simple/deploy/main.tf.tpl",
	)
}

func dataAwsSimpleDeployMainTfTpl() (*asset, error) {
	bytes, err := dataAwsSimpleDeployMainTfTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/aws-simple/deploy/main.tf.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataCommonDevVagrantfileTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x55\xdf\x6f\xdb\x46\x0c\x7e\xd7\x5f\xc1\xca\x6d\xd1\x02\x91\x84\x15\x45\x1f\xb2\x26\x58\x96\xa4\x4b\x80\x2d\x29\xea\x74\x2f\xc3\xe0\x9e\x75\x94\x74\xcb\xf9\xee\x76\x77\x72\xec\x26\xfe\xdf\x47\x9e\xe4\xfc\x6a\xbb\x15\x18\xd0\x06\x12\x45\x7e\x24\x3f\x7e\xa4\x27\xf0\x0b\x1a\xf4\x22\xa2\x84\xf9\x1a\xce\x63\xb4\x3b\x20\x2d\x18\x1b\x01\xa5\x8a\x4f\xb2\x49\x36\x81\x8b\x4e\x05\xa0\x7f\xb1\x43\xf8\x5d\xb4\x5e\x98\xd8\x28\x8d\xd0\x3e\x8e\x85\xc6\xfa\xe4\x25\x71\x89\xda\xba\x05\x9a\x08\xb6\x21\x88\xc8\x10\xc2\x39\xad\x6a\x11\x95\x35\x55\x40\xbf\x54\x35\x96\x70\x1a\x21\x74\xb6\xd7\x32\x25\x9d\x23\x74\xc2\xc8\x82\x93\xa3\x2c\xe1\xc2\xc2\xc2\x4a\xd5\xac\x19\x96\x70\xee\xa5\xdf\x81\x3e\x60\xca\x76\xe0\x1c\x1b\xca\x2c\x1b\x3f\x97\xb5\x35\x8d\x6a\x7b\x8f\x2f\xf2\x57\xf9\x4b\xee\xe8\x66\x30\xdd\x64\x00\xc3\x53\xb9\x5c\x94\x73\xbb\x82\x3d\xc8\x3b\x11\x3a\x55\x5b\xef\x2a\xe7\xb1\x56\x01\xdf\xbc\xce\x33\x72\x9c\xc0\x89\x0d\xd4\x80\xd1\x6b\x30\x18\xaf\xac\xbf\x7c\x10\x3e\xda\x20\x77\x5e\x2d\x89\x87\xd9\x68\xc8\x77\x40\xb9\x5d\xc8\xaf\xaf\x99\x88\x99\x72\x33\x21\xa5\xc7\x10\x60\xb3\x19\x81\xa7\x18\x7b\x07\x02\xc2\xda\xd4\xc4\x5f\x63\xb5\x44\x0f\x8d\xb7\x0b\xb0\xbd\x07\x46\x51\xa6\x05\xa9\xa8\xa0\x68\x3d\xb5\x6f\xa1\x5a\x0e\xdd\x3d\xa8\x61\x00\x98\x8d\x00\x9c\xd2\x89\xd8\x95\x5b\x00\x4a\xb8\x03\xf9\x36\x32\xdf\xa1\x58\x00\x7b\x45\x73\xa3\xfa\x6e\xad\xd0\x7a\xdb\xbb\x7b\x96\xa1\xc8\x63\x23\xe6\x34\xe6\xe9\xf4\x04\x44\xcb\xa3\xa4\xf1\x5e\x09\x2f\x19\x38\x58\x1a\x7f\x8c\xfc\x38\x76\x4f\xbd\x3a\x34\x12\x4d\xad\x30\xa4\x0e\xc2\x5d\xa5\x21\x74\xe5\x18\x3d\x1b\xb0\xf6\x20\xfa\x1e\x87\x44\xef\x6c\x6f\x64\xd2\x05\x6c\x27\x37\xbc\xbd\x50\x0d\x08\xb3\x7e\x49\x5e\xd7\xcf\x92\xba\x88\x11\x50\x86\x1e\xb7\x11\x33\xb2\x84\x92\x78\x86\x67\x1b\x72\xe3\xef\x34\xd2\xca\x92\x1c\xab\x3b\xaf\x82\x88\xa1\x70\x6d\xad\x2b\x0f\xc9\x1a\x89\x2c\x1e\xc6\xbf\x53\xc9\x60\x89\x41\x7a\x78\xe0\xea\xbc\x5d\xaa\xc0\x15\xe6\xa1\x43\xad\x79\xe2\x46\x2b\x83\xc4\x61\x2d\x61\x72\x4d\x01\x1b\x78\xfe\x1c\xe6\x24\xad\xf1\xb5\x5a\x08\x65\xca\xd0\xe5\x43\x33\x44\x15\xf7\x43\x45\x27\x0a\x7e\xb5\x42\x82\xd0\x3a\x8d\xbf\xf1\xa2\xe5\xdd\x09\xd0\xa1\xc7\xd4\x37\xb1\xf0\x80\xe0\xf2\x8e\x92\xad\x37\xf3\xc2\x7a\xbb\x8b\x4e\x8c\x70\xe7\xa3\xe5\xc6\x23\x65\xd9\x6c\xbe\x5a\xc1\xa9\x09\x91\x0b\x98\xf7\x8a\x96\x11\xcd\x52\x79\x6b\x38\xea\x7b\x3b\x7f\x1a\x6a\xaf\x5c\x9c\xd1\x9a\x67\x84\x9d\x65\xf7\x0c\x34\x93\xb7\x6f\xa7\x87\x1f\x4e\xdf\x5f\x64\x01\x23\x14\x34\xf9\x09\xf0\x90\x0a\x5c\x61\xbd\x0b\xfc\xb7\x27\x11\xd5\x76\xb1\xa0\x03\x00\x57\x2a\x76\xc4\x45\x74\x7d\x04\x6d\xdb\x96\x8f\x0c\x3d\xf2\x8d\x90\x2a\x38\x2d\xd6\x28\x33\x8b\x2f\x5e\xc2\x35\x3c\xfd\x09\x5e\xed\x3f\xff\x01\x6e\x06\x4f\x0f\x45\x4c\xd0\xb0\x0f\x15\x11\x52\x99\x5e\xeb\x1f\x61\x73\x9b\x91\xbc\x76\xb7\xd8\x82\xe4\x8b\x8d\x5a\x11\xfe\x82\x36\x94\xb4\x99\x59\x9d\x50\xb1\xee\x2c\xe4\x7f\x70\xc4\x9f\x94\x22\x1f\x11\x7e\x13\x97\x08\x2a\xf2\x02\xc4\x4e\x44\xf8\x34\xee\x0c\x90\xc4\x3f\x41\x6b\x49\xfb\xc3\xd6\xea\xb4\xb4\x7c\x9f\xe8\xb4\xb0\x21\xa9\x68\x40\x25\x8d\xdc\xee\x24\xec\x53\x99\x9d\x5d\xe0\xd6\x52\x95\xac\x1a\x5f\x73\xb6\xc3\x71\x1d\x78\xcf\x78\x0f\xd3\xbc\x45\x60\xf9\x52\x17\xca\x64\xb4\x20\x4f\x68\x7d\xd1\x41\xfe\x31\xe0\xd1\xd9\x94\x28\xca\xa1\xc2\x58\x57\x54\x10\xff\x97\xb3\x61\x7a\xb0\x7f\x8f\x0c\x2a\xcb\xd0\x5c\x87\x6a\xee\x05\xde\x40\xe8\xe9\x5a\x46\x44\x28\xc4\x7f\xc1\x10\x80\xc5\x21\x60\x3c\xe7\x4c\x02\xd0\xa5\x8b\xc2\xc7\xac\x51\x19\x71\x09\xf9\x11\xdd\x1b\x4d\xf2\xe6\x1e\xce\xac\x44\x18\xef\xa2\xa1\xe7\xd9\x12\x7d\x52\xd3\x66\x53\x96\x65\x4e\x13\x85\xab\x96\xf5\xf1\x37\x14\xe7\x8f\x68\x61\xff\x92\x90\xcb\xf6\x33\x74\x31\xba\xb0\x5b\x25\xdb\x5f\xa1\xb4\xbe\xad\x48\x16\xb1\x5a\x7e\x1d\x3b\xf9\x15\xdf\xf8\x58\x90\x7c\xfb\x55\xb1\x7a\xf3\x7a\x44\x1f\xca\xfe\x68\xe8\xcd\x6f\x8b\xde\x56\x37\xd0\x23\x48\x63\x87\x50\x59\x47\x95\xae\x3e\x37\xdf\x2e\x74\x80\x9a\x8e\x97\x92\x4e\xfe\xfb\x83\x8b\x93\x07\x58\xda\x40\x11\x12\xd4\x77\x16\x59\xcd\x95\x49\xae\x50\xf5\xc1\x57\xda\xd6\x42\xdf\xda\xfe\x2f\xac\x5b\x7c\x81\xea\x16\x43\x13\xe3\x7d\xe0\x3e\xd2\x89\x28\x68\x5d\xe8\x3a\x28\xa1\x93\x2a\x0d\xdd\xd9\x25\xd2\xaf\x4f\x7d\x49\x4b\x14\x52\x8f\xb8\x72\xd6\x47\x38\x3a\xfe\xf9\xf4\xe0\x6c\xf6\xee\xc3\xf9\xd9\xc5\xf1\xd9\xd1\x9e\xb1\x46\xf1\x05\x16\x35\x87\xdc\x96\x2c\x5c\x2c\x78\xf6\xbd\x93\xfc\x73\x52\xac\xbf\xf8\xa2\xc6\x13\x55\xac\x1f\x97\x90\x8d\xb7\xe5\x9f\x00\x00\x00\xff\xff\x56\x11\x3c\x9b\xd1\x08\x00\x00"

func dataCommonDevVagrantfileTplBytes() ([]byte, error) {
	return bindataRead(
		_dataCommonDevVagrantfileTpl,
		"data/common/dev/Vagrantfile.tpl",
	)
}

func dataCommonDevVagrantfileTpl() (*asset, error) {
	bytes, err := dataCommonDevVagrantfileTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/common/dev/Vagrantfile.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"data/aws-simple/build/build-node.sh": dataAwsSimpleBuildBuildNodeSh,
	"data/aws-simple/build/template.json.tpl": dataAwsSimpleBuildTemplateJsonTpl,
	"data/aws-simple/deploy/main.tf.tpl": dataAwsSimpleDeployMainTfTpl,
	"data/common/dev/Vagrantfile.tpl": dataCommonDevVagrantfileTpl,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"data": &bintree{nil, map[string]*bintree{
		"aws-simple": &bintree{nil, map[string]*bintree{
			"build": &bintree{nil, map[string]*bintree{
				"build-node.sh": &bintree{dataAwsSimpleBuildBuildNodeSh, map[string]*bintree{
				}},
				"template.json.tpl": &bintree{dataAwsSimpleBuildTemplateJsonTpl, map[string]*bintree{
				}},
			}},
			"deploy": &bintree{nil, map[string]*bintree{
				"main.tf.tpl": &bintree{dataAwsSimpleDeployMainTfTpl, map[string]*bintree{
				}},
			}},
		}},
		"common": &bintree{nil, map[string]*bintree{
			"dev": &bintree{nil, map[string]*bintree{
				"Vagrantfile.tpl": &bintree{dataCommonDevVagrantfileTpl, map[string]*bintree{
				}},
			}},
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

