# Plan y evidencia de pruebas

## Pruebas unitarias

Se utilizan archivos con sufijo `_test.go` y el paquete estándar `testing` para comprobar la lógica de modelos y servicios.

Casos principales:

- Creación y validación de usuarios.
- Registro y rechazo de usuarios duplicados.
- Validación y registro de planes.
- Registro y búsqueda de contenidos.
- Registro de suscripciones cuando existen usuario y plan.
- Rechazo de suscripciones con datos inexistentes.

## Pruebas de servicios HTTP

El archivo `src/api/server_test.go` utiliza `net/http/httptest` para comprobar la API sin depender de un servidor externo.

Casos implementados:

- `POST /usuarios` y consulta posterior del usuario.
- Registro duplicado y respuesta HTTP 409.
- `POST /contenidos` y búsqueda por género.
- `POST /suscripciones` validando usuario y plan existentes.
- `GET /estadisticas/concurrente` para comprobar la respuesta del servicio concurrente.

## Prueba de integración

La integración se verifica ejecutando:

```bash
go test ./...
```

Esta ejecución comprueba conjuntamente los paquetes del proyecto y los servicios HTTP de la API.

## Prueba de aceptación

Durante la demostración se ejecuta la aplicación con:

```bash
go run .
```

Y se verifican desde un cliente HTTP las funcionalidades principales: registrar recursos, consultar datos, buscar contenido, crear suscripciones y consultar estadísticas.

## Concurrencia

El endpoint `GET /estadisticas/concurrente` utiliza goroutines para calcular métricas independientes, un `sync.WaitGroup` para esperar a todas las tareas y un channel para recoger los resultados. Adicionalmente, el servidor incrementa con `sync/atomic` el contador de solicitudes procesadas.

La evidencia final de aceptación se presentará en el video mediante las solicitudes HTTP, sus respuestas JSON y la demostración del endpoint concurrente.
