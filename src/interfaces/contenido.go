package interfaces

// ContenidoGestionable define el comportamiento común de los contenidos del catálogo.
// Pelicula y Serie implementan esta interfaz de forma implícita.
type ContenidoGestionable interface {
	Titulo() string
	Nombre() string
	Genero() string
	Anio() int
	Disponible() bool
	Tipo() string
}
