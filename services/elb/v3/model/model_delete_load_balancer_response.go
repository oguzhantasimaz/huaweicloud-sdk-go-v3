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

// Response Object
type DeleteLoadBalancerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteLoadBalancerResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteLoadBalancerResponse", string(data)}, " ")
}
