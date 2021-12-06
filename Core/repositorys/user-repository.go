package repositorys

import (
	"bete/Core/entity"

	"gorm.io/gorm"
)

//UserRepository is contract what UserRepository can do to db
type UserRepository interface {
	InsertUser(user entity.User) (entity.User, error)

	EditUser(user entity.User) (entity.User, error)
	InsertRole(role entity.Role) error
	EditRole(role entity.Role) (entity.Role, error)
	InsertGroup(group entity.UserSubdetachement) error

	EditGroup(group entity.UserSubdetachement) (entity.UserSubdetachement, error)

	AllUser() ([]entity.User, error)

	PaginationUsers(begin, limit int) ([]entity.User, error)
	DeleteUser(id uint) (bool, error)
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (bool, error)
	FindUserNameLastName(data string) ([]entity.User, error)
	FindByEmail(email string) (entity.User, error)
	IsDuplicateIdentificatio(identification string) bool
	ProfileUser(userId uint) (entity.User, error)
	ChangePassword(user entity.User) error
	DeleteRoleUser(id uint) error

	CountUser() int64
	ListNavigators() ([]entity.User, error)
	ListPioneers() ([]entity.User, error)
	ListFollowersWays() ([]entity.User, error)
	ListScouts() ([]entity.User, error)
	ListKingsScouts(Id uint) ([]entity.User, error)
	ListCommanders() ([]entity.User, error)
	ListMajors() ([]entity.User, error)
	CounKanban() (int64, int64, int64, int64, int64)
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

func (db *userConnection) InsertUser(user entity.User) (entity.User, error) {
	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Save(&user).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan

	return user, err
}

func (db *userConnection) EditUser(user entity.User) (entity.User, error) {

	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Save(&user).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return user, err
}

func (db *userConnection) InsertRole(role entity.Role) error {

	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Save(&role).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return err
}
func (db *userConnection) InsertGroup(group entity.UserSubdetachement) error {

	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Save(&group).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return err
}

func (db *userConnection) EditRole(role entity.Role) (entity.Role, error) {
	var rol entity.Role
	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Where("user_id =?", role.UserId).Updates(&role).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return rol, err
}

func (db *userConnection) EditGroup(gruop entity.UserSubdetachement) (entity.UserSubdetachement, error) {
	var gruo entity.UserSubdetachement
	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Where("user_id =?", gruop.UserId).Updates(&gruop).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return gruo, err
}

func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user entity.User
	var errChan = make(chan error, 1)

	go func() {
		err := db.connection.Preload("SubDetachment").Preload("Roles.Rol.RoleModule.Module").Select("users.id,users.name,users.last_name,users.email,users.image,users.sub_detachment_id,users.church_id,roles.id as idrol").Joins("left join roles on roles.user_id = users.id").Where("email = ?", email).Find(&user).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	if err == nil {
		return user
	}
	return nil
}

func (db *userConnection) IsDuplicateEmail(email string) (bool, error) {
	var user entity.User
	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Where("email = ?", email).Take(&user).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	if err == nil {
		return true, err
	}
	return false, err
}
func (db *userConnection) IsDuplicateIdentificatio(identification string) bool {
	var user entity.User
	var errChan = make(chan error, 1)

	go func() {
		err := db.connection.Where("identification = ?", identification).Take(&user).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan

	if err == nil {
		return true
	}
	return false
}

func (db *userConnection) DeleteUser(id uint) (bool, error) {
	var user entity.User
	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Where("id=?", id).Delete(&user).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	if err == nil {
		return true, err
	}
	return false, err
}
func (db *userConnection) AllUser() ([]entity.User, error) {
	var user []entity.User
	var errChan = make(chan error, 1)

	go func() {

		err := db.connection.Preload("Roles.Rol").Preload("Church").Preload("SubDetachment").Find(&user).Error
		defer entity.Closedb()
		errChan <- err
	}()

	err := <-errChan

	return user, err
}

func (db *userConnection) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Preload("Roles.Rol").Where("email = ?", email).Take(&user).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	if err == nil {
		return user, err
	}
	return user, err
}

func (db *userConnection) ProfileUser(userID uint) (entity.User, error) {
	var user entity.User
	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Preload("Roles.Rol").Preload("Roles.StudyCarried").Preload("MinisterialAcademys").Preload("Church").Find(&user, userID).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	if err == nil {
		return user, err
	}
	return user, err
}

func (db *userConnection) DeleteRoleUser(id uint) error {
	var role entity.Role
	var errChan = make(chan error, 1)

	go func() {
		err := db.connection.Where("user_id=?", id).Delete(&role).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return err
}

func (db *userConnection) CountUser() int64 {
	var count int64

	var errChan = make(chan error, 1)

	go func() {
		err := db.connection.Table("users").Count(&count).Error
		defer entity.Closedb()
		errChan <- err
	}()
	<-errChan
	return count
}
func (db *userConnection) PaginationUsers(begin, limit int) ([]entity.User, error) {
	var user []entity.User
	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Offset(begin).Limit(limit).Order("id desc").Preload("Roles.Rol").Preload("Church").Preload("SubDetachment").Find(&user).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return user, err

}
func (db *userConnection) ChangePassword(user entity.User) error {
	var errChan = make(chan error, 1)
	go func() {
		err := db.connection.Where("id =?", user.Id).Update("password", user.Password).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return err
}

//ListNavigators
func (db *userConnection) ListNavigators() ([]entity.User, error) {
	var user []entity.User
	var errChan = make(chan error, 1)

	go func() {
		err := db.connection.Joins("left join roles on roles.user_id = users.id").Joins("left join sub_detachments on users.sub_detachment_id=sub_detachments.id").Where("sub_detachments.id", 1).Where("roles.rol_id", 29).Order("users.id desc").Group("users.id").Preload("Roles.Rol").Find(&user).Error
		defer entity.Closedb()
		errChan <- err
		close(errChan)
	}()
	err := <-errChan

	return user, err
}

//fin ListNavigators
//ListPioneers
func (db *userConnection) ListPioneers() ([]entity.User, error) {
	var user []entity.User
	var errChan = make(chan error, 1)

	go func() {
		err := db.connection.Joins("left join roles on roles.user_id = users.id").Joins("left join sub_detachments on users.sub_detachment_id=sub_detachments.id").Where("sub_detachments.id", 2).Where("roles.rol_id", 29).Order("users.id desc").Group("users.id").Preload("Roles.Rol").Find(&user).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return user, err
}

//fin ListPioneers
//ListFollowersWays
func (db *userConnection) ListFollowersWays() ([]entity.User, error) {
	var user []entity.User
	var errChan = make(chan error, 1)

	go func() {
		err := db.connection.Joins("left join roles on roles.user_id = users.id").Joins("left join sub_detachments on users.sub_detachment_id=sub_detachments.id").Where("sub_detachments.id", 3).Where("roles.rol_id", 29).Order("users.id desc").Group("users.id").Preload("Roles.Rol").Find(&user).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return user, err
}

//fin ListFollowersWays
//ListScouts
func (db *userConnection) ListScouts() ([]entity.User, error) {
	var user []entity.User
	var errChan = make(chan error, 1)

	go func() {
		err := db.connection.Joins("left join roles on roles.user_id = users.id").Joins("left join sub_detachments on users.sub_detachment_id=sub_detachments.id").Where("sub_detachments.id", 4).Where("roles.rol_id", 29).Order("users.id desc").Group("users.id").Preload("Roles.Rol").Find(&user).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return user, err
}

//fin ListScouts
//ListCommanders
func (db *userConnection) ListCommanders() ([]entity.User, error) {
	var user []entity.User
	var errChan = make(chan error, 1)

	go func() {
		err := db.connection.Joins("left join roles on roles.user_id = users.id").Order("users.id desc").Where("roles.rol_id IN ?", []int{2, 3, 4, 5}).Group("users.id").Preload("Roles.Rol").Find(&user).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return user, err
}

//fin ListCommanders
//ListMajors
func (db *userConnection) ListMajors() ([]entity.User, error) {
	var user []entity.User
	var errChan = make(chan error, 1)

	go func() {
		err := db.connection.Joins("left join roles on roles.user_id = users.id").Distinct("users.id").Order("users.id desc").Where("roles.rol_id IN ?", []int{2, 3}).Preload("Roles.Rol").Preload("Church").Preload("SubDetachment").Find(&user).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return user, err
}

// fin ListMajors

func (db *userConnection) ListKingsScouts(Id uint) ([]entity.User, error) {
	var user []entity.User
	var errChan = make(chan error, 1)

	go func() {
		err := db.connection.Preload("Roles.Rol").Joins("left join roles on roles.user_id = users.id").Joins("left join sub_detachments on users.sub_detachment_id=sub_detachments.id").Where("sub_detachments.id", Id).Where("roles.rol_id", 29).Group("users.id").Find(&user).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan
	return user, err
}

//count scouts of king

func (db *userConnection) CounKanban() (int64, int64, int64, int64, int64) {
	var count_navigators int64
	var count_pioneers int64
	var count_followers int64
	var count_scouts int64
	var count_commanders int64

	db.connection.Joins("left join roles on roles.user_id = users.id").Joins("left join sub_detachments on users.sub_detachment_id=sub_detachments.id").Where("sub_detachments.id", 1).Where("roles.rol_id", 29).Group("users.id").Table("users").Count(&count_navigators)
	db.connection.Joins("left join roles on roles.user_id = users.id").Joins("left join sub_detachments on users.sub_detachment_id=sub_detachments.id").Where("sub_detachments.id", 2).Where("roles.rol_id", 29).Group("users.id").Table("users").Count(&count_pioneers)

	db.connection.Joins("left join roles on roles.user_id = users.id").Joins("left join sub_detachments on users.sub_detachment_id=sub_detachments.id").Where("sub_detachments.id", 3).Where("roles.rol_id", 29).Group("users.id").Table("users").Count(&count_followers)
	db.connection.Joins("left join roles on roles.user_id = users.id").Joins("left join sub_detachments on users.sub_detachment_id=sub_detachments.id").Where("sub_detachments.id", 4).Where("roles.rol_id", 29).Group("users.id").Table("users").Count(&count_scouts)
	//db.connection.Joins("left join roles on roles.user_id = users.id").Joins("left join sub_detachments on users.sub_detachment_id=sub_detachments.id").Where("sub_detachments.id", 2).Where("roles.rol_id", 29).Group("users.id").Table("users").Count(&count_commanders)

	db.connection.Joins("left join roles on roles.user_id = users.id").Where("roles.rol_id IN ?", []int{2, 3, 4, 5}).Group("users.id").Table("users").Count(&count_commanders)
	defer entity.Closedb()
	return count_navigators, count_pioneers, count_followers, count_scouts, count_commanders
}

//FindUserNameLastName
func (db *userConnection) FindUserNameLastName(data string) ([]entity.User, error) {

	var user []entity.User
	var errChan = make(chan error, 1)

	go func() {
		err := db.connection.Preload("Roles.Rol").Preload("Church").Preload("SubDetachment").Where("concat(name,' ',last_name) LIKE ?", "%"+string(data)+"%").Find(&user).Error
		defer entity.Closedb()
		errChan <- err
	}()
	err := <-errChan

	if err == nil {
		return user, err
	}

	return user, err
}

//find FindUserNameLastName
