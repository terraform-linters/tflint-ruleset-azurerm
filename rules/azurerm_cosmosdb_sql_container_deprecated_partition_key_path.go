package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// NewAzurermCosmosdbSQLContainerDeprecatedPartitionKeyPathRule creates a new rule to check for the deprecated attribute partition_key_path.
func NewAzurermCosmosdbSQLContainerDeprecatedPartitionKeyPathRule() *AzurermCosmosdbSQLContainerDeprecatedPartitionKeyPathRule {
    return &AzurermCosmosdbSQLContainerDeprecatedPartitionKeyPathRule{
        resourceType:  "azurerm_cosmosdb_sql_container",
        attributeName: "partition_key_path",
    }
}

// AzurermCosmosdbSQLContainerDeprecatedPartitionKeyPathRule defines a rule that checks for the deprecated partition_key_path in azurerm_cosmosdb_sql_container.
type AzurermCosmosdbSQLContainerDeprecatedPartitionKeyPathRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// Name returns the name of the rule.
func (r *AzurermCosmosdbSQLContainerDeprecatedPartitionKeyPathRule) Name() string {
	return "azurerm_cosmosdb_sql_container_deprecated_partition_key_path"
}

// Enabled returns whether the rule is enabled by default.
func (r *AzurermCosmosdbSQLContainerDeprecatedPartitionKeyPathRule) Enabled() bool {
	return true
}

// Severity returns the severity level of the rule.
func (r *AzurermCosmosdbSQLContainerDeprecatedPartitionKeyPathRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the documentation link for the rule.
func (r *AzurermCosmosdbSQLContainerDeprecatedPartitionKeyPathRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check examines azurerm_cosmosdb_sql_container resources and emits a warning if the deprecated partition_key_path is used.
func (r *AzurermCosmosdbSQLContainerDeprecatedPartitionKeyPathRule) Check(runner tflint.Runner) error {
    resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
        Attributes: []hclext.AttributeSchema{
            {Name: r.attributeName},
        },
    }, nil)
    if err != nil {
        return err
    }

    for _, resource := range resources.Blocks {
        fmt.Printf("Processing resource: %+v\n", resource)
        attribute, exists := resource.Body.Attributes[r.attributeName]
        if !exists {
            continue
        }

        fmt.Printf("Found deprecated attribute: %+v\n", attribute)

        runner.EmitIssue(
            r,
            "`partition_key_path` is deprecated and should be replaced with `partition_key_paths`.",
            attribute.Expr.Range(),
        )
    }
    return nil
}
