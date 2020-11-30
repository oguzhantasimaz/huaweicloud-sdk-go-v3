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
type UpdateL7RuleResponse struct {
	// 请求ID。  注：自动生成 。
	RequestId      *string `json:"request_id,omitempty"`
	Rule           *L7Rule `json:"rule,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateL7RuleResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateL7RuleResponse", string(data)}, " ")
}
