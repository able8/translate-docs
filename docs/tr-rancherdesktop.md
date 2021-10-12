# Kubernetes and container management on the desktop

# Kubernetes 和桌面上的容器管理

[Install Rancher Desktop](https://github.com/rancher-sandbox/rd/releases)

[安装 Rancher Desktop](https://github.com/rancher-sandbox/rd/releases)

#### Great For

#### 非常适合

##### Building, running, pushing, and pulling containers

##### 构建、运行、推送和拉动容器

##### Running your choice of Kubernetes versions

##### 运行您选择的 Kubernetes 版本

##### Local application development in Kubernetes

##### Kubernetes 本地应用开发

## What is Rancher Desktop?

## 什么是Rancher桌面？

Rancher Desktop is an open-source desktop application for Mac and Windows. It provides Kubernetes and container management.
You can choose the version of Kubernetes you want to run. You can build, push, pull, and _run_ container images.
The container images you build can be run by Kubernetes immediately without the need for a registry.

Rancher Desktop 是适用于 Mac 和 Windows 的开源桌面应用程序。它提供 Kubernetes 和容器管理。
您可以选择要运行的 Kubernetes 版本。您可以构建、推送、拉取和 _run_ 容器映像。
您构建的容器镜像可以立即由 Kubernetes 运行，无需注册。

* * *

* * *

## Why Rancher Desktop?

## 为什么是 Rancher 桌面？

### Kubernetes Made Simple

### Kubernetes 变得简单

Getting started with Kubernetes on your desktop can be a project. Especially if you want to match the version of Kubernetes you run locally to the one you run in production. Rancher Desktop makes it as easy as setting a preference.

在桌面上开始使用 Kubernetes 可以是一个项目。特别是如果您想将本地运行的 Kubernetes 版本与生产中运行的版本相匹配。 Rancher Desktop 使设置首选项变得简单。

### Built On Proven Projects

### 建立在经过验证的项目上

Rancher Desktop leverages proven projects to do the dirty work. That includes containerd, k3s, kubectl, and more. These projects have demonstrated themselves as trustworthy and provide a foundation you can trust.

Rancher Desktop 利用经过验证的项目来完成繁重的工作。这包括 containerd、k3s、kubectl 等。这些项目证明了自己值得信赖，并提供了您可以信赖的基础。

### Coupled Container Management

### 耦合容器管理

Container management to build, push, and pull images and run containers. It uses the same container runtime as Kubernetes. Built images are immediately available to use in your local workloads without any pushing, pulling, or copying.

用于构建、推送和拉取镜像以及运行容器的容器管理。它使用与 Kubernetes 相同的容器运行时。构建的镜像可立即用于本地工作负载，无需任何推送、拉取或复制。

* * *

* * *

## How it Works

##  这个怎么运作

Rancher Desktop is an electron based application that wraps other tools while itself providing the user experience to create a simple experience.
On MacOS Rancher Desktop leverages a virtual machine to run containerd and Kubernetes. Windows Subsystem for Linux v2 is leveraged for Windows systems.
All you need to do is download and run the application.

Rancher Desktop 是一个基于电子的应用程序，它包装了其他工具，同时本身提供了用户体验以创建简单的体验。
在 MacOS Rancher Desktop 上利用虚拟机运行 containerd 和 Kubernetes。适用于 Linux v2 的 Windows 子系统用于 Windows 系统。
您需要做的就是下载并运行该应用程序。

![Racher Desktop Architecture](http://rancherdesktop.io/images/Arch_v2.svg)

## Get Started

## 开始

### 1\. Download Rancher Desktop

### 1. 下载 Rancher 桌面版

[Download Rancher Desktop](https://github.com/rancher-sandbox/rd/releases)

[下载 Rancher 桌面](https://github.com/rancher-sandbox/rd/releases)

### 2\. Run The App

### 2. 运行应用程序

Run the app you downloaded and it will take care of the rest.

运行您下载的应用程序，它会处理剩下的事情。

© 2021 SUSE. Rancher Desktop is an open source project of the [SUSE](https://suse.com)[Rancher](https://rancher.com) Engineering group. 

© 2021 SUSE。 Rancher Desktop 是 [SUSE](https://suse.com)[Rancher](https://rancher.com) 工程组的一个开源项目。

