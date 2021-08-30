package helper

import (
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtClaim struct {
	jwt.StandardClaims

	student_number string
}

var key = "Matilda"

func CreatToken(student_number string) string {
	claim := jwtClaim{}

	claim.student_number = student_number
	claim.ExpiresAt = time.Now().Add(200 * time.Hour).Unix()
	claim.IssuedAt = time.Now().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString([]byte(key))

	if err != nil {
		log.Println(err)
	}

	return signedToken
}

func VerifyToken(token string) (string, error) {
	TempToken, err := jwt.ParseWithClaims(token, &jwtClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("Matilda"), nil
	})
	if err != nil {
		return "", errors.New("token解析失败")
	}

	claims, ok := TempToken.Claims.(*jwtClaim)
	if !ok {
		return "", errors.New("发生错误")
	}
	if err := TempToken.Claims.Valid(); err != nil {
		return "", errors.New("发生错误")
	}

	return claims.student_number, nil
}
