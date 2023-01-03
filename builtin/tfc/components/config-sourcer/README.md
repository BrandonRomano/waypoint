<!-- This file was generated via `make gen/integrations-hcl` -->
## terraform-cloud (configsourcer)

Read Terraform state outputs from Terraform Cloud.

### Interface

### Examples

```hcl
config {
  env = {
    "DATABASE_USERNAME" = dynamic("terraform-cloud", {
			organization = "foocorp"
			workspace = "databases"
			output = "db_username"
    })

    "DATABASE_PASSWORD" = dynamic("vault", {
			organization = "foocorp"
			workspace = "databases"
			output = "db_password"
    })

    "DATABASE_HOST" = dynamic("vault", {
			organization = "foocorp"
			workspace = "databases"
			output = "db_hostname"
    })
  }
}
```

