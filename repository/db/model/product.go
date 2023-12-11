package model

import (
	"gmall/repository/cache"
	"gorm.io/gorm"
	"strconv"
)

// Product 商品模型
type Product struct {
	gorm.Model
	Name          string `gorm:"size:255;index"`
	CategoryID    uint   `gorm:"not null"`
	Title         string
	Info          string `gorm:"size:1000"`
	ImgPath       string
	Price         string
	DiscountPrice string
	OnSale        bool `gorm:"default:false"`
	Num           int
	BossID        uint
	BossName      string
	BossAvatar    string
}

// View 获取点击数
func (p *Product) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.ProductViewKey(p.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

// AddView 商品游览
func (p *Product) AddView() {
	// 增加商品点击数
	cache.RedisClient.Incr(cache.ProductViewKey(p.ID))
	// 增加排行点击数
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(p.ID)))
}
