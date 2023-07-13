package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/fs"
	"log"
	"math"
	"os"
	"path/filepath"
	"text/template"
)

var (
	templatePath = filepath.Join("assets", "template.tmpl")
	contactPath  = filepath.Join("assets", "contact.tmpl")
	imagePath    = filepath.Join("assets", "photo.jpg")

	Skills = []string{
		"Go", "C", "Java", "Python 3", "Rust", "Android",
		"PostgreSQL", "Linux", "Git", "OOP", "RESTful API",
		"Test-Driven Development", "SOLID", "CI/CD",
		"Алгоритмы и структуры данных",
	}
)

func FormatInColumns(items []string, nColumns int) [][]string {
	result := make([][]string, nColumns)
	itemsPerColumn := int(math.Ceil(float64(len(items)) / float64(nColumns)))

	for i, v := range items {
		column := i / itemsPerColumn
		result[column] = append(result[column], v)
	}

	return result
}

func ImageToBase64(src []byte) (string, error) {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(src)))

	_, format, err := image.DecodeConfig(bytes.NewReader(src))
	if err != nil {
		return "", err
	}

	base64.StdEncoding.Encode(dst, src)

	return fmt.Sprintf("data:image/%s;base64, %s", format, dst), nil
}

type Contact struct {
	Name, Link, LinkText, Text string
}

func main() {
	t, err := template.ParseFiles(templatePath, contactPath)
	if err != nil {
		log.Fatal(err)
	}

	r, err := os.ReadFile(imagePath)
	if err != nil {
		log.Fatal(err)
	}

	img, err := ImageToBase64(r)
	if err != nil {
		log.Fatal(err)
	}

	err = os.MkdirAll("output", fs.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}

	w, err := os.Create(filepath.Join("output", "resume.html"))
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(w, struct {
		Image        string
		SkillColumns [][]string
		Contacts     []Contact
	}{
		Image:        img,
		SkillColumns: FormatInColumns(Skills, 2),
		Contacts: []Contact{
			{Name: "Phone", Text: "+7(985)750-81-05"},
			{Name: "Telegram", Link: "https://t.me/qizd90"},
			{
				Name:     "Email",
				Link:     "mailto:puzko.e02@gmail.com",
				LinkText: "puzko.e02@gmail.com",
			},
			{Name: "GitHub", Link: "https://github.com/QiZD90"},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
