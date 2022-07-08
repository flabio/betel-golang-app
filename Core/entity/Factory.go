package entity

import "gorm.io/gorm"

func Factory(t int) *gorm.DB {
	switch t {
	case 1:
		return DatabaseConnection()
	case 2:
		return DatabaseConnection()
	default:
		return nil

	}

}
