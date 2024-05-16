<h1 align="center">
  <img alt="logo" src="img/logo.png" width="450px"/><br/>
</h1>
<p align="center">A collection of utility functions that works well when using <a href="https://www.dagger.io/">Dagger</a> and  <a href="https://golang.org/">Go</a><br/><br/>


---

## Why 🤔

This library is a set of reusable functions that can be used when developing [Dagger](https://www.dagger.io/) modules or functions using [Go](https://golang.org/).


## Installation 🛠️

Install it using [Go get](https://golang.org/cmd/go/#hdr-Add_dependencies_to_current_module_and_install_them):

```bash
go get github.com/Excoriate/daggerx
```

### Pre-requisites 📋

- [Go](https://golang.org/doc/install) >= 1.22.3

>**NOTE**: For the tools used in this project, please check the [Makefile](./Makefile), and the [Taskfile](./Taskfile.yml) files. You'll also need [pre-commit](https://pre-commit.com/) installed.

---


## Usage 🚀

```go

```

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
