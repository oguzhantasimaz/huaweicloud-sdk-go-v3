package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSmartChatRoomsRequest Request Object
type ListSmartChatRoomsRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。  格式为(YYYYMMDD'T'HHMMSS'Z')。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 第三方用户ID。 > * 不允许输入中文。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	// 偏移量，表示从此偏移量开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 按智能交互对话直播间名称模糊查询。
	RoomName *string `json:"room_name,omitempty"`

	// 按形象名称模糊查询。
	ModelName *string `json:"model_name,omitempty"`

	// 最近智能交互对话任务起始时间。格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	StartTime *string `json:"start_time,omitempty"`

	// 最近智能交互对话任务结束时间。格式遵循：RFC 3339 如\"2021-01-10T10:43:17Z\"。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ListSmartChatRoomsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSmartChatRoomsRequest struct{}"
	}

	return strings.Join([]string{"ListSmartChatRoomsRequest", string(data)}, " ")
}
