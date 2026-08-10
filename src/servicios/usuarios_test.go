package servicios

import (
	"testing"

	"github.com/PedroRivadeneira/Sistema-Gestion-Streaming/src/modelos"
)

func TestRegistroUsuariosRechazaDuplicados(t *testing.T) {
	registro := NuevoRegistroUsuarios()
	usuario, err := modelos.NuevoUsuario("Pedro", "pedro@streaming.com", 19)
	if err != nil {
		t.Fatalf("no se esperaba un error: %v", err)
	}

	if err = registro.Registrar(usuario); err != nil {
		t.Fatalf("no se esperaba un error al registrar: %v", err)
	}
	if err = registro.Registrar(usuario); err == nil {
		t.Fatal("se esperaba un error al registrar un correo duplicado")
	}
}

func TestRegistroUsuariosUsuarioNoEncontrado(t *testing.T) {
	registro := NuevoRegistroUsuarios()
	_, err := registro.Buscar("noexiste@streaming.com")
	if err == nil {
		t.Fatal("se esperaba un error al buscar un usuario inexistente")
	}
}
