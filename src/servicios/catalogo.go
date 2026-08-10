package servicios

import (
	"errors"
	"strings"

	"github.com/PedroRivadeneira/Sistema-Gestion-Streaming/src/interfaces"
)

// Catalogo administra películas y series mediante una colección polimórfica.
type Catalogo struct {
	contenidos []interfaces.ContenidoGestionable
}

// NuevoCatalogo crea un catálogo vacío.
func NuevoCatalogo() Catalogo {
	return Catalogo{contenidos: make([]interfaces.ContenidoGestionable, 0)}
}

// Agregar incorpora un contenido al catálogo y evita valores nulos.
func (c *Catalogo) Agregar(contenido interfaces.ContenidoGestionable) error {
	if contenido == nil {
		return errors.New("no se puede agregar un contenido vacío")
	}
	c.contenidos = append(c.contenidos, contenido)
	return nil
}

// Buscar devuelve los contenidos cuyo nombre coincide parcialmente con el texto indicado.
func (c Catalogo) Buscar(nombre string) ([]interfaces.ContenidoGestionable, error) {
	nombre = strings.TrimSpace(strings.ToLower(nombre))
	if nombre == "" {
		return nil, errors.New("el nombre de búsqueda no puede estar vacío")
	}

	resultados := make([]interfaces.ContenidoGestionable, 0)
	for _, contenido := range c.contenidos {
		if strings.Contains(strings.ToLower(contenido.Nombre()), nombre) {
			resultados = append(resultados, contenido)
		}
	}

	if len(resultados) == 0 {
		return nil, errors.New("no se encontraron contenidos")
	}
	return resultados, nil
}

// Listar devuelve todos los contenidos registrados.
func (c Catalogo) Listar() []interfaces.ContenidoGestionable {
	return c.contenidos
}

// Cantidad devuelve el número de contenidos del catálogo.
func (c Catalogo) Cantidad() int {
	return len(c.contenidos)
}
