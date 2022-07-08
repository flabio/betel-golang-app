package InterfacesService

import "github.com/dgrijalva/jwt-go"

//JWTService is a contract of what jwtService can do
type IJWTService interface {
	GenerateToken(userId string, RolId string, ChurchId string) string
	ValidateToken(token string) (*jwt.Token, error)
}
