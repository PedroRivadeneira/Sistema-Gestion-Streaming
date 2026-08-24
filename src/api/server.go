package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/PedroRivadeneira/Sistema-Gestion-Streaming/src/interfaces"
	"github.com/PedroRivadeneira/Sistema-Gestion-Streaming/src/modelos"
	"github.com/PedroRivadeneira/Sistema-Gestion-Streaming/src/servicios"
)

// Server expone el sistema de streaming como una API REST.
// El mutex protege los registros, mientras que totalRequests usa atomic
// porque puede incrementarse desde varias solicitudes concurrentes.
type Server struct {
	usuarios       *servicios.RegistroUsuarios
	planes         *servicios.RegistroPlanes
	catalogo       *servicios.Catalogo
	suscripciones  *servicios.RegistroSuscripciones
	mu             sync.RWMutex
	totalRequests  int64
}

// NewServer crea un servidor HTTP usando los servicios existentes del proyecto.
func NewServer(
	usuarios *servicios.RegistroUsuarios,
	planes *servicios.RegistroPlanes,
	catalogo *servicios.Catalogo,
	suscripciones *servicios.RegistroSuscripciones,
) *Server {
	return &Server{
		usuarios:      usuarios,
		planes:        planes,
		catalogo:      catalogo,
		suscripciones: suscripciones,
	}
}

// ServeHTTP enruta las solicitudes y cuenta de forma atómica las peticiones procesadas.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	atomic.AddInt64(&s.totalRequests, 1)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	switch {
	case r.URL.Path == "/" && r.Method == http.MethodGet:
		s.handleInicio(w, r)
	case r.URL.Path == "/health" && r.Method == http.MethodGet:
		writeJSON(w, http.StatusOK, map[string]any{"estado": "activo"})
	case r.URL.Path == "/usuarios" && r.Method == http.MethodPost:
		s.handleCrearUsuario(w, r)
	case r.URL.Path == "/usuarios" && r.Method == http.MethodGet:
		s.handleListarUsuarios(w, r)
	case strings.HasPrefix(r.URL.Path, "/usuarios/") && r.Method == http.MethodGet:
		s.handleBuscarUsuario(w, r)
	case r.URL.Path == "/planes" && r.Method == http.MethodPost:
		s.handleCrearPlan(w, r)
	case r.URL.Path == "/planes" && r.Method == http.MethodGet:
		s.handleListarPlanes(w, r)
	case strings.HasPrefix(r.URL.Path, "/planes/") && r.Method == http.MethodGet:
		s.handleBuscarPlan(w, r)
	case r.URL.Path == "/contenidos" && r.Method == http.MethodPost:
		s.handleCrearContenido(w, r)
	case r.URL.Path == "/contenidos" && r.Method == http.MethodGet:
		s.handleListarContenidos(w, r)
	case r.URL.Path == "/contenidos/buscar" && r.Method == http.MethodGet:
		s.handleBuscarContenidos(w, r)
	case r.URL.Path == "/suscripciones" && r.Method == http.MethodPost:
		s.handleCrearSuscripcion(w, r)
	case r.URL.Path == "/suscripciones" && r.Method == http.MethodGet:
		s.handleListarSuscripciones(w, r)
	case r.URL.Path == "/estadisticas" && r.Method == http.MethodGet:
		s.handleEstadisticas(w, r)
	case r.URL.Path == "/estadisticas/concurrente" && r.Method == http.MethodGet:
		s.handleEstadisticasConcurrentes(w, r)
	default:
		writeError(w, http.StatusNotFound, "ruta o método no disponible")
	}
}

type usuarioRequest struct {
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
	Edad   int    `json:"edad"`
}

type planRequest struct {
	Nombre        string  `json:"nombre"`
	PrecioMensual float64 `json:"precio_mensual"`
	Pantallas     int     `json:"pantallas"`
}

type contenidoRequest struct {
	Tipo             string `json:"tipo"`
	Titulo           string `json:"titulo"`
	Genero           string `json:"genero"`
	Anio             int    `json:"anio"`
	DuracionMinutos int    `json:"duracion_minutos,omitempty"`
	Temporadas       int    `json:"temporadas,omitempty"`
}

type suscripcionRequest struct {
	EmailUsuario string `json:"email_usuario"`
	NombrePlan   string `json:"nombre_plan"`
}

func (s *Server) handleInicio(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{
		"sistema": "Sistema de Gestión de Streaming",
		"version": "API REST - Unidad 4",
		"servicios": []string{
			"POST /usuarios",
			"GET /usuarios",
			"GET /usuarios/{email}",
			"POST /planes",
			"GET /planes",
			"GET /planes/{nombre}",
			"POST /contenidos",
			"GET /contenidos",
			"GET /contenidos/buscar",
			"POST /suscripciones",
			"GET /suscripciones",
			"GET /estadisticas",
			"GET /estadisticas/concurrente",
		},
	})
}

func (s *Server) handleCrearUsuario(w http.ResponseWriter, r *http.Request) {
	var req usuarioRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	usuario, err := modelos.NuevoUsuario(req.Nombre, req.Email, req.Edad)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if err := s.usuarios.Registrar(usuario); err != nil {
		writeError(w, http.StatusConflict, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, usuarioView(usuario))
}

func (s *Server) handleListarUsuarios(w http.ResponseWriter, _ *http.Request) {
	s.mu.RLock()
	usuarios := s.usuarios.Lista()
	s.mu.RUnlock()

	out := make([]map[string]any, 0, len(usuarios))
	for _, usuario := range usuarios {
		out = append(out, usuarioView(usuario))
	}
	writeJSON(w, http.StatusOK, out)
}

func (s *Server) handleBuscarUsuario(w http.ResponseWriter, r *http.Request) {
	email, err := url.PathUnescape(strings.TrimPrefix(r.URL.Path, "/usuarios/"))
	if err != nil || email == "" {
		writeError(w, http.StatusBadRequest, "correo inválido")
		return
	}
	s.mu.RLock()
	usuario, err := s.usuarios.Buscar(email)
	s.mu.RUnlock()
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, usuarioView(usuario))
}

func (s *Server) handleCrearPlan(w http.ResponseWriter, r *http.Request) {
	var req planRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	plan, err := modelos.NuevoPlan(req.Nombre, req.PrecioMensual, req.Pantallas)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if err := s.planes.Registrar(plan); err != nil {
		writeError(w, http.StatusConflict, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, planView(plan))
}

func (s *Server) handleListarPlanes(w http.ResponseWriter, _ *http.Request) {
	s.mu.RLock()
	planes := s.planes.Lista()
	s.mu.RUnlock()
	out := make([]map[string]any, 0, len(planes))
	for _, plan := range planes {
		out = append(out, planView(plan))
	}
	writeJSON(w, http.StatusOK, out)
}

func (s *Server) handleBuscarPlan(w http.ResponseWriter, r *http.Request) {
	nombre, err := url.PathUnescape(strings.TrimPrefix(r.URL.Path, "/planes/"))
	if err != nil || nombre == "" {
		writeError(w, http.StatusBadRequest, "nombre de plan inválido")
		return
	}
	s.mu.RLock()
	plan, err := s.planes.Buscar(nombre)
	s.mu.RUnlock()
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, planView(plan))
}

func (s *Server) handleCrearContenido(w http.ResponseWriter, r *http.Request) {
	var req contenidoRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	var contenido interfaces.ContenidoGestionable
	var err error
	switch strings.ToLower(strings.TrimSpace(req.Tipo)) {
	case "pelicula", "película":
		pelicula, createErr := modelos.NuevaPelicula(req.Titulo, req.Genero, req.Anio, req.DuracionMinutos)
		err = createErr
		if err == nil {
			contenido = pelicula
		}
	case "serie":
		serie, createErr := modelos.NuevaSerie(req.Titulo, req.Genero, req.Anio, req.Temporadas)
		err = createErr
		if err == nil {
			contenido = serie
		}
	default:
		err = errors.New("tipo de contenido no válido; use pelicula o serie")
	}
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	if err := s.catalogo.Agregar(contenido); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, contenidoView(contenido))
}

func (s *Server) handleListarContenidos(w http.ResponseWriter, _ *http.Request) {
	s.mu.RLock()
	contenidos := append([]interfaces.ContenidoGestionable(nil), s.catalogo.Listar()...)
	s.mu.RUnlock()
	out := make([]map[string]any, 0, len(contenidos))
	for _, contenido := range contenidos {
		out = append(out, contenidoView(contenido))
	}
	writeJSON(w, http.StatusOK, out)
}

func (s *Server) handleBuscarContenidos(w http.ResponseWriter, r *http.Request) {
	q := strings.TrimSpace(strings.ToLower(r.URL.Query().Get("q")))
	genero := strings.TrimSpace(strings.ToLower(r.URL.Query().Get("genero")))
	tipo := strings.TrimSpace(strings.ToLower(r.URL.Query().Get("tipo")))
	anioTexto := strings.TrimSpace(r.URL.Query().Get("anio"))
	anio := 0
	if anioTexto != "" {
		parsed, err := strconv.Atoi(anioTexto)
		if err != nil {
			writeError(w, http.StatusBadRequest, "anio debe ser numérico")
			return
		}
		anio = parsed
	}
	if q == "" && genero == "" && tipo == "" && anio == 0 {
		writeError(w, http.StatusBadRequest, "use al menos un filtro: q, genero, tipo o anio")
		return
	}

	s.mu.RLock()
	contenidos := append([]interfaces.ContenidoGestionable(nil), s.catalogo.Listar()...)
	s.mu.RUnlock()

	resultados := make([]map[string]any, 0)
	for _, contenido := range contenidos {
		matches := true
		if q != "" && !strings.Contains(strings.ToLower(contenido.Nombre()), q) {
			matches = false
		}
		if genero != "" && strings.ToLower(contenido.Genero()) != genero {
			matches = false
		}
		if tipo != "" && strings.ToLower(contenido.Tipo()) != normalizarTipo(tipo) {
			matches = false
		}
		if anio != 0 && contenido.Anio() != anio {
			matches = false
		}
		if matches {
			resultados = append(resultados, contenidoView(contenido))
		}
	}

	if len(resultados) == 0 {
		writeError(w, http.StatusNotFound, "no se encontraron contenidos con los filtros indicados")
		return
	}
	writeJSON(w, http.StatusOK, resultados)
}

func (s *Server) handleCrearSuscripcion(w http.ResponseWriter, r *http.Request) {
	var req suscripcionRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if err := s.suscripciones.Registrar(req.EmailUsuario, req.NombrePlan, *s.usuarios, *s.planes); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	lista := s.suscripciones.Lista()
	writeJSON(w, http.StatusCreated, suscripcionView(lista[len(lista)-1]))
}

func (s *Server) handleListarSuscripciones(w http.ResponseWriter, r *http.Request) {
	s.mu.RLock()
	lista := append([]modelos.Suscripcion(nil), s.suscripciones.Lista()...)
	s.mu.RUnlock()
	email := strings.TrimSpace(strings.ToLower(r.URL.Query().Get("email")))
	out := make([]map[string]any, 0, len(lista))
	for _, suscripcion := range lista {
		if email != "" && strings.ToLower(suscripcion.EmailUsuario()) != email {
			continue
		}
		out = append(out, suscripcionView(suscripcion))
	}
	writeJSON(w, http.StatusOK, out)
}

func (s *Server) handleEstadisticas(w http.ResponseWriter, _ *http.Request) {
	s.mu.RLock()
	usuarios := len(s.usuarios.Lista())
	contenidos := s.catalogo.Listar()
	suscripciones := s.suscripciones.Cantidad()
	s.mu.RUnlock()

	peliculas, series := contarTipos(contenidos)
	writeJSON(w, http.StatusOK, map[string]any{
		"usuarios":             usuarios,
		"contenidos":           len(contenidos),
		"peliculas":            peliculas,
		"series":               series,
		"suscripciones":        suscripciones,
		"solicitudes_procesadas": atomic.LoadInt64(&s.totalRequests),
	})
}

// handleEstadisticasConcurrentes usa goroutines, canales y WaitGroup para
// calcular métricas en paralelo y luego sincronizar los resultados.
func (s *Server) handleEstadisticasConcurrentes(w http.ResponseWriter, _ *http.Request) {
	s.mu.RLock()
	usuarios := append([]modelos.Usuario(nil), s.usuarios.Lista()...)
	contenidos := append([]interfaces.ContenidoGestionable(nil), s.catalogo.Listar()...)
	suscripciones := append([]modelos.Suscripcion(nil), s.suscripciones.Lista()...)
	s.mu.RUnlock()

	type metric struct {
		key   string
		value int
	}
	resultados := make(chan metric, 3)
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		resultados <- metric{key: "usuarios", value: len(usuarios)}
	}()
	go func() {
		defer wg.Done()
		resultados <- metric{key: "contenidos", value: len(contenidos)}
	}()
	go func() {
		defer wg.Done()
		resultados <- metric{key: "suscripciones", value: len(suscripciones)}
	}()

	go func() {
		wg.Wait()
		close(resultados)
	}()

	respuesta := map[string]any{
		"modelo": "concurrencia con goroutines + channel + WaitGroup",
		"solicitudes_procesadas": atomic.LoadInt64(&s.totalRequests),
	}
	for resultado := range resultados {
		respuesta[resultado.key] = resultado.value
	}
	writeJSON(w, http.StatusOK, respuesta)
}

func usuarioView(usuario modelos.Usuario) map[string]any {
	return map[string]any{
		"nombre":  usuario.Nombre(),
		"email":   usuario.Email(),
		"edad":    usuario.Edad(),
		"activo":  usuario.Activo(),
	}
}

func planView(plan modelos.Plan) map[string]any {
	return map[string]any{
		"nombre":         plan.Nombre(),
		"precio_mensual": plan.PrecioMensual(),
		"pantallas":      plan.Pantallas(),
	}
}

func contenidoView(contenido interfaces.ContenidoGestionable) map[string]any {
	view := map[string]any{
		"tipo":        contenido.Tipo(),
		"titulo":      contenido.Titulo(),
		"genero":      contenido.Genero(),
		"anio":        contenido.Anio(),
		"disponible":  contenido.Disponible(),
	}
	if pelicula, ok := contenido.(modelos.Pelicula); ok {
		view["duracion_minutos"] = pelicula.Duracion()
	}
	if serie, ok := contenido.(modelos.Serie); ok {
		view["temporadas"] = serie.Temporadas()
	}
	return view
}

func suscripcionView(suscripcion modelos.Suscripcion) map[string]any {
	return map[string]any{
		"email_usuario": suscripcion.EmailUsuario(),
		"nombre_plan":   suscripcion.NombrePlan(),
		"activa":        suscripcion.Activa(),
	}
}

func contarTipos(contenidos []interfaces.ContenidoGestionable) (int, int) {
	peliculas, series := 0, 0
	for _, contenido := range contenidos {
		switch contenido.Tipo() {
		case "Película":
			peliculas++
		case "Serie":
			series++
		}
	}
	return peliculas, series
}

func normalizarTipo(tipo string) string {
	if tipo == "pelicula" || tipo == "película" {
		return "película"
	}
	return "serie"
}

func decodeJSON(r *http.Request, target any) error {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(target)
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]any{
		"error": message,
	})
}
