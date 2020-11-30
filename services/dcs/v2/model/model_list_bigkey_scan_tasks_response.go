/*
 * DCS
 *
 * DCS V2版本API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListBigkeyScanTasksResponse struct {
	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`
	// 总数
	Count *int32 `json:"count,omitempty"`
	// 大key分析记录列表
	Records        *[]RecordsResponse `json:"records,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListBigkeyScanTasksResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListBigkeyScanTasksResponse", string(data)}, " ")
}
