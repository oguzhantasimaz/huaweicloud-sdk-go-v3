package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rgc/v1/model"
)

type RgcClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewRgcClient(hcClient *httpclient.HcHttpClient) *RgcClient {
	return &RgcClient{HcClient: hcClient}
}

func RgcClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// DisableControl 关闭控制策略
//
// 关闭组织下的某个单元的某个控制策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) DisableControl(request *model.DisableControlRequest) (*model.DisableControlResponse, error) {
	requestDef := GenReqDefForDisableControl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableControlResponse), nil
	}
}

// DisableControlInvoker 关闭控制策略
func (c *RgcClient) DisableControlInvoker(request *model.DisableControlRequest) *DisableControlInvoker {
	requestDef := GenReqDefForDisableControl()
	return &DisableControlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableControl 开启控制策略
//
// 给组织下的某个单元开启某个控制策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) EnableControl(request *model.EnableControlRequest) (*model.EnableControlResponse, error) {
	requestDef := GenReqDefForEnableControl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableControlResponse), nil
	}
}

// EnableControlInvoker 开启控制策略
func (c *RgcClient) EnableControlInvoker(request *model.EnableControlRequest) *EnableControlInvoker {
	requestDef := GenReqDefForEnableControl()
	return &EnableControlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConfigRuleCompliance 查询纳管账号的Config规则合规性信息
//
// 查询纳管账号的Config规则合规性信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ListConfigRuleCompliance(request *model.ListConfigRuleComplianceRequest) (*model.ListConfigRuleComplianceResponse, error) {
	requestDef := GenReqDefForListConfigRuleCompliance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConfigRuleComplianceResponse), nil
	}
}

// ListConfigRuleComplianceInvoker 查询纳管账号的Config规则合规性信息
func (c *RgcClient) ListConfigRuleComplianceInvoker(request *model.ListConfigRuleComplianceRequest) *ListConfigRuleComplianceInvoker {
	requestDef := GenReqDefForListConfigRuleCompliance()
	return &ListConfigRuleComplianceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListControlViolations 列出不合规信息
//
// 列出组织里所有不合规的资源信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ListControlViolations(request *model.ListControlViolationsRequest) (*model.ListControlViolationsResponse, error) {
	requestDef := GenReqDefForListControlViolations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListControlViolationsResponse), nil
	}
}

// ListControlViolationsInvoker 列出不合规信息
func (c *RgcClient) ListControlViolationsInvoker(request *model.ListControlViolationsRequest) *ListControlViolationsInvoker {
	requestDef := GenReqDefForListControlViolations()
	return &ListControlViolationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListControls 列出控制策略
//
// 列出RGC服务里所有的预置控制策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ListControls(request *model.ListControlsRequest) (*model.ListControlsResponse, error) {
	requestDef := GenReqDefForListControls()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListControlsResponse), nil
	}
}

// ListControlsInvoker 列出控制策略
func (c *RgcClient) ListControlsInvoker(request *model.ListControlsRequest) *ListControlsInvoker {
	requestDef := GenReqDefForListControls()
	return &ListControlsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListControlsForAccount 列出纳管账号下开启的控制策略
//
// 列出组织里某个纳管账号开启的所有控制策略信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ListControlsForAccount(request *model.ListControlsForAccountRequest) (*model.ListControlsForAccountResponse, error) {
	requestDef := GenReqDefForListControlsForAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListControlsForAccountResponse), nil
	}
}

// ListControlsForAccountInvoker 列出纳管账号下开启的控制策略
func (c *RgcClient) ListControlsForAccountInvoker(request *model.ListControlsForAccountRequest) *ListControlsForAccountInvoker {
	requestDef := GenReqDefForListControlsForAccount()
	return &ListControlsForAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListControlsForOrganizationUnit 列出注册OU下开启的控制策略
//
// 列出组织里某个注册OU开启的所有控制策略信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ListControlsForOrganizationUnit(request *model.ListControlsForOrganizationUnitRequest) (*model.ListControlsForOrganizationUnitResponse, error) {
	requestDef := GenReqDefForListControlsForOrganizationUnit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListControlsForOrganizationUnitResponse), nil
	}
}

// ListControlsForOrganizationUnitInvoker 列出注册OU下开启的控制策略
func (c *RgcClient) ListControlsForOrganizationUnitInvoker(request *model.ListControlsForOrganizationUnitRequest) *ListControlsForOrganizationUnitInvoker {
	requestDef := GenReqDefForListControlsForOrganizationUnit()
	return &ListControlsForOrganizationUnitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDriftDetails 列出漂移信息
//
// 列出Landing Zone的所有漂移详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ListDriftDetails(request *model.ListDriftDetailsRequest) (*model.ListDriftDetailsResponse, error) {
	requestDef := GenReqDefForListDriftDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDriftDetailsResponse), nil
	}
}

// ListDriftDetailsInvoker 列出漂移信息
func (c *RgcClient) ListDriftDetailsInvoker(request *model.ListDriftDetailsRequest) *ListDriftDetailsInvoker {
	requestDef := GenReqDefForListDriftDetails()
	return &ListDriftDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEnabledControls 列出开启的控制策略
//
// 列出组织里开启的所有控制策略信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ListEnabledControls(request *model.ListEnabledControlsRequest) (*model.ListEnabledControlsResponse, error) {
	requestDef := GenReqDefForListEnabledControls()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnabledControlsResponse), nil
	}
}

// ListEnabledControlsInvoker 列出开启的控制策略
func (c *RgcClient) ListEnabledControlsInvoker(request *model.ListEnabledControlsRequest) *ListEnabledControlsInvoker {
	requestDef := GenReqDefForListEnabledControls()
	return &ListEnabledControlsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowComplianceStatusForAccount 查询纳管账号的合规状态
//
// 查询组织里某个纳管账号的资源合规状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowComplianceStatusForAccount(request *model.ShowComplianceStatusForAccountRequest) (*model.ShowComplianceStatusForAccountResponse, error) {
	requestDef := GenReqDefForShowComplianceStatusForAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowComplianceStatusForAccountResponse), nil
	}
}

// ShowComplianceStatusForAccountInvoker 查询纳管账号的合规状态
func (c *RgcClient) ShowComplianceStatusForAccountInvoker(request *model.ShowComplianceStatusForAccountRequest) *ShowComplianceStatusForAccountInvoker {
	requestDef := GenReqDefForShowComplianceStatusForAccount()
	return &ShowComplianceStatusForAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowComplianceStatusForOrganizationUnit 查询注册OU的合规状态
//
// 查询组织里某个注册OU下所有纳管账号的资源合规状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowComplianceStatusForOrganizationUnit(request *model.ShowComplianceStatusForOrganizationUnitRequest) (*model.ShowComplianceStatusForOrganizationUnitResponse, error) {
	requestDef := GenReqDefForShowComplianceStatusForOrganizationUnit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowComplianceStatusForOrganizationUnitResponse), nil
	}
}

// ShowComplianceStatusForOrganizationUnitInvoker 查询注册OU的合规状态
func (c *RgcClient) ShowComplianceStatusForOrganizationUnitInvoker(request *model.ShowComplianceStatusForOrganizationUnitRequest) *ShowComplianceStatusForOrganizationUnitInvoker {
	requestDef := GenReqDefForShowComplianceStatusForOrganizationUnit()
	return &ShowComplianceStatusForOrganizationUnitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowControl 查询控制策略详细信息
//
// 查询单个预置的控制策略详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowControl(request *model.ShowControlRequest) (*model.ShowControlResponse, error) {
	requestDef := GenReqDefForShowControl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowControlResponse), nil
	}
}

// ShowControlInvoker 查询控制策略详细信息
func (c *RgcClient) ShowControlInvoker(request *model.ShowControlRequest) *ShowControlInvoker {
	requestDef := GenReqDefForShowControl()
	return &ShowControlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowControlOperate 查询控制策略操作状态
//
// 根据操作ID查询返回指定ID的操作状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowControlOperate(request *model.ShowControlOperateRequest) (*model.ShowControlOperateResponse, error) {
	requestDef := GenReqDefForShowControlOperate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowControlOperateResponse), nil
	}
}

// ShowControlOperateInvoker 查询控制策略操作状态
func (c *RgcClient) ShowControlOperateInvoker(request *model.ShowControlOperateRequest) *ShowControlOperateInvoker {
	requestDef := GenReqDefForShowControlOperate()
	return &ShowControlOperateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowControlsForOrganizationUnit 查询注册OU开启的控制策略
//
// 查询组织里某个注册OU下开启的某个控制策略的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowControlsForOrganizationUnit(request *model.ShowControlsForOrganizationUnitRequest) (*model.ShowControlsForOrganizationUnitResponse, error) {
	requestDef := GenReqDefForShowControlsForOrganizationUnit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowControlsForOrganizationUnitResponse), nil
	}
}

// ShowControlsForOrganizationUnitInvoker 查询注册OU开启的控制策略
func (c *RgcClient) ShowControlsForOrganizationUnitInvoker(request *model.ShowControlsForOrganizationUnitRequest) *ShowControlsForOrganizationUnitInvoker {
	requestDef := GenReqDefForShowControlsForOrganizationUnit()
	return &ShowControlsForOrganizationUnitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckLaunch 设置Landing Zone前检查
//
// 在设置Landing Zone之前，检查当前区域是否可以设置Landing Zone。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) CheckLaunch(request *model.CheckLaunchRequest) (*model.CheckLaunchResponse, error) {
	requestDef := GenReqDefForCheckLaunch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckLaunchResponse), nil
	}
}

// CheckLaunchInvoker 设置Landing Zone前检查
func (c *RgcClient) CheckLaunchInvoker(request *model.CheckLaunchRequest) *CheckLaunchInvoker {
	requestDef := GenReqDefForCheckLaunch()
	return &CheckLaunchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetupLandingZone 设置Landing Zone
//
// 在当前区域创建或者更新Landing Zone。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) SetupLandingZone(request *model.SetupLandingZoneRequest) (*model.SetupLandingZoneResponse, error) {
	requestDef := GenReqDefForSetupLandingZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetupLandingZoneResponse), nil
	}
}

// SetupLandingZoneInvoker 设置Landing Zone
func (c *RgcClient) SetupLandingZoneInvoker(request *model.SetupLandingZoneRequest) *SetupLandingZoneInvoker {
	requestDef := GenReqDefForSetupLandingZone()
	return &SetupLandingZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAvailableUpdates 查询Landing Zone可更新状态
//
// 查询Landing Zone当前是否需要升级更新。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowAvailableUpdates(request *model.ShowAvailableUpdatesRequest) (*model.ShowAvailableUpdatesResponse, error) {
	requestDef := GenReqDefForShowAvailableUpdates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAvailableUpdatesResponse), nil
	}
}

// ShowAvailableUpdatesInvoker 查询Landing Zone可更新状态
func (c *RgcClient) ShowAvailableUpdatesInvoker(request *model.ShowAvailableUpdatesRequest) *ShowAvailableUpdatesInvoker {
	requestDef := GenReqDefForShowAvailableUpdates()
	return &ShowAvailableUpdatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHomeRegion 查询主区域
//
// 查询Landing Zone的主区域。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowHomeRegion(request *model.ShowHomeRegionRequest) (*model.ShowHomeRegionResponse, error) {
	requestDef := GenReqDefForShowHomeRegion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHomeRegionResponse), nil
	}
}

// ShowHomeRegionInvoker 查询主区域
func (c *RgcClient) ShowHomeRegionInvoker(request *model.ShowHomeRegionRequest) *ShowHomeRegionInvoker {
	requestDef := GenReqDefForShowHomeRegion()
	return &ShowHomeRegionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLandingZoneConfiguration 查询Landing Zone的配置
//
// 查询当前客户的Landing Zone的所有配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowLandingZoneConfiguration(request *model.ShowLandingZoneConfigurationRequest) (*model.ShowLandingZoneConfigurationResponse, error) {
	requestDef := GenReqDefForShowLandingZoneConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLandingZoneConfigurationResponse), nil
	}
}

// ShowLandingZoneConfigurationInvoker 查询Landing Zone的配置
func (c *RgcClient) ShowLandingZoneConfigurationInvoker(request *model.ShowLandingZoneConfigurationRequest) *ShowLandingZoneConfigurationInvoker {
	requestDef := GenReqDefForShowLandingZoneConfiguration()
	return &ShowLandingZoneConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLandingZoneIdentityCenter 查询当前客户的Identity Center用户信息
//
// 查询当前客户的Identity Center用户信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowLandingZoneIdentityCenter(request *model.ShowLandingZoneIdentityCenterRequest) (*model.ShowLandingZoneIdentityCenterResponse, error) {
	requestDef := GenReqDefForShowLandingZoneIdentityCenter()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLandingZoneIdentityCenterResponse), nil
	}
}

// ShowLandingZoneIdentityCenterInvoker 查询当前客户的Identity Center用户信息
func (c *RgcClient) ShowLandingZoneIdentityCenterInvoker(request *model.ShowLandingZoneIdentityCenterRequest) *ShowLandingZoneIdentityCenterInvoker {
	requestDef := GenReqDefForShowLandingZoneIdentityCenter()
	return &ShowLandingZoneIdentityCenterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLandingZoneStatus 查询Landing Zone设置状态
//
// 查询Landing Zone的设置状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowLandingZoneStatus(request *model.ShowLandingZoneStatusRequest) (*model.ShowLandingZoneStatusResponse, error) {
	requestDef := GenReqDefForShowLandingZoneStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLandingZoneStatusResponse), nil
	}
}

// ShowLandingZoneStatusInvoker 查询Landing Zone设置状态
func (c *RgcClient) ShowLandingZoneStatusInvoker(request *model.ShowLandingZoneStatusRequest) *ShowLandingZoneStatusInvoker {
	requestDef := GenReqDefForShowLandingZoneStatus()
	return &ShowLandingZoneStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAccount 创建账号
//
// 在组织里的某个注册OU下创建账号。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) CreateAccount(request *model.CreateAccountRequest) (*model.CreateAccountResponse, error) {
	requestDef := GenReqDefForCreateAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAccountResponse), nil
	}
}

// CreateAccountInvoker 创建账号
func (c *RgcClient) CreateAccountInvoker(request *model.CreateAccountRequest) *CreateAccountInvoker {
	requestDef := GenReqDefForCreateAccount()
	return &CreateAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteManagedOrganizationalUnits 删除注册OU
//
// 在组织里删除已注册OU。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) DeleteManagedOrganizationalUnits(request *model.DeleteManagedOrganizationalUnitsRequest) (*model.DeleteManagedOrganizationalUnitsResponse, error) {
	requestDef := GenReqDefForDeleteManagedOrganizationalUnits()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteManagedOrganizationalUnitsResponse), nil
	}
}

// DeleteManagedOrganizationalUnitsInvoker 删除注册OU
func (c *RgcClient) DeleteManagedOrganizationalUnitsInvoker(request *model.DeleteManagedOrganizationalUnitsRequest) *DeleteManagedOrganizationalUnitsInvoker {
	requestDef := GenReqDefForDeleteManagedOrganizationalUnits()
	return &DeleteManagedOrganizationalUnitsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeregisterOrganizationalUnit 取消注册OU
//
// 将组织里的某个OU从RGC服务里取消注册。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) DeregisterOrganizationalUnit(request *model.DeregisterOrganizationalUnitRequest) (*model.DeregisterOrganizationalUnitResponse, error) {
	requestDef := GenReqDefForDeregisterOrganizationalUnit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeregisterOrganizationalUnitResponse), nil
	}
}

// DeregisterOrganizationalUnitInvoker 取消注册OU
func (c *RgcClient) DeregisterOrganizationalUnitInvoker(request *model.DeregisterOrganizationalUnitRequest) *DeregisterOrganizationalUnitInvoker {
	requestDef := GenReqDefForDeregisterOrganizationalUnit()
	return &DeregisterOrganizationalUnitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnrollAccount 纳管账号
//
// 将组织里的某个账号纳管到RGC服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) EnrollAccount(request *model.EnrollAccountRequest) (*model.EnrollAccountResponse, error) {
	requestDef := GenReqDefForEnrollAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnrollAccountResponse), nil
	}
}

// EnrollAccountInvoker 纳管账号
func (c *RgcClient) EnrollAccountInvoker(request *model.EnrollAccountRequest) *EnrollAccountInvoker {
	requestDef := GenReqDefForEnrollAccount()
	return &EnrollAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListManagedAccounts 列举控制策略生效的纳管账号信息
//
// 列举控制策略生效的纳管账号信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ListManagedAccounts(request *model.ListManagedAccountsRequest) (*model.ListManagedAccountsResponse, error) {
	requestDef := GenReqDefForListManagedAccounts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListManagedAccountsResponse), nil
	}
}

// ListManagedAccountsInvoker 列举控制策略生效的纳管账号信息
func (c *RgcClient) ListManagedAccountsInvoker(request *model.ListManagedAccountsRequest) *ListManagedAccountsInvoker {
	requestDef := GenReqDefForListManagedAccounts()
	return &ListManagedAccountsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListManagedAccountsForParent 列出注册OU下的纳管账号信息
//
// 列出组织里某个注册OU下的所有纳管账号信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ListManagedAccountsForParent(request *model.ListManagedAccountsForParentRequest) (*model.ListManagedAccountsForParentResponse, error) {
	requestDef := GenReqDefForListManagedAccountsForParent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListManagedAccountsForParentResponse), nil
	}
}

// ListManagedAccountsForParentInvoker 列出注册OU下的纳管账号信息
func (c *RgcClient) ListManagedAccountsForParentInvoker(request *model.ListManagedAccountsForParentRequest) *ListManagedAccountsForParentInvoker {
	requestDef := GenReqDefForListManagedAccountsForParent()
	return &ListManagedAccountsForParentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListManagedOrganizationalUnits 列举控制策略生效的注册OU信息
//
// 列举控制策略生效的注册OU信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ListManagedOrganizationalUnits(request *model.ListManagedOrganizationalUnitsRequest) (*model.ListManagedOrganizationalUnitsResponse, error) {
	requestDef := GenReqDefForListManagedOrganizationalUnits()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListManagedOrganizationalUnitsResponse), nil
	}
}

// ListManagedOrganizationalUnitsInvoker 列举控制策略生效的注册OU信息
func (c *RgcClient) ListManagedOrganizationalUnitsInvoker(request *model.ListManagedOrganizationalUnitsRequest) *ListManagedOrganizationalUnitsInvoker {
	requestDef := GenReqDefForListManagedOrganizationalUnits()
	return &ListManagedOrganizationalUnitsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ReRegisterOrganizationalUnit 重新注册OU
//
// 重新注册组织里的某个OU到RGC服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ReRegisterOrganizationalUnit(request *model.ReRegisterOrganizationalUnitRequest) (*model.ReRegisterOrganizationalUnitResponse, error) {
	requestDef := GenReqDefForReRegisterOrganizationalUnit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ReRegisterOrganizationalUnitResponse), nil
	}
}

// ReRegisterOrganizationalUnitInvoker 重新注册OU
func (c *RgcClient) ReRegisterOrganizationalUnitInvoker(request *model.ReRegisterOrganizationalUnitRequest) *ReRegisterOrganizationalUnitInvoker {
	requestDef := GenReqDefForReRegisterOrganizationalUnit()
	return &ReRegisterOrganizationalUnitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RegisterOrganizationalUnit 注册OU
//
// 将组织里的某个OU注册到RGC服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) RegisterOrganizationalUnit(request *model.RegisterOrganizationalUnitRequest) (*model.RegisterOrganizationalUnitResponse, error) {
	requestDef := GenReqDefForRegisterOrganizationalUnit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RegisterOrganizationalUnitResponse), nil
	}
}

// RegisterOrganizationalUnitInvoker 注册OU
func (c *RgcClient) RegisterOrganizationalUnitInvoker(request *model.RegisterOrganizationalUnitRequest) *RegisterOrganizationalUnitInvoker {
	requestDef := GenReqDefForRegisterOrganizationalUnit()
	return &RegisterOrganizationalUnitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowManagedAccount 查询纳管账号信息
//
// 查询组织里某个纳管账号信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowManagedAccount(request *model.ShowManagedAccountRequest) (*model.ShowManagedAccountResponse, error) {
	requestDef := GenReqDefForShowManagedAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowManagedAccountResponse), nil
	}
}

// ShowManagedAccountInvoker 查询纳管账号信息
func (c *RgcClient) ShowManagedAccountInvoker(request *model.ShowManagedAccountRequest) *ShowManagedAccountInvoker {
	requestDef := GenReqDefForShowManagedAccount()
	return &ShowManagedAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowManagedCoreAccount 列出核心纳管账号
//
// 列出组织里的所有核心纳管账号信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowManagedCoreAccount(request *model.ShowManagedCoreAccountRequest) (*model.ShowManagedCoreAccountResponse, error) {
	requestDef := GenReqDefForShowManagedCoreAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowManagedCoreAccountResponse), nil
	}
}

// ShowManagedCoreAccountInvoker 列出核心纳管账号
func (c *RgcClient) ShowManagedCoreAccountInvoker(request *model.ShowManagedCoreAccountRequest) *ShowManagedCoreAccountInvoker {
	requestDef := GenReqDefForShowManagedCoreAccount()
	return &ShowManagedCoreAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowManagedOrganizationalUnit 查询已注册OU信息
//
// 查询在RGC服务里的注册OU信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowManagedOrganizationalUnit(request *model.ShowManagedOrganizationalUnitRequest) (*model.ShowManagedOrganizationalUnitResponse, error) {
	requestDef := GenReqDefForShowManagedOrganizationalUnit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowManagedOrganizationalUnitResponse), nil
	}
}

// ShowManagedOrganizationalUnitInvoker 查询已注册OU信息
func (c *RgcClient) ShowManagedOrganizationalUnitInvoker(request *model.ShowManagedOrganizationalUnitRequest) *ShowManagedOrganizationalUnitInvoker {
	requestDef := GenReqDefForShowManagedOrganizationalUnit()
	return &ShowManagedOrganizationalUnitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOperation 查询注册过程信息
//
// 查询在RGC服务里注册/取消注册的过程信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowOperation(request *model.ShowOperationRequest) (*model.ShowOperationResponse, error) {
	requestDef := GenReqDefForShowOperation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOperationResponse), nil
	}
}

// ShowOperationInvoker 查询注册过程信息
func (c *RgcClient) ShowOperationInvoker(request *model.ShowOperationRequest) *ShowOperationInvoker {
	requestDef := GenReqDefForShowOperation()
	return &ShowOperationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UnEnrollAccount 取消纳管账号
//
// 将组织里的某个账号从RGC服务里取消纳管。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) UnEnrollAccount(request *model.UnEnrollAccountRequest) (*model.UnEnrollAccountResponse, error) {
	requestDef := GenReqDefForUnEnrollAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UnEnrollAccountResponse), nil
	}
}

// UnEnrollAccountInvoker 取消纳管账号
func (c *RgcClient) UnEnrollAccountInvoker(request *model.UnEnrollAccountRequest) *UnEnrollAccountInvoker {
	requestDef := GenReqDefForUnEnrollAccount()
	return &UnEnrollAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateManagedAccount 更新纳管账号
//
// 更新组织里某个已在RGC服务的纳管账号。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) UpdateManagedAccount(request *model.UpdateManagedAccountRequest) (*model.UpdateManagedAccountResponse, error) {
	requestDef := GenReqDefForUpdateManagedAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateManagedAccountResponse), nil
	}
}

// UpdateManagedAccountInvoker 更新纳管账号
func (c *RgcClient) UpdateManagedAccountInvoker(request *model.UpdateManagedAccountRequest) *UpdateManagedAccountInvoker {
	requestDef := GenReqDefForUpdateManagedAccount()
	return &UpdateManagedAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
