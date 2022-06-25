package mock

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewGormDB create gorm.DB mock
func NewGormDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	d, m, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	if err != nil {
		t.Fatalf("sqlmock.New failed %v", err)
	}
	g, err := gorm.Open(
		mysql.New(mysql.Config{Conn: d, SkipInitializeWithVersion: true}),
		&gorm.Config{},
	)
	if err != nil {
		t.Fatalf("gorm.Open failed %v", err)
	}
	return g, m
}
