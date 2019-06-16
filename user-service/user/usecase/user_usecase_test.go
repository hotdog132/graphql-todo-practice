package usecase_test

import (
	"database/sql"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	"github.com/hotdog132/graphql-todo-practice/user-service/models"
	"github.com/hotdog132/graphql-todo-practice/user-service/user"
	"github.com/hotdog132/graphql-todo-practice/user-service/user/repository"
	"github.com/hotdog132/graphql-todo-practice/user-service/user/usecase"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	mock     sqlmock.Sqlmock
	usercase user.Usecase
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	DB, err := gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	DB.LogMode(true)

	r := repository.NewPsqlUserRepository(DB)
	s.usercase = usecase.NewUserUsecase(r)
}

func (s *Suite) TestFetch() {

	queryID := 1
	expectUserName := "user-1"

	s.mock.ExpectQuery(fmt.Sprintf(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE ("users"."id" = %d) ORDER BY "users"."id" ASC LIMIT 1`),
		queryID)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(queryID, expectUserName))

	u, err := s.usercase.Fetch(queryID)

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&models.User{ID: queryID, Name: expectUserName}, u))
}

func (s *Suite) TestFetchAll() {
	queryID1 := 1
	expectUserName1 := "user-1"
	queryID2 := 2
	expectUserName2 := "user-2"

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(queryID1, expectUserName1).AddRow(queryID2, expectUserName2))

	users, err := s.usercase.FetchAll()

	require.NoError(s.T(), err)
	require.Equal(s.T(), 2, len(users))
}

// func (s *Suite) TestStore() {
// 	expectUserName := "user-1"

// 	s.mock.ExpectExec("INSERT INTO users").
// 		WithArgs("expectUserName", time.Now(), time.Now())

// 	u := &models.User{Name: expectUserName}
// 	err := s.usercase.Store(u)

// 	require.NoError(s.T(), err)
// }

func TestInit(t *testing.T) {
	s := new(Suite)
	suite.Run(t, s)
}
