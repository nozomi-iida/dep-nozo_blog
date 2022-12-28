package sqlite_test

import (
	"testing"

	"github.com/nozomi-iida/nozo_blog/domain/user/sqlite"
	"github.com/nozomi-iida/nozo_blog/entity"
	"github.com/nozomi-iida/nozo_blog/test"
	"github.com/nozomi-iida/nozo_blog/valueobject"
)

// TODO: testCaseを増やす
func TestUserSqlite_FindById(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	sq, err := sqlite.New(ts.Filename)
	ps, err := valueobject.NewPassword("password123")
	us, err := entity.NewUser("nozomi", ps)
	_, err = sq.Create(us)
	if err != nil {
		t.Error("create user err:", err)
	}	
	rs, err := sq.FindById(us.GetID())
	if rs.GetID() != us.GetID() {
		t.Error("find by username error", err)
	}
}

func TestUserSqlite_FindByUsername(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	ps, err := valueobject.NewPassword("password123")
	un := "nozomi"
	sq, err := sqlite.New(ts.Filename)
	us, err := entity.NewUser(un, ps)
	_, err = sq.Create(us)
	if err != nil {
		t.Error("create user err:", err)
	}	
	rs, err := sq.FindByUsername(un)
	if rs.GetUsername() != us.GetUsername() {
		t.Error("find by username error", err)
	}
}

func TestUserSqlite_Create(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	sq, err := sqlite.New(ts.Filename)
	if err != nil {
		t.Error("sqlite new err:", err)
	}	
	ps, err := valueobject.NewPassword("password123")
	us, err := entity.NewUser("nozomi", ps)
	if err != nil {
		t.Error("newUser err:", err)
	}	
	_, err = sq.Create(us)
	if err != nil {
		t.Error("create err:", err)
	}	
}
