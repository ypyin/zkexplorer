package main

import (
	"mime"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var extraMimeTypes = map[string]string{
	".icon": "image-x-icon",
	".ttf":  "application/x-font-ttf",
	".woff": "application/x-font-woff",
	".eot":  "application/vnd.ms-fontobject",
	".svg":  "image/svg+xml",
	".html": "text/html; charset-utf-8",
}

func main() {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()
	setupRoutes(router)
	router.Run()
}

func setupRoutes(router *gin.Engine) {
	group := router.Group("/")

	group.GET("/", GetHome)
	group.GET("/static/*path", GetStaticAsset)

	api := group.Group("/api/v1")
	{
		api.GET("/connect", ConnectToZk)
	}

}

func serveStaticAsset(path string, c *gin.Context) {
	data, err := Asset("static" + path)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	ext := filepath.Ext(path)
	contentType := mime.TypeByExtension(ext)
	if contentType == "" {
		contentType = extraMimeTypes[ext]
	}
	if contentType == "" {
		contentType = "text/plain; charset=utf-8"
	}

	c.Data(200, contentType, data)
}

func GetHome(c *gin.Context) {
	serveStaticAsset("/index.html", c)
}

func GetStaticAsset(c *gin.Context) {
	serveStaticAsset(c.Params.ByName("path"), c)
}

func ConnectToZk(c *gin.Context ) {
	serveStaticAsset("/home.html", c)
}

