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
type UpdateMemberResponse struct {
	Member         *MemberResp `json:"member,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o UpdateMemberResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateMemberResponse", string(data)}, " ")
}
