package model

import (
	"cell_phone_store/request"
	"cell_phone_store/structure"
	"encoding/json"
)

var ProductTable = "products"
//产品列表
func ProductList(request request.ProductListRequest) ([]structure.Product,error){
	//offset := (request.Page - 1) * structure.PageSize
	var productList []struct{
		Id        int
		Desc 	  string
		Images    string `gorm:"column:images;type:jsonb"`
		Sort      int
		LabelID   int
		CreatedAt string
		UpdatedAt string
	}
	var products []structure.Product
	db := DB.Table(ProductTable).Model(&structure.Product{})
	db = db.Where("label_id = ?", request.LabelId)
	if request.Desc != "" {
		likeValue := "%" + request.Desc + "%"
		db = db.Where("`desc` LIKE ?", likeValue)
	}
	db = db.Order("`sort` DESC")
	if err := db.Find(&productList).Error; err != nil {
		return nil, err
	}
	for i :=range productList{
		var images []string
		if productList[i].Images == "" {
			images = []string{}
		}else{
			err :=json.Unmarshal([]byte(productList[i].Images),&images)
			if err !=nil {
				images = []string{}
			}
		}
		var productInfo structure.Product
		productInfo.Id = productList[i].Id
		productInfo.Desc = productList[i].Desc
		productInfo.Images = images
		productInfo.Sort = productList[i].Sort
		productInfo.LabelID = productList[i].LabelID
		productInfo.CreatedAt = productList[i].CreatedAt
		productInfo.UpdatedAt = productList[i].UpdatedAt
		products = append(products,productInfo)
	}
	return products, nil
}

//产品新增
func ProductAdd(productAdd structure.ProductAdd) error {
	result :=DB.Table(ProductTable).Create(&productAdd)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//产品详情
func ProductInfo(productInfo request.ProductInfoRequest)(structure.Product,error)  {
	var productRecord struct{
		Id        int
		Desc 	  string
		Images    string `gorm:"column:images;type:jsonb"`
		Sort      int
		LabelID   int
		CreatedAt string
		UpdatedAt string
	}
	result := DB.Table(ProductTable).Where("id = ?",productInfo.Id).First(&productRecord)
	if result.Error != nil {
		return structure.Product{}, result.Error
	}
	var images []string
	if productRecord.Images == "" {
		images = []string{}
	}else{
		err :=json.Unmarshal([]byte(productRecord.Images),&images)
		if err !=nil {
			images = []string{}
		}
	}
	return structure.Product{
		Id: productRecord.Id,
		Desc: productRecord.Desc,
		Images: images,
		Sort: productRecord.Sort,
		LabelID: productRecord.LabelID,
		CreatedAt: productRecord.CreatedAt,
		UpdatedAt: productRecord.UpdatedAt,
	}, nil
}

//产品删除
func ProductDel(productDel request.ProductDelRequest) error {
	result := DB.Table(ProductTable).Where("id = ?",productDel.Id).Delete(nil)
	if result.Error != nil {
		return result.Error
	}
	return nil
}