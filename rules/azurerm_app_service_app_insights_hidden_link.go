package rules

import (
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermAppServiceAppInsightsHiddenLinkRule checks whether lifecycle ignore_changes includes hidden-link tags when Application Insights is configured
type AzurermAppServiceAppInsightsHiddenLinkRule struct {
	tflint.DefaultRule
}

const (
	appServiceIgnoreChangesAttrName              = "ignore_changes"
	appServiceAppSettingsAttrName                = "app_settings"
	appServiceSiteConfigAttrName                 = "site_config"
	appServiceAppInsightsConnectionKey           = "APPLICATIONINSIGHTS_CONNECTION_STRING"
	appServiceAppInsightsInstrumentKey           = "APPINSIGHTS_INSTRUMENTATIONKEY"
	appServiceSiteConfigAppInsightsConnectionKey = "application_insights_connection_string"
	appServiceSiteConfigAppInsightsKey           = "application_insights_key"
)

var appServiceRequiredHiddenLinkTags = []string{
	"hidden-link: /app-insights-conn-string",
	"hidden-link: /app-insights-instrumentation-key",
	"hidden-link: /app-insights-resource-id",
}

var appServiceResourceTypes = []string{
	"azurerm_linux_web_app",
	"azurerm_linux_web_app_slot",
	"azurerm_windows_web_app",
	"azurerm_windows_web_app_slot",
	"azurerm_linux_function_app",
	"azurerm_linux_function_app_slot",
	"azurerm_windows_function_app",
	"azurerm_windows_function_app_slot",
}

// NewAzurermAppServiceAppInsightsHiddenLinkRule returns new rule for checking Application Insights hidden-link configuration
func NewAzurermAppServiceAppInsightsHiddenLinkRule() *AzurermAppServiceAppInsightsHiddenLinkRule {
	return &AzurermAppServiceAppInsightsHiddenLinkRule{}
}

// Name returns the rule name
func (r *AzurermAppServiceAppInsightsHiddenLinkRule) Name() string {
	return "azurerm_app_service_app_insights_hidden_link"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermAppServiceAppInsightsHiddenLinkRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermAppServiceAppInsightsHiddenLinkRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *AzurermAppServiceAppInsightsHiddenLinkRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// checkResourceType checks a specific resource type for Application Insights hidden-link configuration
func (r *AzurermAppServiceAppInsightsHiddenLinkRule) checkResourceType(runner tflint.Runner, resourceType string) error {
	resources, err := runner.GetResourceContent(resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: appServiceAppSettingsAttrName},
		},
		Blocks: []hclext.BlockSchema{
			{
				Type: "lifecycle",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: appServiceIgnoreChangesAttrName},
					},
				},
			},
			{
				Type: "site_config",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: appServiceSiteConfigAppInsightsConnectionKey},
						{Name: appServiceSiteConfigAppInsightsKey},
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

		// Check if Application Insights is configured
		hasAppInsights := false

		// Check in app_settings
		if appSettingsAttr, exists := resource.Body.Attributes[appServiceAppSettingsAttrName]; exists {
			err := runner.EvaluateExpr(appSettingsAttr.Expr, func(val map[string]string) error {
				for key := range val {
					if strings.EqualFold(key, appServiceAppInsightsConnectionKey) || strings.EqualFold(key, appServiceAppInsightsInstrumentKey) {
						hasAppInsights = true
						break
					}
				}
				return nil
			}, nil)

			if err != nil {
				return err
			}
		}

		// Check in site_config block
		if !hasAppInsights {
			for _, block := range resource.Body.Blocks {
				if block.Type == "site_config" {
					// Check for application_insights_connection_string
					if attr, exists := block.Body.Attributes[appServiceSiteConfigAppInsightsConnectionKey]; exists {
						err := runner.EvaluateExpr(attr.Expr, func(val string) error {
							if val != "" {
								hasAppInsights = true
							}
							return nil
						}, nil)
						if err != nil {
							return err
						}
					}

					// Check for application_insights_key
					if !hasAppInsights {
						if attr, exists := block.Body.Attributes[appServiceSiteConfigAppInsightsKey]; exists {
							err := runner.EvaluateExpr(attr.Expr, func(val string) error {
								if val != "" {
									hasAppInsights = true
								}
								return nil
							}, nil)
							if err != nil {
								return err
							}
						}
					}
					break
				}
			}
		}

		// If Application Insights is not configured, skip this resource
		if !hasAppInsights {
			logger.Debug("no Application Insights configuration found", "resource type", resource.Labels[0], "resource name", resource.Labels[1])
			continue
		}

		// Check if lifecycle block exists and contains proper ignore_changes
		hasProperIgnoreChanges := false

		for _, block := range resource.Body.Blocks {
			if block.Type == "lifecycle" {
				if ignoreChangesAttr, ok := block.Body.Attributes[appServiceIgnoreChangesAttrName]; ok {
					err := runner.EvaluateExpr(ignoreChangesAttr.Expr, func(val []string) error {
						// Check if all required hidden-link tags are included in ignore_changes
						foundTags := 0
						for _, requiredTag := range appServiceRequiredHiddenLinkTags {
							for _, ignoreChange := range val {
								if strings.Contains(ignoreChange, requiredTag) {
									foundTags++
									break
								}
							}
						}
						if foundTags == len(appServiceRequiredHiddenLinkTags) {
							hasProperIgnoreChanges = true
						}
						return nil
					}, nil)

					if err != nil {
						return err
					}
					break
				}
			}
		}

		// Emit issue if Application Insights is configured but hidden-link tags are not properly ignored
		if !hasProperIgnoreChanges {
			issue := "When Application Insights is configured, lifecycle { ignore_changes } should include all hidden-link tags: tags[\"hidden-link: /app-insights-conn-string\"], tags[\"hidden-link: /app-insights-instrumentation-key\"], tags[\"hidden-link: /app-insights-resource-id\"]"
			if err := runner.EmitIssue(r, issue, resource.DefRange); err != nil {
				return err
			}
		}
	}

	return nil
}

// Check checks whether hidden-link tags are ignored when Application Insights is configured
func (r *AzurermAppServiceAppInsightsHiddenLinkRule) Check(runner tflint.Runner) error {
	for _, resourceType := range appServiceResourceTypes {
		if err := r.checkResourceType(runner, resourceType); err != nil {
			return err
		}
	}

	return nil
}
