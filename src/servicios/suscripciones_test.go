package servicios

import (
	"testing"

	"github.com/PedroRivadeneira/Sistema-Gestion-Streaming/src/modelos"
)

func TestRegistroSuscripcionesEvitaDuplicadosActivos(t *testing.T) {
	usuarios := NuevoRegistroUsuarios()
	planes := NuevoRegistroPlanes()
	suscripciones := NuevaRegistroSuscripciones()

	usuario, err := modelos.NuevoUsuario("Pedro", "pedro@streaming.com", 19)
	if err != nil {
		t.Fatalf("no se esperaba error al crear usuario: %v", err)
	}
	if err := usuarios.Registrar(usuario); err != nil {
		t.Fatalf("no se esperaba error al registrar usuario: %v", err)
	}

	plan, err := modelos.NuevoPlan("Premium", 12.99, 4)
	if err != nil {
		t.Fatalf("no se esperaba error al crear plan: %v", err)
	}
	if err := planes.Registrar(plan); err != nil {
		t.Fatalf("no se esperaba error al registrar plan: %v", err)
	}

	if err := suscripciones.Registrar("PEDRO@STREAMING.COM", "Premium", usuarios, planes); err != nil {
		t.Fatalf("no se esperaba error al crear la primera suscripción: %v", err)
	}

	if err := suscripciones.Registrar("pedro@streaming.com", "Premium", usuarios, planes); err == nil {
		t.Fatal("se esperaba un error al intentar duplicar una suscripción activa")
	}

	if suscripciones.Cantidad() != 1 {
		t.Fatalf("se esperaba 1 suscripción, se obtuvieron %d", suscripciones.Cantidad())
	}
}
