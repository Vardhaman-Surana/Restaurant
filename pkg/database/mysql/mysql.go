package mysql

import (
	"database/sql"
	"errors"
	"github.com/vds/Restraunt/pkg/encryption"
	"github.com/vds/Restraunt/pkg/models"
	"log"
	_ "github.com/go-sql-driver/mysql"

)

var errEmptyFields = errors.New("Empty fields please fill the fields and try again")
var errInvalidCredentials = errors.New("Incorrect login details")
type MySqlDB struct {
	*sql.DB
}

func NewMySqlDB() (*MySqlDB, error) {
	db,err:=sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/restaurant")
	if err!=nil{
		return nil,err
	}
	mySqlDB:=&MySqlDB{db}
	return mySqlDB,err
}

func (db *MySqlDB) CreateUser(admin *models.Admin) error {
	stmt,err:=db.Prepare("insert into admins(name,password,is_super) values(?,?,?)")
	if err!=nil{
		log.Printf("%v",err)
		return err
	}
	pass,err:=encryption.GenerateHash(admin.Password)
	if err!=nil{
		log.Printf("%v",err)
		return err
	}
	_,err=stmt.Exec(admin.Name,pass,admin.IsSuper)
	if err!=nil{
		log.Printf("%v",err)
		return err
	}
	return nil
}

func (db *MySqlDB) LogIn(cred *models.Credentials,isSuper int) error{
	passHash:=""
	isSuperDB:=0
	if cred.UserName=="" || cred.Password==""{
		return errEmptyFields
	}
	rows,err:=db.Query("select is_super,password from admins where name=?",cred.UserName)
	if err!=nil{
		log.Printf("%v",err)
		return errInvalidCredentials
	}
	rows.Next()
	err=rows.Scan(&isSuperDB,&passHash)
	if err!=nil{
		log.Printf("%v",err)
		return err
	}
	isValid:=encryption.ComparePasswords(passHash,cred.Password)
	if !isValid{
		return errInvalidCredentials
	}
	if isSuperDB!=isSuper{
		return errInvalidCredentials
	}
	return nil
}