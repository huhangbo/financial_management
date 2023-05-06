package service

import (
	"financial_management/model"
	"financial_management/setting"
)

func AddNotes(notes *model.Notes) error {
	var (
		db = setting.GetMySQL()
	)
	if err := db.Create(&notes).Error; err != nil {
		return err
	}
	return nil
}

func DeleteNotes(notesIDs []int) error {
	var (
		db        = setting.GetMySQL()
		notesList []*model.Notes
	)
	if err := db.Delete(&notesList, notesIDs).Error; err != nil {
		return err
	}
	return nil
}

func UpdateNotes(notes *model.Notes) error {
	var (
		db = setting.GetMySQL()
	)
	if err := db.Save(&notes).Error; err != nil {
		return err
	}
	return nil
}

func GetNotesByID(notesID int) *model.Notes {
	var (
		db       = setting.GetMySQL()
		tmpNotes = &model.Notes{
			NotesID: notesID,
		}
	)
	if err := db.First(&tmpNotes).Error; err != nil {
		return nil
	}
	return tmpNotes
}

func GetNotesByIDs(noteIDs []int) []*model.Notes {
	var (
		db        = setting.GetMySQL()
		notesList []*model.Notes
	)
	if err := db.Find(&notesList, noteIDs).Error; err != nil {
		return nil
	}
	return notesList
}

func GetNotesList(userID int, offset int, limit int) []*model.Notes {
	var (
		db        = setting.GetMySQL()
		notesList []*model.Notes
	)
	if err := db.Where("user_id = ?", userID).Find(&notesList).Order("created_at desc").Offset(offset).Limit(limit).Error; err != nil {
		return nil
	}
	return notesList
}
