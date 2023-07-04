package database

import (
	"database/sql"
	"errors"
	"newSite/additional"
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
	Address        Address
	Flat           int
	Baned          int
}

type Address struct {
	ID     int
	Street string
	House  string
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
	Service string
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

	request["SessionSelect"], e = DB.Prepare(`
		SELECT s.hash, u.id, u.name, u.phone, u.account_number, u.current_balance, t.id, t.name, r.id, r.name, a.id, a.street, a.house, u.flat, s.date
		FROM "Session" as s
		JOIN "User" as u ON s.user_id = u.id
		JOIN "Role" as r ON u.role = r.id
		LEFT JOIN "Tariff" as t ON u.current_tariff = t.id
		JOIN "Address" as a ON u.house = a.id
		WHERE u.baned = 0`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["SessionSelectByUserID"], e = DB.Prepare(`SELECT hash FROM "Session" WHERE user_id = $1`)
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

	request["LoginCheck"], e = DB.Prepare(`
		SELECT u."id", u."name", u."phone", u."account_number", u."current_balance", t."id", t."name", r."id", r."name", a.id, a.street, a.house, u.flat
		FROM "User" as u
		JOIN "Role" as r ON u.role = r.id
		LEFT JOIN "Tariff" as t ON u.current_tariff = t.id
		JOIN "Address" as a ON u.house = a.id
		WHERE u.account_number = $1 AND u.password = $2 AND u.baned = 0`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["GetAllUsers"], e = DB.Prepare(`
		SELECT u.id, u.name, u.phone, u.account_number, u.current_balance, t.id, t.name, r.id, r.name, a.id, a.street, a.house, u.flat, u.baned
		FROM "User" as u
		JOIN "Role" as r ON u.role = r.id
		LEFT JOIN "Tariff" as t ON u.current_tariff = t.id
		JOIN "Address" as a ON u.house = a.id
		WHERE u.id != 1
		ORDER BY u.id DESC`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["GetUser"], e = DB.Prepare(`
		SELECT u.id, u.name, u.phone, u.account_number, u.current_balance, t.id, t.name, r.id, r.name, a.id, a.street, a.house, u.flat, u.baned
		FROM "User" as u
		JOIN "Role" as r ON u.role = r.id
		LEFT JOIN "Tariff" as t ON u.current_tariff = t.id
		JOIN "Address" as a ON u.house = a.id
		WHERE u.id = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["GetUserByAccountNumber"], e = DB.Prepare(`
		SELECT u.id, u.name, u.phone, u.account_number, u.current_balance, t.id, t.name, r.id, r.name, a.id, a.street, a.house, u.flat, u.baned
		FROM "User" as u
		JOIN "Role" as r ON u.role = r.id
		LEFT JOIN "Tariff" as t ON u.current_tariff = t.id
		JOIN "Address" as a ON u.house = a.id
		WHERE u.account_number = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["CreateUser"], e = DB.Prepare(`INSERT INTO "User" ("name","phone","account_number","password", "current_balance", "current_tariff", "role", "house", "flat", "baned")
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,0)`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["UpdateUser"], e = DB.Prepare(`UPDATE "User" SET "name"=$1,"phone"=$2,"account_number"=$3, "role"=$4, "current_tariff"=$5, "house"=$6, "flat"=$7 WHERE "id" = $8`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["ChangePasswordUser"], e = DB.Prepare(`UPDATE "User" SET "password"=$1 WHERE "id" = $2`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["BanUser"], e = DB.Prepare(`UPDATE "User" SET "baned" = 1 WHERE "id" = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["UnbanUser"], e = DB.Prepare(`UPDATE "User" SET "baned" = 0 WHERE "id" = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["GetDepositsByID"], e = DB.Prepare(`SELECT d.id, u.id, u.name, d.amount, d.date FROM "Deposit" as d JOIN "User" as u ON d.user_id = u.id WHERE d.user_id = $1 ORDER BY d.id DESC`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["GetExpensesByID"], e = DB.Prepare(`SELECT e.id, u.id, u.name, e.amount, e.service, e.date 
		FROM "Expense" as e 
		JOIN "User" as u ON e.user_id = u.id  
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

	request["CreateDeposit"], e = DB.Prepare(`INSERT INTO "Deposit" ("user_id", "amount", "date") VALUES ($1, $2, $3)`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["UpdateUserBalance"], e = DB.Prepare(`UPDATE "User" SET "current_balance" = $1 WHERE "id" = $2`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["GetExpense"], e = DB.Prepare(`
		SELECT e.id, u.id, u.name, e.amount, e.service, e.date
		FROM "Expense" as e
		JOIN "User" as u ON e.user_id = u.id
		WHERE e.id = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["DeleteExpense"], e = DB.Prepare(`DELETE FROM "Expense" WHERE "id" = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["CreateExpense"], e = DB.Prepare(`INSERT INTO "Expense"("user_id", "amount", "service", "date") VALUES ($1, $2, $3, $4)`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	return errors
}

func CreateExpense(user User, expense Expense) error {
	stmt, ok := request["CreateExpense"]
	if !ok {
		return errors.New("запрос не подотовлен")
	}

	_, e := stmt.Exec(user.ID, expense.Amount, expense.Service, time.Now())
	if e != nil {
		return e
	}

	amount := expense.Amount * -1
	e = UpdateUserBalance(user, amount)
	if e != nil {
		return e
	}

	return nil
}

func DeleteExpense(user User, id string, amount float64) error {
	stmt, ok := request["DeleteExpense"]
	if !ok {
		return errors.New("запрос не подотовлен")
	}

	_, e := stmt.Exec(id)
	if e != nil {
		return e
	}

	e = UpdateUserBalance(user, amount)
	if e != nil {
		return e
	}

	return nil
}

func (deposit *Deposit) CreateDeposit(user User) error {
	e := UpdateUserBalance(user, deposit.Amount)
	if e != nil {
		return e
	}
	stmt, ok := request["CreateDeposit"]
	if !ok {
		return errors.New("запрос не подотовлен")
	}

	_, e = stmt.Exec(user.ID, deposit.Amount, time.Now())
	if e != nil {
		return e
	}

	return nil
}

func UpdateUserBalance(user User, amount float64) error {
	stmt, ok := request["UpdateUserBalance"]
	if !ok {
		return errors.New("запрос не подотовлен")
	}
	_, e := stmt.Exec(user.CurrentBalance+amount, user.ID)
	if e != nil {
		return e
	}

	return nil
}

func GetUserByAccountNumber(accountNumber string) (User, error) {
	stmt, ok := request["GetUserByAccountNumber"]
	if !ok {
		return User{}, errors.New("запрос не подотовлен")
	}

	row := stmt.QueryRow(accountNumber)
	var user User
	var tariffID, flat *int
	var tariffName *string
	e := row.Scan(&user.ID, &user.Name, &user.Phone, &user.AccountNumber, &user.CurrentBalance, &tariffID, &tariffName, &user.Role.ID, &user.Role.Name, &user.Address.ID, &user.Address.Street, &user.Address.House, &flat, &user.Baned)
	if e != nil {
		return User{}, e
	}
	if tariffID != nil {
		user.CurrentTariff.ID = *tariffID
		user.CurrentTariff.Name = *tariffName
	}

	if flat != nil {
		user.Flat = *flat
	}
	return user, nil
}

func GetExpense(id string) (Expense, error) {
	stmt, ok := request["GetExpense"]
	if !ok {
		return Expense{}, errors.New("запрос не подотовлен")
	}

	var expense Expense
	row := stmt.QueryRow(id)
	e := row.Scan(&expense.ID, &expense.User.ID, &expense.User.Name, &expense.Amount, &expense.Service, &expense.Date)
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

	defer rows.Close()

	for rows.Next() {
		var expense Expense
		e = rows.Scan(&expense.ID, &expense.User.ID, &expense.User.Name, &expense.Amount, &expense.Service, &expense.Date)
		if e != nil {
			return nil, e
		}

		expense.Date = expense.Date[0:16]
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

	defer rows.Close()

	for rows.Next() {
		var deposit Deposit
		e = rows.Scan(&deposit.ID, &deposit.User.ID, &deposit.User.Name, &deposit.Amount, &deposit.Date)
		if e != nil {
			return nil, e
		}

		deposit.Date = deposit.Date[0:16]

		deposits = append(deposits, deposit)
	}

	return deposits, nil
}

func UnbanUser(id string) error {
	stmt, ok := request["UnbanUser"]
	if !ok {
		return errors.New("запрос не подотовлен")
	}

	_, e := stmt.Exec(id)
	if e != nil {
		return e
	}

	return nil
}

func BanUser(id string) {
	stmt, ok := request["BanUser"]
	if !ok {
		additional.Logger.Println("запрос не подотовлен")
		return
	}

	_, e := stmt.Exec(id)
	if e != nil {
		additional.Logger.Println(e)
		return
	}

	return
}

func DeleteUserFromSessionMap(id string) {
	stmt, ok := request["SessionSelectByUserID"]
	if !ok {
		additional.Logger.Println("запрос не подотовлен")
		return
	}

	rows, e := stmt.Query(id)
	if e != nil {
		additional.Logger.Println(e)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var session Session
		e = rows.Scan(&session.Hash)
		if e != nil {
			additional.Logger.Println(e)
			return
		}

		go delete(sessionMap, session.Hash)
		go DeleteSession(&session)
	}

	return
}

func (user *User) ChangePasswordUser(id string) error {
	stmt, ok := request["ChangePasswordUser"]
	if !ok {
		return errors.New("запрос не подотовлен")
	}

	_, e := stmt.Exec(user.Password, id)
	if e != nil {
		return e
	}

	return nil
}

func (user *User) UpdateUser(id string) error {
	stmt, ok := request["UpdateUser"]
	if !ok {
		return errors.New("запрос не подотовлен")
	}

	var tariff *int
	if user.CurrentTariff.ID == 0 {
		tariff = nil
	} else {
		tariff = &user.CurrentTariff.ID
	}
	_, e := stmt.Exec(user.Name, user.Phone, user.AccountNumber, user.Role.ID, tariff, user.Address.ID, user.Flat, id)
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

	pass, e := additional.Encrypt(user.Password)
	if e != nil {
		return e
	}

	var tariff *int
	if user.CurrentTariff.ID == 0 {
		tariff = nil
	} else {
		tariff = &user.CurrentTariff.ID
	}
	_, e = stmt.Exec(user.Name, user.Phone, user.AccountNumber, pass, 0, tariff, user.Role.ID, user.Address.ID, user.Flat)
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
	var tariffID, flat *int
	var tariffName *string
	row := stmt.QueryRow(id)
	e := row.Scan(&user.ID, &user.Name, &user.Phone, &user.AccountNumber, &user.CurrentBalance, &tariffID, &tariffName, &user.Role.ID, &user.Role.Name, &user.Address.ID, &user.Address.Street, &user.Address.House, &flat, &user.Baned)
	if e != nil {
		return User{}, e
	}

	if tariffID != nil {
		user.CurrentTariff.ID = *tariffID
		user.CurrentTariff.Name = *tariffName
	}

	if flat != nil {
		user.Flat = *flat
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

	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		var tariffID, flat *int
		var tariffName *string
		e = rows.Scan(&user.ID, &user.Name, &user.Phone, &user.AccountNumber, &user.CurrentBalance, &tariffID, &tariffName, &user.Role.ID, &user.Role.Name, &user.Address.ID, &user.Address.Street, &user.Address.House, &flat, &user.Baned)
		if e != nil {
			return nil, e
		}

		if tariffID != nil {
			user.CurrentTariff.ID = *tariffID
			user.CurrentTariff.Name = *tariffName
		}

		if flat != nil {
			user.Flat = *flat
		}

		users = append(users, user)
	}

	return users, nil
}

func (user *User) UserAuthorizationCheck() bool {
	stmt, ok := request["LoginCheck"]
	if !ok {
		additional.Logger.Println("Запрос не подготовлен")
	}

	row := stmt.QueryRow(user.AccountNumber, user.Password)
	var tariffID, flat *int
	var tariffName *string
	e := row.Scan(&user.ID, &user.Name, &user.Phone, &user.AccountNumber, &user.CurrentBalance, &tariffID, &tariffName, &user.Role.ID, &user.Role.Name, &user.Address.ID, &user.Address.Street, &user.Address.House, &flat)
	if e != nil {
		additional.Logger.Println(e)
		return false
	}

	if tariffID != nil {
		user.CurrentTariff.ID = *tariffID
		user.CurrentTariff.Name = *tariffName
	}

	if flat != nil {
		user.Flat = *flat
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
		additional.Logger.Println(e)
	}

	return
}

func CreateSession(user *User) (string, bool) {
	stmt, ok := request["SessionInsert"]
	if !ok {
		return "", false
	}

	hash, e := additional.GenerateHash(user.AccountNumber)
	if e != nil {
		additional.Logger.Println(e)
		return "", false
	}

	_, e = stmt.Exec(hash, user.ID)
	if e != nil {
		additional.Logger.Println(e)
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
				Address:        user.Address,
				Flat:           user.Flat,
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
		additional.Logger.Println(e)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var session Session
		var tariffID, flat *int
		var tariffName *string
		e = rows.Scan(&session.Hash, &session.User.ID, &session.User.Name, &session.User.Phone, &session.User.AccountNumber, &session.User.CurrentBalance, &tariffID, &tariffName, &session.User.Role.ID, &session.User.Role.Name, &session.User.Address.ID, &session.User.Address.Street, &session.User.Address.House, &flat, &session.Date)
		if e != nil {
			additional.Logger.Println(e)
			return
		}

		if tariffID != nil {
			session.User.CurrentTariff.ID = *tariffID
			session.User.CurrentTariff.Name = *tariffName
		}

		if flat != nil {
			session.User.Flat = *flat
		}

		m[session.Hash] = session
	}
}

func CheckAdmin() {
	stmt, ok := request["GetAdmin"]
	if !ok {
		additional.Logger.Println("Запрос не подготовлен")
		return
	}
	var admin User
	row := stmt.QueryRow()
	e := row.Scan(&admin.ID)
	if e != nil {
		if e == sql.ErrNoRows {
			e = CreateAdmin()
			if e != nil {
				additional.Logger.Println(e)
				return
			}
		} else {
			additional.Logger.Println(e)
			return
		}
	}

}

func CreateAdmin() error {
	stmt, ok := request["CreateUser"]
	if !ok {
		return errors.New("запрос не подготовлен")
	}
	pass := "123"
	pass, e := additional.Encrypt(pass)
	if e != nil {
		return e
	}
	_, e = stmt.Exec("Admin", "1", "Admin", pass, 0, nil, 1, 1, 1)
	if e != nil {
		return e
	}
	return nil
}
