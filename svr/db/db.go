package db

// import "fmt"
import "log"
import "time"
import "github.com/boltdb/bolt"
import "github.com/adayswait/mojo/global"

var db *bolt.DB = nil

func init() {
	initDb()
}

func initDb() {
	var err error
	db, err = bolt.Open("mojo.db", 0600, &bolt.Options{Timeout: 5 * time.Second})
	if err != nil {
		log.Fatal(err)
		return
	}
	if db != nil {
		createBucket(global.BUCKET_USER)
	}
}

func createBucket(bucketName string) error {
	err := db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte(bucketName))
		return nil
	})
	return err
}

func Set(bucketName, key, value string) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		e := b.Put([]byte(key), []byte(value))
		return e
	})
	return err
}
func Get(bucketName, key string) ([]byte, error) {
	var value []byte
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		v := b.Get([]byte(key))
		value = make([]byte, len(v))
		copy(value, v)
		return nil
	})
	return value, err
}
func Delete(bucketName, key string) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		e := b.Delete([]byte(key))
		return e
	})
	return err
}
