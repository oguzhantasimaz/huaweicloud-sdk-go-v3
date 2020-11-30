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

type CouchDbInfo struct {
	// couchDB用户名称
	User *string `json:"user,omitempty"`
}

func (o CouchDbInfo) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CouchDbInfo", string(data)}, " ")
}
