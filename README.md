# Sewer

A simple tool pipes STDOUT to some web services (eg. GitHub, Slack).

## Usage

To posting your command STDOUT as a comment on your GitHub Pull Request:

```sh
#You must set env vars before executing the command
export GITHUB_TOKEN=blah-blah-blah
export CIRCLE_PR_NUMBER=123
export CIRCLE_PROJECT_USERNAME=ainoya
export CIRCLE_PROJECT_REPONAME=sewer

echo 'hello, world' | sewer --drain=github
```

To posting your command STDOUT on slack incoming webhook:

```sh
export SLACK_WEBHOOK_URL=https://hooks.slack.com/services/xxxxxxxxxxxx/xxxxxxxxx
export SLACK_ICON_EMOJI=:tada:
export SLACK_CHANNEL=#general

echo 'hello, world' | sewer --drain=slack
```

You can also post on multiple web services:

```sh
echo 'hello, world' | sewer --drain=slack --drain=github
```

If you want decorating comment, you can use `--template` option:


```sh
$ cat test.tmpl
comment
--------

{{ .Message }}

#Then your pipe input is expanded as {{ .Message }} variable,
$ echo "hello, world" | sewer --drain=slack --template="$(cat test.tmpl)"
comment
---------

hello, world
```
