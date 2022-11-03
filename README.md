# terraform-docs-wrapper

## **Requirements**

```
terraform-docs
```

## How to run

```
go run main.go
```

## Suported Parameters

```
--ignore-dirs
  list of ignored directory to traverse into

--base-path
  module base path, default to current directory 

--config
  your terraform-docs yaml config name
```

## Use Case

You want to use terraform-docs to autogenerate the README.md files for your terraform public/custom modules. Imagine your terraform modules are structured like this, and you want to add the README.md file in each module directory. 

```
.
├── modules
│   └── gcp
│       ├── compute
│       │   └── instance
│       │       ├── main.tf
│       │       ├── output.tf
│       │       ├── variable.tf
│       │       └── versions.tf
│       ├── network
│       │       ├── main.tf
│       │       ├── output.tf
│       │       ├── variable.tf
│       │       └── versions.tf
...
```

Unfortunately, current terraform-docs only support single-nested modules directory structure. In order to compatible with multi-nested modules directory structure, this *single-file-golang-ducktaped-script* was born.

Basically it will traverse through your terraform module directory to find your terraform module path, so you can generate it's README.md using terraform-docs custom module path.
