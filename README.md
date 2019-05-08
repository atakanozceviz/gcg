# GitHub Change Log Generator

This application is a tool to generate changelog from a Github milestone or tag(s)

## Configuration

Edit the config.yaml file to customize the generated output.

## Usage

```
gcg generate [flags]
```

### Options

```
Flags:
  -h, --help               help for generate
  -m, --milestone string   milestone title to get issues and pull requests for
  -r, --repo string        repository name to generate the Changelog for, in the form user/repo
      --since-tag string   issues and pull requests since tag
  -s, --state string       state of the issues and pull requests to get (open,closed or all)
  -t, --token string       personal access token
      --until-tag string   issues and pull requests until tag

Global Flags:
      --config string   config file (default is config.yaml)
```

[Learn more about templating](Templating.md)