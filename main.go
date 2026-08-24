package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PedroRivadeneira/Sistema-Gestion-Streaming/src/api"
	"github.com/PedroRivadeneira/Sistema-Gestion-Streaming/src/modelos"
	"github.com/PedroRivadeneira/Sistema-Gestion-Streaming/src/servicios"
)

func main() {
	usuarios := servicios.NuevoRegistroUsuarios()
	planes := servicios.NuevoRegistroPlanes()
	suscripciones := servicios.NuevaRegistroSuscripciones()
	catalogo := servicios.NuevoCatalogo()

	// Datos iniciales mínimos para que la API pueda demostrarse inmediatamente.
	// Las suscripciones se crean mediante el endpoint POST /suscripciones
	// para poder demostrar esa funcionalidad durante la ejecución.
	usuario, err := modelos.NuevoUsuario("Pedro", "pedro@streaming.com", 19)
	if err != nil {
		log.Fatal(err)
	}
	if err := usuarios.Registrar(usuario); err != nil {
		log.Fatal(err)
	}

	plan, err := modelos.NuevoPlan("Premium", 12.99, 4)
	if err != nil {
		log.Fatal(err)
	}
	if err := planes.Registrar(plan); err != nil {
		log.Fatal(err)
	}

	pelicula, err := modelos.NuevaPelicula("Interestelar", "Ciencia ficción", 2014, 169)
	if err != nil {
		log.Fatal(err)
	}
	serie, err := modelos.NuevaSerie("Stranger Things", "Ciencia ficción", 2016, 4)
	if err != nil {
		log.Fatal(err)
	}
	if err := catalogo.Agregar(pelicula); err != nil {
		log.Fatal(err)
	}
	if err := catalogo.Agregar(serie); err != nil {
		log.Fatal(err)
	}

	server := api.NewServer(&usuarios, &planes, &catalogo, &suscripciones)

	fmt.Println("=== SISTEMA DE GESTIÓN DE STREAMING ===")
	fmt.Println("API REST iniciada en http://localhost:8080")
	fmt.Println("GET  /                     -> documentación básica")
	fmt.Println("GET  /health               -> estado del servicio")
	fmt.Println("POST /usuarios             -> registrar usuario")
	fmt.Println("GET  /usuarios             -> listar usuarios")
	fmt.Println("GET  /usuarios/{email}     -> consultar usuario")
	fmt.Println("POST /planes               -> registrar plan")
	fmt.Println("GET  /planes               -> listar planes")
	fmt.Println("GET  /planes/{nombre}      -> consultar plan")
	fmt.Println("POST /contenidos           -> registrar contenido")
	fmt.Println("GET  /contenidos           -> listar catálogo")
	fmt.Println("GET  /contenidos/buscar    -> búsqueda por filtros")
	fmt.Println("POST /suscripciones        -> crear suscripción")
	fmt.Println("GET  /suscripciones        -> listar suscripciones")
	fmt.Println("GET  /estadisticas         -> estadísticas del sistema")
	fmt.Println("GET  /estadisticas/concurrente -> concurrencia con goroutines")

	log.Fatal(http.ListenAndServe(":8080", server))
}
