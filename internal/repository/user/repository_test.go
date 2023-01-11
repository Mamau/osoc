package user

import (
	"context"
	"database/sql"
	"fmt"
	"osoc/internal/entity"
	"osoc/pkg/fixture"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/assert"
	"osoc/pkg/mysql"
)

// Fixture for positive cases
func setupDBWithUsers(t testing.TB, db *mysql.DB) func(tb testing.TB) {
	dbName := strings.ReplaceAll(uuid.New().String(), "-", "_")
	db.MustExec(fmt.Sprintf("CREATE DATABASE %s", dbName))
	db.MustExec(fmt.Sprintf("USE %s", dbName))
	if err := makeMigrate(db.DB.DB); err != nil {
		t.Fatal(err)
	}
	db.MustExec(`INSERT INTO users (id, name, age, social, created_at) VALUES (1, 'fixture', 20, 'soc', '2022-11-14 10:00:00'),
       (2, 'test2', 20, 'soc', '2022-11-14 10:00:00'),
       (3, 'test3', 20, 'soc', '2022-11-14 10:00:00')`)

	return func(tb testing.TB) {
		db.MustExec(fmt.Sprintf("DROP DATABASE %s", dbName))
	}
}

// Fixture for negative cases
func setupEmptyDB(t testing.TB, db *mysql.DB) func(tb testing.TB) {
	dbName := strings.ReplaceAll(uuid.New().String(), "-", "_")
	db.MustExec(fmt.Sprintf("CREATE DATABASE %s", dbName))
	db.MustExec(fmt.Sprintf("USE %s", dbName))
	if err := makeMigrate(db.DB.DB); err != nil {
		t.Fatal(err)
	}
	return func(tb testing.TB) {
		db.MustExec(fmt.Sprintf("DROP DATABASE %s", dbName))
	}
}

func TestGetUser(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	db, teardownSuite := fixture.SetupMysqlDB(t)
	defer teardownSuite(t)

	cases := []struct {
		Name      string
		SetupFunc func(tb testing.TB, db *mysql.DB) func(tb testing.TB)
		ExpErr    error
		Expected  entity.User
	}{
		{
			Name:      "Success get user",
			SetupFunc: setupDBWithUsers,
			ExpErr:    nil,
			Expected: entity.User{
				ID:        1,
				Age:       20,
				Name:      "fixture",
				Social:    "soc",
				CreatedAt: time.Date(2022, time.November, 14, 10, 0, 0, 0, time.UTC)},
		},
		{
			Name:      "User not found",
			SetupFunc: setupEmptyDB,
			ExpErr:    entity.ErrNotFound,
			Expected:  entity.User{},
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			teardownTest := c.SetupFunc(t, db)
			defer teardownTest(t)

			repo := New(db)
			actual, err := repo.GetUser(context.Background(), 1)
			assert.Equal(t, c.ExpErr, err)
			assert.Equal(t, c.Expected, actual)
		})
	}
}

func TestGetUsers(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	db, teardownSuite := fixture.SetupMysqlDB(t)
	defer teardownSuite(t)

	cases := []struct {
		Name      string
		SetupFunc func(tb testing.TB, db *mysql.DB) func(tb testing.TB)
		ExpErr    error
		Expected  []entity.User
	}{
		{
			Name:      "Success get user",
			SetupFunc: setupDBWithUsers,
			ExpErr:    nil,
			Expected: []entity.User{
				{
					ID:        1,
					Age:       20,
					Name:      "fixture",
					Social:    "soc",
					CreatedAt: time.Date(2022, time.November, 14, 10, 0, 0, 0, time.UTC),
				},
				{
					ID:        2,
					Age:       20,
					Name:      "test2",
					Social:    "soc",
					CreatedAt: time.Date(2022, time.November, 14, 10, 0, 0, 0, time.UTC),
				},
				{
					ID:        3,
					Age:       20,
					Name:      "test3",
					Social:    "soc",
					CreatedAt: time.Date(2022, time.November, 14, 10, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			Name:      "User not found",
			SetupFunc: setupEmptyDB,
			ExpErr:    entity.ErrNotFound,
			Expected:  []entity.User(nil),
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			teardownTest := c.SetupFunc(t, db)
			defer teardownTest(t)

			repo := New(db)
			actual, err := repo.GetUsers(context.Background())
			assert.Equal(t, c.ExpErr, err)
			assert.Equal(t, c.Expected, actual)
		})
	}
}

func makeMigrate(db *sql.DB) error {
	if err := goose.SetDialect("mysql"); err != nil {
		return err
	}
	if err := goose.UpTo(db,
		"../migrations",
		20211015215408); err != nil {
		return err
	}
	return nil
}
