package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

func NewAzurermCosmosdbSqlContainerDeprecatedPartitionKeyPathRule() *AzurermCosmosdbSqlContainerDeprecatedPartitionKeyPathRule {
    return &AzurermCosmosdbSqlContainerDeprecatedPartitionKeyPathRule{
        resourceType:  "azurerm_cosmosdb_sql_container",
        attributeName: "partition_key_path",
    }
}

type AzurermCosmosdbSqlContainerDeprecatedPartitionKeyPathRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

func (r *AzurermCosmosdbSqlContainerDeprecatedPartitionKeyPathRule) Name() string {
	return "azurerm_cosmosdb_sql_container_deprecated_partition_key_path"
}

func (r *AzurermCosmosdbSqlContainerDeprecatedPartitionKeyPathRule) Enabled() bool {
	return true
}

func (r *AzurermCosmosdbSqlContainerDeprecatedPartitionKeyPathRule) Severity() tflint.Severity {
	return tflint.WARNING
}

func (r *AzurermCosmosdbSqlContainerDeprecatedPartitionKeyPathRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *AzurermCosmosdbSqlContainerDeprecatedPartitionKeyPathRule) Check(runner tflint.Runner) error {
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
