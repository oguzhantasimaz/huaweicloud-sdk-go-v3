package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type MigrateNodeResponse struct {
	// API版本，固定值“v3”。

	ApiVersion *string `json:"apiVersion,omitempty"`
	// API类型，固定值“MigrateNodesTask”。

	Kind *string `json:"kind,omitempty"`

	Spec *MigrateNodesSpec `json:"spec,omitempty"`

	Status         *TaskStatus `json:"status,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o MigrateNodeResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MigrateNodeResponse struct{}"
	}

	return strings.Join([]string{"MigrateNodeResponse", string(data)}, " ")
}
