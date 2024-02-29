## Projects
A simple portable tool for initializing programming projects from templates. This tools is written in GoLang and doesn't have any 3rd-party dependencies.


### Installation

```bash
$ go install github.com/moeenn/projects@latest
```

### Usage

```
Usage of ./projects:
  -name string
        Name of project being initialized (default "sandbox")
  -template string
        Project template to use. Valid options are 'cpp', 'js' (default "cpp")
```

### Build from source

```bash
$ go build .
```
