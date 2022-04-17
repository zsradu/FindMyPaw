package app

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/revel/modules"
	"github.com/revel/revel"
	"log"
)

var (
	// AppVersion revel app version (ldflags)
	AppVersion string

	// BuildTime revel app build-time (ldflags)
	BuildTime string
)

// DB object
var DB *gorm.DB

// DB initialization
func initDB() {
	dbInfo, _ := revel.Config.String("db.info")
	db, err := gorm.Open("mysql", dbInfo)
	if err != nil {
		log.Panicf("Failed gorm.Open: %v\n", err)
	}
	DB = db
	// DB.DropTable(&models.User{})
	// DB.CreateTable(&models.User{})
	//DB.DropTable(&models.Animal{})
	//DB.CreateTable(&models.Animal{})
	// DB.DropTable(&models.Location{})

	//var animal0 models.Animal
	//animal0.AnimalType = "Pisică"
	//animal0.Found = false
	//animal0.Breed = "Birmaneză"
	//animal0.FirstX = 44.46420287654894
	//animal0.FirstY = 26.040593950460387
	//animal0.LastX = animal0.FirstX
	//animal0.LastY = animal0.FirstY
	//DB.Create(&animal0)
	//
	//var animal1 models.Animal
	//animal1.AnimalType = "Câine"
	//animal1.Found = false
	//animal1.FirstX = 44.44559714749782
	//animal1.FirstY = 26.040593950460387
	//animal1.LastX = animal1.FirstX
	//animal1.LastY = animal1.FirstY
	//DB.Create(&animal1)
	//
	//var animal2 models.Animal
	//animal2.AnimalType = "Pisică"
	//animal2.Found = false
	//animal2.FirstX = 44.38018967233391
	//animal2.FirstY = 26.106400184624395
	//animal2.LastX = animal2.FirstX
	//animal2.LastY = animal2.FirstY
	//DB.Create(&animal2)

}

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.BeforeAfterFilter,       // Call the before and after filter functions
		revel.ActionInvoker,           // Invoke the action.
	}

	// Register startup functions with OnAppStart
	// revel.DevMode and revel.RunMode only work inside of OnAppStart. See Example Startup Script
	// ( order dependent )
	// revel.OnAppStart(ExampleStartupScript)
	revel.OnAppStart(initDB)
	// revel.OnAppStart(FillCache)
}

// HeaderFilter adds common security headers
// There is a full implementation of a CSRF filter in
// https://github.com/revel/modules/tree/master/csrf
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")
	c.Response.Out.Header().Add("Referrer-Policy", "strict-origin-when-cross-origin")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

//func ExampleStartupScript() {
//	// revel.DevMod and revel.RunMode work here
//	// Use this script to check for dev mode and set dev/prod startup scripts here!
//	if revel.DevMode == true {
//		// Dev mode
//	}
//}
