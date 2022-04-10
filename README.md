# 健康打卡_河海大学版_FC

[![build](https://github.com/yin1999/healthreport_fc/actions/workflows/Build.yml/badge.svg)](https://github.com/yin1999/healthreport_fc/actions/workflows/Build.yml)

本项目用于河海大学健康打卡，此项目为基于[健康打卡_河海大学版](https://github.com/yin1999/healthreport)的函数计算版本。  

目前支持 [阿里云函数计算](https://www.aliyun.com/product/fc) 以及 [腾讯云云函数](https://cloud.tencent.com/product/sls) 平台。

## 开发指南

若想要直接使用（不对程序功能进行自定义），请转至[使用说明](#使用说明)

1. 环境配置，以`CentOS`/`Debian`为例

    - 安装 [Docker](https://docs.docker.com/get-docker/)

    - 安装 git

       ```bash
       # Centos
       sudo yum install -y git

       # Debian/Ubuntu
       sudo apt install -y git
       ```

2. 通过源码下载、编译

```bash
# 下载源码
git clone https://github.com/yin1999/healthreport_fc.git
cd healthreport_fc

# 编译
docker build .
```

## 使用说明

请查看 [**Wiki**](https://github.com/yin1999/healthreport_fc/wiki)。
