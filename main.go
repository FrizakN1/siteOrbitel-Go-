package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"newSite/database"
	"newSite/utils"
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
	router.GET("/", index)
	router.GET("/tarif_for_home", tarif_for_home)
	router.GET("/inettv", inettv)
	router.GET("/tv", tv)
	router.GET("/phone", phone)
	router.GET("/dop_for_home", dop_for_home)
	router.GET("/business/inet-for-ul", inet_for_ul)
	router.GET("/business/phone-for-ul", phone_for_ul)
	router.GET("/business/vpn-for-ul", vpn_for_ul)
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
	routerAdmin.GET("/:obj/create", adminCreate)
	routerAdmin.POST("/:obj/create", adminCreatePOST)

	_ = router.Run("192.168.0.105:8080")
}

func adminCreatePOST(c *gin.Context) {
	session := getSession(c)
	if session.User.Role.ID != 1 {
		c.Redirect(301, "/admin/authorization")
	} else {
		obj := c.Param("obj")
		switch obj {
		case "users":
			var user database.User
			e := c.BindJSON(&user)
			if e != nil {
				utils.Logger.Println(e)
				c.Status(400)
				return
			}
			e = user.CreateUser()
			if e != nil {
				utils.Logger.Println(e)
				c.Status(400)
				return
			}
			c.Status(200)
			break
		case "tariffs":
			var tariff database.Tariff
			e := c.BindJSON(&tariff)
			if e != nil {
				utils.Logger.Println(e)
				c.Status(400)
				return
			}
			e = tariff.CreateTariff()
			if e != nil {
				utils.Logger.Println(e)
				c.Status(400)
				return
			}
			c.Status(200)
			break
		}
	}
}

func adminEdit(c *gin.Context) {
	session := getSession(c)
	if session.User.Role.ID != 1 {
		c.Redirect(301, "/admin/authorization")
	} else {
		obj := c.Param("obj")
		objID := c.Param("id")
		switch obj {
		case "users":
			user, e := database.GetUser(objID)
			if e != nil {
				utils.Logger.Println(e)
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
				utils.Logger.Println(e)
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
				utils.Logger.Println(e)
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
		}
	}
}

func adminCreate(c *gin.Context) {
	session := getSession(c)
	if session.User.Role.ID != 1 {
		c.Redirect(301, "/admin/authorization")
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
		}
	}
}

func adminView(c *gin.Context) {
	session := getSession(c)
	if session.User.Role.ID != 1 {
		c.Redirect(301, "/admin/authorization")
	} else {
		obj := c.Param("obj")
		objID := c.Param("id")

		switch obj {
		case "users":
			user, e := database.GetUser(objID)
			if e != nil {
				utils.Logger.Println(e)
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
				utils.Logger.Println(e)
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
				utils.Logger.Println(e)
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
				utils.Logger.Println(e)
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
				utils.Logger.Println(e)
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
		}

	}
}

func adminObjsByID(c *gin.Context) {
	session := getSession(c)
	if session.User.Role.ID != 1 {
		c.Redirect(301, "/admin/authorization")
	} else {
		obj := c.Param("obj")
		byID := c.Param("id")
		switch obj {
		case "deposits":
			deposits, e := database.GetDepositsByID(byID)
			if e != nil {
				utils.Logger.Println(e)
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
				utils.Logger.Println(e)
			}

			c.HTML(200, "adminExpenses", gin.H{
				"Title":           "Просмотр пользователя: ",
				"Deposits":        expenses,
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
		c.Redirect(301, "/admin/authorization")
	} else {
		obj := c.Param("obj")
		switch obj {
		case "users":
			users, e := database.GetAllUsers()
			if e != nil {
				utils.Logger.Println(e)
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
				utils.Logger.Println(e)
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
				utils.Logger.Println(e)
			}

			c.HTML(200, "adminSettings", gin.H{
				"Title":           "Найстройки",
				"MainActionTitle": "Настройки",
				"Active":          "settings",
				"Settings":        settings,
				"Mode":            nil,
			})
			break
		}
	}
}

func adminIndex(c *gin.Context) {
	c.HTML(200, "adminIndex", nil)
}

func userPersonalAccountGetData(c *gin.Context) {
	session := getSession(c)
	if session.User.ID > 0 {
		deposits, e := database.GetDepositsByID(strconv.Itoa(session.User.ID))
		if e != nil {
			utils.Logger.Println(e)
		}
		expenses, e := database.GetExpensesByID(strconv.Itoa(session.User.ID))
		if e != nil {
			utils.Logger.Println(e)
		}
		personalAccount := database.PersonalAccount{
			User:     session.User,
			Deposits: deposits,
			Expenses: expenses,
		}
		c.JSON(200, personalAccount)
	} else {
		c.Status(403)
	}
}

func userPersonalAccount(c *gin.Context) {
	session := getSession(c)
	if session.User.ID > 0 {
		c.HTML(200, "personal_account", nil)
	} else {
		c.Redirect(301, "/user/authorization")
	}
}

func userAuthorizationCheck(c *gin.Context) {
	session := sessions.Default(c)
	var user database.User
	e := c.BindJSON(&user)
	if e != nil {
		utils.Logger.Println(e)
		c.Status(400)
		return
	}

	user.Password, e = utils.Encrypt(user.Password)
	if e != nil {
		utils.Logger.Println(e)
		c.Status(400)
		return
	}

	status := user.UserAuthorizationCheck()

	if status {
		hash, ok := database.CreateSession(&user)
		if ok {
			session.Set("SessionSecretKey", hash)
			e = session.Save()
			if e != nil {
				utils.Logger.Println(e)
				return
			}
			c.JSON(200, true)
		}
	}

	c.Status(400)
}

func userAuthorization(c *gin.Context) {
	c.HTML(200, "authorization", nil)
}

func tarif_for_home(c *gin.Context) {
	c.HTML(200, "tarif_for_home", nil)
}

func vpn_for_ul(c *gin.Context) {
	c.HTML(200, "vpn_for_ul", nil)
}

func phone_for_ul(c *gin.Context) {
	c.HTML(200, "phone_for_ul", nil)
}

func inet_for_ul(c *gin.Context) {
	c.HTML(200, "inet_for_ul", nil)
}

func dop_for_home(c *gin.Context) {
	c.HTML(200, "dop_for_home", nil)
}

func phone(c *gin.Context) {
	c.HTML(200, "phone", nil)
}

func tv(c *gin.Context) {
	c.HTML(200, "tv", nil)
}

func inettv(c *gin.Context) {
	c.HTML(200, "inettv", nil)
}

func index(c *gin.Context) {
	c.HTML(200, "index", nil)
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
