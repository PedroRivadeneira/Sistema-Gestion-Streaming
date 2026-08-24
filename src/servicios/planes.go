package servicios

import (
	"errors"

	"github.com/PedroRivadeneira/Sistema-Gestion-Streaming/src/modelos"
)

// RegistroPlanes administra los planes disponibles mediante un Map.
type RegistroPlanes struct {
	planes map[string]modelos.Plan
}

// NuevoRegistroPlanes crea un registro vacío de planes.
func NuevoRegistroPlanes() RegistroPlanes {
	return RegistroPlanes{planes: make(map[string]modelos.Plan)}
}

// Registrar agrega un plan usando su nombre como clave única.
func (r *RegistroPlanes) Registrar(plan modelos.Plan) error {
	if _, existe := r.planes[plan.Nombre()]; existe {
		return errors.New("el plan ya está registrado")
	}
	r.planes[plan.Nombre()] = plan
	return nil
}

// Buscar obtiene un plan por su nombre.
func (r RegistroPlanes) Buscar(nombre string) (modelos.Plan, error) {
	plan, existe := r.planes[nombre]
	if !existe {
		return modelos.Plan{}, errors.New("el plan no existe")
	}
	return plan, nil
}

// Lista devuelve todos los planes registrados.
func (r RegistroPlanes) Lista() []modelos.Plan {
	planes := make([]modelos.Plan, 0, len(r.planes))
	for _, plan := range r.planes {
		planes = append(planes, plan)
	}
	return planes
}

// Cantidad devuelve la cantidad de planes registrados.
func (r RegistroPlanes) Cantidad() int {
	return len(r.planes)
}
