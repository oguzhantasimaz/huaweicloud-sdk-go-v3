package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DigitalAssetSummary struct {

	// 资产ID。
	AssetId *string `json:"asset_id,omitempty"`

	// 资产名称。
	AssetName *string `json:"asset_name,omitempty"`

	// 资产类型。 公共资产类型： * VOICE_MODEL：音色模型（仅系统管理员可上传，普通租户仅可查询） * VIDEO：视频文件 * IMAGE：图片文件 * PPT：幻灯片文件 * MUSIC: 音乐 * AUDIO: 音频 * COMMON_FILE：通用文件  分身数字人资产： * HUMAN_MODEL_2D: 分身数字人模型 * BUSINESS_CARD_TEMPLET: 数字人名片模板  3D数字人资产： * HUMAN_MODEL：3D数字人模型 * SCENE：场景模型 * ANIMATION：动作动画 * MATERIAL：风格化素材 * NORMAL_MODEL: 普通模型
	AssetType *DigitalAssetSummaryAssetType `json:"asset_type,omitempty"`

	// 封面图片路径。
	CoverUrl *string `json:"cover_url,omitempty"`
}

func (o DigitalAssetSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DigitalAssetSummary struct{}"
	}

	return strings.Join([]string{"DigitalAssetSummary", string(data)}, " ")
}

type DigitalAssetSummaryAssetType struct {
	value string
}

type DigitalAssetSummaryAssetTypeEnum struct {
	HUMAN_MODEL           DigitalAssetSummaryAssetType
	MODEL                 DigitalAssetSummaryAssetType
	ANIMATION             DigitalAssetSummaryAssetType
	SCENE                 DigitalAssetSummaryAssetType
	PPT                   DigitalAssetSummaryAssetType
	VIDEO                 DigitalAssetSummaryAssetType
	IMAGE                 DigitalAssetSummaryAssetType
	MATERIAL              DigitalAssetSummaryAssetType
	VOICE_MODEL           DigitalAssetSummaryAssetType
	HUMAN_MODEL_2_D       DigitalAssetSummaryAssetType
	BUSINESS_CARD_TEMPLET DigitalAssetSummaryAssetType
	MUSIC                 DigitalAssetSummaryAssetType
	AUDIO                 DigitalAssetSummaryAssetType
}

func GetDigitalAssetSummaryAssetTypeEnum() DigitalAssetSummaryAssetTypeEnum {
	return DigitalAssetSummaryAssetTypeEnum{
		HUMAN_MODEL: DigitalAssetSummaryAssetType{
			value: "HUMAN_MODEL",
		},
		MODEL: DigitalAssetSummaryAssetType{
			value: "MODEL",
		},
		ANIMATION: DigitalAssetSummaryAssetType{
			value: "ANIMATION",
		},
		SCENE: DigitalAssetSummaryAssetType{
			value: "SCENE",
		},
		PPT: DigitalAssetSummaryAssetType{
			value: "PPT",
		},
		VIDEO: DigitalAssetSummaryAssetType{
			value: "VIDEO",
		},
		IMAGE: DigitalAssetSummaryAssetType{
			value: "IMAGE",
		},
		MATERIAL: DigitalAssetSummaryAssetType{
			value: "MATERIAL",
		},
		VOICE_MODEL: DigitalAssetSummaryAssetType{
			value: "VOICE_MODEL",
		},
		HUMAN_MODEL_2_D: DigitalAssetSummaryAssetType{
			value: "HUMAN_MODEL_2D",
		},
		BUSINESS_CARD_TEMPLET: DigitalAssetSummaryAssetType{
			value: "BUSINESS_CARD_TEMPLET",
		},
		MUSIC: DigitalAssetSummaryAssetType{
			value: "MUSIC",
		},
		AUDIO: DigitalAssetSummaryAssetType{
			value: "AUDIO",
		},
	}
}

func (c DigitalAssetSummaryAssetType) Value() string {
	return c.value
}

func (c DigitalAssetSummaryAssetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DigitalAssetSummaryAssetType) UnmarshalJSON(b []byte) error {
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
