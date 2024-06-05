## Projects
A portable CLI tool for initializing programming projects from pre-defined templates. This tools is written in GoLang and doesn't have any 3rd-party dependencies.


### Installation

```bash
$ go install github.com/moeenn/projects@latest
```

### Usage

```
Usage of projects:
  -list
    	Print list of available template names
  -name string
    	Name of project being initialized (default "sandbox")
  -template string
    	Project template to use (default "cpp-cmake")
```

### Available templates

```bash
$ projects -list

Valid templates include: 
 - c
 - cpp-cmake
 - cpp-make
 - javascript (or 'js')
 - typescript (or 'ts')
 - java-gradle
 - python
```

### Build from source

```bash
$ go build .
```

### Todo

- [ ] Allow initializing git repositories
- [ ] Update Python template