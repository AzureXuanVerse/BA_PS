package sdk

import (
	"fmt"
	"github.com/gucooing/BaPs/config"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type SDK struct {
	router     *gin.Engine
	serverinfo *ServerInfo
}

func New(router *gin.Engine) *SDK {
	s := &SDK{
		router: router,
		serverinfo: &ServerInfo{
			loadSync: sync.RWMutex{},
		},
	}

	s.initRouter()
	return s
}

func (s *SDK) initRouter() {
	s.router.LoadHTMLGlob(fmt.Sprintf("%s/templates/*", config.GetConfig().DataPath))
	s.router.Any("/", handleIndex)
	s.router.Any("/index", handleIndex)
	s.router.GET("/r:url", s.connectionGroups)

	s.router.GET("/prod/index.json", index)

	account := s.router.Group("/account")
	{
		account.POST("/yostar_auth_request", s.YostarAuthRequest)
		account.POST("/yostar_auth_submit", s.YostarAuthSubmit)
	}
	user := s.router.Group("/user")
	{
		user.POST("/yostar_createlogin", s.YostarCreatelogin)
		user.POST("/login", s.YostarLogin)
		user.Any("/agreement", agreement)
	}
	app := s.router.Group("/app")
	{
		app.Any("/getSettings", getSettings)
		app.Any("/getCode", getCode)
	}
	hash := s.router.Group("/r77_8q5tn5v8489fubab84a8")
	{
		hash.GET("/TableBundles/TableCatalog.hash", func(c *gin.Context) {
			c.String(200, "0")
		})
		hash.GET("/MediaResources/Catalog/MediaCatalog.hash", func(c *gin.Context) {
			c.String(200, "0")
		})
		hash.GET("/iOS/bundleDownloadInfo.hash", func(c *gin.Context) {
			c.String(200, "0")
		})
		hash.GET("/iOS/catalog_iOS.hash", func(c *gin.Context) {
			c.String(200, "0")
		})
		hash.GET("/TableBundles/TableCatalog.bytes", func(c *gin.Context) {
			c.String(200, "0")
		})
	}
}

func handleIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":  "Ba Ps!",
		"github": "https://github.com/gucooing/BaPs",
	})
}
