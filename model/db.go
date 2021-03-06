package model

import "github.com/jinzhu/gorm"

var Db *gorm.DB

// 初始化数据库
func init() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{}, &ExamPaper{}, &Chat{}, &Transaction{}, &Solution{}, &Message{})
	Db = db
}