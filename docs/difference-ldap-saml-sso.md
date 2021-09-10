# The Difference Between LDAP and SAML SSO

January 28, 2021

What is the difference between [LDAP](http://jumpcloud.com/blog/what-is-ldap) and [SAML SSO (single sign-on)](http://jumpcloud.com/blog/saml-single-sign-sso)? Don’t both LDAP and SAML authenticate users to applications?

While both LDAP and SAML are authentication protocols and are often used for applications, the two are leveraged for very different use cases. In reality, though, organizations don’t often need to choose between using LDAP or SAML, but rather evaluate the most optimal way to leverage both protocols within their IT environment. For most organizations, leveraging a wide range of authentication protocols actually gives them access to more types of IT resources which can ultimately support their business objectives better. The trick, of course, is to do that without increasing the overhead for your IT team.

## A Brief History Lesson

Before we dive into the differences between the two authentication protocols, it’s best to first understand what each are and how they’ve evolved to where they are now. As a note, LDAP and SAML are just two of perhaps a half dozen [_major_ authentication protocols](https://jumpcloud.com/blog/protocols-using-identity-management), and maybe a dozen that are fairly widely used.

### A Brief Overview of LDAP

LDAP ( [Lightweight Directory Access Protocol](https://jumpcloud.com/blog/sso-ldap-integration/)) was created in the early 1990s by our good friend Tim Howes and his colleagues and quickly became one of the foundational authentication protocols used by IT networks. LDAP servers—such as OpenLDAP™ and [389 Directory](https://jumpcloud.com/blog/red-hat-directory-saas/)—are often used as an identity source of truth, also known as an [identity provider (IdP)](https://jumpcloud.com/blog/identity-provider-idp) or directory service. This ability, paired with another authentication protocol called Kerberos and system management abilities using policies and command execution, created the backbone for the traditional, on-prem directory service choice, [Microsoft® Active Directory®](https://jumpcloud.com/blog/what-is-active-directory-anyway/).

The main use of LDAP today is to authenticate users stored in the IdP to on-prem applications or other [Linux® server](https://jumpcloud.com/blog/linux-ldap-server/) processes. LDAP-based applications include OpenVPN, Jenkins, Kubernetes, Docker, Jira, and many others.

Traditionally, IT organizations have been forced to stand up their own LDAP infrastructure on-prem, along with the ancillary services required to keep the LDAP platform secure and operational. As a lightweight protocol, LDAP runs efficiently on systems, and gives IT organizations a great deal of control over authentication and authorization. Implementing it, however, is an arduous technical process, creating significant work upfront for IT admins with tasks such as high availability, performance monitoring, security, and more.

### A Brief Overview of SAML

SAML, on the other hand, was created in the early 2000s with the exclusive purpose of federating identities to web applications. The protocol was instantiated on the fact that there would be an identity provider already existing within an organization (at the time the assumption was Microsoft Active Directory). The SAML protocol didn’t seek to replace the IdP, but rather use it to assert the validity of a user’s identity.

That assertion would be leveraged by a service provider—or web application—via a secure XML exchange. The result was that an on-prem identity, traditionally stored in Active Directory (AD), could be extended to web applications. Vendors used SAML to create software that could extend one user identity from AD to a host of web applications, creating the first generation of Identity-as-a-Service (IDaaS)— [single sign-on (SSO)](https://jumpcloud.com/blog/what-is-single-sign-on) solutions. Examples of applications that support SAML authentication include Salesforce®, Slack, Trello, GitHub, Atlassian solution, and thousands of others.

![](https://jumpcloud.com//wp-content/uploads/2021/01/Hyper-Secure-with-SSO-MFA-aspect-ratio-160x110-1.png)

JumpCloud Single Sign-On

Hundreds of connectors to ensure you can grant access to cloud applications without friction

[Explore SSO Connectors](https://jumpcloud.com/sso-connectors)

Over the years, SAML has been extended to add functionality to provision user access to web applications as well. SAML-based solutions have historically been paired with a core directory service solution.

## The Difference Between LDAP and SAML SSO

When it comes to their areas of influence, LDAP and SAML SSO are as different as they come. LDAP, of course, is mostly focused towards facilitating on-prem authentication and other server processes. SAML extends user credentials to the cloud and other web applications.

While the differences are fairly significant, at their core, LDAP and SAML SSO are of the same ilk. They are effectively serving the same function—to help users connect to their IT resources. Because of this, they are often used in cooperation by IT organizations and have become staples of the identity management industry.

## The Costs of LDAP and SAML SSO

Although they are effective, common methods of LDAP and SAML SSO implementations can be costly to an enterprise’s time and budget. LDAP, as previously mentioned, is notoriously technical to instantiate and requires keen management to properly configure. SAML SSO is often [cloud-hosted](https://jumpcloud.com/blog/hosted-true-sso-single-sign/), but [pricing](https://jumpcloud.com/blog/okta-pricing/) models of these IDaaS solutions can be steep, not to mention the requirement for an IdP adds additional costs.

Thankfully, a new generation of identity provider is supporting these different protocols inside of one centralized cloud-based solution. Rather than face the daunting task of managing a wide range of authentication platforms and protocols, over 100k IT organizations trust [JumpCloud Directory Platform](https://jumpcloud.com/platform/) to accomplish complete identity management from one pane of glass.

## LDAP, SAML SSO, and More with DaaS

By hosting LDAP, SAML, and more from the cloud, a Directory-as-a-Service (DaaS) platform securely authenticates user identities to virtually any device (Windows, Mac®, Linux), application (on-prem or cloud), network, file server (on-prem LDAP Samba-based or cloud SAML-based), and more using a single set of credentials. That means less passwords to remember, less time spent signing in, and more freedom of choice for employees.

Beyond LDAP and SAML, IT organizations can leverage [group policy object (GPO)-like](https://jumpcloud.com/resources/cross-platform-gpo-like-capabilities/) functions to enforce security measures such as full disk encryption (FDE), multi-factor authentication (MFA), and password complexity requirements over user groups and Mac, Windows, and Linux systems. Admins can also use JumpCloud’s [Cloud RADIUS](https://jumpcloud.com/platform/cloud-radius) to tighten up network security with [VLAN tagging](https://jumpcloud.com/blog/what-is-vlan-tagging/) and more.

## The Cost of DaaS Platforms

The entire JumpCloud Directory Platform is available for free for the first 10 users and 10 devices in your organization. Beyond that, the [pricing model](https://jumpcloud.com/pricing/) scales as you do, with bulk discounts for larger organizations, education organizations, non-profits, and managed service providers (MSPs). We also offer a per protocol option (LDAP, SAML, or RADIUS) at a reduced rate.

If you would like to see our cloud directory platform in action before you buy, [try it for free today](https://jumpcloud.com/signup/), for 10 users and 10 devices. You can also [schedule a live demo](https://jumpcloud.com/demo/) of the product, or watch a recorded one [here](https://jumpcloud.com/demo). If you have any additional questions, feel free to [give us a call or send a note.](https://jumpcloud.com/contact/) You can also connect with our 24×7 premium in-app chat support during the first 10 days of your platform use and our engineers will help you.

- [LDAP](http://jumpcloud.com/blog/search?topics=ldap)
- [Single Sign On (SSO)](http://jumpcloud.com/blog/search?topics=single-sign-on-sso)

- [Redefining the Directory](http://jumpcloud.com/blog/search?collections=redefining-the-directory)
