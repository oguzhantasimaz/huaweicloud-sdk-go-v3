/*
 * CloudIDE
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListStacksByTagRequest struct {
	Tags *[]string `json:"tags,omitempty"`
}

func (o ListStacksByTagRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListStacksByTagRequest", string(data)}, " ")
}
