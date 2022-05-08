# MR-Tracker
[![Maintainability](https://api.codeclimate.com/v1/badges/f28b966f3baf7ab66a9d/maintainability)](https://codeclimate.com/github/MamaShip/MR-Tracker/maintainability)
[![Total alerts](https://img.shields.io/lgtm/alerts/g/MamaShip/MR-Tracker.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/MamaShip/MR-Tracker/alerts/)

Find changes between given versions by analyzing merge requests.

Currently only work for Gitlab MRs.

## Usage

### Before

This tool works on Gitlab APIs. Make sure your project visibility is **public**, or generate a `private-token` (see [Gitlab docs](https://docs.gitlab.com/ee/user/project/settings/project_access_tokens.html)) for your project.

### Install

Download the execution file from [release](https://github.com/MamaShip/MR-Tracker/releases) page. And put it in yout `PATH`.

Or if you have golang environment, install by command:

```
go install github.com/MamaShip/MR-Tracker@latest
```

### Run

use `MR-Tracker -h` to see help document.

available options:

| options    | input type | description                                                |
| ---------- | ---------- | ---------------------------------------------------------- |
| `-h`       |            | Print help info                                            |
| `-v`       |            | Print version                                              |
| `-site`    | string     | Set your Gitlab URL                                        |
| `-project` | int        | Set your project ID                                        |
| `-token`   | string     | Set your Gitlab API token for the project                 |
| `-branch`  | string     | This tool automatically analysis MRs on default branch. <br /> If you wanna track changes on other branches, set it by this option |
| `-start`   | string     | Set the tag where you want to compare difference from      |
| `-end`     | string     | Set the tag where you want to compare difference to        |
| `-post`    |            | If this flag is set, result will be posted as gitlab issue |
