package main 

import "composite_pattern/file_system"

/*
1. Composite pattern is used when there is a tree like structure  for objects
2. When one object can have similar objects as child objects 
3. For eg Box can have -> products, box , and again box can have products , box , 
4. FileSystem -> Directory -> Directory,File -> Directory,File ,, where file is the terminal/leaf node in the three structure
5. Now in File System we want to print all relevant direcotries and file for a given directory 
6. So for File Ssytem when we to implement LS() command , we will have two objects File and Directory , where Direcotry has lsit of files
both have Ls() function to print there subobjects name , So if we do not have composite pattern we will have something like Directory having array of objects (directory,files} as an interface type
and when we have compute Ls we call Ls by derefencing each object type like reflect.TypeOf(directory) -> then direcotry.ls , same for file 
7. This now creates if else branch 
8. To solve this we create an interface for this common function Ls()
9. Both File and directory should implement it
10. So file implements FileSystem , Directory -> FileSystem and now in direcotyr we can add multiple type of subobjects(directory and file
in this case) and call there Ls() method no if else required

11. *** Use composite pattern when you know group of objects child objects have common functionalitity like we see in this case

*/

func main() {

	file1 := filesystem.NewFile("Son of Sardar 2")
	file2 := filesystem.NewFile("Phir Hera Pheri")
	file3 := filesystem.NewFile("Ramayana")

	comedyDirectory := filesystem.NewDirectory("comedy",[]filesystem.FileSystem{file1,file2})
	religiousDirectory := filesystem.NewDirectory("religious",[]filesystem.FileSystem{file3})

	movieDirectory := filesystem.NewDirectory("movies",[]filesystem.FileSystem{comedyDirectory,religiousDirectory})

	movieDirectory.Ls()

}