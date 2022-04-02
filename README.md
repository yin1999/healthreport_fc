# 健康打卡_河海大学版_FC

[![build](https://github.com/yin1999/healthreport_fc/actions/workflows/Build.yml/badge.svg)](https://github.com/yin1999/healthreport_fc/actions/workflows/Build.yml)

本项目用于河海大学健康打卡，此项目为基于[健康打卡_河海大学版](https://github.com/yin1999/healthreport)的函数计算版本。  

目前支持 [阿里云函数计算](https://www.aliyun.com/product/fc) 以及 [腾讯云云函数](https://cloud.tencent.com/product/sls) 平台。

## 编译教程

若想要直接使用（不对程序功能进行自定义），请转至[使用说明](#使用说明)

1. 环境配置，以`CentOS`/`Debian`为例

    - 安装Golang[>= 1.16]: [golang.google.cn/doc/install](https://golang.google.cn/doc/install)

    - 安装 git

       ```bash
       # Centos
       sudo yum install -y git

       # Debian/Ubuntu
       sudo apt install -y git
       ```

2. 通过源码下载、编译

    ```bash
    # 配置Goproxy
    go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct

    # 下载源码
    git clone https://github.com/yin1999/healthreport_fc.git
    cd healthreport_fc
    
    # 编译
    go run _script/build.go # 生成二进制文件 xxx-serverless.zip
    #  go run _script/build.go aliyun  # 可指定多个编译编译目标（可选值：aliyun、tencent）
    ```

## 使用说明

### 下载发行版

若已自行编译二进制文件，请跳过此步骤。

直接下载[release](https://github.com/yin1999/healthreport_fc/releases/latest)中的 **xx-serverless.zip** 文件。其中：

- **aliyun** 对应阿里云函数计算
- **tencent** 对应腾讯云 Serverless 云函数

### 函数计算配置

请查看 [**Wiki**](https://github.com/yin1999/healthreport_fc/wiki)。
