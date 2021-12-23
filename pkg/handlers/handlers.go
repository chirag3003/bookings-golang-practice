package handlers

import (
	"fmt"
	"git.chirag.codes/chirag/bookings-golang/pkg/config"
	"git.chirag.codes/chirag/bookings-golang/pkg/models"
	"git.chirag.codes/chirag/bookings-golang/pkg/render"
	"net/http"
)

var Repo *Repository

// TemplateData holds data send from handlers to templates

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(res http.ResponseWriter, req *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "hello world"
	remoteIP := req.RemoteAddr
	m.App.Session.Put(req.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(res, "index", &models.TemplateData{
		StringMap: stringMap,
	})
}
func (m *Repository) About(res http.ResponseWriter, req *http.Request) {
	ip := m.App.Session.GetString(req.Context(), "remote_ip")
	fmt.Fprintf(res, `{"page":"%s"}`, ip)
}
