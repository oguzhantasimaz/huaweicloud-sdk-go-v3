/*
 * As
 *
 * 弹性伸缩API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type EnableOrDisableScalingGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o EnableOrDisableScalingGroupResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"EnableOrDisableScalingGroupResponse", string(data)}, " ")
}
