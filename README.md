# Sewer

A simple tool pipes STDOUT to some web services (eg. GitHub, Slack).

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