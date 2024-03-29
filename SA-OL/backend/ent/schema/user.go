package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique(),
		field.String("email").NotEmpty().Unique(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("foodmenus", Foodmenu.Type).StorageKey(edge.Column("owner_id")),
		edge.To("mealplans", Mealplan.Type).StorageKey(edge.Column("owner_id")),
		edge.To("eatinghistorys", Eatinghistory.Type).StorageKey(edge.Column("user_id")),
	}
}
