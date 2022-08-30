package utils

import (
	"encoding/json"
	"io/ioutil"
)

// "provider_schemas" describes the provider schemas for all providers throughout the configuration tree.
type Schema struct {
	ProviderSchemas ProviderSchemas `json:"provider_schemas"`
}

// "provider_schemas" describes the provider schemas for all providers throughout the configuration tree.
type ProviderSchemas struct {
	Azurerm ProviderSchema `json:"registry.terraform.io/hashicorp/azurerm"`
}

// "resource_schemas" describes the schemas for all resources.
type ProviderSchema struct {
	ResourceSchemas map[string]ResourceSchema `json:"resource_schemas"`
}

// "resource_schemas" describes the schemas for all resources.
type ResourceSchema struct {
	Block BlockSchema `json:"block"`
}

// A block representation contains "attributes" and "block_types" (which represent nested blocks).
type BlockSchema struct {
	Attributes map[string]AttributeSchema `json:"attributes"`
	BlockTypes map[string]ResourceSchema  `json:"block_types"`
}

// AttributeSchema
type AttributeSchema struct {
	Type      interface{} `json:"type"`
	Sensitive bool        `json:"sensitive"`
}

func LoadProviderSchema(path string) ProviderSchema {
	src, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var schema Schema
	if err := json.Unmarshal(src, &schema); err != nil {
		panic(err)
	}
	return schema.ProviderSchemas.Azurerm
}
