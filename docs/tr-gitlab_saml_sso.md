# SAML SSO for GitLab.com groups

# GitLab.com 组的 SAML SSO

Introduced in GitLab 11.0.

在 GitLab 11.0 中引入。

From: https://docs.gitlab.com/ee/user/group/saml_sso/

This page describes SAML for Groups. For instance-wide SAML on self-managed GitLab instances, see [SAML OmniAuth Provider](http://docs.gitlab.com/../../../integration/saml.html).
[View the differences between SaaS and Self-Managed Authentication and Authorization Options](http://docs.gitlab.com/../../../administration/auth/index.html#saas-vs-self-managed-comparison).

本页面介绍了组的 SAML。有关自我管理的 GitLab 实例上的实例范围 SAML，请参阅 [SAML OmniAuth Provider](http://docs.gitlab.com/../../../integration/saml.html)。[查看 SaaS 和自管理身份验证和授权选项之间的差异](http://docs.gitlab.com/../../../administration/auth/index.html#saas-vs-self-managed-比较)。

SAML on GitLab.com allows users to sign in through their SAML identity provider. If the user is not already a member, the sign-in process automatically adds the user to the appropriate group.

GitLab.com 上的 SAML 允许用户通过他们的 SAML 身份提供者登录。如果用户还不是成员，登录过程会自动将用户添加到适当的组。

User synchronization of SAML SSO groups is supported through [SCIM](http://docs.gitlab.com/scim_setup.html). SCIM supports adding and removing users from the GitLab group automatically.
For example, if you remove a user from the SCIM app, SCIM removes that same user from the GitLab group.

通过 [SCIM](http://docs.gitlab.com/scim_setup.html) 支持 SAML SSO 组的用户同步。 SCIM 支持自动在 GitLab 组中添加和删除用户。
例如，如果您从 SCIM 应用程序中删除用户，SCIM 将从 GitLab 组中删除该用户。

SAML SSO is only configurable at the top-level group.

SAML SSO 只能在顶级组中配置。

If required, you can find [a glossary of common terms](http://docs.gitlab.com/../../../integration/saml.html#glossary-of-common-terms).

如果需要，您可以找到[常用术语词汇表](http://docs.gitlab.com/../../../integration/saml.html#glossary-of-common-terms)。

## Configuring your identity provider

## 配置您的身份提供者

1. Navigate to the GitLab group and select**Settings > SAML SSO**.

1. 导航到 GitLab 组并选择**设置 > SAML SSO**。

2. Configure your SAML identity provider using the**Assertion consumer service URL**, **Identifier**, and **GitLab single sign-on URL**.
Alternatively GitLab provides [metadata XML configuration](http://docs.gitlab.com#metadata-configuration).
See [specific identity provider documentation](http://docs.gitlab.com#providers) for more details.

2. 使用**Assertion 消费者服务 URL**、**Identifier** 和 **GitLab 单点登录 URL** 配置您的 SAML 身份提供者。
或者，GitLab 提供 [元数据 XML 配置](http://docs.gitlab.com#metadata-configuration)。
有关更多详细信息，请参阅 [特定身份提供者文档](http://docs.gitlab.com#providers)。

3. Configure the SAML response to include a NameID that uniquely identifies each user.

3. 配置 SAML 响应以包含唯一标识每个用户的 NameID。

4. Configure[required assertions](http://docs.gitlab.com#assertions) at minimum containing the user’s email address.

4. 配置[required assertions](http://docs.gitlab.com#assertions)至少包含用户的电子邮件地址。

5. While the default is enabled for most SAML providers, please ensure the app is set to have service provider
initiated calls in order to link existing GitLab accounts.

5. 虽然大多数 SAML 提供商默认启用，但请确保应用设置为具有服务提供商发起调用以链接现有的 GitLab 帐户。

6. Once the identity provider is set up, move on to[configuring GitLab](http://docs.gitlab.com#configuring-gitlab).


6. 设置身份提供者后，继续[配置 GitLab](http://docs.gitlab.com#configuring-gitlab)。


[![Issuer and callback for configuring SAML identity provider with GitLab.com](http://docs.gitlab.com/img/group_saml_configuration_information.png)](http://docs.gitlab.com/img/group_saml_configuration_information.png)

)

### NameID

### 名称ID

GitLab.com uses the SAML NameID to identify users. The NameID element:

GitLab.com 使用 SAML NameID 来识别用户。 NameID 元素：

- Is a required field in the SAML response.

- 是 SAML 响应中的必填字段。

- Must be unique to each user.

- 每个用户必须是唯一的。

- Must be a persistent value that will never change, such as a randomly generated unique user ID.

- 必须是永远不会改变的持久值，例如随机生成的唯一用户 ID。

- Is case sensitive. The NameID must match exactly on subsequent login attempts, so should not rely on user input that could change between upper and lower case.

- 区分大小写。 NameID 必须在随后的登录尝试中完全匹配，因此不应依赖可能在大小写之间变化的用户输入。

- Should not be an email address or username. We strongly recommend against these as it’s hard to
guarantee it doesn’t ever change, for example, when a person’s name changes. Email addresses are
also case-insensitive, which can result in users being unable to sign in.


- 不应是电子邮件地址或用户名。我们强烈建议不要使用这些，因为这很难
保证它永远不会改变，例如，当一个人的名字改变时。电子邮件地址是
也不区分大小写，这可能会导致用户无法登录。


The relevant field name and recommended value for supported providers are in the [provider specific notes](http://docs.gitlab.com#providers).
appropriate corresponding field.

[provider specific notes](http://docs.gitlab.com#providers) 中提供了支持提供者的相关字段名称和推荐值。
相应的字段。

cautionOnce users have signed into GitLab using the SSO SAML setup, changing the `NameID` breaks the configuration and potentially locks users out of the GitLab group.

注意一旦用户使用 SSO SAML 设置登录到 GitLab，更改“NameID”会破坏配置并可能将用户锁定在 GitLab 组之外。

#### NameID Format

#### NameID 格式

We recommend setting the NameID format to `Persistent` unless using a field (such as email) that requires a different format.

我们建议将 NameID 格式设置为“Persistent”，除非使用需要不同格式的字段（例如电子邮件）。

### Assertions

### 断言

For users to be created with the right information with the improved [user access and management](http://docs.gitlab.com#user-access-and-management),
the user details need to be passed to GitLab as SAML assertions.

为了通过改进的[用户访问和管理](http://docs.gitlab.com#user-access-and-management)创建具有正确信息的用户，用户详细信息需要作为 SAML 断言传递给 GitLab。

At a minimum, the user’s email address _must_ be specified as an assertion named `email` or `mail`.
See [the assertions list](http://docs.gitlab.com/../../../integration/saml.html#assertions) for other available claims.

至少，用户的电子邮件地址_必须_指定为名为“email”或“mail”的断言。
有关其他可用声明，请参阅 [断言列表](http://docs.gitlab.com/../../../integration/saml.html#assertions)。

noteThe `username` assertion is not supported for GitLab.com SaaS integrations.

注意 GitLab.com SaaS 集成不支持“用户名”断言。

### Metadata configuration

### 元数据配置

GitLab provides metadata XML that can be used to configure your identity provider.

GitLab 提供可用于配置身份提供者的元数据 XML。

1. Navigate to the group and select**Settings > SAML SSO**.

1. 导航到组并选择**设置 > SAML SSO**。

2. Copy the provided**GitLab metadata URL**.

2. 复制提供的**GitLab 元数据 URL**。

3. Follow your identity provider’s documentation and paste the metadata URL when it’s requested.


3. 按照您的身份提供商的文档并在请求时粘贴元数据 URL。


## Configuring GitLab 

## 配置 GitLab

After you set up your identity provider to work with GitLab, you must configure GitLab to use it for authentication:

设置身份提供程序以使用 GitLab 后，您必须配置 GitLab 以使用它进行身份验证：

1. Navigate to the group’s **Settings > SAML SSO**.

1. 导航到组的**设置 > SAML SSO**。

2. Find the SSO URL from your identity provider and enter it the**Identity provider single sign-on URL** field.

2. 从您的身份提供商处找到 SSO URL 并将其输入**身份提供商单点登录 URL** 字段。

3. Find and enter the fingerprint for the SAML token signing certificate in the**Certificate** field.

3. 在**证书**字段中查找并输入 SAML 令牌签名证书的指纹。

4. Select the access level to be applied to newly added users in the**Default membership role** field. The default access level is ‘Guest’.

4. 在**默认成员角色**字段中选择要应用于新添加用户的访问级别。默认访问级别为“访客”。

5. Select the**Enable SAML authentication for this group** checkbox.

5. 选中**为该组启用 SAML 身份验证**复选框。

6. Select the**Save changes** button.


6. 选择**保存更改**按钮。


[![Group SAML Settings for GitLab.com](http://docs.gitlab.com/img/group_saml_settings_v13_12.png)](http://docs.gitlab.com/img/group_saml_settings_v13_12.png)

note The certificate [fingerprint algorithm](http://docs.gitlab.com/../../../integration/saml.html#notes-on-configuring-your-identity-provider) must be in SHA1. When configuring the identity provider, use a secure signature algorithm.

注意证书[指纹算法](http://docs.gitlab.com/../../../integration/saml.html#notes-on-configuring-your-identity-provider)必须是SHA1。配置身份提供者时，请使用安全签名算法。

### SSO enforcement

### SSO 强制执行

Version history

版本历史

- [Introduced](https://gitlab.com/gitlab-org/gitlab/-/issues/5291) in GitLab 11.8.

- [介绍](https://gitlab.com/gitlab-org/gitlab/-/issues/5291) 在 GitLab 11.8 中。

- [Improved](https://gitlab.com/gitlab-org/gitlab/-/issues/9255) in GitLab 11.11 with ongoing enforcement in the GitLab UI.

- [改进](https://gitlab.com/gitlab-org/gitlab/-/issues/9255) 在 GitLab 11.11 中，在 GitLab UI 中持续执行。

- [Improved](https://gitlab.com/gitlab-org/gitlab/-/issues/292811) in GitLab 13.8, with an updated timeout experience.

- [改进](https://gitlab.com/gitlab-org/gitlab/-/issues/292811) 在 GitLab 13.8 中，更新了超时体验。

- [Improved](https://gitlab.com/gitlab-org/gitlab/-/issues/211962) in GitLab 13.8 with allowing group owners to not go through SSO.

- [改进](https://gitlab.com/gitlab-org/gitlab/-/issues/211962) 在 GitLab 13.8 中允许组所有者不通过 SSO。

- [Improved](https://gitlab.com/gitlab-org/gitlab/-/issues/9152) in GitLab 13.11 with enforcing open SSO session to use Git if this setting is switched on.


- [已改进](https://gitlab.com/gitlab-org/gitlab/-/issues/9152) 在 GitLab 13.11 中，如果启用此设置，则强制开放 SSO 会话使用 Git。


With this option enabled, users (except owners) must go through your group’s GitLab single sign-on URL if they wish to access group resources through the UI. Users can’t be manually added as members.

启用此选项后，用户（所有者除外）如果希望通过 UI 访问组资源，则必须通过您组的 GitLab 单点登录 URL。用户无法手动添加为成员。

SSO enforcement does not affect sign in or access to any resources outside of the group. Users can view which groups and projects they are a member of without SSO sign in.

SSO 实施不会影响登录或访问组外的任何资源。用户无需 SSO 登录即可查看他们所属的组和项目。

However, users are not prompted to sign in through SSO on each visit. GitLab checks whether a user
has authenticated through SSO. If it’s been more than 1 day since the last sign-in, GitLab
prompts the user to sign in again through SSO.

但是，系统不会在每次访问时提示用户通过 SSO 登录。 GitLab 检查用户是否
已通过 SSO 进行身份验证。如果距离上次登录已经超过 1 天，GitLab
提示用户通过 SSO 再次登录。

We intend to add a similar SSO requirement for [API activity](https://gitlab.com/gitlab-org/gitlab/-/issues/9152).

我们打算为 [API 活动](https://gitlab.com/gitlab-org/gitlab/-/issues/9152) 添加类似的 SSO 要求。

SSO has the following effects when enabled:

SSO 启用时具有以下效果：

- For groups, users can’t share a project in the group outside the top-level group,
even if the project is forked.

- 对于群组，用户不能在顶级群组之外的群组中共享项目，
即使项目是分叉的。

- For a Git activity, users must be signed-in through SSO before they can push to or
pull from a GitLab repository.

- 对于 Git 活动，用户必须先通过 SSO 登录，然后才能推送到或
从 GitLab 存储库中提取。

- Users must be signed-in through SSO before they can pull images using the[Dependency Proxy](http://docs.gitlab.com/../../packages/dependency_proxy/index.html).


- 用户必须通过 SSO 登录才能使用 [依赖代理](http://docs.gitlab.com/../../packages/dependency_proxy/index.html) 拉取图像。


When SSO is enforced, users are not immediately revoked. If the user:

强制执行 SSO 时，不会立即撤销用户。如果用户：

- Is signed out, they cannot access the group after being removed from the identity provider.

- 已注销，从身份提供者中删除后，他们无法访问该组。

- Has an active session, they can continue accessing the group for up to 24 hours until the identity
provider session times out.


- 有一个活跃的会话，他们可以继续访问该组长达 24 小时，直到身份
提供程序会话超时。


When SCIM updates, the user’s access is immediately revoked.

当 SCIM 更新时，用户的访问权限会立即被撤销。

## Providers

## 供应商

The SAML standard means that a wide range of identity providers will work with GitLab. Your identity provider may have relevant documentation. It may be generic SAML documentation, or specifically targeted for GitLab.

SAML 标准意味着广泛的身份提供者将与 GitLab 合作。您的身份提供者可能有相关文件。它可能是通用的 SAML 文档，或者专门针对 GitLab。

When [configuring your identity provider](http://docs.gitlab.com#configuring-your-identity-provider), please consider the notes below for specific providers to help avoid common issues and as a guide for terminology used.

[配置您的身份提供者](http://docs.gitlab.com#configuring-your-identity-provider) 时，请考虑以下特定提供者的注意事项，以帮助避免常见问题并作为所用术语的指南。

For providers not listed below, you can refer to the [instance SAML notes on configuring an identity provider](http://docs.gitlab.com/../../../integration/saml.html#notes-on-configuring-your-identity-provider)
for additional guidance on information your identity provider may require.

对于下面未列出的提供者，您可以参考[关于配置身份提供者的实例 SAML 说明](http://docs.gitlab.com/../../../integration/saml.html#notes-on-configuring-your-identity-provider)
有关您的身份提供者可能需要的信息的其他指导。

GitLab provides the following information for guidance only.
If you have any questions on configuring the SAML app, please contact your provider’s support.

GitLab 提供以下信息仅供参考。
如果您对配置 SAML 应用程序有任何疑问，请联系您的提供商的支持人员。

### Azure setup notes 

### Azure 设置说明

Follow the Azure documentation on [configuring single sign-on to applications](https://docs.microsoft.com/en-us/azure/active-directory/manage-apps/view-applications-portal) with the notes below for consideration.

遵循有关 [配置应用程序单点登录](https://docs.microsoft.com/en-us/azure/active-directory/manage-apps/view-applications-portal) 的 Azure 文档以及以下说明考虑。

For a demo of the Azure SAML setup including SCIM, see [SCIM Provisioning on Azure Using SAML SSO for Groups Demo](https://youtu.be/24-ZxmTeEBU). The video is outdated in regard to
objectID mapping and the [SCIM documentation should be followed](http://docs.gitlab.com/scim_setup.html#azure-configuration-steps).

有关包括 SCIM 在内的 Azure SAML 设置的演示，请参阅 [使用 SAML SSO 在 Azure 上进行组演示的 SCIM 配置](https://youtu.be/24-ZxmTeEBU)。该视频已过时
objectID 映射和 [应遵循 SCIM 文档](http://docs.gitlab.com/scim_setup.html#azure-configuration-steps)。

GitLab Setting
Azure Field
Identifier
Identifier (Entity ID)
Assertion consumer service URL
Reply URL (Assertion Consumer Service URL)
GitLab single sign-on URL
Sign on URL
Identity provider single sign-on URL
Login URL
Certificate fingerprint
Thumbprint

We recommend:

我们推荐：

- **Unique User Identifier (Name identifier)** set to `user.objectID`.

- **唯一用户标识符（名称标识符）** 设置为 `user.objectID`。

- **nameid-format** set to persistent.


- **nameid-format** 设置为持久化。


If using [Group Sync](http://docs.gitlab.com#group-sync), customize the name of the group claim to match the required attribute.

如果使用 [Group Sync](http://docs.gitlab.com#group-sync)，请自定义组声明的名称以匹配所需的属性。

See the [troubleshooting page](http://docs.gitlab.com/../../../administration/troubleshooting/group_saml_scim.html#azure-active-directory) for an example configuration.

有关示例配置，请参阅 [疑难解答页面](http://docs.gitlab.com/../../../administration/troubleshooting/group_saml_scim.html#azure-active-directory)。

### Okta setup notes

### Okta 安装说明

Please follow the Okta documentation on [setting up a SAML application in Okta](https://developer.okta.com/docs/guides/build-sso-integration/saml2/overview/) with the notes below for consideration.

请按照 [在 Okta 中设置 SAML 应用程序](https://developer.okta.com/docs/guides/build-sso-integration/saml2/overview/) 的 Okta 文档以及以下注意事项进行考虑。

For a demo of the Okta SAML setup including SCIM, see [Demo: Okta Group SAML & SCIM setup](https://youtu.be/0ES9HsZq0AQ).

有关包含 SCIM 的 Okta SAML 设置的演示，请参阅 [演示：Okta Group SAML 和 SCIM 设置](https://youtu.be/0ES9HsZq0AQ)。



Under Okta’s **Single sign-on URL** field, check the option **Use this for Recipient URL and Destination URL**.

在 Okta 的 **Single sign-on URL** 字段下，选中选项 **Use this for Recipient URL and Destination URL**。

For NameID, the following settings are recommended; for SCIM, the following settings are required:

对于NameID，推荐如下设置；对于 SCIM，需要以下设置：

- **Application username** (NameID) set to **Custom** `user.getInternalProperty("id")`.

- **应用程序用户名** (NameID) 设置为**自定义** `user.getInternalProperty("id")`。

- **Name ID Format** set to **Persistent**.


- **名称 ID 格式** 设置为 **持久**。


### OneLogin setup notes

### OneLogin 设置说明

OneLogin supports their own [GitLab (SaaS)](https://onelogin.service-now.com/support?id=kb_article&sys_id=92e4160adbf16cd0ca1c400e0b961923&kb_category=50984e84db738300d5505eea4b961913)
application.

OneLogin 支持自己的 [GitLab (SaaS)](https://onelogin.service-now.com/support?id=kb_article&sys_id=92e4160adbf16cd0ca1c400e0b961923&kb_category=50984e84db738300d4b36051)
应用。

If you decide to use the OneLogin generic [SAML Test Connector (Advanced)](https://onelogin.service-now.com/support?id=kb_article&sys_id=b2c19353dbde7b8024c780c74b9619fb&kb_category=93e869b0db185340d5505eea4b961934),
we recommend the [“Use the OneLogin SAML Test Connector” documentation](https://onelogin.service-now.com/support?id=kb_article&sys_id=93f95543db109700d5505eea4b96198f) with the following settings:

如果您决定使用 OneLogin 通用 [SAML 测试连接器（高级）](https://onelogin.service-now.com/support?id=kb_article&sys_id=b2c19353dbde7b8024c780c74b9619fb&kb_category=93e869b0edb14b596350edb14b96350edb14b50
我们建议使用 [“使用 OneLogin SAML 测试连接器”文档](https://onelogin.service-now.com/support?id=kb_article&sys_id=93f95543db109700d5505eea4b96198f) 进行以下设置：



Recommended `NameID` value: `OneLogin ID`.

推荐的`NameID` 值：`OneLogin ID`。

## User access and management

## 用户访问和管理

> [Improved](https://gitlab.com/gitlab-org/gitlab/-/issues/268142) in GitLab 13.7.

> [改进](https://gitlab.com/gitlab-org/gitlab/-/issues/268142) 在 GitLab 13.7 中。

Once Group SSO is configured and enabled, users can access the GitLab.com group through the identity provider’s dashboard. If [SCIM](http://docs.gitlab.com/scim_setup.html) is configured, please see the [user access and linking setup section on the SCIM page](http://docs.gitlab.com/scim_setup.html#user-access-and-linking-setup).

配置并启用组 SSO 后，用户可以通过身份提供者的仪表板访问 GitLab.com 组。如果配置了[SCIM](http://docs.gitlab.com/scim_setup.html)，请参见[SCIM页面上的用户访问和链接设置部分](http://docs.gitlab.com/scim_setup.html)。html#user-access-and-linking-setup)。

When a user tries to sign in with Group SSO, GitLab attempts to find or create a user based on the following:

当用户尝试使用组 SSO 登录时，GitLab 尝试根据以下内容查找或创建用户：

- Find an existing user with a matching SAML identity. This would mean the user either had their account created by[SCIM](http://docs.gitlab.com/scim_setup.html) or they have previously signed in with the group’s SAML IdP.

- 查找具有匹配 SAML 身份的现有用户。这意味着用户的帐户要么是由 [SCIM](http://docs.gitlab.com/scim_setup.html) 创建的，要么他们之前已使用该组的 SAML IdP 登录。

- If there is no conflicting user with the same email address, create a new account automatically.

- 如果没有具有相同电子邮件地址的冲突用户，则自动创建一个新帐户。

- If there is a conflicting user with the same email address, redirect the user to the sign-in page to:
   - Create a new account with another email address.

   - 如果存在具有相同电子邮件地址的冲突用户，请将用户重定向到登录页面以：
  - 使用另一个电子邮件地址创建一个新帐户。

  - Sign-in to their existing account to link the SAML identity.

- 登录到他们现有的帐户以链接 SAML 身份。

### Linking SAML to your existing GitLab.com account

### 将 SAML 链接到您现有的 GitLab.com 帐户

To link SAML to your existing GitLab.com account:

要将 SAML 链接到您现有的 GitLab.com 帐户：

1. Sign in to your GitLab.com account. 

1. 登录您的 GitLab.com 帐户。

2. Locate and visit the**GitLab single sign-on URL** for the group you’re signing in to. A group owner can find this on the group’s **Settings > SAML SSO** page. If the sign-in URL is configured, users can connect to the GitLab app from the identity provider.

2. 找到并访问您要登录的群组的**GitLab 单点登录 URL**。群组所有者可以在群组的 **设置 > SAML SSO** 页面上找到它。如果配置了登录 URL，用户可以从身份提供者连接到 GitLab 应用程序。

3. Select**Authorize**.

3. 选择**授权**。

4. Enter your credentials on the identity provider if prompted.

4. 如果出现提示，请在身份提供者上输入您的凭据。

5. You are then redirected back to GitLab.com and should now have access to the group. In the future, you can use SAML to sign in to GitLab.com.


5. 然后您将被重定向回 GitLab.com，现在应该可以访问该组。将来，您可以使用 SAML 登录 GitLab.com。


On subsequent visits, you should be able to go [sign in to GitLab.com with SAML](http://docs.gitlab.com#signing-in-to-gitlabcom-with-saml) or by visiting links directly. If the **enforce SSO** option is turned on, you are then redirected to sign in through the identity provider.

在后续访问中，您应该能够[使用 SAML 登录 GitLab.com](http://docs.gitlab.com#signing-in-to-gitlabcom-with-saml) 或直接访问链接。如果启用了 **enforce SSO** 选项，您将被重定向以通过身份提供商登录。

### Signing in to GitLab.com with SAML

### 使用 SAML 登录 GitLab.com

1. Sign in to your identity provider.

1. 登录到您的身份提供商。

2. From the list of apps, select the “GitLab.com” app. (The name is set by the administrator of the identity provider.)

2. 从应用程序列表中，选择“GitLab.com”应用程序。 （名称由身份提供者的管理员设置。）

3. You are then signed in to GitLab.com and redirected to the group.


3. 然后您登录到 GitLab.com 并重定向到该组。


### Configure user settings from SAML response

### 从 SAML 响应配置用户设置

[Introduced](https://gitlab.com/gitlab-org/gitlab/-/issues/263661) in GitLab 13.7.

[介绍](https://gitlab.com/gitlab-org/gitlab/-/issues/263661) 在 GitLab 13.7 中。

GitLab allows setting certain user attributes based on values from the SAML response.
This affects users created on first sign-in via Group SAML. Existing users’
attributes are not affected regardless of the values sent in the SAML response.

GitLab 允许根据来自 SAML 响应的值设置某些用户属性。
这会影响通过 Group SAML 首次登录时创建的用户。现有用户的
无论 SAML 响应中发送的值如何，属性都不会受到影响。

#### Supported user attributes

#### 支持的用户属性

- `can_create_group` \- ‘true’ or ‘false’ to indicate whether the user can create
new groups. Default is `true`.

- `can_create_group` \- 'true' 或 'false' 表示用户是否可以创建
新组。默认值为“真”。

- `projects_limit` \- The total number of personal projects a user can create.
A value of `0` means the user cannot create new projects in their personal
namespace. Default is `10000`.


- `projects_limit` \- 用户可以创建的个人项目总数。
“0”值表示用户不能在他们的个人中创建新项目
命名空间。默认值为“10000”。


#### Example SAML response

#### SAML 响应示例

You can find SAML responses in the developer tools or console of your browser,
in base64-encoded format. Use the base64 decoding tool of your choice to
convert the information to XML. An example SAML response is shown here.

您可以在浏览器的开发者工具或控制台中找到 SAML 响应，
以 base64 编码格式。使用您选择的 base64 解码工具
将信息转换为 XML。此处显示了示例 SAML 响应。

```
    <saml2:AttributeStatement>
      <saml2:Attribute Name="email" NameFormat="urn:oasis:names:tc:SAML:2.0:attrname-format:basic">
         <saml2:AttributeValue xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="xs:string">user.email</saml2:AttributeValue>
      </saml2:Attribute>
      <saml2:Attribute Name="first_name" NameFormat="urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified">
         <saml2:AttributeValue xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="xs:string">user.firstName</saml2:AttributeValue>
      </saml2:Attribute>
      <saml2:Attribute Name="last_name" NameFormat="urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified">
         <saml2:AttributeValue xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="xs:string">user.lastName</saml2:AttributeValue>
      </saml2:Attribute>
      <saml2:Attribute Name="can_create_group" NameFormat="urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified">
         <saml2:AttributeValue xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="xs:string">true</saml2:AttributeValue>
      </saml2:Attribute>
      <saml2:Attribute Name="projects_limit" NameFormat="urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified">
         <saml2:AttributeValue xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="xs:string">10</saml2:AttributeValue>
      </saml2:Attribute>
   </saml2:AttributeStatement>
```


### Role

###  角色

Starting from [GitLab 13.3](https://gitlab.com/gitlab-org/gitlab/-/issues/214523), group owners can set a ‘Default membership role’ other than ‘Guest’. To do so, [configure the SAML SSO for the group](http://docs.gitlab.com#configuring-gitlab). That role becomes the starting access level of all users added to the group.

从 [GitLab 13.3](https://gitlab.com/gitlab-org/gitlab/-/issues/214523) 开始，群组所有者可以设置除“Guest”之外的“默认成员角色”。为此，[为组配置 SAML SSO](http://docs.gitlab.com#configuring-gitlab)。该角色成为添加到该组的所有用户的起始访问级别。

Existing members with appropriate privileges can promote or demote users, as needed.

具有适当权限的现有成员可以根据需要提升或降级用户。

If a user is already a member of the group, linking the SAML identity does not change their role.

如果用户已经是该组的成员，则链接 SAML 身份不会改变他们的角色。

### Blocking access

### 阻止访问

Please refer to [Blocking access via SCIM](http://docs.gitlab.com/scim_setup.html#blocking-access).

请参考[通过SCIM阻止访问](http://docs.gitlab.com/scim_setup.html#blocking-access)。

### Unlinking accounts

### 取消关联帐户

Users can unlink SAML for a group from their profile page. This can be helpful if: 

用户可以从他们的个人资料页面取消组的 SAML 链接。如果出现以下情况，这可能会有所帮助：

- You no longer want a group to be able to sign you in to GitLab.com.

- 您不再希望某个群组能够让您登录 GitLab.com。

- Your SAML NameID has changed and so GitLab can no longer find your user.


- 您的 SAML NameID 已更改，因此 GitLab 无法再找到您的用户。


cautionUnlinking an account removes all roles assigned to that user in the group.
If a user re-links their account, roles need to be reassigned.

注意取消帐户链接会删除组中分配给该用户的所有角色。
如果用户重新关联其帐户，则需要重新分配角色。

Groups require at least one owner. If your account is the only owner in the
group, you are not allowed to unlink the account. In that case, set up another user as a
group owner, and then you can unlink the account.

组需要至少一位所有者。如果您的帐户是该帐户中的唯一所有者
组，您不能取消链接帐户。在这种情况下，将另一个用户设置为
群组所有者，然后您可以取消关联帐户。

For example, to unlink the `MyOrg` account:

例如，要取消链接`MyOrg` 帐户：

1. In the top-right corner, select your avatar.

1. 在右上角，选择您的头像。

2. Select**Edit profile**.

2. 选择**编辑个人资料**。

3. On the left sidebar, select**Account**.

3. 在左侧边栏上，选择**帐户**。

4. In the**Social sign-in** section, select **Disconnect** next to the connected account.


4. 在**社交登录**部分，选择已连接帐户旁边的**断开连接**。


[![Unlink Group SAML](http://docs.gitlab.com/img/unlink_group_saml.png)](http://docs.gitlab.com/img/unlink_group_saml.png)

## Group Sync

## 组同步

For a demo of Group Sync using Azure, see [Demo: SAML Group Sync](https://youtu.be/Iqvo2tJfXjg).

有关使用 Azure 进行组同步的演示，请参阅 [演示：SAML 组同步](https://youtu.be/Iqvo2tJfXjg)。

When the SAML response includes a user and their group memberships from the SAML identity provider,
GitLab uses that information to automatically manage that user’s GitLab group memberships.

当 SAML 响应包含来自 SAML 身份提供商的用户及其组成员身份时，
GitLab 使用该信息自动管理该用户的 GitLab 组成员资格。

Ensure your SAML identity provider sends an attribute statement named `Groups` or `groups` like the following:

确保您的 SAML 身份提供商发送名为“Groups”或“groups”的属性语句，如下所示：

```
<saml:AttributeStatement>
  <saml:Attribute Name="Groups">
    <saml:AttributeValue xsi:type="xs:string">Developers</saml:AttributeValue>
    <saml:AttributeValue xsi:type="xs:string">Product Managers</saml:AttributeValue>
  </saml:Attribute>
</saml:AttributeStatement>

```


noteTo inspect the SAML response, you can use one of these [SAML debugging tools](http://docs.gitlab.com#saml-debugging-tools).
Also note that the value for `Groups` or `groups` in the SAML reponse can be either the group name or
the group ID depending what the IdP sends to GitLab.

注意要检查 SAML 响应，您可以使用这些 [SAML 调试工具](http://docs.gitlab.com#saml-debugging-tools) 之一。
另请注意，SAML 响应中“Groups”或“groups”的值可以是组名或
组 ID 取决于 IdP 发送给 GitLab 的内容。

When SAML SSO is enabled for the top-level group, `Maintainer` and `Owner` level users
see a new menu item in group **Settings > SAML Group Links**. You can configure one or more **SAML Group Links** to map
a SAML identity provider group name to a GitLab Access Level. This can be done for the parent group or the subgroups.

为顶级组启用 SAML SSO 后，`Maintainer` 和 `Owner` 级别用户
在 **Settings > SAML Group Links** 组中看到一个新菜单项。您可以配置一个或多个 **SAML Group Links** 来映射
一个 SAML 身份提供者组名称到 GitLab 访问级别。这可以为父组或子组完成。

To link the SAML groups from the `saml:AttributeStatement` example above:

要从上面的“saml:AttributeStatement”示例中链接 SAML 组：

1. Enter the value of`saml:AttributeValue` in the `SAML Group Name` field.

1. 在“SAML Group Name”字段中输入“saml:AttributeValue”的值。

2. Choose the desired`Access Level`.

2. 选择所需的`访问级别`。

3. **Save** the group link.

3. **保存**群链接。

4. Repeat to add additional group links if desired.


4. 如果需要，重复添加其他组链接。


[![SAML Group Links](http://docs.gitlab.com/img/saml_group_links_v13_9.png)](http://docs.gitlab.com/img/saml_group_links_v13_9.png)

If a user is a member of multiple SAML groups mapped to the same GitLab group,
the user gets the highest access level from the groups. For example, if one group
is linked as `Guest` and another `Maintainer`, a user in both groups gets `Maintainer`
access.

如果用户是映射到同一个 GitLab 组的多个 SAML 组的成员，
用户从组中获得最高访问级别。例如，如果一组
被链接为“Guest”和另一个“Maintainer”，两个组中的用户都获得“Maintainer”
使用权。

### Automatic member removal

### 自动删除成员

After a group sync, users who are not members of a mapped SAML group are removed from
the GitLab group.

组同步后，不是映射 SAML 组成员的用户将从
GitLab 小组。

For example, in the following diagram:

例如，在下图中：

- Alex Garcia signs into GitLab and is removed from GitLab Group C because they don’t belong
to SAML Group C.

- Alex Garcia 登录 GitLab 并从 GitLab Group C 中删除，因为它们不属于
到 SAML 组 C。

- Sidney Jones belongs to SAML Group C, but is not added to GitLab Group C because they have
not yet signed in.


- Sidney Jones 属于 SAML 组 C，但未添加到 GitLab 组 C，因为他们有
尚未登录。


## Passwords for users created via SAML SSO for Groups

## 通过 SAML SSO 为群组创建的用户密码

The [Generated passwords for users created through integrated authentication](http://docs.gitlab.com/../../../security/passwords_for_integrated_authentication_methods.html) guide provides an overview of how GitLab generates and sets passwords for users created via SAML SSO for Groups.

[为通过集成身份验证创建的用户生成密码](http://docs.gitlab.com/../../../security/passwords_for_integrated_authentication_methods.html) 指南概述了 GitLab 如何为用户生成和设置密码通过 SAML SSO 为组创建。

## Troubleshooting

##  故障排除

This section contains possible solutions for problems you might encounter.

本节包含您可能遇到的问题的可能解决方案。

### SAML debugging tools

### SAML 调试工具

SAML responses are base64 encoded, so we recommend the following browser plugins to decode them on the fly:

SAML 响应采用 base64 编码，因此我们建议使用以下浏览器插件即时解码它们：

- [SAML tracer for Firefox](https://addons.mozilla.org/en-US/firefox/addon/saml-tracer/)
- [Chrome SAML Panel](https://chrome.google.com/webstore/detail/saml-chrome-panel/paijfdbeoenhembfhkhllainmocckace?hl=en)

- [Firefox 的 SAML 追踪器](https://addons.mozilla.org/en-US/firefox/addon/saml-tracer/)
- [Chrome SAML 面板](https://chrome.google.com/webstore/detail/saml-chrome-panel/paijfdbeoenhembfhkhllinmocckace?hl=en)

Specific attention should be paid to:

应特别注意：

- The[NameID](http://docs.gitlab.com#nameid), which we use to identify which user is signing in. If the user has previously signed in, this [must match the value we have stored](http://docs.gitlab.com#verifying-nameid).

- [NameID](http://docs.gitlab.com#nameid)，我们用它来标识哪个用户正在登录。如果用户以前登录过，这[必须与我们存储的值匹配](http://docs.gitlab.com#verifying-nameid)。

- The presence of a`X509Certificate`, which we require to verify the response signature.

- 存在`X509Certificate`，我们需要它来验证响应签名。

- The`SubjectConfirmation` and `Conditions`, which can cause errors if misconfigured.


- `SubjectConfirmation` 和 `Conditions`，如果配置错误可能会导致错误。


### Verifying configuration 

### 验证配置

For convenience, we’ve included some [example resources](http://docs.gitlab.com/../../../administration/troubleshooting/group_saml_scim.html) used by our Support Team. While they may help you verify the SAML app configuration, they are not guaranteed to reflect the current state of third-party products.

为方便起见，我们包含了一些由我们的支持团队使用的 [示例资源](http://docs.gitlab.com/../../../administration/troubleshooting/group_saml_scim.html)。虽然它们可以帮助您验证 SAML 应用程序配置，但不能保证它们反映第三方产品的当前状态。

### Verifying NameID

### 验证 NameID

In troubleshooting the Group SAML setup, any authenticated user can use the API to verify the NameID GitLab already has linked to the user by visiting [`https://gitlab.com/api/v4/user`](https://gitlab.com/api/v4/user) and checking the `extern_uid` under identities.

在对 Group SAML 设置进行故障排除时，任何经过身份验证的用户都可以通过访问 [`https://gitlab.com/api/v4/user`](https://gitlab.com/api/v4/user) 并检查身份下的“extern_uid”。

Similarly, group members of a role with the appropriate permissions can make use of the [members API](http://docs.gitlab.com/../../../api/members.html) to view group SAML identity information for members of the group.

同样，具有适当权限的角色的组成员可以使用 [members API](http://docs.gitlab.com/../../../api/members.html) 查看组 SAML组成员的身份信息。

This can then be compared to the [NameID](http://docs.gitlab.com#nameid) being sent by the identity provider by decoding the message with a [SAML debugging tool](http://docs.gitlab.com#saml-debugging-tools). We require that these match in order to identify users.

然后可以将其与身份提供者发送的 [NameID](http://docs.gitlab.com#nameid) 进行比较，使用 [SAML 调试工具](http://docs.gitlab.com#saml-调试工具)。我们要求这些匹配才能识别用户。

### Users receive a 404

### 用户收到 404

If a user is trying to sign in for the first time and the GitLab single sign-on URL has not [been configured](http://docs.gitlab.com#configuring-your-identity-provider), they may see a 404.
As outlined in the [user access section](http://docs.gitlab.com#linking-saml-to-your-existing-gitlabcom-account), a group Owner will need to provide the URL to users.

如果用户是第一次尝试登录并且 GitLab 单点登录 URL 尚未 [已配置](http://docs.gitlab.com#configuring-your-identity-provider)，他们可能会看到404.
如 [用户访问部分](http://docs.gitlab.com#linking-saml-to-your-existing-gitlabcom-account) 所述，组所有者需要向用户提供 URL。

### Message: “SAML authentication failed: Extern UID has already been taken”

### 消息：“SAML 身份验证失败：Extern UID 已被占用”

This error suggests you are signed in as a GitLab user but have already linked your SAML identity to a different GitLab user. Sign out and then try to sign in again using the SSO SAML link, which should log you into GitLab with the linked user account.

此错误表明您以 GitLab 用户身份登录，但已将您的 SAML 身份链接到其他 GitLab 用户。注销然后尝试使用 SSO SAML 链接再次登录，该链接应该使用链接的用户帐户将您登录到 GitLab。

If you do not wish to use that GitLab user with the SAML login, you can [unlink the GitLab account from the group’s SAML](http://docs.gitlab.com#unlinking-accounts).

如果您不希望通过 SAML 登录使用该 GitLab 用户，您可以[从组的 SAML 中取消关联 GitLab 帐户](http://docs.gitlab.com#unlinking-accounts)。

### Message: “SAML authentication failed: User has already been taken”

### 消息：“SAML 身份验证失败：用户已被占用”

The user that you’re signed in with already has SAML linked to a different identity.
Here are possible causes and solutions:

您登录的用户已将 SAML 链接到不同的身份。
以下是可能的原因和解决方案：

Cause
Solution
You’ve tried to link multiple SAML identities to the same user, for a given identity provider.
Change the identity that you sign in with. To do so, [unlink the previous SAML identity](http://docs.gitlab.com#unlinking-accounts) from this GitLab account before attempting to sign in again.

原因
解决方案
对于给定的身份提供者，您尝试将多个 SAML 身份链接到同一用户。
更改您登录时使用的身份。为此，在尝试再次登录之前，[取消链接以前的 SAML 身份](http://docs.gitlab.com#unlinking-accounts) 从此 GitLab 帐户。

### Message: “SAML authentication failed: Email has already been taken”

### 消息：“SAML 身份验证失败：电子邮件已被占用”

Cause
Solution
When a user account with the email address already exists in GitLab, but the user does not have the SAML identity tied to their account.
The user will need to [link their account](http://docs.gitlab.com#user-access-and-management).

原因
解决方案
当 GitLab 中已存在具有电子邮件地址的用户帐户，但该用户没有绑定到其帐户的 SAML 身份时。
用户需要[链接他们的帐户](http://docs.gitlab.com#user-access-and-management)。

### Message: “SAML authentication failed: Extern UID has already been taken, User has already been taken”

### 消息：“SAML 身份验证失败：Extern UID 已经被占用，用户已经被占用”

Getting both of these errors at the same time suggests the NameID capitalization provided by the identity provider didn’t exactly match the previous value for that user.

同时出现这两个错误表明身份提供者提供的 NameID 大写与该用户的先前值不完全匹配。

This can be prevented by configuring the [NameID](http://docs.gitlab.com#nameid) to return a consistent value. Fixing this for an individual user involves [unlinking SAML in the GitLab account](http://docs.gitlab.com#unlinking-accounts), although this will cause group membership and to-dos to be lost.

这可以通过配置 [NameID](http://docs.gitlab.com#nameid) 以返回一致的值来防止。为单个用户修复此问题涉及 [在 GitLab 帐户中取消链接 SAML](http://docs.gitlab.com#unlinking-accounts)，尽管这会导致组成员身份和待办事项丢失。

### Message: “Request to link SAML account must be authorized”

### 消息：“必须授权关联 SAML 帐户的请求”

Ensure that the user who is trying to link their GitLab account has been added as a user within the identity provider’s SAML app.

确保尝试链接其 GitLab 帐户的用户已添加为身份提供商的 SAML 应用程序中的用户。

Alternatively, the SAML response may be missing the `InResponseTo` attribute in the
`samlp:Response` tag, which is [expected by the SAML gem](https://github.com/onelogin/ruby-saml/blob/9f710c5028b069bfab4b9e2b66891e0549765af5/lib/onelogin/ruby-saml/response.rb#L307-L316).
The identity provider administrator should ensure that the login is
initiated by the service provider and not the identity provider.

或者，SAML 响应中可能缺少“InResponseTo”属性
`samlp:Response` 标签，这是 [SAML gem 预期的](https://github.com/onelogin/ruby-saml/blob/9f710c5028b069bfab4b9e2b66891e0549765af5/lib/onelogin/ruby-saml/response.rb-L310)。
身份提供者管理员应确保登录是
由服务提供者而非身份提供者发起。

### Stuck in a login “loop”

### 陷入登录“循环”

Ensure that the **GitLab single sign-on URL** has been configured as “Login URL” (or similarly named field) in the identity provider’s SAML app. 

确保 **GitLab 单点登录 URL** 已在身份提供商的 SAML 应用程序中配置为“登录 URL”（或类似命名的字段）。

Alternatively, when users need to [link SAML to their existing GitLab.com account](http://docs.gitlab.com#linking-saml-to-your-existing-gitlabcom-account), provide the **GitLab single sign -on URL** and instruct users not to use the SAML app on first sign in.

或者，当用户需要[将 SAML 链接到他们现有的 GitLab.com 帐户](http://docs.gitlab.com#linking-saml-to-your-existing-gitlabcom-account) 时，提供 **GitLab 单点登录-on URL** 并指示用户在首次登录时不要使用 SAML 应用程序。

### The NameID has changed

### NameID 已更改

Cause
Solution
As mentioned in the [NameID](http://docs.gitlab.com#nameid) section, if the NameID changes for any user, the user can be locked out. This is a common problem when an email address is used as the identifier.
Follow the steps outlined in the [“SAML authentication failed: User has already been taken”](http://docs.gitlab.com#message-saml-authentication-failed-user-has-already-been-taken) section.

原因
解决方案
如 [NameID](http://docs.gitlab.com#nameid) 部分所述，如果任何用户的 NameID 发生更改，则该用户可能会被锁定。当使用电子邮件地址作为标识符时，这是一个常见问题。
按照 [“SAML 身份验证失败：用户已被使用”](http://docs.gitlab.com#message-saml-authentication-failed-user-has-already-been-taken) 部分中概述的步骤进行操作。

### I need to change my SAML app

### 我需要更改我的 SAML 应用程序

If the NameID is identical in both SAML apps, then no change is required.

如果两个 SAML 应用程序中的 NameID 相同，则不需要更改。

Otherwise, to change the SAML app used for sign in, users need to [unlink the current SAML identity](http://docs.gitlab.com#unlinking-accounts) and then [link their identity](http://docs.gitlab.com#user-access-and-management) to the new SAML app.

否则，要更改用于登录的 SAML 应用程序，用户需要[取消链接当前的 SAML 身份](http://docs.gitlab.com#unlinking-accounts) 然后[链接他们的身份](http://docs.gitlab.com#user-access-and-management) 到新的 SAML 应用程序。

### I need additional information to configure my identity provider

### 我需要其他信息来配置我的身份提供者

Many SAML terms can vary between providers. It is possible that the information you are looking for is listed under another name.

许多 SAML 条款可能因提供商而异。您正在查找的信息可能以其他名称列出。

For more information, start with your identity provider’s documentation. Look for their options and examples to see how they configure SAML. This can provide hints on what you’ll need to configure GitLab to work with these providers.

有关更多信息，请从您的身份提供商的文档开始。查找他们的选项和示例以了解他们如何配置 SAML。这可以提供有关配置 GitLab 以与这些提供程序一起工作所需的提示。

It can also help to look at our [more detailed docs for self-managed GitLab](http://docs.gitlab.com/../../../integration/saml.html).
SAML configuration for GitLab.com is mostly the same as for self-managed instances.
However, self-managed GitLab instances use a configuration file that supports more options as described in the external [OmniAuth SAML documentation](https://github.com/omniauth/omniauth-saml/).
Internally that uses the [`ruby-saml` library](https://github.com/onelogin/ruby-saml), so we sometimes check there to verify low level details of less commonly used options.

它还可以帮助查看我们的[自我管理 GitLab 的更详细文档](http://docs.gitlab.com/../../../integration/saml.html)。
GitLab.com 的 SAML 配置与自我管理实例的配置大致相同。
但是，自我管理的 GitLab 实例使用的配置文件支持更多选项，如外部 [OmniAuth SAML 文档](https://github.com/omniauth/omniauth-saml/) 中所述。
在内部使用 [`ruby-saml` 库](https://github.com/onelogin/ruby-saml)，所以我们有时会检查那里以验证不常用选项的低级细节。

It can also help to compare the XML response from your provider with our [example XML used for internal testing](https://gitlab.com/gitlab-org/gitlab/-/blob/master/ee/spec/fixtures/saml/response.xml).

它还可以帮助将来自您的提供商的 XML 响应与我们的 [用于内部测试的示例 XML](https://gitlab.com/gitlab-org/gitlab/-/blob/master/ee/spec/fixtures/saml/response.xml)。

### Searching Rails log

### 搜索 Rails 日志

With access to the rails log or `production_json.log` (available only to GitLab team members for GitLab.com),
you should be able to find the base64 encoded SAML response by searching with the following filters:

可以访问 rails 日志或 `production_json.log`（仅适用于 GitLab.com 的 GitLab 团队成员），
通过使用以下过滤器进行搜索，您应该能够找到 base64 编码的 SAML 响应：

- `json.meta.caller_id`: `Groups::OmniauthCallbacksController#group_saml`
- `json.meta.user` or `json.username`: `username`
- `json.method`: `POST`
- `json.path`: `/groups/GROUP-PATH/-/saml/callback`

- `json.meta.caller_id`：`Groups::OmniauthCallbacksController#group_saml`
- `json.meta.user` 或 `json.username`: `username`
-`json.method`：`POST`
- `json.path`: `/groups/GROUP-PATH/-/saml/callback`

In a relevant log entry, the `json.params` should provide a valid response with:

在相关的日志条目中，`json.params` 应该提供一个有效的响应：

- `"key": "SAMLResponse"` and the `"value": (full SAML response)`,

- `"key": "SAMLResponse"` 和 `"value": (full SAML response)`,

- `"key": "RelayState"` with `"value": "/group-path"`, and

- `"key": "RelayState"` 和 `"value": "/group-path"`，以及

- `"key": "group_id"` with `"value": "group-path"`.


- `"key": "group_id"` 带有 `"value": "group-path"`。


In some cases, if the SAML response is lengthy, you may receive a `"key": "truncated"` with `"value":"..."`.
In these cases, please ask a group owner for a copy of the SAML response from when they select
the “Verify SAML Configuration” button on the group SSO Settings page.

在某些情况下，如果 SAML 响应很长，您可能会收到一个 `"key": "truncated"` 和 `"value":"..."`。
在这些情况下，请向群组所有者索取他们选择时的 SAML 响应副本
组 SSO 设置页面上的“验证 SAML 配置”按钮。

Use a base64 decoder to see a human-readable version of the SAML response. 

使用 base64 解码器查看 SAML 响应的人类可读版本。

