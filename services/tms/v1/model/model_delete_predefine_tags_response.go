/*
 * TMS
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
type DeletePredefineTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePredefineTagsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeletePredefineTagsResponse", string(data)}, " ")
}
