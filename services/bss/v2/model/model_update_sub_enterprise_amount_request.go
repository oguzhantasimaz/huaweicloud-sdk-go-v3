/*
 * Bss
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
type UpdateSubEnterpriseAmountRequest struct {
	Body *TransferEnterpriseMultiAccountReq `json:"body,omitempty"`
}

func (o UpdateSubEnterpriseAmountRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateSubEnterpriseAmountRequest", string(data)}, " ")
}
