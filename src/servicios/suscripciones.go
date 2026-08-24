package servicios

import (
	"errors"
	"strings"

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

// Registrar crea una suscripción solamente si el usuario y el plan existen
// y no existe ya una suscripción activa para la misma combinación.
func (r *RegistroSuscripciones) Registrar(emailUsuario, nombrePlan string, usuarios RegistroUsuarios, planes RegistroPlanes) error {
	emailUsuario = strings.TrimSpace(strings.ToLower(emailUsuario))
	nombrePlan = strings.TrimSpace(nombrePlan)

	if _, err := usuarios.Buscar(emailUsuario); err != nil {
		return errors.New("no se puede crear la suscripción: usuario no encontrado")
	}
	if _, err := planes.Buscar(nombrePlan); err != nil {
		return errors.New("no se puede crear la suscripción: plan no encontrado")
	}

	for _, suscripcion := range r.suscripciones {
		if suscripcion.Activa() &&
			strings.EqualFold(suscripcion.EmailUsuario(), emailUsuario) &&
			strings.EqualFold(suscripcion.NombrePlan(), nombrePlan) {
			return errors.New("ya existe una suscripción activa para este usuario y plan")
		}
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
	return append([]modelos.Suscripcion(nil), r.suscripciones...)
}

// Cantidad devuelve la cantidad de suscripciones registradas.
func (r RegistroSuscripciones) Cantidad() int {
	return len(r.suscripciones)
}
