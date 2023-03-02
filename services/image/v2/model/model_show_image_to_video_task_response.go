package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowImageToVideoTaskResponse struct {

	// 任务id
	TaskId *string `json:"task_id,omitempty"`

	// 任务创建时间，格式为ISO8601：YYYY-MM-DDThh:mm:ssZ
	CreateTime *string `json:"create_time,omitempty"`

	// 任务更新时间，格式为ISO8601：YYYY-MM-DDThh:mm:ssZ
	UpdateTime *string `json:"update_time,omitempty"`

	// 任务当前的状态，分别为SUCCEEDED（运行成功），FAILED（运行失败），RUNNING（运行中）。
	State *string `json:"state,omitempty"`

	Input *TaskInput `json:"input,omitempty"`

	Output *TaskOutput `json:"output,omitempty"`

	Config *ImageToVideoConfig `json:"config,omitempty"`

	Callback *TaskCallback `json:"callback,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowImageToVideoTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageToVideoTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowImageToVideoTaskResponse", string(data)}, " ")
}
