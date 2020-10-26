package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
	"time"
	"github.com/gobuffalo/validate/v3/validators"
)
// Pizza is used by pop to map your pizzas database table to your go code.
type Pizza struct {
    ID uuid.UUID `json:"id" db:"id"`
    Name string `json:"name" db:"name"`
    Size string `json:"size" db:"size"`
    Ingredient1 string `json:"ingredient1" db:"ingredient1"`
    Ingredient2 string `json:"ingredient2" db:"ingredient2"`
    Ingredient3 string `json:"ingredient3" db:"ingredient3"`
    Ingredient4 string `json:"ingredient4" db:"ingredient4"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (p Pizza) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Pizzas is not required by pop and may be deleted
type Pizzas []Pizza

// String is not required by pop and may be deleted
func (p Pizzas) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *Pizza) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: p.Name, Name: "Name"},
		&validators.StringIsPresent{Field: p.Size, Name: "Size"},
		&validators.StringIsPresent{Field: p.Ingredient1, Name: "Ingredient1"},
		&validators.StringIsPresent{Field: p.Ingredient2, Name: "Ingredient2"},
		&validators.StringIsPresent{Field: p.Ingredient3, Name: "Ingredient3"},
		&validators.StringIsPresent{Field: p.Ingredient4, Name: "Ingredient4"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *Pizza) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *Pizza) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
