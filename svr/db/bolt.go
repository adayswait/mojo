package db

// import "fmt"
import "log"
import "time"
import "github.com/boltdb/bolt"
import "github.com/adayswait/mojo/global"

var localDB *bolt.DB = nil

func init() {
	initDb()
}

func initDb() {
	var err error
	localDB, err = bolt.Open("mojo.db", 0600, &bolt.Options{Timeout: 5 * time.Second})
	if err != nil {
		log.Fatal(err)
		return
	}
	if localDB != nil {
		createBucket(global.BUCKET_TOKEN_INFO)
		createBucket(global.BUCKET_USR_PASSWD)
		createBucket(global.BUCKET_USER_TOKEN)
	}
}

func createBucket(bucketName string) error {
	err := localDB.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte(bucketName))
		return nil
	})
	return err
}

func boltSet(bucketName, key, value string) error {
	err := localDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		e := b.Put([]byte(key), []byte(value))
		return e
	})
	return err
}
func boltGet(bucketName, key string) ([]byte, error) {
	var value []byte
	err := localDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		v := b.Get([]byte(key))
		value = make([]byte, len(v))
		copy(value, v)
		return nil
	})
	return value, err
}
func boltDelete(bucketName, key string) error {
	err := localDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		e := b.Delete([]byte(key))
		return e
	})
	return err
}
