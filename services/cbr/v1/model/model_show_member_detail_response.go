/*
 * Cbr
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

// Response Object
type ShowMemberDetailResponse struct {
	// 共享状态
	Status ShowMemberDetailResponseStatus `json:"status,omitempty"`
	// 共享时间，例如:\"2020-02-05T10:38:34.209782\"
	CreatedAt *string `json:"created_at,omitempty"`
	// 更新时间，例如:\"2020-02-05T10:38:34.209782\"
	UpdatedAt *string `json:"updated_at,omitempty"`
	// 备份副本id
	BackupId *string `json:"backup_id,omitempty"`
	// 接受的共享备份副本注册的镜像id
	ImageId *string `json:"image_id,omitempty"`
	// 接受备份共享的项目id
	DestProjectId *string `json:"dest_project_id,omitempty"`
	// 目标端接受共享备份的存储库id
	VaultId *string `json:"vault_id,omitempty"`
	// 共享记录id
	Id *string `json:"id,omitempty"`
}

func (o ShowMemberDetailResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowMemberDetailResponse", string(data)}, " ")
}

type ShowMemberDetailResponseStatus struct {
	value string
}

type ShowMemberDetailResponseStatusEnum struct {
	PENDING  ShowMemberDetailResponseStatus
	ACCEPTED ShowMemberDetailResponseStatus
	REJECTED ShowMemberDetailResponseStatus
}

func GetShowMemberDetailResponseStatusEnum() ShowMemberDetailResponseStatusEnum {
	return ShowMemberDetailResponseStatusEnum{
		PENDING: ShowMemberDetailResponseStatus{
			value: "pending",
		},
		ACCEPTED: ShowMemberDetailResponseStatus{
			value: "accepted",
		},
		REJECTED: ShowMemberDetailResponseStatus{
			value: "rejected",
		},
	}
}

func (c ShowMemberDetailResponseStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ShowMemberDetailResponseStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
