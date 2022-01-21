package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin"                 // web framework adapter
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/postgres" // sql driver
	_ "github.com/GoAdminGroup/themes/sword"                         // ui theme

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/gin-gonic/gin"

	"bin/models"
	"bin/pages"
	"bin/tables"
)

func main() {
	startServer()
}



func startServer() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	r := gin.Default()

	template.AddComp(chartjs.NewChart())

	eng := engine.Default()

	if err := eng.AddConfigFromJSON("./config.json").
		AddGenerator("projects", tables.GetProjectsTable).
		AddGenerator("categories", tables.GetCategoriesTable).
		AddGenerator("projects_categories", tables.GetProjectsCategoriesTable).
		Use(r); err != nil {
		panic(err)
	}

	models.Init(eng.PostgresqlConnection())
		
	r.Static("/uploads", "./uploads")

	eng.HTML("GET", "/admin", pages.GetDashBoard)
	eng.HTML("GET", "/admin/form", pages.GetFormContent)
	eng.HTML("GET", "/admin/table", pages.GetTableContent)

	eng.HTMLFile("GET", "/admin/hello", "./html/hello.tmpl", map[string]interface{}{
		"msg": "Hello world",
	})
		
	//	other routes
	//	for json results
	r.GET("/api/projects/all",showApiProjectsPage)
	r.GET("/api/project/first", showProjectFirst)
	r.GET("/api/project/:project_id", getApiProject)

	//	for html templates results
	//r.LoadHTMLGlob("templates/*")
	//r.GET("/projects/all",showProjectsPage)

	//_ = r.Run(":3000")
	srv := &http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	eng.PostgresqlConnection().Close()
}
