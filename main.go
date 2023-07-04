package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"newSite/additional"
	"newSite/database"
	"newSite/email"
	"strconv"
)

func main() {
	database.ConnectDB()
	database.CheckAdmin()
	//database.CreateTestUser()
	router := gin.Default()
	store := sessions.NewCookieStore([]byte("secretWord"))
	router.Use(sessions.Sessions("session", store))
	router.LoadHTMLGlob("template/**/*.html")
	router.Static("assets", "assets")
	router.NoRoute(noRoute)
	router.GET("/", index)
	router.GET("/abonent_application-:tariff", abonentApplication)
	router.GET("/tarif_for_home", tarif_for_home)
	router.GET("/inettv", inettv)
	router.GET("/tv", tv)
	router.GET("/phone", phone)
	router.GET("/dop_for_home", dop_for_home)
	router.GET("/business/inet-for-ul", inet_for_ul)
	router.GET("/business/phone-for-ul", phone_for_ul)
	router.GET("/business/vpn-for-ul", vpn_for_ul)
	router.GET("/about", about)
	router.GET("/calculator", calculator)
	router.GET("/get_tariffs/:type", getTariffs)
	router.GET("/get_services", getServices)
	router.POST("/send_email-:type", sendEmail)
	router.GET("/addresses/get-all", getAddresses)
	router.GET("/virtual-ats", virtual_ats)
	router.GET("/routers", routers)
	router.GET("/tv-manual", tv_manual)
	router.GET("/oplata", oplata)
	router.GET("/faq", faq)
	router.GET("/replenishment_balance", replenishmentBalance)
	router.POST("/replenishment_balance/deposit", replenishmentBalanceDeposit)
	routerUser := router.Group("/user")
	routerUser.GET("/personal_account", userPersonalAccount)
	routerUser.GET("/personal_account/get_data", userPersonalAccountGetData)
	routerUser.GET("/authorization", userAuthorization)
	routerUser.POST("/exit", exit)
	routerUser.POST("/authorization_check", userAuthorizationCheck)
	routerAdmin := router.Group("/admin")
	routerAdmin.GET("/", adminIndex)
	routerAdmin.GET("/:obj", adminObjs)
	routerAdmin.GET("/:obj/user-:id", adminObjsByID)
	routerAdmin.GET("/:obj/view-:id", adminView)
	routerAdmin.GET("/:obj/edit-:id", adminEdit)
	routerAdmin.POST("/:obj/delete-:id", adminDelete)
	routerAdmin.POST("/expenses/user-:userID/delete-:id", adminDeleteExpense)
	routerAdmin.GET("/expenses/create/user-:userID", adminCreateExpense)
	routerAdmin.POST("/expenses/create/user-:userID", adminCreateExpensePOST)
	routerAdmin.GET("/:obj/create", adminCreate)
	routerAdmin.POST("/:obj/create", adminCreatePOST)
	routerAdmin.POST("/:obj/update-:id", adminUpdatePOST)
	routerAdmin.POST("/users/unban-:id", adminUnbanUser)

	_ = router.Run("192.168.0.105:8080")
}

func adminCreateExpensePOST(c *gin.Context) {
	userID := c.Param("userID")
	var expense database.Expense
	e := c.BindJSON(&expense)
	if e != nil {
		additional.Logger.Println(e)
		c.JSON(400, false)
		return
	}

	user, e := database.GetUser(userID)
	if e != nil {
		additional.Logger.Println(e)
		c.JSON(400, false)
		return
	}

	e = database.CreateExpense(user, expense)
	if e != nil {
		additional.Logger.Println(e)
		c.JSON(400, false)
		return
	}

	c.JSON(200, true)
}

func adminCreateExpense(c *gin.Context) {
	byID := c.Param("userID")

	c.HTML(200, "adminExpensesCreate", gin.H{
		"Title":           "Просмотр пользователя",
		"MainActionTitle": "Пользователи",
		"Active":          "users",
		"Mode":            "view",
		"ByID":            byID,
		"SubMode":         "create",
		"SubTitle":        "История списания",
		"SubActive":       "expenses",
	})
}

func adminDeleteExpense(c *gin.Context) {
	userID := c.Param("userID")
	id := c.Param("id")

	var expense database.Expense
	e := c.BindJSON(&expense)
	if e != nil {
		additional.Logger.Println(e)
		c.JSON(400, false)
		return
	}

	user, e := database.GetUser(userID)
	if e != nil {
		additional.Logger.Println(e)
		c.JSON(400, false)
		return
	}

	e = database.DeleteExpense(user, id, expense.Amount)
	if e != nil {
		additional.Logger.Println(e)
		c.JSON(400, false)
		return
	}

	c.JSON(200, true)
}

func replenishmentBalanceDeposit(c *gin.Context) {
	var deposit database.Deposit
	e := c.BindJSON(&deposit)
	if e != nil {
		additional.Logger.Println(e)
		c.JSON(400, false)
		return
	}

	user, e := database.GetUserByAccountNumber(deposit.User.AccountNumber)
	if e != nil {
		additional.Logger.Println(e)
		c.JSON(400, false)
		return
	}

	e = deposit.CreateDeposit(user)
	if e != nil {
		additional.Logger.Println(e)
		c.JSON(400, false)
		return
	}

	c.JSON(200, true)
}

func replenishmentBalance(c *gin.Context) {
	session := getSession(c)

	c.HTML(200, "replenishmentBalance", gin.H{
		"Settings": database.SettingsMap,
		"SEO":      database.SeoMap["/replenishment_balance"],
		"User":     session.User,
	})
}

func noRoute(c *gin.Context) {
	c.HTML(200, "error", gin.H{
		"StatusCode": 404,
		"StatusText": "Страница не найдена",
	})
}

func adminUnbanUser(c *gin.Context) {
	session := getSession(c)
	if session.User.Role.ID != 1 {
		c.AbortWithStatus(403)
	} else {
		id := c.Param("id")
		e := database.UnbanUser(id)
		if e != nil {
			additional.Logger.Println(e)
			c.JSON(400, false)
			return
		}

		c.JSON(200, true)
	}
}

func getAddresses(c *gin.Context) {
	addresses, e := database.GetAllAddresses()
	if e != nil {
		additional.Logger.Println(e)
		c.JSON(400, false)
		return
	}

	c.JSON(200, addresses)
}

func sendEmail(c *gin.Context) {
	typeEmail := c.Param("type")
	var application email.Application
	e := c.BindJSON(&application)
	if e != nil {
		additional.Logger.Println(e)
		return
	}
	status := email.SendEmail(typeEmail, application)
	if status {
		c.JSON(200, status)
	} else {
		c.JSON(400, status)
	}
}

func getServices(c *gin.Context) {
	services, e := database.GetAllServices()
	if e != nil {
		additional.Logger.Println(e)
		c.JSON(400, false)
		return
	}

	c.JSON(200, services)
}

func getTariffs(c *gin.Context) {
	tariffType := c.Param("type")
	switch tariffType {
	case "all":
		tariffs, e := database.GetAllTariffs()
		if e != nil {
			additional.Logger.Println(e)
			c.JSON(400, false)
			return
		}
		c.JSON(200, tariffs)
		break
	}
}

func calculator(c *gin.Context) {
	c.HTML(200, "calculator", gin.H{
		"Settings": database.SettingsMap,
		"SEO":      database.SeoMap["/calculator"],
	})
}

func adminDelete(c *gin.Context) {
	session := getSession(c)
	if session.User.Role.ID != 1 {
		c.AbortWithStatus(403)
	} else {
		obj := c.Param("obj")
		objID := c.Param("id")
		switch obj {
		case "users":
			go database.BanUser(objID)
			go database.DeleteUserFromSessionMap(objID)
			c.JSON(200, true)
			break
		case "tariffs":
			e := database.DeleteTariff(objID)
			if e != nil {
				additional.Logger.Println(e)
				c.JSON(400, false)
				return
			}
			c.JSON(200, true)
			break
		case "services":
			e := database.DeleteService(objID)
			if e != nil {
				c.JSON(400, false)
				return
			}
			c.JSON(200, true)
			break
		case "addresses":
			e := database.DeleteAddress(objID)
			if e != nil {
				c.JSON(400, false)
				return
			}
			c.JSON(200, true)
			break
		case "faq":
			e := database.DeleteFaq(objID)
			if e != nil {
				c.JSON(400, false)
				return
			}
			c.JSON(200, true)
			break
		}
	}
}

func adminUpdatePOST(c *gin.Context) {
	session := getSession(c)
	if session.User.Role.ID != 1 {
		c.AbortWithStatus(403)
	} else {
		obj := c.Param("obj")
		objID := c.Param("id")
		switch obj {
		case "users":
			var user database.User
			e := c.BindJSON(&user)
			if e != nil {
				additional.Logger.Println(e)
				c.JSON(400, false)
				return
			}
			e = user.UpdateUser(objID)
			if user.Password != "" {
				user.Password, e = additional.Encrypt(user.Password)
				if e != nil {
					additional.Logger.Println(e)
					c.JSON(400, false)
					return
				}

				e = user.ChangePasswordUser(objID)
				if e != nil {
					additional.Logger.Println(e)
					c.JSON(400, false)
					return
				}
			}
			if e != nil {
				additional.Logger.Println(e)
				c.JSON(400, false)
				return
			}
			c.JSON(200, true)
			break
		case "tariffs":
			var tariff database.Tariff
			e := c.BindJSON(&tariff)
			if e != nil {
				additional.Logger.Println(e)
				c.JSON(400, false)
				return
			}
			e = tariff.UpdateTariff(objID)
			if e != nil {
				additional.Logger.Println(e)
				c.JSON(400, false)
				return
			}
			c.JSON(200, true)
			break
		case "settings":
			var settings database.Setting
			e := c.BindJSON(&settings)
			if e != nil {
				additional.Logger.Println(e)
				c.JSON(400, false)
				return
			}
			e = settings.UpdateSettings(objID)
			if e != nil {
				additional.Logger.Println(e)
				c.JSON(400, false)
				return
			}
			c.JSON(200, true)
			break
		case "services":
			var service database.Service
			e := c.BindJSON(&service)
			if e != nil {
				additional.Logger.Println(e)
				c.JSON(400, false)
				return
			}
			e = service.UpdateService(objID)
			if e != nil {
				additional.Logger.Println(e)
				c.JSON(400, false)
				return
			}
			c.JSON(200, true)
			break
		case "addresses":
			var address database.Address
			e := c.BindJSON(&address)
			if e != nil {
				additional.Logger.Println(e)
				c.JSON(400, false)
				return
			}
			e = address.UpdateAddress(objID)
			if e != nil {
				additional.Logger.Println(e)
				c.JSON(400, false)
				return
			}
			c.JSON(200, true)
			break
		case "seo":
			var seo database.SEO
			e := c.BindJSON(&seo)
			if e != nil {
				additional.Logger.Println(e)
				c.JSON(400, false)
				return
			}
			e = seo.UpdateSEO(objID)
			if e != nil {
				additional.Logger.Println(e)
				c.JSON(400, false)
				return
			}
			c.JSON(200, true)
			break
		case "faq":
			var faq database.FAQ
			e := c.BindJSON(&faq)
			if e != nil {
				additional.Logger.Println(e)
				c.JSON(400, false)
				return
			}
			e = faq.UpdateFaq(objID)
			if e != nil {
				additional.Logger.Println(e)
				c.JSON(400, false)
				return
			}
			c.JSON(200, true)
			break
		}
	}
}

func adminCreatePOST(c *gin.Context) {
	session := getSession(c)
	if session.User.Role.ID != 1 {
		c.AbortWithStatus(403)
	} else {
		obj := c.Param("obj")
		switch obj {
		case "users":
			var user database.User
			e := c.BindJSON(&user)
			if e != nil {
				additional.Logger.Println(e)
				c.JSON(400, false)
				return
			}
			e = user.CreateUser()
			if e != nil {
				additional.Logger.Println(e)
				c.JSON(400, false)
				return
			}
			c.JSON(200, true)
			break
		case "tariffs":
			var tariff database.Tariff
			e := c.BindJSON(&tariff)
			if e != nil {
				additional.Logger.Println(e)
				c.JSON(400, false)
				return
			}
			e = tariff.CreateTariff()
			if e != nil {
				additional.Logger.Println(e)
				c.JSON(400, false)
				return
			}
			c.JSON(200, true)
			break
		case "services":
			var service database.Service
			e := c.BindJSON(&service)
			if e != nil {
				additional.Logger.Println(e)
				c.JSON(400, false)
				return
			}
			e = service.CreateService()
			if e != nil {
				additional.Logger.Println(e)
				c.JSON(400, false)
				return
			}
			c.JSON(200, true)
			break
		case "addresses":
			var address database.Address
			e := c.BindJSON(&address)
			if e != nil {
				additional.Logger.Println(e)
				c.JSON(400, false)
				return
			}
			e = address.CreateAddress()
			if e != nil {
				additional.Logger.Println(e)
				c.JSON(400, false)
				return
			}
			c.JSON(200, true)
			break
		//case "seo":
		//	var seo database.SEO
		//	e := c.BindJSON(&seo)
		//	if e != nil {
		//		additional.Logger.Println(e)
		//		c.JSON(400, false)
		//		return
		//	}
		//	e = seo.CreateSEO()
		//	if e != nil {
		//		additional.Logger.Println(e)
		//		c.JSON(400, false)
		//		return
		//	}
		//	c.JSON(200, true)
		//	break
		case "faq":
			var faq database.FAQ
			e := c.BindJSON(&faq)
			if e != nil {
				additional.Logger.Println(e)
				c.JSON(400, false)
				return
			}
			e = faq.CreateFaq()
			if e != nil {
				additional.Logger.Println(e)
				c.JSON(400, false)
				return
			}
			c.JSON(200, true)
			break
		}
	}
}

func adminEdit(c *gin.Context) {
	session := getSession(c)
	if session.User.Role.ID != 1 {
		c.HTML(200, "error", gin.H{
			"StatusCode": 403,
			"StatusText": "У вас недостаточно прав для просмотра этой страницы",
		})
	} else {
		obj := c.Param("obj")
		objID := c.Param("id")
		switch obj {
		case "users":
			user, e := database.GetUser(objID)
			if e != nil {
				additional.Logger.Println(e)
				return
			}

			c.HTML(200, "adminUsersEdit", gin.H{
				"Title":           "Изменение пользователя: ",
				"MainActionTitle": "Пользователи",
				"Active":          "users",
				"Obj":             user,
				"Mode":            "edit",
			})
			break
		case "tariffs":
			tariff, e := database.GetTariff(objID)
			if e != nil {
				additional.Logger.Println(e)
				return
			}

			c.HTML(200, "adminTariffsEdit", gin.H{
				"Title":           "Изменение тарифа: ",
				"MainActionTitle": "Тарифы",
				"Active":          "tariffs",
				"Obj":             tariff,
				"Mode":            "edit",
			})
			break
		case "settings":
			setting, e := database.GetSetting(objID)
			if e != nil {
				additional.Logger.Println(e)
				return
			}

			c.HTML(200, "adminSettingsEdit", gin.H{
				"Title":           "Изменение настроек: ",
				"MainActionTitle": "Настройки",
				"Active":          "settings",
				"Obj":             setting,
				"Mode":            "edit",
			})
			break
		case "services":
			service, e := database.GetService(objID)
			if e != nil {
				additional.Logger.Println(e)
				return
			}

			c.HTML(200, "adminServicesEdit", gin.H{
				"Title":           "Изменение услуг: ",
				"MainActionTitle": "Услуги",
				"Active":          "services",
				"Obj":             service,
				"Mode":            "edit",
			})
			break
		case "addresses":
			address, e := database.GetAddress(objID)
			if e != nil {
				additional.Logger.Println(e)
				return
			}

			c.HTML(200, "adminAddressesEdit", gin.H{
				"Title":           "Изменение адреса: ",
				"MainActionTitle": "Список адресов",
				"Active":          "addresses",
				"Obj":             address,
				"Mode":            "edit",
			})
			break
		case "seo":
			seo, e := database.GetSEO(objID)
			if e != nil {
				additional.Logger.Println(e)
				return
			}

			c.HTML(200, "adminSEOEdit", gin.H{
				"Title":           "Изменение SEO-настроек: ",
				"MainActionTitle": "SEO-настройки",
				"Active":          "seo",
				"Obj":             seo,
				"Mode":            "edit",
			})
			break
		case "faq":
			faq, e := database.GetFaq(objID)
			if e != nil {
				additional.Logger.Println(e)
				return
			}

			c.HTML(200, "adminFaqEdit", gin.H{
				"Title":           "Изменение FAQ: ",
				"MainActionTitle": "FAQ",
				"Active":          "faq",
				"Obj":             faq,
				"Mode":            "edit",
			})
			break
		}
	}
}

func adminCreate(c *gin.Context) {
	session := getSession(c)
	if session.User.Role.ID != 1 {
		c.HTML(200, "error", gin.H{
			"StatusCode": 403,
			"StatusText": "У вас недостаточно прав для просмотра этой страницы",
		})
	} else {
		obj := c.Param("obj")
		switch obj {
		case "users":
			c.HTML(200, "adminUsersCreate", gin.H{
				"Title":           "Создание пользователя",
				"MainActionTitle": "Пользователи",
				"Active":          "users",
				"Mode":            "create",
			})
			break
		case "tariffs":
			c.HTML(200, "adminTariffsCreate", gin.H{
				"Title":           "Создание тарифа",
				"MainActionTitle": "Тарифы",
				"Active":          "tariffs",
				"Mode":            "create",
			})
			break
		case "services":
			c.HTML(200, "adminServicesCreate", gin.H{
				"Title":           "Создание услуги",
				"MainActionTitle": "Услуги",
				"Active":          "services",
				"Mode":            "create",
			})
			break
		case "addresses":
			c.HTML(200, "adminAddressesCreate", gin.H{
				"Title":           "Создание адреса",
				"MainActionTitle": "Список адресов",
				"Active":          "addresses",
				"Mode":            "create",
			})
			break
		//case "seo":
		//	c.HTML(200, "adminSEOCreate", gin.H{
		//		"Title":           "Создание SEO-настроек",
		//		"MainActionTitle": "SEO-настройки",
		//		"Active":          "seo",
		//		"Mode":            "create",
		//	})
		//	break
		case "seo":
			c.HTML(200, "error", gin.H{
				"Title":      "Создание SEO-настроек",
				"StatusCode": 404,
				"StatusText": "Страница не найдена",
			})
			break
		case "faq":
			c.HTML(200, "adminFaqCreate", gin.H{
				"Title":           "Создание FAQ",
				"MainActionTitle": "FAQ",
				"Active":          "faq",
				"Mode":            "create",
			})
			break
		}
	}
}

func adminView(c *gin.Context) {
	session := getSession(c)
	if session.User.Role.ID != 1 {
		c.HTML(200, "error", gin.H{
			"StatusCode": 403,
			"StatusText": "У вас недостаточно прав для просмотра этой страницы",
		})
	} else {
		obj := c.Param("obj")
		objID := c.Param("id")

		switch obj {
		case "users":
			user, e := database.GetUser(objID)
			if e != nil {
				additional.Logger.Println(e)
				return
			}
			c.HTML(200, "adminUsersView", gin.H{
				"Title":           "Просмотр пользователя: ",
				"MainActionTitle": "Пользователи",
				"Active":          "users",
				"Obj":             user,
				"Mode":            "view",
			})
			break
		case "tariffs":
			tariff, e := database.GetTariff(objID)
			if e != nil {
				additional.Logger.Println(e)
				return
			}
			c.HTML(200, "adminTariffsView", gin.H{
				"Title":           "Просмотр тарифа: ",
				"MainActionTitle": "Тарифы",
				"Active":          "tariffs",
				"Obj":             tariff,
				"Mode":            "view",
			})
			break
		case "deposits":
			deposit, e := database.GetDeposit(objID)
			if e != nil {
				additional.Logger.Println(e)
				return
			}
			c.HTML(200, "adminDepositsView", gin.H{
				"Title":           "Просмотр пользователя: ",
				"Obj":             deposit,
				"ByID":            deposit.User.ID,
				"Mode":            "view",
				"SubMode":         "view",
				"SubTitle":        "История пополнения",
				"MainActionTitle": "Пользователи",
				"Active":          "users",
				"SubActive":       "deposits",
			})
			break
		case "expenses":
			expense, e := database.GetExpense(objID)
			if e != nil {
				additional.Logger.Println(e)
				return
			}
			c.HTML(200, "adminExpensesView", gin.H{
				"Title":           "Просмотр пользователя: ",
				"Obj":             expense,
				"ByID":            expense.User.ID,
				"Mode":            "view",
				"SubMode":         "view",
				"SubTitle":        "История списания",
				"MainActionTitle": "Пользователи",
				"Active":          "users",
				"SubActive":       "expenses",
			})
			break
		case "settings":
			setting, e := database.GetSetting(objID)
			if e != nil {
				additional.Logger.Println(e)
				return
			}
			c.HTML(200, "adminSettingsView", gin.H{
				"Title":           "Просмотр настроек: ",
				"MainActionTitle": "Настройки",
				"Active":          "settings",
				"Obj":             setting,
				"Mode":            "view",
			})
			break
		case "services":
			service, e := database.GetService(objID)
			if e != nil {
				additional.Logger.Println(e)
				return
			}
			c.HTML(200, "adminServicesView", gin.H{
				"Title":           "Просмотр услуги: ",
				"MainActionTitle": "Услуги",
				"Active":          "services",
				"Obj":             service,
				"Mode":            "view",
			})
			break
		case "addresses":
			address, e := database.GetAddress(objID)
			if e != nil {
				additional.Logger.Println(e)
				return
			}
			c.HTML(200, "adminAddressesView", gin.H{
				"Title":           "Просмотр адреса: ",
				"MainActionTitle": "Список адресов",
				"Active":          "addresses",
				"Obj":             address,
				"Mode":            "view",
			})
			break
		case "seo":
			seo, e := database.GetSEO(objID)
			if e != nil {
				additional.Logger.Println(e)
				return
			}
			c.HTML(200, "adminSEOView", gin.H{
				"Title":           "Просмотр SEO-настроек: ",
				"MainActionTitle": "SEO-настройки",
				"Active":          "seo",
				"Obj":             seo,
				"Mode":            "view",
			})
			break
		case "faq":
			faq, e := database.GetFaq(objID)
			if e != nil {
				additional.Logger.Println(e)
				return
			}
			c.HTML(200, "adminFaqView", gin.H{
				"Title":           "Просмотр FAQ: ",
				"MainActionTitle": "FAQ",
				"Active":          "faq",
				"Obj":             faq,
				"Mode":            "view",
			})
			break
		}

	}
}

func adminObjsByID(c *gin.Context) {
	session := getSession(c)
	if session.User.Role.ID != 1 {
		c.HTML(200, "error", gin.H{
			"StatusCode": 403,
			"StatusText": "У вас недостаточно прав для просмотра этой страницы",
		})
	} else {
		obj := c.Param("obj")
		byID := c.Param("id")
		switch obj {
		case "deposits":
			deposits, e := database.GetDepositsByID(byID)
			if e != nil {
				additional.Logger.Println(e)
			}

			c.HTML(200, "adminDeposits", gin.H{
				"Title":           "Просмотр пользователя: ",
				"Deposits":        deposits,
				"ByID":            byID,
				"Mode":            "view",
				"SubMode":         "index",
				"SubTitle":        "История пополнения",
				"MainActionTitle": "Пользователи",
				"Active":          "users",
				"SubActive":       "deposits",
			})
			break
		case "expenses":
			expenses, e := database.GetExpensesByID(byID)
			if e != nil {
				additional.Logger.Println(e)
			}

			c.HTML(200, "adminExpenses", gin.H{
				"Title":           "Просмотр пользователя: ",
				"Expanses":        expenses,
				"ByID":            byID,
				"Mode":            "view",
				"SubMode":         "index",
				"SubTitle":        "История списания",
				"MainActionTitle": "Пользователи",
				"Active":          "users",
				"SubActive":       "expenses",
			})
			break
		}
	}
}

func adminObjs(c *gin.Context) {
	session := getSession(c)
	if session.User.Role.ID != 1 {
		c.HTML(200, "error", gin.H{
			"StatusCode": 403,
			"StatusText": "У вас недостаточно прав для просмотра этой страницы",
		})
	} else {
		obj := c.Param("obj")
		switch obj {
		case "users":
			users, e := database.GetAllUsers()
			if e != nil {
				additional.Logger.Println(e)
			}

			c.HTML(200, "adminUsers", gin.H{
				"Title":           "Пользователи",
				"MainActionTitle": "Пользователи",
				"Active":          "users",
				"Users":           users,
				"Mode":            nil,
			})
			break
		case "tariffs":
			tariffs, e := database.GetAllTariffs()
			if e != nil {
				additional.Logger.Println(e)
			}

			c.HTML(200, "adminTariffs", gin.H{
				"Title":           "Тарифы",
				"MainActionTitle": "Тарифы",
				"Active":          "tariffs",
				"Tariffs":         tariffs,
				"Mode":            nil,
			})
			break
		case "settings":
			settings, e := database.GetAllSettings()
			if e != nil {
				additional.Logger.Println(e)
			}

			c.HTML(200, "adminSettings", gin.H{
				"Title":           "Найстройки",
				"MainActionTitle": "Настройки",
				"Active":          "settings",
				"Settings":        settings,
				"Mode":            nil,
			})
			break
		case "services":
			services, e := database.GetAllServices()
			if e != nil {
				additional.Logger.Println(e)
			}

			c.HTML(200, "adminServices", gin.H{
				"Title":           "Услуги",
				"MainActionTitle": "Услуги",
				"Active":          "services",
				"Services":        services,
				"Mode":            nil,
			})
			break
		case "addresses":
			addresses, e := database.GetAllAddresses()
			if e != nil {
				additional.Logger.Println(e)
			}

			c.HTML(200, "adminAddresses", gin.H{
				"Title":           "Список адресов",
				"MainActionTitle": "Список адресов",
				"Active":          "addresses",
				"Addresses":       addresses,
				"Mode":            nil,
			})
			break

		case "seo":
			seo, e := database.GetAllSEO()
			if e != nil {
				additional.Logger.Println(e)
			}

			c.HTML(200, "adminSEO", gin.H{
				"Title":           "SEO-настройки",
				"MainActionTitle": "SEO-настройки",
				"Active":          "seo",
				"SEO":             seo,
				"Mode":            nil,
			})
			break
		case "faq":
			faq, e := database.GetAllFaq()
			if e != nil {
				additional.Logger.Println(e)
			}

			c.HTML(200, "adminFaq", gin.H{
				"Title":           "FAQ",
				"MainActionTitle": "FAQ",
				"Active":          "faq",
				"Faq":             faq,
				"Mode":            nil,
			})
			break
		}
	}
}

func adminIndex(c *gin.Context) {
	session := getSession(c)
	if session.User.Role.ID != 1 {
		c.HTML(200, "authorization", nil)
	} else {
		c.HTML(200, "adminIndex", nil)
	}
}

func userPersonalAccountGetData(c *gin.Context) {
	session := getSession(c)
	if session.User.ID > 0 {
		deposits, e := database.GetDepositsByID(strconv.Itoa(session.User.ID))
		if e != nil {
			additional.Logger.Println(e)
		}
		expenses, e := database.GetExpensesByID(strconv.Itoa(session.User.ID))
		if e != nil {
			additional.Logger.Println(e)
		}
		user, e := database.GetUser(strconv.Itoa(session.User.ID))
		personalAccount := database.PersonalAccount{
			User:     user,
			Deposits: deposits,
			Expenses: expenses,
		}
		c.JSON(200, personalAccount)
	} else {
		c.AbortWithStatus(403)
	}
}

func userPersonalAccount(c *gin.Context) {
	session := getSession(c)
	if session.User.ID == 0 {
		c.HTML(200, "authorization", gin.H{
			"SEO": database.SeoMap["/authorization"],
		})
	} else {
		c.HTML(200, "personal_account", gin.H{
			"User": session.User,
			"SEO":  database.SeoMap["/personal_account"],
		})
	}
}

func userAuthorizationCheck(c *gin.Context) {
	session := sessions.Default(c)
	var user database.User
	e := c.BindJSON(&user)
	if e != nil {
		additional.Logger.Println(e)
		c.JSON(400, false)
		return
	}

	user.Password, e = additional.Encrypt(user.Password)
	if e != nil {
		additional.Logger.Println(e)
		c.JSON(400, false)
		return
	}

	status := user.UserAuthorizationCheck()

	if status {
		hash, ok := database.CreateSession(&user)
		if ok {
			session.Set("SessionSecretKey", hash)
			e = session.Save()
			if e != nil {
				additional.Logger.Println(e)
				c.JSON(500, false)
				return
			}
			c.JSON(200, true)
			return
		}
	}

	c.JSON(400, false)
}

func userAuthorization(c *gin.Context) {
	session := getSession(c)
	if session.User.ID > 0 {
		c.Redirect(301, "/user/personal_account")
	} else {
		c.HTML(200, "authorization", gin.H{
			"SEO": database.SeoMap["/authorization"],
		})
	}
}

func abonentApplication(c *gin.Context) {
	tariffID := c.Param("tariff")

	if tariffID == "0" {
		c.HTML(200, "abonentApplication", gin.H{
			"Tariff":   nil,
			"Settings": database.SettingsMap,
			"SEO":      database.SeoMap["/abonent_application"],
		})
	} else {
		tariff, e := database.GetTariff(tariffID)
		if e != nil {
			additional.Logger.Println(e)
			c.HTML(200, "abonentApplication", gin.H{
				"Tariff":   nil,
				"Settings": database.SettingsMap,
				"SEO":      database.SeoMap["/abonent_application"],
			})
			return
		}
		c.HTML(200, "abonentApplication", gin.H{
			"Tariff":   tariff,
			"Settings": database.SettingsMap,
			"SEO":      database.SeoMap["/abonent_application"],
		})
	}
}

func faq(c *gin.Context) {
	faq, e := database.GetAllFaq()
	if e != nil {
		additional.Logger.Println(e)
		c.HTML(200, "faq", gin.H{
			"Settings": database.SettingsMap,
			"SEO":      database.SeoMap["/faq"],
		})
	}

	c.HTML(200, "faq", gin.H{
		"Settings": database.SettingsMap,
		"Faq":      faq,
		"SEO":      database.SeoMap["/faq"],
	})
}

func oplata(c *gin.Context) {
	c.HTML(200, "oplata", gin.H{
		"Settings": database.SettingsMap,
		"SEO":      database.SeoMap["/oplata"],
	})
}

func tv_manual(c *gin.Context) {
	c.HTML(200, "tv_manual", gin.H{
		"Settings": database.SettingsMap,
		"SEO":      database.SeoMap["/tv-manual"],
	})
}

func routers(c *gin.Context) {
	c.HTML(200, "routers", gin.H{
		"Settings": database.SettingsMap,
		"SEO":      database.SeoMap["/routers"],
	})
}

func virtual_ats(c *gin.Context) {
	c.HTML(200, "virtual_ats", gin.H{
		"Settings": database.SettingsMap,
		"SEO":      database.SeoMap["/virtual-ats"],
	})
}

func about(c *gin.Context) {
	c.HTML(200, "about", gin.H{
		"Settings": database.SettingsMap,
		"SEO":      database.SeoMap["/about"],
	})
}

func tarif_for_home(c *gin.Context) {
	tariffsType2, e := database.GetTariffsByType(2)
	if e != nil {
		additional.Logger.Println(e)
		c.HTML(200, "tv", gin.H{
			"Settings": database.SettingsMap,
			"SEO":      database.SeoMap["/tarif_for_home"],
		})
	}

	c.HTML(200, "tarif_for_home", gin.H{
		"Tariffs":  tariffsType2,
		"Settings": database.SettingsMap,
		"SEO":      database.SeoMap["/tarif_for_home"],
	})
}

func vpn_for_ul(c *gin.Context) {
	c.HTML(200, "vpn_for_ul", gin.H{
		"Settings": database.SettingsMap,
		"SEO":      database.SeoMap["/vpn-for-ul"],
	})
}

func phone_for_ul(c *gin.Context) {
	c.HTML(200, "phone_for_ul", gin.H{
		"Settings": database.SettingsMap,
		"SEO":      database.SeoMap["/phone-for-ul"],
	})
}

func inet_for_ul(c *gin.Context) {
	c.HTML(200, "inet_for_ul", gin.H{
		"Settings": database.SettingsMap,
		"SEO":      database.SeoMap["/inet-for-ul"],
	})
}

func dop_for_home(c *gin.Context) {
	c.HTML(200, "dop_for_home", gin.H{
		"Settings": database.SettingsMap,
		"SEO":      database.SeoMap["/dop_for_home"],
	})
}

func phone(c *gin.Context) {
	c.HTML(200, "phone", gin.H{
		"Settings": database.SettingsMap,
		"SEO":      database.SeoMap["/phone"],
	})
}

func tv(c *gin.Context) {
	tariffsType3, e := database.GetTariffsByType(3)
	if e != nil {
		additional.Logger.Println(e)
		c.HTML(200, "tv", gin.H{
			"Settings": database.SettingsMap,
			"SEO":      database.SeoMap["/tv"],
		})
	}

	c.HTML(200, "tv", gin.H{
		"Tariffs":  tariffsType3,
		"Settings": database.SettingsMap,
		"SEO":      database.SeoMap["/tv"],
	})
}

func inettv(c *gin.Context) {
	tariffsType1, e := database.GetTariffsByType(1)
	if e != nil {
		additional.Logger.Println(e)
		c.HTML(200, "inettv", gin.H{
			"Settings": database.SettingsMap,
			"SEO":      database.SeoMap["/inettv"],
		})
	}

	c.HTML(200, "inettv", gin.H{
		"Tariffs":  tariffsType1,
		"Settings": database.SettingsMap,
		"SEO":      database.SeoMap["/inettv"],
	})
}

func index(c *gin.Context) {
	tariffsType1, e := database.GetTariffsByType(1)
	if e != nil {
		additional.Logger.Println(e)
		c.HTML(200, "index", gin.H{
			"Settings": database.SettingsMap,
			"SEO":      database.SeoMap["/"],
		})
	}

	c.HTML(200, "index", gin.H{
		"Tariffs":  tariffsType1,
		"Settings": database.SettingsMap,
		"SEO":      database.SeoMap["/"],
	})
}

func exit(c *gin.Context) {
	session := sessions.Default(c)
	_session := getSession(c)

	_, ok := session.Get("SessionSecretKey").(string)
	if ok {
		session.Clear()
		_ = session.Save()
		c.SetCookie("hello", "", -1, "/", c.Request.URL.Hostname(), false, true)
		session.Delete("SessionSecretKey")
	}

	database.DeleteSession(_session)

	c.JSON(301, true)
}

func getSession(c *gin.Context) *database.Session {
	_session := sessions.Default(c)

	sessionHash, ok := _session.Get("SessionSecretKey").(string)
	if ok {
		session := database.GetSession(sessionHash)
		if session != nil {
			session.Exists = true
			return session
		}
	}

	return &database.Session{
		Exists: false,
	}
}
