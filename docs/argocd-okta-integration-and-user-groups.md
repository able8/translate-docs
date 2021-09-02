# ArgoCD: Okta integration, and user groups

05/17/2021

The idea is that we don’t add user accounts locally in the ArgoCD’s ConfigMap, but instead will use our Okta users databases and Okta will perform their authentication. And on the ArgoCD side will do users’ authorization, i.e. will check their permission boundaries.

Also, by using the SSO we will be able to create user groups that will have various roles tied to specific Projects.

We will use the [SAML (with Dex)](https://argoproj.github.io/argo-cd/operator-manual/user-management/okta/#saml-with-dex), see also the [What is: SAML – an overview, its structure and requests tracing between a Jenkins and Okta SSO](https://rtfm.co.ua/en/what-is-saml-an-overview-its-structure-and-requests-tracing-between-a-jenkins-and-okta-sso/) post for more details on the SAML.

Contents

- [Okta configuration](http://rtfm.co.ua#Okta_configuration "Okta configuration")
- [ArgoCD configuration](http://rtfm.co.ua#ArgoCD_configuration "ArgoCD configuration")
  - [Roles, and user groups in ArgoCD](http://rtfm.co.ua#Roles_and_user_groups_in_ArgoCD "Roles, and user groups in ArgoCD")
  - [ArgoCD CLI login](http://rtfm.co.ua#ArgoCD_CLI_login "ArgoCD CLI login")
- [Errors and problems](http://rtfm.co.ua#Errors_and_problems "Errors and problems")
  - [Bad Request – User session error](http://rtfm.co.ua#Bad_Request_-_User_session_error "Bad Request – User session error")
  - [Failed to authenticate: no attribute with name “group”: [email]](http://rtfm.co.ua#Failed_to_authenticate_no_attribute_with_name_group_email "Failed to authenticate: no attribute with name “group”: [email]")

# Okta configuration

Go to your Okta account, and create a SAML application:

[![](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_124710.png)](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_124710.png)

[![](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_124758.png)](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_124758.png)

For the logo, I’ve used this pic:

[![](https://rtfm.co.ua/wp-content/uploads/2021/05/argo-horizontal-color.png)](https://rtfm.co.ua/wp-content/uploads/2021/05/argo-horizontal-color.png) On the next page, set the following:

- _Single sign on URL_ and _Audience URI (SP Entity ID)_:  your ArgoCD URL like _argourl.com/api/dex/callback_
- в_Attribute Statements_: name – _email_, format Basic, Value – _user.email_
- в_Group Attribute Statements_: name – _group_, format Basic, in the Filter you can use a regex as it’s shown in the [documentation](https://argoproj.github.io/argo-cd/operator-manual/user-management/okta/#saml-with-dex) – then, Okta groups will be filtered before passing them to the ArgoCD during  [Authentication context](https://rtfm.co.ua/en/what-is-saml-an-overview-its-structure-and-requests-tracing-between-a-jenkins-and-okta-sso/#Authentication_context), or filter nothing by setting the regex as “\*”:

[![](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_132008.png)](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_132008.png)

Or you can filter a list of groups by using the `|`:

[![](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_133047.png)](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_133047.png)

On the next page, set the _I’m an Okta customer adding an internal app_, and click Finish:

[![](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_125608.png)](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_125608.png)

Assign the Application to a user or group:

[![](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_125657.png)](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_125657.png)

Switch to the Sign On tab, and click on the _View Setup Instructions_ button:

[![](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_125737.png)](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_125737.png) Now, let’s go to our ArgoCD instance.

# ArgoCD configuration

Edit the `argocd-cm` ConfigMap:

kubectl -n dev-1-18-devops-argocd-ns edit configmap argocd-cm

Set:

- _url_: (can be already set, just check its value) – ArgoCD’s instance external URL
- _ssoURL_: Identity Provider Single Sign-On URL from the page that was opened after you’ve clicked the _View Setup Instructions_ button
- _caData_: an Okta’s certificate from the same page encoded with base64, for example on the [https://www.base64encode.org](https://www.base64encode.org/) site

Now, your config will look like this:

[![](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_130354.png)](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_130354.png)

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

Edit the `argocd-rbac-cm` ConfigMap, add the _DevOps_ group (will get from the Okta) mapping to the `role:admin`:

```
g, DevOps, role:admin
```

Also, here you can disable the _admin_ user, and add Backend and Web groups bindings to applications on their projects:

```
...
data:
admin.enabled: "false"
policy.csv: |
    p, role:backend-app-admin, applications, *, Backend/*, allow
    p, role:web-app-admin, applications, *, Web/*, allow
    g, DevOps, role:admin
policy.default: role:''
scopes: '[email,groups]'
...
```

Keep in mind, that RBAC is case sensitive, so DevOps != devops. So you have to specify a group here the same as it is set in the Okta.

Save, and try to log in:

[![](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_131242.png)](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_131242.png)

[![](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_150422-1.png)](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_150422-1.png)

## ArgoCD CLI login

To use CLI with SSO, just add the  `--sso` option, and in the `--username` – specify your Okta’s login:

argocd login dev-1-18.argocd.example.com --sso --username arseniy@example.com

Opening browser for authentication

...

Authentication successful

'arseniy@example.com' logged in successfully

Context 'dev-1-18.argocd.example.com' updated

A default browser will be opened to authenticate you with Okta.

# Errors and problems

## Bad Request – User session error

When using the [Dex](https://dexidp.io/docs/connectors/saml/) keep in mind, that it doesn’t allow the Provider-initated login, i.e. you can log in to the ArgoCD instance from your Okta – you’ll get the 400 error:

[![](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_135016.png)](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_135016.png)

Thus, worth hiding the application from users:

[![](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_135122.png)](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_135122.png)

## Failed to authenticate: no attribute with name “group”: [email]

In case of error like “ **Failed to authenticate: no attribute with name “group”: [email]**“:

[![](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_131544.png)](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_131544.png)

Use the [saml-chrome-panel](https://chrome.google.com/webstore/detail/saml-chrome-panel/paijfdbeoenhembfhkhllainmocckace), and check attributes that are sent in the `AttributeStatement` from the Okta:

[![](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_150903.png)](https://rtfm.co.ua/wp-content/uploads/2021/05/Screenshot_20210514_150903.png)

```
<saml2:AttributeStatement xmlns:saml2="urn:oasis:names:tc:SAML:2.0:assertion">
    <saml2:Attribute Name="email" NameFormat="urn:oasis:names:tc:SAML:2.0:attrname-format:basic">
        <saml2:AttributeValue xmlns:xs="http://www.w3.org/2001/XMLSchema"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="xs:string">arseniy@example.com</saml2:AttributeValue>
    </saml2:Attribute>
</saml2:AttributeStatement>
```

In this example, I’ve made a misprint in the Okta config and it doesn’t send the `<saml2:Attribute Name="group"` attribute.

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