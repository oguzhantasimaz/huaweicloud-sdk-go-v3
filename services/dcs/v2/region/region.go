package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
)

var (
	AF_SOUTH_1 = region.NewRegion("af-south-1",
		"https://dcs.af-south-1.myhuaweicloud.com")
	CN_NORTH_2 = region.NewRegion("cn-north-2",
		"https://dcs.cn-north-2.myhuaweicloud.com")
	CN_NORTH_4 = region.NewRegion("cn-north-4",
		"https://dcs.cn-north-4.myhuaweicloud.com")
	CN_NORTH_1 = region.NewRegion("cn-north-1",
		"https://dcs.cn-north-1.myhuaweicloud.com")
	CN_EAST_2 = region.NewRegion("cn-east-2",
		"https://dcs.cn-east-2.myhuaweicloud.com")
	CN_EAST_3 = region.NewRegion("cn-east-3",
		"https://dcs.cn-east-3.myhuaweicloud.com")
	CN_SOUTH_1 = region.NewRegion("cn-south-1",
		"https://dcs.cn-south-1.myhuaweicloud.com")
	CN_SOUTH_2 = region.NewRegion("cn-south-2",
		"https://dcs.cn-south-2.myhuaweicloud.com")
	CN_SOUTHWEST_2 = region.NewRegion("cn-southwest-2",
		"https://dcs.cn-southwest-2.myhuaweicloud.com")
	AP_SOUTHEAST_2 = region.NewRegion("ap-southeast-2",
		"https://dcs.ap-southeast-2.myhuaweicloud.com")
	AP_SOUTHEAST_1 = region.NewRegion("ap-southeast-1",
		"https://dcs.ap-southeast-1.myhuaweicloud.com")
	AP_SOUTHEAST_3 = region.NewRegion("ap-southeast-3",
		"https://dcs.ap-southeast-3.myhuaweicloud.com")
	RU_NORTHWEST_2 = region.NewRegion("ru-northwest-2",
		"https://dcs.ru-northwest-2.myhuaweicloud.com")
	SA_BRAZIL_1 = region.NewRegion("sa-brazil-1",
		"https://dcs.sa-brazil-1.myhuaweicloud.com")
	LA_NORTH_2 = region.NewRegion("la-north-2",
		"https://dcs.la-north-2.myhuaweicloud.com")
	LA_SOUTH_2 = region.NewRegion("la-south-2",
		"https://dcs.la-south-2.myhuaweicloud.com")
	CN_NORTH_9 = region.NewRegion("cn-north-9",
		"https://dcs.cn-north-9.myhuaweicloud.com")
	NA_MEXICO_1 = region.NewRegion("na-mexico-1",
		"https://dcs.na-mexico-1.myhuaweicloud.com")
	AE_AD_1 = region.NewRegion("ae-ad-1",
		"https://dcs.ae-ad-1.myhuaweicloud.com")
	TR_WEST_1 = region.NewRegion("tr-west-1",
		"https://dcs.tr-west-1.myhuaweicloud.com")
	AP_SOUTHEAST_4 = region.NewRegion("ap-southeast-4",
		"https://dcs.ap-southeast-4.myhuaweicloud.com")
	EU_WEST_101 = region.NewRegion("eu-west-101",
		"https://dcs.eu-west-101.myhuaweicloud.com")
	ME_EAST_1 = region.NewRegion("me-east-1",
		"https://dcs.me-east-1.myhuaweicloud.com")
)

var staticFields = map[string]*region.Region{
	"af-south-1":     AF_SOUTH_1,
	"cn-north-2":     CN_NORTH_2,
	"cn-north-4":     CN_NORTH_4,
	"cn-north-1":     CN_NORTH_1,
	"cn-east-2":      CN_EAST_2,
	"cn-east-3":      CN_EAST_3,
	"cn-south-1":     CN_SOUTH_1,
	"cn-south-2":     CN_SOUTH_2,
	"cn-southwest-2": CN_SOUTHWEST_2,
	"ap-southeast-2": AP_SOUTHEAST_2,
	"ap-southeast-1": AP_SOUTHEAST_1,
	"ap-southeast-3": AP_SOUTHEAST_3,
	"ru-northwest-2": RU_NORTHWEST_2,
	"sa-brazil-1":    SA_BRAZIL_1,
	"la-north-2":     LA_NORTH_2,
	"la-south-2":     LA_SOUTH_2,
	"cn-north-9":     CN_NORTH_9,
	"na-mexico-1":    NA_MEXICO_1,
	"ae-ad-1":        AE_AD_1,
	"tr-west-1":      TR_WEST_1,
	"ap-southeast-4": AP_SOUTHEAST_4,
	"eu-west-101":    EU_WEST_101,
	"me-east-1":      ME_EAST_1,
}

var provider = region.DefaultProviderChain("DCS")

func ValueOf(regionId string) *region.Region {
	if regionId == "" {
		panic("unexpected empty parameter: regionId")
	}

	reg := provider.GetRegion(regionId)
	if reg != nil {
		return reg
	}

	if _, ok := staticFields[regionId]; ok {
		return staticFields[regionId]
	}
	panic(fmt.Sprintf("unexpected regionId: %s", regionId))
}
