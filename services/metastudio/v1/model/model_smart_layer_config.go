package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SmartLayerConfig 智能图层配置。
type SmartLayerConfig struct {

	// 图层类型。 - IMAGE： 素材图片图层 - VIDEO： 素材视频图层
	LayerType SmartLayerConfigLayerType `json:"layer_type"`

	Position *LayerPositionConfig `json:"position"`

	Size *LayerSizeConfig `json:"size,omitempty"`

	ImageConfig *SmartImageLayerConfig `json:"image_config,omitempty"`

	VideoConfig *SmartVideoLayerConfig `json:"video_config,omitempty"`
}

func (o SmartLayerConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartLayerConfig struct{}"
	}

	return strings.Join([]string{"SmartLayerConfig", string(data)}, " ")
}

type SmartLayerConfigLayerType struct {
	value string
}

type SmartLayerConfigLayerTypeEnum struct {
	IMAGE SmartLayerConfigLayerType
	VIDEO SmartLayerConfigLayerType
}

func GetSmartLayerConfigLayerTypeEnum() SmartLayerConfigLayerTypeEnum {
	return SmartLayerConfigLayerTypeEnum{
		IMAGE: SmartLayerConfigLayerType{
			value: "IMAGE",
		},
		VIDEO: SmartLayerConfigLayerType{
			value: "VIDEO",
		},
	}
}

func (c SmartLayerConfigLayerType) Value() string {
	return c.value
}

func (c SmartLayerConfigLayerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SmartLayerConfigLayerType) UnmarshalJSON(b []byte) error {
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
