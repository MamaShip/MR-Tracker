# MR-Tracker

[English](README.md) | 中文

[![Maintainability](https://api.codeclimate.com/v1/badges/f28b966f3baf7ab66a9d/maintainability)](https://codeclimate.com/github/MamaShip/MR-Tracker/maintainability)
[![Total alerts](https://img.shields.io/lgtm/alerts/g/MamaShip/MR-Tracker.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/MamaShip/MR-Tracker/alerts/)
[![Go Reference](https://pkg.go.dev/badge/github.com/MamaShip/MR-Tracker.svg)](https://pkg.go.dev/github.com/MamaShip/MR-Tracker)

基于 MR 标题生成2个版本之间的改动变化 changelog。

目前仅支持分析 Gitlab 仓库。

## 使用

### 前置准备

此工具依赖于 Gitlab API。请先确认你的 Gitlab 仓库是**公开**的；对于私有仓库，你需要先为仓库生成相应访问权限的 `private-token` (详见 [官方文档](https://docs.gitlab.com/ee/user/project/settings/project_access_tokens.html))。

另外，Gitlab API 需要使用到 **Project ID**。你可以在仓库的首页找到它：
![where to find project ID](images/project_id.png)

### 在 CI 中使用 MR-Tracker

推荐的使用方式：直接用 **docker 镜像([mr-tracker](https://hub.docker.com/r/mamaship/mr-tracker))**，镜像内已经打包了最新的 MR-Tracker，可集成进 CI Pipeline，工作流自动触发 changelog 生成。

使用 Docker 镜像 `mamaship/mr-tracker:latest`，可以直接在 CI job 里调用 `MR-Tracker` 命令。

`.gitlab-ci.yml` 配置样例: 

```
gen-changelog:
  image: mamaship/mr-tracker:latest
  rules:
    - if: $CI_COMMIT_TAG
  script:
    - MR-Tracker -site YOUR_GITLAB_DOMAIN -project YOUR_PROJECT_ID -token YOUR_TOKEN -latest $CI_COMMIT_TAG -post
```

### 命令行程序

#### 安装

如果你已有 Go 语言开发环境，可以执行以下命令安装：

```
go install github.com/MamaShip/MR-Tracker@latest
```

或者你可以从 [Release](https://github.com/MamaShip/MR-Tracker/releases) 页面下载对应平台的可执行文件，把它放到你的`PATH` 路径下（或者直接运行它）。

#### 运行

必须通过 `-project` 参数传递 `Project ID` 给程序，才能执行分析。

你可以执行 `MR-Tracker -h` 来查看详细的命令行参数文档。

**基础选项**:

| 参数名    | 输入类型 | 功能                                                |
| ---------- | ---------- | ---------------------------------------------------------- |
| `-h`       |            | 打印帮助信息                                            |
| `-v`       |            | 打印版本号                                              |
| `-project` | int        | 设置 Project ID                                        |
| `-site`    | string     | 设置 Gitlab 域名 (默认值: `gitlab.com`)                |
| `-start`   | string     | 设置分析的开始范围（tag 名）                              |
| `-end`     | string     | 设置分析的结束范围（tag 名）                                |
| `-output`  | string     | 将分析结果以 markdown 格式输出到文件 |
| `-simple`  |            | 简化输出内容 |

样例:

```
MR-Tracker -project 278964 -start v14.10.0-ee -end v14.10.1-ee -output changes.md
```

**额外选项**:

| 参数名    | 输入类型 | 功能                                                |
| ---------- | ---------- | ---------------------------------------------------------- |
| `-token`*   | string     | 设置你的 Gitlab API token。<br /> 私有仓库必须有 token 才可能执行分析。 |
| `-post`    |            | 将分析结果发布到该仓库的 Issue 页面。 |
| `-latest`    | string | 给定一个版本 tag，自动尝试找出仓库内前一个[语义化版本号](https://semver.org/)，并对两个版本间的变化执行分析。 |
| `-branch`  | string     | MR-Tracker 默认基于仓库的 default 分支进行 MR 分析。 <br /> 如果想要关注其它分支的 MR 变动，通过此参数来指定分支名。 |

*你也可以使用环境变量来存储 token:
```
export MR_TRACKER_TOKEN=XXXXXXXX
```
MR-Tracker 会从环境变量内读取名为 `MR_TRACKER_TOKEN` 的值用作 token（如果找到的话）。

## 受启发于

- [https://github.com/github-changelog-generator/github-changelog-generator](https://github.com/github-changelog-generator/github-changelog-generator)
- [https://github.com/eirture/walle](https://github.com/eirture/walle)