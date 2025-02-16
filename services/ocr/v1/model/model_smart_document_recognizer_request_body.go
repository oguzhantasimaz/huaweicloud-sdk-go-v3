package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SmartDocumentRecognizerRequestBody struct {

	// 与url二选一。图片或PDF格式，base64编码，要求base64编码后大小不超过10M。 图像尺寸不小于15×15像素，最长边不超过8192像素，支持JPG/PNG/BMP/TIFF格式。 PDF以144dpi的分辨率转为图像进行文档解析，需符合上述图像尺寸规定。若PDF有多页，当前仅对第1页进行识别。
	Data *string `json:"data,omitempty"`

	// 与data二选一。 图片或PDF的URL路径，目前仅支持华为云上OBS提供的匿名公开授权访问的URL以及公网URL。 > 说明： - 接口响应时间依赖于图片的下载时间，如果图片下载时间过长，会返回接口调用失败。 - 请保证被检测图片所在的存储服务稳定可靠，推荐使用OBS服务存储图片数据。 - url中不能存在中文字符，若存在，中文需要进行utf8编码。
	Url *string `json:"url,omitempty"`

	// 是否进行键值对（key-value）提取。若是，结果会以“kv_result”这一关键字返回。
	Kv *bool `json:"kv,omitempty"`

	// 是否进行表格识别。此处表格特指逻辑表格，通常具有M行N列的形式，且第一行或第一列为表头。若是，结果会以“table_result”这一关键字返回。
	Table *bool `json:"table,omitempty"`

	// 是否进行版面分析。若是，结果会以“layout_result”这一关键字返回。
	Layout *bool `json:"layout,omitempty"`

	// 仅当table为True时有效。是否返回表格转换Microsoft Excel的Base64编码字段。
	ReturnExcel *bool `json:"return_excel,omitempty"`
}

func (o SmartDocumentRecognizerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartDocumentRecognizerRequestBody struct{}"
	}

	return strings.Join([]string{"SmartDocumentRecognizerRequestBody", string(data)}, " ")
}
