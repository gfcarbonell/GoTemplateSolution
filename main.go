package main

import (
	_ "TemplateSolution/docs"
	"TemplateSolution/src/IoC"
	"TemplateSolution/src/Shared/Utils"
	"TemplateSolution/src/TemplateProject/InterfaceAdapters/Gateways/Persistences/SQLServer"
	"TemplateSolution/src/TemplateProject/InterfaceAdapters/Routers"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"log"
	"strconv"
	"time"
)

// @title Swagger Template API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9500
// @BasePath /v1
// @schemes http
func main() {
	// Load Variables
	Utils.LoadSetting(".\\settings")

	// Connect Database
	db := SQLServer.DbSQLServer{}
	defer db.Close()
	cn, error := db.Open()

	if error != nil {
		log.Fatalln(error)
	}

	// Injections Dependency
	ioc := IoC.NewDependencyContainer(cn)

	// Create Routers
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Logging
	date := time.Now().Format("2006-02-01")
	filenameLog := fmt.Sprintf(".\\logs\\%s.log", date)
	Utils.Logging(filenameLog)

	// Routers
	v1 := e.Group("/v1/api/templates")
	Routers.NewTemplateRouter(v1, ioc.NewTemplateContainer().NewTemplateController())
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Run Server
	log.Println("Server listen at http://localhost:" + strconv.Itoa(Utils.SETTING.Server.Port))
	if err := e.Start(":" + strconv.Itoa(Utils.SETTING.Server.Port)); err != nil {
		log.Fatalln(err)
	}
}
