package modelos

import "errors"

// Plan representa un plan de suscripción disponible en el sistema.
type Plan struct {
	nombre       string
	precioMensual float64
	pantallas    int
}

// NuevoPlan crea un plan validando sus datos principales.
func NuevoPlan(nombre string, precioMensual float64, pantallas int) (Plan, error) {
	plan := Plan{}
	if err := plan.SetNombre(nombre); err != nil {
		return Plan{}, err
	}
	if err := plan.SetPrecioMensual(precioMensual); err != nil {
		return Plan{}, err
	}
	if err := plan.SetPantallas(pantallas); err != nil {
		return Plan{}, err
	}
	return plan, nil
}

func (p *Plan) SetNombre(nombre string) error {
	if nombre == "" {
		return errors.New("el nombre del plan no puede estar vacío")
	}
	p.nombre = nombre
	return nil
}

func (p *Plan) SetPrecioMensual(precio float64) error {
	if precio < 0 {
		return errors.New("el precio mensual no puede ser negativo")
	}
	p.precioMensual = precio
	return nil
}

func (p *Plan) SetPantallas(pantallas int) error {
	if pantallas < 1 {
		return errors.New("el plan debe tener al menos una pantalla")
	}
	p.pantallas = pantallas
	return nil
}

func (p Plan) Nombre() string { return p.nombre }
func (p Plan) PrecioMensual() float64 { return p.precioMensual }
func (p Plan) Pantallas() int { return p.pantallas }
