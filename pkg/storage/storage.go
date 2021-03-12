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
func Get(bucket, key []byte, decoder func([]byte) error) error {
	return db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		if b == nil {
			return decoder(nil)
		}
		v := b.Get(key)
		return decoder(v)
	})
}

// Put 放入
func Put(bucket, key, value []byte) error {
	return db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists(bucket)
		b := tx.Bucket(bucket)
		err := b.Put(key, []byte(value))
		return err
	})
}

// Delete 删除
func Delete(bucket, key []byte) error {
	return db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists(bucket)
		b := tx.Bucket(bucket)
		err := b.Delete(key)
		return err
	})
}

// GetByPrefix 根据前缀获取
func GetByPrefix(bucket, prefix []byte, decoder func([]byte, []byte) error) error {
	return db.Update(func(tx *bolt.Tx) error {
		// tx.CreateBucketIfNotExists(bucket)
		b := tx.Bucket(bucket)
		if b == nil {
			return nil
		}
		c := b.Cursor()
		for k, v := c.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = c.Next() {
			err := decoder(k, v)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// HasByPrefix 根据前缀判断是否存在
func HasByPrefix(bucket, prefix []byte) (bool, error) {
	var r bool
	err := db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists(bucket)
		c := tx.Bucket(bucket).Cursor()
		for k, _ := c.Seek([]byte(prefix)); k != nil && bytes.HasPrefix(k, []byte(prefix)); k, _ = c.Next() {
			r = true
		}
		return nil
	})
	return r, err
}

// Incr
func Incr(bucket, key []byte, incr int) (int, error) {
	var r int
	err := db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists(bucket)
		b := tx.Bucket(bucket)
		v := b.Get(key)
		if v == nil {
			r = incr
			// v = IntToBytes(incr)
		} else {
			r = BytesToInt(v) + incr
		}
		err := b.Put(key, IntToBytes(r))
		return err
	})
	return r, err
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
