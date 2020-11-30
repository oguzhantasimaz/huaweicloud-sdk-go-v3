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

// 监控数据信息
type MetricItemResultApi struct {
	Metric *MetricDemision `json:"metric,omitempty"`
	// 监控数据信息
	DataPoints *[]MetricDataPoints `json:"dataPoints,omitempty"`
}

func (o MetricItemResultApi) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"MetricItemResultApi", string(data)}, " ")
}
