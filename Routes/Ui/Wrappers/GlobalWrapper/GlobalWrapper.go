package GlobalWrapper

import (
	"Polybub/Utilities"
	"bytes"
	"html/template"
	"path/filepath"
)

type globalWrapperVariant struct {
	Title string
	Body  template.HTML
}

func GetSafeHtml(filePath string, data any) (string, error) {
	body, err := template.
		New(filepath.Base(filePath)).
		ParseFiles(filePath)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	body.Execute(buf, data)

	return buf.String(), nil
}

func GetWrappedTemplate(safeHtml string) (string, error) {
	wrapPath := "Routes/Ui/Wrappers/GlobalWrapper/global.html"
	wrap, err := template.ParseFiles(wrapPath)
	if err != nil {
		return "", err
	}

	variant := globalWrapperVariant{
		Title: Utilities.GlobalConfig.Domain,
		Body:  template.HTML(safeHtml),
	}

	buf := new(bytes.Buffer)
	wrap.Execute(buf, variant)

	return buf.String(), nil
}
