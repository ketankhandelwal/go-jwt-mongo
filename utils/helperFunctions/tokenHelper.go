package helper

import (
	"fmt"
	"log"
	"context"
	"os"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go-jwt-mongo/db"

)

type SignedDetails struct {
	Email string
	First_name string
	Last_name string
	Uid string
	User_type string
	jwt.StandardClaims // its a struct

}



var userCollection *mongo.Collection = db.OpenCollection(db.Client , "user")

var secretKey string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email string, firstName string, lastName string, userType string , uid string) () {
	claims := &SignedDetails{
		Email: email,
		First_name: firstName,
		Last_name: lastName,
		Uid: uid,
		User_type : userType,
		StandardClaims : jwt.StandardClaims{
			ExpiresAt : time.Now().Local().Add(time.Hour * time.Duration(24)).Unix()
		}

	}

	refreshClaims := &SignedDetails{
	StandardClaims : jwt.StandardClaims{
		ExpiresAt : time.Now().Add(time.Hour * time.Duration(168)).Unix()
	}


	}

	token, err :=jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))
	refreshToken , err :=jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(secretKey))

	if err != nil{
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}





