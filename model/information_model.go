package model

import (
	"cell_phone_store/request"
	"cell_phone_store/structure"
	"encoding/json"
)
var InformationTable = "informations"

//获取资讯列表
func InformationList(InformationListRequest request.InformationListRequest) ([]structure.Information,error) {
	offset := (InformationListRequest.Page - 1) * structure.PageSize
	var informationList []struct {
		Id        int
		MasterMap string
		Title     string
		Text      string
		Images    string `gorm:"column:images;type:jsonb"`
		CreatedAt string
		UpdatedAt string
	}
	var information []structure.Information
	result := DB.Table(InformationTable).Order("created_at DESC").Limit(structure.PageSize).Offset(offset).Find(&informationList)
	if result.Error != nil {
		return nil, result.Error
	}
	for i := range informationList {
		var images []string
		if informationList[i].Images == "" {
			images = []string{}
		} else {
			err := json.Unmarshal([]byte(informationList[i].Images), &images)
			if err != nil {
				images = []string{}
			}
		}
		var informationInfo structure.Information
		informationInfo.Id = informationList[i].Id
		informationInfo.MasterMap = informationList[i].MasterMap
		informationInfo.Title = informationList[i].Title
		informationInfo.Text = informationList[i].Text
		informationInfo.Images = images
		informationInfo.CreatedAt = informationList[i].CreatedAt
		informationInfo.UpdatedAt = informationList[i].UpdatedAt
		information = append(information,informationInfo)
	}
	return information, nil
}

//资讯新增
func InformationAdd(informationAdd structure.InformationAdd) error {
	result := DB.Table(InformationTable).Create(&informationAdd)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//资讯详情
func InformationInfo(informationInfo request.InformationInfoRequest) (structure.Information,error) {
	var informationRecord struct {
		ConfigValue string
		Id        int
		MasterMap string
		Title     string
		Text      string
		Images   string `gorm:"column:images;type:jsonb"`
		CreatedAt string
		UpdatedAt string
	}

	result := DB.Table(InformationTable).Where("id = ?",informationInfo.Id).First(&informationRecord)
	if result.Error != nil {
		return structure.Information{}, result.Error
	}
	var images []string
	if informationRecord.Images == "" {
		images = []string{}
	} else {
		err := json.Unmarshal([]byte(informationRecord.Images), &images)
		if err != nil {
			images = []string{}
		}
	}
	return structure.Information{
		Id: informationRecord.Id,
		MasterMap: informationRecord.MasterMap,
		Title: informationRecord.Title,
		Text: informationRecord.Text,
		Images: images,
		CreatedAt: informationRecord.CreatedAt,
		UpdatedAt: informationRecord.UpdatedAt,
	}, nil
}

//资讯删除
func InformationDel(informationDel request.InformationDelRequest) error {
	result := DB.Table(InformationTable).Where("id = ?",informationDel.Id).Delete(nil)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
