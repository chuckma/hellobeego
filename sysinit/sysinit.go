package sysinit

import (
	"hellobeego/utils"
	"time"
	cache "github.com/patrickmn/go-cache"
	_"hellobeego/utils"
)

func init() {
	// init cache

	utils.Cache=cache.New(60*time.Minute,120*time.Minute)

	initDB()

}
