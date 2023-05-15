package database

import (
	"database/sql"
	"errors"
)

type Tariff struct {
	ID             int
	Type           int
	Price          float64
	Name           string
	Description    string
	Speed          int
	DigitalChannel int
	AnalogChannel  int
	Image          string
	Color          string
}

type Service struct {
	ID   int
	Name string
}

type Setting struct {
	ID          int
	Key         string
	Value       string
	Description string
}

var request map[string]*sql.Stmt

func prepareRequest() []string {
	if request == nil {
		request = make(map[string]*sql.Stmt)
	}
	errors := make([]string, 0)
	var e error

	request["GetAllTariffs"], e = DB.Prepare(`SELECT * FROM "Tariff" WHERE "id" != 1 ORDER BY id DESC`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["GetTariff"], e = DB.Prepare(`SELECT * FROM "Tariff" WHERE "id" = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["CreateTariff"], e = DB.Prepare(`INSERT INTO "Tariff" ("type","price","name","description", "speed", "digital_channel", "analog_channel", "image", "color")
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["GetAllSettings"], e = DB.Prepare(`SELECT * FROM "Settings"`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["GetSetting"], e = DB.Prepare(`SELECT * FROM "Settings" WHERE "id" = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	return errors
}

func GetSetting(id string) (Setting, error) {
	stmt, ok := request["GetSetting"]
	if !ok {
		return Setting{}, errors.New("запрос не подотовлен")
	}

	row := stmt.QueryRow(id)

	var setting Setting
	e := row.Scan(&setting.ID, &setting.Key, &setting.Value, &setting.Description)
	if e != nil {
		return Setting{}, e
	}

	return setting, nil
}

func GetAllSettings() ([]Setting, error) {
	var settings []Setting
	stmt, ok := request["GetAllSettings"]
	if !ok {
		return nil, errors.New("запрос не подотовлен")
	}

	rows, e := stmt.Query()
	if e != nil {
		return nil, e
	}

	for rows.Next() {
		var setting Setting
		e = rows.Scan(&setting.ID, &setting.Key, &setting.Value, &setting.Description)
		if e != nil {
			return nil, e
		}
		settings = append(settings, setting)
	}

	return settings, nil
}

func (tariff *Tariff) CreateTariff() error {
	stmt, ok := request["CreateTariff"]
	if !ok {
		return errors.New("запрос не подотовлен")
	}

	_, e := stmt.Exec(tariff.Type, tariff.Price, tariff.Name, tariff.Description, tariff.Speed, tariff.DigitalChannel, tariff.AnalogChannel, tariff.Image, tariff.Color)
	if e != nil {
		return e
	}

	return nil
}

func GetTariff(id string) (Tariff, error) {
	stmt, ok := request["GetTariff"]
	if !ok {
		return Tariff{}, errors.New("запрос не подготовлен")
	}

	var tariff Tariff
	row := stmt.QueryRow(id)
	e := row.Scan(&tariff.ID, &tariff.Type, &tariff.Price, &tariff.Name, &tariff.Description, &tariff.Speed, &tariff.DigitalChannel, &tariff.AnalogChannel, &tariff.Image, &tariff.Color)
	if e != nil {
		return Tariff{}, e
	}

	return tariff, nil
}

func GetAllTariffs() ([]Tariff, error) {
	stmt, ok := request["GetAllTariffs"]
	if !ok {
		return nil, errors.New("запрос не подготовлен")
	}

	rows, e := stmt.Query()
	if e != nil {
		return nil, e
	}

	var tariffs []Tariff
	for rows.Next() {
		var tariff Tariff
		e = rows.Scan(&tariff.ID, &tariff.Type, &tariff.Price, &tariff.Name, &tariff.Description, &tariff.Speed, &tariff.DigitalChannel, &tariff.AnalogChannel, &tariff.Image, &tariff.Color)
		if e != nil {
			return nil, e
		}
		tariffs = append(tariffs, tariff)
	}

	return tariffs, nil
}
