package main

import (
	"fmt"
	"iterator_pattern/users"
)


/*
1. Iterator pattern use this to iterator over complex data structures lets say heap or grapg 
or tree 

2. Below is the simple example of iterator over user list , so we define a linear traversal iterator for 
user list , userItertor which is imolementing Iterator interface 

3. Now we have a User Collection  calss which has the iterator and it also implements craateIterator for Collection interface 

4. we can modify the class to change iterator in run time so we can suppor different type of reporter 
for eg if we have something like UserheapIterator added , we can pass this as parametere in CreateIterator function 
to leverage the iterator functionality provided through interface 

*/

func main() {

	dummyUsers := []users.User{}
	dummyUsers = append(dummyUsers,
		users.NewUser(1,"Shubham"),users.NewUser(2,"Nishi"),users.NewUser(3,"Kaur"))

	itr := users.NewUserCollection(dummyUsers).CreateIterator()

	for ;itr.HasNext(); {
		fmt.Println(itr.Next())
	}
}