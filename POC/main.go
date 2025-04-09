package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

/* Global Variables */
/* Inputs */
var location string
var project_name string

/* Other variables */
var hclFile *hclwrite.File
var rootBody *hclwrite.Body
var resource string
var vars map[string]string
var projects_folder string = "projects"

// Usage
func main() {
	// Constants (would be as input variables)
	location = "us-central1"
	project_name = "your-project"

	// Make a directory for your project (in `projects/`)
	makeProjectDirectory(projects_folder, project_name)

	// Create provider.tf
	hclFile = hclwrite.NewEmptyFile()
	rootBody = hclFile.Body()
	addProviderBlock(rootBody, project_name, location)
	writeToFile(hclFile, fmt.Sprintf("%s/%s/provider.tf", projects_folder, project_name))

	// Create infrastructure.tf
	//   Add terraform state file home in a new bucket in the central project with the requested project name in the bucket name
	hclFile = hclwrite.NewEmptyFile()
	rootBody = hclFile.Body()
	configureTerraformStateBucket(rootBody, project_name)
	writeToFile(hclFile, fmt.Sprintf("%s/%s/infrastructure.tf", projects_folder, project_name))

	// Create files for the resources
	hclFile = hclwrite.NewEmptyFile()
	rootBody = hclFile.Body()

	// Add multiple modules
	// GCP Bucket example
	resource = "gcp_bucket"
	vars = map[string]string{
		"project":  project_name,
		"location": location,
		"name":     "my-bucket", // as an input variable
	}
	addModule(rootBody, resource, fmt.Sprintf("./../../modules/%s", "resource"), vars)

	// GCP Function example
	resource = "gcp_function"
	vars = map[string]string{
		"project":     project_name,
		"location":    location,
		"name":        "my-function",                // as an input variable
		"runtime":     "nodejs20",                   // as an input variable
		"description": "my function can do things!", // as an input variable
	}
	addModule(rootBody, resource, fmt.Sprintf("./../../modules/%s", resource), vars)
	writeToFile(hclFile, fmt.Sprintf("%s/%s/%s.tf", projects_folder, project_name, project_name))
}

func makeProjectDirectory(projects_folder string, project_name string) {
	err := os.MkdirAll(fmt.Sprintf("%s/%s", projects_folder, project_name), 0755)
	if err != nil {
		fmt.Println("Error creating directories:", err)
		return
	}
}
func addProviderBlock(rootBody *hclwrite.Body, project_name string, region string) {
	// Add provider block
	provider := rootBody.AppendNewBlock("provider", []string{"google"})
	providerBody := provider.Body()
	providerBody.SetAttributeValue("project", cty.StringVal(project_name))
	providerBody.SetAttributeValue("region", cty.StringVal(region))
}

func configureTerraformStateBucket(rootBody *hclwrite.Body, project_name string) {
	// Add terraform state file home in a new bucket in the central project with the requested project name in the bucket name
	resource := "gcp_bucket"
	resource_name := "gcp_bucket_state"
	vars := map[string]string{
		"project": "central-infrastructure",
		"name":    fmt.Sprintf("terraform-%s", project_name),
	}
	addModule(rootBody, resource_name, fmt.Sprintf("./../../modules/%s", resource), vars)
}

func addModule(rootBody *hclwrite.Body, name string, source string, vars map[string]string) {
	module := rootBody.AppendNewBlock("module", []string{name})
	moduleBody := module.Body()

	moduleBody.SetAttributeValue("source", cty.StringVal(source))

	for key, value := range vars {
		moduleBody.SetAttributeValue(key, cty.StringVal(value))
	}
}

func writeToFile(hclFile *hclwrite.File, file_name string) {
	// Write to file
	if err := os.WriteFile(file_name, hclFile.Bytes(), 0644); err != nil {
		fmt.Printf("Error writing file: %s\n", err)
		return
	}
}
