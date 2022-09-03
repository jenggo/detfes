package pkg

import (
	"math/rand"
	"net"
	"time"

	"detfes/vars"

	"github.com/bradfitz/gomemcache/memcache"
)

func cacheInvalidKey(input string) bool {
	host := net.JoinHostPort(vars.Config.Memcached.Host, vars.Config.Memcached.Port)

	mc := memcache.New(host)

	if _, err := mc.Get(input); err == memcache.ErrCacheMiss {
		expire := gtr(3600, 604800) // Between 1hrs to 7days
		val := []byte(input)

		_ = mc.Set(&memcache.Item{Key: input, Value: val, Expiration: expire})

		return false
	}

	return true
}

func gtr(min int, max int) int32 {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(max-min+1) + min
	return int32(r)
}
