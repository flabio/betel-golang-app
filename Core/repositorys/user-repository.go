package repositorys

import (
	"bete/Core/Interfaces"
	"bete/Core/entity"
	constantvariables "bete/Infrastructure/constantVariables"
	"sync"
)

//NewUserRepository is creates a new instance of UserRepository

func NewUserRepository() Interfaces.IUser {
	var (
		_OPEN *OpenConnections
		_ONCE sync.Once
	)
	_ONCE.Do(func() {
		_OPEN = &OpenConnections{

			connection: entity.Factory(constantvariables.OPTION_FACTORY_DB),
		}
	})
	return _OPEN
}

/*
@param user, is a struct of User
*/
func (db *OpenConnections) SetInsertUser(user entity.User) (entity.User, error) {
	db.mux.Lock()
	err := db.connection.Save(&user).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return user, err
}

/*
@param user, is a struct of User
*/
func (db *OpenConnections) SetEditUser(user entity.User, Id uint) (entity.User, error) {
	db.mux.Lock()
	err := db.connection.Where("id=?", Id).Save(&user).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return user, err
}

/*
@param group, is a struct of UserSubdetachement
*/
func (db *OpenConnections) SetInsertGroup(group entity.UserSubdetachement) error {
	db.mux.Lock()
	err := db.connection.Save(&group).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return err
}

/*
@param gruop, is a struct of UserSubdetachement
*/
func (db *OpenConnections) SetEditGroup(gruop entity.UserSubdetachement) (entity.UserSubdetachement, error) {
	var gruo entity.UserSubdetachement
	db.mux.Lock()
	err := db.connection.Where("user_id =?", gruop.UserId).Updates(&gruop).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return gruo, err
}

/*
@param email, is a string of user
@param password, is a string of user
*/

func (db *OpenConnections) VerifyCredential(email string, password string) interface{} {
	var user entity.User
	db.mux.Lock()
	err := db.connection.Preload("Rol").Preload("Church").Preload("City").Where("email = ?", email).
		Find(&user).Error
	defer entity.Closedb()
	defer db.mux.Unlock()

	if err == nil {
		return user
	}
	return nil
}

/*
@param email, is a string of user
*/
func (db *OpenConnections) IsDuplicateEmail(email string) (bool, error) {
	var user entity.User
	db.mux.Lock()
	err := db.connection.Where("email = ?", email).Take(&user).Error
	defer entity.Closedb()
	defer db.mux.Unlock()

	if err == nil {
		return true, err
	}
	return false, err
}

/*
@param identification, is a string of user
*/
func (db *OpenConnections) IsDuplicateIdentificatio(identification string) bool {
	var user entity.User
	db.mux.Lock()
	err := db.connection.Where("identification = ?", identification).Take(&user).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	if err == nil {
		return true
	}
	return false
}

/*
@param id, is a uint of user
*/
func (db *OpenConnections) SetRemoveUser(id uint) (bool, error) {
	var user entity.User
	db.mux.Lock()
	err := db.connection.Where("id=?", id).Delete(&user).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	if err == nil {
		return true, err
	}
	return false, err
}
func (db *OpenConnections) GetAllUser() ([]entity.User, error) {
	var user []entity.User
	db.mux.Lock()
	err := db.connection.Preload("Rol").
		Preload("Church").
		Preload("City").
		Find(&user).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return user, err
}

/*
@param email, is a string of user
*/
func (db *OpenConnections) GetFindByEmail(email string) (entity.User, error) {
	var user entity.User
	db.mux.Lock()
	err := db.connection.Preload("Rol").Where("email = ?", email).Take(&user).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	if err == nil {
		return user, err
	}
	return user, err
}

/*
@param userID, is a uint of user
*/
func (db *OpenConnections) GetProfileUser(userID uint) (entity.User, error) {
	var user entity.User
	db.mux.Lock()
	err := db.connection.Preload("Rol").
		Preload("Church").
		Find(&user, userID).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	if err == nil {
		return user, err
	}
	return user, err
}

func (db *OpenConnections) GetCountUser() int64 {
	var count int64
	db.mux.Lock()
	db.connection.Table("users").Count(&count)
	defer entity.Closedb()
	defer db.mux.Unlock()
	return count
}

/*
@param begin
@param limit is a int
*/
func (db *OpenConnections) GetPaginationUsers(begin, limit int) ([]entity.User, error) {
	var user []entity.User
	db.mux.Lock()
	err := db.connection.Offset(begin).
		Limit(limit).
		Order("id desc").
		Preload("Rol").
		Preload("Church").
		Find(&user).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return user, err

}

/*
@param user, is a struct of user
*/
func (db *OpenConnections) SetChangePassword(user entity.User) error {
	db.mux.Lock()
	err := db.connection.Where("id =?", user.Id).Update("password", user.Password).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return err
}

// ListNavigators
func (db *OpenConnections) GetListNavigators() ([]entity.User, error) {
	var user []entity.User
	db.mux.Lock()
	err := db.connection.
		Where("rol_id", constantvariables.KING_SCOUTS).
		Order("users.id desc").
		Group("users.id").
		Preload("Rol").
		Find(&user).Error
	defer entity.Closedb()
	defer db.mux.Unlock()

	return user, err
}

//fin ListNavigators

// ListPioneers
func (db *OpenConnections) GetListPioneers() ([]entity.User, error) {
	var user []entity.User
	db.mux.Lock()
	err := db.connection.
		Where("rol_id", constantvariables.KING_SCOUTS).
		Order("users.id desc").
		Group("users.id").
		Preload("Rol").
		Find(&user).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return user, err
}

//fin ListPioneers

// ListFollowersWays
func (db *OpenConnections) GetListFollowersWays() ([]entity.User, error) {
	var user []entity.User
	db.mux.Lock()
	err := db.connection.
		Where("rol_id", constantvariables.KING_SCOUTS).
		Order("users.id desc").
		Group("users.id").
		Preload("Rol").
		Find(&user).Error
	defer entity.Closedb()
	defer db.mux.Unlock()

	return user, err
}

//fin ListFollowersWays

// ListScouts
func (db *OpenConnections) GetListScouts() ([]entity.User, error) {
	var user []entity.User
	db.mux.Lock()
	err := db.connection.
		Where("rol_id", constantvariables.KING_SCOUTS).
		Order("users.id desc").
		Group("users.id").
		Preload("Rol").
		Find(&user).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return user, err
}

//fin ListScouts

// ListCommanders
func (db *OpenConnections) GetListCommanders() ([]entity.User, error) {
	var user []entity.User
	db.mux.Lock()
	err := db.connection.
		Order("users.id desc").
		Where("rol_id IN ?", []int{constantvariables.FIRST_MAJOR_ROL, constantvariables.SECOND_MAJOR_ROL, constantvariables.SECOND_COMMANDERS_ROL, constantvariables.SECOND_COMMANDERS_ROL}).
		Group("users.id").
		Preload("Rol").
		Find(&user).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return user, err
}

//fin ListCommanders

// ListMajors
func (db *OpenConnections) GetListMajors() ([]entity.User, error) {
	var user []entity.User
	db.mux.Lock()
	err := db.connection.
		Distinct("users.id").
		Order("users.id desc").
		Where("rol_id IN ?", []int{constantvariables.FIRST_MAJOR_ROL, constantvariables.SECOND_MAJOR_ROL}).
		Preload("Rol").Preload("Church").Preload("SubDetachment").
		Find(&user).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return user, err
}

// fin ListMajors
/*
@param Id, is a uint of user
*/
func (db *OpenConnections) GetListKingsScouts(Id uint) ([]entity.User, error) {
	var user []entity.User
	db.mux.Lock()
	err := db.connection.Preload("Rol").
		Where("rol_id", constantvariables.KING_SCOUTS).
		Group("users.id").Find(&user).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	return user, err
}

//count scouts of king

func (db *OpenConnections) GetCounKanban() (int64, int64, int64, int64, int64) {
	var count_navigators int64
	var count_pioneers int64
	var count_followers int64
	var count_scouts int64
	var count_commanders int64
	db.mux.Lock()
	db.connection.
		Where("rol_id", constantvariables.KING_SCOUTS).
		Group("users.id").
		Table("users").Count(&count_navigators)

	db.connection.
		Where("rol_id", constantvariables.KING_SCOUTS).
		Group("users.id").
		Table("users").Count(&count_pioneers)

	db.connection.
		Where("rol_id", constantvariables.KING_SCOUTS).
		Group("users.id").
		Table("users").Count(&count_followers)

	db.connection.
		Where("rol_id", constantvariables.KING_SCOUTS).
		Group("users.id").
		Table("users").Count(&count_scouts)
	//db.connection.Joins("left join roles on roles.user_id = users.id").Joins("left join sub_detachments on users.sub_detachment_id=sub_detachments.id").Where("sub_detachments.id", 2).Where("roles.rol_id", 29).Group("users.id").Table("users").Count(&count_commanders)

	db.connection.
		Where("rol_id IN ?", []int{constantvariables.NAVIGANTORS_ROL, constantvariables.PIONEERS_ROL, constantvariables.SECOND_COMMANDERS_ROL, constantvariables.SCOUTS_ROL}).
		Group("users.id").
		Table("users").Count(&count_commanders)
	defer entity.Closedb()
	defer db.mux.Unlock()
	return count_navigators, count_pioneers, count_followers, count_scouts, count_commanders
}

//FindUserNameLastName
/*
@param data, is a string
*/
func (db *OpenConnections) GetFindUserNameLastName(data string) ([]entity.User, error) {

	var user []entity.User
	db.mux.Lock()
	err := db.connection.Preload("Rol").
		Preload("Church").
		Where("concat(name,' ',last_name) LIKE ?", "%"+string(data)+"%").
		Find(&user).Error
	defer entity.Closedb()
	defer db.mux.Unlock()
	if err == nil {
		return user, err
	}

	return user, err
}

//find FindUserNameLastName
