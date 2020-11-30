/*
 * IAM
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
type ShowDomainQuotaResponse struct {
	Quotas         *QuotaResult `json:"quotas,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowDomainQuotaResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowDomainQuotaResponse", string(data)}, " ")
}
