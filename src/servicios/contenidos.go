package servicios

import "github.com/PedroRivadeneira/Sistema-Gestion-Streaming/src/interfaces"

// MostrarContenido demuestra el polimorfismo mediante la interfaz ContenidoGestionable.
// Puede recibir una película, una serie u otro tipo que cumpla el contrato de la interfaz.
func MostrarContenido(contenido interfaces.ContenidoGestionable) string {
	return contenido.Tipo() + ": " + contenido.Titulo()
}
