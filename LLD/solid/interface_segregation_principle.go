package solid



type FileWorker interface {
     Reader()
	 Writer()
}

// some filestypes (lets say confidential earning report) can only be read  hence 
// no need multiple interface can be segreagted as 

// Segregating interface so only required interfaces are implenented
type Reader interface {
	Read()
}

type Writer interface {
	Write()
}