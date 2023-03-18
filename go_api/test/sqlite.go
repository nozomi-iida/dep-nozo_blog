package test

import (
	"database/sql"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/nozomi-iida/nozo_blog/domain/user/sqlite"
	"github.com/nozomi-iida/nozo_blog/entity"
	"github.com/nozomi-iida/nozo_blog/valueobject"
	migrate "github.com/rubenv/sql-migrate"
)

type TestSqlite struct {
	Filename string
	Remove   func() error
}

func ConnectDB(t *testing.T) TestSqlite {
	f, err := os.CreateTemp("", "go-sqlite3-test-")
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
	filename := f.Name()
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		t.Fatal("sql open error:", err)
	}
	// リファクタの余地あり
	_, b, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(b), "../../../")
	migrations := &migrate.FileMigrationSource{
		Dir: root + "/nozo_blog/go_api/db/migrations",
	}
	_, err = migrate.Exec(db, "sqlite3", migrations, migrate.Up)
	if err != nil {
		t.Fatal("migrate exec error:", err)
	}
	defer func() {
		err = db.Close()
		if err != nil {
			t.Fatal("db close error:", err)
		}
	}()

	return TestSqlite{Filename: filename, Remove: func() error {
		return os.Remove(filename)
	}}
}

type userOption func(*entity.User)

func SetUsername(username string) userOption {
	return func(user *entity.User) {
		user.Username = username
	}
}

func CreateUser(t *testing.T, fileName string, options ...userOption) entity.User {
	ps, err := valueobject.NewPassword("password123")
	user, err := entity.NewUser("nozomi", ps)
	for _, option := range options {
		option(&user)
	}
	sq, err := sqlite.New(fileName)
	_, err = sq.Create(user)
	if err != nil {
		t.Error("create user err:", err)
	}

	return user
}
