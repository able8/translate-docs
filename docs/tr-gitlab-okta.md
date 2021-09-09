## GitLab and Okta

From: https://about.gitlab.com/handbook/business-technology/okta/

## 什么是Okta？

From the Okta website

来自 Okta 网站

> Okta is the foundation for secure connections between people and technology.
> It’s a service that gives employees, customers, and partners secure access to the tools they need to do their most important work.

> Okta 是人与技术之间安全连接的基础。
> 这是一项服务，可让员工、客户和合作伙伴安全地访问他们完成最重要工作所需的工具。

In practice - Okta is an Identity and Single Sign On solution for applications and Cloud entities.
It allows GitLab to consolidate authentication and authorisation to applications we use daily through a single dashboard and ensure a consistent, secure and auditable login experience for all our GitLab team members.

在实践中 - Okta 是适用于应用程序和云实体的身份和单点登录解决方案。
它允许 GitLab 通过单个仪表板整合对我们日常使用的应用程序的身份验证和授权，并确保我们所有 GitLab 团队成员获得一致、安全和可审核的登录体验。

### How is GitLab using Okta?

### GitLab 如何使用 Okta？

GitLab is using Okta for a few key goals :

GitLab 使用 Okta 来实现几个关键目标：

- We can use Okta to enable Zero-Trust based authentication controls upon our assets, so that we can allow authorised connections to key assets with a greater degree of certainty.
- We can better manage the login process to the 80+ and growing cloud applications that we use within our tech stack.
- We can better manage the provisioning and deprovisioning process for our users to access these application, by use of automation and integration into our HRIS system.
- We can make trust and risk based decisions on authentication requirements to key assets, and adapt these to ensure a consistent user experience.

- 我们可以使用 Okta 对我们的资产启用基于零信任的身份验证控制，以便我们可以以更高的确定性允许与关键资产的授权连接。
- 我们可以更好地管理我们在技术堆栈中使用的 80 多个且不断增长的云应用程序的登录过程。
- 通过使用自动化和集成到我们的 HRIS 系统中，我们可以更好地管理用户访问这些应用程序的配置和取消配置过程。
- 我们可以对关键资产的身份验证要求做出基于信任和风险的决策，并调整这些决策以确保一致的用户体验。

### What are the benefits to me using Okta as a user?

### 作为用户使用 Okta 对我有什么好处？

- A single Dashboard that is provided to all users, with all the applications you need in a single place.
- Managed SSO and Multi-Factor Authentication that learns and adapts to your login patterns, making life simpler to access the assets you need.
- Transparent Security controls with a friendly user experience.

- 提供给所有用户的单一仪表板，在一个地方提供您需要的所有应用程序。
- 托管 SSO 和多因素身份验证，可学习并适应您的登录模式，让访问您需要的资产变得更简单。
- 具有友好用户体验的透明安全控制。

### What are the benefits to me as an application administrator to using Okta?

### 作为应用程序管理员，使用 Okta 对我有什么好处？

- Automated provisioning and group management
- Ability to transparently manage shared credentials to web applications without disclosing the credentials to users
- Centralised access for users, making it easy to add, remove and change the application profile without the need to update all users.

- 自动配置和组管理
- 能够透明地管理 Web 应用程序的共享凭据，而无需向用户披露凭据
- 用户集中访问，无需更新所有用户即可轻松添加、删除和更改应用程序配置文件。

## How do I get my Okta account set up? 

## 如何设置我的 Okta 帐户？

All GitLab team-members will have an Okta account set up as part of their onboarding process. You should already have an activation email in both your Gmail and Personal Accounts. For efficiency, please follow the onboarding process for setting up Okta and set up 1Password first and follow that up with Okta. Please also set up Okta from your computer rather than your mobile or the mobile app, as you will be guided to set up the mobile app as part of the onboarding process.

所有 GitLab 团队成员都将设置一个 Okta 帐户，作为其入职流程的一部分。您的 Gmail 和个人帐户中应该已经有一封激活电子邮件。为提高效率，请遵循设置 Okta 的入门流程并先设置 1Password，然后再使用 Okta。还请从您的计算机而不是您的手机或移动应用程序设置 Okta，因为将指导您设置移动应用程序作为入职流程的一部分。

Follow the GitLab Okta [Getting Started Guide](https://docs.google.com/document/d/1x2NJan0job5nM5tT8HF6yofg-Y2aAsSVKc6qNnCuoxo/) and [FAQs](http://about.gitlab.com/handbook/business-technology/okta/okta-enduser-faq/).

遵循 GitLab Okta [入门指南](https://docs.google.com/document/d/1x2NJan0job5nM5tT8HF6yofg-Y2aAsSVKc6qNnCuoxo/) 和 [FAQs](http://about.gitlab.com/handbook/business-technology/okta/okta-enduser-faq/)。

We have also prepared Introductory Videos on [Okta Setup](https://youtu.be/upJ4p3lKYKw), [Setting up MFA/Yubikeys](https://youtu.be/9UyKml_aO3s), [Configuring Applications](https://youtu.be/xS2CarGUPLc) and [Dashboard Tips](https://youtu.be/xQQwa_pbe2U).

我们还准备了关于 [Okta 设置](https://youtu.be/upJ4p3lKYKw)、[设置MFA/Yubikeys](https://youtu.be/9UyKml_aO3s)、[配置应用程序](https:///youtu.be/xS2CarGUPLc) 和 [仪表板提示](https://youtu.be/xQQwa_pbe2U)。

We recommend particularly that once your account is set up, you set up an additional MFA factor (either YubiKey or Google Authenticator/TOTP) in case there's an issue with one of your MFA factors.

我们特别建议您在设置帐户后，设置额外的 MFA 因素（YubiKey 或 Google 身份验证器/TOTP），以防您的 MFA 因素之一出现问题。

### Setting up my Okta Account requires me to use the app Okta Verify on my phone, and I don't like that.

### 设置我的 Okta 帐户需要我在手机上使用 Okta 验证应用程序，我不喜欢那样。

Our Okta implementation defaults to using Okta Verify as the Required MFA factor.
Okta Verify is a safe and secure application that allows push notifications and one-time tokencodes on your phone to validate your login.
It is supported on iPhone, Android and Windows Phones.

我们的 Okta 实现默认使用 Okta Verify 作为必需的 MFA 因子。
Okta Verify 是一款安全可靠的应用程序，它允许在您的手机上使用推送通知和一次性令牌代码来验证您的登录信息。
它在 iPhone、Android 和 Windows 手机上受支持。

For some people, there are issues with installing a verification app on their phone.
If there is some reason that this is not appropriate for your geography or other reasons, please submit an issue to [Opt Out](https://gitlab.com/gitlab-com/business-technology/team-member-enablement/issue-tracker/issues/new?issuable_template=okta_verify_optout) and we can add you to an authentication group that will make Okta Verify optional.
Please note that we still recommend that you set up at least two MFA factors, in case something happens to one of your factors.

对于某些人来说，在手机上安装验证应用会出现问题。
如果由于某些原因这不适合您的地理位置或其他原因，请提交问题到 [Opt Out](https://gitlab.com/gitlab-com/business-technology/team-member-enablement/issue-tracker/issues/new?issuable_template=okta_verify_optout)，我们可以将您添加到身份验证组，这将使 Okta 验证成为可选。
请注意，我们仍然建议您至少设置两个 MFA 因子，以防其中一个因子发生问题。

## I forgot my password/my login doesn't work, what do I do?

## 我忘记了密码/我的登录名无效，我该怎么办？

There is a "need help signing in?" button on the login screen.
If you expand this there is a link to an automated password reset process via email.
You will need to know the answers to your security question(s) to use this.

有一个“需要帮助登录？”登录屏幕上的按钮。
如果您展开它，则会有一个通过电子邮件自动重置密码过程的链接。
您需要知道安全问题的答案才能使用它。

We recommend that you store your Okta password in 1Password as well as your Security Questions there.
Please review the [1Password Guidelines](http://about.gitlab.com/handbook/security/#1password-guidelines) for best ways to use Okta and 1Password together.

我们建议您将 Okta 密码以及您的安全问题存储在 1Password 中。
请查看 [1Password 指南](http://about.gitlab.com/handbook/security/#1password-guidelines) 以了解结合使用 Okta 和 1Password 的最佳方法。

### I forgot my Security Questions, how do I reset my password?

### 我忘记了我的安全问题，如何重置我的密码？

Firstly, review the [1Password Guidelines](http://about.gitlab.com/handbook/security/#1password-guidelines).
Then head to `#it_help` in Slack and ask for a temporary password to be issued.
You will be issued a temporary password at which point you can reset your access.

首先，查看[1Password Guidelines](http://about.gitlab.com/handbook/security/#1password-guidelines)。
然后前往 Slack 中的 `#it_help` 并要求提供一个临时密码。
您将获得一个临时密码，此时您可以重置您的访问权限。

### I changed my phone and now can't do MFA, what do I do?

### 我换了手机，现在不能做MFA，我该怎么办？

No worries! You can easily reset your own MFA code for Okta if you did not wipe/return your old phone yet.

不用担心！如果您还没有擦除/归还旧手机，您可以轻松地为 Okta 重置您自己的 MFA 代码。

Firstly, sign into your Okta webpage by going to [gitlab.okta.com](https://gitlab.okta.com) use your email, password, and the MFA code on your old phone.

首先，前往 [gitlab.okta.com](https://gitlab.okta.com) 使用您的电子邮件、密码和旧手机上的 MFA 代码登录您的 Okta 网页。

Once you're on the Okta webpage click on your name and then click settings.

进入 Okta 网页后，单击您的姓名，然后单击设置。

![Okta Settings](http://about.gitlab.com/handbook/business-technology/Okta-settings.png)

Scroll down until you see "Extra Verification", once you're there click "remove" to disable that specific MFA and then click setup to configure the new MFA code on your new phone.

向下滚动，直到看到“额外验证”，在那里单击“删除”以禁用该特定 MFA，然后单击设置以在新手机上配置新的 MFA 代码。

![Okta 2FA](http://about.gitlab.com/handbook/business-technology/Okta-2FA.png)

If you wiped and returned your old mobile device you could use a [Yubikey](https://www.yubico.com/products/) as another form of authentication (if you have one set one up). Use that to access your settings page and follow the steps above to reset your Okta MFA.

如果您擦除并归还旧移动设备，则可以使用 [Yubikey](https://www.yubico.com/products/)作为另一种身份验证形式（如果您已设置)。使用它来访问您的设置页面并按照上述步骤重置您的 Okta MFA。

Lost all your MFA Factors?
Head to `#it_help` in Slack and ask for a MFA Reset.
Once your Factors have been reset, please set up at least two MFA factors (Yubikey or Google Authenticator, see [this video](https://youtu.be/9UyKml_aO3s)).

失去了所有的 MFA 因素？
前往 Slack 中的“#it_help”并要求重置 MFA。
重置因子后，请设置至少两个 MFA 因子（Yubikey 或 Google 身份验证器，请参阅 [此视频](https://youtu.be/9UyKml_aO3s))。

### Managing Okta Access Using Google Groups 

### 使用 Google Groups 管理 Okta Access

The GitLab Team Member Enablement team has created a new process for Owners and Provisioners to manage access to Okta applications. If you are listed as an Owner/Provisioner for an application in the [tech stack](https://docs.google.com/spreadsheets/d/1mTNZHsK3TWzQdeFqkITKA0pHADjuurv37XMuHv12hDU/edit#gid=1906611965) you will be using the method below to add users to a Google group, which will then sync this group to Okta and assign the application to users. This process was created to empower business application owners to effect Access Requests which require Okta application assignment.

GitLab 团队成员启用团队为所有者和供应商创建了一个新流程来管理对 Okta 应用程序的访问。如果您在 [技术堆栈](https://docs.google.com/spreadsheets/d/1mTNZHsK3TWzQdeFqkITKA0pHADjuurv37XMuHv12hDU/edit#gid=1906611965) 中被列为应用程序的所有者/供应商，您将使用以下方法添加用户添加到 Google 群组，然后该群组会将此群组同步到 Okta 并将应用程序分配给用户。创建此流程是为了使业务应用程序所有者能够影响需要 Okta 应用程序分配的访问请求。

- Sign in to[Google Groups](https://groups.google.com/).
- Click My groups.
- Click the name of the group were you want to add/remove a user. (Note that all Google groups which manage users in Okta application start with okta-xxxxx-users)
- Next press the `People tab` on the left side and select `Members`.
![Screenshot](https://gitlab.com/plaurinavicius/image-sources-for-runbooks/-/raw/master/Screenshot_2020-11-05_at_14.27.05.png)

- This will show you all the members currently in the group.
- To add a member press the`Add Members` button. To remove access mouse over a user and press on the little white box that appears, this will mark the user. After that on the right side press the remove member button (Looks like a circle with a horizontal line across).

- 登录[Google 网上论坛](https://groups.google.com/)。
- 单击我的组。
- 单击要添加/删除用户的组的名称。 （请注意，所有在 Okta 应用程序中管理用户的 Google 群组都以 okta-xxxxx-users 开头）
- 接下来按左侧的“人物标签”并选择“成员”。
- 这将显示当前组中的所有成员。
- 要添加成员，请按`添加成员`按钮。要删除用户上的访问鼠标并按下出现的小白框，这将标记用户。然后在右侧按下删除成员按钮（看起来像一个带有水平线的圆圈）。

When a member is added/removed from the group it may take up to 1 hour for the sync to happen between Google and Okta. Once the sync happens the user will see the application in Okta, if removed the opposite.
If you have any questions or require assistance please reach out to the IT team in the #it-help Slack channel.

从群组中添加/删除成员后，Google 和 Okta 之间的同步可能需要长达 1 小时的时间。一旦同步发生，用户将在 Okta 中看到应用程序，如果删除相反。
如果您有任何问题或需要帮助，请通过 #it-help Slack 频道联系 IT 团队。

### My Okta account has been locked out because of failed attempts, what do I do?

### 由于尝试失败，我的 Okta 帐户已被锁定，我该怎么办？

Head to `#it_help` and ask to have your account unlocked.
As a precaution, you will also need to change your Okta Password.

前往`#it_help` 并要求解锁您的帐户。
作为预防措施，您还需要更改 Okta 密码。

## Why isn't an application I need available in Okta?

## 为什么 Okta 中没有我需要的应用程序？

Create a [new application setup issue](https://gitlab.com/gitlab-com/business-technology/change-management/issues/new?issuable_template=change_management_okta) and fill in as much information as you can.

创建一个[新的应用程序设置问题](https://gitlab.com/gitlab-com/business-technology/change-management/issues/new?issuable_template=change_management_okta) 并尽可能多地填写信息。

Okta is currently configured with assigned groups/roles based on a team member's role/group.
Refer to the [Access Change Request](http://about.gitlab.com/handbook/business-technology/team-member-enablement/onboarding-access-requests/access-requests/#access-change-request) section of the handbook for additional information on why an application may not be available in Okta.

Okta 当前根据团队成员的角色/组配置了分配的组/角色。
参考[访问变更请求](http://about.gitlab.com/handbook/business-technology/team-member-enablement/onboarding-access-requests/access-requests/#access-change-request)部分有关为什么应用程序可能无法在 Okta 中使用的其他信息的手册。

### How do I get my application set up within Okta?

### 如何在 Okta 中设置我的应用程序？

If you are an application owner please submit a [new application setup issue](https://gitlab.com/gitlab-com/business-technology/change-management/issues/new?issuable_template=change_management_okta) on the Okta project page for your application.
We will work with you to verify details and provide setup instructions.

如果您是应用程序所有者，请在 Okta 项目页面上提交 [新应用程序设置问题](https://gitlab.com/gitlab-com/business-technology/change-management/issues/new?issuable_template=change_management_okta)你的申请。
我们将与您一起验证详细信息并提供设置说明。

### I have an application that uses a shared password for my team, can I move this to Okta?

### 我有一个应用程序为我的团队使用共享密码，我可以将它移到 Okta 吗？

Yes you can!
Submit a [new application setup issue](https://gitlab.com/gitlab-com/business-technology/change-management/issues/new?issuable_template=change_management_okta) on the Okta project page for your application.
We will work with you to verify details and provide setup instructions.

是的你可以！
在 Okta 项目页面上为您的应用程序提交 [新应用程序设置问题](https://gitlab.com/gitlab-com/business-technology/change-management/issues/new?issuable_template=change_management_okta)。
我们将与您一起验证详细信息并提供设置说明。

## I'm getting asked to MFA authenticate a lot, is that normal?

## 我经常被要求进行 MFA 身份验证，这正常吗？

The way we have Okta setup should require you to authenticate once with MFA when you start your working day, and that session should last for the rest of your work day.
It's recommended that you login via the [Okta Dashboard](https://gitlab.okta.com/) at the beginning of your day, and then use either the dashboard or the Okta plugin for applications during your work day.

我们设置 Okta 的方式应该要求您在开始工作日时使用 MFA 进行一次身份验证，并且该会话应该持续到您工作日的剩余时间。
建议您在一天开始时通过 [Okta Dashboard](https://gitlab.okta.com/) 登录，然后在您的工作日使用仪表板或 Okta 插件进行应用程序。

For some applications, we enforce an additional MFA step periodically because of the sensitivity of the data in them.
We are also trialling a risk-based authentication algorithm that may ask you to re-authenticate if anomalous behaviour is detected on your account or Okta detects an unusual login pattern.
At this stage, BambooHR and Greenhouse require an additional authentication step. 

对于某些应用程序，由于其中数据的敏感性，我们会定期执行额外的 MFA 步骤。
我们还在试用基于风险的身份验证算法，如果在您的帐户中检测到异常行为或 Okta 检测到异常登录模式，该算法可能会要求您重新进行身份验证。
在此阶段，BambooHR 和 Greenhouse 需要额外的身份验证步骤。

If you are having problems with being asked for multiple MFA authentications during the day, please [log an issue](https://gitlab.com/gitlab-com/business-technology/change-management/issues) and we can look into it.

如果您在白天被要求进行多次 MFA 身份验证时遇到问题，请[记录问题](https://gitlab.com/gitlab-com/business-technology/change-management/issues)，我们可以调查它。

### Why does GitLab.com ask for an additional MFA when I login via Okta?

### 当我通过 Okta 登录时，为什么 GitLab.com 要求额外的 MFA？

Your gitlab.com account will have 2FA installed as required by our policy.
Note that the 2FA for GitLab.com is different to the MFA you use to log into Okta.
[This issue](https://gitlab.com/gitlab-com/gl-infra/infrastructure/issues/7397) has been opened to propose a solution.

您的 gitlab.com 帐户将按照我们的政策要求安装 2FA。
请注意，GitLab.com 的 2FA 与您用于登录 Okta 的 MFA 不同。
[本期](https://gitlab.com/gitlab-com/gl-infra/infrastructure/issues/7397) 已开放提出解决方案。

## Where do I go if I have any questions?

## 如果我有任何问题，我应该去哪里？

- For Okta help, setup and integration questions:`#it_help` slack channel

- 有关 Okta 帮助、设置和集成问题：`#it_help` 松弛频道
