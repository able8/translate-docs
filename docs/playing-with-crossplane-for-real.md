# Playing with Crossplane, for real

From https://prune998.medium.com/playing-with-crossplane-for-real-f591e66065ae 

Aug 25, 2022



![img](https://miro.medium.com/max/1400/1*c8jhmTqn9FGm7vFyZywWmA.png)

It’s 2022 ! We all know about the necessity of creating your Infra AS CODE. I guess we can all agree that [Terraform (TF)](https://www.hashicorp.com/products/terraform) is the leader in this field.

that said, TF is far from easy to use. This is actually not the fault of TF itself (or its parent company, [Hashicorp](https://www.hashicorp.com/)), but to the fact that each provider is so specific that it’s impossible to build up something clean out of it.

So we end up with complicated code, tricks to make things happen or not (like using `count`) and the two biggest pain point:

- having to manage a state file which holds the result of the last apply
- globally not possible to re-use code or build simple abstraction to create a  resource in different clouds, like a K8s cluster in AWS and Google.

The global consensus is that TF is too complicated for Devs team to use and maintain well. It’s down to the Ops to operate it, or at least, try to  provide `modules` so Devs can work with what is supposed to be an abstraction.

## Then enters [Crossplane](https://crossplane.github.io/) !

Before diving into Crossplane, let’s put some notice, or as we like to call it:

# TL ; DR

- I’m no Crossplane expert and only spent few days playing with it for a  Proof Of Concept. Please, help me better understand/use it if you think  i’m wrong
- OSS Crossplane Providers are really limited
- Jet Providers are on par with Terraform providers, but lack docs and may be buggy and are not well supported or updated by the community
- Still unsure how to replicate 100% of what you can do with TF

**Overall conclusion is that it’s not production ready for me**

Please follow up for details

# Crossplane

As stated on their website: **The cloud native control plane framework**

Saying that Crossplane is a K8s native way to do IaC would be limitative.  Crossplane is far more than that. Crossplane, to me, is a way to build  simple abstractions in front of complex stuff, like infrastructure or  deployments.

As TF, Crossplane is build up on [**Providers**](https://github.com/crossplane-contrib) ([see official list here](https://crossplane.github.io/docs/v1.9/concepts/providers.html)), which extends Crossplane with new *Managed Resource Type* to work with.

But the strength of Crossplane lies in its `Composition` feature. It is exactly what the name implies: you build up a new type  of resource by combining (composing) other type of resources. Ex: create a K8s cluster and a nodepool, or create a DB Instance, a DB schema and a DB user, all at one.

Let’s grab some pictures from the official docs:

![https://crossplane.github.io/docs/v1.9/concepts/composition.html#overview](https://miro.medium.com/max/1400/1*RotpDcpSi0UuYq28IdD35g.png)

You `claim` a Postresql Instance, that references a `Composite Resource` that will trigger the creation of a `CloudSQL Instance` and a `Firewall Rule` to access it. Neat !

It’s a little bit more complicated, so here’s another picture from the doc that is supposed to be closer to reality:

![https://crossplane.github.io/docs/v1.9/concepts/composition.html#how-it-works](https://miro.medium.com/max/1400/1*8fN6cZ0OX_xI--8UmtRtIA.png)

AH ! There we go !

You can go read the docs at [https://crossplane.github.io/docs/v1.9/concepts/composition.html](https://crossplane.github.io/docs/v1.9/concepts/composition.html#overview). I personally read that again and again, and wasn’t able to fully  understand the real thing until I played with it, and built my own  schema:

![img](https://miro.medium.com/max/1400/1*S-r3Y4mhb6W9TyNUHQbnFQ.png)

Let me break things down:

1. **Infra**: Deploy a Provider, like GKE, AWS, Helm, K8s, even a [Terraform provider](https://github.com/crossplane-contrib/provider-terraform)
2. The provider created a set of CRDs corresponding to each Cloud resource it manages (yellow boxes)
3. **Infra**: Create a CompositeResourceDefinition (XRD) which creates an interface with a limited set of parameters to tweak
4. Crossplane will create and maintain two new CRDs based on the XRD: a Claim and a  CompositeResource (XR) (green boxes). Crossplane will start watching and reconciling CR based on those CRDs
5. **Infra**: Create a Composition, which will reference a source XRD and a list of  Resources to created (from the CRDs created by the Provider). It’s a  sort of templating resources with values from the interface (the XRD)
6. **Dev**: Claim a resource (purple box) -> a Claim is actually a CustomResource of a type maintained by Crossplane
7. Crossplane will create a CompositeResource (XR) based on the Claim
8. Crossplane will create CustomResources (CR) which are instances of the Provider’s  resources, based on the content of the CompositeResource (XR) (red  boxes)
9. Provider will reconcile the resources he manages, and call GCP API (in case of  the GCP provider) to create the resources declared in the CR

This is really powerful, and the only limitation is actually in what a provider can do.

Talking of which, I guess you see me coming, it’s also the biggest problem  Crossplane has: it all depends on what a Provider can do !

Oh, by the way, I’ll be a speaker at [KubeCon North America 2022 in Detroit](https://events.linuxfoundation.org/kubecon-cloudnativecon-north-america/) , please [check my talk here](https://events.linuxfoundation.org/kubecon-cloudnativecon-north-america/program/schedule/), register for Thursday, October 27 • 2:30pm — 4:00pm:

[Tutorial:  Set Up Your Shell For Kubernetes Productivity And Be Efficient Quickly —  Sebastien “Prune” Thomas, Wunderkind & Archy Ayrat Khayretdinov,  Google](https://kccncna2022.sched.com/event/182F7/tutorial-set-up-your-shell-for-kubernetes-productivity-and-be-efficient-quickly-sebastien-prune-thomas-wunderkind-archy-ayrat-khayretdinov-google)

![img](https://miro.medium.com/max/1400/1*MK4vEEFIs4PEHg4BAMGbOQ.png)

# Question 1: Providers

As stated, `Providers` are the part that manages some resources. In fact, it’s a `Pod` that will be deployed along Crossplane, will create CRD for the  resource it manages and watch on them. Each instance of a managed  resource will make the `Provider` create and maintain a resource. For the GCP provider, that means calling the GCP API and create some resources.

I’m not diving into `Providers`configuration here. It involves installing the provider, giving it a global config  and some more specific variations of the config, like having a `Dev` , `Staging` and `Prod`provider for GCP. I would maybe have to create a mix of configurations for each of my `env` * `project` so I can ensure some devs can deploy to `Project A` in `Dev` but not on `Prod` or in `Project B` . This is something I haven’t explored yet… Please, again, comment.

You can usually check for Provider’s CRD (supported Resources) at https://doc.crds.dev/.
Ex: 
\- crossplane own CRDs: https://doc.crds.dev/github.com/crossplane/crossplane@v1.9.0
\- crossplane GCP Provider (official): https://doc.crds.dev/github.com/crossplane/provider-gcp@v0.21.0

Let’s dive in the GCP provider:

![img](https://miro.medium.com/max/1400/1*0e89L90VOy-KlB0WYNlRpw.png)

28 CRDs discovered ? wait, what ? only 28 different GCP resources are managed by Crossplane ?

short answer: YES :(

So, for example, you can create a `CloudSQLInstance` which represents a Postgres or MySQL Database Instance:

![img](https://miro.medium.com/max/1400/1*iJ-uhzAQA4ycAjJs_AG7jw.png)

But then, you just can’t create any specific DB, User or anything else  related to this DB Instance. Because, well, the resources were not  ported into the Provider.

AWS official Provider is a little better, with 172 resources.

This situation is just normal. Building up a provider is tedious, there are a lot of different resources, and Crossplane team is quite small compared to this huge work.
At first, I guess the Cloud providers wanted to  help, but it feels that in the end, they prefered build their own  version of Crossplane for themselves. That is when Google created it [Config Connector](https://cloud.google.com/config-connector/docs/overview) app, which in term is a lightweight version of Crossplane, or [AWS ACK](https://aws.amazon.com/blogs/containers/aws-controllers-for-kubernetes-ack/).

But Crossplane team is smart, so they asked `who else in the world is already maintaining Providers ?` and the answer was, as often… Terraform !

And the [TerraJet](https://github.com/crossplane/terrajet) project was born !

TerraJet (Jet) is a way to convert Terraform Providers into Crossplane  Providers, hiding the TF mechanics. With Jet, TF is run in the  background and the TF State file is, let’s say, split and stored along  each Crossplane Resource.

## *Smart ?*

Well, the GCP Jet Provider counts 438 resources, including all that we need  to manage SQL DBs in GCP. AWS Jet Provider counts a wooping 780  resources ! I guess it’s more than you’ll ever use.

![img](https://miro.medium.com/max/1400/1*Arl6hA77k5MWXpXImMayjg.png)

# Question 2: Docs ?

It seems to be a lot of resources to use ! While TF docs are, well, not  that bad in the end, and you have a LOT of blog posts and examples to  play with, it’s not the same story for Crossplane.

The [Composition Docs](https://crossplane.github.io/docs/v1.9/concepts/composition.html), for example, showcase few different stuffs that will work out of the  box if you copy/paste. But then, you want your own stuff, maybe not a DB instance. So you start digging in… and reach the [Composition Reference doc](https://crossplane.github.io/docs/v1.9/reference/composition.html). 
Once again, you’re driven through the same example, with little more detail  that will just confuse you even more (at least it confused me):

- Should I create Claims or XR ?
- what is this `compositionRef` about ?
- where should I set `writeConnectionSecretToRef` ? in the Claim ? The XR ? the resource in the Composition ?
- What are the `secret values` that the provider is returning ?

Well, I guess the product is new, and not widely used yet, but man, EVERY  doc, blog, talk, demo is using the same stuff. They all create a single  bucket, [a single DB instance, a VPC with few subnets](https://crossplane.github.io/docs/v1.9/getting-started/create-configuration.html#create-compositions), and in a rare case, [some K8s clusters](https://github.com/upbound/platform-ref-multi-k8s).

To me, nothing that related to my needs.

And because what I want is not (yet?) supported in the official providers, I have to use the GCP Jet Provider.

# Jet Providers, straight from Terraform

Don’t expect to find docs on the Jet Providers… because they are transpositions from TF providers, the doc is minimal, usually [a copy of the TF doc itself](https://registry.terraform.io/providers/hashicorp/google/latest/docs), if any.

For my POC, I wanted something simple, and luckily, something that looks  like what most of the examples are based on: a Postgres DB (CloudSQL).  Of course, I want the Instance, plus a DB, plus few Users, hopefully  using GCP WorkloadIdentity (IAM ServiceAccount binding).

Note that this is a really limited need so far. Only 3 resourceTypes are at  play here. The real use-case would have been creating a Pub/Sub, a  ComposerV2, some cloud-functions, and all the IAM stuff that allow one  to trigger the other. All that at once, by Claiming ONE `Composition` .

## XRD

I started creating a XRD:

```
apiVersion: apiextensions.crossplane.io/v1
kind: CompositeResourceDefinition
metadata:
  name: xjetpostgresqls.database.wk
spec:
  group: database.wk
  names:
    kind: XJetPostgreSQL
    plural: xjetpostgresqls
  claimNames:
    kind: JetPostgreSQL
    plural: jetpostgresqls
  versions:
  - name: v1alpha1
    served: true
    referenceable: true
    schema:
      openAPIV3Schema:
        type: object
        properties:
          spec:
            type: object
            properties:
              parameters:
                type: object
                properties:
                  storageGB:
                    type: integer
                    description: size of the Database in GB - integer
                  dbName:
                    type: string
                    description: name of the new DB inside the DB instance - string
                  instanceSize:
                    type: string
                    description: instance size - string
                    enum:
                      - small
                      - medium
                      - large
                required:
                  - storageGB
                  - dbName
                  - instanceSize
            required:
              - parameters
```

Here I can already see different problems. I wanted to create a DB, and a DB user with the same name, but what if I want 2 DBs ? 3 DBs? 2 users per  DB ?

I guess I have to re-write the schema to use something like:

```
properties:
  dbs:
    type: array
    items:
      type: object
      properties:
        name:
          type: string
          description: name of the new DB inside the DB instance - string
        users:
          type: array
          items:
             properties:
               name:
                 type: string
```

Whatever, modelling those is not straightforward… but, well, it’s a one time  effort. Totally worth it. Take your time building this, as it defines  the parameters that your Dev team will use to create resources. Whatever is not defined here will use the defaults from the composition or the  Provider. You’re creating your abstraction.

## Compositions

Now we can create the Composition, which will take the values defined by the XRD and apply them to the Cloud Resources we need.

```
apiVersion: apiextensions.crossplane.io/v1
kind: Composition
metadata:
  name: jetpostgresql.gcp.database.wk
  labels:
    provider: gcp
    crossplane.io/xrd: xjetpostgresql.database.wk
spec:
  # should I set this here ? Please help
  # writeConnectionSecretsToNamespace: crossplane
  compositeTypeRef:
    apiVersion: database.wk/v1alpha1
    kind: XJetPostgreSQL
  resources:
    - name: cloudsqlinstance
      base:
        apiVersion: sql.gcp.jet.crossplane.io/v1alpha2
        kind: DatabaseInstance
        metadata:
          annotations: 
            crossplane.io/external-name: "crossplanesqlinstance"
        spec:
          providerConfigRef:
            name: crossplane-provider-jet-gcp
          deletionPolicy: Delete
          forProvider:
            databaseVersion: POSTGRES_14
            region: us-central1
            deletionProtection: false
            settings:
            - tier: db-custom-1-3840
              diskType: PD_SSD
              diskSize: 20
              ipConfiguration:
                - ipv4Enabled: true
                  authorizedNetworks:
                    - value: "0.0.0.0/0"
            userLabels:
              creator: crossplane
              owner: prune
          writeConnectionSecretToRef:
            namespace: crossplane
            name: cloudsqlinstance
      patches:
        # set diskSize based on the Claim
        - fromFieldPath: "spec.parameters.storageGB"
          toFieldPath: "spec.forProvider.settings[0].diskSize"
        # set the secret name to the claim name
        - fromFieldPath: "metadata.labels[crossplane.io/claim-name]"
          toFieldPath: "spec.writeConnectionSecretToRef.name"
          transforms:
            - type: string
              string:
                fmt: "%s-pginstance"
        # change secret namespace to the one of the claim
        - fromFieldPath: "metadata.labels[crossplane.io/claim-namespace]"
          toFieldPath: "spec.writeConnectionSecretToRef.namespace"
        # set label app = name of the original claim
        - fromFieldPath: "metadata.labels[crossplane.io/claim-name]"
          toFieldPath: "metadata.labels[crossplane.io/app]"
        # set the name of the external resource to be the name of the claim
        - fromFieldPath: "metadata.labels[crossplane.io/claim-name]"
          toFieldPath: "metadata.annotations[crossplane.io/external-name]"
        # set instance size to the one defined in the claim
        - fromFieldPath: "spec.parameters.instanceSize"
          toFieldPath: "spec.forProvider.settings[0].tier"
          transforms:
            - type: map
              map:
                small: db-custom-1-3840
                medium: db-custom-2-7680
                large: db-custom-4-15360
          policy:
            fromFieldPath: Required
    - name: cloudsqldb
      base:
        apiVersion: sql.gcp.jet.crossplane.io/v1alpha2
        kind: Database
        metadata:
          annotations: 
            crossplane.io/external-name: "crossplanesqldb"
        spec:
          providerConfigRef:
            name: crossplane-provider-jet-gcp
          deletionPolicy: Delete
          forProvider:
            instanceSelector:
              MatchControllerRef: true
          writeConnectionSecretToRef:
            namespace: crossplane
            name: cloudsqldb
      patches:
        # set the secret name to the claim name
        - fromFieldPath: "metadata.labels[crossplane.io/claim-name]"
          toFieldPath: "spec.writeConnectionSecretToRef.name"
          transforms:
            - type: string
              string:
                fmt: "%s-pgdb"
        # change secret namespace to the one of the claim
        - fromFieldPath: "metadata.labels[crossplane.io/claim-namespace]"
          toFieldPath: "spec.writeConnectionSecretToRef.namespace"
        # set the name of the DB resource to be the name defined in the claim
        - fromFieldPath: "spec.parameters.dbName"
          toFieldPath: "metadata.annotations[crossplane.io/external-name]"
        # set app Label
        - fromFieldPath: "metadata.labels[crossplane.io/claim-name]"
          toFieldPath: "metadata.labels[crossplane.io/app]"
    - name: cloudsqldbuser
      base:
        apiVersion: sql.gcp.jet.crossplane.io/v1alpha2
        kind: User
        metadata:
          annotations: 
            # set the name of the DB User, this is hardcoded for demo but should come from the CRD
            crossplane.io/external-name: "existing-sa-for-db@my-project.iam"
        spec:
          providerConfigRef:
            name: crossplane-provider-jet-gcp
          deletionPolicy: Delete
          forProvider:
            instanceSelector:
              MatchControllerRef: true
            type: CLOUD_IAM_SERVICE_ACCOUNT
          writeConnectionSecretToRef:
            namespace: crossplane
            name: cloudsqluser
      patches:
        # set the secret name to the claim name
        - fromFieldPath: "metadata.labels[crossplane.io/claim-name]"
          toFieldPath: "spec.writeConnectionSecretToRef.name"
          transforms:
            - type: string
              string:
                fmt: "%s-pguser"
        # change secret namespace to the one of the claim
        - fromFieldPath: "metadata.labels[crossplane.io/claim-namespace]"
          toFieldPath: "spec.writeConnectionSecretToRef.namespace"
        # set the name of the DB User, this is hardcoded for demo but should come from the Claim CRD
        # - fromFieldPath: "spec.parameters.dbName"
        #   toFieldPath: "metadata.annotations[crossplane.io/external-name]"
        # set app Label
        - fromFieldPath: "metadata.labels[crossplane.io/claim-name]"
          toFieldPath: "metadata.labels[crossplane.io/app]"
```

Here you can see that it’s not trivial either. A lot is going on here.  Globally, you just list all the Provider’s Resources you want to  instanciate, give default parameters you want to enforce, and patch some others from user-supplied values or other resources values.

Let’s go back to the issue I had describing the DBs and Users: Let’s say I  have a list of DBs, and each DB has a list of Users… Then I need to  iterate over the values and create as many `Database.sql.gcp.jet.crossplane.io` and `User.sql.gcp.jet.crossplane.io` that is in the list. I would achieve that in TF using a `for_each` but here, in Crossplane… no idea (Help me if you know how).

You can also see that each resource has `writeConnectionSecretToRef` which points to a secret `name` and `namespace` which in turn, should hold all the sensible values that the Resource  may create (like DB password, SSL certs, maybe even URLs, and whatever  you want to add into the secret).

I’m still unsure if it’s the right way of doing it. Most examples are writing the secrets in the `crossplane` namespace, or at least in a namespace fixed in the composition… This  does not make sense to me. I want the secret to be in the same `namespace` as the Developer’s claim that is creating the resource… This is making the `writeConnectionSecretsToNamespace` value that is set at the `Composition` `Spec` level (I commented it out).

I guess it’s the moment to tell you about the Claims.

# Claims

The claim is the easiest, at first, as it’s the abstraction: it is really limited.
It’s meant to be created by your developers that will want a DB along their  apps, or a Pub/Sub, or whatever resource you want them to have full  control on.

It is the only resource that is *namespaced*. This is made so you can define RBACs and allow this team (this namespace) to hold DBs, but no Pub/Sub, for example.

Here’s mine:

```
apiVersion: database.wk/v1alpha1
kind: JetPostgreSQL
metadata:
  namespace: test-namespace
  name: jet-db-claim
spec:
  parameters:
    storageGB: 25
    dbName: xrdb
    instanceSize: small # small, medium, large
  writeConnectionSecretToRef:
    name: jet-db-claim-details
```

Neat ! Simple ! Thanks abstractions ! Devs only specify what they care about, and you take care of all the boring stuff !

Going back to Secrets, you see that I also defined which secret to write  stuff into… but it seems this value is not an override of what is in the `Composition`. So I patched the composition to actually create the secret with a name derived from the `Claim` name and in the same `Namespace` as the `Claim` :

```
# set the secret name to reference the claim name
- fromFieldPath: "metadata.labels[crossplane.io/claim-name]"
  toFieldPath: "spec.writeConnectionSecretToRef.name"
  transforms:
    - type: string
      string:
        fmt: "%s-pginstance"
# change secret namespace to the one of the claim
- fromFieldPath: "metadata.labels[crossplane.io/claim-namespace]"
  toFieldPath: "spec.writeConnectionSecretToRef.namespace"
```

I first tried to use `metadata.name` to get the claim name, but, in fact, the `Composition` is not templated from the `Claim` but from the intermediate `XR` that is created from the `Claim` . Refer to my schema above if needed. So, in the `XR` the only way to get back to the `Claim` metadata is by looking at some specific `labels` like `crossplane.io/claim-name`.

# Question 3: Provider’s execution

Once you have all that, the Providers are going to start doing things:  reconciling between what you created in K8s and the reality in the  Cloud.

For that, the official OSS Providers seems to do a great work without much surprise. 
For the Jet provider, at least the GCP one I tested, it’s another story. I started mine with `--debug` to better understand.
It’s k8s native, so also `k describe` on the many resources (XRD, Composition, XR, Claim and intermediate CR) to get the logs and events attached to each.

I would also suggest that you add the `--debug` option to Crossplane itself to fully understand when and where the problem is when rendering the Composition.

Crossplane keeps reconcilling the resource, so you may end up with tons of logs.

In the end, I wasn’t able to create a working DB user using Workload  Identity. Not sure why, but once the user created, Crossplane was trying to change the `username` to the name of the claim… and I was constantly seeing the `crossplane.io/external-name` value being switched from what I asked in the Composition and the Claim name…

I also wasn’t able to understand what kind of values was the Provider /  the Resource returning that I could put in a Secret. I think this is a  difference between the OSS Provider and the Jet Provider… with Jet, it’s like with TF: YOU have to create the secret and provide it to K8s… so  maybe I only used resources that had nothing to return as a secret value ?

# Question 4: Support

I asked many questions on Slack, I even opened an issue… and globally I  had no answer. I also tried to participate and read through the other  questions/bugs/PRs… and it does not feel the community is heavily  active. Don’t get me wrong, I’m not saying the project is dead or  whatever. Its more a feeling… Ask about TF and you’ll have 15 answers in the hour…

I talked about docs and examples, and again it’s really limited at the moment.

If i’m about to switch my whole IaC in a new tool, I want to be sure that it’s active and reactive.

I just reached to [Upbound](https://www.upbound.io/), the company behind Crossplane, to see if the paid plan and support would help me solve my issues. 
I’m really pleased to see that I had an answer quickly along with an  appointment with an engineer. According to what I was told, the OSS  Providers (and Jet) are way behind what Upbound offers to paid  customers. Maybe it’s just the way to go ? After all, TF also have a  paid subscription, which is far from cheap… and if I put all my IaC  somewhere, maybe a little support is a good idea.

# Conclusion

My POC was really limited in time, and there’s so much more I wish I had time to cover…

My conclusion is that I could use Crossplane for my Developer’s infra, IF the need is really small…

For example, if you only create one DB and one User per DB Instance.

Another way of doing would certainly be to not Compose multiple resources in one `Composition` but have one `Composition` per resource… we I will end up with one `MyDBInstance` , one `MyDB`, two or more `MyDbUser` in each namespace/dev project… I’m even not sure if this model would  work… but at least, it would allow me to build an abstraction with  limited levers that I can present to the Dev’s teams while still  allowing flexibility.

At the moment I would say that this POC just showed me that Crossplane is not an out-of-the-box easy replacement to TF.

I will keep playing with Crossplane on the side, until I fully understand it. I will also investigate the GCP official tool.