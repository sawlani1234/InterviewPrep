package filesystem

import "fmt"



type Directory struct {

	name string
	FileSystems []FileSystem
}

func NewDirectory(name string,FileSystems []FileSystem) Directory {
	return Directory{name,FileSystems}
}


func (d Directory) Ls() {

	fmt.Println("Directory name :",d.name)
	for _,fs := range d.FileSystems {
		fs.Ls()
	}
}