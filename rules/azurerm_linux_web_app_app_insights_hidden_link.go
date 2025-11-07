package rules

import (
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermLinuxWebAppAppInsightsHiddenLinkRule checks whether lifecycle ignore_changes includes hidden-link tags when Application Insights is configured
type AzurermLinuxWebAppAppInsightsHiddenLinkRule struct {
	tflint.DefaultRule
}

const (
	ignoreChangesAttrName    = "ignore_changes"
	appSettingsAttrName      = "app_settings"
	appInsightsConnectionKey = "APPLICATIONINSIGHTS_CONNECTION_STRING"
	appInsightsInstrumentKey = "APPINSIGHTS_INSTRUMENTATIONKEY"
)

var requiredHiddenLinkTags = []string{
	"hidden-link: /app-insights-conn-string",
	"hidden-link: /app-insights-instrumentation-key",
	"hidden-link: /app-insights-resource-id",
}

// NewAzurermLinuxWebAppAppInsightsHiddenLinkRule returns new rule for checking Application Insights hidden-link configuration
func NewAzurermLinuxWebAppAppInsightsHiddenLinkRule() *AzurermLinuxWebAppAppInsightsHiddenLinkRule {
	return &AzurermLinuxWebAppAppInsightsHiddenLinkRule{}
}

// Name returns the rule name
func (r *AzurermLinuxWebAppAppInsightsHiddenLinkRule) Name() string {
	return "azurerm_linux_web_app_app_insights_hidden_link"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermLinuxWebAppAppInsightsHiddenLinkRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermLinuxWebAppAppInsightsHiddenLinkRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *AzurermLinuxWebAppAppInsightsHiddenLinkRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// checkResourceType checks a specific resource type for Application Insights hidden-link configuration
func (r *AzurermLinuxWebAppAppInsightsHiddenLinkRule) checkResourceType(runner tflint.Runner, resourceType string) error {
	resources, err := runner.GetResourceContent(resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: appSettingsAttrName},
		},
		Blocks: []hclext.BlockSchema{
			{
				Type: "lifecycle",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: ignoreChangesAttrName},
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
		if appSettingsAttr, exists := resource.Body.Attributes[appSettingsAttrName]; exists {
			err := runner.EvaluateExpr(appSettingsAttr.Expr, func(val map[string]string) error {
				for key := range val {
					if strings.EqualFold(key, appInsightsConnectionKey) || strings.EqualFold(key, appInsightsInstrumentKey) {
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
				if ignoreChangesAttr, ok := block.Body.Attributes[ignoreChangesAttrName]; ok {
					err := runner.EvaluateExpr(ignoreChangesAttr.Expr, func(val []string) error {
						// Check if all required hidden-link tags are included in ignore_changes
						foundTags := 0
						for _, requiredTag := range requiredHiddenLinkTags {
							for _, ignoreChange := range val {
								if strings.Contains(ignoreChange, requiredTag) {
									foundTags++
									break
								}
							}
						}
						if foundTags == len(requiredHiddenLinkTags) {
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
func (r *AzurermLinuxWebAppAppInsightsHiddenLinkRule) Check(runner tflint.Runner) error {
	resourceTypes := []string{
		"azurerm_linux_web_app",
		"azurerm_linux_web_app_slot",
	}

	for _, resourceType := range resourceTypes {
		if err := r.checkResourceType(runner, resourceType); err != nil {
			return err
		}
	}

	return nil
}
