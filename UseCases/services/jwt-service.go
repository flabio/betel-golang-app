package services

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//JWTService is a contract of what jwtService can do
type JWTService interface {
	GenerateToken(userId string, RolId string, SubDetachmentId string, ChurchId string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserId          string `json:"id"`
	RolId           string `json:"rol"`
	SubDetachmentId string `json:"subdetachmentid"`
	ChurchId        string `json:"churchid"`
	jwt.StandardClaims
}
type jwtService struct {
	secretKey string
	issuer    string
}

//NewJWTService method is creates a new instance of JWTService
func NewJWTService() JWTService {
	return &jwtService{
		issuer: "test",

		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "test"
	}
	return secretKey
}
func (j *jwtService) GenerateToken(UserId string, RolId string, SubDetachmentId string, ChurchId string) string {
	claims := &jwtCustomClaim{
		UserId,
		RolId,
		SubDetachmentId,
		ChurchId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			//ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
			Issuer:   j.issuer,
			Id:       UserId,
			IssuedAt: time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		checkError(err)
		//fmt.Println(err.Error())
	}
	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {

	return jwt.Parse(
		token, func(t_ *jwt.Token) (interface{}, error) {

			if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {

				return nil, fmt.Errorf("Unexpected signig method %v", t_.Header["alg"])
			}
			return []byte(j.secretKey), nil
		})
}
