package controllers

import (
	"fmt" 
	"log"  
	"time" 
	"context" 
	"github.com/gin-gonic/gin" 
	"strconv" 
	"net/http"
	"github.com/go-playground/validator/v10"  
	helper "go-jwt-mongo/utils/helperFunctions" 
	"golang.org/x/crypto/bcrypt" 
	"go-jwt-mongo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go-jwt-mongo/db"

)

var userCollection *mongo.Collection = db.OpenCollection(db.client ,"user")
var validate = validator.New()

func Signup() gin.HandlerFunc{

	return func (c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Backgound(), 100 *time.Second)
		var user models.User
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error()
			})

			return

		}

		validateErr := validate.Struct(user)
		if validateErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err":validateErr.Error()
			})

			return
		}

		 count ,err := userCollection.CountDocuments(ctx , bson.M{"email" :user.Email })
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error":err.Error()
			})
		}

		count ,err := userCollection.CountDocuments(ctx , bson.M{"phone" :user.Phone })
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error":err.Error()
			})
		}

		if count > 0 {
		
			c.JSON(http.StatusBadRequest, gin.H{
				"error":"Email or Phone already exists"
			})

		}

		user.Created_at = time.Parse(time.RFC3339 , time.Now().Format(time.RFC3339))
		user.Updated_at = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()

		user.user_id = user.ID.Hex()
		token, refreshToken = helper.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name , user.User_type )
		user.Token = &token
		user.Refresh_token = &refreshToken
		resultInsertionNum , InsertErr := userCollection.InsertOne(ctx, user)

		if InsertErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":"User Not Created"
			})

			return
		}
		defer cancel()
		c.JSON(http.StatusOk, resultInsertionNum)





	}

}

func Login() gin.HandlerFunc{
	return func(c *gin.Context){
		ctx , cancel = context.WithTimeout(context.Backgound() , 100 *time.Second )
		var user models.User
		var foundUser models.User
		c.BindJSON(&user)
		err := userCollection.FindOne(ctx , bson.M{
			"email":user.Email
		}).Decode(&foundUser)
		defer cancel()

		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"error":err.Error()
			})
		}

		passwordIsValid , msg := VerifyPassword(*user.Password, *foundUser.Password)
		defer cancel()

		passwordIsValid != true {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":msg
			})

			return 
		}

		if foundUser.Email == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":"User NOt FOund"
			})
		}


	}

}

func GetAllUsers(){

}

func GetUserByID() gin.HandlerFunc{ // HandlerFunc return a Function itself
	return func(c *gin.Context){
		userId := c.Param("user_id")

	err := helper.MatchUserTypeToUId(c, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error()
		})
		return
	}

	var ctx , cancel = context.WithTimeout(context.Backgound() , 100*time.Second)

	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"user_id":userId}).Decode(&user) // goland do not understand JSON that's why we are decoding because MOngoDB send data as JSON
	defer cancel()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errro":err.Error()
		})
	}

	return c.JSON(http.StatusOk, user)



	}
	
}

func HashPassword(password string) string{
	err := bcrypt.GenerateFromPassword([]byte(password),14)
	if err != nil {
		log.Panic(err)
		return
	}

	return string(bytes)

}

func VerifyPassword(userPassword string, foundPassword string) (bool, string){
	err := bcrypt.CompareHashAndPassword([]byte(foundPassword) , []byte(userPassword))
	check := true
	msg := "Password is correct!! "

	if err != nil{
		check = false
		msg  = "Wrong Password"
		
	}
	return check ,msg


	
}
