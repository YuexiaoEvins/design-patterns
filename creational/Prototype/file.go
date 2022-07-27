package Prototype

type INode interface {
	Path() string
	Clone() INode
}

type File struct {
	path string
}

func (f *File) Path() string {
	return f.path
}

func (f *File) Clone() INode {
	return &File{path: f.path + "_clone"}
}
