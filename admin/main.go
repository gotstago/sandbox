package main

import (
	"fmt"
	"net/http"

	"github.com/gotstago/sandbox/admin/model"
    "github.com/gotstago/sandbox/admin/app/models"
    "github.com/gotstago/sandbox/admin/config/routes"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qor/admin"
	"github.com/qor/qor"
	// "github.com/qor/render"
)

var DB *gorm.DB
var DataTypes = []string{"uint", "int", "string"}

func main() {
	DB, _ := gorm.Open("sqlite3", "demo.db")
    //l10n.RegisterCallbacks(&DB)
	DB.AutoMigrate(&models.User{}, &models.Product{}, &model.Entity{}, &model.EntityAttribute{})

	// Register route
	mux := http.NewServeMux()
    
    mux.Handle("/", routes.Router())
	
    // Initalize
	Admin := admin.New(&qor.Config{DB: DB})

	Admin.SetSiteName("Lighthouse Buy and Sell")

	// Create resources from GORM-backend model
	Admin.AddResource(&models.User{})
	Admin.AddResource(&models.Product{}, &admin.Config{Menu: []string{"Product Management"}})

	// Add Dashboard
	Admin.AddMenu(&admin.Menu{Name: "Dashboard", Link: "/admin"})

	// a mount to /admin, so visit `/admin` to view the admin interface
	Admin.MountTo("/admin", mux)

	//initialize design
	Design := admin.New(&qor.Config{DB: DB})

	Design.SetSiteName("Lighthouse Buy and Sell - Design")

	// Create resources from GORM-backend model
	//Design.AddResource(&Entity{})
	entity := Design.AddResource(&model.Entity{}, &admin.Config{Menu: []string{"Entity Management"}})
    entityAttributeMeta := entity.Meta(&admin.Meta{Name: "EntityAttributes"})
	entityAttributeMeta.Resource.Meta(&admin.Meta{Name: "DataType", Type: "select_one", Collection: DataTypes})
    //entity.Meta(&admin.Meta{Name: "DataType", Type: "select_one", Collection: DataTypes})
    //Design.Meta(&admin.Meta{Name: "DataType", Type: "select_one", Collection: DataTypes})

	// Add Dashboard
	Design.AddMenu(&admin.Menu{Name: "Dashboard", Link: "/design"})

	Design.MountTo("/design", mux)
    
    //config.View.Layout

	fmt.Println("Listening on: 9001")
	//http.ListenAndServe(":9001", mux)
	if err := http.ListenAndServe(":9001", mux); err != nil {
		panic(err)
	}
}
