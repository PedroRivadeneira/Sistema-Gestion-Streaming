package modelos

import "errors"

// Contenido representa la información común de cualquier elemento del catálogo.
// Sus campos privados permiten mantener la encapsulación y controlar los cambios mediante métodos.
type Contenido struct {
	titulo       string
	genero       string
	anio         int
	disponible   bool
}

// NuevoContenido crea un contenido validando sus datos básicos.
func NuevoContenido(titulo, genero string, anio int) (Contenido, error) {
	contenido := Contenido{}

	if err := contenido.SetTitulo(titulo); err != nil {
		return Contenido{}, err
	}
	if err := contenido.SetGenero(genero); err != nil {
		return Contenido{}, err
	}
	if err := contenido.SetAnio(anio); err != nil {
		return Contenido{}, err
	}

	contenido.disponible = true
	return contenido, nil
}

// SetTitulo actualiza el título después de comprobar que no esté vacío.
func (c *Contenido) SetTitulo(titulo string) error {
	if titulo == "" {
		return errors.New("el título no puede estar vacío")
	}

	c.titulo = titulo
	return nil
}

// SetGenero actualiza el género después de comprobar que no esté vacío.
func (c *Contenido) SetGenero(genero string) error {
	if genero == "" {
		return errors.New("el género no puede estar vacío")
	}

	c.genero = genero
	return nil
}

// SetAnio actualiza el año de lanzamiento con una validación básica.
func (c *Contenido) SetAnio(anio int) error {
	if anio < 1888 || anio > 2100 {
		return errors.New("el año del contenido no es válido")
	}

	c.anio = anio
	return nil
}

// Titulo devuelve el título del contenido.
func (c Contenido) Titulo() string {
	return c.titulo
}

// Genero devuelve el género del contenido.
func (c Contenido) Genero() string {
	return c.genero
}

// Anio devuelve el año de lanzamiento del contenido.
func (c Contenido) Anio() int {
	return c.anio
}

// Disponible indica si el contenido está disponible en el catálogo.
func (c Contenido) Disponible() bool {
	return c.disponible
}

// Activar vuelve a poner el contenido a disposición de los usuarios.
func (c *Contenido) Activar() {
	c.disponible = true
}

// Desactivar retira temporalmente el contenido del catálogo.
func (c *Contenido) Desactivar() {
	c.disponible = false
}
