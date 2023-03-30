package postgresql_repository_test

import (
	"database/sql/driver"
	"fmt"
	entities "mygram/domains/entity"
	repository "mygram/infrastructures/repository/postgres"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}


func TestUserRepository_SuccessGetUserByEmail(t *testing.T){

	db, mock, err := sqlmock.New()
	if err != nil {
		// t.Fatal("error creating mock database: %v",err)
		fmt.Println(err)
	}

	gormDB,err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}),&gorm.Config{})

	if err != nil {
        t.Fatalf("error opening GORM database: %v", err)
    }

	user:=entities.User{
		Id:        uuid.New(),
		UserName:  "Adit",
		Email:     "adit@mail.com",
		Password:  "12345",
		Age:       21,
	}

	userRepository := repository.NewUserRepositoryPostgres(gormDB)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "user" WHERE "email" = $1 ORDER BY "user"."id" LIMIT 1`)). //parsing the string as regex and requires escape characters gunakan QuoteMeta
	WithArgs("adit@mail.com").
	WillReturnRows(sqlmock.
		NewRows([]string{"id","username","email","age","password"}).
		AddRow(user.Id.String(),user.UserName,user.Email,user.Age,user.Password),
	)
	result,err:=userRepository.GetUserByEmail("adit@mail.com")
	assert.NotNil(t,result)
	assert.Nil(t,err)
}

func TestUserRepository_NotFoundGetUserByEmail(t *testing.T){
	
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println(err)
	}

	gormDB,err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}),&gorm.Config{})

	if err != nil {
        t.Fatalf("error opening GORM database: %v", err)
    }

	userRepository := repository.NewUserRepositoryPostgres(gormDB)
	// HANDLE ROW NOT FOUND

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "user" WHERE "email" = $1 ORDER BY "user"."id" LIMIT 1`)).
	WithArgs("adit@mail.com").
	WillReturnRows(sqlmock.
		NewRows([]string{"id"}), // guanakan fungsi ini
	)
	result,err:=userRepository.GetUserByEmail("adit@mail.com")
	assert.Nil(t,result)
	assert.NotNil(t,err)
}

/*
Test Insert User

*/

func TestUserRepository_InsertSuccess(t *testing.T){
	// var (
	// 	id = uuid.New()
	// 	username ="adit"
	// 	password = "12345"
	// 	email = "aditya@mail.com"
	// 	age = 25
	// 	createdAt = time.Now()
	// 	updatedAt = time.Now()
	// )
	// var user entities.User

	user:= &entities.User{
		Id:        uuid.New(),
		UserName:  "adit",
		Email:     "aditya@mail.com",
		Password:  "12345",
		Age:       26,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db,mock,err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB,err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}),&gorm.Config{})

	if err != nil {
        t.Fatalf("error opening GORM database: %v", err)
    }

	userRepository := repository.NewUserRepositoryPostgres(gormDB)
	// query := "INSERT INTO USER (id,username,password,email,age,created_at,updated_at) VALUES($1,$2,$3,$4,$5,$6,$7)"
	query := `INSERT INTO "user" ("id","username","email","password","age","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6,$7)`
	// query := `INSERT INTO "user" \("id","username","email","password","age","created_at","updated_at"\) VALUES\(\$1,\$2,\$3,\$4,\$5,\$6,\$7\)`
	// mock.ExpectBegin()
	// mock.ExpectQuery(regexp.QuoteMeta(query)).
	// // WithArgs(id,username,password,email,age,createdAt,updatedAt).
	// WithArgs(user.Id.String(),user.UserName,user.Email,user.Password,user.Age,user.CreatedAt,user.UpdatedAt).
	// WillReturnRows(
	// 	sqlmock.NewRows([]string{"id"}),
	// )
	// mock.ExpectCommit()
	mock.ExpectBegin()
	// prep := mock.ExpectPrepare(regexp.QuoteMeta(query))
	mock.ExpectExec(regexp.QuoteMeta(query)).
	// WithArgs(id,username,password,email,age,createdAt,updatedAt).
	WithArgs(&user.Id,user.UserName,user.Email,user.Password,user.Age,AnyTime{},AnyTime{}).
	WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()
	err = userRepository.Insert(user)
	assert.Nil(t,err)
}