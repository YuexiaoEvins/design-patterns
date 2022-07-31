package Composite

import "fmt"

type Component interface {
	search(string)
}

type File struct {
	name string
}

func (f *File) search(name string) {
	fmt.Printf("Searching for keyword %s in file %s\n", name, f.name)
}

func (f *File) getName() string {
	return f.name
}

type Folder struct {
	name       string
	components []Component
}

func (f *Folder) search(name string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", name, f.name)
	for _, c := range f.components {
		c.search(name)
	}
}

func (f *Folder) getName() string {
	return f.name
}
func (f *Folder) add(c Component) {
	f.components = append(f.components, c)
}
