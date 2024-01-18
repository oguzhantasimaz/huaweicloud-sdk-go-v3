package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowManagedOrganizationalUnitRequest Request Object
type ShowManagedOrganizationalUnitRequest struct {

	// 注册OU ID。
	ManagedOrganizationUnitId string `json:"managed_organization_unit_id"`
}

func (o ShowManagedOrganizationalUnitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowManagedOrganizationalUnitRequest struct{}"
	}

	return strings.Join([]string{"ShowManagedOrganizationalUnitRequest", string(data)}, " ")
}
