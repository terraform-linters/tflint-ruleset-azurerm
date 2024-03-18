mapping "azurerm_network_packet_capture" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2023-09-01/networkWatcher.json"

  target_resource_id        = PacketCaptureParameters.target
  maximum_bytes_per_packet  = any //PacketCaptureParameters.bytesToCapturePerPacket
  maximum_bytes_per_session = any //PacketCaptureParameters.totalBytesPerSession
  maximum_capture_duration  = PacketCaptureParameters.timeLimitInSeconds
}
