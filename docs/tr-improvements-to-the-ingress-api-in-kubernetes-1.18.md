# Improvements to the Ingress API in Kubernetes 1.18

# Kubernetes 1.18 中 Ingress API 的改进

Thursday, April 02, 2020

**Authors:** Rob Scott (Google), Christopher M Luciano (IBM)

The Ingress API in Kubernetes has enabled a large number of controllers to provide simple and powerful ways to manage inbound network traffic to Kubernetes workloads. In Kubernetes 1.18, we've made 3 significant additions to this API:

Kubernetes 中的 Ingress API 使大量控制器能够提供简单而强大的方法来管理到 Kubernetes 工作负载的入站网络流量。在 Kubernetes 1.18 中，我们对该 API 进行了 3 项重要的添加：

- A new`pathType` field that can specify how Ingress paths should be matched.
- A new`IngressClass` resource that can specify how Ingresses should be implemented by controllers.
- Support for wildcards in hostnames.

- 一个新的`pathType` 字段，可以指定入口路径的匹配方式。
- 一个新的`IngressClass` 资源，可以指定控制器应该如何实现 Ingress。
- 支持主机名中的通配符。

## Better Path Matching With Path Types

## 更好的路径匹配路径类型

The new concept of a path type allows you to specify how a path should be matched. There are three supported types:

路径类型的新概念允许您指定路径的匹配方式。支持三种类型：

- **ImplementationSpecific (default):** With this path type, matching is up to the controller implementing the `IngressClass`. Implementations can treat this as a separate `pathType` or treat it identically to the `Prefix` or `Exact` path types.
- **Exact:** Matches the URL path exactly and with case sensitivity.
- **Prefix:** Matches based on a URL path prefix split by `/`. Matching is case sensitive and done on a path element by element basis.

- **ImplementationSpecific（默认）：** 对于这种路径类型，匹配取决于实现 `IngressClass` 的控制器。实现可以将其视为单独的“pathType”或将其视为与“Prefix”或“Exact”路径类型相同。
- **Exact:** 精确匹配 URL 路径并区分大小写。
- **Prefix:** 基于被`/`分割的URL路径前缀匹配。匹配区分大小写，并在逐个路径元素的基础上完成。

## Extended Configuration With Ingress Classes

## 带有入口类的扩展配置

The Ingress resource was designed with simplicity in mind, providing a simple set of fields that would be applicable in all use cases. Over time, as use cases evolved, implementations began to rely on a long list of custom annotations for further configuration. The new `IngressClass` resource provides a way to replace some of those annotations.

Ingress 资源的设计考虑到了简单性，提供了一组适用于所有用例的简单字段。随着时间的推移，随着用例的发展，实现开始依赖一长串自定义注释进行进一步配置。新的 IngressClass 资源提供了一种替换其中一些注释的方法。

Each `IngressClass` specifies which controller should implement Ingresses of the class and can reference a custom resource with additional parameters.

每个 IngressClass 指定哪个控制器应该实现类的 Ingress，并且可以引用带有附加参数的自定义资源。

```yaml
apiVersion: "networking.k8s.io/v1beta1"
kind: "IngressClass"
metadata:
  name: "external-lb"
spec:
  controller: "example.com/ingress-controller"
  parameters:
    apiGroup: "k8s.example.com/v1alpha"
    kind: "IngressParameters"
    name: "external-lb"
```


### Specifying the Class of an Ingress

### 指定入口的类

A new `ingressClassName` field has been added to the Ingress spec that is used to reference the `IngressClass` that should be used to implement this Ingress.

Ingress 规范中添加了一个新的“ingressClassName”字段，用于引用应该用于实现此 Ingress 的“IngressClass”。

### Deprecating the Ingress Class Annotation

### 弃用 Ingress 类注释

Before the `IngressClass` resource was added in Kubernetes 1.18, a similar concept of Ingress class was often specified with a `kubernetes.io/ingress.class` annotation on the Ingress. Although this annotation was never formally defined, it was widely supported by Ingress controllers, and should now be considered formally deprecated.

在 Kubernetes 1.18 中添加 `IngressClass` 资源之前，通常在 Ingress 上使用 `kubernetes.io/ingress.class` 注释指定类似的 Ingress 类概念。虽然这个注解从未正式定义，但它得到了 Ingress 控制器的广泛支持，现在应该被视为正式弃用。

### Setting a Default IngressClass

### 设置默认 IngressClass

It’s possible to mark a specific `IngressClass` as default in a cluster. Setting the
`ingressclass.kubernetes.io/is-default-class` annotation to true on an
IngressClass resource will ensure that new Ingresses without an `ingressClassName` specified will be assigned this default `IngressClass`.

可以将特定的“IngressClass”标记为集群中的默认值。设置
`ingressclass.kubernetes.io/is-default-class` 注释为 true
IngressClass 资源将确保没有指定 `ingressClassName` 的新 Ingress 将被分配这个默认的 `IngressClass`。

## Support for Hostname Wildcards

## 支持主机名通配符

Many Ingress providers have supported wildcard hostname matching like `*.foo.com` matching `app1.foo.com`, but until now the spec assumed an exact FQDN match of the host. Hosts can now be precise matches (for example “ `foo.bar.com`”) or a wildcard (for example “ `*.foo.com`”). Precise matches require that the http host header matches the Host setting. Wildcard matches require the http host header is equal to the suffix of the wildcard rule.

许多 Ingress 提供者都支持通配符主机名匹配，比如 `*.foo.com` 匹配 `app1.foo.com`，但直到现在规范假设主机的 FQDN 完全匹配。主机现在可以是精确匹配（例如“`foo.bar.com`”）或通配符（例如“`*.foo.com`”）。精确匹配要求 http 主机标头与主机设置匹配。通配符匹配要求 http 主机头等于通配符规则的后缀。

HostHost headerMatch?`*.foo.com``bar.foo.com`Matches based on shared suffix`*.foo.com``baz.bar.foo.com`No match, wildcard only covers a single DNS label`* .foo.com``foo.com`No match, wildcard only covers a single DNS label

HostHost headerMatch?`*.foo.com``bar.foo.com`基于共享后缀匹配`*.foo.com``baz.bar.foo.com`不匹配，通配符只覆盖单个DNS标签`* .foo.com``foo.com`不匹配，通配符只覆盖单个DNS标签

### Putting it All Together

### 把它们放在一起

These new Ingress features allow for much more configurability. Here’s an example of an Ingress that makes use of pathType, `ingressClassName`, and a hostname wildcard:

这些新的 Ingress 功能允许更多的可配置性。这是一个使用 pathType、`ingressClassName` 和主机名通配符的 Ingress 示例：

```yaml
apiVersion: "networking.k8s.io/v1beta1"
kind: "Ingress"
metadata:
  name: "example-ingress"
spec:
  ingressClassName: "external-lb"
  rules:
  - host: "*.example.com"
    http:
      paths:
      - path: "/example"
        pathType: "Prefix"
        backend:
          serviceName: "example-service"
          servicePort: 80
```


### Ingress Controller Support 

### 入口控制器支持

Since these features are new in Kubernetes 1.18, each Ingress controller implementation will need some time to develop support for these new features. Check the documentation for your preferred Ingress controllers to see when they will support this new functionality.

由于这些功能是 Kubernetes 1.18 中的新功能，因此每个 Ingress 控制器实现都需要一些时间来开发对这些新功能的支持。查看您首选的 Ingress 控制器的文档，了解它们何时支持此新功能。

## The Future of Ingress

## Ingress 的未来

The Ingress API is on pace to graduate from beta to a stable API in Kubernetes 1.19. It will continue to provide a simple way to manage inbound network traffic for Kubernetes workloads. This API has intentionally been kept simple and lightweight, but there has been a desire for greater configurability for more advanced use cases.

Ingress API 即将从测试版升级为 Kubernetes 1.19 中的稳定 API。它将继续提供一种简单的方法来管理 Kubernetes 工作负载的入站网络流量。这个 API 有意保持简单和轻量级，但人们希望为更高级的用例提供更大的可配置性。

Work is currently underway on a new highly configurable set of APIs that will provide an alternative to Ingress in the future. These APIs are being referred to as the new “Service APIs”. They are not intended to replace any existing APIs, but instead provide a more configurable alternative for complex use cases. For more information, check out the [Service APIs repo on GitHub](http://github.com/kubernetes-sigs/service-apis).

目前正在开发一组新的高度可配置的 API，这些 API 将在未来提供 Ingress 的替代方案。这些 API 被称为新的“服务 API”。它们并不打算替换任何现有的 API，而是为复杂的用例提供更可配置的替代方案。有关更多信息，请查看 [GitHub 上的服务 API 存储库](http://github.com/kubernetes-sigs/service-apis)。

- [← Previous](http://kubernetes.io/blog/2020/04/01/kubernetes-1.18-feature-server-side-apply-beta-2/)
[Next →](http://kubernetes.io/blog/2020/04/03/kubernetes-1-18-feature-windows-csi-support-alpha/) 

- [← 上一页](http://kubernetes.io/blog/2020/04/01/kubernetes-1.18-feature-server-side-apply-beta-2/)
[下一步→](http://kubernetes.io/blog/2020/04/03/kubernetes-1-18-feature-windows-csi-support-alpha/)

