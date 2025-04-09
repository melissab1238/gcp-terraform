# gcp-terraform

Messing around with GCP and terraform

See `POC/` folder for setup instructions.

# STAR

|               | Description                                                                                                                                                                                                                                                                                        |
| ------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **S**ituation | A developer wants to quickly test out an idea in GCP. I am a cloud engineer on the infrastructure team.                                                                                                                                                                                            |
| **T**ask      | Design and build a full-service application to automate the provisioning of the resources requested by the developer. Consider network security, automation, local development, and easy destruction of resources upon completion of the developer's idea to save.                                 |
| **A**ction    | Create a frontend app that capture's developers requests, connects with central state storage management, configures network security, creates a new GCP project for the developer's "playground", tracks cost, and allows for easy deletion of the projects once the developer is done "playing". |
| **R**esult    | coming soon!                                                                                                                                                                                                                                                                                       |

## Best practices

- state management
  - versioning on the state bucket
  - use state locking
- access control
  - least privilege
- monitoring
  - monitor resource creation across projects

## Options

- Databases
  - SQL
  - NoSQL Firestore
- Storage
  - Buckets
- Serverless
  - Cloud Run
  - Cloud Functions
- Containerized
  - Kubernetes Engine (GKE)
  -
