package main

type Nodo[T any] struct {
	dato      T
	siguiente *Nodo[T]
}

func crearNodo[T any](dato T, siguiente *Nodo[T]) *Nodo[T] {
	return &Nodo[T]{dato: dato, siguiente: siguiente}
}

type Pila[T any] struct {
	ultimo *Nodo[T]
}

func crearPila[T any]() *Pila[T] {
	return &Pila[T]{ultimo: nil}
}

func (p *Pila[T]) Push(dato T) {
	p.ultimo = crearNodo(dato, p.ultimo)
}

func (p *Pila[T]) Pop() (dato T, ok bool) {
	if p.ultimo == nil {
		var zero T
		return zero, false
	}
	dato = p.ultimo.dato
	p.ultimo = p.ultimo.siguiente
	return dato, true
}

type Cola[T any] struct {
	primero *Nodo[T]
	ultimo  *Nodo[T]
}

func crearCola[T any]() *Cola[T] {
	return &Cola[T]{primero: nil, ultimo: nil}
}

func (c *Cola[T]) Enqueue(dato T) {
	nuevo := crearNodo(dato, nil)
	if c.primero == nil {
		c.primero = nuevo
		c.ultimo = nuevo
		return
	}
	c.ultimo.siguiente = nuevo
	c.ultimo = nuevo
}

func (c *Cola[T]) Dequeue() (dato T, ok bool) {
	if c.primero == nil {
		var zero T
		return zero, false
	}
	dato = c.primero.dato
	c.primero = c.primero.siguiente
	if c.primero == nil {
		c.ultimo = nil
	}
	return dato, true
}

