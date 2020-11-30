/*
 * IMS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateWholeImageResponse struct {
	// 异步任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateWholeImageResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateWholeImageResponse", string(data)}, " ")
}
