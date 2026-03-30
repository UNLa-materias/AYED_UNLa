package main

import "fmt"

type Destino struct {
	ciudad   string
	latitud  float32
	longitud float32
}

func crearDestino(ciudad string, latitud, longitud float32) *Destino {
	return &Destino{ciudad: ciudad, latitud: latitud, longitud: longitud}
}

func (d *Destino) GetCiudad() string           { return d.ciudad }
func (d *Destino) GetLatitud() float32         { return d.latitud }
func (d *Destino) GetLongitud() float32        { return d.longitud }
func (d *Destino) SetCiudad(c string)          { d.ciudad = c }
func (d *Destino) SetLatitud(lat float32)      { d.latitud = lat }
func (d *Destino) SetLongitud(lon float32)     { d.longitud = lon }

func mostrarDestino(d *Destino) {
	fmt.Printf("\n---------DESTINO---------\n")
	fmt.Printf("\nCIUDAD: %s\n", d.ciudad)
	fmt.Printf("\nLATITUD: %f\n", d.latitud)
	fmt.Printf("\nLONGITUD: %f\n", d.longitud)
}

