package main

import (
	"embed"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/QiZD90/resume-generator/internal/generator"
)

//go:embed assets/template.gohtml
var fs embed.FS

type Application struct {
	errorLog  *log.Logger
	infoLog   *log.Logger
	generator *generator.Generator
}

func main() {
	errorLog := log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime)

	tmpl, err := template.ParseFS(fs, "assets/template.gohtml")
	if err != nil {
		errorLog.Fatal(err)
	}

	generator := generator.Generator{
		PageTemplate: tmpl,
	}

	app := &Application{
		errorLog:  errorLog,
		infoLog:   infoLog,
		generator: &generator,
	}

	infoLog.Printf("Starting server on :80")
	if err := http.ListenAndServe(":80", app.GetRouter()); err != nil {
		errorLog.Fatal(err)
	}
}
