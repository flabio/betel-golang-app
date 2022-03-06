package repositorys

import (
	"bete/Core/entity"
	constantvariables "bete/Infrastructure/constantVariables"

	"gorm.io/gorm"
)

//UserRepository is contract what UserRepository can do to db
type UserRepository interface {
	SetInsertUser(user entity.User) (entity.User, error)

	SetEditUser(user entity.User) (entity.User, error)
	SetInsertRole(role entity.Role) error
	SetEditRole(role entity.Role) (entity.Role, error)
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
	SetRemoveRoleUser(id uint) error

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

type userConnection struct {
	connection *gorm.DB
}

//NewUserRepository is creates a new instance of UserRepository

func NewUserRepository() UserRepository {
	var db *gorm.DB = entity.DatabaseConnection()
	return &userConnection{
		connection: db,
	}
}

var errChanUser = make(chan error, constantvariables.CHAN_VALUE)

func (db *userConnection) SetInsertUser(user entity.User) (entity.User, error) {

	go func() {
		err := db.connection.Save(&user).Error
		defer entity.Closedb()
		errChanUser <- err
	}()
	err := <-errChanUser

	return user, err
}

func (db *userConnection) SetEditUser(user entity.User) (entity.User, error) {
	go func() {
		err := db.connection.Save(&user).Error
		defer entity.Closedb()
		errChanUser <- err
	}()
	err := <-errChanUser
	return user, err
}

func (db *userConnection) SetInsertRole(role entity.Role) error {

	go func() {
		err := db.connection.Save(&role).Error
		defer entity.Closedb()
		errChanUser <- err
	}()
	err := <-errChanUser
	return err
}
func (db *userConnection) SetInsertGroup(group entity.UserSubdetachement) error {

	go func() {
		err := db.connection.Save(&group).Error
		defer entity.Closedb()
		errChanUser <- err
	}()
	err := <-errChanUser
	return err
}

func (db *userConnection) SetEditRole(role entity.Role) (entity.Role, error) {
	var rol entity.Role

	go func() {
		err := db.connection.Where("user_id =?", role.UserId).Updates(&role).Error
		defer entity.Closedb()
		errChanUser <- err
	}()
	err := <-errChanUser
	return rol, err
}

func (db *userConnection) SetEditGroup(gruop entity.UserSubdetachement) (entity.UserSubdetachement, error) {
	var gruo entity.UserSubdetachement

	go func() {
		err := db.connection.Where("user_id =?", gruop.UserId).Updates(&gruop).Error
		defer entity.Closedb()
		errChanUser <- err
	}()
	err := <-errChanUser
	return gruo, err
}

func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user entity.User

	go func() {
		err := db.connection.Preload("SubDetachment").
			Preload("Roles.Rol.RoleModule.Module").
			Select("users.id,users.name,users.last_name,users.email,users.image,users.sub_detachment_id,users.church_id,roles.id as idrol").
			Joins("left join roles on roles.user_id = users.id").
			Where("email = ?", email).
			Find(&user).Error
		defer entity.Closedb()
		errChanUser <- err
	}()
	err := <-errChanUser
	if err == nil {
		return user
	}
	return nil
}

func (db *userConnection) IsDuplicateEmail(email string) (bool, error) {
	var user entity.User

	go func() {
		err := db.connection.Where("email = ?", email).Take(&user).Error
		defer entity.Closedb()
		errChanUser <- err
	}()
	err := <-errChanUser
	if err == nil {
		return true, err
	}
	return false, err
}
func (db *userConnection) IsDuplicateIdentificatio(identification string) bool {
	var user entity.User

	go func() {
		err := db.connection.Where("identification = ?", identification).Take(&user).Error
		defer entity.Closedb()
		errChanUser <- err
	}()
	err := <-errChanUser

	if err == nil {
		return true
	}
	return false
}

func (db *userConnection) SetRemoveUser(id uint) (bool, error) {
	var user entity.User

	go func() {
		err := db.connection.Where("id=?", id).Delete(&user).Error
		defer entity.Closedb()
		errChanUser <- err
	}()
	err := <-errChanUser
	if err == nil {
		return true, err
	}
	return false, err
}
func (db *userConnection) GetAllUser() ([]entity.User, error) {
	var user []entity.User
	go func() {
		err := db.connection.Preload("Roles.Rol").
			Preload("Church").
			Preload("SubDetachment").
			Find(&user).Error
		defer entity.Closedb()
		errChanUser <- err
	}()
	err := <-errChanUser
	return user, err
}

func (db *userConnection) GetFindByEmail(email string) (entity.User, error) {
	var user entity.User

	go func() {
		err := db.connection.Preload("Roles.Rol").Where("email = ?", email).Take(&user).Error
		defer entity.Closedb()
		errChanUser <- err
	}()
	err := <-errChanUser
	if err == nil {
		return user, err
	}
	return user, err
}

func (db *userConnection) GetProfileUser(userID uint) (entity.User, error) {
	var user entity.User

	go func() {
		err := db.connection.Preload("City").
			Preload("SubDetachment").
			Preload("Roles.Rol").
			Preload("Roles.StudyCarried").
			Preload("MinisterialAcademys").
			Preload("Church").
			Find(&user, userID).Error
		defer entity.Closedb()
		errChanUser <- err
	}()
	err := <-errChanUser
	if err == nil {
		return user, err
	}
	return user, err
}

func (db *userConnection) SetRemoveRoleUser(id uint) error {
	var role entity.Role

	go func() {
		err := db.connection.Where("user_id=?", id).Delete(&role).Error
		defer entity.Closedb()
		errChanUser <- err
	}()
	err := <-errChanUser
	return err
}

func (db *userConnection) GetCountUser() int64 {
	var count int64

	go func() {
		err := db.connection.Table("users").Count(&count).Error
		defer entity.Closedb()
		errChanUser <- err
	}()
	<-errChanUser
	return count
}
func (db *userConnection) GetPaginationUsers(begin, limit int) ([]entity.User, error) {
	var user []entity.User

	go func() {
		err := db.connection.Offset(begin).
			Limit(limit).
			Order("id desc").
			Preload("Roles.Rol").
			Preload("Church").
			Preload("SubDetachment").
			Find(&user).Error
		defer entity.Closedb()
		errChanUser <- err
	}()
	err := <-errChanUser
	return user, err

}
func (db *userConnection) SetChangePassword(user entity.User) error {

	go func() {
		err := db.connection.Where("id =?", user.Id).Update("password", user.Password).Error
		defer entity.Closedb()
		errChanUser <- err
	}()
	err := <-errChanUser
	return err
}

//ListNavigators
func (db *userConnection) GetListNavigators() ([]entity.User, error) {
	var user []entity.User

	go func() {
		err := db.connection.Joins("left join roles on roles.user_id = users.id").
			Joins("left join sub_detachments on users.sub_detachment_id=sub_detachments.id").
			Where("sub_detachments.id", constantvariables.NAVIGANTORS_SUB_DETACHMENT).
			Where("roles.rol_id", constantvariables.KING_SCOUTS).
			Order("users.id desc").
			Group("users.id").
			Preload("Roles.Rol").
			Find(&user).Error
		defer entity.Closedb()
		errChanUser <- err
		close(errChanUser)
	}()
	err := <-errChanUser

	return user, err
}

//fin ListNavigators
//ListPioneers
func (db *userConnection) GetListPioneers() ([]entity.User, error) {
	var user []entity.User

	go func() {
		err := db.connection.Joins("left join roles on roles.user_id = users.id").
			Joins("left join sub_detachments on users.sub_detachment_id=sub_detachments.id").
			Where("sub_detachments.id", constantvariables.PIONEERS_SUB_DETACHMENT).
			Where("roles.rol_id", constantvariables.KING_SCOUTS).
			Order("users.id desc").
			Group("users.id").
			Preload("Roles.Rol").
			Find(&user).Error
		defer entity.Closedb()
		errChanUser <- err
	}()
	err := <-errChanUser
	return user, err
}

//fin ListPioneers
//ListFollowersWays
func (db *userConnection) GetListFollowersWays() ([]entity.User, error) {
	var user []entity.User

	go func() {
		err := db.connection.Joins("left join roles on roles.user_id = users.id").
			Joins("left join sub_detachments on users.sub_detachment_id=sub_detachments.id").
			Where("sub_detachments.id", constantvariables.PATH_FOLLOWERS_SUB_DETACHMENT).
			Where("roles.rol_id", constantvariables.KING_SCOUTS).
			Order("users.id desc").
			Group("users.id").
			Preload("Roles.Rol").
			Find(&user).Error
		defer entity.Closedb()
		errChanUser <- err
	}()
	err := <-errChanUser
	return user, err
}

//fin ListFollowersWays
//ListScouts
func (db *userConnection) GetListScouts() ([]entity.User, error) {
	var user []entity.User

	go func() {
		err := db.connection.Joins("left join roles on roles.user_id = users.id").
			Joins("left join sub_detachments on users.sub_detachment_id=sub_detachments.id").
			Where("sub_detachments.id", constantvariables.SCOUTS_SUB_DETACHMENT).
			Where("roles.rol_id", constantvariables.KING_SCOUTS).
			Order("users.id desc").
			Group("users.id").
			Preload("Roles.Rol").
			Find(&user).Error
		defer entity.Closedb()
		errChanUser <- err
	}()
	err := <-errChanUser
	return user, err
}

//fin ListScouts
//ListCommanders
func (db *userConnection) GetListCommanders() ([]entity.User, error) {
	var user []entity.User

	go func() {
		err := db.connection.Joins("left join roles on roles.user_id = users.id").
			Order("users.id desc").
			Where("roles.rol_id IN ?", []int{constantvariables.FIRST_MAJOR_ROL, constantvariables.SECOND_MAJOR_ROL, constantvariables.SECOND_COMMANDERS_ROL, constantvariables.SECOND_COMMANDERS_ROL}).
			Group("users.id").
			Preload("Roles.Rol").
			Find(&user).Error
		defer entity.Closedb()
		errChanUser <- err
	}()
	err := <-errChanUser
	return user, err
}

//fin ListCommanders
//ListMajors
func (db *userConnection) GetListMajors() ([]entity.User, error) {
	var user []entity.User

	go func() {
		err := db.connection.Joins("left join roles on roles.user_id = users.id").
			Distinct("users.id").
			Order("users.id desc").
			Where("roles.rol_id IN ?", []int{constantvariables.FIRST_MAJOR_ROL, constantvariables.SECOND_MAJOR_ROL}).
			Preload("Roles.Rol").Preload("Church").Preload("SubDetachment").
			Find(&user).Error
		defer entity.Closedb()
		errChanUser <- err
	}()
	err := <-errChanUser
	return user, err
}

// fin ListMajors

func (db *userConnection) GetListKingsScouts(Id uint) ([]entity.User, error) {
	var user []entity.User

	go func() {
		err := db.connection.Preload("Roles.Rol").
			Joins("left join roles on roles.user_id = users.id").
			Joins("left join sub_detachments on users.sub_detachment_id=sub_detachments.id").
			Where("sub_detachments.id", Id).
			Where("roles.rol_id", constantvariables.KING_SCOUTS).
			Group("users.id").Find(&user).Error
		defer entity.Closedb()
		errChanUser <- err
	}()
	err := <-errChanUser
	return user, err
}

//count scouts of king

func (db *userConnection) GetCounKanban() (int64, int64, int64, int64, int64) {
	var count_navigators int64
	var count_pioneers int64
	var count_followers int64
	var count_scouts int64
	var count_commanders int64

	db.connection.Joins("left join roles on roles.user_id = users.id").
		Joins("left join sub_detachments on users.sub_detachment_id=sub_detachments.id").
		Where("sub_detachments.id", constantvariables.NAVIGANTORS_SUB_DETACHMENT).Where("roles.rol_id", constantvariables.KING_SCOUTS).
		Group("users.id").
		Table("users").Count(&count_navigators)

	db.connection.Joins("left join roles on roles.user_id = users.id").
		Joins("left join sub_detachments on users.sub_detachment_id=sub_detachments.id").
		Where("sub_detachments.id", constantvariables.PIONEERS_SUB_DETACHMENT).
		Where("roles.rol_id", constantvariables.KING_SCOUTS).
		Group("users.id").
		Table("users").Count(&count_pioneers)

	db.connection.Joins("left join roles on roles.user_id = users.id").
		Joins("left join sub_detachments on users.sub_detachment_id=sub_detachments.id").
		Where("sub_detachments.id", constantvariables.PATH_FOLLOWERS_SUB_DETACHMENT).
		Where("roles.rol_id", constantvariables.KING_SCOUTS).
		Group("users.id").
		Table("users").Count(&count_followers)

	db.connection.Joins("left join roles on roles.user_id = users.id").
		Joins("left join sub_detachments on users.sub_detachment_id=sub_detachments.id").
		Where("sub_detachments.id", constantvariables.SCOUTS_SUB_DETACHMENT).
		Where("roles.rol_id", constantvariables.KING_SCOUTS).
		Group("users.id").
		Table("users").Count(&count_scouts)
	//db.connection.Joins("left join roles on roles.user_id = users.id").Joins("left join sub_detachments on users.sub_detachment_id=sub_detachments.id").Where("sub_detachments.id", 2).Where("roles.rol_id", 29).Group("users.id").Table("users").Count(&count_commanders)

	db.connection.Joins("left join roles on roles.user_id = users.id").
		Where("roles.rol_id IN ?", []int{constantvariables.NAVIGANTORS_ROL, constantvariables.PIONEERS_ROL, constantvariables.SECOND_COMMANDERS_ROL, constantvariables.SCOUTS_ROL}).
		Group("users.id").
		Table("users").Count(&count_commanders)
	defer entity.Closedb()
	return count_navigators, count_pioneers, count_followers, count_scouts, count_commanders
}

//FindUserNameLastName
func (db *userConnection) GetFindUserNameLastName(data string) ([]entity.User, error) {

	var user []entity.User

	go func() {
		err := db.connection.Preload("Roles.Rol").
			Preload("Church").
			Preload("SubDetachment").
			Where("concat(name,' ',last_name) LIKE ?", "%"+string(data)+"%").
			Find(&user).Error
		defer entity.Closedb()
		errChanUser <- err
	}()
	err := <-errChanUser

	if err == nil {
		return user, err
	}

	return user, err
}

//find FindUserNameLastName
