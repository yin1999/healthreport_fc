# 健康打卡_河海大学版_FC

[![build](https://github.com/yin1999/healthreport_fc/actions/workflows/Build.yml/badge.svg)](https://github.com/yin1999/healthreport_fc/actions/workflows/Build.yml)

本项目用于河海大学健康打卡，此项目为基于[健康打卡_河海大学版](https://github.com/yin1999/healthreport)的函数计算版本。  

当前仅适用于[阿里云函数计算](https://fc.console.aliyun.com/)，若使用腾讯云的相关产品，请自行修改代码

## 编译教程

若想要直接使用，请转至[下载发行版](#下载发行版)

1. 环境配置，以CentOS 7为例

    安装软件：Golang[>= 1.16]、git、make、zip

       yum install -y golang git make zip

2. 通过源码下载、编译

       go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct  #配置Goproxy

       #下载
       git clone https://github.com/yin1999/healthreport_fc.git
       cd healthreport_fc

3. 编译

       make # 生成二进制文件bootstrap.zip

## 使用说明

### 下载发行版

若已自行编译二进制文件，请跳过此步骤。

直接下载[release](https://github.com/yin1999/healthreport_fc/releases/latest)中的**bootstrap.zip**文件

### 函数计算配置

请按照以下步骤进行配置，配置完成后即可

1. [点击](https://fc.console.aliyun.com/)进入函数计算控制台，选择**服务及函数**，在对话框中输入必要信息，然后点击下一步

|项目|示例|说明|
|---|---|---|
|服务名称|healthReport|随意填写，符合要求即可|
|服务描述|健康填报|随意填写|
|绑定日志|可选项|可用于查看运行日志信息，会产生一定费用|
|链路追踪|不勾选|未使用|

2. 新建函数方式选择：**事件函数**，点击下一步

3. 按以下要求进行配置，然后点击下一步

|项目|示例|说明|
|---|---|---|
|函数名称|healthReport|随意填写，符合要求即可，因使用二进制文件，**函数名称**无需正确配置|
|运行环境|Custom Runtime|需选择**Custom Runtime**，以上传自定义运行时文件|
|上传代码|代码包上传 (选择二进制文件：**bootstrap.zip**)|上传已编译完成的自定义运行时|
|函数实例类型|弹性实例|弹性实例已能满足性能需求|
|函数入口|index.handler|默认无需修改，自定义运行时未使用该项配置|
|函数执行内存|128MB|该运行时仅占用20MB以内内存，选择最小的128MB即可|
|超时时间|45秒|函数设定打卡超时时间默认为30秒，比该值稍大即可|
|单实例并发度|1|无并发需求，单次触发执行一次|

4. 跳转至函数界面后，选择**触发器**选项，点击**创建触发器**，触发器配置如下，完成配置后点击**确定**即可

|项目|示例|说明|
|---|---|---|
|服务类型|定时触发器|选择定时触发器以每天自动触发健康填报函数|
|触发器名称|timeTrigger|随意填写，符合要求即可|
|触发器版本|LATEST|默认无需修改|
|Cron表达式|5 1/30 2-4 * * *|自行设定，示例表达式指在**UTC**时间下，每天**02:01:05-04:01:05**（北京时间为**UTC+8**，即**10:01:05-12:01:05**），每隔**30分钟**自动触发函数运行，完成一次健康填报（因每一次健康填报不一定成功，可选择设定较短的触发时间30可缩短，以在当天完成多次填报），可在[在线工具](https://tool.lu/crontab/)（类型选择Java）进行表达式设定查看|
|启用触发器|是|启用后即可触发函数运行，以后若需暂时停用触发器，直接将其设置成**否**即可|
|触发消息|例："1862X10XXX Password"(不包含双引号) |填写学校健康填报用户名、密码(用户名与密码之间使用一个**换行符**、**空格**或**制表符**隔开即可)，示例指用户名为**1862X10XXX**，密码为**Password**|

5. 已完成触发器配置
