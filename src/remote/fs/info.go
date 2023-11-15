package fs

import (
	iofs "io/fs"
	"os"
	"time"

	pb "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"
)

// info represents information about a file/directory
type info struct {
	name     string
	isDir    bool
	size     int64
	modTime  time.Time
	mode     os.FileMode
	typeMode os.FileMode
}

func (i *info) Type() iofs.FileMode {
	return i.typeMode
}

func (i *info) Info() (iofs.FileInfo, error) {
	return i, nil
}

func (i *info) Name() string {
	return i.name
}

func (i *info) Size() int64 {
	return i.size
}

func (i *info) Mode() iofs.FileMode {
	return i.mode
}

func (i *info) ModTime() time.Time {
	return i.modTime
}

func (i *info) IsDir() bool {
	return i.isDir
}

func (i *info) Sys() any {
	return nil
}

// withProperties safely sets the node info if it's available.
func (i *info) withProperties(nodeProperties *pb.NodeProperties) *info {
	if nodeProperties == nil {
		return i
	}

	if nodeProperties.UnixMode != nil {
		i.mode = os.FileMode(nodeProperties.UnixMode.Value)
	}

	if nodeProperties.Mtime != nil {
		i.modTime = nodeProperties.Mtime.AsTime()
	}
	return i
}