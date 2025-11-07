package rules

import (
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermWindowsWebAppAppInsightsHiddenLinkRule checks whether lifecycle ignore_changes includes hidden-link tags when Application Insights is configured
type AzurermWindowsWebAppAppInsightsHiddenLinkRule struct {
	tflint.DefaultRule
}

const (
	windowsIgnoreChangesAttrName    = "ignore_changes"
	windowsAppSettingsAttrName      = "app_settings"
	windowsAppInsightsConnectionKey = "APPLICATIONINSIGHTS_CONNECTION_STRING"
	windowsAppInsightsInstrumentKey = "APPINSIGHTS_INSTRUMENTATIONKEY"
)

var windowsRequiredHiddenLinkTags = []string{
	"hidden-link: /app-insights-conn-string",
	"hidden-link: /app-insights-instrumentation-key",
	"hidden-link: /app-insights-resource-id",
}

// NewAzurermWindowsWebAppAppInsightsHiddenLinkRule returns new rule for checking Application Insights hidden-link configuration
func NewAzurermWindowsWebAppAppInsightsHiddenLinkRule() *AzurermWindowsWebAppAppInsightsHiddenLinkRule {
	return &AzurermWindowsWebAppAppInsightsHiddenLinkRule{}
}

// Name returns the rule name
func (r *AzurermWindowsWebAppAppInsightsHiddenLinkRule) Name() string {
	return "azurerm_windows_web_app_app_insights_hidden_link"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermWindowsWebAppAppInsightsHiddenLinkRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermWindowsWebAppAppInsightsHiddenLinkRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *AzurermWindowsWebAppAppInsightsHiddenLinkRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether hidden-link tags are ignored when Application Insights is configured
func (r *AzurermWindowsWebAppAppInsightsHiddenLinkRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("azurerm_windows_web_app", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: windowsAppSettingsAttrName},
		},
		Blocks: []hclext.BlockSchema{
			{
				Type: "lifecycle",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: windowsIgnoreChangesAttrName},
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
		if appSettingsAttr, exists := resource.Body.Attributes[windowsAppSettingsAttrName]; exists {
			err := runner.EvaluateExpr(appSettingsAttr.Expr, func(val map[string]string) error {
				for key := range val {
					if strings.EqualFold(key, windowsAppInsightsConnectionKey) || strings.EqualFold(key, windowsAppInsightsInstrumentKey) {
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

		// If Application Insights is not configured, skip this resource
		if !hasAppInsights {
			logger.Debug("no Application Insights configuration found", "resource type", resource.Labels[0], "resource name", resource.Labels[1])
			continue
		}

		// Check if lifecycle block exists and contains proper ignore_changes
		hasProperIgnoreChanges := false

		for _, block := range resource.Body.Blocks {
			if block.Type == "lifecycle" {
				if ignoreChangesAttr, ok := block.Body.Attributes[windowsIgnoreChangesAttrName]; ok {
					err := runner.EvaluateExpr(ignoreChangesAttr.Expr, func(val []string) error {
						// Check if all required hidden-link tags are included in ignore_changes
						foundTags := 0
						for _, requiredTag := range windowsRequiredHiddenLinkTags {
							for _, ignoreChange := range val {
								if strings.Contains(ignoreChange, requiredTag) {
									foundTags++
									break
								}
							}
						}
						if foundTags == len(windowsRequiredHiddenLinkTags) {
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
