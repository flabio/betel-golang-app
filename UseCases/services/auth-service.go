package services

import (
	"bete/Core/entity"
	"bete/Core/repositorys"
	"bete/UseCases/dto"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

//authService is a contract about something that this service can do
type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.UserDTO) entity.User
	FindByEmail(email string) (entity.User, error)
	IsDuplicateEmail(email string) (bool, error)
	GetIdRol(id uint) uint
}

type authService struct {
	userRepository repositorys.UserRepository
}

//NewauthService creates a new instance of authService
func NewAuthService() AuthService {

	var userRepo repositorys.UserRepository = repositorys.NewUserRepository()
	return &authService{
		userRepository: userRepo,
	}
}

func (authService *authService) VerifyCredential(email string, password string) interface{} {
	res := authService.userRepository.VerifyCredential(email, password)

	if v, ok := res.(entity.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if comparedPassword {
			if v.Email == email && comparedPassword {
				return res
			}
		}
		return nil
	}
	return nil
}

func (authService *authService) CreateUser(user dto.UserDTO) entity.User {
	userToCreate := entity.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	checkError(err)

	res, err := authService.userRepository.InsertUser(userToCreate)

	return res
}

func (authService *authService) FindByEmail(email string) (entity.User, error) {
	return authService.userRepository.FindByEmail(email)
}

func (authService *authService) IsDuplicateEmail(email string) (bool, error) {
	return authService.userRepository.IsDuplicateEmail(email)

}

func (authService *authService) GetIdRol(id uint) uint {

	user, _ := authService.userRepository.ProfileUser(id)
	return user.Roles.RolId
}
func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	bcrypt.CompareHashAndPassword(byteHash, plainPassword)

	return true
}
