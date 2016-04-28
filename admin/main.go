package main

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qor/admin"
	"github.com/qor/qor"
    //"github.com/gotstago/lighthouse/admin/db"
    "github.com/gotstago/lighthouse/admin/config"
)

// Create a GORM-backend model
type User struct {
	gorm.Model
	Name string
}

// Create another GORM-backend model
type Product struct {
	gorm.Model
	Name        string
	Description string
}
var DB *gorm.DB

func main() {
	DB, _ := gorm.Open("sqlite3", "demo.db")
	DB.AutoMigrate(&User{}, &Product{}, &config.Entity{})

	// Initalize
	Admin := admin.New(&qor.Config{DB: DB})
    
    Admin.SetSiteName("Lighthouse Buy and Sell")

	// Create resources from GORM-backend model
	Admin.AddResource(&User{})
	Admin.AddResource(&Product{}, &admin.Config{Menu: []string{"Product Management"}})

	// Add Dashboard
	Admin.AddMenu(&admin.Menu{Name: "Dashboard", Link: "/admin"})

	// Register route
	mux := http.NewServeMux()
	// a mount to /admin, so visit `/admin` to view the admin interface
	Admin.MountTo("/admin", mux)
    
    
    //initialize design
    Design := admin.New(&qor.Config{DB: DB})
    
    Design.SetSiteName("Lighthouse Buy and Sell - Design")

	// Create resources from GORM-backend model
	//Design.AddResource(&Entity{})
	Design.AddResource(&config.Entity{}, &admin.Config{Menu: []string{"Entity Management"}})

	// Add Dashboard
	Design.AddMenu(&admin.Menu{Name: "Dashboard", Link: "/design"})

    Design.MountTo("/design", mux)

	fmt.Println("Listening on: 9001")
	http.ListenAndServe(":9001", mux)
}
