package servicios

import (
	"testing"

	"github.com/PedroRivadeneira/Sistema-Gestion-Streaming/src/modelos"
)

func TestMostrarContenidoConPelicula(t *testing.T) {
	pelicula, err := modelos.NuevaPelicula("Interestelar", "Ciencia ficción", 2014, 169)
	if err != nil {
		t.Fatalf("no se esperaba un error: %v", err)
	}

	resultado := MostrarContenido(pelicula)
	if resultado != "Película: Interestelar" {
		t.Fatalf("resultado inesperado: %s", resultado)
	}
}

func TestMostrarContenidoConSerie(t *testing.T) {
	serie, err := modelos.NuevaSerie("Stranger Things", "Ciencia ficción", 2016, 4)
	if err != nil {
		t.Fatalf("no se esperaba un error: %v", err)
	}

	resultado := MostrarContenido(serie)
	if resultado != "Serie: Stranger Things" {
		t.Fatalf("resultado inesperado: %s", resultado)
	}
}
