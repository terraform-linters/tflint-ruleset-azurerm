mapping "azurerm_application_gateway" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/applicationGateway.json"

  zones        = ApplicationGateway.zones
  enable_http2 = ApplicationGatewayPropertiesFormat.enableHttp2
}

mapping "azurerm_bastion_host" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/bastionHost.json"

  name                = BastionHostName
  resource_group_name = ResourceGroupName
}

mapping "azurerm_express_route_circuit" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/expressRouteCircuit.json"

  service_provider_name    = ExpressRouteCircuitServiceProviderProperties.serviceProviderName
  peering_location         = ExpressRouteCircuitServiceProviderProperties.peeringLocation
  bandwidth_in_mbps        = ExpressRouteCircuitServiceProviderProperties.bandwidthInMbps
  allow_classic_operations = ExpressRouteCircuitPropertiesFormat.allowClassicOperations
}

mapping "azurerm_express_route_circuit_peering" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/expressRouteCircuit.json"

  peering_type                  = ExpressRoutePeeringType
  primary_peer_address_prefix   = ExpressRouteCircuitPeeringPropertiesFormat.primaryPeerAddressPrefix
  secondary_peer_address_prefix = ExpressRouteCircuitPeeringPropertiesFormat.secondaryPeerAddressPrefix
  vlan_id                       = ExpressRouteCircuitPeeringPropertiesFormat.vlanId
  shared_key                    = ExpressRouteCircuitPeeringPropertiesFormat.sharedKey
  peer_asn                      = ExpressRouteCircuitPeeringPropertiesFormat.peerASN
}

mapping "azurerm_express_route_gateway" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2019-12-01/expressRouteGateway.json"

  virtual_hub_id = VirtualHubId.id
}

mapping "azurerm_firewall" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/azureFirewall.json"

  zones = AzureFirewall.zones
}

mapping "azurerm_firewall_application_rule_collection" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/azureFirewall.json"

  priority = AzureFirewallApplicationRuleCollectionPropertiesFormat.priority
  action   = AzureFirewallRCActionType
}

mapping "azurerm_firewall_nat_rule_collection" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/azureFirewall.json"

  priority = AzureFirewallNatRuleCollectionProperties.priority
  action   = AzureFirewallNatRCActionType
}

mapping "azurerm_firewall_network_rule_collection" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/azureFirewall.json"

  priority = AzureFirewallNetworkRuleCollectionPropertiesFormat.priority
  action   = AzureFirewallRCActionType
}

mapping "azurerm_local_network_gateway" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/virtualNetworkGateway.json"

  gateway_address = LocalNetworkGatewayPropertiesFormat.gatewayIpAddress
}

mapping "azurerm_nat_gateway" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/natGateway.json"

  idle_timeout_in_minutes = NatGatewayPropertiesFormat.idleTimeoutInMinutes
  public_ip_address_ids   = NatGatewayPropertiesFormat.publicIpAddresses
  public_ip_prefix_ids    = NatGatewayPropertiesFormat.publicIpPrefixes
  sku_name                = NatGatewaySku.name
  zones                   = NatGateway.zones
}

mapping "azurerm_network_interface" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/networkInterface.json"

  dns_servers                   = NetworkInterfaceDnsSettings.dnsServers
  enable_ip_forwarding          = NetworkInterfacePropertiesFormat.enableIPForwarding
  enable_accelerated_networking = NetworkInterfacePropertiesFormat.enableAcceleratedNetworking
  internal_dns_name_label       = NetworkInterfaceDnsSettings.internalDnsNameLabel
}

mapping "azurerm_network_packet_capture" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/networkWatcher.json"

  target_resource_id        = PacketCaptureParameters.target
  maximum_bytes_per_packet  = PacketCaptureParameters.bytesToCapturePerPacket
  maximum_bytes_per_session = PacketCaptureParameters.totalBytesPerSession
  maximum_capture_duration  = PacketCaptureParameters.timeLimitInSeconds
}

mapping "azurerm_network_security_rule" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/networkSecurityGroup.json"

  description                                = SecurityRulePropertiesFormat.description
  protocol                                   = SecurityRulePropertiesFormat.protocol
  source_port_range                          = SecurityRulePropertiesFormat.sourcePortRange
  source_port_ranges                         = SecurityRulePropertiesFormat.sourcePortRanges
  destination_port_range                     = SecurityRulePropertiesFormat.destinationPortRange
  destination_port_ranges                    = SecurityRulePropertiesFormat.destinationPortRanges
  source_address_prefix                      = SecurityRulePropertiesFormat.sourceAddressPrefix
  source_address_prefixes                    = SecurityRulePropertiesFormat.sourceAddressPrefixes
  source_application_security_group_ids      = SecurityRulePropertiesFormat.sourceApplicationSecurityGroups
  destination_address_prefix                 = SecurityRulePropertiesFormat.destinationAddressPrefix
  destination_address_prefixes               = SecurityRulePropertiesFormat.destinationAddressPrefixes
  destination_application_security_group_ids = SecurityRulePropertiesFormat.destinationApplicationSecurityGroups
  access                                     = SecurityRuleAccess
  priority                                   = SecurityRulePropertiesFormat.priority
  direction                                  = SecurityRuleDirection
}

mapping "azurerm_network_watcher_flow_log" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/networkWatcher.json"

  network_security_group_id = FlowLogPropertiesFormat.targetResourceId
  storage_account_id        = FlowLogPropertiesFormat.storageId
  enabled                   = FlowLogPropertiesFormat.enabled
  version                   = FlowLogFormatParameters.version
}

mapping "azurerm_packet_capture" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/networkWatcher.json"

  target_resource_id        = PacketCaptureParameters.target
  maximum_bytes_per_packet  = PacketCaptureParameters.bytesToCapturePerPacket
  maximum_bytes_per_session = PacketCaptureParameters.totalBytesPerSession
  maximum_capture_duration  = PacketCaptureParameters.timeLimitInSeconds
}

mapping "azurerm_point_to_site_vpn_gateway" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/virtualWan.json"

  scale_unit = P2SVpnGatewayProperties.vpnGatewayScaleUnit
}

mapping "azurerm_private_link_service" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/privateLinkService.json"

  load_balancer_frontend_ip_configuration_ids = PrivateLinkServiceProperties.loadBalancerFrontendIpConfigurations
  enable_proxy_protocol                       = PrivateLinkServiceProperties.enableProxyProtocol
}

mapping "azurerm_public_ip" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/publicIpAddress.json"

  sku                     = PublicIPAddressSku.name
  idle_timeout_in_minutes = PublicIPAddressPropertiesFormat.idleTimeoutInMinutes
  domain_name_label       = PublicIPAddressDnsSettings.domainNameLabel
  reverse_fqdn            = PublicIPAddressDnsSettings.reverseFqdn
  zones                   = PublicIPAddress.zones
}

mapping "azurerm_public_ip_prefix" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/publicIpPrefix.json"

  sku           = PublicIPPrefixSku.name
  prefix_length = PublicIPPrefixPropertiesFormat.prefixLength
  zones         = PublicIPPrefix.zones
}

mapping "azurerm_route" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/routeTable.json"

  address_prefix         = RoutePropertiesFormat.addressPrefix
  next_hop_type          = RouteNextHopType
  next_hop_in_ip_address = RoutePropertiesFormat.nextHopIpAddress
}

mapping "azurerm_route_table" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/routeTable.json"

  disable_bgp_route_propagation = RouteTablePropertiesFormat.disableBgpRoutePropagation
}

mapping "azurerm_subnet" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/virtualNetwork.json"

  name              = Subnet.name
  address_prefix    = SubnetPropertiesFormat.addressPrefix
  address_prefixes  = SubnetPropertiesFormat.addressPrefixes
  service_endpoints = SubnetPropertiesFormat.serviceEndpoints
}

mapping "azurerm_traffic_manager_endpoint" {
  import_path = "azure-rest-api-specs/specification/trafficmanager/resource-manager/Microsoft.Network/stable/2018-04-01/trafficmanager.json"

  endpoint_status     = EndpointProperties.endpointStatus
  target              = EndpointProperties.target
  target_resource_id  = EndpointProperties.targetResourceId
  weight              = EndpointProperties.weight
  priority            = EndpointProperties.priority
  endpoint_location   = EndpointProperties.endpointLocation
  min_child_endpoints = EndpointProperties.minChildEndpoints
  geo_mappings        = EndpointProperties.geoMapping
}

mapping "azurerm_traffic_manager_profile" {
  import_path = "azure-rest-api-specs/specification/trafficmanager/resource-manager/Microsoft.Network/stable/2018-04-01/trafficmanager.json"

  profile_status         = ProfileProperties.profileStatus
  traffic_routing_method = ProfileProperties.trafficRoutingMethod
}

mapping "azurerm_virtual_hub" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/virtualWan.json"

  address_prefix = VirtualHubProperties.addressPrefix
}

mapping "azurerm_virtual_hub_connection" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/virtualWan.json"

  hub_to_vitual_network_traffic_allowed          = HubVirtualNetworkConnectionProperties.allowHubToRemoteVnetTransit
  vitual_network_to_hub_gateways_traffic_allowed = HubVirtualNetworkConnectionProperties.allowRemoteVnetToUseHubVnetGateways
  internet_security_enabled                      = HubVirtualNetworkConnectionProperties.enableInternetSecurity
}

mapping "azurerm_virtual_network" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/virtualNetwork.json"

  dns_servers = DhcpOptions.dnsServers
}

mapping "azurerm_virtual_network_gateway" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/virtualNetworkGateway.json"

  type          = VirtualNetworkGatewayPropertiesFormat.gatewayType
  vpn_type      = VirtualNetworkGatewayPropertiesFormat.vpnType
  enable_bgp    = VirtualNetworkGatewayPropertiesFormat.enableBgp
  active_active = VirtualNetworkGatewayPropertiesFormat.activeActive
  sku           = VirtualNetworkGatewaySku.name
  generation    = VirtualNetworkGatewayPropertiesFormat.vpnGatewayGeneration
}

mapping "azurerm_virtual_network_gateway_connection" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/virtualNetworkGateway.json"

  type                               = VirtualNetworkGatewayConnectionType
  authorization_key                  = VirtualNetworkGatewayConnectionPropertiesFormat.authorizationKey
  routing_weight                     = VirtualNetworkGatewayConnectionPropertiesFormat.routingWeight
  shared_key                         = VirtualNetworkGatewayConnectionPropertiesFormat.sharedKey
  connection_protocol                = ConnectionProtocol
  enable_bgp                         = VirtualNetworkGatewayConnectionPropertiesFormat.enableBgp
  express_route_gateway_bypass       = VirtualNetworkGatewayConnectionPropertiesFormat.expressRouteGatewayBypass
  use_policy_based_traffic_selectors = VirtualNetworkGatewayConnectionPropertiesFormat.usePolicyBasedTrafficSelectors
}

mapping "azurerm_virtual_network_peering" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/virtualNetwork.json"

  allow_virtual_network_access = VirtualNetworkPeeringPropertiesFormat.allowVirtualNetworkAccess
  allow_forwarded_traffic      = VirtualNetworkPeeringPropertiesFormat.allowForwardedTraffic
  allow_gateway_transit        = VirtualNetworkPeeringPropertiesFormat.allowGatewayTransit
  use_remote_gateways          = VirtualNetworkPeeringPropertiesFormat.useRemoteGateways
}

mapping "azurerm_virtual_wan" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/virtualWan.json"

  disable_vpn_encryption            = VirtualWanProperties.disableVpnEncryption
  allow_branch_to_branch_traffic    = VirtualWanProperties.allowBranchToBranchTraffic
  allow_vnet_to_vnet_traffic        = VirtualWanProperties.allowVnetToVnetTraffic
  office365_local_breakout_category = OfficeTrafficCategory
  type                              = VirtualWanProperties.type
}

mapping "azurerm_vpn_gateway" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/virtualWan.json"

  scale_unit = VpnGatewayProperties.vpnGatewayScaleUnit
}

mapping "azurerm_vpn_server_configuration" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/virtualWan.json"

  vpn_authentication_types = VpnServerConfigurationProperties.vpnAuthenticationTypes
  vpn_protocols            = VpnServerConfigurationProperties.vpnProtocols
}
