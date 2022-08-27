package Interfaces

import "bete/Core/entity"

type IUser interface {
	SetInsertUser(user entity.User) (entity.User, error)
	SetEditUser(user entity.User, Id uint) (entity.User, error)
	SetInsertGroup(group entity.UserSubdetachement) error
	SetEditGroup(group entity.UserSubdetachement) (entity.UserSubdetachement, error)
	GetAllUser() ([]entity.User, error)
	GetPaginationUsers(begin, limit int) ([]entity.User, error)
	SetRemoveUser(id uint) (bool, error)
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (bool, error)
	GetFindUserNameLastName(data string) ([]entity.User, error)
	GetFindByEmail(email string) (entity.User, error)
	IsDuplicateIdentificatio(identification string) bool
	GetProfileUser(userId uint) (entity.User, error)
	SetChangePassword(user entity.User) error
	GetCountUser() int64
	GetListNavigators() ([]entity.User, error)
	GetListPioneers() ([]entity.User, error)
	GetListFollowersWays() ([]entity.User, error)
	GetListScouts() ([]entity.User, error)
	GetListKingsScouts(Id uint) ([]entity.User, error)
	GetListCommanders() ([]entity.User, error)
	GetListMajors() ([]entity.User, error)
	GetCounKanban() (int64, int64, int64, int64, int64)
}
