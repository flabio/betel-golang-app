package repositorys

import (
	"sync"

	"gorm.io/gorm"
)

type OpenConnections struct {
	connection *gorm.DB
	mux        sync.Mutex
}
