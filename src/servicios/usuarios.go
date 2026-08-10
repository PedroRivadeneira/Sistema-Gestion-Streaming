package servicios

import (
	"errors"

	"github.com/PedroRivadeneira/Sistema-Gestion-Streaming/src/modelos"
)

// RegistroUsuarios administra usuarios mediante un map para facilitar búsquedas por correo.
type RegistroUsuarios struct {
	usuarios map[string]modelos.Usuario
}

// NuevoRegistroUsuarios crea un registro vacío de usuarios.
func NuevoRegistroUsuarios() RegistroUsuarios {
	return RegistroUsuarios{usuarios: make(map[string]modelos.Usuario)}
}

// Registrar agrega un usuario y evita duplicar su correo.
func (r *RegistroUsuarios) Registrar(usuario modelos.Usuario) error {
	correo := usuario.Email()
	if _, existe := r.usuarios[correo]; existe {
		return errors.New("ya existe un usuario con ese correo")
	}

	r.usuarios[correo] = usuario
	return nil
}

// Buscar obtiene un usuario mediante su correo.
func (r RegistroUsuarios) Buscar(correo string) (modelos.Usuario, error) {
	usuario, existe := r.usuarios[correo]
	if !existe {
		return modelos.Usuario{}, errors.New("usuario no encontrado")
	}
	return usuario, nil
}

// Lista devuelve todos los usuarios registrados.
func (r RegistroUsuarios) Lista() []modelos.Usuario {
	usuarios := make([]modelos.Usuario, 0, len(r.usuarios))
	for _, usuario := range r.usuarios {
		usuarios = append(usuarios, usuario)
	}
	return usuarios
}
