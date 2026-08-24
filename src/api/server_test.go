package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PedroRivadeneira/Sistema-Gestion-Streaming/src/modelos"
	"github.com/PedroRivadeneira/Sistema-Gestion-Streaming/src/servicios"
)

func newTestServer() *Server {
	usuarios := servicios.NuevoRegistroUsuarios()
	planes := servicios.NuevoRegistroPlanes()
	catalogo := servicios.NuevoCatalogo()
	suscripciones := servicios.NuevaRegistroSuscripciones()
	return NewServer(&usuarios, &planes, &catalogo, &suscripciones)
}

func doJSONRequest(t *testing.T, server *Server, method, path string, body any) *httptest.ResponseRecorder {
	t.Helper()
	var payload *bytes.Reader
	if body == nil {
		payload = bytes.NewReader(nil)
	} else {
		data, err := json.Marshal(body)
		if err != nil {
			t.Fatal(err)
		}
		payload = bytes.NewReader(data)
	}
	req := httptest.NewRequest(method, path, payload)
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	server.ServeHTTP(res, req)
	return res
}

func TestCrearYBuscarUsuario(t *testing.T) {
	server := newTestServer()

	res := doJSONRequest(t, server, http.MethodPost, "/usuarios", usuarioRequest{
		Nombre: "Pedro",
		Email:  "pedro@streaming.com",
		Edad:   19,
	})
	if res.Code != http.StatusCreated {
		t.Fatalf("esperado 201, recibido %d", res.Code)
	}

	res = doJSONRequest(t, server, http.MethodGet, "/usuarios/pedro@streaming.com", nil)
	if res.Code != http.StatusOK {
		t.Fatalf("esperado 200, recibido %d", res.Code)
	}
}

func TestUsuarioDuplicadoDevuelveConflict(t *testing.T) {
	server := newTestServer()
	body := usuarioRequest{Nombre: "Pedro", Email: "pedro@streaming.com", Edad: 19}
	_ = doJSONRequest(t, server, http.MethodPost, "/usuarios", body)
	res := doJSONRequest(t, server, http.MethodPost, "/usuarios", body)
	if res.Code != http.StatusConflict {
		t.Fatalf("esperado 409, recibido %d", res.Code)
	}
}

func TestCrearYBuscarContenido(t *testing.T) {
	server := newTestServer()

	res := doJSONRequest(t, server, http.MethodPost, "/contenidos", contenidoRequest{
		Tipo:             "pelicula",
		Titulo:           "Interestelar",
		Genero:           "Ciencia ficción",
		Anio:             2014,
		DuracionMinutos: 169,
	})
	if res.Code != http.StatusCreated {
		t.Fatalf("esperado 201, recibido %d", res.Code)
	}

	res = doJSONRequest(t, server, http.MethodGet, "/contenidos/buscar?genero=Ciencia%20ficci%C3%B3n", nil)
	if res.Code != http.StatusOK {
		t.Fatalf("esperado 200, recibido %d", res.Code)
	}
}

func TestCrearSuscripcionRequiereUsuarioYPlan(t *testing.T) {
	server := newTestServer()

	_ = doJSONRequest(t, server, http.MethodPost, "/usuarios", usuarioRequest{
		Nombre: "Pedro",
		Email:  "pedro@streaming.com",
		Edad:   19,
	})
	_ = doJSONRequest(t, server, http.MethodPost, "/planes", planRequest{
		Nombre:        "Premium",
		PrecioMensual: 12.99,
		Pantallas:     4,
	})

	res := doJSONRequest(t, server, http.MethodPost, "/suscripciones", suscripcionRequest{
		EmailUsuario: "pedro@streaming.com",
		NombrePlan:   "Premium",
	})
	if res.Code != http.StatusCreated {
		t.Fatalf("esperado 201, recibido %d", res.Code)
	}
}

func TestSuscripcionDuplicadaDevuelveBadRequest(t *testing.T) {
	server := newTestServer()

	_ = doJSONRequest(t, server, http.MethodPost, "/usuarios", usuarioRequest{
		Nombre: "Pedro",
		Email:  "pedro@streaming.com",
		Edad:   19,
	})
	_ = doJSONRequest(t, server, http.MethodPost, "/planes", planRequest{
		Nombre:        "Premium",
		PrecioMensual: 12.99,
		Pantallas:     4,
	})

	body := suscripcionRequest{EmailUsuario: "pedro@streaming.com", NombrePlan: "Premium"}
	first := doJSONRequest(t, server, http.MethodPost, "/suscripciones", body)
	if first.Code != http.StatusCreated {
		t.Fatalf("esperado 201 en la primera suscripción, recibido %d", first.Code)
	}

	second := doJSONRequest(t, server, http.MethodPost, "/suscripciones", body)
	if second.Code != http.StatusBadRequest {
		t.Fatalf("esperado 400 en la suscripción duplicada, recibido %d", second.Code)
	}
}

func TestEstadisticasConcurrentes(t *testing.T) {
	server := newTestServer()

	usuario, err := modelos.NuevoUsuario("Pedro", "pedro@streaming.com", 19)
	if err != nil {
		t.Fatal(err)
	}
	plan, err := modelos.NuevoPlan("Premium", 12.99, 4)
	if err != nil {
		t.Fatal(err)
	}
	pelicula, err := modelos.NuevaPelicula("Interestelar", "Ciencia ficción", 2014, 169)
	if err != nil {
		t.Fatal(err)
	}

	server.mu.Lock()
	if err := server.usuarios.Registrar(usuario); err != nil {
		t.Fatal(err)
	}
	if err := server.planes.Registrar(plan); err != nil {
		t.Fatal(err)
	}
	if err := server.catalogo.Agregar(pelicula); err != nil {
		t.Fatal(err)
	}
	if err := server.suscripciones.Registrar(usuario.Email(), plan.Nombre(), *server.usuarios, *server.planes); err != nil {
		t.Fatal(err)
	}
	server.mu.Unlock()

	res := doJSONRequest(t, server, http.MethodGet, "/estadisticas/concurrente", nil)
	if res.Code != http.StatusOK {
		t.Fatalf("esperado 200, recibido %d", res.Code)
	}
}
