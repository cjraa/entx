package entx

import (
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

type ServiceExtension struct {
	entc.DefaultExtension
}

func (*ServiceExtension) Templates() []*gen.Template {
	return []*gen.Template{
		gen.MustParse(gen.NewTemplate("greet").ParseFiles("templates/greet.tmpl")),
	}
}
