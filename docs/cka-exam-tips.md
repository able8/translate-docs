# How I prepared & passed the Certified Kubernetes Administrator (CKA) Exam

Tips for preparing for the Certified Kubernetes Administrator(CKA) exam.

> **NOTE:**
>
> This blog post was published in **2019** when the exam environment was based on the version **v1.16** of Kubernetes, which makes it a bit outdated by now.
>
> However, you might still find it useful as it includes general information that is not version dependent and some external resources that are still regularly updated.

* * *

Finishing the year as a Certified Kubernetes Administrator(CKA) was my personal goal for 2019, and I was able to take the exam around the end of the year and pass with a score of **91%**. 🎉

In this blog post, I wanted to share some useful resources that helped me pass the CKA exam, and a few tips that can help you prepare and hopefully pass if you are also planning to take it.

> **DISCLAIMER:** This post is a bit long because I tried to dump all the knowledge and experience I gathered when preparing for the CKA exam. So brace yourself 😀

# The Certified Kubernetes Administrator Exam

With the exploding adoption of Kubernetes, the [Certified Kubernetes Administrator](https://www.cncf.io/certification/cka/) program was created by the [Cloud Native Computing Foundation](https://www.cncf.io/)(CNCF) in collaboration with the Linux Foundation to allow Kubernetes users to demonstrate that they have the necessary skills and knowledge to perform the tasks and responsibilities of a Kubernetes administrator.

## The exam's format

The good thing about it is that it's 100% hands-on. It's an online proctored exam where you are asked to perform certain tasks on the command line.

The [Candidate Handbook](https://training.linuxfoundation.org/go/cka-ckad-candidate-handbook) is your definitive source for any details about the exam. So make sure to read it **thoroughly**.

Here is a short list of points worth mentioning:

- You need a**steady** internet connection.
- You would need a**webcam** and a **microphone** which are required by the proctor.
- You would need a government issued **ID**, or a passport.
- The exam consists of**24 questions** that you can solve in no specific order.
- The duration of the exam is**3 hours**.
- The pass mark is**74%.**
- You need to use the**Chrome browser**.
- You have**one free retake** in case you don't pass on your first try 🎉

## The curriculum

Unlike the Certified [Kubernetes Application Developer](https://www.cncf.io/certification/ckad/)(CKAD) exam, the CKA exam focuses more on cluster administration rather than deploying and managing applications on Kubernetes.

The exam's curriculum is usually updated quarterly, you can always find the latest version at:

Will I receive the hardware or come to the office to pick it up?

The CKA exam covers the following topics:

- Application Lifecycle Management – 8%
- Installation, Configuration & Validation – 12%
- Core Concepts – 19%
- Networking – 11%
- Scheduling – 5%
- Security – 12%
- Cluster Maintenance – 11%
- Logging / Monitoring – 5%
- Storage – 7%
- Troubleshooting – 10%

## The exam environment

The day of the exam, you will have multiple cluster presented for you, and with each question you will be provided with name of the cluster where you should try to solve the question.

Below is the list of the clusters provided to the candidate from the latest [Exam Tips](http://training.linuxfoundation.org/go//Important-Tips-CKA-CKAD) document available at the CKA CNCF page at the time of writing this post:

**Cluster****Members****CNI****Description**k8s1 master, 2 workersflannelk8s clusterhk8s1 master, 2 workerscalicok8s clusterbk8s1 master, 1 workerflannelk8s clusterwk8s1 master, 2 workersflannelk8s clusterek8s1 master, 2 workersflannelk8s clusterik8s1 master, 1 base nodeloopbackk8s cluster - missing worker node

The Kubernetes version running on the exam environment is currently **v1.16** at the time of writing this post, and the Linux distribution is **Ubuntu 16**.

# Preparing for the exam

The first step in preparing for the CKA exam(or any exam) is understanding what it is about.

So make sure to read all the documents provided in the CKA Program page at [https://www.cncf.io/certification/cka/](https://www.cncf.io/certification/cka/) :

- [Candidate Handbook](https://training.linuxfoundation.org/go/cka-ckad-candidate-handbook)
- [Curriculum Overview](https://github.com/cncf/curriculum)
- [Exam Tips](http://training.linuxfoundation.org/go//Important-Tips-CKA-CKAD)
- [Frequently Asked Questions](http://training.linuxfoundation.org/go/cka-ckad-faq)

## Pre-requisites

Although the CKA exam is about Kubernetes, it also requires some basic sysadmin skills. So, you need be comfortable with the Linux command line and have a minimum knowledge on how to use the following tools:

- `systemd` for managing system services. Basic knowledge would be enough IMHO, but very important especially for troubleshooting cluster components. There is a nice tutorial series for that provided by the DigitalOcean people:

Systemd Essentials: Working with Services, Units, and the Journal \| DigitalOcean

In recent years, Linux distributions have increasingly migrated away from other init systems to systemd. The systemd suite of tools provides a fast and flexible init model for managing an entire machine, from boot onwards. In this guide, we’ll give you a quick run...\


- `vim` for editing files on the command line. Although you could change the default text editor by setting the value of `$EDITOR` to nano if that's what you are most comfortable with, vim can give you a productive boost during the exam.
- `tmux` since you only get one console during the exam, being able to have multiple panes open at the same time might be helpful. Personally, I didn't really need or use tmux during the exam, so if you don't use it already in your day to day work, I don't recommend learning it for the sake of the exam.
- `openssl` for generating keys, CSRs, certificates etc.. You will probably need it during the exam for security related questions. So make sure you train yourself to use it at least for those basic use cases.

## Getting ready for the exam

In this section, I am going to provide some tips on how to prepare for the exam and also list some useful resources that helped me and might help you get fit for the exam day.

### kubectl

Since the CKA exam is 100% practical, you need to make sure you are confident enough with `kubectl`. That's mostly what you will be using during the exam, and since you are already reading this post, chances are you are already using kubectl or at least experimenting with it.

You need to be quick on the command line since you will have limited time for solving the questions during the exam, so knowing how to perform the following quickly with kubectl is crucial:

- Checking the config, switching and creating contexts
- Creating, editing and deleting kubernetes resources
- Viewing, finding and inspecting resources
- Updating and patching resources
- Interacting with pods, nodes and cluster

A lot of useful `kubectl` command examples can be found in the [kubectl cheatsheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/) available in the official kubernetes documentation.

It is also very useful to know how to use the `kubectl run` command to create resources quickly, saving time by avoiding to write yaml files(who likes that, right?). You can also use it to generate yaml files if you need to edit something before actually creating the kubernetes object by using the `--dry-run`  and the `-o yaml` options combined. Some details about the `kubectl run` usage can be found [here](https://kubernetes.io/docs/reference/kubectl/conventions/#kubectl-run).

If you come from the Docker world and still starting with Kubernetes, then the **[kubectl for Docker Users](https://kubernetes.io/docs/reference/kubectl/docker-cli-to-kubectl/)** documentation page is definitely worth checking out.

Getting familiar with **[JSONPath](https://kubernetes.io/docs/reference/kubectl/jsonpath/)** template would be also helpful. Combining kubectl and jsonpath enables you to easily extract resource information in a format that you can specify.

Finally, make sure to practice a lot with kubectl, whether it is on local kubernetes clusters with minikube, docker-desktop or on the cloud. That is very crucial for the exam.

### Learning resources

If you are planning to take the CKA exam, them you probably have already searched around the internet for some resources and found plenty. So in this section, I am only going to list the resources that I have found most informative and helpful for me to pass the exam.

**_The Kubernetes Documentation_**

The most important resource is the [official kubernetes documentation](https://kubernetes.io/docs/home/); that's your definitive source of information. And since you are **allowed** to access it during the exam, it's really important that you know how to easily navigate it and quickly search for what you need. Make sure to get well accustomed to it.

Also, make sure you get to do most if not all of the **tasks** listed [here](https://kubernetes.io/docs/tasks/).

> It's really useful to join the kubernetes slack community at [https://slack.k8s.io/](https://slack.k8s.io/). There is a slack channel dedicated to CKA exam questions named **#cka-exam-prep.**
>
> The members there are really nice and helpful and would answer any questions you have.

**Kubernetes The Hard Way(KTHW)**

The [kubernetes-the-hard-way](https://github.com/kelseyhightower/kubernetes-the-hard-way) repo was created by Kelsey Hightower to provide a guide for bootstrapping a Kubernetes cluster on Google Cloud Platform. It helps you understand the internals of a Kubernetes cluster, which would be really important especially for troubleshooting.

Make sure to go through it at least once while trying to understand **every** step on the way.

If you don't want to use GCP, there is another fork that relies on vagrant and can be found here:

**Online Courses**

There are a couple of online course available that can help you prepare for the CKA exam.

I was able to try 3 of them while preparing for the exam:

- [Kubernetes Fundamentals (LFS258)](https://training.linuxfoundation.org/training/kubernetes-fundamentals/) which was part of the CKA exam bundle I purchased.

I only liked the practice labs, otherwise the course was very boring and you'd better read the kubernetes documentation rather than reading slides. So _IMHO_ totally not worth taking.
- [Linuxacademy: Cloud Native Certified Kubernetes Administrator (CKA)](https://linuxacademy.com/course/cloud-native-certified-kubernetes-administrator-cka/):

I found this course really good at first. However, after a while I found myself watching the instructor mostly typing commands in the terminal so I got disconnected and stopped following the course. I also tried the mock exams, but I found them a bit limited.
- Udemy:[Certified Kubernetes Administrator (CKA) with Practice Tests](https://www.udemy.com/course/certified-kubernetes-administrator-with-practice-tests/)

This was the most comprehensive course for me in this list. It covered all the topics, and the instructor made sure to explain all the Kubernetes concepts(and also other concepts) thouroughly.

The practice labs are really good since you are provided with an environment and your answers are checked there automatically.

The mock exams were also a great preparation for the exam.

**I cannot _recommend_ this course enough!**

**Additional Resources**

The [Kubernetes in Action](https://www.manning.com/books/kubernetes-in-action) book by Marko Lukša is definitely worth reading to have a good understanding of Kubernetes.

There is a google spreadsheet created by the community that compiles a lot of useful resources that can be found [here](https://bit.ly/2IdKwIc).

Some additional useful Github repositories:

\- [https://github.com/stretchcloud/cka-lab-practice](https://github.com/stretchcloud/cka-lab-practice)

\- [https://github.com/walidshaari/Kubernetes-Certified-Administrator](https://github.com/walidshaari/Kubernetes-Certified-Administrator)

\- [https://github.com/krzko/awesome-cka](https://github.com/krzko/awesome-cka)

\- [https://github.com/David-VTUK/CKA-StudyGuide](https://github.com/David-VTUK/CKA-StudyGuide)

## Tips for the exam day

In this section, I am going to provide a few tips for the day of the exam:

- You are allowed to open one additional browser tab in addition to the exam interface and you can use it to browse the kubernetes documentation._Bookmarks_ are also allowed, so make sure to create some bookmarks in chromes for the documentation sections that you think you would need in the exam beforehand.
- You don't have to solve the questions in a specific order. So you can start with the easiest to build some confidence, but that's a matter of personal preference.
- There is built-in notepad in the exam interface which might be handy since you're not allowed to write on paper during the exam. You can use it to write the questions' numbers so that you keep track of the ones you didn't solve and get back to them later.
- If you are taking the exam with a laptop, use an external monitor if your laptop screen is tiny. You would need all the space you can get for the terminal.
- Make sure to go to the restroom before starting the exam. During the 3 hours, you would only be able to take a break if your proctor allows it but the timer would never stop for you.
- Have some water in a bottle without a label, or a transparent glass. Anything other than that is not allowed.
- Take the exam in a quiet room, on a clean desk. Remove any electronics from the desk and make sure that absolutely no one enters the room during the exam.

You will be asked by the proctor to show him around the room using the webcam.
- Finally:**_GOOD LUCK!_**

# Conclusion

In this post, I tried to provide some tips and resources for preparing the CKA exam based on my experience.

I hope this article would be useful for you and please let me know in the comments if it somehow helped you to pass the exam.

### Subscribe to Mehdi Yedes' blog

Get the latest posts delivered right to your inbox

Subscribe
