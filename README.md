<h1 align="center">
  <img alt="logo" src="img/tftest-logo.png" width="300px"/><br/>
</h1>
<p align="center">A collection of functions and common patterns to test <a href="https://www.terraform.io/">Terraform</a> code using <a href="https://golang.org/">Go</a>, and <a href="https://terratest.gruntwork.io">Terratest</a>.<br/><br/>


---

## Why 🤔

This library wraps Terratest to provide simplified functions for common testing patterns in Terraform modules.
The **problem** that this library aims to solve is to make it so easy to write
([terraform](https://www.terraform.io/)) tests that you don't need to be a profficient Go developer to write them.

## Installation 🛠️

Install it using [Go get](https://golang.org/cmd/go/#hdr-Add_dependencies_to_current_module_and_install_them):

```bash
go get github.com/Excoriate/tftest
```

### Pre-requisites 📋

- [Go](https://golang.org/doc/install) >= 1.18

>**NOTE**: For the tools used in this project, please check the [Makefile](./Makefile), and the [Taskfile](./Taskfile.yml) files. You'll also need [pre-commit](https://pre-commit.com/) installed.

---


## Usage 🚀

### Simple 'Plan' scenario

```go
package simple

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/excoriate/tftest/pkg/scenario"
)

func TestSimpleOptionsPlanScenario(t *testing.T) {
    s, err := scenario.New(t, "../../data/tf-random")
    assert.NoErrorf(t, err, "Failed to create scenario: %s", err)

    s.Stg.PlanStage(t, s.GetTerraformOptions())
}

```

### Expecting changes on plan time for certain resources

```go

import (
"testing"
"github.com/stretchr/testify/assert"
"github.com/excoriate/tftest/pkg/scenario"
)


func TestSpecificResourcesExpectedChanges(t *testing.T) {
    s, err := scenario.New(t, "../../data/tf-random")
    assert.NoErrorf(t, err, "Failed to create scenario: %s", err)

    s.Stg.PlanWithSpecificResourcesThatWillChange(t, s.GetTerraformOptions(), []string{"random_id.this"})
}

```

### Full lifecycle (init, plan, destroy) scenario

```go
import (
"testing"
"github.com/stretchr/testify/assert"
"github.com/excoriate/tftest/pkg/scenario"
)

func TestLifecycle(t *testing.T) {
    s, err := scenario.New(t, "../../data/tf-random")
    assert.NoErrorf(t, err, "Failed to create scenario: %s", err)

    defer s.Stg.DestroyStage(t, s.GetTerraformOptions())

    s.Stg.PlanStageWithAnySortOfChanges(t, s.GetTerraformOptions())
    s.Stg.ApplyStage(t, s.GetTerraformOptions())
}
```

### Expecting a variable that should have an expected value on Plan time

```go
import (
"testing"
"github.com/stretchr/testify/assert"
"github.com/excoriate/tftest/pkg/scenario"
)

func TestWithVarOptionsValid(t *testing.T) {
	workdir := "../../data/tf-random"
	s, err := scenario.NewWithOptions(t, workdir,
		scenario.WithVarFiles(workdir, "fixtures/override-random-password.tfvars"))

	assert.NoErrorf(t, err, "Failed to create scenario: %s", err)

	s.Stg.PlanWithSpecificVariableValueToExpect(t, s.GetTerraformOptions(), "random_length_password", "25")
}

```

More examples will be added in the [examples](./test/examples) folder.

---

## APIs 📚

There are more API(s) available for the following common patterns:

- [x] A strict validation of the **terraform directory** (it validates whether it's an actual terraform module, directory integrity and others).
- [x] A strict validation of `terraform.tfvars` files.
- [x] Cloud Provider API(s). Currently, [AWS](https://aws.amazon.com/) is supported.


## Roadmap 🗓️

- [ ] Add more tests
- [ ] Add a set of pre-defined error messages

>**Note**: This is still work in progress, however, I'll be happy to receive any feedback or contribution. Ensure you've read the [contributing guide](./CONTRIBUTING.md) before doing so.


## Contributing

Please read our [contributing guide](./CONTRIBUTING.md).
