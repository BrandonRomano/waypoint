<!-- This file was generated via `make gen/integrations-hcl` -->
## aws-ssm (configsourcer)

Read configuration values from AWS SSM Parameter Store.

### Interface

### Examples

```hcl
config {
  env = {
    PORT = dynamic("aws-ssm", {
	  path = "port"
	})
  }
}
```

