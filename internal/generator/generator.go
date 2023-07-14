package generator

import (
	_ "image/jpeg"
	_ "image/png"
	"io"
	"text/template"

	"github.com/QiZD90/resume-generator/internal/entity"
)

type Generator struct {
	PageTemplate *template.Template
}

func (g *Generator) Generate(info entity.ResumeInfo, w io.Writer) error {
	croppedImage, err := CropAndEncodeImage(info.Image, info.Crop)
	if err != nil {
		return err
	}

	err = g.PageTemplate.Execute(w, struct {
		entity.ResumeInfo
		CroppedImage  string
		SkillsColumns [][]string
	}{
		ResumeInfo:    info,
		CroppedImage:  croppedImage,
		SkillsColumns: DivideInColumns(info.Skills, info.Columns),
	})

	return err
}
