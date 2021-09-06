# NTP in a Kubernetes cluster

Anyone who doesn’t know what is NTP (Network Time Protocol), directly from Wikipedia,

“The Network Time Protocol is a networking protocol for clock synchronization between computer systems over packet-switched, variable-latency data networks.”

In this blog, I am using OpenNTPD which is a FREE implementation of NTP. It provides

- the ability to sync the local clock to remote NTP servers
- can act as NTP server itself, redistributing the local clock.

# Problem Description

I was curious if it is possible to run NTPD using openntpd in the Kubernetes cluster. Found a [handy blog](http://blog.oddbit.com/post/2015-10-09-running-ntp-in-a-container/). In this blog, the author is showing how you can run NTP in a container. So I am going to follow this blog to create a build process for NTPD docker image and use that image to run NTPD in the Kubernetes cluster.

# Solutions

## Test with Docker container

First, create a `Dockerfile` using the following content.

```
FROM alpine
RUN apk update
RUN apk add openntpd
ENTRYPOINT ["ntpd"]

```

Now you can build docker image using the following command

```
docker build -t goglides/ntpd .

```

Let’s test it using the following

```
docker run goglides/ntpd -d

Output:
ntpd: can't set priority: Permission denied
reset adjtime failed: Operation not permitted
creating new /var/db/ntpd.drift
adjtimex failed: Operation not permitted
ntp engine ready
constraint request to 172.217.9.196
constraint request to 2607:f8b0:4004:806::2004
tls connect failed: 2607:f8b0:4004:806::2004 (www.google.com): connect: Address not available
no constraint reply from 2607:f8b0:4004:806::2004 received in time, next query 900s
tls connect failed: 172.217.9.196 (www.google.com): ssl verify memory setup failure
no constraint reply from 172.217.9.196 received in time, next query 900s

```

Apparently, our container is not working as expected, if you look at following errors,

```
ntpd: can't set priority: Permission denied
reset adjtime failed: Operation not permitted
creating new /var/db/ntpd.drift
adjtimex failed: Operation not permitted

```

which is basically saying not enough permission is because docker is not able to run binary, we can fix this by adding `--privileged` flag as follows

```
docker run --privileged goglides/ntpd -d

Output:
creating new /var/db/ntpd.drift
ntp engine ready
constraint request to 2607:f8b0:4004:806::2004
constraint request to 172.217.9.196
tls connect failed: 2607:f8b0:4004:806::2004 (www.google.com): connect: Address not available
no constraint reply from 2607:f8b0:4004:806::2004 received in time, next query 900s
tls connect failed: 172.217.9.196 (www.google.com): ssl verify memory setup failure
no constraint reply from 172.217.9.196 received in time, next query 900s

```

And for SSL issue I keep seeing it, I tried adding following in Dockerfile which didn’t fix the issue,

```
RUN apk add ca-certificates
RUN update-ca-certificates

```

There is an open bug in the official repo, [https://gitlab.alpinelinux.org/alpine/aports/issues/9635](https://gitlab.alpinelinux.org/alpine/aports/issues/9635) which mentioning why we are seeing this issue.

The main cause is mentioned below (copy/pasted from the ticket)

> This happens because of the chroot(2) call in ntpd(8) child. ntpd(8) calls open(“/tmp/libtlscompat\*\*\*”, …) after chrooting to /var/empty and fails because there is no /var/empty/tmp. Doing mkdir -m 1777 /var/empty/tmp helps, but I don’t know if this is the right approach. UPD 1: There is a way to specify the directory where ntpd(8) should chroot with ./configure … –with-privsep-path=, so it would make sense to create /var/lib/openntpd with tmp inside owned by ntp:ntp. UPD 2: Actually, there is no call to open(“/tmp/libtls\*\*\*”, …) on OpenBSD. I guess that this is standalone-only thing. UPD 3: The offending function SSL\_CTX\_load\_verify\_mem resides here. Calling mkstemp(3) from there happens after chrooting and ntpd(8) does not expect that since vanilla libressl does all this in memory:

```
/* From libressl-2.9.1 */

int
SSL_CTX_load_verify_mem(SSL_CTX *ctx, void *buf, int len)
{
    return (X509_STORE_load_mem(ctx->cert_store, buf, len));
}

```

> UPD 4: OpenNTPD does not use –with-privsep-path=in code, it just for hints. Instead it chroots to the home directory of a running user, so it’s necessary to setup a dedicated user (e.g. openntpd) with a home directory pointing to e.g. /var/lib/openntpd.

So solution is to add `mkdir -m 1777 /var/empty/tmp` as workaround for now. My final `Dockerfile` looks like this.

```
FROM alpine:3.11.3
RUN apk update
RUN apk add openntpd
RUN mkdir -m 1777 /var/empty/tmp
ENTRYPOINT ["ntpd"]

```

Now I am seeing following output

```
docker run --privileged goglides/ntpd -d

Output:
creating new /var/db/ntpd.drift
ntp engine ready
constraint request to 172.217.9.196
constraint request to 2607:f8b0:4004:806::2004
tls connect failed: 2607:f8b0:4004:806::2004 (www.google.com): connect: Address not available
no constraint reply from 2607:f8b0:4004:806::2004 received in time, next query 900s
constraint reply from 172.217.9.196: offset -1.288009
reply from 66.228.58.20: offset -0.292780 delay 0.079884, next query 8s
reply from 72.87.88.202: offset -0.299418 delay 0.088550, next query 6s
reply from 138.68.201.49: offset -0.288903 delay 0.137101, next query 9s

```

The next step is to provide a configuration file. I am simply going to create `/tmp/ntpd.conf` conf file with `servers pool.ntp.org` content and mount that to the container. For detail configuration options, please visit [ntpd.conf man page](http://man.openbsd.org/cgi-bin/man.cgi/OpenBSD-current/man5/ntpd.conf.5?query=ntpd.conf)

```
docker run -v /tmp/ntpd.conf:/etc/ntpd.conf --privileged goglides/ntpd -d -f /etc/ntpd.conf

Output:
creating new /var/db/ntpd.drift
ntp engine ready
reply from 69.89.207.99: offset -0.356652 delay 0.122182, next query 33s
reply from 96.8.121.205: offset -0.356918 delay 0.153844, next query 30s
reply from 138.236.128.112: offset -0.346649 delay 0.108738, next query 30s
reply from 96.8.121.205: offset -0.359319 delay 0.153514, next query 34s
reply from 204.11.201.10: offset -0.359962 delay 0.152213, next query 32s
reply from 69.89.207.99: offset -0.362657 delay 0.121926, next query 32s
adjusting local clock by 0.064548s

```

Alright our test is complete let’s run this container in the background with following

```
docker run --name ntpd \
    -v /tmp/ntpd.conf:/etc/ntpd.conf \
    --restart=always -d --privileged \
    goglides/ntpd -d -s -f /etc/ntpd.conf

```

The `-s` means “Try to set the time immediately at startup.” You can check if the process is running or not using following,

```
docker ps

Output:
CONTAINER ID        IMAGE               COMMAND                  CREATED              STATUS              PORTS               NAMES
035cf608e489        goglides/ntpd       "ntpd -d -s -f /etc/…"   About a minute ago   Up About a minute                       ntpd

```

and logs

```
docker logs ntpd

Output:
creating new /var/db/ntpd.drift
adjtimex adjusted frequency by 0.000000ppm
ntp engine ready
reply from 45.55.217.50: offset -0.395447 delay 0.139999, next query 7s
reply from 50.205.244.112: offset -0.392487 delay 0.139775, next query 6s
set local clock to Tue Mar  3 20:18:48 UTC 2020 (offset -0.395447s)
reply from 216.240.36.24: negative delay -0.247795s, next query 3217s
reply from 23.129.64.159: negative delay -0.245678s, next query 3019s
reply from 50.205.244.112: offset 0.024988 delay 0.082805, next query 5s
reply from 45.55.217.50: offset 0.019650 delay 0.093784, next query 7s
reply from 45.55.217.50: offset 0.009099 delay 0.085599, next query 9s
peer 45.55.217.50 now valid

```

## Deploy to Kubernetes cluster

Since we validate NTP is working in a container, I should able to deploy this in the Kubernetes cluster easily. For this, I have to modify `Dockerfile` little bit to address some of the challenges.

- Run container in Foreground We can easily achieve this by changing our docker entrypoint, by adding -d flag. As per man page, `-d` means do not daemonize. If this option is specified, ntpd will run in the foreground and log to stderr.
- Supply configuration file I am thinking about supplying configuration using configmap. I have to modify Dockerfile to access file config using `-f` flag. After that create configmap with actual config, mount the volume and use that custom configuration.

For this created a custom `entrypoint.sh` to handle some logic.

```
#! /bin/sh
if [ -z "${NTP_CONF_FILE}" ]
then
    NTP_CONF_FILE="/etc/ntpd.conf"
fi
ntpd -v -d -s -f ${NTP_CONF_FILE}

```

and modified `Dockerfile` to consume this.

```
FROM alpine:3.11.3
RUN apk update
RUN apk add openntpd
RUN mkdir -m 1777 /var/empty/tmp
ADD ./entrypoint.sh ./entrypoint.sh
RUN chmod 755 ./entrypoint.sh
ENTRYPOINT ["./entrypoint.sh"]

```

After that, I created a Kubernetes DaemonSet object to use this docker image. Save following content to `ds-ntpd.yaml`

```
apiVersion: v1
kind: ConfigMap
metadata:
name: ntpd-config
data:
ntpd.conf: |
    servers pool.ntp.org
---
apiVersion: apps/v1
kind: DaemonSet
metadata:
name: ntpd
labels:
    k8s-app: ntpd
    created-by: tech.goglides.com
spec:
selector:
    matchLabels:
      name: ntpd
template:
    metadata:
      labels:
        name: ntpd
    spec:
      tolerations:
      - key: node-role.kubernetes.io/master
        effect: NoSchedule
      containers:
      - name: ntp-sync
        image: goglides/ntpd
        imagePullPolicy: Never
        resources:
          limits:
            memory: 20Mi
            cpu: 20m
          requests:
            cpu: 10m
            memory: 10Mi
        securityContext:
          privileged: true
        env:
          - name: NTP_CONF_FILE
            value: /app/ntpd.conf
        volumeMounts:
          - name: ntpd-config
            mountPath: /app/
      volumes:
        - name: ntpd-config
          configMap:
            name: ntpd-config

```

Now finally apply this manifest to Kubernetes using the following,

```
kubectl apply -f ds-ntpd.yaml

Output:
configmap/ntpd-config created
daemonset.apps/ntpd created

```

You can verify if pods are running or not using following,

```
kubectl get pods

Output:
NAME         READY   STATUS    RESTARTS   AGE
ntpd-b7csn   1/1     Running   0          68s

```

And check logs,

```
kubectl logs ntpd-b7csn

Output:
creating new /var/db/ntpd.drift
adjtimex adjusted frequency by 0.000000ppm
ntp engine ready
reply from 51.158.147.92: offset -0.715358 delay 0.126846, next query 6s
reply from 176.9.40.131: offset -0.722062 delay 0.148438, next query 9s
reply from 5.186.65.2: offset -0.716915 delay 0.148267, next query 7s
reply from 212.25.15.128: offset -0.703432 delay 0.160469, next query 8s
set local clock to Tue Mar  3 23:09:38 UTC 2020 (offset -0.715358s)

```
