mapping "azurerm_advanced_threat_protection" {
  import_path = "azure-rest-api-specs/specification/security/resource-manager/Microsoft.Security/stable/2019-01-01/advancedThreatProtectionSettings.json"

  enabled = AdvancedThreatProtectionProperties.isEnabled
}

mapping "azurerm_security_center_subscription_pricing" {
  import_path = "azure-rest-api-specs/specification/security/resource-manager/Microsoft.Security/stable/2018-06-01/pricings.json"

  tier = PricingProperties.pricingTier
}
