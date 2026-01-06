package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermAppServiceMissingAutoHealSettingRule checks whether auto_heal_setting is configured in site_config
type AzurermAppServiceMissingAutoHealSettingRule struct {
	tflint.DefaultRule
}

const (
	autoHealSettingBlockName = "auto_heal_setting"
)

var autoHealResourceTypes = []string{
	"azurerm_linux_web_app",
	"azurerm_linux_web_app_slot",
	"azurerm_windows_web_app",
	"azurerm_windows_web_app_slot",
}

// NewAzurermAppServiceMissingAutoHealSettingRule returns new rule for checking auto_heal_setting configuration
func NewAzurermAppServiceMissingAutoHealSettingRule() *AzurermAppServiceMissingAutoHealSettingRule {
	return &AzurermAppServiceMissingAutoHealSettingRule{}
}

// Name returns the rule name
func (r *AzurermAppServiceMissingAutoHealSettingRule) Name() string {
	return "azurerm_app_service_missing_auto_heal_setting"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermAppServiceMissingAutoHealSettingRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermAppServiceMissingAutoHealSettingRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *AzurermAppServiceMissingAutoHealSettingRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// checkResourceType checks a specific resource type for auto_heal_setting configuration
func (r *AzurermAppServiceMissingAutoHealSettingRule) checkResourceType(runner tflint.Runner, resourceType string) error {
	resources, err := runner.GetResourceContent(resourceType, &hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type: "site_config",
				Body: &hclext.BodySchema{
					Blocks: []hclext.BlockSchema{
						{
							Type: autoHealSettingBlockName,
						},
					},
				},
			},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		logger.Debug("checking", "resource type", resource.Labels[0], "resource name", resource.Labels[1])

		// Check if site_config block exists
		hasSiteConfig := false
		hasAutoHealSetting := false

		for _, block := range resource.Body.Blocks {
			if block.Type == "site_config" {
				hasSiteConfig = true

				// Check for auto_heal_setting block
				for _, siteConfigBlock := range block.Body.Blocks {
					if siteConfigBlock.Type == "auto_heal_setting" {
						hasAutoHealSetting = true
						break
					}
				}

				break
			}
		}

		// If site_config doesn't exist, skip this resource
		if !hasSiteConfig {
			logger.Debug("no site_config block found", "resource type", resource.Labels[0], "resource name", resource.Labels[1])
			continue
		}

		// Emit issue if auto_heal_setting is not configured in site_config
		if !hasAutoHealSetting {
			issue := "auto_heal_setting should be configured in site_config block for robust app services."
			if err := runner.EmitIssue(r, issue, resource.DefRange); err != nil {
				return err
			}
		}
	}

	return nil
}

// Check checks whether auto_heal_setting is configured in site_config
func (r *AzurermAppServiceMissingAutoHealSettingRule) Check(runner tflint.Runner) error {
	for _, resourceType := range autoHealResourceTypes {
		if err := r.checkResourceType(runner, resourceType); err != nil {
			return err
		}
	}

	return nil
}
