package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteComponentRequest Request Object
type DeleteComponentRequest struct {

	// 应用ID。
	ApplicationId string `json:"application_id"`

	// 组件ID。
	ComponentId string `json:"component_id"`

	// 企业项目ID。  - 创建环境时，环境会绑定企业项目ID。      - 最大长度36字节，带“-”连字符的UUID格式，或者是字符串“0”。     - 该字段不传（或传为字符串“0”）时，则查询默认企业项目下的资源。  > 关于企业项目ID的获取及企业项目特性的详细信息，请参见《[企业管理服务用户指南](https://support.huaweicloud.com/usermanual-em/zh-cn_topic_0126101490.html)》。
	XEnterpriseProjectID *string `json:"X-Enterprise-Project-ID,omitempty"`

	// 环境ID。      - 获取环境ID，通过《[云应用引擎API参考](https://support.huaweicloud.com/api-cae/ListEnvironments.html)》的“获取环境列表”章节获取环境信息。     - 请求响应成功后在响应体的items数组中的一个元素即为一个环境的信息，其中id字段即是环境ID。
	XEnvironmentID string `json:"X-Environment-ID"`
}

func (o DeleteComponentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteComponentRequest struct{}"
	}

	return strings.Join([]string{"DeleteComponentRequest", string(data)}, " ")
}
