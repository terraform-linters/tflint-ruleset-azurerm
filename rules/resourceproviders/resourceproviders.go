// Package resourceproviders contains the Azure resource provider namespaces used by azurerm resources,
// and the namespaces registered by each resource_provider_registrations mode of the azurerm provider.
//
// resources.go and registrations.go are generated from a local checkout of hashicorp/terraform-provider-azurerm:
//
//	go run -tags generators ./generator -azurerm-dir /path/to/terraform-provider-azurerm
package resourceproviders

import "strings"

// RegisteredByDefault contains the namespaces Azure registers in every subscription.
// https://learn.microsoft.com/azure/azure-resource-manager/management/azure-services-resource-providers
var RegisteredByDefault = []string{
	"Microsoft.ADHybridHealthService",
	"Microsoft.Authorization",
	"Microsoft.Billing",
	"Microsoft.ChangeSafety",
	"Microsoft.ClassicSubscription",
	"Microsoft.Commerce",
	"Microsoft.Consumption",
	"Microsoft.CostManagement",
	"Microsoft.Features",
	"Microsoft.MarketplaceOrdering",
	"Microsoft.Portal",
	"Microsoft.ResourceGraph",
	"Microsoft.ResourceNotification",
	"Microsoft.Resources",
	"Microsoft.SerialConsole",
	"microsoft.support",
}

// azureCasing contains the namespaces whose casing in the Azure Resource Manager API differs from the azurerm
// documentation or registration sets. Check the casing with: az provider show -n <namespace> --query namespace
var azureCasing = []string{
	"GitHub.Network",
	"Microsoft.Automanage",
	"Microsoft.SCVMM",
	"NGINX.NGINXPLUS",
	"microsoft.aadiam",
	"microsoft.insights",
}

var canonical = canonicalNamespaces()

// Canonical returns the namespace with the casing the Azure Resource Manager API uses, which the azurerm provider
// requires as it compares namespaces case-sensitively. It returns false when the namespace is not known.
func Canonical(namespace string) (string, bool) {
	value, ok := canonical[strings.ToLower(namespace)]
	return value, ok
}

// canonicalNamespaces maps lowercase namespaces to their canonical casing. Later sources take precedence.
func canonicalNamespaces() map[string]string {
	result := map[string]string{}
	add := func(namespaces ...string) {
		for _, namespace := range namespaces {
			result[strings.ToLower(namespace)] = namespace
		}
	}
	for _, namespaces := range Registrations {
		add(namespaces...)
	}
	for _, namespaces := range Resources {
		add(namespaces...)
	}
	add(RegisteredByDefault...)
	add(azureCasing...)
	return result
}
