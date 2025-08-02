package filesystem

import "fmt"


type File struct  {
	name string 
}

func NewFile(name string) File {
	return File{name}
}

func (f File) Ls() {
	fmt.Println("file name ",f.name)
}