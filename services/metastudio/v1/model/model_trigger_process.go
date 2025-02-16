package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TriggerProcess 触发器处理
type TriggerProcess struct {

	// 处理抑制时长。单位秒。 -1 表示整场直播 0 表示无抑制，每次都触发
	TimeWindow *int32 `json:"time_window,omitempty"`

	// 回复类型。 * SYSTEM_REPLY：系统自动回复设置的话术。 * CALLBACK：回调给其他服务，携带设置的话术。 * SHOW_LAYER: 显示叠加图层，不影响话术。
	ReplyMode *TriggerProcessReplyMode `json:"reply_mode,omitempty"`

	LayerConfig *SmartLayerConfig `json:"layer_config,omitempty"`

	// 回复话术集
	ReplyTexts *[]string `json:"reply_texts,omitempty"`

	// 回复音频集。填写audio_url。
	ReplyAudios *[]ReplyAudioInfo `json:"reply_audios,omitempty"`

	// 回复次序 - RANDOM：随机 - ORDER：顺序循环
	ReplyOrder *TriggerProcessReplyOrder `json:"reply_order,omitempty"`
}

func (o TriggerProcess) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TriggerProcess struct{}"
	}

	return strings.Join([]string{"TriggerProcess", string(data)}, " ")
}

type TriggerProcessReplyMode struct {
	value string
}

type TriggerProcessReplyModeEnum struct {
	SYSTEM_REPLY TriggerProcessReplyMode
	CALLBACK     TriggerProcessReplyMode
	SHOW_LAYER   TriggerProcessReplyMode
}

func GetTriggerProcessReplyModeEnum() TriggerProcessReplyModeEnum {
	return TriggerProcessReplyModeEnum{
		SYSTEM_REPLY: TriggerProcessReplyMode{
			value: "SYSTEM_REPLY",
		},
		CALLBACK: TriggerProcessReplyMode{
			value: "CALLBACK",
		},
		SHOW_LAYER: TriggerProcessReplyMode{
			value: "SHOW_LAYER",
		},
	}
}

func (c TriggerProcessReplyMode) Value() string {
	return c.value
}

func (c TriggerProcessReplyMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TriggerProcessReplyMode) UnmarshalJSON(b []byte) error {
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

type TriggerProcessReplyOrder struct {
	value string
}

type TriggerProcessReplyOrderEnum struct {
	RANDOM TriggerProcessReplyOrder
	ORDER  TriggerProcessReplyOrder
}

func GetTriggerProcessReplyOrderEnum() TriggerProcessReplyOrderEnum {
	return TriggerProcessReplyOrderEnum{
		RANDOM: TriggerProcessReplyOrder{
			value: "RANDOM",
		},
		ORDER: TriggerProcessReplyOrder{
			value: "ORDER",
		},
	}
}

func (c TriggerProcessReplyOrder) Value() string {
	return c.value
}

func (c TriggerProcessReplyOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TriggerProcessReplyOrder) UnmarshalJSON(b []byte) error {
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
