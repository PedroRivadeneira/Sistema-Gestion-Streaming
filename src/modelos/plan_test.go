package modelos

import "testing"

func TestNuevoPlanValido(t *testing.T) {
	plan, err := NuevoPlan("Premium", 12.99, 4)
	if err != nil {
		t.Fatalf("no se esperaba un error: %v", err)
	}
	if plan.Nombre() != "Premium" {
		t.Fatalf("nombre inesperado: %s", plan.Nombre())
	}
}

func TestNuevoPlanInvalido(t *testing.T) {
	_, err := NuevoPlan("", 12.99, 4)
	if err == nil {
		t.Fatal("se esperaba un error para un nombre vacío")
	}
}
