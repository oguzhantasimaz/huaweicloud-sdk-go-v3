package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeregisterOrganizationalUnitRequest Request Object
type DeregisterOrganizationalUnitRequest struct {

	// 注册OU ID。
	ManagedOrganizationUnitId string `json:"managed_organization_unit_id"`
}

func (o DeregisterOrganizationalUnitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeregisterOrganizationalUnitRequest struct{}"
	}

	return strings.Join([]string{"DeregisterOrganizationalUnitRequest", string(data)}, " ")
}
