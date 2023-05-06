package service

import (
	"financial_management/model"
	"financial_management/setting"
)

func AddNews(news *model.News) error {
	var (
		db = setting.GetMySQL()
	)
	if err := db.Create(&news).Error; err != nil {
		return err
	}
	return nil
}

func DeleteNews(newsIDs []int) error {
	var (
		db       = setting.GetMySQL()
		newsList []*model.News
	)
	if err := db.Delete(&newsList, newsIDs).Error; err != nil {
		return err
	}
	return nil
}

func UpdateNews(news *model.News) error {
	var (
		db = setting.GetMySQL()
	)
	if err := db.Save(&news).Error; err != nil {
		return err
	}
	return nil
}

func GetNewsByID(newsID int) *model.News {
	var (
		db      = setting.GetMySQL()
		tmpNews *model.News
	)
	if err := db.First(&tmpNews, newsID).Error; err != nil {
		return nil
	}
	return tmpNews
}

func GetNewsByIDs(noteIDs []int) []*model.News {
	var (
		db        = setting.GetMySQL()
		notesList []*model.News
	)
	if err := db.Find(&notesList, noteIDs).Error; err != nil {
		return nil
	}
	return notesList
}

func GetNewsList(offset int, limit int) []*model.News {
	var (
		db        = setting.GetMySQL()
		notesList []*model.News
	)
	if err := db.Find(&notesList).Order("creat_at desc").Offset(offset).Limit(limit).Error; err != nil {
		return nil
	}
	return notesList
}
