package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DrugJobDto struct {

	// 作业id
	Id *string `json:"id,omitempty"`

	// 作业的名称，取值范围：[1,63]，允许大小写字母、数字、以及特殊字符中划线(-)
	Name *string `json:"name,omitempty"`

	// 作业标签
	Labels *[]string `json:"labels,omitempty"`

	// 作业状态
	Status *string `json:"status,omitempty"`

	// 作业类型
	Type *string `json:"type,omitempty"`

	// 作业创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 作业结束时间
	FinishTime *string `json:"finish_time,omitempty"`

	// 作业结束时间
	StartTime *string `json:"start_time,omitempty"`

	// 失败提示，当作业执行失败时会返回
	FailedMessage *string `json:"failed_message,omitempty"`

	// 创建任务的用户名称
	UserName *string `json:"user_name,omitempty"`

	// 作业结果输出目录
	OutputDir *string `json:"output_dir,omitempty"`

	// 预估功能调用消耗次数
	ExpectChargeNum *int32 `json:"expect_charge_num,omitempty"`

	// 实际功能调用消耗次数
	RealChargeNum *int32 `json:"real_charge_num,omitempty"`

	Progress *Progress `json:"progress,omitempty"`
}

func (o DrugJobDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DrugJobDto struct{}"
	}

	return strings.Join([]string{"DrugJobDto", string(data)}, " ")
}
