package v4

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/projectman/v4/model"
)

type ProjectManClient struct {
	HcClient *http_client.HcHttpClient
}

func NewProjectManClient(hcClient *http_client.HcHttpClient) *ProjectManClient {
	return &ProjectManClient{HcClient: hcClient}
}

func ProjectManClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// AddApplyJoinProjectForAgc AGC调用 当前用户申请加入项目
//
// AGC调用 当前用户申请加入项目, 申请的用户id写在header中
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) AddApplyJoinProjectForAgc(request *model.AddApplyJoinProjectForAgcRequest) (*model.AddApplyJoinProjectForAgcResponse, error) {
	requestDef := GenReqDefForAddApplyJoinProjectForAgc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddApplyJoinProjectForAgcResponse), nil
	}
}

// AddApplyJoinProjectForAgcInvoker AGC调用 当前用户申请加入项目
func (c *ProjectManClient) AddApplyJoinProjectForAgcInvoker(request *model.AddApplyJoinProjectForAgcRequest) *AddApplyJoinProjectForAgcInvoker {
	requestDef := GenReqDefForAddApplyJoinProjectForAgc()
	return &AddApplyJoinProjectForAgcInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddMemberV4 添加项目成员
//
// 添加项目成员,可以添加跨租户成员
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) AddMemberV4(request *model.AddMemberV4Request) (*model.AddMemberV4Response, error) {
	requestDef := GenReqDefForAddMemberV4()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddMemberV4Response), nil
	}
}

// AddMemberV4Invoker 添加项目成员
func (c *ProjectManClient) AddMemberV4Invoker(request *model.AddMemberV4Request) *AddMemberV4Invoker {
	requestDef := GenReqDefForAddMemberV4()
	return &AddMemberV4Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchAddMembersV4 批量添加项目成员
//
// 批量添加项目成员，只能添加和项目创建者同一租户下的成员，不正确的用户id会略过，添加的用户超过权限的，默认角色设置为7
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) BatchAddMembersV4(request *model.BatchAddMembersV4Request) (*model.BatchAddMembersV4Response, error) {
	requestDef := GenReqDefForBatchAddMembersV4()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddMembersV4Response), nil
	}
}

// BatchAddMembersV4Invoker 批量添加项目成员
func (c *ProjectManClient) BatchAddMembersV4Invoker(request *model.BatchAddMembersV4Request) *BatchAddMembersV4Invoker {
	requestDef := GenReqDefForBatchAddMembersV4()
	return &BatchAddMembersV4Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteMembersV4 批量删除项目成员
//
// 批量删除项目成员
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) BatchDeleteMembersV4(request *model.BatchDeleteMembersV4Request) (*model.BatchDeleteMembersV4Response, error) {
	requestDef := GenReqDefForBatchDeleteMembersV4()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteMembersV4Response), nil
	}
}

// BatchDeleteMembersV4Invoker 批量删除项目成员
func (c *ProjectManClient) BatchDeleteMembersV4Invoker(request *model.BatchDeleteMembersV4Request) *BatchDeleteMembersV4Invoker {
	requestDef := GenReqDefForBatchDeleteMembersV4()
	return &BatchDeleteMembersV4Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateChildNickNames 更新子用户昵称
//
// 拥有te_admin角色的用户可以更新其他用户的昵称
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) BatchUpdateChildNickNames(request *model.BatchUpdateChildNickNamesRequest) (*model.BatchUpdateChildNickNamesResponse, error) {
	requestDef := GenReqDefForBatchUpdateChildNickNames()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateChildNickNamesResponse), nil
	}
}

// BatchUpdateChildNickNamesInvoker 更新子用户昵称
func (c *ProjectManClient) BatchUpdateChildNickNamesInvoker(request *model.BatchUpdateChildNickNamesRequest) *BatchUpdateChildNickNamesInvoker {
	requestDef := GenReqDefForBatchUpdateChildNickNames()
	return &BatchUpdateChildNickNamesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckProjectNameV4 检查项目名称是否存在
//
// 检查项目名称是否存在
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) CheckProjectNameV4(request *model.CheckProjectNameV4Request) (*model.CheckProjectNameV4Response, error) {
	requestDef := GenReqDefForCheckProjectNameV4()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckProjectNameV4Response), nil
	}
}

// CheckProjectNameV4Invoker 检查项目名称是否存在
func (c *ProjectManClient) CheckProjectNameV4Invoker(request *model.CheckProjectNameV4Request) *CheckProjectNameV4Invoker {
	requestDef := GenReqDefForCheckProjectNameV4()
	return &CheckProjectNameV4Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateProjectV4 创建项目
//
// 创建项目
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) CreateProjectV4(request *model.CreateProjectV4Request) (*model.CreateProjectV4Response, error) {
	requestDef := GenReqDefForCreateProjectV4()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateProjectV4Response), nil
	}
}

// CreateProjectV4Invoker 创建项目
func (c *ProjectManClient) CreateProjectV4Invoker(request *model.CreateProjectV4Request) *CreateProjectV4Invoker {
	requestDef := GenReqDefForCreateProjectV4()
	return &CreateProjectV4Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteProjectV4 删除项目
//
// 删除项目
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) DeleteProjectV4(request *model.DeleteProjectV4Request) (*model.DeleteProjectV4Response, error) {
	requestDef := GenReqDefForDeleteProjectV4()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteProjectV4Response), nil
	}
}

// DeleteProjectV4Invoker 删除项目
func (c *ProjectManClient) DeleteProjectV4Invoker(request *model.DeleteProjectV4Request) *DeleteProjectV4Invoker {
	requestDef := GenReqDefForDeleteProjectV4()
	return &DeleteProjectV4Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDomainNotAddedProjectsV4 获取租户没有加入的项目
//
// 获取租户没有加入的项目
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) ListDomainNotAddedProjectsV4(request *model.ListDomainNotAddedProjectsV4Request) (*model.ListDomainNotAddedProjectsV4Response, error) {
	requestDef := GenReqDefForListDomainNotAddedProjectsV4()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDomainNotAddedProjectsV4Response), nil
	}
}

// ListDomainNotAddedProjectsV4Invoker 获取租户没有加入的项目
func (c *ProjectManClient) ListDomainNotAddedProjectsV4Invoker(request *model.ListDomainNotAddedProjectsV4Request) *ListDomainNotAddedProjectsV4Invoker {
	requestDef := GenReqDefForListDomainNotAddedProjectsV4()
	return &ListDomainNotAddedProjectsV4Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectBugStaticsV4 获取bug统计信息
//
// 获取bug统计信息，按模块统计
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) ListProjectBugStaticsV4(request *model.ListProjectBugStaticsV4Request) (*model.ListProjectBugStaticsV4Response, error) {
	requestDef := GenReqDefForListProjectBugStaticsV4()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectBugStaticsV4Response), nil
	}
}

// ListProjectBugStaticsV4Invoker 获取bug统计信息
func (c *ProjectManClient) ListProjectBugStaticsV4Invoker(request *model.ListProjectBugStaticsV4Request) *ListProjectBugStaticsV4Invoker {
	requestDef := GenReqDefForListProjectBugStaticsV4()
	return &ListProjectBugStaticsV4Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectDemandStaticV4 获取需求统计信息
//
// 获取需求统计信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) ListProjectDemandStaticV4(request *model.ListProjectDemandStaticV4Request) (*model.ListProjectDemandStaticV4Response, error) {
	requestDef := GenReqDefForListProjectDemandStaticV4()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectDemandStaticV4Response), nil
	}
}

// ListProjectDemandStaticV4Invoker 获取需求统计信息
func (c *ProjectManClient) ListProjectDemandStaticV4Invoker(request *model.ListProjectDemandStaticV4Request) *ListProjectDemandStaticV4Invoker {
	requestDef := GenReqDefForListProjectDemandStaticV4()
	return &ListProjectDemandStaticV4Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectMembersV4 获取指定项目的成员用户列表
//
// 获取项目成员列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) ListProjectMembersV4(request *model.ListProjectMembersV4Request) (*model.ListProjectMembersV4Response, error) {
	requestDef := GenReqDefForListProjectMembersV4()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectMembersV4Response), nil
	}
}

// ListProjectMembersV4Invoker 获取指定项目的成员用户列表
func (c *ProjectManClient) ListProjectMembersV4Invoker(request *model.ListProjectMembersV4Request) *ListProjectMembersV4Invoker {
	requestDef := GenReqDefForListProjectMembersV4()
	return &ListProjectMembersV4Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectsV4 查询项目列表
//
// 查询项目列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) ListProjectsV4(request *model.ListProjectsV4Request) (*model.ListProjectsV4Response, error) {
	requestDef := GenReqDefForListProjectsV4()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectsV4Response), nil
	}
}

// ListProjectsV4Invoker 查询项目列表
func (c *ProjectManClient) ListProjectsV4Invoker(request *model.ListProjectsV4Request) *ListProjectsV4Invoker {
	requestDef := GenReqDefForListProjectsV4()
	return &ListProjectsV4Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RemoveProject 主动退出项目
//
// 项目成员主动退出项目，项目创建者不能退出
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) RemoveProject(request *model.RemoveProjectRequest) (*model.RemoveProjectResponse, error) {
	requestDef := GenReqDefForRemoveProject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveProjectResponse), nil
	}
}

// RemoveProjectInvoker 主动退出项目
func (c *ProjectManClient) RemoveProjectInvoker(request *model.RemoveProjectRequest) *RemoveProjectInvoker {
	requestDef := GenReqDefForRemoveProject()
	return &RemoveProjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBugDensityV2 查询缺陷密度
//
// 查询缺陷密度
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) ShowBugDensityV2(request *model.ShowBugDensityV2Request) (*model.ShowBugDensityV2Response, error) {
	requestDef := GenReqDefForShowBugDensityV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBugDensityV2Response), nil
	}
}

// ShowBugDensityV2Invoker 查询缺陷密度
func (c *ProjectManClient) ShowBugDensityV2Invoker(request *model.ShowBugDensityV2Request) *ShowBugDensityV2Invoker {
	requestDef := GenReqDefForShowBugDensityV2()
	return &ShowBugDensityV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBugsPerDeveloper 查询人均bug
//
// 查询人均bug
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) ShowBugsPerDeveloper(request *model.ShowBugsPerDeveloperRequest) (*model.ShowBugsPerDeveloperResponse, error) {
	requestDef := GenReqDefForShowBugsPerDeveloper()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBugsPerDeveloperResponse), nil
	}
}

// ShowBugsPerDeveloperInvoker 查询人均bug
func (c *ProjectManClient) ShowBugsPerDeveloperInvoker(request *model.ShowBugsPerDeveloperRequest) *ShowBugsPerDeveloperInvoker {
	requestDef := GenReqDefForShowBugsPerDeveloper()
	return &ShowBugsPerDeveloperInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCompletionRate 查询需求按时完成率
//
// 查询需求按时完成率
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) ShowCompletionRate(request *model.ShowCompletionRateRequest) (*model.ShowCompletionRateResponse, error) {
	requestDef := GenReqDefForShowCompletionRate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCompletionRateResponse), nil
	}
}

// ShowCompletionRateInvoker 查询需求按时完成率
func (c *ProjectManClient) ShowCompletionRateInvoker(request *model.ShowCompletionRateRequest) *ShowCompletionRateInvoker {
	requestDef := GenReqDefForShowCompletionRate()
	return &ShowCompletionRateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCurUserInfo 获取当前用户信息
//
// 获取当前用户信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) ShowCurUserInfo(request *model.ShowCurUserInfoRequest) (*model.ShowCurUserInfoResponse, error) {
	requestDef := GenReqDefForShowCurUserInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCurUserInfoResponse), nil
	}
}

// ShowCurUserInfoInvoker 获取当前用户信息
func (c *ProjectManClient) ShowCurUserInfoInvoker(request *model.ShowCurUserInfoRequest) *ShowCurUserInfoInvoker {
	requestDef := GenReqDefForShowCurUserInfo()
	return &ShowCurUserInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCurUserRole 获取当前用户角色
//
// 获取用户在项目中的角色
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) ShowCurUserRole(request *model.ShowCurUserRoleRequest) (*model.ShowCurUserRoleResponse, error) {
	requestDef := GenReqDefForShowCurUserRole()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCurUserRoleResponse), nil
	}
}

// ShowCurUserRoleInvoker 获取当前用户角色
func (c *ProjectManClient) ShowCurUserRoleInvoker(request *model.ShowCurUserRoleRequest) *ShowCurUserRoleInvoker {
	requestDef := GenReqDefForShowCurUserRole()
	return &ShowCurUserRoleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectInfoV4 获取项目详情
//
// 获取项目详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) ShowProjectInfoV4(request *model.ShowProjectInfoV4Request) (*model.ShowProjectInfoV4Response, error) {
	requestDef := GenReqDefForShowProjectInfoV4()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectInfoV4Response), nil
	}
}

// ShowProjectInfoV4Invoker 获取项目详情
func (c *ProjectManClient) ShowProjectInfoV4Invoker(request *model.ShowProjectInfoV4Request) *ShowProjectInfoV4Invoker {
	requestDef := GenReqDefForShowProjectInfoV4()
	return &ShowProjectInfoV4Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectSummaryV4 获取项目概览
//
// 获取项目概览
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) ShowProjectSummaryV4(request *model.ShowProjectSummaryV4Request) (*model.ShowProjectSummaryV4Response, error) {
	requestDef := GenReqDefForShowProjectSummaryV4()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectSummaryV4Response), nil
	}
}

// ShowProjectSummaryV4Invoker 获取项目概览
func (c *ProjectManClient) ShowProjectSummaryV4Invoker(request *model.ShowProjectSummaryV4Request) *ShowProjectSummaryV4Invoker {
	requestDef := GenReqDefForShowProjectSummaryV4()
	return &ShowProjectSummaryV4Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMembesRoleV4 更新成员在项目中的角色
//
// 更新成员在项目中的角色
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) UpdateMembesRoleV4(request *model.UpdateMembesRoleV4Request) (*model.UpdateMembesRoleV4Response, error) {
	requestDef := GenReqDefForUpdateMembesRoleV4()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMembesRoleV4Response), nil
	}
}

// UpdateMembesRoleV4Invoker 更新成员在项目中的角色
func (c *ProjectManClient) UpdateMembesRoleV4Invoker(request *model.UpdateMembesRoleV4Request) *UpdateMembesRoleV4Invoker {
	requestDef := GenReqDefForUpdateMembesRoleV4()
	return &UpdateMembesRoleV4Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateNickNameV4 更新用户昵称
//
// 更新用户昵称
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) UpdateNickNameV4(request *model.UpdateNickNameV4Request) (*model.UpdateNickNameV4Response, error) {
	requestDef := GenReqDefForUpdateNickNameV4()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNickNameV4Response), nil
	}
}

// UpdateNickNameV4Invoker 更新用户昵称
func (c *ProjectManClient) UpdateNickNameV4Invoker(request *model.UpdateNickNameV4Request) *UpdateNickNameV4Invoker {
	requestDef := GenReqDefForUpdateNickNameV4()
	return &UpdateNickNameV4Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProjectV4 更新项目
//
// 更新项目
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) UpdateProjectV4(request *model.UpdateProjectV4Request) (*model.UpdateProjectV4Response, error) {
	requestDef := GenReqDefForUpdateProjectV4()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProjectV4Response), nil
	}
}

// UpdateProjectV4Invoker 更新项目
func (c *ProjectManClient) UpdateProjectV4Invoker(request *model.UpdateProjectV4Request) *UpdateProjectV4Invoker {
	requestDef := GenReqDefForUpdateProjectV4()
	return &UpdateProjectV4Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteIssuesV4 批量删除工作项
//
// 批量删除工作项
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) BatchDeleteIssuesV4(request *model.BatchDeleteIssuesV4Request) (*model.BatchDeleteIssuesV4Response, error) {
	requestDef := GenReqDefForBatchDeleteIssuesV4()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteIssuesV4Response), nil
	}
}

// BatchDeleteIssuesV4Invoker 批量删除工作项
func (c *ProjectManClient) BatchDeleteIssuesV4Invoker(request *model.BatchDeleteIssuesV4Request) *BatchDeleteIssuesV4Invoker {
	requestDef := GenReqDefForBatchDeleteIssuesV4()
	return &BatchDeleteIssuesV4Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteIterationsV4 批量删除项目的迭代
//
// 批量删除项目的迭代
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) BatchDeleteIterationsV4(request *model.BatchDeleteIterationsV4Request) (*model.BatchDeleteIterationsV4Response, error) {
	requestDef := GenReqDefForBatchDeleteIterationsV4()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteIterationsV4Response), nil
	}
}

// BatchDeleteIterationsV4Invoker 批量删除项目的迭代
func (c *ProjectManClient) BatchDeleteIterationsV4Invoker(request *model.BatchDeleteIterationsV4Request) *BatchDeleteIterationsV4Invoker {
	requestDef := GenReqDefForBatchDeleteIterationsV4()
	return &BatchDeleteIterationsV4Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCustomfields 创建工作项类型自定义字段
//
// 创建工作项类型自定义字段
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) CreateCustomfields(request *model.CreateCustomfieldsRequest) (*model.CreateCustomfieldsResponse, error) {
	requestDef := GenReqDefForCreateCustomfields()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCustomfieldsResponse), nil
	}
}

// CreateCustomfieldsInvoker 创建工作项类型自定义字段
func (c *ProjectManClient) CreateCustomfieldsInvoker(request *model.CreateCustomfieldsRequest) *CreateCustomfieldsInvoker {
	requestDef := GenReqDefForCreateCustomfields()
	return &CreateCustomfieldsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateIssueV4 创建工作项
//
// 创建工作项
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) CreateIssueV4(request *model.CreateIssueV4Request) (*model.CreateIssueV4Response, error) {
	requestDef := GenReqDefForCreateIssueV4()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateIssueV4Response), nil
	}
}

// CreateIssueV4Invoker 创建工作项
func (c *ProjectManClient) CreateIssueV4Invoker(request *model.CreateIssueV4Request) *CreateIssueV4Invoker {
	requestDef := GenReqDefForCreateIssueV4()
	return &CreateIssueV4Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateIterationV4 创建Scrum项目迭代
//
// 创建Scrum项目迭代
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) CreateIterationV4(request *model.CreateIterationV4Request) (*model.CreateIterationV4Response, error) {
	requestDef := GenReqDefForCreateIterationV4()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateIterationV4Response), nil
	}
}

// CreateIterationV4Invoker 创建Scrum项目迭代
func (c *ProjectManClient) CreateIterationV4Invoker(request *model.CreateIterationV4Request) *CreateIterationV4Invoker {
	requestDef := GenReqDefForCreateIterationV4()
	return &CreateIterationV4Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSystemIssueV4 细粒度权限用户创建工作项
//
// 拥有IAM细粒度权限（projectmanConfig:systemSettingField:set）且在devcloud项目中有创建工作项的权限的用户可以设置工作项的创建者
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) CreateSystemIssueV4(request *model.CreateSystemIssueV4Request) (*model.CreateSystemIssueV4Response, error) {
	requestDef := GenReqDefForCreateSystemIssueV4()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSystemIssueV4Response), nil
	}
}

// CreateSystemIssueV4Invoker 细粒度权限用户创建工作项
func (c *ProjectManClient) CreateSystemIssueV4Invoker(request *model.CreateSystemIssueV4Request) *CreateSystemIssueV4Invoker {
	requestDef := GenReqDefForCreateSystemIssueV4()
	return &CreateSystemIssueV4Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteIssueV4 删除工作项
//
// 删除工作项
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) DeleteIssueV4(request *model.DeleteIssueV4Request) (*model.DeleteIssueV4Response, error) {
	requestDef := GenReqDefForDeleteIssueV4()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteIssueV4Response), nil
	}
}

// DeleteIssueV4Invoker 删除工作项
func (c *ProjectManClient) DeleteIssueV4Invoker(request *model.DeleteIssueV4Request) *DeleteIssueV4Invoker {
	requestDef := GenReqDefForDeleteIssueV4()
	return &DeleteIssueV4Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteIterationV4 删除项目迭代
//
// 删除项目迭代
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) DeleteIterationV4(request *model.DeleteIterationV4Request) (*model.DeleteIterationV4Response, error) {
	requestDef := GenReqDefForDeleteIterationV4()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteIterationV4Response), nil
	}
}

// DeleteIterationV4Invoker 删除项目迭代
func (c *ProjectManClient) DeleteIterationV4Invoker(request *model.DeleteIterationV4Request) *DeleteIterationV4Invoker {
	requestDef := GenReqDefForDeleteIterationV4()
	return &DeleteIterationV4Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListChildIssuesV4 获取子工作项
//
// 获取子工作项
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) ListChildIssuesV4(request *model.ListChildIssuesV4Request) (*model.ListChildIssuesV4Response, error) {
	requestDef := GenReqDefForListChildIssuesV4()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListChildIssuesV4Response), nil
	}
}

// ListChildIssuesV4Invoker 获取子工作项
func (c *ProjectManClient) ListChildIssuesV4Invoker(request *model.ListChildIssuesV4Request) *ListChildIssuesV4Invoker {
	requestDef := GenReqDefForListChildIssuesV4()
	return &ListChildIssuesV4Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIssueCommentsV4 获取指定工作项的评论列表
//
// 获取工作项的评论
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) ListIssueCommentsV4(request *model.ListIssueCommentsV4Request) (*model.ListIssueCommentsV4Response, error) {
	requestDef := GenReqDefForListIssueCommentsV4()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIssueCommentsV4Response), nil
	}
}

// ListIssueCommentsV4Invoker 获取指定工作项的评论列表
func (c *ProjectManClient) ListIssueCommentsV4Invoker(request *model.ListIssueCommentsV4Request) *ListIssueCommentsV4Invoker {
	requestDef := GenReqDefForListIssueCommentsV4()
	return &ListIssueCommentsV4Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIssueRecordsV4 获取工作项历史记录
//
// 获取工作项历史记录
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) ListIssueRecordsV4(request *model.ListIssueRecordsV4Request) (*model.ListIssueRecordsV4Response, error) {
	requestDef := GenReqDefForListIssueRecordsV4()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIssueRecordsV4Response), nil
	}
}

// ListIssueRecordsV4Invoker 获取工作项历史记录
func (c *ProjectManClient) ListIssueRecordsV4Invoker(request *model.ListIssueRecordsV4Request) *ListIssueRecordsV4Invoker {
	requestDef := GenReqDefForListIssueRecordsV4()
	return &ListIssueRecordsV4Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIssuesV4 高级查询工作项
//
// 根据筛选条件查询工作项
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) ListIssuesV4(request *model.ListIssuesV4Request) (*model.ListIssuesV4Response, error) {
	requestDef := GenReqDefForListIssuesV4()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIssuesV4Response), nil
	}
}

// ListIssuesV4Invoker 高级查询工作项
func (c *ProjectManClient) ListIssuesV4Invoker(request *model.ListIssuesV4Request) *ListIssuesV4Invoker {
	requestDef := GenReqDefForListIssuesV4()
	return &ListIssuesV4Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIterationHistories 查看迭代历史记录
//
// 查看迭代历史记录
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) ListIterationHistories(request *model.ListIterationHistoriesRequest) (*model.ListIterationHistoriesResponse, error) {
	requestDef := GenReqDefForListIterationHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIterationHistoriesResponse), nil
	}
}

// ListIterationHistoriesInvoker 查看迭代历史记录
func (c *ProjectManClient) ListIterationHistoriesInvoker(request *model.ListIterationHistoriesRequest) *ListIterationHistoriesInvoker {
	requestDef := GenReqDefForListIterationHistories()
	return &ListIterationHistoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectIterationsV4 获取指定项目的迭代列表
//
// 获取项目迭代
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) ListProjectIterationsV4(request *model.ListProjectIterationsV4Request) (*model.ListProjectIterationsV4Response, error) {
	requestDef := GenReqDefForListProjectIterationsV4()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectIterationsV4Response), nil
	}
}

// ListProjectIterationsV4Invoker 获取指定项目的迭代列表
func (c *ProjectManClient) ListProjectIterationsV4Invoker(request *model.ListProjectIterationsV4Request) *ListProjectIterationsV4Invoker {
	requestDef := GenReqDefForListProjectIterationsV4()
	return &ListProjectIterationsV4Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectWorkHours 按用户查询工时（多项目）
//
// 按用户查询工时（多项目）
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) ListProjectWorkHours(request *model.ListProjectWorkHoursRequest) (*model.ListProjectWorkHoursResponse, error) {
	requestDef := GenReqDefForListProjectWorkHours()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectWorkHoursResponse), nil
	}
}

// ListProjectWorkHoursInvoker 按用户查询工时（多项目）
func (c *ProjectManClient) ListProjectWorkHoursInvoker(request *model.ListProjectWorkHoursRequest) *ListProjectWorkHoursInvoker {
	requestDef := GenReqDefForListProjectWorkHours()
	return &ListProjectWorkHoursInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowIssueCompletionRate 获取工作项完成率
//
// 获取工作项的完成率
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) ShowIssueCompletionRate(request *model.ShowIssueCompletionRateRequest) (*model.ShowIssueCompletionRateResponse, error) {
	requestDef := GenReqDefForShowIssueCompletionRate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIssueCompletionRateResponse), nil
	}
}

// ShowIssueCompletionRateInvoker 获取工作项完成率
func (c *ProjectManClient) ShowIssueCompletionRateInvoker(request *model.ShowIssueCompletionRateRequest) *ShowIssueCompletionRateInvoker {
	requestDef := GenReqDefForShowIssueCompletionRate()
	return &ShowIssueCompletionRateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowIssueV4 查询工作项详情
//
// 查询工作项详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) ShowIssueV4(request *model.ShowIssueV4Request) (*model.ShowIssueV4Response, error) {
	requestDef := GenReqDefForShowIssueV4()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIssueV4Response), nil
	}
}

// ShowIssueV4Invoker 查询工作项详情
func (c *ProjectManClient) ShowIssueV4Invoker(request *model.ShowIssueV4Request) *ShowIssueV4Invoker {
	requestDef := GenReqDefForShowIssueV4()
	return &ShowIssueV4Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowIterationV4 查看迭代详情
//
// 查看迭代详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) ShowIterationV4(request *model.ShowIterationV4Request) (*model.ShowIterationV4Response, error) {
	requestDef := GenReqDefForShowIterationV4()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIterationV4Response), nil
	}
}

// ShowIterationV4Invoker 查看迭代详情
func (c *ProjectManClient) ShowIterationV4Invoker(request *model.ShowIterationV4Request) *ShowIterationV4Invoker {
	requestDef := GenReqDefForShowIterationV4()
	return &ShowIterationV4Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectWorkHours 按用户查询工时（单项目）
//
// 按用户查询工时（单项目）
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) ShowProjectWorkHours(request *model.ShowProjectWorkHoursRequest) (*model.ShowProjectWorkHoursResponse, error) {
	requestDef := GenReqDefForShowProjectWorkHours()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectWorkHoursResponse), nil
	}
}

// ShowProjectWorkHoursInvoker 按用户查询工时（单项目）
func (c *ProjectManClient) ShowProjectWorkHoursInvoker(request *model.ShowProjectWorkHoursRequest) *ShowProjectWorkHoursInvoker {
	requestDef := GenReqDefForShowProjectWorkHours()
	return &ShowProjectWorkHoursInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateIssueV4 更新工作项
//
// 更新工作项
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) UpdateIssueV4(request *model.UpdateIssueV4Request) (*model.UpdateIssueV4Response, error) {
	requestDef := GenReqDefForUpdateIssueV4()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateIssueV4Response), nil
	}
}

// UpdateIssueV4Invoker 更新工作项
func (c *ProjectManClient) UpdateIssueV4Invoker(request *model.UpdateIssueV4Request) *UpdateIssueV4Invoker {
	requestDef := GenReqDefForUpdateIssueV4()
	return &UpdateIssueV4Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateIterationV4 更新Scrum项目迭代
//
// 更新Scrum项目迭代
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) UpdateIterationV4(request *model.UpdateIterationV4Request) (*model.UpdateIterationV4Response, error) {
	requestDef := GenReqDefForUpdateIterationV4()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateIterationV4Response), nil
	}
}

// UpdateIterationV4Invoker 更新Scrum项目迭代
func (c *ProjectManClient) UpdateIterationV4Invoker(request *model.UpdateIterationV4Request) *UpdateIterationV4Invoker {
	requestDef := GenReqDefForUpdateIterationV4()
	return &UpdateIterationV4Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadIssueImg 上传图片
//
// 上传图片
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ProjectManClient) UploadIssueImg(request *model.UploadIssueImgRequest) (*model.UploadIssueImgResponse, error) {
	requestDef := GenReqDefForUploadIssueImg()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadIssueImgResponse), nil
	}
}

// UploadIssueImgInvoker 上传图片
func (c *ProjectManClient) UploadIssueImgInvoker(request *model.UploadIssueImgRequest) *UploadIssueImgInvoker {
	requestDef := GenReqDefForUploadIssueImg()
	return &UploadIssueImgInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
