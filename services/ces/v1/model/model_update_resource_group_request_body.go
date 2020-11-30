/*
 * CES
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 更新资源分组，请求参数。
type UpdateResourceGroupRequestBody struct {
	// 资源分组的名称；长度为1-128，只能包含0-9/a-z/A-Z/_/-或汉字；如：ResourceGroup-Test01。
	GroupName string `json:"group_name"`
	// 更新资源分组选择一个或者多个资源。
	Resources []CreateResourceGroup `json:"resources"`
}

func (o UpdateResourceGroupRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateResourceGroupRequestBody", string(data)}, " ")
}
