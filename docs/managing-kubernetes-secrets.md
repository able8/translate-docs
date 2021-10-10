# Protecting Kubernetes Secrets: A Practical Guide

January 15, 2019

While secrets are critical for the operation of production systems, exposing those secrets puts those systems at risk. Kubernetes does not provide robust mechanisms to encrypt, manage, and share secrets across a Kubernetes cluster. You will probably leverage secrets management solutions like Vault, but you’ll quickly discover they do not provide a full solution in a container environment.
This guide will help you understand the complex problem of secrets management and security and the solutions at your disposal.
In this article you will learn:

- [What are Kubernetes secrets and how they work](http://blog.aquasec.com#what)
- [Problems with the built-in secrets mechanism in Kubernetes](http://blog.aquasec.com#problems)
- [Third-party secrets management tools and their limitations](http://blog.aquasec.com#tools)
- [Using cloud native security platforms to secure secrets at the container level](http://blog.aquasec.com#aqua)

### What are Kubernetes Secrets and are they Secure?

It is very bad practice to store sensitive data, such as passwords, authentication tokens and SSH keys, in plaintext on a container. However, containers need this data to perform basic operations like integrating with other systems. To this end, Kubernetes provides an object called Secret, which you can use to store sensitive data. Below is an example of a secret.

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

However, secrets give you more control over access and usage of passwords, keys, etc. Kubernetes can either mount secrets separately from the pods that use them, or save them as environment variables.

The built-in secrets mechanism in Kubernetes provides basic security capabilities such as:

- **Enabling encryption at rest** (see [documentation](https://kubernetes.io/docs/tasks/administer-cluster/encrypt-data/))
- **Defining authorization policies** (see [documentation](https://kubernetes.io/docs/reference/access-authn-authz/authorization/))
- **White listing access to specific container instances** (see [documentation](https://kubernetes.io/docs/reference/access-authn-authz/rbac/#referring-to-resources))

However these basic security measures are not enforced by default, and even if they are enabled, are not enough for most organizations. Many are complementing Kubernetes with third-party secret management tools—this will be the subject of the second part of this blog.

### How Do Kubernetes Secrets Work?

There are two types of secrets in Kubernetes:

- **Built-in secrets**—Kubernetes Service Accounts automatically create secrets and attach them to containers with API Credentials. This mechanism can be disabled or overridden if it raises security concerns (see the [official documentation](https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/)).
- **Custom secrets**—you can define your own sensitive data and create a secret to store it.

There are two primary ways of creating Kubernetes secrets:

- **Automatically—** using kubectl create secret **—** specify one or more files that include sensitive data, and these files are then packaged as a secret.

kubectl create secret generic db-user-pass --from-file=./username.txt --from-file=./password.txt

- **Manually**—create a secret in a JSON or YAML file, then write the code of the object and create a secret from it using kubectl apply



kubectl get secret mysecret -o yaml

**To decode a secret** from Base64 encoding to plaintext, use kubectl get secret to view the content of the secret. Then decode the sensitive data like this:
echo 'MWYyZDFlMmU2N2Rm' \| base64 --decode

### Problems with the Built-in Secrets Mechanism in Kubernetes

The Kubernetes project [documented](https://kubernetes.io/docs/concepts/configuration/secret/#security-properties) several security risks affecting the built-in Kubernetes secrets mechanism, which users should pay attention to:

- **Securing etcd**—secret data is stored in etcd. By default, etcd data is not encrypted and neither are your secrets. You should enable encryption at rest, limit access to etcd to admin users only, and safely dispose of disks where etcd data was formerly stored.
- **Use SSL/TLS**—when running etcd in a cluster, you must use secure peer-to-peer communication.
- **You can’t share the manifest file or check it into a repo**—commonly, secrets are configured using JSON or YAML files, with the secret encoded in base64. If you share or check in these manifest files, the secret is compromised. This makes it difficult to manage secrets as part of your development workflow.
- **Ensure applications don’t expose the secret**—even if the secret is stored and transmitted securely, an application consuming the secret may store it unsecurely, may log it, or transmit it to a third party.
- **Users who consume a secret can see its value**—any user who creates a pod that uses a secret has access to that secret, even if the API Server policy does not allow the user to view the secret.
- **Root exploit—** anyone with root access on any node can read any secret, because they can impersonate the kubelet. Kubernetes does not currently send secrets on a “need to know” basis; it exposes secrets to anyone on any node with root access.

In addition, Kubernetes secrets have some important usability issues, pointed out in the excellent post by [Omer Levi Hevroni](https://itnext.io/can-kubernetes-keep-a-secret-it-all-depends-what-tool-youre-using-498e5dee9c25):

- **No visibility or change management**—secrets are critical to the operation of many services and can break a service in production (e.g., if credentials are missing or wrong). In order to manage a development pipeline and troubleshoot production issues, you need to track changes to secrets. Kubernetes provides an audit mechanism but it’s not straightforward, and there is no way to track changes to secrets using version control.
- **Secrets mounted as volumes are unwieldy**—secrets can be stored as environment variables or mounted as a volume. The former technique is widely agreed to be less secure. If you opt for volumes, things quickly get complex when you have a large number of keys. Kubernetes creates one file per key, and you need to read all these files from within the application. There are workarounds, but they can be equally complex.
- **Not a zero-trust system—** once a user is allowed to receive a secret, that user receives the secret decrypted. It would be much better to define granular permissions that specify who can encrypt a secret and prevent anyone from directly decrypting a secret. Secrets should only be decrypted on demand, when they are actually needed. Unfortunately, Kubernetes does not allow this.

For these and other reasons, most practitioners are opting for third-party management tools to help them take control of Kubernetes secrets.

### Third-Party Secret Management Tools

Here are some popular tools that can help you achieve better security and usability for secrets. Their primary value is that they provide a centralized, robust mechanism for encrypting secrets and granting access to secrets. All these tools support Kubernetes, but they are not purpose-built for container workloads.

Secret management solutions fall into two broad categories:

- **Cloud provider tools**—including [AWS Secrets Manager](https://aws.amazon.com/secrets-manager/), [Google Cloud Platform KMS](https://cloud.google.com/kms/) and [Azure Key Vault](https://azure.microsoft.com/en-in/services/key-vault/)—help encrypt secrets within each cloud environment, automatically rotate secret values, manage access to secrets with policies, and perform central auditing of secrets.
- **Open source tools**— [Hashicorp Vault](https://www.vaultproject.io/) provides secrets management and data protection, with advanced features like dynamic secrets, namespaces, leases, and revocation for secrets data. Other options are [CyberArk](https://www.conjur.org/) (with an open source version called [Conjur](https://www.conjur.org/)) and [Confidant](https://lyft.github.io/confidant/).

### Limitations of Secret Management in a Container Environment

There is no doubt that secrets management tools can provide value for Kubernetes projects. However, even when you are armed with a secret management solution, you will still be limited in your ability to manage and monitor secrets.

Secrets management systems have the following serious limitations in a container environment:

- **Managing which containers have access to which secrets**—secrets should only be accessible to containers that actually need them. Secrets management tools don’t provide a mechanism to map secrets to relevant containers, and don’t let you know which container is actually using which secret.
- **Retrieving secrets from the vault only when actually needed**—nobody wants a secret inside a container image, because this exposes the secret to many more users and processes than necessary. Secrets should be transmitted only when the appropriate container runs, and not before.
- **Storing secrets in containers in a secure way**—even when a container receives a secret, it should not be stored on disk or accessible at the host level. Secrets should only be stored in memory, should only be accessible to the specific container that needs them, and should disappear when the container shuts down. Secrets management tools cannot do this because they don’t have direct access to individual containers.
- **Pushing updated secrets directly to containers**—secrets management tools are great at rotating and modifying secrets to reduce the possibility of exposure. But they can’t automatically update the relevant containers with new values of secrets. Containers can hold out-of-date secrets, which may lead to production issues, or you may be able to update containers after a restart, which results in uptime issues.

### Aqua Security: Secrets Management for Kubernetes and Cloud Native Environments

Aqua Security is a cloud native security platform that secures containerized and serverless applications, from the CI/CD pipeline to runtime production environments.

Aqua's cloud native security platform provides a secrets management solution that lets you centrally control secrets and how containers access those secrets. Aqua can inject secrets into a running container, ensuring that secrets only run in a container’s memory, which is more secure.

![vault](https://blog.aquasec.com/hs-fs/hubfs/Blog/Kubernetes%20Secrets/vault.png?width=1153&name=vault.png)

Aqua integrates with secret management tools including Amazon KMS, HashiCorp Vault, Azure Key Vault, and CyberArk Enterprise Password Vault. If you don’t have an existing secret store, Aqua provides its own encrypted database for storing secrets.

**Here are other ways Aqua can improve security and usability of secrets in Kubernetes:**

1. **Encryption at Rest**—secrets are encrypted at rest through third-party storage. Access to secrets is provided using access tokens or IAM roles of the current host.
2. **No Persistence**—Aqua securely delivers secrets to containers, encrypted at rest, loading them in memory with no persistence on disk, where they are only visible to the container that needs them.
3. **Access Control**—you can control user access to secrets (which may be integrated with Active Directory), and group together containers with similar security features.
4. **Write Only**—once you create a secret, it cannot be seen via the web interface or the API. You can update the content of the secret or delete it, but not read it.
5. **Rotation**—when you modify a secret, the new value is updated in real time on the running container. You do not have to restart the container for the changes in secrets to propagate.
6. **Revocation**—when the vault revokes a secret, Aqua remove those secrets from the containers that use them, with no need to restart the containers.
7. **SSH Encryption**—Aqua enforces mutually authenticated, SHH encrypted communications within a Kubernetes cluster. The Enforcer uses its public key to verify the Gateway and the Enforcer makes itself known via an authentication token.
8. **Administration and usage audit**—Each administrative access event to a secret is logged, and you can see which containers had access to which secrets during runtime.

To see how a container security platform can upgrade your Kubernetes secrets management and help you create a secure development pipeline, [learn more](https://www.aquasec.com/products/aqua-cloud-native-security-platform/) about the Aqua Cloud Native Security Platform.

Rani is the VP of Product Marketing and Strategy at Aqua. Rani has worked in enterprise software companies more than 25 years, spanning project management, product management and marketing, including a decade as VP of marketing for innovative startups in the cyber-security and cloud arenas. Previously Rani was also a management consultant in the London office of Booz & Co. He holds an MBA from INSEAD in Fontainebleau, France. Rani is an avid wine geek, and a slightly less avid painter and electronic music composer.