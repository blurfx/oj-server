package sq

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"testing"
	"time"
)

func TestSq(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")

	defer db.Close()

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	sq := NewDb(db)
	sq.Exec("CREATE TABLE user (`id` INT NOT NULL PRIMARY KEY, `name` varchar(255) NOT NULL, `verified` BOOLEAN NULL, `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP)")
	sq.Exec("INSERT INTO `user` (`id`, `name`, `created_at`) VALUES(1, 'User One', '1996-02-21 12:34')")
	sq.Exec("INSERT INTO `user` (`id`, `name`, `verified`) VALUES(2, '사용자 2', 1)")
	rows, err := sq.Query("SELECT * FROM user ORDER BY id")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	type User struct {
		Id         int          `db:"id"`
		Username   string       `db:"name"`
		IsVerified sql.NullBool `db:"verified"`
		CreatedAt  time.Time    `db:"created_at"`
	}

	for rows.Next() {
		var (
			id         int
			name       string
			isVerified sql.NullBool
			createdAt  time.Time
			user       User
		)
		err := rows.Scan(&id, &name, &isVerified, &createdAt)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
		err = rows.Scan(&id, &name, &isVerified, &createdAt)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
		err = rows.ScanStruct(&user)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
		t.Logf("[%10v] id: %v, name: %v, verified:%v(valid:%v), created_at: %v\n", "Scan", id, name, isVerified.Bool, isVerified.Valid, createdAt)
		t.Logf("[%10v] id: %v, name: %v, verified:%v, created_at:%v\n", "ScanStruct", user.Id, user.Username, user.IsVerified, user.CreatedAt)
		if id != user.Id || name != user.Username || isVerified.Bool != user.IsVerified.Bool || createdAt != user.CreatedAt {
			t.Errorf("failed to assertion.\nId: expected='%v' given='%v'\nName: expected='%s' given='%s'", id, user.Id, name, user.Username)
			t.FailNow()
		}
	}

	var (
		id         int
		name       string
		isVerified sql.NullBool
		createdAt  time.Time
		user       User
	)
	row := sq.QueryRow("SELECT * FROM user ORDER BY id")
	err = row.Scan(&id, &name, &isVerified, &createdAt)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	row = sq.QueryRow("SELECT * FROM user ORDER BY id")
	err = row.ScanStruct(&user)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("[%10v] id: %v, name: %v, verified:%v(valid:%v), created_at: %v\n", "Scan", id, name, isVerified.Bool, isVerified.Valid, createdAt)
	t.Logf("[%10v] id: %v, name: %v, verified:%v, created_at:%v\n", "ScanStruct", user.Id, user.Username, user.IsVerified, user.CreatedAt)
	if id != user.Id || name != user.Username || isVerified.Bool != user.IsVerified.Bool || createdAt != user.CreatedAt {
		t.Errorf("failed to assertion.\nId: expected='%v' given='%v'\nName: expected='%s' given='%s'", id, user.Id, name, user.Username)
		t.FailNow()
	}
}
