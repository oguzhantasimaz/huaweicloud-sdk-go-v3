package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
	"sort"
	"strings"
)

var (
	CN_NORTH_4 = region.NewRegion("cn-north-4",
		"https://image.cn-north-4.myhuaweicloud.com",
		"https://image.cn-north-4.myhuaweicloud.cn")
	CN_NORTH_1 = region.NewRegion("cn-north-1",
		"https://image.cn-north-1.myhuaweicloud.com",
		"https://image.cn-north-1.myhuaweicloud.cn")
	AP_SOUTHEAST_1 = region.NewRegion("ap-southeast-1",
		"https://image.ap-southeast-1.myhuaweicloud.com",
		"https://image.ap-southeast-1.myhuaweicloud.cn")
	AP_SOUTHEAST_3 = region.NewRegion("ap-southeast-3",
		"https://image.ap-southeast-3.myhuaweicloud.com",
		"https://image.ap-southeast-3.myhuaweicloud.cn")
	CN_EAST_3 = region.NewRegion("cn-east-3",
		"https://image.cn-east-3.myhuaweicloud.com",
		"https://image.cn-east-3.myhuaweicloud.cn")
)

var staticFields = map[string]*region.Region{
	"cn-north-4":     CN_NORTH_4,
	"cn-north-1":     CN_NORTH_1,
	"ap-southeast-1": AP_SOUTHEAST_1,
	"ap-southeast-3": AP_SOUTHEAST_3,
	"cn-east-3":      CN_EAST_3,
}

var provider = region.DefaultProviderChain("IMAGE")

func getRegionIds() []string {
	ids := make([]string, 0, len(staticFields))
	for key := range staticFields {
		ids = append(ids, key)
	}
	sort.Strings(ids)
	return ids
}

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
	panic(fmt.Sprintf("region id '%s' is not in the following supported regions of service 'Image': [%s]", regionId, strings.Join(getRegionIds(), ", ")))
}
