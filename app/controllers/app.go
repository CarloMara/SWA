package controllers

import (
	"encoding/binary"
	// "encoding/hex"
	// "fmt"
	"github.com/revel/revel"
	// "github.com/withmandala/go-log"
	bolt "go.etcd.io/bbolt"
	"my-app/app"
	"my-app/app/routes"
	// "reflect"
	"strconv"
	"time"
)

// type Note struct {
// 	id   int
// 	data string
// }

// type Notebook struct {
// 	notes Note
// }

type App struct {
	*revel.Controller
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(ts []byte) uint64 {
	var i uint64
	i = binary.BigEndian.Uint64(ts)
	return i
}

func (c App) List() revel.Result {

	notes := make(map[string]string)

	app.Db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("Notes"))

		b.ForEach(func(k, v []byte) error {
			log := c.Log.New()
			// fmt.Printf("key=%s, value=%s\n", k, v)	reflect.TypeOf(btoi(k))		reflect.TypeOf(strconv.Itoa(btoi(k)))	hex.Dump(k)
			// log.Info("elements ", "[]keyvalue", k, "int64(keyvalue)", btoi(k))
			log.Infof("elements. int64k: %s", btoi(k))
			notes[strconv.FormatInt(int64(btoi(k)), 10)] = string(v)
			// notes[time.Unix(btoi(k), 10)] = string(v)
			return nil
		})
		return nil
	})
	// fmt.Print("map:", notes)
	// notes := map[int]string{00: string("prova"), 01: string("prova2")}

	return c.Render(notes)
}

// this is not the best thing, as that page is almost a static one
func (c App) NewGet() revel.Result {
	return c.Render()
}

func (c App) New() revel.Result {

	ts := time.Now().Unix()

	data := c.Params.Form.Get("Data")

	log := c.Log.New("New")

	app.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Notes"))
		if b == nil {
			log.Error("nope, tx bucket didn't work")
			return nil
		}
		err := b.Put(itob(int(ts)), []byte(data))
		return err
	})
	return c.Redirect(routes.App.List())
}

func (c App) Edit() revel.Result {
	return c.Render()

}

func (c App) Done() revel.Result {
	return c.Render()
}

func (c App) Delate() revel.Result {
	return c.Render()
}
