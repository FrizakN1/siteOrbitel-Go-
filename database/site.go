package database

import (
	"database/sql"
	"errors"
	"newSite/utils"
)

type Tariff struct {
	ID             int
	Type           Type
	Price          float64
	Name           string
	Description    string
	Speed          int
	DigitalChannel int
	AnalogChannel  int
	Image          string
	Color          string
}

type Type struct {
	ID   int
	Name string
}

type Service struct {
	ID        int
	Name      string
	Note      string
	FullPrice float64
	RentPrice float64
	Type      Type
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

	request["GetAllTariffs"], e = DB.Prepare(`SELECT t.id,ty.id,ty.name,t.price,t.name,t.description,t.speed,t.digital_channel,t.analog_channel,t.image,t.color 
												FROM "Tariff" as t
												JOIN "Tariff_type" as ty ON t.type_id = ty.id
												ORDER BY t.id DESC`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["GetTariff"], e = DB.Prepare(`SELECT t.id,ty.id,ty.name,t.price,t.name,t.description,t.speed,t.digital_channel,t.analog_channel,t.image,t.color 
												FROM "Tariff" as t
												JOIN "Tariff_type" as ty ON t.type_id = ty.id 
												WHERE t.id = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["CreateTariff"], e = DB.Prepare(`INSERT INTO "Tariff" ("type_id","price","name","description", "speed", "digital_channel", "analog_channel", "image", "color")
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["UpdateTariff"], e = DB.Prepare(`UPDATE "Tariff" SET "type_id"=$1,"price"=$2,"name"=$3,"description"=$4, "speed"=$5, "digital_channel"=$6, "analog_channel"=$7, "image"=$8, "color"=$9 WHERE "id" = $10`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["DeleteTariff"], e = DB.Prepare(`DELETE FROM "Tariff" WHERE "id" = $1`)
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

	request["UpdateSetting"], e = DB.Prepare(`UPDATE "Settings" SET "value"=$1,"description"=$2 WHERE "id" = $3`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["GetAllServices"], e = DB.Prepare(`SELECT s.id,s.name,s.note,s.full_price,s.rent_price,ty.id,ty.name 
												FROM "Service" as s
												JOIN "Service_type" as ty ON s.type_id = ty.id`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["GetService"], e = DB.Prepare(`SELECT s.id,s.name,s.note,s.full_price,s.rent_price,ty.id,ty.name 
												FROM "Service" as s
												JOIN "Service_type" as ty ON s.type_id = ty.id
												WHERE s.id = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	return errors
}

func GetService(id string) (Service, error) {
	stmt, ok := request["GetService"]
	if !ok {
		return Service{}, errors.New("запрос не подотовлен")
	}

	row := stmt.QueryRow(id)

	var service Service
	var fullPrice, rentPrice *float64
	var note *string
	e := row.Scan(&service.ID, &service.Name, &note, &fullPrice, &rentPrice, &service.Type.ID, &service.Type.Name)
	if e != nil {
		return Service{}, e
	}
	if note != nil {
		service.Note = *note
	}
	if fullPrice != nil {
		service.FullPrice = *fullPrice
	}
	if rentPrice != nil {
		service.RentPrice = *rentPrice
	}

	return service, nil
}

func GetAllServices() ([]Service, error) {
	stmt, ok := request["GetAllServices"]
	if !ok {
		return nil, errors.New("запрос не подготовлен")
	}

	rows, e := stmt.Query()
	if e != nil {
		return nil, e
	}

	var services []Service

	for rows.Next() {
		var service Service
		var note *string
		var fullPrice, rentPrice *float64
		e = rows.Scan(&service.ID, &service.Name, &note, &fullPrice, &rentPrice, &service.Type.ID, &service.Type.Name)
		if e != nil {
			utils.Logger.Println(e)
			return nil, e
		}
		if note != nil {
			service.Note = *note
		}
		if fullPrice != nil {
			service.FullPrice = *fullPrice
		}
		if rentPrice != nil {
			service.RentPrice = *rentPrice
		}

		services = append(services, service)
	}

	return services, nil
}

func (settings *Setting) UpdateSettings(id string) error {
	stmt, ok := request["UpdateSetting"]
	if !ok {
		return errors.New("запрос не подотовлен")
	}

	_, e := stmt.Exec(settings.Value, settings.Description, id)
	if e != nil {
		return e
	}

	return nil
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

func DeleteTariff(id string) bool {
	stmt, ok := request["DeleteTariff"]
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

func (tariff *Tariff) UpdateTariff(id string) error {
	stmt, ok := request["UpdateTariff"]
	if !ok {
		return errors.New("запрос не подотовлен")
	}

	_, e := stmt.Exec(tariff.Type.ID, tariff.Price, tariff.Name, tariff.Description, tariff.Speed, tariff.DigitalChannel, tariff.AnalogChannel, tariff.Image, tariff.Color, id)
	if e != nil {
		return e
	}

	return nil
}

func (tariff *Tariff) CreateTariff() error {
	stmt, ok := request["CreateTariff"]
	if !ok {
		return errors.New("запрос не подотовлен")
	}

	_, e := stmt.Exec(tariff.Type.ID, tariff.Price, tariff.Name, tariff.Description, tariff.Speed, tariff.DigitalChannel, tariff.AnalogChannel, tariff.Image, tariff.Color)
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
	var digitalChannel, analogChannel, speed *int
	var image *string
	e := row.Scan(&tariff.ID, &tariff.Type.ID, &tariff.Type.Name, &tariff.Price, &tariff.Name, &tariff.Description, &speed, &digitalChannel, &analogChannel, &image, &tariff.Color)
	if e != nil {
		return Tariff{}, e
	}

	if digitalChannel != nil {
		tariff.DigitalChannel = *digitalChannel
		tariff.AnalogChannel = *analogChannel
		tariff.Speed = *speed
	}
	if image != nil {
		tariff.Image = *image
		tariff.Speed = *speed
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
		var digitalChannel, analogChannel, speed *int
		var image *string
		e = rows.Scan(&tariff.ID, &tariff.Type.ID, &tariff.Type.Name, &tariff.Price, &tariff.Name, &tariff.Description, &speed, &digitalChannel, &analogChannel, &image, &tariff.Color)
		if e != nil {
			return nil, e
		}

		if digitalChannel != nil {
			tariff.DigitalChannel = *digitalChannel
			tariff.AnalogChannel = *analogChannel
			tariff.Speed = *speed
		}
		if image != nil {
			tariff.Image = *image
			tariff.Speed = *speed
		}

		tariffs = append(tariffs, tariff)
	}

	return tariffs, nil
}
