package servicios

import (
	"testing"

	"github.com/PedroRivadeneira/Sistema-Gestion-Streaming/src/modelos"
)

func TestCatalogoAgregarYBuscar(t *testing.T) {
	catalogo := NuevoCatalogo()
	pelicula, err := modelos.NuevaPelicula("Interestelar", "Ciencia ficción", 2014, 169)
	if err != nil {
		t.Fatalf("no se esperaba un error: %v", err)
	}
	if err := catalogo.Agregar(pelicula); err != nil {
		t.Fatalf("no se esperaba un error: %v", err)
	}

	resultados, err := catalogo.Buscar("interest")
	if err != nil {
		t.Fatalf("no se esperaba un error: %v", err)
	}
	if len(resultados) != 1 {
		t.Fatalf("se esperaba un resultado, se obtuvieron %d", len(resultados))
	}
}

func TestCatalogoBuscarSinResultados(t *testing.T) {
	catalogo := NuevoCatalogo()
	_, err := catalogo.Buscar("inexistente")
	if err == nil {
		t.Fatal("se esperaba un error cuando no existen resultados")
	}
}
