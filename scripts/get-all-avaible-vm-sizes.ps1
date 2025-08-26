# You have to install the Az Modules with the following commands:
# `Install-Module Az`
# `Import-Module Az`
# Also make sure, that you already did an authentication!

$locations = Get-AzLocation
$vmSizes = [System.Collections.Generic.HashSet[string]]@()
foreach ($location in $locations.Location) {
    $vmSizeNames = (Get-AzVMSize -Location $location).Name
    foreach ($name in $vmSizeNames) {
        $vmSizes.Add($name)
    }
}

Write-Output $vmSizes
