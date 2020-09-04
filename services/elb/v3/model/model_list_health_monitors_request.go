/*
 * ELB
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
type ListHealthMonitorsRequest struct {
	Marker              *string  `json:"marker,omitempty"`
	Limit               *int32   `json:"limit,omitempty"`
	PageReverse         *bool    `json:"page_reverse,omitempty"`
	Id                  []string `json:"id,omitempty"`
	MonitorPort         []int32  `json:"monitor_port,omitempty"`
	DomainName          []string `json:"domain_name,omitempty"`
	Name                []string `json:"name,omitempty"`
	Delay               []int32  `json:"delay,omitempty"`
	MaxRetries          []int32  `json:"max_retries,omitempty"`
	AdminStateUp        *bool    `json:"admin_state_up,omitempty"`
	MaxRetriesDown      []int32  `json:"max_retries_down,omitempty"`
	Timeout             *int32   `json:"timeout,omitempty"`
	Type                []string `json:"type,omitempty"`
	ExpectedCodes       []string `json:"expected_codes,omitempty"`
	UrlPath             []string `json:"url_path,omitempty"`
	HttpMethod          []string `json:"http_method,omitempty"`
	EnterpriseProjectId []string `json:"enterprise_project_id,omitempty"`
}

func (o ListHealthMonitorsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListHealthMonitorsRequest", string(data)}, " ")
}
