package app

import (
	"time"

	"gorm.io/gorm"
)

var DB *gorm.DB
var KeyAccess string
var KeyRefresh string
var ExpiryAgeAccess time.Duration
var ExpiryAgeRefresh time.Duration
