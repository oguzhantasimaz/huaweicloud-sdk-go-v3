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
type CreateDataImageResponse struct {
	// 异步任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDataImageResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateDataImageResponse", string(data)}, " ")
}
