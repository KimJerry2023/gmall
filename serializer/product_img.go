package serializer

import (
	"gmall/conf"
	"gmall/repository/db/model"
)

type ProductImg struct {
	ProductID uint   `json:"product_id" form:"prodct_id"`
	ImgPath   string `json:"img_path" form:"img_path"`
}

func BuildProductImg(item *model.ProductImg) ProductImg {
	pImg := ProductImg{
		ProductID: item.ProductID,
		ImgPath:   conf.PhotoHost + conf.HttpPort + conf.ProductPhotoPath + item.ImgPath,
	}
	return pImg
}

func BuildProductImgs(items []*model.ProductImg) (productImgs []ProductImg) {
	for _, item := range items {
		product := BuildProductImg(item)
		productImgs = append(productImgs, product)
	}
	return productImgs
}
