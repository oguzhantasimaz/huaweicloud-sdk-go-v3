package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/aos/v1/model"
)

type ApplyExecutionPlanInvoker struct {
	*invoker.BaseInvoker
}

func (i *ApplyExecutionPlanInvoker) Invoke() (*model.ApplyExecutionPlanResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ApplyExecutionPlanResponse), nil
	}
}

type CreateExecutionPlanInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateExecutionPlanInvoker) Invoke() (*model.CreateExecutionPlanResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateExecutionPlanResponse), nil
	}
}

type DeleteExecutionPlanInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteExecutionPlanInvoker) Invoke() (*model.DeleteExecutionPlanResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteExecutionPlanResponse), nil
	}
}

type EstimateExecutionPlanPriceInvoker struct {
	*invoker.BaseInvoker
}

func (i *EstimateExecutionPlanPriceInvoker) Invoke() (*model.EstimateExecutionPlanPriceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EstimateExecutionPlanPriceResponse), nil
	}
}

type GetExecutionPlanInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetExecutionPlanInvoker) Invoke() (*model.GetExecutionPlanResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetExecutionPlanResponse), nil
	}
}

type GetExecutionPlanMetadataInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetExecutionPlanMetadataInvoker) Invoke() (*model.GetExecutionPlanMetadataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetExecutionPlanMetadataResponse), nil
	}
}

type ListExecutionPlansInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListExecutionPlansInvoker) Invoke() (*model.ListExecutionPlansResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListExecutionPlansResponse), nil
	}
}

type ContinueDeployStackInvoker struct {
	*invoker.BaseInvoker
}

func (i *ContinueDeployStackInvoker) Invoke() (*model.ContinueDeployStackResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ContinueDeployStackResponse), nil
	}
}

type ContinueRollbackStackInvoker struct {
	*invoker.BaseInvoker
}

func (i *ContinueRollbackStackInvoker) Invoke() (*model.ContinueRollbackStackResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ContinueRollbackStackResponse), nil
	}
}

type CreateStackInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateStackInvoker) Invoke() (*model.CreateStackResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateStackResponse), nil
	}
}

type DeleteStackInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteStackInvoker) Invoke() (*model.DeleteStackResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteStackResponse), nil
	}
}

type DeleteStackEnhancedInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteStackEnhancedInvoker) Invoke() (*model.DeleteStackEnhancedResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteStackEnhancedResponse), nil
	}
}

type DeployStackInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeployStackInvoker) Invoke() (*model.DeployStackResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeployStackResponse), nil
	}
}

type GetStackMetadataInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetStackMetadataInvoker) Invoke() (*model.GetStackMetadataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetStackMetadataResponse), nil
	}
}

type GetStackTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetStackTemplateInvoker) Invoke() (*model.GetStackTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetStackTemplateResponse), nil
	}
}

type ListStackEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStackEventsInvoker) Invoke() (*model.ListStackEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStackEventsResponse), nil
	}
}

type ListStackOutputsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStackOutputsInvoker) Invoke() (*model.ListStackOutputsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStackOutputsResponse), nil
	}
}

type ListStackResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStackResourcesInvoker) Invoke() (*model.ListStackResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStackResourcesResponse), nil
	}
}

type ListStacksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStacksInvoker) Invoke() (*model.ListStacksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStacksResponse), nil
	}
}

type UpdateStackInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateStackInvoker) Invoke() (*model.UpdateStackResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateStackResponse), nil
	}
}

type CreateStackInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateStackInstanceInvoker) Invoke() (*model.CreateStackInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateStackInstanceResponse), nil
	}
}

type CreateStackSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateStackSetInvoker) Invoke() (*model.CreateStackSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateStackSetResponse), nil
	}
}

type DeleteStackInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteStackInstanceInvoker) Invoke() (*model.DeleteStackInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteStackInstanceResponse), nil
	}
}

type DeleteStackInstanceDeprecatedInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteStackInstanceDeprecatedInvoker) Invoke() (*model.DeleteStackInstanceDeprecatedResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteStackInstanceDeprecatedResponse), nil
	}
}

type DeleteStackSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteStackSetInvoker) Invoke() (*model.DeleteStackSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteStackSetResponse), nil
	}
}

type DeployStackSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeployStackSetInvoker) Invoke() (*model.DeployStackSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeployStackSetResponse), nil
	}
}

type ListStackInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStackInstancesInvoker) Invoke() (*model.ListStackInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStackInstancesResponse), nil
	}
}

type ListStackSetOperationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStackSetOperationsInvoker) Invoke() (*model.ListStackSetOperationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStackSetOperationsResponse), nil
	}
}

type ListStackSetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStackSetsInvoker) Invoke() (*model.ListStackSetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStackSetsResponse), nil
	}
}

type ShowStackInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStackInstanceInvoker) Invoke() (*model.ShowStackInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStackInstanceResponse), nil
	}
}

type ShowStackSetMetadataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStackSetMetadataInvoker) Invoke() (*model.ShowStackSetMetadataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStackSetMetadataResponse), nil
	}
}

type ShowStackSetOperationMetadataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStackSetOperationMetadataInvoker) Invoke() (*model.ShowStackSetOperationMetadataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStackSetOperationMetadataResponse), nil
	}
}

type ShowStackSetTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStackSetTemplateInvoker) Invoke() (*model.ShowStackSetTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStackSetTemplateResponse), nil
	}
}

type UpdateStackInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateStackInstancesInvoker) Invoke() (*model.UpdateStackInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateStackInstancesResponse), nil
	}
}

type UpdateStackSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateStackSetInvoker) Invoke() (*model.UpdateStackSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateStackSetResponse), nil
	}
}

type ParseTemplateVariablesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ParseTemplateVariablesInvoker) Invoke() (*model.ParseTemplateVariablesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ParseTemplateVariablesResponse), nil
	}
}

type DeleteTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTemplateInvoker) Invoke() (*model.DeleteTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTemplateResponse), nil
	}
}

type DeleteTemplateVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTemplateVersionInvoker) Invoke() (*model.DeleteTemplateVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTemplateVersionResponse), nil
	}
}

type ListTemplateVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTemplateVersionsInvoker) Invoke() (*model.ListTemplateVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTemplateVersionsResponse), nil
	}
}

type ListTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTemplatesInvoker) Invoke() (*model.ListTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTemplatesResponse), nil
	}
}

type ShowTemplateMetadataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTemplateMetadataInvoker) Invoke() (*model.ShowTemplateMetadataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTemplateMetadataResponse), nil
	}
}

type ShowTemplateVersionContentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTemplateVersionContentInvoker) Invoke() (*model.ShowTemplateVersionContentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTemplateVersionContentResponse), nil
	}
}

type ShowTemplateVersionMetadataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTemplateVersionMetadataInvoker) Invoke() (*model.ShowTemplateVersionMetadataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTemplateVersionMetadataResponse), nil
	}
}

type UpdateTemplateMetadataInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTemplateMetadataInvoker) Invoke() (*model.UpdateTemplateMetadataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTemplateMetadataResponse), nil
	}
}
