package database

import "github.com/vds/Restraunt/pkg/models"

type Database interface {
	CreateUser(admin *models.Admin) error
	LogIn(cred *models.Credentials,isSuper int) error
}
