package repository

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/rkunihiro/gormgql/mock"
)

func TestPostRepo_FindByID(t *testing.T) {
	g, m := mock.NewGormDB(t)
	m.ExpectQuery(regexp.QuoteMeta("SELECT")).WithArgs(123).WillReturnRows(
		sqlmock.NewRows([]string{"id", "authorId", "title", "posted"}).
			AddRow(456, 789, "hoge", time.Now()),
	)
	r := NewPostRepository(g)
	post, err := r.FindByID(123)
	assert.Nil(t, err, "FindById return unexpected error %v", err)
	assert.Equal(t, post.ID, 456, "unexpected post.ID")
	assert.Equal(t, post.Title, "hoge", "unexpected post.Name")
	assert.Nil(t, m.ExpectationsWereMet())
}

func TestPostRepo_Find(t *testing.T) {
	g, m := mock.NewGormDB(t)
	m.ExpectQuery(regexp.QuoteMeta("SELECT")).WillReturnRows(
		sqlmock.NewRows([]string{"id", "authorId", "title", "posted"}).
			AddRow(1, 1, "foo", time.Now()).AddRow(2, 2, "bar", time.Now()),
	)
	r := NewPostRepository(g)
	posts, err := r.Find()
	assert.Nil(t, err, "Find return unexpected error %v", err)
	assert.Equal(t, len(posts), 2)
	assert.Nil(t, m.ExpectationsWereMet())
}
