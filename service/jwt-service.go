package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type JWTService interface {
	GenerateToken(name, password string) string
	ValidateToken(tokenString string) (*jwt.Token, error)
}
// jwtCustomClaims are custom claims extending default ones.
type jwtCustomClaims struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "mysite.com",
	}
}

func getSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (j *jwtService) GenerateToken(name, password string) string  {
//Set custom and standard claims
	claims:= &jwtCustomClaims{
		name,
		password,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour*48).Unix(),
			Issuer: j.issuer,
			IssuedAt: time.Now().Unix(),
		},
	}
	// create token with claims
	token:= jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//generate encoded token using the secret signing key
	t, err:= token.SignedString([]byte(j.secretKey))
	if err!=nil{
		panic(err)
	}
	return t
}

func (j *jwtService)ValidateToken(token string) (*jwt.Token, error)  {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		//signing method validation
		if _, ok := token.Method.(*jwt.SigningMethodHMAC);!ok{
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		//return the secret signing key
		return []byte(j.secretKey), nil
	})

}