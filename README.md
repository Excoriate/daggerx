<h1 align="center">
  <img alt="logo" src="img/logo.png" width="450px"/><br/>
</h1>
<p align="center">A collection of utility functions that works well when using <a href="https://www.dagger.io/">Dagger</a> and  <a href="https://golang.org/">Go</a><br/><br/>

---
[![Go Reference](https://pkg.go.dev/badge/github.com/Excoriate/daggerx.svg)](https://pkg.go.dev/github.com/Excoriate/daggerx)
[![Contributors](https://img.shields.io/github/contributors/badges/shields)](https://github.com/Excoriate/daggerx/graphs/contributors)
[![Activity](https://img.shields.io/github/commit-activity/m/Excoriate/daggerx)](https://github.com/Excoriate/daggerx/pulse)

## Why 🤔

This library is a set of reusable functions that can be used when developing [Dagger](https://www.dagger.io/) modules or functions using [Go](https://golang.org/). The objective is to speed up the development process and make it easier to write Dagger modules using Go.


## Installation 🛠️

Install it using [Go get](https://golang.org/cmd/go/#hdr-Add_dependencies_to_current_module_and_install_them):

```bash
go get github.com/Excoriate/daggerx
```

### Pre-requisites 📋

- [Go](https://golang.org/doc/install) >= ~1.22

>**NOTE**: For the tools used in this project, please check the [Makefile](./Makefile), and the [Taskfile](./Taskfile.yml) files. You'll also need [pre-commit](https://pre-commit.com/) installed.

---


## Usage 🚀

### Convert map to slice of DaggerEnvVars

```go
package main

import (
    "fmt"
    "github.com/Excoriate/daggerx/pkg/envvars"
    "github.com/Excoriate/daggerx/pkg/types"
)

func main() {
    envVarsMap := map[string]string{
        "key1": "value1",
        "key2": "value2",
        "key3": "value3",
    }

    envVarsSlice, err := envvars.ToDaggerEnvVarsFromMap(envVarsMap)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Environment Variables Slice:", envVarsSlice)
}

```

### Generate validated Dagger commands

```go
package main

import (
    "fmt"
    "github.com/Excoriate/daggerx/pkg/execmd"
)

func main() {
    cmd, err := execmd.GenerateCommand("terraform", "plan", "-var", "foo=bar")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Generated Command:", *cmd)
}
````

### Convert slice of key=value strings to slice of DaggerEnvVars

```go
package main

import (
    "fmt"
    "github.com/Excoriate/daggerx/pkg/envvars"
    "github.com/Excoriate/daggerx/pkg/types"
)

func main() {
    envVarsSlice := []string{"key1=value1", "key2=value2", "key3=value3"}

    daggerEnvVarsSlice, err := envvars.ToDaggerEnvVarsFromSlice(envVarsSlice)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Dagger Environment Variables Slice:", daggerEnvVarsSlice)
}

```

---


## Contributing

Please read our [contributing guide](./CONTRIBUTING.md).
