package v1

import (
	"github.com/gin-gonic/gin"
	"gmall/consts"
	"gmall/pkg/utils"
	"gmall/service"
)

// CreateAddress 新增收货地址
func CreateAddress(c *gin.Context) {
	addressService := service.AddressService{}
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&addressService); err != nil {
		res := addressService.Create(c, claim.ID)
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}

// GetAddress 展示某个收货地址
func GetAddress(c *gin.Context) {
	addressService := service.AddressService{}
	res := addressService.Show(c, c.Param("id"))
	c.JSON(consts.StatusOK, res)
}

// ListAddress 展示收货地址
func ListAddress(c *gin.Context) {
	addressService := service.AddressService{}
	claim, _ := utils.ParseToken(c.GetHeader("Authorizaiton"))
	if err := c.ShouldBind(&addressService); err != nil {
		res := addressService.List(c.Request.Context(), claim.ID)
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}

// UpdateAddress 修改收货地址
func UpdateAddress(c *gin.Context) {
	addressService := service.AddressService{}
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&addressService); err == nil {
		res := addressService.Update(c.Request.Context(), claim.ID, c.Param("id"))
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}

// DeleteAddress 删除收获地址
func DeleteAddress(c *gin.Context) {
	addressService := service.AddressService{}
	if err := c.ShouldBind(&addressService); err == nil {
		res := addressService.Delete(c.Request.Context(), c.Param("id"))
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}
