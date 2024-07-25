package model

import (
	"cell_phone_store/request"
	"cell_phone_store/structure"
)

var LabelTable = "labels"

//标签列表
func LabelList() ([]structure.Label,error) {
	var labels []structure.Label
	result := DB.Table(LabelTable).Order("sort DESC").Find(&labels)
	if result.Error != nil {
		return nil, result.Error
	}
	return labels, nil
}

//标签新增
func LabelAdd(labelAdd structure.LabelAdd) error {
	result :=DB.Table(LabelTable).Create(&labelAdd)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//标签删除
func LabelDel(labelDel request.LabelDelRequest) error {
	result :=DB.Table(LabelTable).Where("id = ? ",labelDel.Id).Delete(nil)
	if result.Error != nil {
		return result.Error
	}
	return nil
}