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
type UpdateL7ruleResponse struct {
	Rule           *L7ruleResp `json:"rule,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o UpdateL7ruleResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateL7ruleResponse", string(data)}, " ")
}
