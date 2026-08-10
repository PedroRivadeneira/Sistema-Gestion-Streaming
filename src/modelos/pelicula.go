package modelos

// Pelicula representa un contenido de tipo película.
type Pelicula struct {
	Contenido
	duracionMinutos int
}

// NuevaPelicula crea una película validando los datos principales.
func NuevaPelicula(titulo, genero string, anio, duracionMinutos int) (Pelicula, error) {
	contenido, err := NuevoContenido(titulo, genero, anio)
	if err != nil {
		return Pelicula{}, err
	}

	pelicula := Pelicula{Contenido: contenido}
	if err := pelicula.SetDuracion(duracionMinutos); err != nil {
		return Pelicula{}, err
	}

	return pelicula, nil
}

// SetDuracion actualiza la duración de la película.
func (p *Pelicula) SetDuracion(duracion int) error {
	if duracion <= 0 {
		return errorDuracionInvalida()
	}
	p.duracionMinutos = duracion
	return nil
}

// Duracion devuelve la duración de la película en minutos.
func (p Pelicula) Duracion() int {
	return p.duracionMinutos
}

// Tipo devuelve el tipo de contenido.
func (p Pelicula) Tipo() string {
	return "Película"
}
