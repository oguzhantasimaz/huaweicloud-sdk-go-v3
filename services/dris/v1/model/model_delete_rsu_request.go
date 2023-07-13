package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRsuRequest Request Object
type DeleteRsuRequest struct {

	// \"**参数说明**：实例ID。dris物理实例的唯一标识。获取方法参见[获取Instance-Id](https://support.huaweicloud.com/api-v2x/v2x_04_0030.html)。  **取值范围**：仅支持数字，小写字母和横杠（-）的组合，长度36。\"
	InstanceId *string `json:"Instance-Id,omitempty"`

	// **参数说明**：RSU的唯一标识符，在平台创建RSU时由平台生成。
	RsuId string `json:"rsu_id"`
}

func (o DeleteRsuRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRsuRequest struct{}"
	}

	return strings.Join([]string{"DeleteRsuRequest", string(data)}, " ")
}
