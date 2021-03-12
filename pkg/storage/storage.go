package storage

import (
	"bytes"
	"encoding/binary"

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
		tx.CreateBucketIfNotExists([]byte(bucket))
		b := tx.Bucket([]byte(bucket))
		v := b.Get([]byte(key))
		return decoder(string(v))
	})
}

// Put 放入
func Put(bucket, key, value string) error {
	return db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte(bucket))
		b := tx.Bucket([]byte(bucket))
		err := b.Put([]byte(key), []byte(value))
		return err
	})
}

// Delete 删除
func Delete(bucket, key string) error {
	return db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte(bucket))
		b := tx.Bucket([]byte(bucket))
		err := b.Delete([]byte(key))
		return err
	})
}

// GetByPrefix 根据前缀获取
func GetByPrefix(bucket, prefix string, decoder func(string, string) error) error {
	return db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte(bucket))
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

// HasByPrefix 根据前缀判断是否存在
func HasByPrefix(bucket, prefix string) (bool, error) {
	var r bool
	err := db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte(bucket))
		c := tx.Bucket([]byte(bucket)).Cursor()
		for k, _ := c.Seek([]byte(prefix)); k != nil && bytes.HasPrefix(k, []byte(prefix)); k, _ = c.Next() {
			r = true
		}
		return nil
	})
	return r, err
}

// Incr
func Incr(bucket, key string, incr int) (int, error) {
	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte(bucket))
		b := tx.Bucket([]byte(bucket))
		v := b.Get([]byte(key))
		if v == nil {
			v = IntToBytes(incr)
		}
		v = IntToBytes(BytesToInt(v) + incr)
		err := b.Put([]byte(key), v)
		return err
	})
}

//整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}
