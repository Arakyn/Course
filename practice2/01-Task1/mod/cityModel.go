package mod

import "gorm.io/gorm"

type State struct {
	gorm.Model
	City    string
	Country string
}
