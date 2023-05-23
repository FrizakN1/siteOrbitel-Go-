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
	Phone          string
	AccountNumber  string
	Password       string
	CurrentBalance float64
	CurrentTariff  Tariff
	Role           Role
}

type Role struct {
	ID   int
	Name string
}

type Deposit struct {
	ID     int
	User   User
	Amount float64
	Date   string
}

type Expense struct {
	ID      int
	User    User
	Amount  float64
	Service Service
	Date    string
}

type PersonalAccount struct {
	User     User
	Expenses []Expense
	Deposits []Deposit
}

var sessionMap map[string]Session

func prepareUser() []string {
	sessionMap = make(map[string]Session)
	if request == nil {
		request = make(map[string]*sql.Stmt)
	}
	errors := make([]string, 0)
	var e error

	//request["SessionSelect"], e = DB.Prepare(`-- SELECT "hash", "id", "name", "phone", "account_number", "current_balance", "current_tariff", "role", "date" FROM "Session" AS s INNER JOIN "User" AS u ON u."id" = s."user"`)
	request["SessionSelect"], e = DB.Prepare(`
		SELECT s.hash, u.id, u.name, u.phone, u.account_number, u.current_balance, t.id, t.name, r.id, r.name, s.date
		FROM "Session" as s
		JOIN "User" as u ON s.user_id = u.id
		JOIN "Role" as r ON u.role = r.id
		LEFT JOIN "Tariff" as t ON u.current_tariff = t.id`)
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

	request["CreateTestUser"], e = DB.Prepare(`INSERT INTO "User" ("name", "phone", "account_number", "password", "current_balance", "current_tariff", "role") VALUES ($1,$2,$3,$4,$5,$6,$7)`)

	request["LoginCheck"], e = DB.Prepare(`
		SELECT u."id", u."name", u."phone", u."account_number", u."current_balance", t."id", t."name", r."id", r."name"
		FROM "User" as u
		JOIN "Role" as r ON u.role = r.id
		LEFT JOIN "Tariff" as t ON u.current_tariff = t.id
		WHERE u.account_number = $1 AND u.password = $2`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["GetAllUsers"], e = DB.Prepare(`
		SELECT u.id, u.name, u.phone, u.account_number, u.current_balance, t.id, t.name, r.id, r.name
		FROM "User" as u
		JOIN "Role" as r ON u.role = r.id
		LEFT JOIN "Tariff" as t ON u.current_tariff = t.id
		WHERE u.id != 1
		ORDER BY u.id DESC`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["GetUser"], e = DB.Prepare(`
		SELECT u.id, u.name, u.phone, u.account_number, u.current_balance, t.id, t.name, r.id, r.name
		FROM "User" as u
		JOIN "Role" as r ON u.role = r.id
		LEFT JOIN "Tariff" as t ON u.current_tariff = t.id
		WHERE u.id = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["CreateUser"], e = DB.Prepare(`INSERT INTO "User" ("name","phone","account_number","password", "current_balance", "current_tariff", "role")
		VALUES ($1,$2,$3,$4,$5,$6,$7)`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["UpdateUser"], e = DB.Prepare(`UPDATE "User" SET "name"=$1,"phone"=$2,"account_number"=$3, "role"=$4 WHERE "id" = $5`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["DeleteUser"], e = DB.Prepare(`DELETE FROM "User" WHERE "id" = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["GetDepositsByID"], e = DB.Prepare(`SELECT d.id, u.id, u.name, d.amount, d.date FROM "Deposit" as d JOIN "User" as u ON d.user_id = u.id WHERE d.user_id = $1 ORDER BY d.id DESC`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["GetExpensesByID"], e = DB.Prepare(`SELECT e.id, u.id, u.name, e.amount, s.id, s.name, e.date 
		FROM "Expense" as e 
		JOIN "User" as u ON e.user_id = u.id  
		JOIN "Service" as s ON e.service_id = s.id
		WHERE e.user_id = $1
		ORDER BY e.id DESC`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["GetDeposit"], e = DB.Prepare(`
		SELECT d.id, u.id, u.name, d.amount, d.date
		FROM "Deposit" as d
		JOIN "User" as u ON d.user_id = u.id
		WHERE d.id = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["GetExpense"], e = DB.Prepare(`
		SELECT e.id, u.id, u.name, e.amount, s.id, s.name, e.date
		FROM "Expense" as e
		JOIN "User" as u ON e.user_id = u.id
		JOIN "Service" as s ON e.service_id = s.id
		WHERE e.id = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	return errors
}

func GetExpense(id string) (Expense, error) {
	stmt, ok := request["GetDeposit"]
	if !ok {
		return Expense{}, errors.New("запрос не подотовлен")
	}

	var expense Expense
	row := stmt.QueryRow(id)
	e := row.Scan(&expense.ID, &expense.User.ID, &expense.User.Name, &expense.Amount, &expense.Service.ID, &expense.Service.Name, &expense.Date)
	if e != nil {
		return Expense{}, e
	}

	return expense, nil
}

func GetDeposit(id string) (Deposit, error) {
	stmt, ok := request["GetDeposit"]
	if !ok {
		return Deposit{}, errors.New("запрос не подотовлен")
	}

	var deposit Deposit
	row := stmt.QueryRow(id)
	e := row.Scan(&deposit.ID, &deposit.User.ID, &deposit.User.Name, &deposit.Amount, &deposit.Date)
	if e != nil {
		return Deposit{}, e
	}

	return deposit, nil
}

func GetExpensesByID(userID string) ([]Expense, error) {
	var expenses []Expense
	stmt, ok := request["GetExpensesByID"]
	if !ok {
		return nil, errors.New("запрос не подотовлен")
	}

	rows, e := stmt.Query(userID)
	if e != nil {
		return nil, e
	}

	for rows.Next() {
		var expense Expense
		e = rows.Scan(&expense.ID, &expense.User.ID, &expense.User.Name, &expense.Amount, &expense.Service.ID, &expense.Service.Name, &expense.Date)
		if e != nil {
			return nil, e
		}
		expenses = append(expenses, expense)
	}

	return expenses, nil
}

func GetDepositsByID(userID string) ([]Deposit, error) {
	var deposits []Deposit
	stmt, ok := request["GetDepositsByID"]
	if !ok {
		return nil, errors.New("запрос не подотовлен")
	}

	rows, e := stmt.Query(userID)
	if e != nil {
		return nil, e
	}

	for rows.Next() {
		var deposit Deposit
		e = rows.Scan(&deposit.ID, &deposit.User.ID, &deposit.User.Name, &deposit.Amount, &deposit.Date)
		if e != nil {
			return nil, e
		}
		deposits = append(deposits, deposit)
	}

	return deposits, nil
}

func DeleteUser(id string) bool {
	stmt, ok := request["DeleteUser"]
	if !ok {
		return false
	}

	_, e := stmt.Exec(id)
	if e != nil {
		utils.Logger.Println(e)
		return false
	}

	return true
}

func (user *User) UpdateUser(id string) error {
	stmt, ok := request["UpdateUser"]
	if !ok {
		return errors.New("запрос не подотовлен")
	}

	_, e := stmt.Exec(user.Name, user.Phone, user.AccountNumber, user.Role.ID, id)
	if e != nil {
		return e
	}

	return nil
}

func (user *User) CreateUser() error {
	stmt, ok := request["CreateUser"]
	if !ok {
		return errors.New("запрос не подотовлен")
	}

	pass, e := utils.Encrypt(user.Password)
	if e != nil {
		return e
	}

	_, e = stmt.Exec(user.Name, user.Phone, user.AccountNumber, pass, 0, nil, user.Role.ID)
	if e != nil {
		return e
	}

	return nil
}

func GetUser(id string) (User, error) {
	stmt, ok := request["GetUser"]
	if !ok {
		return User{}, errors.New("запрос не подготовлен")
	}

	var user User
	var tariffID *int
	var tariffName *string
	row := stmt.QueryRow(id)
	e := row.Scan(&user.ID, &user.Name, &user.Phone, &user.AccountNumber, &user.CurrentBalance, &tariffID, &tariffName, &user.Role.ID, &user.Role.Name)
	if e != nil {
		return User{}, e
	}

	if tariffID != nil {
		user.CurrentTariff.ID = *tariffID
		user.CurrentTariff.Name = *tariffName
	}

	return user, nil
}

func GetAllUsers() ([]User, error) {
	stmt, ok := request["GetAllUsers"]
	if !ok {
		return nil, errors.New("запрос не подготовлен")
	}

	rows, e := stmt.Query()
	if e != nil {
		return nil, e
	}

	var users []User
	for rows.Next() {
		var user User
		var tariffID *int
		var tariffName *string
		e = rows.Scan(&user.ID, &user.Name, &user.Phone, &user.AccountNumber, &user.CurrentBalance, &tariffID, &tariffName, &user.Role.ID, &user.Role.Name)
		if e != nil {
			return nil, e
		}

		if tariffID != nil {
			user.CurrentTariff.ID = *tariffID
			user.CurrentTariff.Name = *tariffName
		}

		users = append(users, user)
	}

	return users, nil
}

func (user *User) UserAuthorizationCheck() bool {
	stmt, ok := request["LoginCheck"]
	if !ok {
		utils.Logger.Println("Запрос не подготовлен")
	}

	row := stmt.QueryRow(user.AccountNumber, user.Password)
	var tariffID *int
	var tariffName *string
	e := row.Scan(&user.ID, &user.Name, &user.Phone, &user.AccountNumber, &user.CurrentBalance, &tariffID, &tariffName, &user.Role.ID, &user.Role.Name)
	if e != nil {
		utils.Logger.Println(e)
		return false
	}

	if tariffID != nil {
		user.CurrentTariff.ID = *tariffID
		user.CurrentTariff.Name = *tariffName
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
		var tariffID *int
		var tariffName *string
		e = rows.Scan(&session.Hash, &session.User.ID, &session.User.Name, &session.User.Phone, &session.User.AccountNumber, &session.User.CurrentBalance, &tariffID, &tariffName, &session.User.Role.ID, &session.User.Role.Name, &session.Date)
		if e != nil {
			utils.Logger.Println(e)
			return
		}

		if tariffID != nil {
			session.User.CurrentTariff.ID = *tariffID
			session.User.CurrentTariff.Name = *tariffName
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
	if e != nil {
		return e
	}
	_, e = stmt.Exec("Admin", "1", "Admin", pass, 0, nil, 1)
	if e != nil {
		return e
	}
	return nil
}

func CreateTestUser() {
	stmt, ok := request["CreateTestUser"]
	if !ok {
		utils.Logger.Println("запрос не подготовлен")
		return
	}
	pass := "123"
	pass, e := utils.Encrypt(pass)
	if e != nil {
		utils.Logger.Println(e)
		return
	}
	_, e = stmt.Exec("Иванов Иван Иванович", "+79195978629", "77788869", pass, 628.20, 2, 3)
	if e != nil {
		utils.Logger.Println(e)
		return
	}
}
