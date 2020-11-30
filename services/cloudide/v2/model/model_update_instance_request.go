/*
 * CloudIDE
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateInstanceRequest struct {
	InstanceId string               `json:"instance_id"`
	Body       *InstanceUpdateParam `json:"body,omitempty"`
}

func (o UpdateInstanceRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateInstanceRequest", string(data)}, " ")
}
