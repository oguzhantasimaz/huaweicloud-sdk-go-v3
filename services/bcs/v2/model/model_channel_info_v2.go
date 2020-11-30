/*
 * BCS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type ChannelInfoV2 struct {
	// 通道名
	Name *string `json:"name,omitempty"`
	// 通道中组织名
	OrgNames *[]string `json:"org_names,omitempty"`
	// 通道描述
	Description *string `json:"description,omitempty"`
}

func (o ChannelInfoV2) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ChannelInfoV2", string(data)}, " ")
}
