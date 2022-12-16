package sqlite_test

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nozomi-iida/nozo_blog/domain/user/sqlite"
	"github.com/nozomi-iida/nozo_blog/entity"
	"github.com/nozomi-iida/nozo_blog/test"
)

func TestUserSqlite_Create(t *testing.T) {
	ts, err := test.ConnectDB(t)
	defer ts.Remove()
	sr, err := sqlite.New(ts.Filename)
	if err != nil {
		t.Error("sqlite new err:", err)
	}	
	us, err := entity.NewUser("nozomi", "password123")
	if err != nil {
		t.Error("newUser err:", err)
	}	
	_, err = sr.Create(us)
	if err != nil {
		t.Error("create err:", err)
	}	
}
