package sq

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"testing"
	"time"
)

type testUser struct {
	Id              int          `db:"id"`
	Username        string       `db:"name"`
	IsVerified      sql.NullBool `db:"verified"`
	CreatedAt       time.Time    `db:"created_at"`
	DeprecatedField int
}

func createTestDb() *DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	sq := NewDb(db)

	return sq
}

func TestQueryRow(t *testing.T) {
	db := createTestDb()
	defer db.Close()

	db.Exec("CREATE TABLE user (`id` INT NOT NULL PRIMARY KEY, `name` varchar(255) NOT NULL, `verified` BOOLEAN NULL, `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP)")
	db.Exec("INSERT INTO `user` (`id`, `name`, `created_at`) VALUES(1, 'User One', '1996-02-21 12:34')")
	db.Exec("INSERT INTO `user` (`id`, `name`, `verified`) VALUES(2, '사용자 2', 1)")

	var (
		id         int
		name       string
		isVerified sql.NullBool
		createdAt  time.Time
		user       testUser
	)

	row := db.QueryRow("SELECT * FROM user ORDER BY id")
	err := row.Scan(&id, &name, &isVerified, &createdAt)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	row = db.QueryRow("SELECT * FROM user ORDER BY id")
	err = row.ScanStruct(&user)
	if err != nil {
		t.Error(err)
		return
	}

	if id != user.Id || name != user.Username || isVerified.Bool != user.IsVerified.Bool || createdAt != user.CreatedAt {
		t.Logf("[%10v] id: %v, name: %v, verified:%v(valid:%v), created_at: %v\n", "Scan", id, name, isVerified.Bool, isVerified.Valid, createdAt)
		t.Logf("[%10v] id: %v, name: %v, verified:%v, created_at:%v\n", "ScanStruct", user.Id, user.Username, user.IsVerified, user.CreatedAt)
		t.Errorf("failed to assertion.\nId: expected='%v' given='%v'\nName: expected='%s' given='%s'", id, user.Id, name, user.Username)
		return
	}
}

func TestQuery(t *testing.T) {
	db := createTestDb()

	defer db.Close()

	t1, _ := time.Parse("2006-01-02 15:04", "1996-02-21 12:34")
	t2, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	db.Exec("CREATE TABLE user (`id` INT NOT NULL PRIMARY KEY, `name` varchar(255) NOT NULL, `verified` BOOLEAN NULL, `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP)")
	db.Exec("INSERT INTO `user` (`id`, `name`, `created_at`) VALUES(1, 'User One', '1996-02-21 12:34')")
	db.Exec("INSERT INTO `user` (`id`, `name`, `verified`) VALUES(2, '사용자 2', 1)")

	rows, err := db.Query("SELECT * FROM user ORDER BY id")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	expected := []testUser{
		{1, "User One", sql.NullBool{false, false}, t1, 0},
		{2, "사용자 2", sql.NullBool{true, true}, t2, 0},
	}

	for i := 0; rows.Next(); i++ {
		var user testUser

		err = rows.ScanStruct(&user)
		if err != nil {
			t.Error(err)
			return
		}

		if expected[i] != user {
			t.Error("failed to assertion.")
			t.Logf("[%10v] %v\n", "Expected", expected[i])
			t.Logf("[%10v] %v\n", "Actual", user)
			return
		}
	}

}
