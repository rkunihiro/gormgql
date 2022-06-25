package repository

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/rkunihiro/gormgql/mock"
)

func TestUserRepo_FindByID(t *testing.T) {
	g, m := mock.NewGormDB(t)
	m.ExpectQuery(regexp.QuoteMeta("SELECT")).WithArgs(123).WillReturnRows(
		sqlmock.NewRows([]string{"id", "name"}).AddRow(456, "hoge"),
	)
	r := NewUserRepository(g)
	user, err := r.FindByID(123)
	assert.Nil(t, err, "FindById return unexpected error %v", err)
	assert.Equal(t, user.ID, 456, "unexpected user.ID")
	assert.Equal(t, user.Name, "hoge", "unexpected user.Name")
	assert.Nil(t, m.ExpectationsWereMet())
}

func TestUserRepo_Find(t *testing.T) {
	g, m := mock.NewGormDB(t)
	m.ExpectQuery(regexp.QuoteMeta("SELECT")).WillReturnRows(
		sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "foo").AddRow(2, "bar"),
	)
	r := NewUserRepository(g)
	users, err := r.Find()
	assert.Nil(t, err, "Find return unexpected error %v", err)
	assert.Equal(t, len(users), 2)
	assert.Nil(t, m.ExpectationsWereMet())
}
