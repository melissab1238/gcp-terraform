# Go Terraform Generator

This project acts as an API and creates a GCP project with resources with terraform. The state files for the project's configuration is stored in a central-infrastructure (limited access) GCP project.

## Setup

### Initialize a new Go module

`go mod init terraform-generator`

### Get required dependencies

- `go get github.com/hashicorp/hcl/v2/hclwrite`
- `go get github.com/zclconf/go-cty/cty`

### Run it

`go run main.go`

### Validate with terraform init

- `cd projects\your-project\`
- `terraform init`

## Architecture

Central infrastructure GCP project

- bucket for each project's state files
- network config
- IAM roles

Modules (for now)

- gcp_bucket
- gcp_function

This service

- a GCP Cloud Run Golang service
- takes in inputs
  - location
  - region
  - project_name
  - resources [] > see `gcp.proto` for example API configuration

# Future (If I had more time)
- make the project bucket in the central project be used as tf backend
- make the go app into an API
  - likely HTTP REST, but could use gRPC instead
  - use input variables in code, refactor
- automatically create a PR on trigger
- create GHA that triggers on PR creation, runs `terraform init` and `terraform plan`, also adds reviewers from the public cloud team to review
- put this service on GCP cloud run as a function
- add more resource types
- make a simple frontend that calls out to the API
- deal with IAM/auth; the developer requesting the project creation will automatically be granted access to the project, unique google ID in API request
- add more than just "create a project with these resources" endpoint. for instance, maybe the developer wants to add another bucket or change a bucket name and have that code change be in terraform.
- consider generating a `variables.tf` file and `*.tfvars` files in the project folders (`projects/xxx/`) instead of passing in variables

Notes

- developers would have the ability to make PRs to their projects

Considerations/Questions/Unknowns/Other thoughts

- would the terraform code for the developer's project be stored in one repo (like this one), where the `projects/` folder holds all the terraform code? or would they be separated? or held in each developer's repo where their project code is being held?
