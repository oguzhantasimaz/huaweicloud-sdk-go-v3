/*
 * IAM
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
type KeystoneShowCatalogResponse struct {
	// 服务目录信息列表。
	Catalog        *[]Catalog `json:"catalog,omitempty"`
	Links          *LinksSelf `json:"links,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o KeystoneShowCatalogResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneShowCatalogResponse", string(data)}, " ")
}
