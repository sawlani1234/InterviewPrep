/*

1. Its a creational pattern used when we need copy of the objects
2. It saves to manually copy when there are thousands of instance variable/struct variable one by one
otherwise. u may miss to copy some variable
3. some variables might be private which hence it wont be possib;e to set those variables
4. Prootype pattern comes to rescue which says thaat class itself has an Clone() implementation
and it is the repsonsbility of class to provide object clone

5. In go mostly use shallow copy for direct assignment or deep copy 3rd party methods 

*/

package main

import "encoding/json"
import "fmt"


type Prototype interface {
	Clone() interface{}
}


type Student struct {
	name string
	roll_no int 
	address string
}

func NewStudent(name string ,roll_no int, address string) Student {
	return Student{name,roll_no,address}
}

func (s Student) Clone() interface{} {
	return NewStudent(s.name,s.roll_no,s.address)
}

func Jsonify(abc any) string {
	by,err := json.Marshal(abc)
	if err != nil {
		fmt.Println(err)
	}

	return string(by)
}


func main() {
	s := NewStudent("Shubham",661,"Dakshin")
	fmt.Println(s)
	fmt.Println(s.Clone())
}
