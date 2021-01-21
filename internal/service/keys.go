package service

import (
	"github.com/mukesh0513/RxSecure/internal/database"
	"github.com/mukesh0513/RxSecure/internal/model"
	"log"

	"github.com/gin-gonic/gin"
)

func GetTokenizeValue(c *gin.Context, args model.GetApiParams) (model.Keys, error) {
	var post model.Keys

	if err := database.DB.Where("token = ?", args.Token).Find(&post).Error; err != nil {
		log.Println(err)
		return post, err
	}
	return post, nil
}

func CreateToken(c *gin.Context, post *model.Keys) (*model.Keys, error) {

	//Logic to generate token.

	return post, nil
}
