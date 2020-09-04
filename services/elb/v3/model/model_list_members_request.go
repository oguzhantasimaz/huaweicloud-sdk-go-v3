/*
 * ELB
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListMembersRequest struct {
	PoolId              string   `json:"pool_id"`
	Marker              *string  `json:"marker,omitempty"`
	Limit               *int32   `json:"limit,omitempty"`
	PageReverse         *bool    `json:"page_reverse,omitempty"`
	Name                []string `json:"name,omitempty"`
	Weight              []int32  `json:"weight,omitempty"`
	AdminStateUp        *bool    `json:"admin_state_up,omitempty"`
	SubnetCidrId        []string `json:"subnet_cidr_id,omitempty"`
	Address             []string `json:"address,omitempty"`
	ProtocolPort        []int32  `json:"protocol_port,omitempty"`
	Id                  []string `json:"id,omitempty"`
	OperatingStatus     []string `json:"operating_status,omitempty"`
	EnterpriseProjectId []string `json:"enterprise_project_id,omitempty"`
}

func (o ListMembersRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListMembersRequest", string(data)}, " ")
}
