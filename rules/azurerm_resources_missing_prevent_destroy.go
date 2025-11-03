package rules

import (
	"slices"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermResourcesMissingPreventDestroyRule checks whether resources have lifecycle prevent_destroy = true
type AzurermResourcesMissingPreventDestroyRule struct {
	tflint.DefaultRule
}

type azurermResourcePreventDestroyRuleConfig struct {
	Resources []string `hclext:"resources,optional"`
	Exclude   []string `hclext:"exclude,optional"`
}

const (
	lifecycleBlockName     = "lifecycle"
	preventDestroyAttrName = "prevent_destroy"
)

var defaultPreventDestroyResources = []string{
	"azurerm_backup_container_storage_account",
	"azurerm_backup_policy_file_share",
	"azurerm_backup_policy_vm",
	"azurerm_backup_protected_vm",
	"azurerm_cosmosdb_account",
	"azurerm_cosmosdb_cassandra_cluster",
	"azurerm_cosmosdb_cassandra_table",
	"azurerm_cosmosdb_mongo_database",
	"azurerm_cosmosdb_postgresql_cluster",
	"azurerm_cosmosdb_sql_container",
	"azurerm_cosmosdb_sql_database",
	"azurerm_cosmosdb_table",
	"azurerm_databricks_workspace",
	"azurerm_eventhub",
	"azurerm_eventhub_namespace",
	"azurerm_iothub",
	"azurerm_key_vault",
	"azurerm_key_vault_certificate",
	"azurerm_key_vault_key",
	"azurerm_key_vault_secret",
	"azurerm_mssql_database",
	"azurerm_mssql_server",
	"azurerm_mssql_virtual_machine",
	"azurerm_mysql_flexible_database",
	"azurerm_mysql_flexible_server",
	"azurerm_postgresql_flexible_server",
	"azurerm_postgresql_flexible_server_database",
	"azurerm_postgresql_server",
	"azurerm_servicebus_namespace",
	"azurerm_servicebus_queue",
	"azurerm_servicebus_topic",
	"azurerm_storage_account",
	"azurerm_storage_blob",
	"azurerm_storage_container",
	"azurerm_storage_queue",
	"azurerm_storage_share",
	"azurerm_storage_share_directory",
	"azurerm_storage_share_file",
	"azurerm_storage_table",
}

// NewAzurermResourcesMissingPreventDestroyRule returns new rule for checking prevent_destroy lifecycle
func NewAzurermResourcesMissingPreventDestroyRule() *AzurermResourcesMissingPreventDestroyRule {
	return &AzurermResourcesMissingPreventDestroyRule{}
}

// Name returns the rule name
func (r *AzurermResourcesMissingPreventDestroyRule) Name() string {
	return "azurerm_resources_missing_prevent_destroy"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermResourcesMissingPreventDestroyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermResourcesMissingPreventDestroyRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *AzurermResourcesMissingPreventDestroyRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks resources for missing lifecycle prevent_destroy configuration
func (r *AzurermResourcesMissingPreventDestroyRule) Check(runner tflint.Runner) error {
	config := azurermResourcePreventDestroyRuleConfig{}

	if err := runner.DecodeRuleConfig(r.Name(), &config); err != nil {
		return err
	}

	// Use default resources if none specified in config
	resourcesToCheck := config.Resources
	if len(resourcesToCheck) == 0 {
		resourcesToCheck = defaultPreventDestroyResources
	}

	for _, resourceType := range resourcesToCheck {
		// Skip this resource if its type is excluded in configuration
		if slices.Contains(config.Exclude, resourceType) {
			logger.Debug("excluding", "resourceType", resourceType)
			continue
		}

		resources, err := runner.GetResourceContent(resourceType, &hclext.BodySchema{
			Blocks: []hclext.BlockSchema{
				{
					Type: lifecycleBlockName,
					Body: &hclext.BodySchema{
						Attributes: []hclext.AttributeSchema{
							{Name: preventDestroyAttrName},
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

			hasPreventDestroy := false

			// Check if lifecycle block exists
			for _, block := range resource.Body.Blocks {
				if block.Type == lifecycleBlockName {
					// Check if prevent_destroy attribute exists in lifecycle block
					if attribute, ok := block.Body.Attributes[preventDestroyAttrName]; ok {
						logger.Debug("found prevent_destroy attribute", "resource type", resource.Labels[0], "resource name", resource.Labels[1])

						err := runner.EvaluateExpr(attribute.Expr, func(val bool) error {
							hasPreventDestroy = true
							return nil
						}, nil)

						if err != nil {
							return err
						}
						break
					}
				}
			}

			// Emit issue if prevent_destroy is missing or set to false
			if !hasPreventDestroy {
				issue := "Resource is missing lifecycle { prevent_destroy = true }. This resource contains data that should be protected from accidental deletion."
				if err := runner.EmitIssue(r, issue, resource.DefRange); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
