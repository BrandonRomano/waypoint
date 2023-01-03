<!-- This file was generated via `make gen/integrations-hcl` -->
## docker (task)

Launch a Docker container as a task.

If a Docker server is available (either locally or via environment variables
such as "DOCKER_HOST"), then it will be used to start the container.

### Interface

### Examples

```hcl
task {
  use "docker" {
		force_pull = true
  }
}
```

