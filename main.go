package main

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"newSite/database"
	"newSite/utils"
)

func main() {
	database.ConnectDB()
	database.CheckAdmin()
	router := gin.Default()
	store := sessions.NewCookieStore([]byte("secretWord"))
	router.Use(sessions.Sessions("session", store))
	router.LoadHTMLGlob("template/*.html")
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
	routerUser.GET("/authorization", userAuthorization)
	routerUser.POST("/authorization_check", userAuthorizationCheck)
	_ = router.Run("localhost:8080")
}

func userPersonalAccount(c *gin.Context) {

	c.HTML(200, "personal_account", nil)
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

	fmt.Println(user)

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
