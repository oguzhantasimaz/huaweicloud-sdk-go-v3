package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateVideoSynthesisTaskResponse struct {

	// 任务唯一标识
	TaskId *string `json:"task_id,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateVideoSynthesisTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVideoSynthesisTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateVideoSynthesisTaskResponse", string(data)}, " ")
}
