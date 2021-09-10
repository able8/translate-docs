# The Difference Between LDAP and SAML SSO

# LDAP 和 SAML SSO 的区别

January 28, 2021

What is the difference between [LDAP](http://jumpcloud.com/blog/what-is-ldap) and [SAML SSO (single sign-on)](http://jumpcloud.com/blog/saml-single-sign-sso)? Don’t both LDAP and SAML authenticate users to applications?

[LDAP](http://jumpcloud.com/blog/what-is-ldap) 和 [SAML SSO（单点登录）](http://jumpcloud.com/blog/saml-single)有什么区别-sign-sso)？ LDAP 和 SAML 不是都对应用程序的用户进行身份验证吗？

While both LDAP and SAML are authentication protocols and are often used for applications, the two are leveraged for very different use cases. In reality, though, organizations don’t often need to choose between using LDAP or SAML, but rather evaluate the most optimal way to leverage both protocols within their IT environment. For most organizations, leveraging a wide range of authentication protocols actually gives them access to more types of IT resources which can ultimately support their business objectives better. The trick, of course, is to do that without increasing the overhead for your IT team.

虽然 LDAP 和 SAML 都是身份验证协议，并且通常用于应用程序，但两者用于非常不同的用例。但实际上，组织通常不需要在使用 LDAP 或 SAML 之间做出选择，而是评估在其 IT 环境中利用这两种协议的最佳方式。对于大多数组织而言，利用广泛的身份验证协议实际上可以让他们访问更多类型的 IT 资源，最终可以更好地支持他们的业务目标。当然，诀窍是在不增加 IT 团队开销的情况下做到这一点。

## A Brief History Lesson

## 简史课

Before we dive into the differences between the two authentication protocols, it’s best to first understand what each are and how they’ve evolved to where they are now. As a note, LDAP and SAML are just two of perhaps a half dozen [_major_ authentication protocols](https://jumpcloud.com/blog/protocols-using-identity-management), and maybe a dozen that are fairly widely used.

在我们深入探讨两种身份验证协议之间的差异之前，最好先了解每种协议是什么以及它们是如何发展到现在的。需要注意的是，LDAP 和 SAML 只是大约六种 [_主要_身份验证协议](https://jumpcloud.com/blog/protocols-using-identity-management) 中的两种，而且可能还有十几种使用相当广泛。

### A Brief Overview of LDAP

### LDAP 的简要概述

LDAP ( [Lightweight Directory Access Protocol](https://jumpcloud.com/blog/sso-ldap-integration/)) was created in the early 1990s by our good friend Tim Howes and his colleagues and quickly became one of the foundational authentication protocols used by IT networks. LDAP servers—such as OpenLDAP™ and [389 Directory](https://jumpcloud.com/blog/red-hat-directory-saas/)—are often used as an identity source of truth, also known as an [identity provider (IdP)](https://jumpcloud.com/blog/identity-provider-idp) or directory service. This ability, paired with another authentication protocol called Kerberos and system management abilities using policies and command execution, created the backbone for the traditional, on-prem directory service choice, [Microsoft® Active Directory®](https://jumpcloud.com/blog/what-is-active-directory-anyway/).

LDAP（[轻量级目录访问协议](https://jumpcloud.com/blog/sso-ldap-integration/)）由我们的好朋友 Tim Howes 和他的同事在 1990 年代初期创建，并迅速成为基础认证之一IT 网络使用的协议。 LDAP 服务器（例如 OpenLDAP™ 和 [389 Directory](https://jumpcloud.com/blog/red-hat-directory-saas/)）通常用作身份真实来源，也称为[身份提供者(IdP)](https://jumpcloud.com/blog/identity-provider-idp) 或目录服务。这种能力与另一种称为 Kerberos 的身份验证协议以及使用策略和命令执行的系统管理能力相结合，为传统的本地目录服务选择 [Microsoft® Active Directory®](https://jumpcloud.com/博客/what-is-active-directory-anyway/)。

The main use of LDAP today is to authenticate users stored in the IdP to on-prem applications or other [Linux® server](https://jumpcloud.com/blog/linux-ldap-server/) processes. LDAP-based applications include OpenVPN, Jenkins, Kubernetes, Docker, Jira, and many others.

今天 LDAP 的主要用途是将存储在 IdP 中的用户验证到本地应用程序或其他 [Linux® 服务器](https://jumpcloud.com/blog/linux-ldap-server/) 进程。基于 LDAP 的应用程序包括 OpenVPN、Jenkins、Kubernetes、Docker、Jira 等。

Traditionally, IT organizations have been forced to stand up their own LDAP infrastructure on-prem, along with the ancillary services required to keep the LDAP platform secure and operational. As a lightweight protocol, LDAP runs efficiently on systems, and gives IT organizations a great deal of control over authentication and authorization. Implementing it, however, is an arduous technical process, creating significant work upfront for IT admins with tasks such as high availability, performance monitoring, security, and more.

传统上，IT 组织被迫在本地建立自己的 LDAP 基础设施，以及保持 LDAP 平台安全和运行所需的辅助服务。作为一种轻量级协议，LDAP 在系统上高效运行，并为 IT 组织提供了对身份验证和授权的大量控制。然而，实施它是一个艰巨的技术过程，为 IT 管理员创建了大量的前期工作，这些任务包括高可用性、性能监控、安全性等。

### A Brief Overview of SAML

### SAML 的简要概述

SAML, on the other hand, was created in the early 2000s with the exclusive purpose of federating identities to web applications. The protocol was instantiated on the fact that there would be an identity provider already existing within an organization (at the time the assumption was Microsoft Active Directory). The SAML protocol didn’t seek to replace the IdP, but rather use it to assert the validity of a user’s identity. 

另一方面，SAML 是在 2000 年代初创建的，其唯一目的是将身份联合到 Web 应用程序。该协议是基于这样一个事实来实例化的，即组织内已经存在一个身份提供者（当时假设是 Microsoft Active Directory）。 SAML 协议并不寻求取代 IdP，而是使用它来断言用户身份的有效性。

That assertion would be leveraged by a service provider—or web application—via a secure XML exchange. The result was that an on-prem identity, traditionally stored in Active Directory (AD), could be extended to web applications. Vendors used SAML to create software that could extend one user identity from AD to a host of web applications, creating the first generation of Identity-as-a-Service (IDaaS)— [single sign-on (SSO)](https://jumpcloud.com/blog/what-is-single-sign-on) solutions. Examples of applications that support SAML authentication include Salesforce®, Slack, Trello, GitHub, Atlassian solution, and thousands of others.

服务提供商或 Web 应用程序将通过安全的 XML 交换来利用该断言。结果是，传统上存储在 Active Directory (AD) 中的本地身份可以扩展到 Web 应用程序。供应商使用 SAML 创建可以将一个用户身份从 AD 扩展到一系列 Web 应用程序的软件，从而创建了第一代身份即服务 (IDaaS) — [单点登录 (SSO)](https://jumpcloud.com/blog/what-is-single-sign-on) 解决方案。支持 SAML 身份验证的应用程序示例包括 Salesforce®、Slack、Trello、GitHub、Atlassian 解决方案等。

![](https://jumpcloud.com//wp-content/uploads/2021/01/Hyper-Secure-with-SSO-MFA-aspect-ratio-160x110-1.png)

JumpCloud Single Sign-On

JumpCloud 单点登录

Hundreds of connectors to ensure you can grant access to cloud applications without friction

数百个连接器，确保您可以毫无障碍地授予对云应用程序的访问权限

[Explore SSO Connectors](https://jumpcloud.com/sso-connectors)

[探索 SSO 连接器](https://jumpcloud.com/sso-connectors)

Over the years, SAML has been extended to add functionality to provision user access to web applications as well. SAML-based solutions have historically been paired with a core directory service solution.

多年来，SAML 已经扩展为添加功能，以提供用户对 Web 应用程序的访问权限。基于 SAML 的解决方案历来与核心目录服务解决方案配对。

## The Difference Between LDAP and SAML SSO

## LDAP 和 SAML SSO 的区别

When it comes to their areas of influence, LDAP and SAML SSO are as different as they come. LDAP, of course, is mostly focused towards facilitating on-prem authentication and other server processes. SAML extends user credentials to the cloud and other web applications.

就其影响领域而言，LDAP 和 SAML SSO 各不相同。当然，LDAP 主要侧重于促进本地身份验证和其他服务器进程。 SAML 将用户凭据扩展到云和其他 Web 应用程序。

While the differences are fairly significant, at their core, LDAP and SAML SSO are of the same ilk. They are effectively serving the same function—to help users connect to their IT resources. Because of this, they are often used in cooperation by IT organizations and have become staples of the identity management industry.

虽然差异相当显着，但从本质上讲，LDAP 和 SAML SSO 是同一类。它们有效地提供相同的功能——帮助用户连接到他们的 IT 资源。因此，它们经常被 IT 组织用于合作，并已成为身份管理行业的主要内容。

## The Costs of LDAP and SAML SSO

## LDAP 和 SAML SSO 的成本

Although they are effective, common methods of LDAP and SAML SSO implementations can be costly to an enterprise’s time and budget. LDAP, as previously mentioned, is notoriously technical to instantiate and requires keen management to properly configure. SAML SSO is often [cloud-hosted](https://jumpcloud.com/blog/hosted-true-sso-single-sign/), but [pricing](https://jumpcloud.com/blog/okta-pricing/) models of these IDaaS solutions can be steep, not to mention the requirement for an IdP adds additional costs.

尽管它们是有效的，但 LDAP 和 SAML SSO 实施的常用方法可能会增加企业的时间和预算。如前所述，LDAP 的实例化技术是出了名的，需要敏锐的管理才能正确配置。 SAML SSO 通常是 [云托管](https://jumpcloud.com/blog/hosted-true-sso-single-sign/)，但 [定价](https://jumpcloud.com/blog/okta-pricing/) 这些 IDaaS 解决方案的模型可能很陡峭，更不用说对 IdP 的要求会增加额外的成本。

Thankfully, a new generation of identity provider is supporting these different protocols inside of one centralized cloud-based solution. Rather than face the daunting task of managing a wide range of authentication platforms and protocols, over 100k IT organizations trust [JumpCloud Directory Platform](https://jumpcloud.com/platform/) to accomplish complete identity management from one pane of glass.

值得庆幸的是，新一代身份提供商正在一个基于云的集中式解决方案中支持这些不同的协议。超过 10 万 IT 组织信任 [JumpCloud 目录平台](https://jumpcloud.com/platform/) 从一个管理平台完成完整的身份管理，而不是面临管理各种身份验证平台和协议的艰巨任务。

## LDAP, SAML SSO, and More with DaaS

## LDAP、SAML SSO 等 DaaS

By hosting LDAP, SAML, and more from the cloud, a Directory-as-a-Service (DaaS) platform securely authenticates user identities to virtually any device (Windows, Mac®, Linux), application (on-prem or cloud), network, file server (on-prem LDAP Samba-based or cloud SAML-based), and more using a single set of credentials. That means less passwords to remember, less time spent signing in, and more freedom of choice for employees.

通过从云托管 LDAP、SAML 等，目录即服务 (DaaS) 平台可以安全地验证几乎任何设备（Windows、Mac®、Linux）、应用程序（本地或云）的用户身份，网络、文件服务器（基于本地 LDAP Samba 或基于云 SAML 的），以及更多使用一组凭据。这意味着要记住的密码更少，登录所花费的时间更少，员工的选择自由度更高。

Beyond LDAP and SAML, IT organizations can leverage [group policy object (GPO)-like](https://jumpcloud.com/resources/cross-platform-gpo-like-capabilities/) functions to enforce security measures such as full disk encryption (FDE), multi-factor authentication (MFA), and password complexity requirements over user groups and Mac, Windows, and Linux systems. Admins can also use JumpCloud's [Cloud RADIUS](https://jumpcloud.com/platform/cloud-radius) to tighten up network security with [VLAN tagging](https://jumpcloud.com/blog/what-is-vlan-tagging/) and more.

除了 LDAP 和 SAML，IT 组织还可以利用 [组策略对象 (GPO)-like](https://jumpcloud.com/resources/cross-platform-gpo-like-capabilities/) 功能来实施安全措施，例如完整磁盘用户组和 Mac、Windows 和 Linux 系统的加密 (FDE)、多因素身份验证 (MFA) 和密码复杂性要求。管理员还可以使用 JumpCloud 的 [Cloud RADIUS](https://jumpcloud.com/platform/cloud-radius) 通过 [VLAN 标记](https://jumpcloud.com/blog/what-is-vlan)加强网络安全-标记/)等等。

## The Cost of DaaS Platforms 

## DaaS 平台的成本

The entire JumpCloud Directory Platform is available for free for the first 10 users and 10 devices in your organization. Beyond that, the [pricing model](https://jumpcloud.com/pricing/) scales as you do, with bulk discounts for larger organizations, education organizations, non-profits, and managed service providers (MSPs). We also offer a per protocol option (LDAP, SAML, or RADIUS) at a reduced rate.

您组织中的前 10 位用户和 10 台设备可以免费使用整个 JumpCloud Directory Platform。除此之外，[定价模型](https://jumpcloud.com/pricing/) 随您扩展，为大型组织、教育组织、非营利组织和托管服务提供商 (MSP) 提供批量折扣。我们还以较低的价格提供每个协议选项（LDAP、SAML 或 RADIUS)。

If you would like to see our cloud directory platform in action before you buy, [try it for free today](https://jumpcloud.com/signup/), for 10 users and 10 devices. You can also [schedule a live demo](https://jumpcloud.com/demo/) of the product, or watch a recorded one [here](https://jumpcloud.com/demo). If you have any additional questions, feel free to [give us a call or send a note.](https://jumpcloud.com/contact/) You can also connect with our 24×7 premium in-app chat support during the first 10 days of your platform use and our engineers will help you.

如果您想在购买之前查看我们的云目录平台的运行情况，[今天免费试用](https://jumpcloud.com/signup/)，适用于 10 个用户和 10 个设备。您还可以[安排产品的现场演示](https://jumpcloud.com/demo/)，或观看录制的[此处](https://jumpcloud.com/demo)。如果您有任何其他问题，请随时[给我们打电话或发送便条。](https://jumpcloud.com/contact/) 您还可以在活动期间联系我们的 24×7 高级应用内聊天支持在您使用平台的前 10 天，我们的工程师将为您提供帮助。

- [LDAP](http://jumpcloud.com/blog/search?topics=ldap)
- [Single Sign On (SSO)](http://jumpcloud.com/blog/search?topics=single-sign-on-sso)

- [LDAP](http://jumpcloud.com/blog/search?topics=ldap)
- [单点登录 (SSO)](http://jumpcloud.com/blog/search?topics=single-sign-on-sso)

- [Redefining the Directory](http://jumpcloud.com/blog/search?collections=redefining-the-directory) 

- [重新定义目录](http://jumpcloud.com/blog/search?collections=redefining-the-directory)

