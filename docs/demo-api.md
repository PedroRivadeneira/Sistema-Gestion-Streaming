# Guion técnico para demostrar la API

## 1. Iniciar el servidor

```bash
go run .
```

Debe aparecer:

```text
API REST iniciada en http://localhost:8080
```

## 2. Verificar estado

En el navegador:

```text
http://localhost:8080/health
```

Respuesta esperada:

```json
{"estado":"activo"}
```

## 3. Registrar usuario

PowerShell:

```powershell
Invoke-RestMethod -Method Post -Uri http://localhost:8080/usuarios -ContentType "application/json" -Body '{"nombre":"Ana","email":"ana@streaming.com","edad":22}'
```

## 4. Consultar usuarios

```text
GET http://localhost:8080/usuarios
```

## 5. Registrar contenido

```powershell
Invoke-RestMethod -Method Post -Uri http://localhost:8080/contenidos -ContentType "application/json" -Body '{"tipo":"serie","titulo":"Dark","genero":"Ciencia ficción","anio":2017,"temporadas":3}'
```

## 6. Buscar contenido por filtros

```text
GET http://localhost:8080/contenidos/buscar?genero=Ciencia%20ficci%C3%B3n
```

También se puede combinar con `tipo`, `anio` o `q`.

## 7. Crear suscripción

```powershell
Invoke-RestMethod -Method Post -Uri http://localhost:8080/suscripciones -ContentType "application/json" -Body '{"email_usuario":"ana@streaming.com","nombre_plan":"Premium"}'
```

## 8. Consultar estadísticas

```text
GET http://localhost:8080/estadisticas
```

## 9. Demostrar concurrencia

```text
GET http://localhost:8080/estadisticas/concurrente
```

La respuesta identifica el uso de **goroutines + channel + WaitGroup** y muestra el contador de solicitudes procesadas mediante `sync/atomic`.
