package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type schema struct {
	ProviderSchema providerSchema `json:"provider_schemas"`
}

type providerSchema struct {
	AzureRM provider `json:"registry.terraform.io/hashicorp/azurerm"`
}

type provider struct {
	ResourceSchemas map[string]resourceSchema `json:"resource_schemas"`
}

type resourceSchema struct {
	Block block `json:"block"`
}

type block struct {
	Attributes map[string]attribute      `json:"attributes"`
	BlockTypes map[string]resourceSchema `json:"block_types"`
}

type attribute struct {
	Type      interface{} `json:"type"`
	Sensitive bool        `json:"sensitive"`
}

func loadProviderSchema() provider {
	schemaPath := getFullPath("schema/schema.json")
	fmt.Println(schemaPath)
	src, err := ioutil.ReadFile(schemaPath)

	if err != nil {
		panic(err)
	}

	var schema schema
	if err := json.Unmarshal(src, &schema); err != nil {
		panic(err)
	}
	return schema.ProviderSchema.AzureRM
}
