/*
 * EPS
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
type DisableEnterpriseProjectResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DisableEnterpriseProjectResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DisableEnterpriseProjectResponse", string(data)}, " ")
}
