# Protecting Kubernetes Secrets: A Practical Guide

# 保护 Kubernetes 的秘密：实用指南

January 15, 2019

2019 年 1 月 15 日

While secrets are critical for the operation of production systems, exposing those secrets puts those systems at risk. Kubernetes does not provide robust mechanisms to encrypt, manage, and share secrets across a Kubernetes cluster. You will probably leverage secrets management solutions like Vault, but you’ll quickly discover they do not provide a full solution in a container environment.
This guide will help you understand the complex problem of secrets management and security and the solutions at your disposal.
In this article you will learn:

虽然机密对于生产系统的运行至关重要，但暴露这些机密会使这些系统面临风险。 Kubernetes 不提供在 Kubernetes 集群中加密、管理和共享机密的强大机制。您可能会利用 Vault 等机密管理解决方案，但您很快就会发现它们并没有在容器环境中提供完整的解决方案。
本指南将帮助您了解机密管理和安全的复杂问题以及您可以使用的解决方案。
在本文中，您将了解到：

- [What are Kubernetes secrets and how they work](http://blog.aquasec.com#what)
- [Problems with the built-in secrets mechanism in Kubernetes](http://blog.aquasec.com#problems)
- [Third-party secrets management tools and their limitations](http://blog.aquasec.com#tools)
- [Using cloud native security platforms to secure secrets at the container level](http://blog.aquasec.com#aqua)

- [什么是 Kubernetes 的秘密及其工作原理](http://blog.aquasec.com#what)
- [Kubernetes 内置秘密机制的问题](http://blog.aquasec.com#problems)
- [第三方机密管理工具及其限制](http://blog.aquasec.com#tools)
- [使用云原生安全平台在容器级别保护机密](http://blog.aquasec.com#aqua)

### What are Kubernetes Secrets and are they Secure?

### Kubernetes 的秘密是什么？它们安全吗？

It is very bad practice to store sensitive data, such as passwords, authentication tokens and SSH keys, in plaintext on a container. However, containers need this data to perform basic operations like integrating with other systems. To this end, Kubernetes provides an object called Secret, which you can use to store sensitive data. Below is an example of a secret.

在容器中以明文形式存储敏感数据（例如密码、身份验证令牌和 SSH 密钥）是非常糟糕的做法。但是，容器需要这些数据来执行基本操作，例如与其他系统集成。为此，Kubernetes 提供了一个名为 Secret 的对象，您可以使用它来存储敏感数据。下面是一个秘密的例子。

```yml
apiVersion: v1
data:
 super-secret: aHR0cDovL2pvYnMuc29sdXRvLmNvbS9hcHBseS95MDVYSXEvUHJvZHVjdGlvbi1FbmdpbmVlci1EZXZPcHM=
kind: Secret
metadata:
username: admin
password: mUV2Sw7eLbfp
type: Opaque
```

Placing sensitive info into a secret object does not automatically make it secure. By default, data in Kubernetes secrets is stored in Base64 encoding, which is practically the same as plaintext.

将敏感信息放入秘密对象并不会自动使其安全。默认情况下，Kubernetes Secret 中的数据以 Base64 编码存储，这实际上与明文相同。

However, secrets give you more control over access and usage of passwords, keys, etc. Kubernetes can either mount secrets separately from the pods that use them, or save them as environment variables.

但是，secrets 使您可以更好地控制密码、密钥等的访问和使用。Kubernetes 可以将 secrets 与使用它们的 pod 分开挂载，或者将它们保存为环境变量。

The built-in secrets mechanism in Kubernetes provides basic security capabilities such as:

Kubernetes 中内置的秘密机制提供了基本的安全功能，例如：

- **Enabling encryption at rest** (see [documentation](https://kubernetes.io/docs/tasks/administer-cluster/encrypt-data/))
- **Defining authorization policies** (see [documentation](https://kubernetes.io/docs/reference/access-authn-authz/authorization/))
- **White listing access to specific container instances** (see [documentation](https://kubernetes.io/docs/reference/access-authn-authz/rbac/#referring-to-resources))

- **启用静态加密**（参见 [文档](https://kubernetes.io/docs/tasks/administer-cluster/encrypt-data/))
- **定义授权策略**（参见[文档](https://kubernetes.io/docs/reference/access-authn-authz/authorization/))
- **对特定容器实例的白名单访问**（请参阅[文档](https://kubernetes.io/docs/reference/access-authn-authz/rbac/#referring-to-resources))

However these basic security measures are not enforced by default, and even if they are enabled, are not enough for most organizations. Many are complementing Kubernetes with third-party secret management tools—this will be the subject of the second part of this blog.

然而，这些基本安全措施在默认情况下不会强制执行，即使启用了它们，对于大多数组织来说也不够。许多人正在用第三方秘密管理工具补充 Kubernetes——这将是本博客第二部分的主题。

### How Do Kubernetes Secrets Work?

### Kubernetes Secrets 如何工作？

There are two types of secrets in Kubernetes:

Kubernetes 中有两种类型的秘密：

- **Built-in secrets**—Kubernetes Service Accounts automatically create secrets and attach them to containers with API Credentials. This mechanism can be disabled or overridden if it raises security concerns (see the [official documentation](https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/)).
- **Custom secrets**—you can define your own sensitive data and create a secret to store it.

- **内置机密** - Kubernetes 服务帐户自动创建机密并将其附加到具有 API 凭据的容器。如果此机制引起安全问题，则可以禁用或覆盖此机制（请参阅 [官方文档](https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/))。
- **自定义秘密**——您可以定义自己的敏感数据并创建一个秘密来存储它。

There are two primary ways of creating Kubernetes secrets:

创建 Kubernetes Secret 有两种主要方式：

- **Automatically—** using kubectl create secret **—** specify one or more files that include sensitive data, and these files are then packaged as a secret.

- **自动——** 使用 kubectl create secret **——** 指定一个或多个包含敏感数据的文件，然后将这些文件打包为机密。

kubectl create secret generic db-user-pass --from-file=./username.txt --from-file=./password.txt

kubectl 创建秘密通用 db-user-pass --from-file=./username.txt --from-file=./password.txt

- **Manually**—create a secret in a JSON or YAML file, then write the code of the object and create a secret from it using kubectl apply

- **手动**——在 JSON 或 YAML 文件中创建一个秘密，然后编写对象的代码并使用 kubectl apply 从中创建一个秘密

kubectl get secret mysecret -o yaml

kubectl 获取秘密 mysecret -o yaml

**To decode a secret** from Base64 encoding to plaintext, use kubectl get secret to view the content of the secret. Then decode the sensitive data like this:
echo 'MWYyZDFlMmU2N2Rm' \| base64 --decode

**要将秘密**从 Base64 编码解码为明文，请使用 kubectl get secret 查看秘密的内容。然后像这样解码敏感数据：
回声'MWYyZDFlMmU2N2Rm' \| base64 --解码

### Problems with the Built-in Secrets Mechanism in Kubernetes

### Kubernetes 内置 Secrets 机制的问题

The Kubernetes project [documented](https://kubernetes.io/docs/concepts/configuration/secret/#security-properties) several security risks affecting the built-in Kubernetes secrets mechanism, which users should pay attention to: 

Kubernetes 项目[文档化](https://kubernetes.io/docs/concepts/configuration/secret/#security-properties) 影响内置 Kubernetes Secrets 机制的几个安全风险，用户应注意：

- **Securing etcd**—secret data is stored in etcd. By default, etcd data is not encrypted and neither are your secrets. You should enable encryption at rest, limit access to etcd to admin users only, and safely dispose of disks where etcd data was formerly stored.
- **Use SSL/TLS**—when running etcd in a cluster, you must use secure peer-to-peer communication.
- **You can’t share the manifest file or check it into a repo**—commonly, secrets are configured using JSON or YAML files, with the secret encoded in base64. If you share or check in these manifest files, the secret is compromised. This makes it difficult to manage secrets as part of your development workflow.
- **Ensure applications don’t expose the secret**—even if the secret is stored and transmitted securely, an application consuming the secret may store it unsecurely, may log it, or transmit it to a third party.
- **Users who consume a secret can see its value**—any user who creates a pod that uses a secret has access to that secret, even if the API Server policy does not allow the user to view the secret.
- **Root exploit—** anyone with root access on any node can read any secret, because they can impersonate the kubelet. Kubernetes does not currently send secrets on a “need to know” basis; it exposes secrets to anyone on any node with root access.

- **Securing etcd**——秘密数据存储在etcd中。默认情况下，etcd 数据未加密，您的机密也未加密。您应该启用静态加密，将 etcd 的访问权限限制为仅限管理员用户，并安全地处理以前存储 etcd 数据的磁盘。
- **使用 SSL/TLS**——在集群中运行 etcd 时，必须使用安全的对等通信。
- **您无法共享清单文件或将其签入存储库** - 通常，机密使用 JSON 或 YAML 文件配置，机密以 base64 编码。如果您共享或签入这些清单文件，则机密会泄露。这使得在开发工作流程中管理机密变得困难。
- **确保应用程序不会泄露机密**——即使机密被安全地存储和传输，使用机密的应用程序也可能不安全地存储它、记录它或将它传输给第三方。
- **使用秘密的用户可以看到它的价值**——任何创建使用秘密的 pod 的用户都可以访问该秘密，即使 API 服务器策略不允许用户查看秘密。
- **Root 漏洞利用——** 在任何节点上拥有 root 访问权限的任何人都可以读取任何机密，因为他们可以模拟 kubelet。 Kubernetes 目前不会在“需要知道”的基础上发送机密；它向任何具有 root 访问权限的节点上的任何人公开秘密。

In addition, Kubernetes secrets have some important usability issues, pointed out in the excellent post by [Omer Levi Hevroni](https://itnext.io/can-kubernetes-keep-a-secret-it-all-depends-what-tool-youre-using-498e5dee9c25):

此外，Kubernetes Secrets 有一些重要的可用性问题，在 [Omer Levi Hevroni] 的优秀帖子中指出(https://itnext.io/can-kubernetes-keep-a-secret-it-all-depends-what- tool-youre-using-498e5dee9c25）：

- **No visibility or change management**—secrets are critical to the operation of many services and can break a service in production (e.g., if credentials are missing or wrong). In order to manage a development pipeline and troubleshoot production issues, you need to track changes to secrets. Kubernetes provides an audit mechanism but it’s not straightforward, and there is no way to track changes to secrets using version control.
- **Secrets mounted as volumes are unwieldy**—secrets can be stored as environment variables or mounted as a volume. The former technique is widely agreed to be less secure. If you opt for volumes, things quickly get complex when you have a large number of keys. Kubernetes creates one file per key, and you need to read all these files from within the application. There are workarounds, but they can be equally complex.
- **Not a zero-trust system—** once a user is allowed to receive a secret, that user receives the secret decrypted. It would be much better to define granular permissions that specify who can encrypt a secret and prevent anyone from directly decrypting a secret. Secrets should only be decrypted on demand, when they are actually needed. Unfortunately, Kubernetes does not allow this.

- **无可见性或变更管理**——秘密对许多服务的运行至关重要，可能会破坏生产中的服务（例如，如果凭据丢失或错误）。为了管理开发管道和解决生产问题，您需要跟踪对机密的更改。 Kubernetes 提供了一种审计机制，但它并不简单，并且无法使用版本控制来跟踪对机密的更改。
- **作为卷安装的机密很笨拙**——机密可以存储为环境变量或作为卷安装。前一种技术被广泛认为不太安全。如果您选择卷，当您拥有大量密钥时，事情很快就会变得复杂。 Kubernetes 为每个键创建一个文件，您需要从应用程序中读取所有这些文件。有一些解决方法，但它们可能同样复杂。
- **不是零信任系统——** 一旦允许用户接收机密，该用户就会收到解密后的机密。最好定义细化权限，指定谁可以加密机密并防止任何人直接解密机密。秘密应该只在实际需要时才根据需要解密。不幸的是，Kubernetes 不允许这样做。

For these and other reasons, most practitioners are opting for third-party management tools to help them take control of Kubernetes secrets.

由于这些和其他原因，大多数从业者选择第三方管理工具来帮助他们控制 Kubernetes 的秘密。

### Third-Party Secret Management Tools

### 第三方机密管理工具

Here are some popular tools that can help you achieve better security and usability for secrets. Their primary value is that they provide a centralized, robust mechanism for encrypting secrets and granting access to secrets. All these tools support Kubernetes, but they are not purpose-built for container workloads.

以下是一些流行的工具，可以帮助您实现更好的机密安全性和可用性。它们的主要价值在于它们提供了一种集中的、强大的机制来加密秘密和授予对秘密的访问权。所有这些工具都支持 Kubernetes，但它们并不是专门为容器工作负载而构建的。

Secret management solutions fall into two broad categories:

机密管理解决方案分为两大类：

- **Cloud provider tools**—including [AWS Secrets Manager](https://aws.amazon.com/secrets-manager/), [Google Cloud Platform KMS](https://cloud.google.com/kms/) and [Azure Key Vault](https://azure.microsoft.com/en-in/services/key-vault/)—help encrypt secrets within each cloud environment, automatically rotate secret values, manage access to secrets with policies , and perform central auditing of secrets. 

- **云提供商工具**——包括 [AWS Secrets Manager](https://aws.amazon.com/secrets-manager/)、[Google Cloud Platform KMS](https://cloud.google.com/kms/) 和 [Azure Key Vault](https://azure.microsoft.com/en-in/services/key-vault/)—帮助加密每个云环境中的机密、自动轮换机密值、使用策略管理对机密的访问，并对机密进行集中审计。

- **Open source tools**— [Hashicorp Vault](https://www.vaultproject.io/) provides secrets management and data protection, with advanced features like dynamic secrets, namespaces, leases, and revocation for secrets data. Other options are [CyberArk](https://www.conjur.org/) (with an open source version called [Conjur](https://www.conjur.org/)) and [Confidant](https://lyft.github.io/confidant/).

- **开源工具**— [Hashicorp Vault](https://www.vaultproject.io/) 提供机密管理和数据保护，具有动态机密、命名空间、租用和机密数据撤销等高级功能。其他选项有 [CyberArk](https://www.conjur.org/)（开源版本称为[Conjur](https://www.conjur.org/)）和 [Confidant](https://lyft.github.io/confidant/)。

### Limitations of Secret Management in a Container Environment

### 容器环境中机密管理的局限性

There is no doubt that secrets management tools can provide value for Kubernetes projects. However, even when you are armed with a secret management solution, you will still be limited in your ability to manage and monitor secrets.

毫无疑问，秘密管理工具可以为 Kubernetes 项目提供价值。但是，即使您拥有机密管理解决方案，您管理和监控机密的能力仍然会受到限制。

Secrets management systems have the following serious limitations in a container environment:

Secrets 管理系统在容器环境中存在以下严重限制：

- **Managing which containers have access to which secrets**—secrets should only be accessible to containers that actually need them. Secrets management tools don’t provide a mechanism to map secrets to relevant containers, and don’t let you know which container is actually using which secret.
- **Retrieving secrets from the vault only when actually needed**—nobody wants a secret inside a container image, because this exposes the secret to many more users and processes than necessary. Secrets should be transmitted only when the appropriate container runs, and not before.
- **Storing secrets in containers in a secure way**—even when a container receives a secret, it should not be stored on disk or accessible at the host level. Secrets should only be stored in memory, should only be accessible to the specific container that needs them, and should disappear when the container shuts down. Secrets management tools cannot do this because they don’t have direct access to individual containers.
- **Pushing updated secrets directly to containers**—secrets management tools are great at rotating and modifying secrets to reduce the possibility of exposure. But they can’t automatically update the relevant containers with new values of secrets. Containers can hold out-of-date secrets, which may lead to production issues, or you may be able to update containers after a restart, which results in uptime issues.

- **管理哪些容器可以访问哪些秘密**——秘密应该只有真正需要它们的容器才能访问。机密管理工具不提供将机密映射到相关容器的机制，也不会让您知道哪个容器实际使用了哪个机密。
- **仅在实际需要时才从保管库中检索机密**——没有人希望在容器映像中包含机密，因为这会将机密暴露给不必要的用户和进程。只有在适当的容器运行时才应该传输秘密，而不是之前。
- **以安全的方式在容器中存储机密**——即使容器收到机密，也不应将其存储在磁盘上或在主机级别访问。秘密应该只存储在内存中，应该只能被需要它们的特定容器访问，并且应该在容器关闭时消失。秘密管理工具无法做到这一点，因为它们无法直接访问单个容器。
- **将更新的秘密直接推送到容器**——秘密管理工具非常擅长轮换和修改秘密以减少暴露的可能性。但是他们不能用新的秘密值自动更新相关的容器。容器可能保存过时的机密，这可能会导致生产问题，或者您可能能够在重新启动后更新容器，从而导致正常运行时间问题。

### Aqua Security: Secrets Management for Kubernetes and Cloud Native Environments

### Aqua Security：Kubernetes 和云原生环境的秘密管理

Aqua Security is a cloud native security platform that secures containerized and serverless applications, from the CI/CD pipeline to runtime production environments.

Aqua Security 是一个云原生安全平台，可保护从 CI/CD 管道到运行时生产环境的容器化和无服务器应用程序。

Aqua's cloud native security platform provides a secrets management solution that lets you centrally control secrets and how containers access those secrets. Aqua can inject secrets into a running container, ensuring that secrets only run in a container’s memory, which is more secure.

Aqua 的云原生安全平台提供了一个机密管理解决方案，让您可以集中控制机密以及容器如何访问这些机密。 Aqua可以将secrets注入正在运行的容器中，保证secrets只在容器内存中运行，更加安全。

![vault](https://blog.aquasec.com/hs-fs/hubfs/Blog/Kubernetes%20Secrets/vault.png?width=1153&name=vault.png)

Aqua integrates with secret management tools including Amazon KMS, HashiCorp Vault, Azure Key Vault, and CyberArk Enterprise Password Vault. If you don’t have an existing secret store, Aqua provides its own encrypted database for storing secrets.

Aqua 与机密管理工具集成，包括 Amazon KMS、HashiCorp Vault、Azure Key Vault 和 CyberArk Enterprise Password Vault。如果您没有现有的秘密存储，Aqua 提供了自己的加密数据库来存储秘密。

**Here are other ways Aqua can improve security and usability of secrets in Kubernetes:**

**以下是 Aqua 可以提高 Kubernetes 中机密的安全性和可用性的其他方式：**

1. **Encryption at Rest**—secrets are encrypted at rest through third-party storage. Access to secrets is provided using access tokens or IAM roles of the current host.
2. **No Persistence**—Aqua securely delivers secrets to containers, encrypted at rest, loading them in memory with no persistence on disk, where they are only visible to the container that needs them.
3. **Access Control**—you can control user access to secrets (which may be integrated with Active Directory), and group together containers with similar security features.
4. **Write Only**—once you create a secret, it cannot be seen via the web interface or the API. You can update the content of the secret or delete it, but not read it.
5. **Rotation**—when you modify a secret, the new value is updated in real time on the running container. You do not have to restart the container for the changes in secrets to propagate.
6. **Revocation**—when the vault revokes a secret, Aqua remove those secrets from the containers that use them, with no need to restart the containers. 

1. **静态加密**——秘密通过第三方存储进行静态加密。使用访问令牌或当前主机的 IAM 角色提供对机密的访问。
2. **无持久性**——Aqua 安全地将秘密传送到容器，静态加密，将它们加载到内存中而没有持久性在磁盘上，只有需要它们的容器才能看到它们。
3. **访问控制**——您可以控制用户对机密（可能与 Active Directory 集成）的访问，并将具有相似安全功能的容器组合在一起。
4. **Write Only**——一旦你创建了一个秘密，它就无法通过网络界面或 API 看到。您可以更新密钥的内容或删除它，但不能读取它。
5. **Rotation**——当你修改一个secret时，新值会在运行的容器上实时更新。您不必重新启动容器即可传播机密更改。
6. **撤销**——当保管库撤销一个秘密时，Aqua 从使用它们的容器中删除这些秘密，而无需重新启动容器。

7. **SSH Encryption**—Aqua enforces mutually authenticated, SHH encrypted communications within a Kubernetes cluster. The Enforcer uses its public key to verify the Gateway and the Enforcer makes itself known via an authentication token.
8. **Administration and usage audit**—Each administrative access event to a secret is logged, and you can see which containers had access to which secrets during runtime.

7. **SSH 加密**——Aqua 在 Kubernetes 集群内强制执行相互认证的 SHH 加密通信。 Enforcer 使用其公钥来验证网关，而 Enforcer 通过身份验证令牌使自己为人所知。
8. **管理和使用审计**——记录对机密的每个管理访问事件，您可以查看哪些容器在运行时访问了哪些机密。

To see how a container security platform can upgrade your Kubernetes secrets management and help you create a secure development pipeline, [learn more](https://www.aquasec.com/products/aqua-cloud-native-security-platform/) about the Aqua Cloud Native Security Platform.

要了解容器安全平台如何升级您的 Kubernetes 机密管理并帮助您创建安全的开发管道，[了解更多](https://www.aquasec.com/products/aqua-cloud-native-security-platform/)关于 Aqua Cloud 原生安全平台。

Rani is the VP of Product Marketing and Strategy at Aqua. Rani has worked in enterprise software companies more than 25 years, spanning project management, product management and marketing, including a decade as VP of marketing for innovative startups in the cyber-security and cloud arenas. Previously Rani was also a management consultant in the London office of Booz & Co. He holds an MBA from INSEAD in Fontainebleau, France. Rani is an avid wine geek, and a slightly less avid painter and electronic music composer. 

Rani 是 Aqua 的产品营销和战略副总裁。 Rani 在企业软件公司工作超过 25 年，涉及项目管理、产品管理和营销，其中包括在网络安全和云领域的创新初创公司担任营销副总裁十年。此前，Rani 还是 Booz & Co 伦敦办事处的管理顾问。他拥有法国枫丹白露欧洲工商管理学院的 MBA 学位。 Rani 是一位狂热的葡萄酒爱好者，也是一位不太狂热的画家和电子音乐作曲家。

