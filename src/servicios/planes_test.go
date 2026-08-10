package servicios

import (
	"testing"

	"github.com/PedroRivadeneira/Sistema-Gestion-Streaming/src/modelos"
)

func TestRegistroPlanesRechazaDuplicados(t *testing.T) {
	registro := NuevoRegistroPlanes()
	plan, err := modelos.NuevoPlan("Premium", 12.99, 4)
	if err != nil {
		t.Fatalf("no se esperaba un error: %v", err)
	}

	if err = registro.Registrar(plan); err != nil {
		t.Fatalf("no se esperaba un error al registrar: %v", err)
	}
	if err = registro.Registrar(plan); err == nil {
		t.Fatal("se esperaba un error al registrar un plan duplicado")
	}
}

func TestRegistroPlanesPlanNoEncontrado(t *testing.T) {
	registro := NuevoRegistroPlanes()
	_, err := registro.Buscar("Inexistente")
	if err == nil {
		t.Fatal("se esperaba un error al buscar un plan inexistente")
	}
}
