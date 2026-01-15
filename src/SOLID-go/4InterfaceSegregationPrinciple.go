package main

/*
	- many small interfaces are better than one big interface
	- don't force classes to implement method they don't use
*/

/* bad example */
type Worker interface {
	Work() error
	Eat() error
}


/* good example */ 
type Workable interface {
	Work() error
}

type Eatable interface {
	Eat() error
}