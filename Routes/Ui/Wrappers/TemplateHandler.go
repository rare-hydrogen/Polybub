package Wrappers

import (
	"Polybub/Routes/Jsend"
	"Polybub/Routes/Ui/Components/ComponentTemplates"
	"Polybub/Routes/Ui/Pages/PageTemplates"
	"Polybub/Routes/Ui/Wrappers/GlobalWrapper"
	"Polybub/Routes/Ui/Wrappers/WrapperTemplates"
	"Polybub/Utilities"
	"bytes"
	"html/template"
	"io/fs"
	"net/http"
	"os"
)

// Helper to make sure only GET methods are used on pages
func confirmVerb(w http.ResponseWriter, req *http.Request) {
	// Confirm the only allowed HTTP Verb is GET
	if req.Method != "GET" {
		Jsend.MethodNotAllowed(w, req)
		return
	}
}

// New function for calling pages
func getTemplate(pageName string, data any) (string, error) {
	// Prepare all three template locations for use, whether they are embedded or hot-reloaded
	var compDir fs.FS
	var pageDir fs.FS
	var wrapperDir fs.FS
	if Utilities.GlobalConfig.Env == "production" {
		compDir = ComponentTemplates.Embeds
		pageDir = PageTemplates.Embeds
		wrapperDir = WrapperTemplates.Embeds
	} else {
		compDir = os.DirFS("./Routes/Ui/Components/ComponentTemplates/")
		pageDir = os.DirFS("./Routes/Ui/Pages/PageTemplates/")
		wrapperDir = os.DirFS("./Routes/Ui/Wrappers/WrapperTemplates/")
	}

	// Combine all templates into one giant template, so any can be called from anywhere
	templates := template.New("root")
	templates.Funcs(GlobalWrapper.FuncMap)
	templates = template.Must(templates.ParseFS(compDir, "*.html"))
	templates = template.Must(templates.ParseFS(pageDir, "*.html"))
	templates = template.Must(templates.ParseFS(wrapperDir, "*.html"))

	// Apply data to chosen page
	buf := new(bytes.Buffer)
	err := templates.ExecuteTemplate(buf, pageName, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

// Helper to streamline page instantiation
func TemplateHandler(w http.ResponseWriter, req *http.Request, pageName string, data any) {
	confirmVerb(w, req)
	wrappedBody, err := getTemplate(pageName, data)
	if err != nil {
		Jsend.InternalServerError(w, req, err)
		return
	}
	Jsend.Ui(w, req, wrappedBody)
}
