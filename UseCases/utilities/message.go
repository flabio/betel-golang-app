package utilities

type MessageRequired struct {
	Name string
}
type GetMessageRequired interface {
	RequiredName() string
}

func (msg MessageRequired) RequiredId() string {
	return "The id is required"
}
func (msg MessageRequired) RequiredUserId() string {
	return "The id user is required"
}

func (msg MessageRequired) RequiredSubDetachmentId() string {
	return "The id of subdetachment is required"
}

func (msg MessageRequired) RequiredName() string {
	return "The name is required"
}

func (msg MessageRequired) RequiredLastName() string {
	return "The last name is required"
}
func (msg MessageRequired) RequiredIdentification() string {
	return "The identification is required"
}
func (msg MessageRequired) RequiredExistIdentification() string {
	return "The identification already exists"
}
func (msg MessageRequired) RequiredTypeIdentification() string {
	return "The type identification is required"
}
func (msg MessageRequired) RequiredBirthplace() string {
	return "The birthplace is required"
}
func (msg MessageRequired) RequiredGender() string {
	return "The gender is required"
}
func (msg MessageRequired) RequiredCivilStatus() string {
	return "The civilstatus is required"
}
func (msg MessageRequired) RequiredPosition() string {
	return "The position is required"
}

func (msg MessageRequired) RequiredOccupation() string {
	return "The occupation is required"
}
func (msg MessageRequired) RequiredEmail() string {
	return "The email is required"
}
func (msg MessageRequired) RequiredExistEmail() string {
	return "The email already exists"
}

func (msg MessageRequired) RequiredRol() string {
	return "The rol is required"
}
func (msg MessageRequired) RequiredDetachment() string {
	return "The detachment is required"
}
func (msg MessageRequired) RequiredSubDetachment() string {
	return "The sub detachment is required"
}
func (msg MessageRequired) RequiredChurch() string {
	return "The church is required"
}

func (msg MessageRequired) RequiredPassword() string {
	return "The password is required"
}
func (msg MessageRequired) RequiredPasswordConfirm() string {
	return "The passwords are not the same"
}
func (msg MessageRequired) RequiredBirthday() string {
	return "Please check the date of birth"
}
func (msg MessageRequired) RequiredWeek() string {
	return "The week is required"
}
func (msg MessageRequired) ExtisByUserWeek() string {
	return "You already have this week assigned"
}
func (msg MessageRequired) RequiredState() string {
	return "The state is required"
}
func (msg MessageRequired) RequiredDescription() string {
	return "The description is required"
}
