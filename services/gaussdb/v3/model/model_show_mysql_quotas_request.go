package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowMysqlQuotasRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。 取值范围：0 - 10000

	Offset *string `json:"offset,omitempty"`
	// 查询记录数。默认为10，不能为负数，最小值为1，最大值为100。

	Limit *string `json:"limit,omitempty"`
	// 企业项目名称。

	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`
}

func (o ShowMysqlQuotasRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowMysqlQuotasRequest struct{}"
	}

	return strings.Join([]string{"ShowMysqlQuotasRequest", string(data)}, " ")
}
