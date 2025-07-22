package main

import (
	"fmt"
	"sync"
	"time"
)

/*
1. Singleton pattern says that each class/struct should have only one object . this case is used in
lets say database initalisation , file obj etc

2. It also provides global access to that instance

Downside

1. Violates SRP as it tries do two things described above
2.


*/
/*
1. Singleton pattern using sync.Mutex
2. Singleton pattern using sync.Once
3. Makes Unit Test dificult owing to global state
*/


var DbIInitSingleton *DbInit

var lock = &sync.Mutex{}
var once = &sync.Once{}

type DbInit struct {

}

func initialiseDbInstance(i int) {
     
	lock.Lock()
	defer lock.Unlock()

	if DbIInitSingleton == nil {
		fmt.Println("created instance fuck yeah by :v",i)
		DbIInitSingleton = &DbInit{}
	} else {
		fmt.Println("Instance already created : ",i)
	}

	
}

var DbIInitSingletonV2 DbInit

func initialiseDbInstanceV2(i int) {

	once.Do(func() {
		fmt.Println("created instance fuck yeah by :v",i)
		DbIInitSingleton = &DbInit{}
	})
}

func main() {


	for i:=0;i<20;i++ {
		go initialiseDbInstance(i)
	}

	time.Sleep(5*time.Second)

	for i:=0;i<20;i++ {
		go initialiseDbInstanceV2(i)
	}

   time.Sleep(5*time.Second)

}