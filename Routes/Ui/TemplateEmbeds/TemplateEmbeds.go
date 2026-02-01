package TemplateEmbeds

import (
	"embed"
)

//go:embed ComponentEmbeds/*.html
var ComponentEmbeds embed.FS

//go:embed PageEmbeds/*.html
var PageEmbeds embed.FS

//go:embed WrapperEmbeds/*.html
var WrapperEmbeds embed.FS
