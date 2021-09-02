# Argo CD RBAC Configuration[¶](https://argo-cd-docs.readthedocs.io/en/latest/operator-manual/rbac/#rbac-configuration)

# Argo CD RBAC 配置[¶](https://argo-cd-docs.readthedocs.io/en/latest/operator-manual/rbac/#rbac-configuration)

From: https://argo-cd-docs.readthedocs.io/en/latest/operator-manual/rbac/

The RBAC feature enables restriction of access to Argo CD resources. Argo CD does not have its own user management system and has only one built-in user `admin`. The `admin` user is a superuser and it has unrestricted access to the system. RBAC requires [SSO configuration](https://argo-cd-docs.readthedocs.io/en/latest/operator-manual/sso/). Once SSO is configured, additional RBAC roles can be defined, and SSO groups can man be mapped to roles.

RBAC 功能可以限制对 Argo CD 资源的访问。 Argo CD 没有自己的用户管理系统，只有一个内置用户 `admin`。 `admin` 用户是超级用户，可以不受限制地访问系统。 RBAC 需要 [SSO 配置](https://argo-cd-docs.readthedocs.io/en/latest/operator-manual/sso/)。配置 SSO 后，可以定义其他 RBAC 角色，并且可以将 SSO 组映射到角色。

## Basic Built-in Roles[¶](https://argo-cd-docs.readthedocs.io/en/latest/operator-manual/rbac/#basic-built-in-roles)

## 基本内置角色[¶](https://argo-cd-docs.readthedocs.io/en/latest/operator-manual/rbac/#basic-built-in-roles)

Argo CD has two pre-defined roles but RBAC configuration allows defining roles and groups (see below).

Argo CD 有两个预定义的角色，但 RBAC 配置允许定义角色和组（见下文）。

- `role:readonly` - read-only access to all resources
- `role:admin` - unrestricted access to all resources

- `role:readonly` - 对所有资源的只读访问
- `role:admin` - 不受限制地访问所有资源

These default built-in role definitions can be seen in [builtin-policy.csv](https://github.com/argoproj/argo-cd/blob/master/assets/builtin-policy.csv)

这些默认的内置角色定义可以在 [builtin-policy.csv](https://github.com/argoproj/argo-cd/blob/master/assets/builtin-policy.csv) 中看到

### RBAC Permission Structure[¶](https://argo-cd-docs.readthedocs.io/en/latest/operator-manual/rbac/#rbac-permission-structure)

### RBAC 权限结构[¶](https://argo-cd-docs.readthedocs.io/en/latest/operator-manual/rbac/#rbac-permission-structure)

Breaking down the permissions definition differs slightly between applications and every other resource type in Argo CD.

Argo CD 中的应用程序和其他所有资源类型之间的权限定义细分略有不同。

- All resources *except* applications permissions (see next bullet):

   `p, <role/user/group>, <resource>, <action>, <object>`

- 所有资源*除了*应用程序权限（见下一个项目符号）：

  `p、<角色/用户/组>、<资源>、<动作>、<对象>`

- Applications (which belong to an AppProject):

   `p, <role/user/group>, <resource>, <action>, <appproject>/<object>`

- 应用程序（属于 AppProject）：

  `p、<role/user/group>、<resource>、<action>、<appproject>/<object>`

### RBAC Resources and Actions[¶](https://argo-cd-docs.readthedocs.io/en/latest/operator-manual/rbac/#rbac-resources-and-actions)

### RBAC 资源和操作[¶](https://argo-cd-docs.readthedocs.io/en/latest/operator-manual/rbac/#rbac-resources-and-actions)

Resources: `clusters`, `projects`, `applications`, `repositories`, `certificates`

资源：`clusters`、`projects`、`applications`、`repositories`、`certificates`

Actions: `get`, `create`, `update`, `delete`, `sync`, `override`, `action`

动作：`get`、`create`、`update`、`delete`、`sync`、`override`、`action`

## Tying It All Together[¶](https://argo-cd-docs.readthedocs.io/en/latest/operator-manual/rbac/#tying-it-all-together)

## 将它们捆绑在一起[¶](https://argo-cd-docs.readthedocs.io/en/latest/operator-manual/rbac/#tying-it-all-together)

Additional roles and groups can be configured in `argocd-rbac-cm` ConfigMap. The example below configures a custom role, named `org-admin`. The role is assigned to any user which belongs to `your-github-org:your-team` group. All other users get the default policy of `role:readonly`, which cannot modify Argo CD settings.

可以在 `argocd-rbac-cm` ConfigMap 中配置其他角色和组。下面的示例配置了一个名为“org-admin”的自定义角色。该角色被分配给属于“your-github-org:your-team”组的任何用户。所有其他用户获得“role:readonly”的默认策略，该策略无法修改 Argo CD 设置。

*ArgoCD ConfigMap `argocd-rbac-cm` Example:*

```
apiVersion: v1
kind: ConfigMap
metadata:
  name: argocd-rbac-cm
  namespace: argocd
data:
  policy.default: role:readonly
  policy.csv: |
    p, role:org-admin, applications, *, */*, allow
    p, role:org-admin, clusters, get, *, allow
    p, role:org-admin, repositories, get, *, allow
    p, role:org-admin, repositories, create, *, allow
    p, role:org-admin, repositories, update, *, allow
    p, role:org-admin, repositories, delete, *, allow

    g, your-github-org:your-team, role:org-admin
```


------

Another `policy.csv` example might look as follows:

```
p, role:staging-db-admins, applications, create, staging-db-admins/*, allow
p, role:staging-db-admins, applications, delete, staging-db-admins/*, allow
p, role:staging-db-admins, applications, get, staging-db-admins/*, allow
p, role:staging-db-admins, applications, override, staging-db-admins/*, allow
p, role:staging-db-admins, applications, sync, staging-db-admins/*, allow
p, role:staging-db-admins, applications, update, staging-db-admins/*, allow
p, role:staging-db-admins, projects, get, staging-db-admins, allow
g, db-admins, role:staging-db-admins
```


This example defines a *role* called `staging-db-admins` with *seven permissions* that allow that role to perform the *actions* (`create`/`delete`/`get`/`override`/`sync` /`update` applications, and `get` appprojects) against `*` (all) objects in the `staging-db-admins` Argo CD AppProject.

此示例定义了一个名为 `staging-db-admins` 的 *角色*，具有*七种权限*，允许该角色执行*操作*（`create`/`delete`/`get`/`override`/`sync` /`update` 应用程序和 `get` appprojects）针对 `staging-db-admins` Argo CD AppProject 中的 `*`（所有）对象。

## Anonymous Access[¶](https://argo-cd-docs.readthedocs.io/en/latest/operator-manual/rbac/#anonymous-access) 

## 匿名访问[¶](https://argo-cd-docs.readthedocs.io/en/latest/operator-manual/rbac/#anonymous-access)

The anonymous access to Argo CD can be enabled using `users.anonymous.enabled` field in `argocd-cm` (see [argocd-cm.yaml](https://argo-cd-docs.readthedocs.io/en/latest/operator-manual/argocd-cm.yaml)). The anonymous users get default role permissions specified by `policy.default` in `argocd-rbac-cm.yaml`. For read-only access you'll want `policy.default: role:readonly` as above 

可以使用 `argocd-cm` 中的 `users.anonymous.enabled` 字段启用对 Argo CD 的匿名访问（参见 [argocd-cm.yaml](https://argo-cd-docs.readthedocs.io/en/最新/操作员手册/argocd-cm.yaml）)。匿名用户获得由 `argocd-rbac-cm.yaml` 中的 `policy.default` 指定的默认角色权限。对于只读访问，您需要如上所述的`policy.default: role:readonly`


