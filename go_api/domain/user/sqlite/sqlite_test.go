package sqlite_test

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nozomi-iida/nozo_blog/domain/user/sqlite"
	"github.com/nozomi-iida/nozo_blog/entity"
	"github.com/nozomi-iida/nozo_blog/test"
	"github.com/nozomi-iida/nozo_blog/valueobject"
)

func TestUserSqlite_Create(t *testing.T) {
	ts, err := test.ConnectDB(t)
	defer ts.Remove()
	sr, err := sqlite.New(ts.Filename)
	if err != nil {
		t.Error("sqlite new err:", err)
	}	
	ps, err := valueobject.NewPassword("password123")
	us, err := entity.NewUser("nozomi", ps)
	if err != nil {
		t.Error("newUser err:", err)
	}	
	_, err = sr.Create(us)
	if err != nil {
		t.Error("create err:", err)
	}	
}
