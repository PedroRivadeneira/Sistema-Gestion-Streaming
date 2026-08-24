# Sistema de Gestión de Streaming

Proyecto individual de la asignatura **Programación Orientada a Objetos 2** de Ingeniería en Ciberseguridad (UIDE).

**Autor:** Pedro Rivadeneira  
**Fecha:** 23 de agosto de 2026  
**Lenguaje:** Go  
**Objetivo:** desarrollar un sistema de gestión de streaming que administre usuarios, contenido multimedia, planes y suscripciones, integrando los conocimientos de las cuatro unidades de la asignatura.

## Principales funcionalidades

- Registro y consulta de usuarios con validaciones y encapsulación.
- Gestión de planes de suscripción.
- Gestión polimórfica de películas y series mediante interfaces.
- Búsqueda de contenidos por título, género, tipo y año.
- Registro y consulta de suscripciones.
- Estadísticas generales del sistema.
- Estadísticas concurrentes usando goroutines, channels, `sync.WaitGroup` y `sync/atomic`.
- API REST con JSON y métodos HTTP GET y POST.
- Manejo de errores y respuestas HTTP con códigos adecuados.
- Pruebas unitarias y pruebas de servicios HTTP con `httptest`.

## Servicios web

La API se ejecuta con `net/http` en `http://localhost:8080`.

- `GET /` – información de la API y rutas disponibles.
- `GET /health` – estado del servicio.
- `POST /usuarios` – registrar usuario.
- `GET /usuarios` – listar usuarios.
- `GET /usuarios/{email}` – consultar usuario.
- `POST /planes` – registrar plan.
- `GET /planes` – listar planes.
- `GET /planes/{nombre}` – consultar plan.
- `POST /contenidos` – registrar película o serie.
- `GET /contenidos` – consultar catálogo.
- `GET /contenidos/buscar` – búsqueda mediante filtros.
- `POST /suscripciones` – registrar suscripción.
- `GET /suscripciones` – consultar suscripciones.
- `GET /estadisticas` – estadísticas generales.
- `GET /estadisticas/concurrente` – estadísticas calculadas con concurrencia.

## Ejecución

```bash
go run .
```

Después de iniciar el programa, abrir:

```text
http://localhost:8080
```

## Pruebas

```bash
go test ./...
```

El proyecto incluye pruebas sobre modelos y servicios, además de pruebas de endpoints HTTP mediante `net/http/httptest`.

## Integración de las unidades

- **Unidad 1:** funciones, estructuras de datos, Maps y Slices.
- **Unidad 2:** structs, métodos, encapsulación relacionada con los modelos e interfaces.
- **Unidad 3:** setters, validaciones, manejo de errores, interfaces y polimorfismo.
- **Unidad 4:** goroutines, channels, `sync.WaitGroup`, `sync/atomic`, servicios HTTP, REST, JSON y testing.

## Documentación complementaria

- [Diagrama de clases](docs/diagrama-clases.md)
- [Plan de pruebas](docs/pruebas.md)
- [Visión futura](docs/vision-futura.md)
