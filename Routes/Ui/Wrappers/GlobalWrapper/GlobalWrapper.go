package GlobalWrapper

import (
	"Polybub/Routes/Ui/Wrappers/WrapperTemplates"
	"Polybub/Utilities"
	"bytes"
	"html/template"
)

type globalWrapperVariant struct {
	Title string
	Body  template.HTML
}

func GetSafeHtml(b []byte, data any) (string, error) {
	temp := template.New("")
	temp.Funcs(FuncMap)
	body, err := temp.Parse(string(template.HTML(b)))
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	body.Execute(buf, data)

	return buf.String(), nil
}

func GetWrappedTemplate(safeHtml string) (string, error) {
	temp := template.New("")
	wrappedBytes, _ := WrapperTemplates.GetTemplate("private.html")
	wrap, err := temp.Parse(string(template.HTML(wrappedBytes)))
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

func GetPublicTemplate(safeHtml string) (string, error) {
	temp := template.New("")
	wrappedBytes, _ := WrapperTemplates.GetTemplate("global.html")
	wrap, err := temp.Parse(string(template.HTML(wrappedBytes)))
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
