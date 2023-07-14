package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/QiZD90/resume-generator/internal/entity"
	"github.com/QiZD90/resume-generator/internal/pdf"
	"github.com/go-chi/chi/v5"
)

func (app *Application) InternalServerError(w http.ResponseWriter, err error) {
	app.errorLog.Print(err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) Index(w http.ResponseWriter, r *http.Request) {
}

func (app *Application) Create(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1 * 1024 * 1024)

	file, _, err := r.FormFile("image")
	if err != nil {
		app.InternalServerError(w, err)
		return
	}
	defer file.Close()

	img, err := io.ReadAll(file)
	if err != nil {
		app.InternalServerError(w, err)
		return
	}

	file2, _, err := r.FormFile("info")
	if err != nil {
		app.InternalServerError(w, err)
		return
	}
	defer file2.Close()

	infoRaw, err := io.ReadAll(file2)
	if err != nil {
		app.InternalServerError(w, err)
		return
	}

	info := entity.ResumeInfo{}
	err = json.Unmarshal(infoRaw, &info)
	if err != nil {
		app.InternalServerError(w, err)
		return
	}
	info.Image = img

	buffer := bytes.Buffer{}
	err = app.generator.Generate(info, &buffer)
	if err != nil {
		app.InternalServerError(w, err)
		return
	}

	err = pdf.Generate(w, &buffer)
	if err != nil {
		app.InternalServerError(w, err)
		return
	}
}

func (app *Application) GetRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", app.Index)
	r.Post("/create", app.Create)

	return r
}
