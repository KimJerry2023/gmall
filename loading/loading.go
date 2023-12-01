package loading

import (
	"gmall/pkg/utils"
	"gmall/repository/db/dao"
)

func Loading() {
	dao.InitMySQL()
	utils.InitLog()
}
