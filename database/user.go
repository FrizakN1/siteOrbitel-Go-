package database

import (
	"database/sql"
	"errors"
	"newSite/utils"
	"time"
)

type Session struct {
	Hash   string
	User   User
	Date   string
	Exists bool
}

type User struct {
	ID             int
	Name           string
	Phone          int64
	AccountNumber  string
	Password       string
	CurrentBalance float64
	CurrentTariff  *Tariff
	Role           Role
}

type Role struct {
	ID   int
	Name string
}

var request map[string]*sql.Stmt
var sessionMap map[string]Session

func prepareUser() []string {
	sessionMap = make(map[string]Session)
	request = make(map[string]*sql.Stmt)
	errors := make([]string, 0)
	var e error

	//request["SessionSelect"], e = DB.Prepare(`-- SELECT "hash", "id", "name", "phone", "account_number", "current_balance", "current_tariff", "role", "date" FROM "Session" AS s INNER JOIN "User" AS u ON u."id" = s."user"`)
	request["SessionSelect"], e = DB.Prepare(`
		SELECT s.hash, u.id, u.name, u.phone, u.account_number, u.current_balance, t.id, t.name, r.id, r.name, s.date
		FROM "Session" as s
		JOIN "User" as u ON s.user_id = u.id
		JOIN "Role" as r ON u.role = r.id
		JOIN "Tariff" as t ON u.current_tariff = t.id`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["SessionInsert"], e = DB.Prepare(`INSERT INTO "Session" ("hash", "user_id", "date") VALUES ($1, $2, CURRENT_TIMESTAMP)`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["SessionDelete"], e = DB.Prepare(`DELETE FROM "Session" WHERE "hash" = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["GetAdmin"], e = DB.Prepare(`SELECT "id" FROM "User" WHERE "role" = 1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["CreateAdmin"], e = DB.Prepare(`INSERT INTO "User" ("name", "phone", "account_number", "password", "current_balance", "current_tariff", "role") VALUES ($1,$2,$3,$4,$5,$6,$7)`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["LoginCheck"], e = DB.Prepare(`
		SELECT u."id", u."name", u."phone", u."account_number", u."current_balance", u."current_tariff", r."id", r."name"
		FROM "User" as u
		JOIN "Role" as r ON u.role = r.id
		WHERE u.account_number = $1 AND u.password = $2`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	return errors
}

func (user *User) UserAuthorizationCheck() bool {
	stmt, ok := request["LoginCheck"]
	if !ok {
		utils.Logger.Println("Запрос не подготовлен")
	}

	row := stmt.QueryRow(user.AccountNumber, user.Password)
	e := row.Scan(&user.ID, &user.Name, &user.Phone, &user.AccountNumber, &user.CurrentBalance, &user.CurrentTariff, &user.Role.ID, &user.Role.Name)
	if e != nil {
		utils.Logger.Println(e)
		return false
	}

	return true
}

func GetSession(hash string) *Session {
	session, ok := sessionMap[hash]
	if ok {
		return &session
	}

	return nil
}

func DeleteSession(s *Session) {
	stmt, ok := request["SessionDelete"]
	if !ok {
		return
	}

	_, e := stmt.Exec(s.Hash)
	if e != nil {
		utils.Logger.Println(e)
	}

	return
}

func CreateSession(user *User) (string, bool) {
	stmt, ok := request["SessionInsert"]
	if !ok {
		return "", false
	}

	hash, e := utils.GenerateHash(user.AccountNumber)
	if e != nil {
		utils.Logger.Println(e)
		return "", false
	}

	_, e = stmt.Exec(hash, user.ID)
	if e != nil {
		utils.Logger.Println(e)
		return "", false
	}

	if sessionMap != nil {
		sessionMap[hash] = Session{
			Hash: hash,
			User: User{
				ID:             user.ID,
				Name:           user.Name,
				Phone:          user.Phone,
				AccountNumber:  user.AccountNumber,
				Password:       "",
				CurrentBalance: user.CurrentBalance,
				CurrentTariff:  user.CurrentTariff,
				Role:           user.Role,
			},
			Date: time.Now().String()[:19],
		}
	}

	return hash, true
}

func LoadSession(m map[string]Session) {
	stmt, ok := request["SessionSelect"]
	if !ok {
		return
	}

	rows, e := stmt.Query()
	if e != nil {
		utils.Logger.Println(e)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var session Session
		e = rows.Scan(&session.Hash, &session.User.ID, &session.User.Name, &session.User.Phone, &session.User.AccountNumber, &session.User.CurrentBalance, &session.User.CurrentTariff.ID, &session.User.CurrentTariff.Name, &session.User.Role.ID, &session.User.Role.Name, &session.Date)
		if e != nil {
			utils.Logger.Println(e)
			return
		}

		m[session.Hash] = session
	}
}

func CheckAdmin() {
	stmt, ok := request["GetAdmin"]
	if !ok {
		utils.Logger.Println("Запрос не подготовлен")
		return
	}
	var admin User
	row := stmt.QueryRow()
	e := row.Scan(&admin.ID)
	if e != nil {
		if e == sql.ErrNoRows {
			e = CreateAdmin()
			if e != nil {
				utils.Logger.Println(e)
				return
			}
		} else {
			utils.Logger.Println(e)
			return
		}
	}

}

func CreateAdmin() error {
	stmt, ok := request["CreateAdmin"]
	if !ok {
		return errors.New("запрос не подготовлен")
	}
	pass := "123"
	pass, e := utils.Encrypt(pass)
	_, e = stmt.Exec("Admin", "1", "Admin", pass, 0, 0, 1)
	if e != nil {
		return e
	}
	return nil
}
