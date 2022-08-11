mapping "azurerm_shared_image" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2019-12-01/gallery.json"

  resource_group_name   = ResourceGroupNameParameter
  os_type               = GalleryImageProperties.osType
  description           = GalleryImageProperties.description
  eula                  = GalleryImageProperties.eula
  privacy_statement_uri = GalleryImageProperties.privacyStatementUri
  release_note_uri      = GalleryImageProperties.releaseNoteUri
}
