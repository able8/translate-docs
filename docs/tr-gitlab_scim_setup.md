# SCIM provisioning using SAML SSO for GitLab.com groups

# 使用 SAML SSO 为 GitLab.com 组配置 SCIM

[Introduced](https://gitlab.com/gitlab-org/gitlab/-/merge_requests/9388) in GitLab Premium 11.10.

[介绍](https://gitlab.com/gitlab-org/gitlab/-/merge_requests/9388) 在 GitLab Premium 11.10 中。

From: https://docs.gitlab.com/ee/user/group/saml_sso/scim_setup.html

System for Cross-domain Identity Management (SCIM), is an open standard that enables the
automation of user provisioning. When SCIM is provisioned for a GitLab group, membership of
that group is synchronized between GitLab and the identity provider.

跨域身份管理系统 (SCIM) 是一个开放标准，它使用户配置自动化。当为 GitLab 组配置 SCIM 时，
该组在 GitLab 和身份提供者之间同步。

The GitLab [SCIM API](http://docs.gitlab.com/../../../api/scim.html) implements part of [the RFC7644 protocol](https://tools.ietf.org/html/rfc7644).

GitLab [SCIM API](http://docs.gitlab.com/../../../api/scim.html) 实现了 [RFC7644 协议](https://tools.ietf.org)的一部分/html/rfc7644)。

## Features

##  特征

The following actions are available:
- Create users
- Deactivate users


以下操作可用：
- 创建用户
- 停用用户


The following identity providers are supported:
- Azure
- Okta


支持以下身份提供者：
- 天蓝色
- Okta


## Requirements

##  要求

- [Group Single Sign-On](http://docs.gitlab.com/index.html) must be configured.


- 必须配置[组单点登录](http://docs.gitlab.com/index.html)。


## GitLab configuration

## GitLab 配置

Once [Group Single Sign-On](http://docs.gitlab.com/index.html) has been configured, we can:

[Group Single Sign-On](http://docs.gitlab.com/index.html) 配置完成后，我们可以：

1. Navigate to the group and click**Administration > SAML SSO**.
2. Click on the**Generate a SCIM token** button.
3. Save the token and URL so they can be used in the next step.


1. 导航到该组并单击**管理 > SAML SSO**。
2. 单击**生成 SCIM 令牌** 按钮。
3. 保存令牌和 URL，以便在下一步中使用它们。


[![SCIM token configuration](http://docs.gitlab.com/img/scim_token_v13_3.png)](http://docs.gitlab.com/img/scim_token_v13_3.png)

## Identity Provider configuration

## 身份提供者配置

- [Azure](http://docs.gitlab.com#azure-configuration-steps)
- [Okta](http://docs.gitlab.com#okta-configuration-steps)

- [Azure](http://docs.gitlab.com#azure-configuration-steps)
- [Okta](http://docs.gitlab.com#okta-configuration-steps)

### Azure configuration steps

### Azure 配置步骤

The SAML application that was created during [Single sign-on](http://docs.gitlab.com/index.html) setup for [Azure](https://docs.microsoft.com/en-us/azure/active-directory/manage-apps/view-applications-portal) now needs to be set up for SCIM.

在 [单点登录](http://docs.gitlab.com/index.html) 设置 [Azure](https://docs.microsoft.com/en-us/azure/) 期间创建的 SAML 应用程序active-directory/manage-apps/view-applications-portal) 现在需要为 SCIM 设置。

1. Set up automatic provisioning and administrative credentials by following the
[Azure's SCIM setup documentation](https://docs.microsoft.com/en-us/azure/active-directory/app-provisioning/use-scim-to-provision-users-and-groups#provisioning-users-and-groups-to-applications-that-support-scim).


1. 按照以下步骤设置自动配置和管理凭据
[Azure 的 SCIM 设置文档](https://docs.microsoft.com/en-us/azure/active-directory/app-provisioning/use-scim-to-provision-users-and-groups#provisioning-users-and-groups-to-applications-that-support-scim)。


During this configuration, note the following:

在此配置过程中，请注意以下事项：

- The`Tenant URL` and `secret token` are the ones retrieved in the
[previous step](http://docs.gitlab.com#gitlab-configuration).

- `Tenant URL` 和 `secret token` 是在
[上一步](http://docs.gitlab.com#gitlab-configuration)。

- It is recommended to set a notification email and check the**Send an email notification when a failure occurs** checkbox.

- 建议设置通知邮件并勾选**发生故障时发送邮件通知**复选框。

- For mappings, we only leave`Synchronize Azure Active Directory Users to AppName` enabled.
`Synchronize Azure Active Directory Groups to AppName` is usually disabled. However, this
does not mean Azure AD users cannot be provisioned in groups. Leaving it enabled does not break
the SCIM user provisioning, but causes errors in Azure AD that may be confusing and misleading.


- 对于映射，我们只启用`Synchronize Azure Active Directory Users to AppName`。
“将 Azure Active Directory 组同步到 AppName”通常被禁用。然而，这
并不意味着不能按组配置 Azure AD 用户。启用它不会中断
SCIM 用户配置，但会导致 Azure AD 中的错误，这些错误可能会造成混淆和误导。


You can then test the connection by clicking on **Test Connection**. If the connection is successful, be sure to save your configuration before moving on. See below for [troubleshooting](http://docs.gitlab.com#troubleshooting).

然后，您可以通过单击 **Test Connection** 来测试连接。如果连接成功，请确保在继续之前保存您的配置。请参阅下面的 [疑难解答](http://docs.gitlab.com#troubleshooting)。

#### Configure attribute mapping

#### 配置属性映射

Follow [Azure documentation to configure the attribute mapping](https://docs.microsoft.com/en-us/azure/active-directory/app-provisioning/customize-application-attributes).

按照 [Azure 文档配置属性映射](https://docs.microsoft.com/en-us/azure/active-directory/app-provisioning/customize-application-attributes)。

The following table below provides an attribute mapping known to work with GitLab. If
your SAML configuration differs from [the recommended SAML settings](http://docs.gitlab.com/index.html#azure-setup-notes),
modify the corresponding `customappsso` settings accordingly. If a mapping is not listed in the
table, use the Azure defaults. For a list of required attributes, refer to the [SCIM API documentation](http://docs.gitlab.com/../../../api/scim.html).

下表提供了已知可用于 GitLab 的属性映射。如果
您的 SAML 配置不同于 [推荐的 SAML 设置](http://docs.gitlab.com/index.html#azure-setup-notes)，
相应地修改相应的“customappsso”设置。如果映射未列在
表，使用 Azure 默认值。有关必需属性的列表，请参阅 [SCIM API 文档](http://docs.gitlab.com/../../../api/scim.html)。

Azure Active Directory Attribute
`customappsso` Attribute
Matching precedence
`objectId``externalId`1
`userPrincipalName``emails[type eq "work"].value``mailNickname``userName`

Azure Active Directory 属性
`customappsso` 属性
匹配优先级
`objectId`externalId`1
`userPrincipalName``emails[type eq "work"].value``mailNickname``userName`

For guidance, you can view [an example configuration in the troubleshooting reference](http://docs.gitlab.com/../../../administration/troubleshooting/group_saml_scim.html#azure-active-directory).

如需指导，您可以查看[故障排除参考中的示例配置](http://docs.gitlab.com/../../../administration/troubleshooting/group_saml_scim.html#azure-active-directory)。

1. Below the mapping list click on**Show advanced options > Edit attribute list for AppName**.

1. 在映射列表下方单击**显示高级选项 > 编辑 AppName 的属性列表**。

2. Ensure the `id` is the primary and required field, and `externalId` is also required.




2. 确保 `id` 是主要和必填字段，并且 `externalId` 也是必需的。




note `username` should neither be primary nor required as we don’t support
that field on GitLab SCIM yet.

注意 `username` 既不应该是主要的也不应该是必需的，因为我们不支持
GitLab SCIM 上的那个领域呢。

3. Save all changes.

3. 保存所有更改。

4. In the **Provisioning** step, set the `Provisioning Status` to `On`.




4. 在 **Provisioning** 步骤中，将 `Provisioning Status` 设置为 `On`。




noteYou can control what is actually synced by selecting the `Scope`. For example, 

注意您可以通过选择“范围”来控制实际同步的内容。例如，

`Sync only assigned users and groups` only syncs the users assigned to
the application ( `Users and groups`), otherwise, it syncs the whole Active Directory.


`仅同步分配的用户和组`仅同步分配给的用户
应用程序（“用户和组”），否则，它会同步整个 Active Directory。


Once enabled, the synchronization details and any errors appears on the
bottom of the **Provisioning** screen, together with a link to the audit events.

启用后，同步详细信息和任何错误都会出现在
**Provisioning** 屏幕底部，以及指向审计事件的链接。

cautionOnce synchronized, changing the field mapped to `id` and `externalId` may cause a number of errors. These include provisioning errors, duplicate users, and may prevent existing users from accessing the GitLab group.

注意一旦同步，更改映射到 `id` 和 `externalId` 的字段可能会导致许多错误。这些包括配置错误、重复用户，并可能阻止现有用户访问 GitLab 组。

### Okta configuration steps

### Okta 配置步骤

Before you start this section, complete the [GitLab configuration](http://docs.gitlab.com#gitlab-configuration) process.
Make sure that you’ve also set up a SAML application for [Okta](https://developer.okta.com/docs/guides/build-sso-integration/saml2/overview/),
as described in the [Okta setup notes](http://docs.gitlab.com/index.html#okta-setup-notes)

在开始本节之前，请先完成 [GitLab 配置](http://docs.gitlab.com#gitlab-configuration) 流程。
确保您还为 [Okta] 设置了 SAML 应用程序（https://developer.okta.com/docs/guides/build-sso-integration/saml2/overview/），
如 [Okta 设置说明](http://docs.gitlab.com/index.html#okta-setup-notes) 中所述

Make sure that the Okta setup matches our documentation exactly, especially the NameID
configuration. Otherwise, the Okta SCIM app may not work properly.

确保 Okta 设置与我们的文档完全匹配，尤其是 NameID
配置。否则，Okta SCIM 应用程序可能无法正常工作。

01. Sign in to Okta.

01. 登录 Okta。

02. If you see an **Admin** button in the top right, click the button. This will
     ensure you are in the Admin area.




     02. 如果您在右上角看到 **Admin** 按钮，请单击该按钮。这会
    确保您在管理区域。




    noteIf you’re using the Developer Console, click **Developer Console** in the top
     bar and select **Classic UI**. Otherwise, you may not see the buttons described
     in the following steps:

注意如果您使用的是 Developer Console，请单击顶部的 **Developer Console**
    栏并选择 **Classic UI**。否则，您可能看不到描述的按钮
    在以下步骤中：

03. In the**Application** tab, click **Add Application**.

03. 在**应用程序** 选项卡中，单击**添加应用程序**。

04. Search for**GitLab**, find and click on the ‘GitLab’ application.

04. 搜索**GitLab**，找到并点击‘GitLab’应用程序。

05. On the GitLab application overview page, click**Add**.

05. 在 GitLab 应用程序概览页面上，单击**添加**。

06. Under**Application Visibility** select both checkboxes. Currently the GitLab application does not support SAML authentication so the icon should not be shown to users.

06. 在**Application Visibility** 下选中两个复选框。目前，GitLab 应用程序不支持 SAML 身份验证，因此不应向用户显示该图标。

07. Click**Done** to finish adding the application.

07. 单击**完成**以完成添加应用程序。

08. In the**Provisioning** tab, click **Configure API integration**.

08. 在**Provisioning** 选项卡中，单击**Configure API integration**。

09. Select**Enable API integration**.

     09. 选择**启用 API 集成**。

    - For**Base URL** enter the URL obtained from the GitLab SCIM configuration page

     - 对于**Base URL** 输入从 GitLab SCIM 配置页面获取的 URL

    - For**API Token** enter the SCIM token obtained from the GitLab SCIM configuration page
10. Click ‘Test API Credentials’ to verify configuration.

- 对于**API Token** 输入从 GitLab SCIM 配置页面获取的 SCIM 令牌
10. 单击“Test API Credentials”以验证配置。

11. Click**Save** to apply the settings.

11. 单击**保存**以应用设置。

12. After saving the API integration details, new settings tabs appear on the left. Choose**To App**.

12. 保存 API 集成详细信息后，左侧会出现新的设置选项卡。选择**到应用**。

13. Click**Edit**.

13. 单击**编辑**。

14. Check the box to**Enable** for both **Create Users** and **Deactivate Users**.

14. 为 **Create Users** 和 **Deactivate Users** 选中复选框以 **Enable**。

15. Click**Save**.

15. 单击**保存**。

16. Assign users in the**Assignments** tab. Assigned users are created and
     managed in your GitLab group.


16. 在**Assignments** 选项卡中分配用户。已创建分配的用户并
    在您的 GitLab 组中管理。


#### Okta Known Issues

#### Okta 已知问题

The Okta GitLab application currently only supports SCIM. Continue
using the separate Okta [SAML SSO](http://docs.gitlab.com/index.html) configuration along with the new SCIM
application described above.

Okta GitLab 应用程序目前仅支持 SCIM。继续
使用单独的 Okta [SAML SSO](http://docs.gitlab.com/index.html) 配置以及新的 SCIM
应用如上所述。

### OneLogin

### 登录

OneLogin provides a “GitLab (SaaS)” app in their catalog, which includes a SCIM integration.
As the app is developed by OneLogin, please reach out to OneLogin if you encounter issues.

OneLogin 在其目录中提供了一个“GitLab (SaaS)”应用程序，其中包括 SCIM 集成。
由于该应用程序由 OneLogin 开发，如果您遇到问题，请联系 OneLogin。

## User access and linking setup

## 用户访问和链接设置

During the synchronization process, all of your users get GitLab accounts, welcoming them
to their respective groups, with an invitation email. When implementing SCIM provisioning,
you may want to warn your security-conscious employees about this email.

在同步过程中，您的所有用户都会获得 GitLab 帐户，欢迎他们
到他们各自的组，用邀请电子邮件。在实施 SCIM 配置时，
您可能想就这封电子邮件警告您的具有安全意识的员工。

The following diagram is a general outline on what happens when you add users to your SCIM app:

下图概述了将用户添加到 SCIM 应用程序时会发生的情况：

graph TD
A[Add User to SCIM app] -->\|IdP sends user info to GitLab\| B(GitLab: Does the email exists?)
B -->\|No\| C[GitLab creates user with SCIM identity]
B -->\|Yes\| D[GitLab sends message back 'Email exists']

图TD
A[Add User to SCIM app] -->\|IdP 发送用户信息到 GitLab\| B（GitLab：电子邮件是否存在？）
B -->\|否\| C[GitLab 创建具有 SCIM 身份的用户]
B -->\|是\| D[GitLab 发回消息'电子邮件存在']

During provisioning:

在配置期间：

- Both primary and secondary emails are considered when checking whether a GitLab user account exists.

- 检查 GitLab 用户帐户是否存在时，会同时考虑主要和次要电子邮件。

- Duplicate usernames are also handled, by adding suffix`1` upon user creation. For example,
due to already existing `test_user` username, `test_user1` is used.


- 通过在用户创建时添加后缀`1`，也可以处理重复的用户名。例如，
由于已经存在 `test_user` 用户名，使用 `test_user1`。


As long as [Group SAML](http://docs.gitlab.com/index.html) has been configured, existing GitLab.com users can link to their accounts in one of the following ways:

只要配置了[Group SAML](http://docs.gitlab.com/index.html)，现有的GitLab.com用户就可以通过以下方式之一链接到他们的账户：

- By updating their_primary_ email address in their GitLab.com user account to match their identity provider’s user profile email address.

- 通过更新他们的 GitLab.com 用户帐户中的 their_primary_ 电子邮件地址以匹配其身份提供者的用户配置文件电子邮件地址。

- By following these steps:

- 按照以下步骤操作：

1. Sign in to GitLab.com if needed.

1. 如果需要，请登录 GitLab.com。

2. Click on the GitLab app in the identity provider’s dashboard or visit the**GitLab single sign-on URL**.

2. 单击身份提供者仪表板中的 GitLab 应用程序或访问**GitLab 单点登录 URL**。

3. Click on the**Authorize** button. 

3. 单击**授权**按钮。

We recommend users do this prior to turning on sync, because while synchronization is active, there may be provisioning errors for existing users.

我们建议用户在打开同步之前执行此操作，因为当同步处于活动状态时，现有用户可能会出现配置错误。

New users and existing users on subsequent visits can access the group through the identify provider’s dashboard or by visiting links directly.

后续访问的新用户和现有用户可以通过身份提供商的仪表板或直接访问链接访问该组。

[In GitLab 14.0 and later](https://gitlab.com/gitlab-org/gitlab/-/issues/325712), GitLab users created with a SCIM identity display with an **Enterprise** badge in the **Members ** view.

[在 GitLab 14.0 及更高版本中](https://gitlab.com/gitlab-org/gitlab/-/issues/325712)，使用 SCIM 身份显示创建的 GitLab 用户在 **Members 中带有 **Enterprise** 徽章** 看法。

[![Enterprise badge for users created with a SCIM identity](http://docs.gitlab.com/img/member_enterprise_badge_v14_0.png)](http://docs.gitlab.com/img/member_enterprise_badge_v14_0.png)

For role information, please see the [Group SAML page](http://docs.gitlab.com/index.html#user-access-and-management)

角色信息请查看 [Group SAML页面](http://docs.gitlab.com/index.html#user-access-and-management)

### Blocking access

### 阻止访问

To rescind access to the top-level group, all sub-groups, and projects, remove or deactivate the user
on the identity provider. SCIM providers generally update GitLab with the changes on demand, which
is minutes at most. The user’s membership is revoked and they immediately lose access.

要撤销对顶级组、所有子组和项目的访问权限，请删除或停用用户
在身份提供者上。 SCIM 提供商通常会根据需要更新 GitLab，从而
最多是分钟。用户的会员资格被撤销，他们立即失去访问权限。

noteDeprovisioning does not delete the GitLab user account.

noteDeprovisioning 不会删除 GitLab 用户帐户。

graph TD
A[Remove User from SCIM app] -->\|IdP sends request to GitLab\| B(GitLab: Is the user part of the group?)
B -->\|No\| C[Nothing to do]
B -->\|Yes\| D[GitLab removes user from GitLab group]

图TD
A[Remove User from SCIM app] -->\|IdP 向 GitLab 发送请求\| B（GitLab：用户是组的一部分吗？）
B -->\|否\| C[无事可做]
B -->\|是\| D[GitLab 从 GitLab 组中删除用户]

## Troubleshooting

##  故障排除

This section contains possible solutions for problems you might encounter.

本节包含您可能遇到的问题的可能解决方案。

### How come I can’t add a user after I removed them?

### 为什么我删除用户后无法添加用户？

As outlined in the [Blocking access section](http://docs.gitlab.com#blocking-access), when you remove a user, they are removed from the group. However, their account is not deleted.

如 [阻止访问部分](http://docs.gitlab.com#blocking-access) 中所述，当您删除用户时，他们将从组中删除。但是，他们的帐户不会被删除。

When the user is added back to the SCIM app, GitLab cannot create a new user because the user already exists.

当用户重新添加到 SCIM 应用程序时，GitLab 无法创建新用户，因为该用户已存在。

Solution: Have a user sign in directly to GitLab, then [manually link](http://docs.gitlab.com#user-access-and-linking-setup) their account.

解决方案：让用户直接登录 GitLab，然后 [手动链接](http://docs.gitlab.com#user-access-and-linking-setup) 他们的帐户。

### How do I diagnose why a user is unable to sign in

### 如何诊断用户无法登录的原因

Ensure that the user has been added to the SCIM app.

确保用户已添加到 SCIM 应用程序。

If you receive “User is not linked to a SAML account”, then most likely the user already exists in GitLab. Have the user follow the [User access and linking setup](http://docs.gitlab.com#user-access-and-linking-setup) instructions.

如果您收到“用户未链接到 SAML 帐户”，则很可能该用户已存在于 GitLab 中。让用户按照 [用户访问和链接设置](http://docs.gitlab.com#user-access-and-linking-setup) 说明进行操作。

The **Identity** ( `extern_uid`) value stored by GitLab is updated by SCIM whenever `id` or `externalId` changes. Users cannot sign in unless the GitLab Identity ( `extern_uid`) value matches the `NameId` sent by SAML.

GitLab 存储的 **Identity**（`extern_uid`）值在 `id` 或 `externalId` 发生变化时由 SCIM 更新。除非 GitLab 身份 (`extern_uid`) 值与 SAML 发送的 `NameId` 匹配，否则用户无法登录。

This value is also used by SCIM to match users on the `id`, and is updated by SCIM whenever the `id` or `externalId` values change.

SCIM 还使用此值来匹配“id”上的用户，并且每当“id”或“externalId”值发生变化时，SCIM 都会更新该值。

It is important that this SCIM `id` and SCIM `externalId` are configured to the same value as the SAML `NameId`. SAML responses can be traced using [debugging tools](http://docs.gitlab.com/index.html#saml-debugging-tools), and any errors can be checked against our [SAML troubleshooting docs](http://docs.gitlab.com/index.html#troubleshooting).

重要的是此 SCIM `id` 和 SCIM `externalId` 被配置为与 SAML `NameId` 相同的值。可以使用 [调试工具](http://docs.gitlab.com/index.html#saml-debugging-tools) 跟踪 SAML 响应，并且可以根据我们的 [SAML 故障排除文档](http://docs.gitlab.com/index.html#troubleshooting)。

### How do I verify user’s SAML NameId matches the SCIM externalId

### 如何验证用户的 SAML NameId 与 SCIM externalId 匹配

Group owners can see the list of users and the `externalId` stored for each user in the group SAML SSO Settings page.

组所有者可以在组 SAML SSO 设置页面中查看用户列表和为每个用户存储的“externalId”。

A possible alternative is to use the [SCIM API](http://docs.gitlab.com/../../../api/scim.html#get-a-list-of-scim-provisioned-users) to manually retrieve the `externalId` we have stored for users, also called the `external_uid` or `NameId`.

一种可能的替代方法是使用 [SCIM API](http://docs.gitlab.com/../../../api/scim.html#get-a-list-of-scim-provisioned-users) 手动检索我们为用户存储的“externalId”，也称为“external_uid”或“NameId”。

To see how the `external_uid` compares to the value returned as the SAML NameId, you can have the user use a [SAML Tracer](http://docs.gitlab.com/index.html#saml-debugging-tools).

要查看 `external_uid` 如何与作为 SAML NameId 返回的值进行比较，您可以让用户使用 [SAML Tracer](http://docs.gitlab.com/index.html#saml-debugging-tools)。

### Update or fix mismatched SCIM externalId and SAML NameId

### 更新或修复不匹配的 SCIM externalId 和 SAML NameId

Whether the value was changed or you need to map to a different field, ensure `id`, `externalId`, and `NameId` all map to the same field.

无论是更改值还是需要映射到不同的字段，请确保 `id`、`externalId` 和 `NameId` 都映射到相同的字段。

If the GitLab `externalId` doesn't match the SAML NameId, it needs to be updated in order for the user to sign in. Ideally your identity provider is configured to do such an update, but in some cases it may be unable to do so, such as when looking up a user fails due to an ID change.

如果 GitLab `externalId` 与 SAML NameId 不匹配，则需要更新它才能让用户登录。理想情况下，您的身份提供者配置为执行此类更新，但在某些情况下可能无法执行所以，例如当查找用户因 ID 更改而失败时。

Be cautious if you revise the fields used by your SCIM identity provider, typically `id` and `externalId`.
We use these IDs to look up users. If the identity provider does not know the current values for these fields,
that provider may create duplicate users.

如果您修改 SCIM 身份提供者使用的字段，通常是“id”和“externalId”，请务必小心。
我们使用这些 ID 来查找用户。如果身份提供者不知道这些字段的当前值，
该提供商可能会创建重复的用户。

If the `externalId` for a user is not correct, and also doesn’t match the SAML NameID,
you can address the problem in the following ways: 

如果用户的 `externalId` 不正确，并且与 SAML NameID 不匹配，
您可以通过以下方式解决问题：

- You can have users unlink and relink themselves, based on the[“SAML authentication failed: User has already been taken”](http://docs.gitlab.com/index.html#message-saml-authentication-failed-user-has-already-been-taken) section.

- 您可以让用户根据[“SAML 身份验证失败：用户已经被占用”](http://docs.gitlab.com/index.html#message-saml-authentication-failed-user-已经采取)部分。

- You can unlink all users simultaneously, by removing all users from the SAML app while provisioning is turned on.

- 您可以通过在启用配置时从 SAML 应用程序中删除所有用户来同时取消所有用户的链接。

- It may be possible to use the[SCIM API](http://docs.gitlab.com/../../../api/scim.html#update-a-single-scim-provisioned-user) to manually correct the `externalId` stored for users to match the SAML `NameId`.
To look up a user, you need to know the desired value that matches the `NameId` as well as the current `externalId`.


- 可以使用[SCIM API](http://docs.gitlab.com/../../../api/scim.html#update-a-single-scim-provisioned-user)手动更正为用户存储的 `externalId` 以匹配 SAML `NameId`。
要查找用户，您需要知道与“NameId”以及当前“externalId”相匹配的所需值。


It is important not to update these to incorrect values, since this causes users to be unable to sign in. It is also important not to assign a value to the wrong user, as this causes users to get signed into the wrong account.

重要的是不要将这些更新为不正确的值，因为这会导致用户无法登录。同样重要的是不要将值分配给错误的用户，因为这会导致用户登录到错误的帐户。

### I need to change my SCIM app

### 我需要更改我的 SCIM 应用程序

Individual users can follow the instructions in the [“SAML authentication failed: User has already been taken”](http://docs.gitlab.com/index.html#i-need-to-change-my-saml-app) section.

个人用户可以按照[“SAML身份验证失败：用户已经被占用”](http://docs.gitlab.com/index.html#i-need-to-change-my-saml-app)中的说明进行操作部分。

Alternatively, users can be removed from the SCIM app which de-links all removed users. Sync can then be turned on for the new SCIM app to [link existing users](http://docs.gitlab.com#user-access-and-linking-setup).

或者，可以从 SCIM 应用程序中删除用户，该应用程序会取消所有已删除用户的链接。然后可以为新的 SCIM 应用程序打开同步以[链接现有用户](http://docs.gitlab.com#user-access-and-linking-setup)。

### The SCIM app is throwing `"User has already been taken","status":409` error message

### SCIM 应用程序抛出“用户已经被占用”、“状态”：409` 错误消息

Changing the SAML or SCIM configuration or provider can cause the following problems:

更改 SAML 或 SCIM 配置或提供程序可能会导致以下问题：

Problem
Solution
SAML and SCIM identity mismatch.
First [verify that the user's SAML NameId matches the SCIM externalId](http://docs.gitlab.com#how-do-i-verify-users-saml-nameid-matches-the-scim-externalid) and then [update or fix the mismatched SCIM externalId and SAML NameId](http://docs.gitlab.com#update-or-fix-mismatched-scim-externalid-and-saml-nameid).
SCIM identity mismatch between GitLab and the Identify Provider SCIM app.
You can confirm whether you're hitting the error because of your SCIM identity mismatch between your SCIM app and GitLab.com by using [SCIM API](http://docs.gitlab.com/../../../api/scim.html#update-a-single-scim-provisioned-user) which shows up in the `id` key and compares it with the user `externalId` in the SCIM app. You can use the same [SCIM API](http://docs.gitlab.com/../../../api/scim.html#update-a-single-scim-provisioned-user) to update the SCIM `id` for the user on GitLab.com.

问题
解决方案
SAML 和 SCIM 身份不匹配。
首先[验证用户的 SAML NameId 是否与 SCIM externalId 匹配](http://docs.gitlab.com#how-do-i-verify-users-saml-nameid-matches-the-scim-externalid) 然后 [更新或修复不匹配的 SCIM externalId 和 SAML NameId](http://docs.gitlab.com#update-or-fix-mismatched-scim-externalid-and-saml-nameid)。
GitLab 和识别提供者 SCIM 应用程序之间的 SCIM 身份不匹配。
您可以使用 [SCIM API](http://docs.gitlab.com/../../../api/scim.html#update-a-single-scim-provisioned-user) 显示在 `id` 键中，并将其与 SCIM 应用程序中的用户 `externalId` 进行比较。您可以使用相同的 [SCIM API](http://docs.gitlab.com/../../../api/scim.html#update-a-single-scim-provisioned-user) 来更新GitLab.com 上用户的 SCIM `id`。

### Azure

### 蔚蓝

#### How do I verify my SCIM configuration is correct?

#### 如何验证我的 SCIM 配置是否正确？

Review the following:

查看以下内容：

- Ensure that the SCIM value for`id` matches the SAML value for `NameId`.

- 确保`id` 的SCIM 值与`NameId` 的SAML 值匹配。

- Ensure that the SCIM value for`externalId` matches the SAML value for `NameId`.


- 确保`externalId` 的SCIM 值与`NameId` 的SAML 值匹配。


Review the following SCIM parameters for sensible values:

查看以下 SCIM 参数以获取合理值：

- `userName`
- `displayName`
- `emails[type eq "work"].value`

- `用户名`
-`显示名称`
- `emails[type eq "work"].value`

#### Testing Azure connection: invalid credentials

#### 测试 Azure 连接：凭据无效

When testing the connection, you may encounter an error: **You appear to have entered invalid credentials. Please confirm you are using the correct information for an administrative account**. If `Tenant URL` and `secret token` are correct, check whether your group path contains characters that may be considered invalid JSON primitives (such as `.`). Removing such characters from the group path typically resolves the error.

测试连接时，您可能会遇到错误：**您似乎输入了无效凭据。请确认您使用的是管理帐户的正确信息**。如果“租户 URL”和“秘密令牌”正确，请检查您的组路径是否包含可能被视为无效 JSON 原语的字符（例如“.”）。从组路径中删除此类字符通常可以解决错误。

#### (Field) can’t be blank sync error

####（字段）不能为空白同步错误

When checking the Audit Events for the Provisioning, you can sometimes see the
error `Namespace can't be blank, Name can't be blank, and User can't be blank.`

检查供应的审核事件时，您有时可以看到
错误`命名空间不能为空，名称不能为空，用户不能为空。`

This is likely caused because not all required fields (such as first name and last name) are present for all users being mapped.

这可能是因为并非所有被映射的用户都存在所有必填字段（例如名字和姓氏）。

As a workaround, try an alternate mapping:

作为解决方法，请尝试替代映射：

1. Follow the Azure mapping instructions from above.

1. 按照上面的 Azure 映射说明进行操作。

2. Delete the`name.formatted` target attribute entry.

2. 删除`name.formatted` 目标属性条目。

3. Change the`displayName` source attribute to have `name.formatted` target attribute. 

3. 将`displayName` 源属性更改为具有`name.formatted` 目标属性。

