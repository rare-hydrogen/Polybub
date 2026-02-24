package WrapperTemplates

import (
	"Polybub/Utilities"
	"embed"
	"os"
)

//go:embed *.html
var Embeds embed.FS

func GetTemplate(filename string) ([]byte, error) {
	if Utilities.GlobalConfig.Env == "production" {
		return Embeds.ReadFile(filename)
	} else {
		return os.ReadFile("./Routes/Ui/Wrappers/WrapperTemplates/" + filename)
	}
}
