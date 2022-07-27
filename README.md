# MR-Tracker

English | [中文](README.md)

[![Maintainability](https://api.codeclimate.com/v1/badges/f28b966f3baf7ab66a9d/maintainability)](https://codeclimate.com/github/MamaShip/MR-Tracker/maintainability)
[![Total alerts](https://img.shields.io/lgtm/alerts/g/MamaShip/MR-Tracker.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/MamaShip/MR-Tracker/alerts/)
[![Go Reference](https://pkg.go.dev/badge/github.com/MamaShip/MR-Tracker.svg)](https://pkg.go.dev/github.com/MamaShip/MR-Tracker)

Find changes between given versions by analyzing merge requests.

Only work for Gitlab MRs.

## Usage

### Before

This tool works on Gitlab APIs. Make sure your gitlab project visibility is **public**, OR generate a `private-token` (see [Gitlab docs](https://docs.gitlab.com/ee/user/project/settings/project_access_tokens.html)) for your project.

The **project ID** is required for Gitlab APIs to work. You can find it from your project home page:
![where to find project ID](images/project_id.png)

### Using MR-Tracker in CI Pipeline

It is recommended to use **docker image([mr-tracker](https://hub.docker.com/r/mamaship/mr-tracker))** in CI Pipeline to automatically generate changelog in your workflow.

Docker image `mamaship/mr-tracker:latest` makes the command `MR-Tracker` available for CI jobs.

`.gitlab-ci.yml` Example: 

```
gen-changelog:
  image: mamaship/mr-tracker:latest
  rules:
    - if: $CI_COMMIT_TAG
  script:
    - MR-Tracker -site YOUR_GITLAB_DOMAIN -project YOUR_PROJECT_ID -token YOUR_TOKEN -latest $CI_COMMIT_TAG -post
```

### CLI

#### Install

If you have golang environment, install by command:

```
go install github.com/MamaShip/MR-Tracker@latest
```

Or you can download the executable file from [release](https://github.com/MamaShip/MR-Tracker/releases) page. And put it in your `PATH` (or run it directly).

#### Run

The `project ID` must be set by `-project` flag for program to run.

Type `MR-Tracker -h` to see detail instructions.

**Basic options**:

| options    | input type | description                                                |
| ---------- | ---------- | ---------------------------------------------------------- |
| `-h`       |            | Print help info                                            |
| `-v`       |            | Print version                                              |
| `-project` | int        | Set your project ID                                        |
| `-site`    | string     | Set your Gitlab URL (default: `gitlab.com`)                |
| `-start`   | string     | Set the tag to start analyze                               |
| `-end`     | string     | Set the tag to stop analyze                                |
| `-output`  | string     | Set the output file to save the changes in markdown format |
| `-simple`  |            | If this flag is set, the markdown output will be simplified |

example:

```
MR-Tracker -project 278964 -start v14.10.0-ee -end v14.10.1-ee -output changes.md
```

**Additional options**:

| options    | input type | description                                                |
| ---------- | ---------- | ---------------------------------------------------------- |
| `-token`*   | string     | Set your Gitlab API token for the project.<br /> This is essential for non-public repos. |
| `-post`    |            | If this flag is set, result will be posted as gitlab issue. <br /> The Gitlab API token is needed for authentication. |
| `-latest`    | string | Given a version tag, changes of latest [Semantic Version](https://semver.org/) will be analyzed.(ignoring pre-release & build identifier) |
| `-branch`  | string     | MR-Tracker automatically analysis MRs on default branch. <br /> If you wanna track changes on other branches, set it by this option. |

*You can use environment variables to store the private token:
```
export MR_TRACKER_TOKEN=XXXXXXXX
```
MR-Tracker will read `MR_TRACKER_TOKEN` from `ENV` if available.

## Inspired by

- [https://github.com/github-changelog-generator/github-changelog-generator](https://github.com/github-changelog-generator/github-changelog-generator)
- [https://github.com/eirture/walle](https://github.com/eirture/walle)