package snippetrepo

import "gorm.io/gorm"

type Repo struct {
	DbClient *gorm.DB
}
