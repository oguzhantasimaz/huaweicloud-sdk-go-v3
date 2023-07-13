package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateValueListRequest Request Object
type CreateValueListRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *CreateValueListRequestBody `json:"body,omitempty"`
}

func (o CreateValueListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateValueListRequest struct{}"
	}

	return strings.Join([]string{"CreateValueListRequest", string(data)}, " ")
}
