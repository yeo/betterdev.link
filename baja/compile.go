package baja

import (
	"log"

	"html/template"
	"os"
)

//TODO Move theme reference to config with DI
func Compile(source string) error {
	log.Println("Done", source)

	t, err := template.ParseFiles("themes/yeo/layout.tmpl", "themes/yeo/index.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	if err := t.ExecuteTemplate(os.Stdout, "base", nil); err != nil {
		log.Fatal(err)
	}

	return nil
}
