mapping "azurerm_packet_capture" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-07-01/networkWatcher.json"

  target_resource_id        = PacketCaptureParameters.target
  maximum_bytes_per_packet  = PacketCaptureParameters.bytesToCapturePerPacket
  maximum_bytes_per_session = PacketCaptureParameters.totalBytesPerSession
  maximum_capture_duration  = PacketCaptureParameters.timeLimitInSeconds
}
