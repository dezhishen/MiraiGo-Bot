package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var c = cache.New(5*time.Minute, 10*time.Minute)

// Get 获取
func Get(key string) (interface{}, bool) {
	v, exists := c.Get(key)
	if exists {
		return *v.(*interface{}), true
	} else {
		return nil, false
	}
}

// Set 设置
func Set(key string, value interface{}, expiration time.Duration) {
	c.Set(key, &value, expiration)
}

// Delete 删除
func Delete(key string) {
	c.Delete(key)
}
