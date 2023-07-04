package database

import (
	"database/sql"
	"errors"
	"newSite/additional"
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

type SEO struct {
	ID          int
	Title       string
	Keywords    string
	Description string
	Uri         string
}

type FAQ struct {
	ID       int
	Question string
	Answer   string
}

var request map[string]*sql.Stmt
var SettingsMap map[string]string
var SeoMap map[string]SEO

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

	request["GetTariffsByType"], e = DB.Prepare(`SELECT t.id,ty.id,ty.name,t.price,t.name,t.description,t.speed,t.digital_channel,t.analog_channel,t.image,t.color 
												FROM "Tariff" as t
												JOIN "Tariff_type" as ty ON t.type_id = ty.id
												WHERE t.type_id = $1
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

	request["GetAllSettings"], e = DB.Prepare(`SELECT * FROM "Settings" ORDER BY "id"`)
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
												JOIN "Service_type" as ty ON s.type_id = ty.id ORDER BY s.id DESC`)
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

	request["CreateService"], e = DB.Prepare(`INSERT INTO "Service" (name, note, full_price, rent_price, type_id) VALUES ($1,$2,$3,$4,$5)`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["UpdateService"], e = DB.Prepare(`UPDATE "Service" SET name = $1, note = $2, full_price = $3, rent_price = $4, type_id = $5 WHERE id = $6`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["DeleteService"], e = DB.Prepare(`DELETE FROM "Service" WHERE "id" = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["GetAllAddresses"], e = DB.Prepare(`SELECT id,street,house FROM "Address" ORDER BY id`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["GetAddress"], e = DB.Prepare(`SELECT id,street,house FROM "Address" WHERE id = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["CreateAddress"], e = DB.Prepare(`INSERT INTO "Address" ("street", "house") VALUES ($1,$2)`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["UpdateAddress"], e = DB.Prepare(`UPDATE "Address" SET "street" = $1, "house" = $2 WHERE id = $3`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["DeleteAddress"], e = DB.Prepare(`DELETE FROM "Address" WHERE "id" = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["GetAllSEO"], e = DB.Prepare(`SELECT id,title,keywords,description,uri FROM "SEO" ORDER BY id`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["GetSEO"], e = DB.Prepare(`SELECT id,title,keywords,description,uri FROM "SEO" WHERE id = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["CreateSEO"], e = DB.Prepare(`INSERT INTO "SEO" (title,keywords,description,uri) VALUES ($1,$2,$3,$4)`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["UpdateSEO"], e = DB.Prepare(`UPDATE "SEO" SET title = $1, keywords = $2, description = $3 WHERE id = $4`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["GetAllFaq"], e = DB.Prepare(`SELECT id, question, answer FROM "Faq" ORDER BY id`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["GetFaq"], e = DB.Prepare(`SELECT id, question, answer FROM "Faq" WHERE id = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["CreateFaq"], e = DB.Prepare(`INSERT INTO "Faq" (question, answer) VALUES ($1,$2)`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["UpdateFaq"], e = DB.Prepare(`UPDATE "Faq" SET question = $1, answer = $2 WHERE id = $3`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	request["DeleteFaq"], e = DB.Prepare(`DELETE FROM "Faq" WHERE "id" = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	return errors
}

func DeleteFaq(id string) error {
	stmt, ok := request["DeleteFaq"]
	if !ok {
		return errors.New("запрос не подготовлен")
	}

	_, e := stmt.Exec(id)
	if e != nil {
		return e
	}

	return nil
}

func (faq *FAQ) UpdateFaq(id string) error {
	stmt, ok := request["UpdateFaq"]
	if !ok {
		return errors.New("запрос не подготовлен")
	}

	_, e := stmt.Exec(faq.Question, faq.Answer, id)
	if e != nil {
		return e
	}

	return nil
}

func (faq *FAQ) CreateFaq() error {
	stmt, ok := request["CreateFaq"]
	if !ok {
		return errors.New("запрос не подготовлен")
	}

	_, e := stmt.Exec(faq.Question, faq.Answer)
	if e != nil {
		return e
	}

	return nil
}

func GetFaq(id string) (FAQ, error) {
	stmt, ok := request["GetFaq"]
	if !ok {
		return FAQ{}, errors.New("запрос не подотовлен")
	}

	row := stmt.QueryRow(id)

	var faq FAQ
	e := row.Scan(&faq.ID, &faq.Question, &faq.Answer)
	if e != nil {
		return FAQ{}, e
	}

	return faq, nil
}

func GetAllFaq() ([]FAQ, error) {
	stmt, ok := request["GetAllFaq"]
	if !ok {
		return nil, errors.New("запрос не подготовлен")
	}

	rows, e := stmt.Query()
	if e != nil {
		return nil, e
	}

	defer rows.Close()

	var faq []FAQ

	for rows.Next() {
		var faqRow FAQ
		e = rows.Scan(&faqRow.ID, &faqRow.Question, &faqRow.Answer)
		if e != nil {
			additional.Logger.Println(e)
			return nil, e
		}

		faq = append(faq, faqRow)
	}

	return faq, nil
}

func (seo *SEO) UpdateSEO(id string) error {
	stmt, ok := request["UpdateSEO"]
	if !ok {
		return errors.New("запрос не подготовлен")
	}

	SeoMap[seo.Uri] = *seo

	_, e := stmt.Exec(seo.Title, seo.Keywords, seo.Description, id)
	if e != nil {
		return e
	}

	return nil
}

func (seo *SEO) CreateSEO() error {
	stmt, ok := request["CreateSEO"]
	if !ok {
		return errors.New("запрос не подготовлен")
	}

	_, e := stmt.Exec(seo.Title, seo.Keywords, seo.Description, seo.Uri)
	if e != nil {
		return e
	}

	return nil
}

func GetSEO(id string) (SEO, error) {
	stmt, ok := request["GetSEO"]
	if !ok {
		return SEO{}, errors.New("запрос не подготовлен")
	}

	var seo SEO
	row := stmt.QueryRow(id)
	e := row.Scan(&seo.ID, &seo.Title, &seo.Keywords, &seo.Description, &seo.Uri)
	if e != nil {
		return SEO{}, e
	}

	return seo, nil
}

func GetAllSEO() ([]SEO, error) {
	stmt, ok := request["GetAllSEO"]
	if !ok {
		return nil, errors.New("запрос не подготовлен")
	}

	rows, e := stmt.Query()
	if e != nil {
		return nil, e
	}

	defer rows.Close()

	var allSEO []SEO
	for rows.Next() {
		var seo SEO
		e = rows.Scan(&seo.ID, &seo.Title, &seo.Keywords, &seo.Description, &seo.Uri)
		if e != nil {
			return nil, e
		}

		allSEO = append(allSEO, seo)
	}

	return allSEO, nil
}

func DeleteAddress(id string) error {
	stmt, ok := request["DeleteAddress"]
	if !ok {
		return errors.New("запрос не подготовлен")
	}

	_, e := stmt.Exec(id)
	if e != nil {
		return e
	}

	return nil
}

func (address *Address) UpdateAddress(id string) error {
	stmt, ok := request["UpdateAddress"]
	if !ok {
		return errors.New("запрос не подготовлен")
	}

	_, e := stmt.Exec(address.Street, address.House, id)
	if e != nil {
		return e
	}

	return nil
}

func (address *Address) CreateAddress() error {
	stmt, ok := request["CreateAddress"]
	if !ok {
		return errors.New("запрос не подготовлен")
	}

	_, e := stmt.Exec(address.Street, address.House)
	if e != nil {
		return e
	}

	return nil
}

func GetAddress(id string) (Address, error) {
	stmt, ok := request["GetAddress"]
	if !ok {
		return Address{}, errors.New("запрос не подготовлен")
	}

	var address Address
	row := stmt.QueryRow(id)
	e := row.Scan(&address.ID, &address.Street, &address.House)
	if e != nil {
		return Address{}, e
	}

	return address, nil
}

func GetAllAddresses() ([]Address, error) {
	stmt, ok := request["GetAllAddresses"]
	if !ok {
		return nil, errors.New("запрос не подготовлен")
	}

	rows, e := stmt.Query()
	if e != nil {
		return nil, e
	}

	defer rows.Close()

	var addresses []Address
	for rows.Next() {
		var address Address
		e = rows.Scan(&address.ID, &address.Street, &address.House)
		if e != nil {
			return nil, e
		}

		addresses = append(addresses, address)
	}

	return addresses, nil
}

func GetTariffsByType(typeID int) ([]Tariff, error) {
	stmt, ok := request["GetTariffsByType"]
	if !ok {
		return nil, errors.New("запрос не подготовлен")
	}

	rows, e := stmt.Query(typeID)
	if e != nil {
		return nil, e
	}

	defer rows.Close()

	var tariffs []Tariff
	for rows.Next() {
		var tariff Tariff
		var digitalChannel, analogChannel, speed *int
		var image *string
		e = rows.Scan(&tariff.ID, &tariff.Type.ID, &tariff.Type.Name, &tariff.Price, &tariff.Name, &tariff.Description, &tariff.Speed, &tariff.DigitalChannel, &tariff.AnalogChannel, &tariff.Image, &tariff.Color)
		if e != nil {
			return nil, e
		}
		if digitalChannel != nil {
			tariff.DigitalChannel = *digitalChannel
		}
		if analogChannel != nil {
			tariff.AnalogChannel = *analogChannel
		}
		if speed != nil {
			tariff.Speed = *speed
		}
		if image != nil {
			tariff.Image = *image
		}

		tariffs = append(tariffs, tariff)
	}

	return tariffs, nil
}

func DeleteService(id string) error {
	stmt, ok := request["DeleteService"]
	if !ok {
		return errors.New("запрос не подготовлен")
	}

	_, e := stmt.Exec(id)
	if e != nil {
		return e
	}

	return nil
}

func (service *Service) UpdateService(id string) error {
	stmt, ok := request["UpdateService"]
	if !ok {
		return errors.New("запрос не подготовлен")
	}

	_, e := stmt.Exec(service.Name, service.Note, service.FullPrice, service.RentPrice, service.Type.ID, id)
	if e != nil {
		return e
	}

	return nil
}

func (service *Service) CreateService() error {
	stmt, ok := request["CreateService"]
	if !ok {
		return errors.New("запрос не подготовлен")
	}

	_, e := stmt.Exec(service.Name, service.Note, service.FullPrice, service.RentPrice, service.Type.ID)
	if e != nil {
		return e
	}

	return nil
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

	defer rows.Close()

	var services []Service

	for rows.Next() {
		var service Service
		var note *string
		var fullPrice, rentPrice *float64
		e = rows.Scan(&service.ID, &service.Name, &note, &fullPrice, &rentPrice, &service.Type.ID, &service.Type.Name)
		if e != nil {
			additional.Logger.Println(e)
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

	SettingsMap[settings.Key] = settings.Value

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

	defer rows.Close()

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

func DeleteTariff(id string) error {
	stmt, ok := request["DeleteTariff"]
	if !ok {
		return errors.New("запрос не подготовлен")
	}

	_, e := stmt.Exec(id)
	if e != nil {
		return e
	}

	return nil
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
	}
	if analogChannel != nil {
		tariff.AnalogChannel = *analogChannel
	}
	if speed != nil {
		tariff.Speed = *speed
	}
	if image != nil {
		tariff.Image = *image
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

	defer rows.Close()

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
		}
		if analogChannel != nil {
			tariff.AnalogChannel = *analogChannel
		}
		if speed != nil {
			tariff.Speed = *speed
		}
		if image != nil {
			tariff.Image = *image
		}

		tariffs = append(tariffs, tariff)
	}

	return tariffs, nil
}

func LoadSettings(m *map[string]string) {
	stmt, ok := request["GetAllSettings"]
	if !ok {
		additional.Logger.Println("запрос не подотовлен")
		return
	}

	rows, e := stmt.Query()
	if e != nil {
		additional.Logger.Println(e)
		return
	}

	defer rows.Close()

	if *m == nil {
		*m = make(map[string]string)
	}

	for rows.Next() {
		var setting Setting
		e = rows.Scan(&setting.ID, &setting.Key, &setting.Value, &setting.Description)
		if e != nil {
			additional.Logger.Println(e)
			return
		}

		(*m)[setting.Key] = setting.Value
	}
}

func LoadSEO(m *map[string]SEO) {
	stmt, ok := request["GetAllSEO"]
	if !ok {
		additional.Logger.Println("запрос не подотовлен")
		return
	}

	rows, e := stmt.Query()
	if e != nil {
		additional.Logger.Println(e)
		return
	}

	defer rows.Close()

	if *m == nil {
		*m = make(map[string]SEO)
	}

	for rows.Next() {
		var seo SEO
		e = rows.Scan(&seo.ID, &seo.Title, &seo.Keywords, &seo.Description, &seo.Uri)
		if e != nil {
			additional.Logger.Println(e)
			return
		}

		(*m)[seo.Uri] = seo
	}
}
