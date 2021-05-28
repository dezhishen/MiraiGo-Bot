package cache

import (
	"testing"
	"time"
)

func TestPutAndGet(t *testing.T) {
	Set("1", "2", 1*time.Minute)
	v, _ := Get("1")
	str := v.(string)
	print(str)
}
