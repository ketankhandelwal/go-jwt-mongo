package helper

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// func CheckUserType(ctx *gin.Context , role string) (err error){

// }

func MatchUserTypeToUId(ctx *gin.Context, user_id string) (err error) {

	uID := ctx.GetString("user_id")
	user_type = ctx.GetString("user_type")
	err := nil

	if user_type == "USER" && uID != user_id {
		err = errors.New("Unauthorize")
		return err

	}
}
