# 健康打卡_河海大学版_FC

[![build](https://github.com/yin1999/healthreport_fc/actions/workflows/Build.yml/badge.svg)](https://github.com/yin1999/healthreport_fc/actions/workflows/Build.yml)

本项目用于河海大学健康打卡，此项目为基于[健康打卡_河海大学版](https://github.com/yin1999/healthreport)的函数计算版本。  

当前仅适用于[阿里云函数计算](https://fc.console.aliyun.com/)，若使用腾讯云的相关产品，请自行修改代码

## 编译教程

**请在Linux系统下编译**，Windows操作系统下可正常执行交叉编译，但在无其他依赖的情况下无法授予可执行权限（无法在函数计算中正常运行）。  
若想要直接使用，请转至[使用说明](#使用说明)

1. 环境配置，以`CentOS`/`Debian`为例

    - 安装Golang[>= 1.16]: [golang.google.cn/doc/install](https://golang.google.cn/doc/install)

    - 安装git、make、zip

       ```bash
       # Centos
       sudo yum install -y git make zip

       # Debian/Ubuntu
       sudo apt install -y git make zip
       ```

2. 通过源码下载、编译

    ```bash
    # 配置Goproxy
    go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct

    # 下载源码
    git clone https://github.com/yin1999/healthreport_fc.git
    cd healthreport_fc
    
    # 编译
    make # 生成二进制文件 xxx-serverless.zip
    # make targets=aliyun  # 可使用targets指定编译目标（可选值：aliyun、tencent）
    ```

## 使用说明

### 下载发行版

若已自行编译二进制文件，请跳过此步骤。

直接下载[release](https://github.com/yin1999/healthreport_fc/releases/latest)中的 **xx-serverless.zip** 文件。其中：

- **aliyun** 对应阿里云函数计算
- **tencent** 对应腾讯云 Serverless 云函数

### 函数计算配置

请查看 [**Wiki**](https://github.com/yin1999/healthreport_fc/wiki)。
