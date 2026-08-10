package main

import (
	"fmt"

	"github.com/PedroRivadeneira/Sistema-Gestion-Streaming/src/modelos"
	"github.com/PedroRivadeneira/Sistema-Gestion-Streaming/src/servicios"
)

func main() {
	usuarios := servicios.NuevoRegistroUsuarios()
	planes := servicios.NuevoRegistroPlanes()
	suscripciones := servicios.NuevaRegistroSuscripciones()
	catalogo := servicios.NuevoCatalogo()

	// Se crean los datos principales del sistema utilizando validaciones.
	usuario, err := modelos.NuevoUsuario("Pedro", "pedro@streaming.com", 19)
	if err != nil {
		fmt.Println("Error al crear usuario:", err)
		return
	}
	if err = usuarios.Registrar(usuario); err != nil {
		fmt.Println("Error al registrar usuario:", err)
		return
	}

	plan, err := modelos.NuevoPlan("Premium", 12.99, 4)
	if err != nil {
		fmt.Println("Error al crear plan:", err)
		return
	}
	if err = planes.Registrar(plan); err != nil {
		fmt.Println("Error al registrar plan:", err)
		return
	}

	pelicula, err := modelos.NuevaPelicula("Interestelar", "Ciencia ficción", 2014, 169)
	if err != nil {
		fmt.Println("Error al crear película:", err)
		return
	}
	serie, err := modelos.NuevaSerie("Stranger Things", "Ciencia ficción", 2016, 4)
	if err != nil {
		fmt.Println("Error al crear serie:", err)
		return
	}

	// El catálogo utiliza una interfaz común para almacenar películas y series.
	if err = catalogo.Agregar(pelicula); err != nil {
		fmt.Println("Error al agregar película:", err)
		return
	}
	if err = catalogo.Agregar(serie); err != nil {
		fmt.Println("Error al agregar serie:", err)
		return
	}

	fmt.Println("CATÁLOGO")
	for _, contenido := range catalogo.Listar() {
		fmt.Println(servicios.MostrarContenido(contenido))
	}

	// Se prueba la búsqueda del catálogo y su manejo de errores.
	resultados, err := catalogo.Buscar("interest")
	if err != nil {
		fmt.Println("Error en búsqueda:", err)
	} else {
		fmt.Println("Resultados de búsqueda:", len(resultados))
	}

	// La suscripción comprueba que el usuario y el plan existan antes de registrarse.
	if err = suscripciones.Registrar(usuario.Email(), plan.Nombre(), usuarios, planes); err != nil {
		fmt.Println("Error al registrar suscripción:", err)
		return
	}

	fmt.Println("Contenidos registrados:", catalogo.Cantidad())
	fmt.Println("Suscripciones registradas:", suscripciones.Cantidad())
}
