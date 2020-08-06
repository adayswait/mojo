package db

import "fmt"
import "log"
import "time"
import "strconv"
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
		createBucket(global.BUCKET_USR_OPSLOG)
		createBucket(global.BUCKET_OPS_MACINI)
		createBucket(global.BUCKET_OPS_DEPINI)
		createBucket(global.BUCKET_OPS_DEVINI)
		createBucket(global.BUCKET_ITEMS_DESC)
		createBucket(global.BUCKET_OPS_DEPBIL)
	}
}

func createBucket(bucketName string) error {
	rwLock.Lock()
	err := localDB.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte(bucketName))
		return nil
	})
	defer rwLock.Unlock()
	return err
}

func boltSet(bucketName, key, value string) error {
	rwLock.Lock()
	err := localDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return fmt.Errorf("bucket:%s is nil", bucketName)
		}
		if len(key) == 0 {
			id, _ := b.NextSequence()
			key = strconv.FormatUint(id, 10)
		}
		e := b.Put([]byte(key), []byte(value))
		return e
	})
	defer rwLock.Unlock()
	return err
}
func boltGet(bucketName, key string) ([]byte, error) {
	var value []byte
	rwLock.RLock()
	err := localDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return fmt.Errorf("bucket:%s is nil", bucketName)
		}
		v := b.Get([]byte(key))
		value = make([]byte, len(v))
		copy(value, v)
		return nil
	})
	defer rwLock.RUnlock()
	return value, err
}
func boltDelete(bucketName, key string) error {
	rwLock.Lock()
	err := localDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return fmt.Errorf("bucket:%s is nil", bucketName)
		}
		e := b.Delete([]byte(key))
		return e
	})
	defer rwLock.Unlock()
	return err
}
