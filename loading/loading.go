package loading

import (
	"gmall/pkg/utils"
	"gmall/repository/cache"
	"gmall/repository/db/dao"
)

func Loading() {
	dao.InitMySQL()
	cache.InitCache()
	utils.InitLog()
}
