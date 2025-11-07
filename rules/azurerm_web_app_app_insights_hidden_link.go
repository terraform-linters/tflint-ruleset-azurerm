package rules

import (
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermWebAppAppInsightsHiddenLinkRule checks whether lifecycle ignore_changes includes hidden-link tags when Application Insights is configured
type AzurermWebAppAppInsightsHiddenLinkRule struct {
	tflint.DefaultRule
}

const (
	webAppIgnoreChangesAttrName    = "ignore_changes"
	webAppAppSettingsAttrName      = "app_settings"
	webAppAppInsightsConnectionKey = "APPLICATIONINSIGHTS_CONNECTION_STRING"
	webAppAppInsightsInstrumentKey = "APPINSIGHTS_INSTRUMENTATIONKEY"
)

var webAppRequiredHiddenLinkTags = []string{
	"hidden-link: /app-insights-conn-string",
	"hidden-link: /app-insights-instrumentation-key",
	"hidden-link: /app-insights-resource-id",
}

var webAppResourceTypes = []string{
	"azurerm_linux_web_app",
	"azurerm_linux_web_app_slot",
	"azurerm_windows_web_app",
	"azurerm_windows_web_app_slot",
}

// NewAzurermWebAppAppInsightsHiddenLinkRule returns new rule for checking Application Insights hidden-link configuration
func NewAzurermWebAppAppInsightsHiddenLinkRule() *AzurermWebAppAppInsightsHiddenLinkRule {
	return &AzurermWebAppAppInsightsHiddenLinkRule{}
}

// Name returns the rule name
func (r *AzurermWebAppAppInsightsHiddenLinkRule) Name() string {
	return "azurerm_web_app_app_insights_hidden_link"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermWebAppAppInsightsHiddenLinkRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermWebAppAppInsightsHiddenLinkRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *AzurermWebAppAppInsightsHiddenLinkRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// checkResourceType checks a specific resource type for Application Insights hidden-link configuration
func (r *AzurermWebAppAppInsightsHiddenLinkRule) checkResourceType(runner tflint.Runner, resourceType string) error {
	resources, err := runner.GetResourceContent(resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: webAppAppSettingsAttrName},
		},
		Blocks: []hclext.BlockSchema{
			{
				Type: "lifecycle",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: webAppIgnoreChangesAttrName},
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
		if appSettingsAttr, exists := resource.Body.Attributes[webAppAppSettingsAttrName]; exists {
			err := runner.EvaluateExpr(appSettingsAttr.Expr, func(val map[string]string) error {
				for key := range val {
					if strings.EqualFold(key, webAppAppInsightsConnectionKey) || strings.EqualFold(key, webAppAppInsightsInstrumentKey) {
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
				if ignoreChangesAttr, ok := block.Body.Attributes[webAppIgnoreChangesAttrName]; ok {
					err := runner.EvaluateExpr(ignoreChangesAttr.Expr, func(val []string) error {
						// Check if all required hidden-link tags are included in ignore_changes
						foundTags := 0
						for _, requiredTag := range webAppRequiredHiddenLinkTags {
							for _, ignoreChange := range val {
								if strings.Contains(ignoreChange, requiredTag) {
									foundTags++
									break
								}
							}
						}
						if foundTags == len(webAppRequiredHiddenLinkTags) {
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
func (r *AzurermWebAppAppInsightsHiddenLinkRule) Check(runner tflint.Runner) error {
	for _, resourceType := range webAppResourceTypes {
		if err := r.checkResourceType(runner, resourceType); err != nil {
			return err
		}
	}

	return nil
}
