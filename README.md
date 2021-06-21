# licensepush

[![Built with WeBuild](https://raw.githubusercontent.com/webuild-community/badge/master/svg/WeBuild.svg)](https://webuild.community)

It's the tool for pushing license information or any text into your source code

## Installation

### Go get

In case we have installed Golang, we can use Go tool to install licensepush:

```
  go install github.com/ledongthuc/licensepush
```

## Usage

### Prerequisite

We need to setup configuration that contains license text. From the directory of project, create file toml with content:

```
license = """License content."""
```
### Run

Run the command:

```
  licensepush ~/project/source_code/ --config=~/project/source_code/.licensepush.yml
```

It push license content in `~/project/source_code/.licensepush.yml` to all code files in `~/project/source_code/`

Currently, we support code files:

*.go*

```
/*
 * License content.
*/
```

*.js*

```
/*
 * License content.
*/
```

*.css*

```
/*
 * License content.
*/
```

*.html*

```
<!--
 License content.
-->
```
