package config

import (
	// "fmt"
	// "net/http"

	"github.com/jinzhu/gorm"
	//_ "github.com/mattn/go-sqlite3"
	// "github.com/qor/admin"
	// "github.com/qor/qor"
    //"github.com/gotstago/lighthouse/admin"
)

// Create a GORM-backend model
// type User struct {
// 	gorm.Model
// 	Name string
// }

// Create another GORM-backend model
type Entity struct {
	gorm.Model
	Name        string
	Description string
}
// var Design *admin.Admin


// func init() {
// 	// DB, _ := gorm.Open("sqlite3", "demo.db")
// 	// DB.AutoMigrate(&User{}, &Product{})
//     //db.DB.AutoMigrate(&Entity{})

// 	// Initalize
// 	Design = admin.New(&qor.Config{DB: admin.DB})
    
//     Design.SetSiteName("Lighthouse Buy and Sell - Design")

// 	// Create resources from GORM-backend model
// 	//Design.AddResource(&Entity{})
// 	Design.AddResource(&Entity{}, &admin.Config{Menu: []string{"Entity Management"}})

// 	// Add Dashboard
// 	Design.AddMenu(&admin.Menu{Name: "Dashboard", Link: "/design"})

// 	// Register route
// 	//mux := http.NewServeMux()
// 	// a mount to /admin, so visit `/admin` to view the admin interface
	

// 	// fmt.Println("Listening on: 9001")
// 	// http.ListenAndServe(":9001", mux)
// }
