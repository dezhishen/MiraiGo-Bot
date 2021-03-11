package storage

import (
	"bytes"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

func init() {
	database, err := bolt.Open("plugin-storage.db", 0600, nil)
	if err != nil {
		panic(err)
	}
	db = database
}

// Get 获取
func Get(bucket, key string, decoder func(string) error) error {
	return db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		v := b.Get([]byte(key))
		return decoder(string(v))
	})
}

// Put 放入
func Put(bucket, key, value string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		err := b.Put([]byte(key), []byte(value))
		return err
	})
}

// Delete 删除
func Delete(bucket, key string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		err := b.Delete([]byte(key))
		return err
	})
}

// GetByPrefix 根据前缀获取
func GetByPrefix(bucket, prefix string, decoder func(string, string) error) error {
	return db.Update(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte(bucket)).Cursor()
		for k, v := c.Seek([]byte(prefix)); k != nil && bytes.HasPrefix(k, []byte(prefix)); k, v = c.Next() {
			err := decoder(string(k), string(v))
			if err != nil {
				return err
			}
		}
		return nil
	})
}
