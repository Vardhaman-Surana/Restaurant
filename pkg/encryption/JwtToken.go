package encryption

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/vds/Restraunt/pkg/models"
	"time"
)

func CreateToken(userName string) (string,error){
	jwtKey:=[]byte("SecretKey")
	expirationTime:=time.Now().Add(5*time.Minute)
	claims:=&models.Claims{
		UserName:userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt:expirationTime.Unix(),
		},
	}
	//remember to change it later
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err!=nil{
		return "",err
	}
	return tokenString,nil
}
