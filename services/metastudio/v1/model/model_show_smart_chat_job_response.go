package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSmartChatJobResponse Response Object
type ShowSmartChatJobResponse struct {

	// 数字人智能交互对话任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 数字人智能交互对话任务的状态。 * WAITING: 等待 * PROCESSING: 处理中 * SUCCEED: 成功 * FAILED: 失败 * CANCELED: 取消 * HEARTBEAT: 心跳
	State *ShowSmartChatJobResponseState `json:"state,omitempty"`

	// 数字人智能交互对话时长，单位秒。
	Duration *float32 `json:"duration,omitempty"`

	// 数字人智能交互对话任务开始时间。格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	StartTime *string `json:"start_time,omitempty"`

	// 数字人智能交互对话任务结束时间。格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	EndTime *string `json:"end_time,omitempty"`

	ErrorInfo *ErrorResponse `json:"error_info,omitempty"`

	// 数字人智能交互对话任务创建时间。格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	CreateTime *string `json:"create_time,omitempty"`

	// 数字人智能交互对话任务最后更新时间。格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	LastupdateTime *string `json:"lastupdate_time,omitempty"`

	RtcRoomInfo *RtcRoomInfoList `json:"rtc_room_info,omitempty"`

	ChatSubtitleConfig *ChatSubtitleConfig `json:"chat_subtitle_config,omitempty"`

	VideoConfig *ChatVideoConfigRsp `json:"video_config,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSmartChatJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSmartChatJobResponse struct{}"
	}

	return strings.Join([]string{"ShowSmartChatJobResponse", string(data)}, " ")
}

type ShowSmartChatJobResponseState struct {
	value string
}

type ShowSmartChatJobResponseStateEnum struct {
	WAITING    ShowSmartChatJobResponseState
	PROCESSING ShowSmartChatJobResponseState
	SUCCEED    ShowSmartChatJobResponseState
	FAILED     ShowSmartChatJobResponseState
	CANCELED   ShowSmartChatJobResponseState
	HEARTBEAT  ShowSmartChatJobResponseState
}

func GetShowSmartChatJobResponseStateEnum() ShowSmartChatJobResponseStateEnum {
	return ShowSmartChatJobResponseStateEnum{
		WAITING: ShowSmartChatJobResponseState{
			value: "WAITING",
		},
		PROCESSING: ShowSmartChatJobResponseState{
			value: "PROCESSING",
		},
		SUCCEED: ShowSmartChatJobResponseState{
			value: "SUCCEED",
		},
		FAILED: ShowSmartChatJobResponseState{
			value: "FAILED",
		},
		CANCELED: ShowSmartChatJobResponseState{
			value: "CANCELED",
		},
		HEARTBEAT: ShowSmartChatJobResponseState{
			value: "HEARTBEAT",
		},
	}
}

func (c ShowSmartChatJobResponseState) Value() string {
	return c.value
}

func (c ShowSmartChatJobResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSmartChatJobResponseState) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
