package solid

// import "reflect"


type Shape interface{} 


type Square struct {
	length int 
  
}

type Rectangle struct {
	length int 
	breadth int 
}


// Violates Open Closed Principle for getting area
// // Calculates area of given shape 
// type Area struct {

// }


// func (a Area) getArea(s Shape) int {

// 	  if reflect.TypeOf(s) == reflect.TypeOf(Square{}) {
// 		  return s.(Square).length * s.(Square).length

// 	  } else if reflect.TypeOf(s) == reflect.TypeOf(Rectangle{}) {
// 		return s.(Rectangle).length * s.(Rectangle).breadth
// 	  }
   

// 	  return -1 
// }

// Now each new shape will just implement new funciton to get its area
type Area interface {
	GetArea() int 
}

func (s Square) GetArea() int {
	return s.length*s.length
}

func (r Rectangle) GetArea() int {
	return r.length * r.breadth
}