package pgdatabase

import (
	"database/sql"
	"log"
)

func CreateNotesToDatabase(session_id, note string) (bool, uint32, error) {
	var id uint32
	sqlQuery := `INSERT INTO notes(session_id, note) VALUES($1, $2) RETURNING id;`
	if err := DB.QueryRow(sqlQuery, session_id, note).Scan(&id); err != nil {
		log.Println("CreateNotesToDatabase QueryRow " + err.Error())
		return false, id, err
	}
	return true, id, nil
}
func GetNotesBySessionIdFromDatabase(session_id string) ([]Notes, error) {
	var notes Notes
	var notesNull NotesNull
	var notes_array = []Notes{}
	sqlQuery := `SELECT id, session_id, note FROM notes WHERE session_id = $1`
	row, errQuery := DB.Query(sqlQuery, session_id)
	if err != nil {
		log.Println("GetNotesBySessionIdFromDatabase Query ", errQuery.Error())
		return notes_array, errQuery
	}
	for row.Next() {
		if errScan := row.Scan(&notes.ID, &notes.SessionId, &notesNull.Note); errScan != nil {
			log.Println("GetNotesBySessionIdFromDatabase Scan ", errScan.Error())
			return notes_array, errScan
		}

		notes.Note = notesNull.Note.String
		notes_array = append(notes_array, notes)
	}
	defer row.Close()
	return notes_array, nil
}
func DeleteNoteToDatabase(sid string, id uint32) (bool, error) {
	sqlQuery := `DELETE FROM notes WHERE session_id = $1 AND id = $2`
	if _, err := DB.Exec(sqlQuery, sid, id); err != nil {
		log.Println("DeleteNoteToDatabase Exec ", err.Error())
		return false, err
	}
	return true, nil
}
func GetNotesByIdFromDatabase(id string) (Notes, error) {
	var notes Notes
	var notesNull NotesNull
	sqlQuery := `SELECT id, session_id, note FROM notes WHERE id = $1`
	if err := DB.QueryRow(sqlQuery, id).Scan(&notes.ID, &notes.SessionId, &notesNull.Note); err != nil {
		log.Println("GetNotesByIdFromDatabase Scan ", err.Error())
		return notes, err
	}
	notes.Note = notesNull.Note.String

	return notes, nil
}
func VerifyNotesIdToDatabase(id uint32) bool {
	SqlQuery := `SELECT id FROM notes WHERE id = $1`
	if err := DB.QueryRow(SqlQuery, id).Scan(&id); err != nil && err == sql.ErrNoRows {
		log.Println("VerifyNotesIdToDatabase QueryRow ", err.Error())
		return false
	}
	return true
}
