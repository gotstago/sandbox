package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gotstago/activity"
	"github.com/gotstago/sandbox/admin/app/models"
	"github.com/gotstago/sandbox/admin/config/routes"
	"github.com/gotstago/sandbox/admin/model"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/qor/worker"
	// "github.com/qor/render"
)

var DB *gorm.DB
var DataTypes = []string{"uint", "int", "string"}
var Design *admin.Admin

func main() {
	DB, _ := gorm.Open("sqlite3", "demo.db")
	//l10n.RegisterCallbacks(&DB)
	DB.AutoMigrate(&models.User{}, &models.Product{}, &model.Entity{}, &model.EntityAttribute{}, &model.AttributeType{})
	//AutoMigrate(&activity.QorActivity{})
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
	Design = admin.New(&qor.Config{DB: DB})

	Design.SetSiteName("Lighthouse Buy and Sell - Design")

	// Create resources from GORM-backend model
	//Design.AddResource(&Entity{})
	entity := Design.AddResource(&model.Entity{}, &admin.Config{Menu: []string{"Entity Management"}})
	// Register Activity for Entity resource
	activity.Register(entity)

	entity.Action(&admin.Action{
		Name: "Disable",
		Handle: func(arg *admin.ActionArgument) error {
			for _, record := range arg.FindSelectedRecords() {
				arg.Context.DB.Model(record.(*model.Entity)).Update("enabled", false)
			}
			return nil
		},
		Visible: func(record interface{}, context *admin.Context) bool {
			if entity, ok := record.(*model.Entity); ok {
				return entity.Enabled == true
			}
			return true
		},
		Modes: []string{"index", "edit", "menu_item"},
		// Method: "POST",
	})

	entity.Action(&admin.Action{
		Name: "Enable",
		Handle: func(arg *admin.ActionArgument) error {
			for _, record := range arg.FindSelectedRecords() {
				arg.Context.DB.Model(record.(*model.Entity)).Update("enabled", true)
			}
			return nil
		},
		Visible: func(record interface{}, context *admin.Context) bool {
			if entity, ok := record.(*model.Entity); ok {
				return entity.Enabled == false
			}
			return true
		},
		Modes: []string{"index", "edit", "menu_item"},
	})
	Design.AddResource(&model.AttributeType{}, &admin.Config{Menu: []string{"Lookups"}})
	//entityAttributeMeta := entity.Meta(&admin.Meta{Name: "EntityAttributes"})
	//entityAttributeMeta.Resource.Meta(&admin.Meta{Name: "DataType", Type: "select_one", Collection: DataTypes})
	//entity.Meta(&admin.Meta{Name: "DataType", Type: "select_one", Collection: DataTypes})
	//Design.Meta(&admin.Meta{Name: "DataType", Type: "select_one", Collection: DataTypes})

	// Add Dashboard
	Design.AddMenu(&admin.Menu{Name: "Dashboard", Link: "/design"})

	Design.MountTo("/design", mux)

	// Add Worker
	Worker := worker.New(worker.Config{Admin:Design})
	//getWorker()

	type sendNewsletterArgument struct {
		Subject      string
		Content      string `sql:"size:65532"`
		SendPassword string
		worker.Schedule
	}
    
    Worker.RegisterJob(&worker.Job{
		Name: "Send Newsletter",
		Handler: func(argument interface{}, qorJob worker.QorJobInterface) error {
			qorJob.AddLog("Started sending newsletters...")
			qorJob.AddLog(fmt.Sprintf("Argument: %+v", argument.(*sendNewsletterArgument)))
			for i := 1; i <= 100; i++ {
				time.Sleep(100 * time.Millisecond)
				qorJob.AddLog(fmt.Sprintf("Sending newsletter %v...", i))
				qorJob.SetProgress(uint(i))
			}
			qorJob.AddLog("Finished send newsletters")
			return nil
		},
		Resource: Design.NewResource(&sendNewsletterArgument{}),
	})

	Design.AddResource(Worker)

	//config.View.Layout

	fmt.Println("Listening on: 9001")
	//http.ListenAndServe(":9001", mux)
	if err := http.ListenAndServe(":9001", mux); err != nil {
		panic(err)
	}
}

// func getWorker() *worker.Worker {
// 	Worker := worker.New()

// 	type sendNewsletterArgument struct {
// 		Subject      string
// 		Content      string `sql:"size:65532"`
// 		SendPassword string
// 		worker.Schedule
// 	}

// 	Worker.RegisterJob(&worker.Job{
// 		Name: "Send Newsletter",
// 		Handler: func(argument interface{}, qorJob worker.QorJobInterface) error {
// 			qorJob.AddLog("Started sending newsletters...")
// 			qorJob.AddLog(fmt.Sprintf("Argument: %+v", argument.(*sendNewsletterArgument)))
// 			for i := 1; i <= 100; i++ {
// 				time.Sleep(100 * time.Millisecond)
// 				qorJob.AddLog(fmt.Sprintf("Sending newsletter %v...", i))
// 				qorJob.SetProgress(uint(i))
// 			}
// 			qorJob.AddLog("Finished send newsletters")
// 			return nil
// 		},
// 		Resource: Design.NewResource(&sendNewsletterArgument{}),
// 	})

// 	return Worker
// }
