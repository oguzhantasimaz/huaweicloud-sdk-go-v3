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

// Response Object
type CreateIpGroupResponse struct {
	Ipgroup *IpGroup `json:"ipgroup,omitempty"`
	// 请求ID。  注：自动生成 。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateIpGroupResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateIpGroupResponse", string(data)}, " ")
}
