package templates

import (
	"io"
	"text/template"

	"github.com/ehutchllew/template.ts/cmd/models"
)

func GenerateRoot(w io.Writer) {
	tsTmpl := template.Must(template.New("tsconfig").Parse(tsConfigTemplate))
	tsTmpl.Execute(w, models.UserAnswers{Test: "Hello World!"})
}
