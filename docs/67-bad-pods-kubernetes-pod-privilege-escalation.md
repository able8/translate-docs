# Bad Pods: Kubernetes Pod Privilege Escalation

on Jan 19, 2021 5:26:38 AM

What are the risks associated with overly permissive pod creation in Kubernetes? The answer varies based on which of the host’s namespaces and security contexts are allowed. In this post, I will describe eight insecure pod configurations and the corresponding methods to perform privilege escalation. This article and the accompanying [repository](https://github.com/BishopFox/badPods) were created to help penetration testers and administrators better understand common misconfiguration scenarios.

If you are an administrator, I hope that this post gives you the confidence to apply restrictive controls around pod creation by default. I also hope it helps you consider isolating any pods that need access to the host’s resources to a namespace that is only accessible to administrators using the principle of least privilege.

If you are a penetration tester, I hope this post provides you with some ideas on how to demonstrate the impact of an overly permissive pod security policy. And I hope that the repository gives you some easy-to-use manifests and actionable steps to achieve those goals.

**Executive Summary:**

One of the foundations of information security is the "principal of least privilege." This means that every user, system process, or application needs to operate using the least set of privileges required to do a task. When privileges are configured where they greatly exceed what is required, attackers can take advantage of these situations to access sensitive data, compromise systems, or escalate those privileges to conduct lateral movement in a network.

Kubernetes and other new "DevOps" technologies are complex to implement properly and are often deployed misconfigured or configured with more permissions than necessary. The lesson, as we have demonstrated from our "Bad Pods" research, is that if you are using Kubernetes in your infrastructure, you need to find out from your development team how they are configuring and hardening this environment.

## HARDENING PODS: HOW RISKY CAN A SINGLE ATTRIBUTE BE?

When it comes to Kubernetes security best practices, every checklist worth its salt mentions that you want to use the principle of least privilege when provisioning pods. But how can we enforce granular security controls and how do we evaluate the risk of each attribute?

A Kubernetes administrator can enforce the principle of least privilege using admission controllers. For example, there’s a built-in Kubernetes controller called [PodSecurityPolicy](https://kubernetes.io/docs/concepts/policy/pod-security-policy/) and also a popular third-party admission controller called [OPA Gatekeeper](https://github.com/open-policy-agent/gatekeeper). Admission controllers allow you to deny a pod entry into the cluster if it has more permissions than the policy allows.

However, even though the controls exist to define and enforce policy, the real-world security implications of allowing each specific attribute is not always understood, and quite often, pod creation is not as locked down as it needs to be.

As a penetration tester, you might find yourself with access to create pods on a cluster where there is no policy enforcement. This is what I like to refer to as “easy mode.” Use [this manifest](https://raesene.github.io/blog/2019/04/01/The-most-pointless-kubernetes-command-ever/) from Rory McCune ( [@raesene](https://twitter.com/raesene)), [this command](https://twitter.com/mauilion/status/1129468485480751104) from Duffie Cooley ( [@mauilion](https://twitter.com/mauilion)), or the [node-shell](https://github.com/kvaps/kubectl-node-shell) krew plugin and you will have fully interactive privileged code execution on the underlying host. It doesn’t get easier than that!

But what if you can create a pod with just? What can you do in each case? Let’s take a look!

![Kubernetes bad pods](https://labs.bishopfox.com/hs-fs/hubfs/Kubernetes%20-%20PAD%20PODS%20-%20300%20PPI%20-%20Option%2001.png?width=700&name=Kubernetes%20-%20PAD%20PODS%20-%20300%20PPI%20-%20Option%2001.png)

## BAD PODS - ATTRIBUTES AND THEIR WORST-CASE SECURITY IMPACT

The pods below are loosely ordered from highest to lowest security impact. Note that the generic attack paths that could affect any Kubernetes pod (e.g., checking to see if the pod can access the cloud provider’s metadata service or identifying misconfigured Kubernetes RBAC) are covered in [Bad Pod #8: Nothing allowed](http://labs.bishopfox.com#pod8).

## THE BAD PODS LINEUP

**Pods**

[Bad Pod #1: Everything allowed](http://labs.bishopfox.com#Pod1)

[Bad Pod #2: Privileged and hostPid](http://labs.bishopfox.com#Pod2)

[Bad Pod #3: Privileged only](http://labs.bishopfox.com#Pod3)

[Bad Pod #4: hostPath only](http://labs.bishopfox.com#Pod4)

[Bad Pod #5: hostPid only](http://labs.bishopfox.com#Pod5)

[Bad Pod #6: hostNetwork only](http://labs.bishopfox.com#pod6)

[Bad Pod #7: hostIPC only](http://labs.bishopfox.com#pod7)

[Bad Pod #8: Nothing allowed](http://labs.bishopfox.com#pod8)

## BAD POD \#1: EVERYTHING ALLOWED

![BAD POD #1: EVERYTHING ALLOWED](https://labs.bishopfox.com/hs-fs/hubfs/Check%20Box%20-%20Door%20Idea%20FINAL%2001%20-%20300PPI%20(1).jpg?width=700&name=Check%20Box%20-%20Door%20Idea%20FINAL%2001%20-%20300PPI%20(1).jpg)

### What’s the worst that can happen?

Multiple paths to full cluster compromise

### How?

The pod you create mounts the host’s filesystem to the pod. You’ll have the best luck if you can schedule your pod on a control-plane node using the [nodeName](https://kubernetes.io/docs/tasks/configure-pod-container/assign-pods-nodes/#create-a-pod-that-gets-scheduled-to-specific-node) selector in your manifest. You then `exec` into your pod and `chroot` to the directory where you mounted the host’s filesystem. You now have root on the node running your pod.

- **Read secrets from `etcd`** — If you can run your pod on a control-plane node using the `nodeName` selector in the pod spec, you might have easy access to the `etcd` database, which contains the configuration for the cluster, including all secrets.
- **Hunt for privileged service account tokens**— Even if you can only schedule your pod on the worker node, you can also access any secret mounted within any pod on the node you are on. In a production cluster, even on a worker node, there is usually at least one pod that has a mounted token that is bound to a service account that is bound to a clusterrolebinding, which gives you access to do things like create pods or view secrets in all namespaces

Some additional privilege escalation patterns are outlined in the [README](https://github.com/BishopFox/badPods/tree/main/manifests/everything-allowed) document linked below and also in [Bad Pod #4: hostPath](https://labs.bishopfox.com/tech-blog/bad-pods-kubernetes-pod-privilege-escalation?hs_preview=qrRGstOq-40671699438#Pod4).

### Usage and exploitation examples

[https://github.com/BishopFox/badPods/tree/main/manifests/everything-allowed](https://github.com/BishopFox/badPods/tree/main/manifests/everything-allowed)

### References and further reading

- [The Most Pointless Kubernetes Command Ever](https://raesene.github.io/blog/2019/04/01/The-most-pointless-kubernetes-command-ever/)
- [Secure Kubernetes - KubeCon NA 2019 CTF](https://securekubernetes.com/)
- [Deep Dive into Real-World Kubernetes Threats](https://research.nccgroup.com/2020/02/12/command-and-kubectl-talk-follow-up/)
- [Compromising Kubernetes Cluster by Exploiting RBAC Permissions](https://www.youtube.com/watch?v=1LMo0CftVC4) ( [slides](https://published-prd.lanyonevents.com/published/rsaus20/sessionsFiles/18100/2020_USA20_DSO-W01_01_Compromising%20Kubernetes%20Cluster%20by%20Exploiting%20RBAC%20Permissions.pdf))
- [The Path Less Traveled: Abusing Kubernetes Defaults](https://www.youtube.com/watch?v=HmoVSmTIOxM) & [Corresponding Repo](https://github.com/mauilion/blackhat-2019)

## BAD POD \#2: PRIVILEGED AND HOSTPID

![BAD POD #2: PRIVILEGED AND HOSTPID](https://labs.bishopfox.com/hs-fs/hubfs/Check%20Box%20-%20Door%20Idea%20FINAL%2002%20-%20300PPI%20(1).jpg?width=700&name=Check%20Box%20-%20Door%20Idea%20FINAL%2002%20-%20300PPI%20(1).jpg)

### What’s the worst that can happen?

Multiple paths to full cluster compromise

### How?

In this scenario, the only thing that changes from the everything-allowed pod is how you gain root access to the host. Rather than `chroot`ing to the host’s filesystem, you can use `nsenter` to get a root shell on the node running your pod.

**Why does it work?**

- **Privileged** — The `privileged: true` container-level security context breaks down almost all the walls that containers are supposed to provide; however, the PID namespace is one of the few walls that stands. Without `hostPID`, `nsenter` would only work to enter the namespaces of a process running within the container. For more examples on what you can do if you only have `privileged: true`, refer to the next example [Bad Pod #3: Privileged Only](http://labs.bishopfox.com#Pod3).

- **Privileged + `hostPID`** — When both `hostPID: true` and `privileged: true` are set, the pod can see all of the processes on the host, and you can enter the `init` system (PID 1) on the host. From there, you can execute your shell on the node.


Once you are root on the host, the privilege escalation paths are all the same as described in [Bad Pod # 1: Everything-allowed](http://labs.bishopfox.com#Pod1).

### Usage and exploitation examples

[https://github.com/BishopFox/badPods/tree/main/manifests/priv-and-hostpid](https://github.com/BishopFox/badPods/tree/main/manifests/priv-and-hostpid)

### References and further reading

- [Duffie Cooley's Nsenter Pod Tweet](https://twitter.com/mauilion/status/1129468485480751104)
- [The Path Less Traveled: Abusing Kubernetes Defaults](https://www.youtube.com/watch?v=HmoVSmTIOxM) & [Corresponding Repo](https://github.com/mauilion/blackhat-2019)
- [Node-shell Krew Plugin](https://github.com/kvaps/kubectl-node-shell)

## BAD POD \#3: PRIVILEGED ONLY

![BAD POD #3: PRIVILEGED ONLY](https://labs.bishopfox.com/hs-fs/hubfs/Check%20Box%20-%20Door%20Idea%20FINAL%2003%20-%20300PPI%20(1).jpg?width=700&name=Check%20Box%20-%20Door%20Idea%20FINAL%2003%20-%20300PPI%20(1).jpg)

### What’s the worst that can happen?

Multiple paths to full cluster compromise

### How?

If you only have `privileged: true`, there are two paths you can take:

- **Mount the host’s filesystem** — In privileged mode, `/dev` on the host is accessible in your pod. You can mount the disk that contains the host’s filesystem into your pod using the `mount` command. In my experience, this gives you a limited view of the filesystem though. Some files, and therefore privesc paths, are not accessible from your privileged pod unless you escalate to a full shell on the node. That said, it is easy enough that you might as well mount the device and see what you can see.

- **Exploit `cgroup`** **user mode helper programs**— Your best bet is to get interactive root access on the node, but you must jump through a few hoops first. You can use Felix Wilhelm's exploit PoC [undock.sh](https://twitter.com/_fel1x/status/1151487051986087936) to execute one command a time, or you can use Brandon Edwards and Nick Freeman’s version from their talk [A Compendium of Container Escapes](https://www.youtube.com/watch?v=BQlqita2D2s), which forces the host to connect back to the listener on the pod for an easy upgrade to interactive root access on the host. Another option is to use the Metasploit module [Docker Privileged Container Escape](https://www.rapid7.com/db/modules/exploit/linux/local/docker_privileged_container_escape/), which uses the same exploit to upgrade a shell received from a container to a shell on the host.


Whichever option you choose, the Kubernetes privilege escalation paths are largely the same as the [Bad Pod #1: Everything-allowed](http://labs.bishopfox.com#Pod1).

### Usage and exploitation examples

[https://github.com/BishopFox/badPods/tree/main/manifests/priv](https://github.com/BishopFox/badPods/tree/main/manifests/priv)

### References and further reading

- [Felix Wilhelm's Cgroup Usermode Helper Exploit](https://twitter.com/_fel1x/status/1151487051986087936)
- [Understanding Docker Container Escapes](https://blog.trailofbits.com/2019/07/19/understanding-docker-container-escapes/)
- [A Compendium of Container Escapes](https://www.youtube.com/watch?v=BQlqita2D2s)
- [Docker Privileged Container Escape Metasploit Module](https://www.rapid7.com/db/modules/exploit/linux/local/docker_privileged_container_escape/)

## BAD POD \#4: HOSTPATH ONLY

![BAD POD #4: HOSTPATH ONLY](https://labs.bishopfox.com/hs-fs/hubfs/Check%20Box%20-%20Door%20Idea%20FINAL%2004%20-%20300PPI%20(1).jpg?width=700&name=Check%20Box%20-%20Door%20Idea%20FINAL%2004%20-%20300PPI%20(1).jpg)

### What’s the worst that can happen?

Multiple paths to full cluster compromise

### How?

In this case, even if you don’t have access to the host’s process or network namespaces, if the administrators have not limited what you can mount, you can mount the entire host’s filesystem into your pod, giving you read/write access on the host’s filesystem. This allows you to execute most of the same privilege escalation paths outlined above. There are so many paths available that Ian Coldwater and Duffie Cooley gave an awesome Black Hat 2019 talk about it titled “ [The Path Less Traveled: Abusing Kubernetes Defaults!](https://www.youtube.com/watch?v=HmoVSmTIOxM)”

Here are some privileged escalation paths that apply any time you have access to a Kubernetes node’s filesystem:

- **Look for `kubeconfig` files on the host filesystem** — If you are lucky, you will find a `cluster-admin` config with full access to everything.

- **Access the tokens from all pods on the node** — Use something like `kubectl auth can-i --list` or [access-matrix](https://github.com/corneliusweig/rakkess) to see if any of the pods have tokens that give you more permissions than you currently have. Look for tokens that have permissions to get secrets or create pods, deployments, etc., in `kube-system`, or that allow you to create clusterrolebindings.

- **Add your SSH key** — If you have network access to SSH to the node, you can add your public key to the node and SSH to it for full interactive access.

- **Crack hashed passwords** — Crack hashes in `/etc/shadow`; see if you can use them to access other nodes.


### Usage and exploitation examples

[https://github.com/BishopFox/badPods/tree/main/manifests/hostpath](https://github.com/BishopFox/badPods/tree/main/manifests/hostpath)

### References and further reading

- [The Path Less Traveled: Abusing Kubernetes Defaults](https://www.youtube.com/watch?v=HmoVSmTIOxM) & [Corresponding Repo](https://github.com/mauilion/blackhat-2019)
- [Secure Kubernetes - KubeCon NA 2019 CTF](https://securekubernetes.com/)
- [Deep Dive into Real-World Kubernetes Threats](https://research.nccgroup.com/2020/02/12/command-and-kubectl-talk-follow-up/)
- [Compromising Kubernetes Cluster by Exploiting RBAC Permissions](https://www.youtube.com/watch?v=1LMo0CftVC4) ( [slides](https://published-prd.lanyonevents.com/published/rsaus20/sessionsFiles/18100/2020_USA20_DSO-W01_01_Compromising%20Kubernetes%20Cluster%20by%20Exploiting%20RBAC%20Permissions.pdf))

## BAD POD \#5: HOSTPID ONLY

![BAD POD #5: HOSTPID ONLY](https://labs.bishopfox.com/hs-fs/hubfs/Check%20Box%20-%20Door%20Idea%20FINAL%2005%20-%20300PPI%20(1).jpg?width=700&name=Check%20Box%20-%20Door%20Idea%20FINAL%2005%20-%20300PPI%20(1).jpg)

### What’s the worst that can happen?

Application or cluster credential leaks if an application in the cluster is configured incorrectly. Denial of service via process termination.

### How?

There’s no clear path to get root on the node with only `hostPID`, but there are still some good post-exploitation opportunities.

- **View processes on the host** — When you run `ps` from within a pod that has `hostPID: true`, you see all the processes running on the host, including processes running in each pod.

- **Look for passwords, tokens, keys, etc.** — If you are lucky, you will find credentials and you can then use them to escalate privileges in the cluster, to escalate privileges to services supported by the cluster, or to escalate privileges to services that communicate with cluster-hosted applications. It’s a long shot, but you might find a Kubernetes service account token or some other authentication material that will allow you to access other namespaces and eventually escalate all the way to cluster admin.

- **Kill processes** — You can also kill any process on the node (presenting a denial-of-service risk). Because of this risk though, I would advise against it on a penetration test!


### Usage and exploitation examples

[https://github.com/BishopFox/badPods/tree/main/manifests/hostpid](https://github.com/BishopFox/badPods/tree/main/manifests/hostpid)

## BAD POD \#6: HOSTNETWORK ONLY

![BAD POD #6: HOSTNETWORK ONLY](https://labs.bishopfox.com/hs-fs/hubfs/Check%20Box%20-%20Door%20Idea%20FINAL%2006%20-%20300PPI%20(2).jpg?width=700&name=Check%20Box%20-%20Door%20Idea%20FINAL%2006%20-%20300PPI%20(2).jpg)

### What’s the worst that can happen?

Potential path to cluster compromise

### How?

If you only have `hostNetwork: true`, you can’t get privileged code execution on the host directly, but if you cross your fingers, you might still find a path to cluster admin. There are three potential escalation paths:

- **Sniff traffic** — You can use `tcpdump` to sniff unencrypted traffic on any interface on the host. You might get lucky and find service account tokens or other sensitive information that is transmitted over unencrypted channels.

- **Access services bound to localhost** — You can also reach services that only listen on the host’s loopback interface or that are otherwise blocked by network policies. These services might turn into a fruitful privilege escalation path.

- **Bypass network policy** — If a restrictive network policy is applied to the namespace, deploying a pod with `hostNetwork: true` allows you to bypass the restrictions. This works because you are bound to the host's network interfaces and not the pods.


### Usage and exploitation examples

[https://github.com/BishopFox/badPods/tree/main/manifests/hostnetwork](https://github.com/BishopFox/badPods/tree/main/manifests/hostnetwork)

## BAD POD \#7: HOSTIPC ONLY

![BAD POD #7: HOSTIPC ONLY](https://labs.bishopfox.com/hs-fs/hubfs/Check%20Box%20-%20Door%20Idea%20FINAL%2007%20-%20300PPI%20(1).jpg?width=700&name=Check%20Box%20-%20Door%20Idea%20FINAL%2007%20-%20300PPI%20(1).jpg)

### What’s the worst that can happen?

Ability to access data used by any pods that also use the host’s IPC namespace

### How?

If any process on the host or any processes in a pod uses the host’s inter-process communication mechanisms (shared memory, semaphore arrays, message queues, etc.), you’ll be able to read/write to those same mechanisms. The first place you'll want to look is `/dev/shm`, as it is shared between any pod with `hostIPC: true` and the host. You'll also want to check out the other IPC mechanisms with `ipcs`.

- **Inspect /dev/shm** — Look for any files in this shared memory location.

- **Inspect existing IPC facilities** — You can check to see if any IPC facilities are being used with `/usr/bin/ipcs`.


### Usage and exploitation examples

[https://github.com/BishopFox/badPods/tree/main/manifests/hostipc](https://github.com/BishopFox/badPods/tree/main/manifests/hostipc)

## BAD POD \#8: NOTHING ALLOWED

![BAD POD #8: NOTHING ALLOWED](https://labs.bishopfox.com/hs-fs/hubfs/Check%20Box%20-%20Door%20Idea%20FINAL%2008%20-%20300PPI%20(1).jpg?width=700&name=Check%20Box%20-%20Door%20Idea%20FINAL%2008%20-%20300PPI%20(1).jpg)

### What’s the worst that can happen?

Multiple potential paths to full cluster compromise

### How?

To close our bad Pods lineup, there are plenty of attack paths that should be investigated any time you can create a pod or simply have access to a pod, even if there are no security attributes enabled. Here are some things to look for whenever you have access to a Kubernetes pod:

- **Accessible cloud metadata** — If the pod is cloud hosted, try to access the cloud metadata service. You might get access to the IAM credentials associated with the node or even just find a cloud IAM credential created specifically for that pod. In either case, this can be your path to escalate in the cluster, in the cloud environment, or in both.

- **Overly permissive service accounts** — If the namespace’s default service account is mounted to `/var/run/secrets/kubernetes.io/serviceaccount/token` in your pod and is overly permissive, use that token to further escalate your privileges within the cluster.

- **Misconfigured Kubernetes components** — If either [the apiserver or the kubelets have `anonymous-auth` set to](https://labs.f-secure.com/blog/attacking-kubernetes-through-kubelet/) [`true`](https://labs.f-secure.com/blog/attacking-kubernetes-through-kubelet/) and there are no network policy controls preventing it, you can interact with them directly without authentication.

- **Kernel, container engine, or Kubernetes exploits** — An unpatched exploit in the underlying kernel, in the container engine, or in Kubernetes can potentially allow a container escape or access to the Kubernetes cluster without any additional permissions.

- **Hunt for vulnerable services**— Your pod will likely see a different view of the network services running in the cluster than you can see from the machine you used to create the pod. You can hunt for vulnerable services and applications by proxying your traffic through the pod.


### Usage and exploitation examples

[https://github.com/BishopFox/badPods/tree/main/manifests/nothing-allowed](https://github.com/BishopFox/badPods/tree/main/manifests/nothing-allowed)

### References and further reading

- [Secure Kubernetes - KubeCon NA 2019 CTF](https://securekubernetes.com/)
- [Kubernetes Goat](https://madhuakula.com/kubernetes-goat/)
- [Attacking Kubernetes through Kubelet](https://labs.f-secure.com/blog/attacking-kubernetes-through-kubelet/)
- [Deep Dive into Real-World Kubernetes Threats](https://research.nccgroup.com/2020/02/12/command-and-kubectl-talk-follow-up/)
- [A Compendium of Container Escapes](https://www.youtube.com/watch?v=BQlqita2D2s)
- [CVE-2020-8558 POC](https://github.com/tabbysable/POC-2020-8558)

## CONCLUSION

Apart from the [Bad Pod #8: Nothing Allowed](http://labs.bishopfox.com#pod8) example, all of the privilege escalation paths covered in this blog post (and the [respective repository](https://github.com/BishopFox/badPods)) can be mitigated with restrictive pod security policies.

Additionally, there are many other defense-in-depth security controls available to Kubernetes administrators that can reduce the impact of or completely thwart certain attack paths even when an attacker has access to some or all of the host namespaces and capabilities (e.g., disabling the automatic mounting of service account tokens or requiring all pods to run as non-root by enforcing `MustRunAsNonRoot=true` and `allowPrivilegeEscalation=false`). As is always the case with penetration testing, your mileage may vary.

Administrators are sometimes hard pressed to defend security best practices without examples that demonstrate the security implications of risky configurations. I hope the examples laid out in this post and the manifests contained in the [Bad Pods repository](https://github.com/BishopFox/badPods) help you enforce the principle of least privilege when it comes to Kubernetes pod creation in your organization.