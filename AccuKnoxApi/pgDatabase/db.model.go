package pgdatabase

import (
	"database/sql"
)

type (
	Users struct {
		ID        uint32 `json:"userId"`
		Name      string `json:"name"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		SessionId string `json:"sessionId"`
	}
	UsersNull struct {
		Name      sql.NullString
		SessionId sql.NullString
		UpdatedBy sql.NullString
	}
	Notes struct {
		ID        uint32 `json:"id"`
		SessionId string `json:"sessionId"`
		Note      string `json:"note"`
	}
	NotesNull struct {
		Note sql.NullString
	}
)
