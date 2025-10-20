package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"github.com/giantswarm/apiextensions/v6/pkg/apis/provider/v1alpha1"
	"github.com/santhosh-tekuri/jsonschema/v6"
)

func main() {
	_ = v1alpha1.AzureConfig{}

	schema, err := jsonschema.UnmarshalJSON(strings.NewReader(`{"type": "string", "contentEncoding": "hex"}`))
	if err != nil {
		log.Fatal(err)
	}
	inst := "abcxyz"

	c := jsonschema.NewCompiler()
	c.RegisterContentEncoding(&jsonschema.Decoder{
		Name:   "hex",
		Decode: hex.DecodeString,
	})
	c.AssertContent()
	if err := c.AddResource("schema.json", schema); err != nil {
		log.Fatal(err)
	}
	sch, err := c.Compile("schema.json")
	if err != nil {
		log.Fatal(err)
	}
	err = sch.Validate(inst)
	fmt.Println("valid:", err == nil)
}
