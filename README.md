# Sewer

A simple tool pipes STDOUT to some web services (eg. GitHub, Slack).

[![wercker status](https://app.wercker.com/status/3a74c0ee858eaeffbe9aba92fa868042/m/master "wercker status")](https://app.wercker.com/project/byKey/3a74c0ee858eaeffbe9aba92fa868042)

## Usage

To posting your command STDOUT as a comment on your GitHub Pull Request:

```sh
#You must set env vars before executing the command
#export GITHUB_TOKEN=blah-blah-blah
#export CI_PULL_REQUEST=123
#export CIRCLE_PROJECT_USERNAME=ainoya
#export CIRCLE_PROJECT_REPONAME=sewer

echo 'hello, world' | sewer --github
```