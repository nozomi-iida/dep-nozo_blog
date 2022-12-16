package domain_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/rubenv/sql-migrate"
)

func TempFilename(t testing.TB) string  {
	f, err := os.CreateTemp("", "go-sqlite3-test-")
	if err != nil {
		t.Fatal(err)
	}
	// ファイルを閉じ、I/Oに使用できないようにする
	f.Close()
	return f.Name()
}

func DoTestOpen(t *testing.T, option string) (string, error)  {
	tempFilename := TempFilename(t)	
	url := tempFilename + option

	defer func ()  {
		err := os.Remove(tempFilename)
		if err != nil {
			t.Error("temp file remove error:", err)
		}
	}()

	db, err := sql.Open("sqlite3", url)
	if err != nil {
		return "failed to open database:", err
	}

	defer func ()  {
		err = db.Close()
		if err != nil {
			t.Error("db close error:", err)
		}
	}()

	err = db.Ping()
	if err != nil {
		return "ping error:", err
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "db/migrations",
	}
	
	_, err = migrate.Exec(db, "sqlite3", migrations, migrate.Up)
	if err != nil {
		return "migrate error:", err
	}

	if stat, err := os.Stat(tempFilename); err != nil || stat.IsDir() {
		return "failed to create db file", nil
	}

	return "", nil
}

func TestOpen(t *testing.T)  {
	// トランザクションのロックの解除方法を指定するオプションらしいw
	cases := map[string]bool{
		"":                   true,
		"?_txlock=immediate": true,
		"?_txlock=deferred":  true,
		"?_txlock=exclusive": true,
		"?_txlock=bogus":     false,
	}
	for option, expectedPass := range cases {
		result, err := DoTestOpen(t, option)
		if result == "" {
			if !expectedPass {
				errmsg := fmt.Sprintf("_txlock error not caught at dbOpen with option %s", option)
				t.Fatal(errmsg)
			} else if expectedPass {
				if err == nil {
					t.Fatal(result)
				} else {
					t.Fatal(result, err)
				}
			}
		}
	}
}
