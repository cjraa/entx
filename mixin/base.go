package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type ID struct{ ent.Schema }

func (ID) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
	}
}

type Base struct{ ent.Schema }

func (Base) Fields() []ent.Field {
	return append(
		ID{}.Fields(),
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Immutable(),
	)
}

var _ ent.Mixin = (*ID)(nil)
var _ ent.Mixin = (*Base)(nil)
