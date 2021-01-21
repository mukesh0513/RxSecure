package service

import (
	"github.com/mukesh0513/RxSecure/internal/model"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetTokenizeValue(c *gin.Context, db *gorm.DB, args model.GetApiParams) (model.Keys, error) {
	var post model.Keys

	if err := db.Where("token = ?", args.Token).Find(&post).Error; err != nil {
		log.Println(err)
		return post, err
	}
	return post, nil
}

func CreateToken(c *gin.Context, db *gorm.DB, post *model.Keys) (*model.Keys, error) {

	//Logic to generate token.

	return post, nil
}
