package servicios

import (
	"errors"

	"github.com/PedroRivadeneira/Sistema-Gestion-Streaming/src/modelos"
)

// RegistroSuscripciones administra las suscripciones activas y registradas.
type RegistroSuscripciones struct {
	suscripciones []modelos.Suscripcion
}

// NuevaRegistroSuscripciones crea un registro vacío de suscripciones.
func NuevaRegistroSuscripciones() RegistroSuscripciones {
	return RegistroSuscripciones{
		suscripciones: make([]modelos.Suscripcion, 0),
	}
}

// Registrar crea una suscripción solamente si el usuario y el plan existen.
func (r *RegistroSuscripciones) Registrar(emailUsuario, nombrePlan string, usuarios RegistroUsuarios, planes RegistroPlanes) error {
	if _, err := usuarios.Buscar(emailUsuario); err != nil {
		return errors.New("no se puede crear la suscripción: usuario no encontrado")
	}
	if _, err := planes.Buscar(nombrePlan); err != nil {
		return errors.New("no se puede crear la suscripción: plan no encontrado")
	}

	suscripcion, err := modelos.NuevaSuscripcion(emailUsuario, nombrePlan)
	if err != nil {
		return err
	}

	r.suscripciones = append(r.suscripciones, suscripcion)
	return nil
}

// Lista devuelve las suscripciones registradas.
func (r RegistroSuscripciones) Lista() []modelos.Suscripcion {
	return r.suscripciones
}

// Cantidad devuelve la cantidad de suscripciones registradas.
func (r RegistroSuscripciones) Cantidad() int {
	return len(r.suscripciones)
}
