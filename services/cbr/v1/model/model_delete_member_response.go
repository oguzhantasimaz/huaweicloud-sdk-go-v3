/*
 * Cbr
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
type DeleteMemberResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteMemberResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteMemberResponse", string(data)}, " ")
}
