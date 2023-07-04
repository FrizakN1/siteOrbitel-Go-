package email

import (
	"gopkg.in/gomail.v2"
	"newSite/additional"
	"newSite/database"
	"strconv"
)

type Application struct {
	Name             string
	OrganizationName string
	Phone            string
	Address          string
	Services         []string
	Note             string
	Tariff           database.Tariff
}

func SendEmail(typeEmail string, application Application) bool {
	switch typeEmail {
	case "business":
		htmlContent := `<html>
			<head>
				<style>
					body {
						font-family: Arial, sans-serif;
						color: #333;
					}
					h1 {
						color: #0177fd;
					}
				</style>
			</head>
			<body>
				<h1>Бизнес-заявка</h1>
				<p>Имя заказщика: ` + application.Name + `</p>
				<p>Название организации: ` + application.OrganizationName + `</p>
				<p>Контактный телефон: ` + application.Phone + `</p>
				<p>Адрес подключения: ` + application.Address + `</p>
				<p>Услуги: </p>
				<ul>`
		for _, service := range application.Services {
			htmlContent += "<li>" + service + "</li>"
		}

		htmlContent += `</ul></body></html>`

		d := gomail.NewDialer("smtp.yandex.ru", 465, "orbitel-application@yandex.ru", "qsftihaafsgojbko")

		m := gomail.NewMessage()
		m.SetHeader("From", "orbitel-application@yandex.ru")
		m.SetHeader("To", "orbitel-application@yandex.ru")
		m.SetHeader("Subject", "Бизнес-заявка")
		m.SetBody("text/html", htmlContent)

		if e := d.DialAndSend(m); e != nil {
			additional.Logger.Println(e)
			return false
		}
		break
	case "abonent":
		htmlContent := `<html>
			<head>
				<style>
					body {
						font-family: Arial, sans-serif;
						color: #333;
					}
					h1 {
						color: #0177fd;
					}
				</style>
			</head>
			<body>
				<h1>Абонентская-заявка</h1>
				<p>Имя абонента: ` + application.Name + `</p>
				<p>Контактный телефон: ` + application.Phone + `</p>
				<p>Адрес подключения: ` + application.Address + `</p>`

		if application.Note != "" {
			htmlContent += "<p>Комментарий: " + application.Note + "</p>"
		}

		if application.Tariff.ID != 0 {
			tariff, e := database.GetTariff(strconv.Itoa(application.Tariff.ID))
			if e != nil {
				additional.Logger.Println(e)
				return false
			}
			htmlContent += "<p>Тариф: " + tariff.Name + "</p>"
		} else {
			htmlContent += "<p>Тариф: не выбран</p>"
		}

		htmlContent += `</body></html>`

		d := gomail.NewDialer("smtp.yandex.ru", 465, "orbitel-application@yandex.ru", "qsftihaafsgojbko")

		m := gomail.NewMessage()
		m.SetHeader("From", "orbitel-application@yandex.ru")
		m.SetHeader("To", "orbitel-application@yandex.ru")
		m.SetHeader("Subject", "Абонентская-заявка")
		m.SetBody("text/html", htmlContent)

		if e := d.DialAndSend(m); e != nil {
			additional.Logger.Println(e)
			return false
		}
		break
	}
	return true
}
