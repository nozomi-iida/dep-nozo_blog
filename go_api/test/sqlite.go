package test

import (
	"database/sql"
	"os"
	"testing"

	migrate "github.com/rubenv/sql-migrate"
)

type TestSqlite struct {
	Filename string
	Remove func() error
}

func ConnectDB(t *testing.T) (TestSqlite, error) {
	
	f, err := os.CreateTemp("", "go-sqlite3-test-")
	if err != nil {
		t.Fatal(err)
	}
	// ファイルを閉じ、I/Oに使用できないようにする
	f.Close()
	filename := f.Name()
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return TestSqlite{}, err
	}	
	migrations := &migrate.FileMigrationSource{
		Dir: "../../../db/migrations",
	}
	_, err = migrate.Exec(db, "sqlite3", migrations, migrate.Up)
	if err != nil {
		return TestSqlite{}, err
	}	
	defer func ()  {
		err = db.Close()
		if err != nil {
			t.Error("db close error:", err)
		}
	}()
	
	return TestSqlite{Filename: filename, Remove: func() error {
		return os.Remove(filename)
	}}, nil
}
