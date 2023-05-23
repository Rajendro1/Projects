package notes

import (
	"log"

	"github.com/Rajendro1/Projects/AccuKnoxApi/config"
	pgdatabase "github.com/Rajendro1/Projects/AccuKnoxApi/pgDatabase"
	"github.com/gin-gonic/gin"
)

func CreateNotes() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var input = notes{}
		if err := c.ShouldBindJSON(&input); err == nil {
			if pgdatabase.VerifyUsersSessionIDToDatabase(input.SID) {
				createNote, noteID, createNoteErr := pgdatabase.CreateNotesToDatabase(input.SID, input.Note)
				if createNote && createNoteErr == nil {
					c.JSON(config.SUCCESS, gin.H{
						"status":  config.SUCCESS,
						"message": "Note created successfully",
						"id":      noteID,
					})
				} else {
					c.JSON(config.SUCCESS, gin.H{
						"status":  config.DBERROR,
						"message": "Sorry! we can't create the note",
						"error":   createNoteErr.Error(),
					})
				}
			} else {
				c.JSON(config.UNAUTHORIZE, gin.H{
					"status":  config.UNAUTHORIZE,
					"message": "Please give valid session id",
				})
			}
		} else {
			log.Println("input ERROR ", err.Error())
			c.JSON(config.INPUTERROR, gin.H{
				"status":  config.INPUTERROR,
				"message": "input error",
				"error":   err.Error(),
			})
		}
	}
	return gin.HandlerFunc(fn)
}
func DeleteNote() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var input = notes{}
		if err := c.ShouldBindJSON(&input); err == nil {
			if pgdatabase.VerifyUsersSessionIDToDatabase(input.SID) && pgdatabase.VerifyNotesIdToDatabase(input.ID) {
				deleteNote, deleteNoteErr := pgdatabase.DeleteNoteToDatabase(input.SID, input.ID)
				if deleteNote && deleteNoteErr == nil {
					c.JSON(config.SUCCESS, gin.H{
						"status":  config.SUCCESS,
						"message": "Your notes deleted successfully",
					})
				} else {
					c.JSON(config.SUCCESS, gin.H{
						"status":  config.DBERROR,
						"message": "Sorry! we can't delete your notes",
					})
				}
			} else if !pgdatabase.VerifyUsersSessionIDToDatabase(input.SID) {
				c.JSON(config.UNAUTHORIZE, gin.H{
					"status":  config.UNAUTHORIZE,
					"message": "Please give valid session id",
				})
			} else if !pgdatabase.VerifyNotesIdToDatabase(input.ID) {
				c.JSON(config.INPUTERROR, gin.H{
					"status":  config.UNAUTHORIZE,
					"message": "Please give valid notes id",
				})
			}
		} else {
			log.Println("input ERROR ", err.Error())
			c.JSON(config.INPUTERROR, gin.H{
				"status":  config.INPUTERROR,
				"message": "input error",
				"error":   err.Error(),
			})
		}
	}
	return gin.HandlerFunc(fn)
}
func GetNotes() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var input = notes{}
		if err := c.ShouldBindJSON(&input); err == nil {
			if pgdatabase.VerifyUsersSessionIDToDatabase(input.SID) {
				getNotesBySid, getNotesBySidErr := pgdatabase.GetNotesBySessionIdFromDatabase(input.SID)
				if getNotesBySidErr == nil {
					c.JSON(config.SUCCESS, gin.H{
						"status":  config.SUCCESS,
						"message": "Notes getting successfully",
						"notes":   getNotesBySid,
					})
				} else {
					c.JSON(config.SUCCESS, gin.H{
						"status":  config.DBERROR,
						"message": "Sorry! we can't get your notes",
						"error":   getNotesBySidErr.Error(),
					})
				}
			} else if !pgdatabase.VerifyUsersSessionIDToDatabase(input.SID) {
				c.JSON(config.UNAUTHORIZE, gin.H{
					"status":  config.UNAUTHORIZE,
					"message": "Please give valid session id",
				})
			}
		} else {
			log.Println("input ERROR ", err.Error())
			c.JSON(config.INPUTERROR, gin.H{
				"status":  config.INPUTERROR,
				"message": "input error",
				"error":   err.Error(),
			})
		}
	}
	return gin.HandlerFunc(fn)
}
