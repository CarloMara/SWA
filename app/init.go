package app

import (
	"encoding/binary"
	"fmt"
	"github.com/revel/revel"
	bolt "go.etcd.io/bbolt"
	"log"
	"reflect"
	// "strconv"
	// "my-app/app/controllers"
)

var (
	// AppVersion revel app version (ldflags)
	AppVersion string

	// BuildTime revel app build-time (ldflags)
	BuildTime string

	//bolt database
	Db *bolt.DB
)

// (type func(*bbolt.Tx)) as type func(*bbolt.Tx) error in argument to db.Update"
// type func(*bbolt.Tx)
func name() {

}
func InitDB() {
	var err error
	Db, err = bolt.Open("/home/carlo/Documents/web_frameworks_comparison/go/revel/src/my-app/db/notes.db", 0666, nil)
	if err != nil {
		log.Fatal("fail to open db")
	}
	fmt.Println(reflect.TypeOf(Db))

	Db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Notes"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	// Db.Update(func(tx *bolt.Tx) error {
	// 	b := tx.Bucket([]byte("Notes"))

	// 	id, _ := b.NextSequence()
	// 	fmt.Println("%s", id)

	// 	err := b.Put(itob(int(id)), []byte("NOTA zero"))
	// 	return err
	// })

	// Db.Update(func(tx *bolt.Tx) error {
	// 	b := tx.Bucket([]byte("Notes"))

	// 	id, _ := b.NextSequence()

	// 	err := b.Put(itob(int(id)), []byte("NOTA zero"))
	// 	return err
	// })

	// if err != nil {
	// 	return err
	// }
	// return nil
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
	revel.OnAppStart(InitDB)
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

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
