mapping "azurerm_shared_image_gallery" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-03-03/gallery.json"

  resource_group_name = ResourceGroupNameParameter
  description         = GalleryImageProperties.description
}
