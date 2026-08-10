package modelos

import "errors"

// Serie representa un contenido de tipo serie.
type Serie struct {
	Contenido
	temporadas int
}

// NuevaSerie crea una serie validando los datos principales.
func NuevaSerie(titulo, genero string, anio, temporadas int) (Serie, error) {
	contenido, err := NuevoContenido(titulo, genero, anio)
	if err != nil {
		return Serie{}, err
	}

	serie := Serie{Contenido: contenido}
	if err := serie.SetTemporadas(temporadas); err != nil {
		return Serie{}, err
	}

	return serie, nil
}

// SetTemporadas actualiza el número de temporadas de la serie.
func (s *Serie) SetTemporadas(temporadas int) error {
	if temporadas <= 0 {
		return errors.New("una serie debe tener al menos una temporada")
	}
	 s.temporadas = temporadas
	return nil
}

// Temporadas devuelve el número de temporadas de la serie.
func (s Serie) Temporadas() int {
	return s.temporadas
}

// Tipo devuelve el tipo de contenido.
func (s Serie) Tipo() string {
	return "Serie"
}
