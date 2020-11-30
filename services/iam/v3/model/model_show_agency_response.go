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
type ShowAgencyResponse struct {
	Agency         *AgencyResult `json:"agency,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowAgencyResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowAgencyResponse", string(data)}, " ")
}
