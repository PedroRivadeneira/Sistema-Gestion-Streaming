package main

import (
	"fmt"

	"github.com/PedroRivadeneira/Sistema-Gestion-Streaming/src/interfaces"
	"github.com/PedroRivadeneira/Sistema-Gestion-Streaming/src/modelos"
	"github.com/PedroRivadeneira/Sistema-Gestion-Streaming/src/servicios"
)

func main() {
	pelicula, err := modelos.NuevaPelicula("Interestelar", "Ciencia ficción", 2014, 169)
	if err != nil {
		fmt.Println("Error al crear la película:", err)
		return
	}

	serie, err := modelos.NuevaSerie("Stranger Things", "Ciencia ficción", 2016, 4)
	if err != nil {
		fmt.Println("Error al crear la serie:", err)
		return
	}

	// Ambas estructuras pueden tratarse mediante la misma interfaz.
	contenidos := []interfaces.ContenidoGestionable{pelicula, serie}

	for _, contenido := range contenidos {
		fmt.Println(servicios.MostrarContenido(contenido))
	}
}
