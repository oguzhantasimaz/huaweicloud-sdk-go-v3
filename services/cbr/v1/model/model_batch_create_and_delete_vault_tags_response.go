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
type BatchCreateAndDeleteVaultTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateAndDeleteVaultTagsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchCreateAndDeleteVaultTagsResponse", string(data)}, " ")
}
