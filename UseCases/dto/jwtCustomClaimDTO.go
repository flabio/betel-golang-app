package dto

import "github.com/dgrijalva/jwt-go"

type JwtCustomClaimDTO struct {
	UserId string `json:"id"`
	RolId  string `json:"rol"`

	ChurchId string `json:"churchid"`
	jwt.StandardClaims
}

// type jwtServiceDTO struct {
// 	secretKey string
// 	issuer    string
// }
