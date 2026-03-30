package main

type Nodo[T any] struct {
	dato      T
	siguiente *Nodo[T]
}

func crearNodo[T any](dato T, siguiente *Nodo[T]) *Nodo[T] {
	return &Nodo[T]{dato: dato, siguiente: siguiente}
}

