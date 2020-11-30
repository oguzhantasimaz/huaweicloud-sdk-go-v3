/*
 * BCS
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
type ListEntityMetricRequest struct {
	BlockchainId string                       `json:"blockchain_id"`
	Body         *ListEntityMetricRequestBody `json:"body,omitempty"`
}

func (o ListEntityMetricRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListEntityMetricRequest", string(data)}, " ")
}
