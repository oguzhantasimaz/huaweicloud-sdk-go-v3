package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOptmJobResponse Response Object
type ShowOptmJobResponse struct {
	BasicInfo *DrugJobDto `json:"basic_info,omitempty"`

	// 分子SMILES表达式
	Smiles *string `json:"smiles,omitempty"`

	MoleculeFile *DrugFile `json:"molecule_file,omitempty"`

	// 生成分子数量
	NumTrials *int32 `json:"num_trials,omitempty"`

	BindingSite *BindSiteDto `json:"binding_site,omitempty"`

	// 弱约束集合
	WeakConstraints *[]WeakConstraintDto `json:"weak_constraints,omitempty"`

	// 强约束集合
	StrongConstraints *[]StrongConstraintDto `json:"strong_constraints,omitempty"`

	// 初始化采样权重
	SamplerMixinWeight *float32 `json:"sampler_mixin_weight,omitempty"`

	// 模型列表
	Models         *[]BasicDrugModel `json:"models,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowOptmJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOptmJobResponse struct{}"
	}

	return strings.Join([]string{"ShowOptmJobResponse", string(data)}, " ")
}
