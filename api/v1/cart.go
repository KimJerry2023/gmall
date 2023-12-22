package v1

import (
	"github.com/gin-gonic/gin"
	"gmall/consts"
	"gmall/pkg/utils"
	"gmall/service"
)

func CreateCart(c *gin.Context) {
	createCartService := service.CartService{}
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createCartService); err != nil {
		res := createCartService.Create(c.Request.Context(), claim.ID)
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}

// ShowCarts 购物车详细信息
func ShowCarts(c *gin.Context) {
	showCartsService := service.CartService{}
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	res := showCartsService.Show(c.Request.Context(), claim.ID)
	c.JSON(consts.StatusOK, res)
}

// UpdateCart 修改购物车信息
func UpdateCart(c *gin.Context) {
	updateCartService := service.CartService{}
	if err := c.ShouldBind(&updateCartService); err == nil {
		res := updateCartService.Update(c.Request.Context(), c.Param("id"))
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}

// DeleteCart 删除购物车
func DeleteCart(c *gin.Context) {
	deleteCartService := service.CartService{}
	if err := c.ShouldBind(&deleteCartService); err == nil {
		res := deleteCartService.Delete(c.Request.Context())
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}
