package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

// CompletionType represents the different types of completions
type CompletionType string

const (
	CompletionTypeTVShow    CompletionType = "TV Show"
	CompletionTypeVideoGame CompletionType = "Video Game"
	CompletionTypeBook      CompletionType = "Book"
	CompletionTypeAudioBook CompletionType = "Audio Book"
	CompletionTypeEvent     CompletionType = "Event"
)

// GetCompletionTypes returns all available completion types
func GetCompletionTypes() []CompletionType {
	return []CompletionType{
		CompletionTypeTVShow,
		CompletionTypeVideoGame,
		CompletionTypeBook,
		CompletionTypeAudioBook,
		CompletionTypeEvent,
	}
}

// Completion is used by pop to map your completions database table to your go code.
type Completion struct {
	ID          uuid.UUID      `json:"id" db:"id"`
	Name        string         `json:"name" db:"name"`
	Type        CompletionType `json:"type" db:"type"`
	Completions int            `json:"completions" db:"completions"`
	CompletedAt time.Time      `json:"completed_at" db:"completed_at"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (c Completion) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Completions is not required by pop and may be deleted
type Completions []Completion

// String is not required by pop and may be deleted
func (c Completions) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *Completion) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: c.Name, Name: "Name"},
		&validators.StringIsPresent{Field: string(c.Type), Name: "Type"},
		&validators.StringInclusion{Field: string(c.Type), Name: "Type", List: []string{
			string(CompletionTypeTVShow),
			string(CompletionTypeVideoGame),
			string(CompletionTypeBook),
			string(CompletionTypeAudioBook),
			string(CompletionTypeEvent),
		}},
		&validators.IntIsPresent{Field: c.Completions, Name: "Completions"},
		&validators.TimeIsPresent{Field: c.CompletedAt, Name: "CompletedAt"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *Completion) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *Completion) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
