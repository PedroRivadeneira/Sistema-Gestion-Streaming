package modelos

import "errors"

// Usuario representa a una persona registrada en el sistema de streaming.
// Los campos son privados para aplicar encapsulación.
type Usuario struct {
	nombre   string
	email    string
	edad     int
	activo   bool
}

// NuevoUsuario crea un usuario con los datos iniciales proporcionados.
func NuevoUsuario(nombre, email string, edad int) (Usuario, error) {
	usuario := Usuario{}

	if err := usuario.SetNombre(nombre); err != nil {
		return Usuario{}, err
	}
	if err := usuario.SetEmail(email); err != nil {
		return Usuario{}, err
	}
	if err := usuario.SetEdad(edad); err != nil {
		return Usuario{}, err
	}

	usuario.activo = true
	return usuario, nil
}

// SetNombre actualiza el nombre después de validar que no esté vacío.
func (u *Usuario) SetNombre(nombre string) error {
	if nombre == "" {
		return errors.New("el nombre no puede estar vacío")
	}

	u.nombre = nombre
	return nil
}

// SetEmail actualiza el correo después de validar que no esté vacío.
func (u *Usuario) SetEmail(email string) error {
	if email == "" {
		return errors.New("el correo no puede estar vacío")
	}

	u.email = email
	return nil
}

// SetEdad actualiza la edad después de validar un rango básico.
func (u *Usuario) SetEdad(edad int) error {
	if edad < 1 || edad > 120 {
		return errors.New("la edad debe estar entre 1 y 120 años")
	}

	u.edad = edad
	return nil
}

// Nombre devuelve el nombre del usuario.
func (u Usuario) Nombre() string {
	return u.nombre
}

// Email devuelve el correo del usuario.
func (u Usuario) Email() string {
	return u.email
}

// Edad devuelve la edad del usuario.
func (u Usuario) Edad() int {
	return u.edad
}

// Activo indica si el usuario se encuentra activo.
func (u Usuario) Activo() bool {
	return u.activo
}

// Desactivar cambia el estado del usuario a inactivo.
func (u *Usuario) Desactivar() {
	u.activo = false
}
