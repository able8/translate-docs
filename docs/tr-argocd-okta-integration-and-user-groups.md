# ArgoCD: Okta integration, and user groups

# ArgoCD：Okta 集成和用户组

05/17/2021 From: https://rtfm.co.ua/en/argocd-okta-integration-and-user-groups/

The idea is that we don’t add user accounts locally in the ArgoCD’s ConfigMap, but instead will use our Okta users databases and Okta will perform their authentication. And on the ArgoCD side will do users’ authorization, i.e. will check their permission boundaries.

这个想法是我们不会在 ArgoCD 的 ConfigMap 本地添加用户帐户，而是使用我们的 Okta 用户数据库，Okta 将执行他们的身份验证。而在 ArgoCD 端会做用户的授权，即会检查他们的权限边界。

Also, by using the SSO we will be able to create user groups that will have various roles tied to specific Projects.

此外，通过使用 SSO，我们将能够创建具有与特定项目相关联的各种角色的用户组。

We will use the [SAML (with Dex)](https://argoproj.github.io/argo-cd/operator-manual/user-management/okta/#saml-with-dex), see also the [What is : SAML – an overview, its structure and requests tracing between a Jenkins and Okta SSO](https://rtfm.co.ua/en/what-is-saml-an-overview-its-structure-and-requests-tracing-between-a-jenkins-and-okta-sso/) post for more details on the SAML.

我们将使用 [SAML (with Dex)](https://argoproj.github.io/argo-cd/operator-manual/user-management/okta/#saml-with-dex)，另见 [什么是: SAML – Jenkins 和 Okta SSO 之间的概述、其结构和请求跟踪](https://rtfm.co.ua/en/what-is-saml-an-overview-its-structure-and-requests-tracing-between-a-jenkins-and-okta-sso/) 发布有关 SAML 的更多详细信息。

Contents

内容

- [Okta configuration](http://rtfm.co.ua#Okta_configuration "Okta configuration")
- [ArgoCD configuration](http://rtfm.co.ua#ArgoCD_configuration "ArgoCD configuration")
   - [Roles, and user groups in ArgoCD](http://rtfm.co.ua#Roles_and_user_groups_in_ArgoCD "Roles, and user groups in ArgoCD")
   - [ArgoCD CLI login](http://rtfm.co.ua#ArgoCD_CLI_login "ArgoCD CLI login")
- [Errors and problems](http://rtfm.co.ua#Errors_and_problems "Errors and problems")
   - [Bad Request – User session error](http://rtfm.co.ua#Bad_Request_-_User_session_error "Bad Request – User session error")
   - [Failed to authenticate: no attribute with name “group”: [email]](http://rtfm.co.ua#Failed_to_authenticate_no_attribute_with_name_group_email "Failed to authenticate: no attribute with name “group”: [email]")

- [Okta 配置](http://rtfm.co.ua#Okta_configuration "Okta 配置")
- [ArgoCD 配置](http://rtfm.co.ua#ArgoCD_configuration "ArgoCD 配置")
  - [ArgoCD 中的角色和用户组](http://rtfm.co.ua#Roles_and_user_groups_in_ArgoCD "ArgoCD 中的角色和用户组")
  - [ArgoCD CLI 登录](http://rtfm.co.ua#ArgoCD_CLI_login "ArgoCD CLI 登录")
- [错误和问题](http://rtfm.co.ua#Errors_and_problems"错误和问题")
  - [错误请求 – 用户会话错误](http://rtfm.co.ua#Bad_Request_-_User_session_error "错误请求 – 用户会话错误")
  - [验证失败：没有名称为“组”的属性：[email]](http://rtfm.co.ua#Failed_to_authenticate_no_attribute_with_name_group_email“验证失败：没有名称为“组”的属性：[email]”)

# Okta configuration

# Okta 配置

Go to your Okta account, and create a SAML application:

[![](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_124710.png)](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_124710.png)

转到您的 Okta 帐户，并创建一个 SAML 应用程序：

[![](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_124758.png)](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_124758.png)

For the logo, I’ve used this pic:

[![](https://rtfm.co.ua/wp-content/uploads/2021/05/argo-horizontal-color.png)](https://rtfm.co.ua/wp-content/uploads/2021/05/argo-horizontal-color.png) On the next page, set the following:

- _Single sign on URL_ and Audience URI (SP Entity ID)_:  your ArgoCD URL like _argourl.com/api/dex/callback_
- Attribute Statements_: name – _email_, format Basic, Value – _user.email_
- Group Attribute Statements_: name – _group_, format Basic, in the Filter you can use a regex as it's shown in the [documentation](https://argoproj.github.io/argo-cd/operator-manual/user-management/okta/#saml-with-dex) – then, Okta groups will be filtered before passing them to the ArgoCD during  [Authentication context](https://rtfm.co.ua/en/what-is-saml-an-overview-its-structure-and-requests-tracing-between-a-jenkins-and-okta-sso/#Authentication_context), or filter nothing by setting the regex as “\*”:

[![](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_132008.png)](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_132008.png)

对于徽标，我使用了这张图片：

在下一页，设置以下内容：

- 单点登录 URL 和 _Audience URI（SP 实体 ID）_：您的 ArgoCD URL，如 _argourl.com/api/dex/callback_
- Attribute Statements_：名称 - _email_，格式基本，值 - _user.email_
- Group Attribute Statements_：名称 – _group_，格式基本，在过滤器中，您可以使用 [文档](https://argoproj.github.io/argo-cd/operator-manual/user-management)中所示的正则表达式/okta/#saml-with-dex) – 然后，Okta 组将在 [身份验证上下文](https://rtfm.co.ua/en/what-is-saml-an-Overview-its-structure-and-requests-tracing-between-a-jenkins-and-okta-sso/#Authentication_context)，或者通过将正则表达式设置为“\*”来过滤任何内容：

Or you can filter a list of groups by using the `|`:

[![](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_133047.png)](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_133047.png)

或者，您可以使用 `|` 过滤组列表：

On the next page, set the _I’m an Okta customer adding an internal app_, and click Finish:

[![](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_125608.png)](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_125608.png)

在下一页上，设置_我是 Okta 客户添加内部应用程序_，然后单击完成：

Assign the Application to a user or group:

[![](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_125657.png)](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_125657.png)

将应用程序分配给用户或组：

Switch to the Sign On tab, and click on the _View Setup Instructions_ button:

[![](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_125737.png)](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_125737.png) 

Now, let's go to our ArgoCD instance.

切换到登录选项卡，然后单击_查看设置说明_按钮：

 现在，让我们转到我们的 ArgoCD 实例。

# ArgoCD configuration

# ArgoCD 配置

Edit the `argocd-cm` ConfigMap:

```
kubectl -n dev-1-18-devops-argocd-ns edit configmap argocd-cm
```

编辑 `argocd-cm` ConfigMap：

Set:

- _url_: (can be already set, just check its value) – ArgoCD’s instance external URL 

- _ssoURL_: Identity Provider Single Sign-On URL from the page that was opened after you’ve clicked the _View Setup Instructions_ button
- _caData_: an Okta’s certificate from the same page encoded with base64, for example on the [https://www.base64encode.org](https://www.base64encode.org/) site


Now, your config will look like this:

[![](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_130354.png)](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_130354.png)

现在，您的配置将如下所示：

```
apiVersion: v1
data:
url: https://dev-1-18.argocd.example.com
dex.config: |
    logger:
      level: debug
      format: json
    connectors:
    - type: saml
      id: okta
      name: Okta
      config:
        ssoURL: https://okta.example.com/app/appname/exk9f6o07dKjN0huC357/sso/saml
        caData: |
          LS0tLS***FLS0tLS0=
        usernameAttr: email
        emailAttr: email
        groupsAttr: group
...
```


## Roles, and user groups in ArgoCD

## ArgoCD 中的角色和用户组

Edit the `argocd-rbac-cm` ConfigMap, add the _DevOps_ group (will get from the Okta) mapping to the `role:admin`:

```
g, DevOps, role:admin
```


Also, here you can disable the _admin_ user, and add Backend and Web groups bindings to applications on their projects:

```
data:
admin.enabled: "false"
policy.csv: |
    p, role:backend-app-admin, applications, *, Backend/*, allow
    p, role:web-app-admin, applications, *, Web/*, allow
    g, DevOps, role:admin
policy.default: role:''
scopes: '[email,groups]'
```


Keep in mind, that RBAC is case sensitive, so DevOps != devops. So you have to specify a group here the same as it is set in the Okta.

请记住，RBAC 区分大小写，因此 DevOps != devops。因此，您必须在此处指定一个与 Okta 中设置相同的组。

Save, and try to log in:

[![](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_131242.png)](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_131242.png)

保存并尝试登录：

[![](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_150422-1.png)](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_150422-1.png)

## ArgoCD CLI login

## ArgoCD CLI 登录

To use CLI with SSO, just add the  `--sso` option, and in the `--username` – specify your Okta’s login:

要将 CLI 与 SSO 结合使用，只需添加 `--sso` 选项，并在 `--username` 中指定您的 Okta 登录名：


```
argocd login dev-1-18.argocd.example.com --sso --username arseniy@example.com
Opening browser for authentication
...
Authentication successful
'arseniy@example.com' logged in successfully
Context 'dev-1-18.argocd.example.com' updated
```

A default browser will be opened to authenticate you with Okta.

将打开默认浏览器以使用 Okta 对您进行身份验证。

# Errors and problems

# 错误和问题

## Bad Request – User session error

## 错误请求 – 用户会话错误

When using the [Dex](https://dexidp.io/docs/connectors/saml/) keep in mind, that it doesn't allow the Provider-initated login, ie you can log in to the ArgoCD instance from your Okta – you'll get the 400 error:

[![](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_135016.png)](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_135016.png)

使用 [Dex](https://dexidp.io/docs/connectors/saml/) 时请记住，它不允许提供者启动的登录，即您可以从 Okta 登录到 ArgoCD 实例– 你会得到 400 错误：

Thus, worth hiding the application from users:

[![](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_135122.png)](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_135122.png)

因此，值得对用户隐藏应用程序：

## Failed to authenticate: no attribute with name “group”: [email]

## 验证失败：没有名为“group”的属性：[email]

In case of error like “ **Failed to authenticate: no attribute with name “group”: [email]**“:

[![](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_131544.png)](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_131544.png)

如果出现“**无法验证：没有名称为“组”的属性：[email]**”之类的错误：

Use the [saml-chrome-panel](https://chrome.google.com/webstore/detail/saml-chrome-panel/paijfdbeoenhembfhkhllainmocckace), and check attributes that are sent in the `AttributeStatement` from the Okta:

[![](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_150903.png)](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_150903.png)

使用 [saml-chrome-panel](https://chrome.google.com/webstore/detail/saml-chrome-panel/paijfdbeoenhembfhkhllinmocckace)，并检查从 Okta 发送到 `AttributeStatement` 的属性：

```
<saml2:AttributeStatement xmlns:saml2="urn:oasis:names:tc:SAML:2.0:assertion">
    <saml2:Attribute Name="email" NameFormat="urn:oasis:names:tc:SAML:2.0:attrname-format:basic">
        <saml2:AttributeValue xmlns:xs="http://www.w3.org/2001/XMLSchema"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="xs:string">arseniy@example.com</saml2:AttributeValue>
    </saml2:Attribute>
</saml2:AttributeStatement>
```


In this example, I’ve made a misprint in the Okta config and it doesn’t send the `<saml2:Attribute Name="group"` attribute.

在这个例子中，我在 Okta 配置中打印错误，它没有发送 `<saml2:Attribute Name="group"` 属性。

The entry must be like the next:

```
<saml2:AttributeStatement xmlns:saml2="urn:oasis:names:tc:SAML:2.0:assertion">
    <saml2:Attribute Name="email" NameFormat="urn:oasis:names:tc:SAML:2.0:attrname-format:basic">
        <saml2:AttributeValue xmlns:xs="http://www.w3.org/2001/XMLSchema"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="xs:string">arseniy@example.com</saml2:AttributeValue>
    </saml2:Attribute>
    <saml2:Attribute Name="group" NameFormat="urn:oasis:names:tc:SAML:2.0:attrname-format:basic">
        <saml2:AttributeValue xmlns:xs="http://www.w3.org/2001/XMLSchema"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="xs:string">DevOps</saml2:AttributeValue>
    </saml2:Attribute>
</saml2:AttributeStatement>
```

Done. 

