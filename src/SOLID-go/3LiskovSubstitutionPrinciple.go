package main

/*
	- child classess must be usable without breaking parent behaviour
	- A type implementing an interface must be fully substitutable
*/

/* bad example */
type BirdA interface {
	Fly() error
}

type PenguinA struct {}
func (p *PenguinA) Fly() error {
	return nil
}


/* good example */
type BirdB interface {}

type Flyable interface {
	Fly() error
}

type Sparrow struct {}
func (s *Sparrow) Fly() error {
	return nil
}

type PenguinB struct{}