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

// Response Object
type ListProjectTemplatesResponse struct {
	// 模板列表
	Templates *[]ProjectTemplates `json:"templates,omitempty"`
	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListProjectTemplatesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListProjectTemplatesResponse", string(data)}, " ")
}
