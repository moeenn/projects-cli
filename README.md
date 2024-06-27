## Projects
A portable CLI tool for initializing programming projects from pre-defined templates. This tools is written in GoLang and doesn't have any 3rd-party dependencies.

#### Important

**Note**: This project has now been archived and have been superseded by [projects-rs](https://github.com/moeenn/projects-rs) which is a one-to-one port of this codebase to rust. This codebase will no longer receive any updates.


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
