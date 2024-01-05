[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](https://go.dev/)
[![Latest release](https://badgen.net/github/release/arnoldj-devops/ghvars)](https://github.com/arnoldj-devops/ghvars/releases)
[![GitHub license](https://img.shields.io/github/license/arnoldj-devops/ghvars.svg)](https://github.com/arnoldj-devops/ghvars/blob/master/LICENSE)

# ghvars
CLI tool to view, list, set, and delete github variables with ease

# Install ghvars

## MacOS

```bash
brew install arnoldj-devops/tools/ghvars
```

## Ubuntu

```bash
curl -s https://raw.githubusercontent.com/arnoldj-devops/ghvars/master/scripts/install.sh | bash
```

# Usage

**To fetch github variables** <br />

```bash
ghvars get
```

**To set github variables from file**

By default set command runs for all github environments. To run for specific environment, use `-e option` <br />
<br />
**To disconnect instance** <br />

```bash
ghvars set -e dev
```

<br />

**For all commands** <br />

```bash
ghvars --help
```

 <br />

# Prerequisites:

- [gh](https://cli.github.com/)
