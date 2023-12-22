package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	api "gmall/api/v1"
	"gmall/middleware"
	"gmall/serializer"
	"net/http"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("jimmy's gmall"))
	r.Use(middleware.Cors())
	r.Use(sessions.Sessions("mysession", store))
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, serializer.Response{
				Status: 200,
				Msg:    "success",
				Data:   "Go lang",
			})
		})
		// 用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		// 商品操作
		v1.GET("products", api.ListProducts)
		v1.GET("product/:id", api.ShowProduct)
		v1.POST("products", api.SearchProducts)
		v1.GET("imgs/:id", api.ListProductImg)

		authed := v1.Group("/") // 需要登录保护
		authed.Use(middleware.JWT())
		{
			// 用户操作
			authed.PUT("user", api.UserUpdate)
			authed.POST("avatar", api.UploadAvatar)

			// 收货地址操作
			authed.POST("addresses", api.CreateAddress)
			authed.GET("addresses/:id", api.GetAddress)
			authed.GET("addresses", api.ListAddress)
			authed.PUT("addresses/:id", api.UpdateAddress)
			authed.DELETE("addresses/:id", api.DeleteAddress)

			// 商品操作
			authed.POST("product", api.CreateProduct)
			authed.PUT("product/:id", api.UpdateProduct)
			authed.DELETE("product/:id", api.DeleteProduct)

			// 购物车
			authed.POST("carts", api.CreateCart)
			authed.GET("carts", api.ShowCarts)
			authed.PUT("carts/:id", api.UpdateCart) // 购物车id
			authed.DELETE("carts/:id", api.DeleteCart)
		}
	}
	return r
}
