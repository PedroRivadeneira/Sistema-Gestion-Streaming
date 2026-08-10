package modelos

import "errors"

// Suscripcion representa la relación entre un usuario y un plan de streaming.
type Suscripcion struct {
	emailUsuario string
	nombrePlan   string
	activa       bool
}

// NuevaSuscripcion crea una suscripción validando los datos necesarios.
func NuevaSuscripcion(emailUsuario, nombrePlan string) (Suscripcion, error) {
	if emailUsuario == "" {
		return Suscripcion{}, errors.New("el correo del usuario no puede estar vacío")
	}
	if nombrePlan == "" {
		return Suscripcion{}, errors.New("el nombre del plan no puede estar vacío")
	}

	return Suscripcion{
		emailUsuario: emailUsuario,
		nombrePlan:   nombrePlan,
		activa:       true,
	}, nil
}

// EmailUsuario devuelve el correo del usuario asociado.
func (s Suscripcion) EmailUsuario() string {
	return s.emailUsuario
}

// NombrePlan devuelve el nombre del plan contratado.
func (s Suscripcion) NombrePlan() string {
	return s.nombrePlan
}

// Activa indica si la suscripción está activa.
func (s Suscripcion) Activa() bool {
	return s.activa
}

// Cancelar desactiva la suscripción.
func (s *Suscripcion) Cancelar() {
	s.activa = false
}
