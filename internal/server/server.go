package server

import (
	"embed"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/thisisamr/SysWatch/app"
	"github.com/thisisamr/SysWatch/internal/metrics"
	"github.com/thisisamr/SysWatch/internal/ws"
)

//go:embed static/*
var static embed.FS

type Subscriber struct {
	Msgchan chan []byte
}
type service struct {
	Subscribers map[string]Subscriber
	Router      *chi.Mux
}

func homeHandler(provider metrics.StatProvider) http.HandlerFunc {

	f := func(w http.ResponseWriter, r *http.Request) {
		data, err := metrics.GatherAllMetrics(provider)
		if err != nil {
		}
		app.Page(*data).Render(r.Context(), w)
	}
	return f
}
func (s *service) Subscribe() {
}
func (s *service) AddSubscriber() {

}

func Abouthandler(w http.ResponseWriter, r *http.Request) {
	app.AboutPage().Render(r.Context(), w)
}
func Contacthandler(w http.ResponseWriter, r *http.Request) {
	app.ContactPage().Render(r.Context(), w)
}
func NewServer(provider metrics.StatProvider) *service {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	//css embeded into the binaryy
	r.Handle("/static/*", http.FileServer(http.FS(static)))
	r.Get("/", homeHandler(provider))
	r.Get("/about", Abouthandler)
	r.Get("/contact", Contacthandler)

	r.Get("/ws", ws.ClockWsHandler)
	r.Get("/metrics", ws.MetricsHandler(provider))

	return &service{
		Subscribers: make(map[string]Subscriber),
		Router:      r,
	}
}

// func gatherMetricsJSON(provider metrics.StatProvider) ([]byte, error) {
// 	metrics, err := metrics.GatherAllMetrics(provider)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return json.Marshal(metrics)
// }
