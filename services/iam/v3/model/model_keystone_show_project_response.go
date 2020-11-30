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
type KeystoneShowProjectResponse struct {
	Project        *ProjectResult `json:"project,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o KeystoneShowProjectResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneShowProjectResponse", string(data)}, " ")
}
