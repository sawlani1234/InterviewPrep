package solid

import "fmt"

type Book struct {
	title string 
	author string 
}

func (b Book) PrintBook() {
	fmt.Println(b.author + " : " + b.title)
}

// func (b Book) Save() {. 
// 	// Save to db. // Violates SRP
// }



type Repository struct {

}

func (r Repository) Save(obj interface{}) {
		// Save Obj
}

