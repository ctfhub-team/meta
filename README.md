# Meta 题目信息

[![Build Status](https://github.com/ctfhub-team/meta/actions/workflows/test.yml/badge.svg)](https://github.com/ctfhub-team/meta/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/ctfhub-team/meta)](https://goreportcard.com/report/github.com/ctfhub-team/meta)
[![GoDoc](https://pkg.go.dev/badge/github.com/ctfhub-team/meta?status.svg)](https://pkg.go.dev/github.com/ctfhub-team/meta?tab=doc)

> [题目环境生成器 - https://github.com/ctfhub-team/challenge_generate](https://github.com/ctfhub-team/challenge_generate)

> [环境制作说明 - https://www.wolai.com/ctfhub/3DvnJJtPbHyyDtVkDaW1yz](https://www.wolai.com/ctfhub/3DvnJJtPbHyyDtVkDaW1yz)

## 说明

```yaml
author:
  # 制作者ID
  name: l1n3
  # 制作者邮箱
  contact: yw9381@163.com
task:
  # 题目镜像名称，名称中应当不包含下划线(_)，空格( )和中文，下划线和空格用-代替
  id: 2022_hitcon_web_rce-me
  # 题目显示名称
  name: rce me
  # 题目类型， 包含 con, file, ext
  # con类型为容器类型，指这个题目存在容器, file类型为文件类型，指这个题目只有附件, ext类型为外部类型，指这个题目是一个外部链接
  # con及ext类型附件可有可无，file类型必须有附件
  type: con
  # 题目分类，首字母需要大写，每个题目应当只有一个分类，分类参考如下
  # 例如 Web, Pwn, Reverse, Misc, Crypto, Forensics, Blockchain, Mobile, ICS, IoT
  category: Web
  # 题目描述
  description: |
    please rce me
    try try try!
  # 题目难度，难度分为 签到, 简单, 中等, 困难
  level: 签到
  # 题目flag，如是静态flag在此处填写具体的flag值，如是动态flag则此处留空
  flag:
  # 题目来源，来源格式: 年份-比赛名称简写-题目类型-题目显示名称
  # 例如 2021年强网杯的Web类的babysqli，则为2021-QWB-Web-baysqli
  # 例如 2019年SCTF的Pwn类的babyheap，则为2019-SCTF-Pwn-bayheap
  refer: 2022-HitCon-Web-rce
  # 题目标签，标签应当尽可能体现题目考点，如无标签则此处为空数组
  tags:
    - Web
    - HitCon
    - 2022
  # 题目提示，如无提示则此处为空数组
  hints:
    - asdasdaz
    - asdas
    - asdad
# 关于容器的配置，如果没有容器则无须此部分
container:
  # 镜像名称，需要和task_name保持一致，如果有多个容器，请在末尾添加标识符，例如web,db或是1,2
  # 例如 challenge_web_sqli_basic-web
  - image: xxxxxxxx
    # 需要对外暴露的端口, 协议为tcp, udp, http, 因K8S在同一个pod中的多容器不能监听同一个端口，故不同的容器暴露的端口需不相同，可视作所有容器共享一个IP地址及端口资源
    port:
      - 10000/tcp
      - 20000/udp
    # 资源配置，参考docker对资源的定义
    resource:
      cpu: 250m
      mem: 256Mi
  # 例如 challenge_web_sqli_basic-db
  - image: xxxxxx
    # 需要对外暴露的端口
    port:
      - 10000/tcp
      - 20000/udp
    # 资源配置，参考docker对资源的定义，cpu 250m即为0.25C
    resource:
      cpu: 250m
      mem: 256Mi
```

## Multiple Yaml 多文件合并

```yaml
---
yaml content 1
---
yaml content 2
```
