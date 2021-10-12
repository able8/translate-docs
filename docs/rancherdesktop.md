# Kubernetes and container management on the desktop

[Install Rancher Desktop](https://github.com/rancher-sandbox/rd/releases)

#### Great For

##### Building, running, pushing, and pulling containers

##### Running your choice of Kubernetes versions

##### Local application development in Kubernetes

## What is Rancher Desktop?

Rancher Desktop is an open-source desktop application for Mac and Windows. It provides Kubernetes and container management.
You can choose the version of Kubernetes you want to run. You can build, push, pull, and _run_ container images.
The container images you build can be run by Kubernetes immediately without the need for a registry.

* * *

## Why Rancher Desktop?

### Kubernetes Made Simple

Getting started with Kubernetes on your desktop can be a project. Especially if you want to match the version of Kubernetes you run locally to the one you run in production. Rancher Desktop makes it as easy as setting a preference.

### Built On Proven Projects

Rancher Desktop leverages proven projects to do the dirty work. That includes containerd, k3s, kubectl, and more. These projects have demonstrated themselves as trustworthy and provide a foundation you can trust.

### Coupled Container Management

Container management to build, push, and pull images and run containers. It uses the same container runtime as Kubernetes. Built images are immediately available to use in your local workloads without any pushing, pulling, or copying.

* * *

## How it Works

Rancher Desktop is an electron based application that wraps other tools while itself providing the user experience to create a simple experience.
On MacOS Rancher Desktop leverages a virtual machine to run containerd and Kubernetes. Windows Subsystem for Linux v2 is leveraged for Windows systems.
All you need to do is download and run the application.

![Racher Desktop Architecture](http://rancherdesktop.io/images/Arch_v2.svg)

## Get Started

### 1\. Download Rancher Desktop

[Download Rancher Desktop](https://github.com/rancher-sandbox/rd/releases)

### 2\. Run The App

Run the app you downloaded and it will take care of the rest.

© 2021 SUSE. Rancher Desktop is an open source project of the [SUSE](https://suse.com) [Rancher](https://rancher.com) Engineering group.
