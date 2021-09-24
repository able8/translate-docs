# How to Prepare for CKAD and CKA Certification? 				

InfraCloud Team

Jun 10th, 2021

While preparing for the CNCF’s CKAD or CKA certification, there could be numerous doubts, which exam to appear first, what resources to refer to, what are common mistakes to avoid, etc.. Especially, if you  don’t have previous knowledge or hands-on experience with Kubernetes,  this could be a tough situation to be in. At InfraCloud, engineers are  highly encouraged to appear for these exams and get certified. This blog post is a collaborative effort from the recently certified Infranauts  to share all the insights straight from - before registering for the  exam to the next steps after clearing the exam successfully.

With around 50% developers CKA and CKAD certified, we share our  experiences, study material, Do’s and Dont’s, FAQ, etc. about the exam.  If you’re willing to start your journey in Kubernetes and are aiming to  be certified, you will find this blog helpful.

Shall we begin then?

## CKA vs CKAD certification! Which one should I take first?

First things first! Before thinking of CKA or CKAD, having knowledge  of Kubernetes basics is crucial. If you’re a complete newbie and do not  know what on earth is Kubernetes? It is highly recommended to get  familiar with the basic concepts of Kubernetes before you book your exam dates.

If you’re starting out your journey into the cloud native and Kubernetes world, you can opt for [Kubernetes For The Absolute Beginner - Hands On by Mumshad Mannambeth](https://kodekloud.com/courses/kubernetes-for-the-absolute-beginners-hands-on/) course to get yourself comfortable with the basics of Kubernetes.

Although one might still get confused about which exam one should go  for first - CKA or CKAD? Is CKA harder than CKAD? And what is the exact  difference in terms of exam curriculum and the difficulty level of both  exams? To keep things simple, you can look at it this way:

- **CKAD** is for those interested in the design, build & configuration of cloud native applications using Kubernetes.
- While **CKA** exam is mainly for those, who want to build, manage the Kubernetes infrastructure.

Some of us started digging deeper over the internet and eventually  find out that CKAD is broadly a subset of CKA. A few Infranauts decided  to prepare for CKAD first and then for CKA, as they were ultimately  aiming to gain all the possible knowledge around Kubernetes. It  eventually turned out to be the right decision as CKAD also helped them  improve speed and muscle memory over Kubernetes commands, which is  critical for both exams. This helped them a lot in the CKA exam as most  of us managed to complete almost all questions with 20-30 mins still  left with us to solve the only couple of complex questions that had been flagged and skipped for later.

## Study material

Many Infranauts have taken the [CKAD course](https://www.udemy.com/course/certified-kubernetes-application-developer/) by Mumshad Mannambeth available on Udemy. After clearing CKAD with that course, some of us also opted for his another [CKA course](https://www.udemy.com/course/certified-kubernetes-administrator-with-practice-tests/) for the CKA certification. The primary reason to go through these courses was the number of  recommendations it had from the community. Few of the certified  Infranauts had also watched a few of his videos on YouTube and always  found it easy to follow. So yes, these courses are highly recommended to anyone planning to get CKA / CKAD certified.

As you might guess, this is not the only course in the market - there are plenty of them. Study materials are available in other formats like blogs, GitHub Repos, YouTube videos containing different scenario-based practice questions/answers. You may choose at your own will like [this one](https://medium.com/bb-tutorials-and-thoughts/practice-enough-with-these-questions-for-the-ckad-exam-2f42d1228552) that some of us followed from [Medium](https://medium.com/) for the CKAD exam. You should also check out the video by [Harshit Singhvi on exam experience and tips](https://www.youtube.com/watch?v=FZ3VQC-aRmI) on YouTube.

At the end of the day, what matters the most is how much time you spend **practicing** different scenarios. We can’t emphasize enough how important the labs are. Irrespective of the course you take, they all have labs and mock tests. Don’t skip any of them. They are all equally important.

## Mistakes to avoid

- Don’t get trapped into registering the exam date six months in  advance and then start the actual study. If you’re not fully prepared by the time nearing your exam date, you might get nervous and feel to  postpone the exam to study more and more. This cycle never ends.
- Even if you started late, don’t be inconsistent with your study. One would study one day and skip for another two days. Don’t do that. Be consistent and stick to a schedule.
- Do NOT skip labs. One might get trapped thinking it would be nice to go through all the lectures at once and then come back to labs. Nope, not a good idea!
- Completing a course is one thing. But you won’t know how much you understand until you take the mock tests. So don’t avoid taking them until you’re left with last 3-4 days to study.
- Don’t keep postponing till the last day of validity of your particular exam voucher. If you delay it till the last day of voucher, and if you couldn’t clear the exam on the first attempt, you’ll leave no *retake* option. So, prepare, practice, and appear for the exam at least a week before the last date.

## Do’s for CKA and CKAD Certification

- The most obvious one - ***practice\***, ***practice\*** and some more ***practice\***.
- Make it a habit to take a lecture and do the lab that follows with it.
- If possible, complete the courses first, and when you think you  are ready for the exam, go ahead and apply for the certification. This gives you a year ***after\*** you have completed the course. It is enough time to test your skills.
- The discussions in the Slack forum of KodeKloud CKA and CKAD group also helped us play and try different variety of questions/scenarios  during preparation. So, keep an eye on it. Plus, feel free to seek  suggestions from the forum in case of any doubts.
- Practice with complete commands instead of binding an `alias` for like everything.
- In each mock test, try to complete the test 15 min before the deadline. This gives you time to revisit the questions.
- Monitor your time as you practice.
- Do the labs and mock tests repeatedly (at least three times). Identify what is slowing you down and plan accordingly.
- On the day of the exam, make sure to clean your desk and not have anything (apart from a transparent water bottle).
- On the exam day, keep an alternative internet source handy in case of Wi-Fi internet goes down (trust me, it happens more than you  believe)
- In the exam, if you analyze that any particular question is going  to take more than 6-7 mins to solve, flag/mark it to solve for later and come back once you solve the rest.

## Don’ts for CKA and CKAD Certification

- Don’t overwhelm yourself with an `alias` for everything.
- Don’t panic if you are stuck; simply flag/mark the particular question and move ahead. You can always come back to it later.
- Don’t give the exam on the last day (like many folks do).
- At the time of the exam, other than your system, don’t have anything on the table.
- Don’t overkill yourself with Kubernetes the Hard Way unless you have time.
- Don’t ignore the candidate handbook. Read it and follow the instructions.

## Final tips & tricks

- The exam clusters are set up with `kubeadm` mostly in the Ubuntu environment.

- Do check out the [CKA and CKAD environment](https://docs.linuxfoundation.org/tc-docs/certification/tips-cka-and-ckad) details and practice accordingly.

- Try to use auto-completion while running Kubernetes commands.  It will be much more helpful and effective in the exam.

- Here is how you can set up the auto-completion:    

  ```
  $ alias k='kubectl'
  
  $ source <(kubectl completion bash)
  $ echo "source <(kubectl completion bash)" >> ~/.bashrc
  
  $ complete -F __start_kubectl k
  ```

- The above commands can be found in the [kubectl cheat sheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/). And, **yes** you can use them (some of us have used them ourselves), as they are listed in the docs.

- As you do your labs, you must search the docs. Especially for resources like PV, PVC, Ingress, RBAC resources etc.

- This trains your brain to remember the correct links, in case you need to refer to the docs during the exam. Saves a lot of time.

- If you have time during preparation, take both the courses CKA and CKAD. This will surely help.

- Get familiar with `vi` or `vim`. We’re not sure if `nano` will be available, so prefer `vi` or `vim`. GNU Emacs is available but the usual key bindings like `C-p`, `C-n` don’t work inside most of the browsers.

- You don’t need `tmux` or `screen`, if you have saved your aliases or other bash settings in `~/.bashrc`.

- Get used to pasting text using the mouse middle/center key. The mouse secondary key doesn’t work in the exam environment.

## Useful commands

Here is a *nonextensive* list of commands that you will definitely need during practice or exam:

```
# list running processes
$ ps -aux

# search for string in the output
$ ps -aux | grep -i 'string' 

# search for multiple expressions in the output (exp can be plain a string too)
$ ps -aux | grep -e 'exp-one' -e 'exp-two'

# get details about network interfaces
$ ifconfig

# list network interfaces and their IP address
$ ip a

# get the route details
$ ip r

# check service status and also show logs
$ systemctl status kubelet

# restart a service
$ systemctl restart kubelet

# reload the service daemon, if you changed the service file
$ systemctl daemon reload

# detailed logs of the service
$ journalctl -u kubelet

# list out ports, protocol and what processes are listening on those ports
$ netstat -tunlp
```

## How to book the CKA and CKAD exams?

You can book the exams from [Linux Foundation](https://training.linuxfoundation.org/full-catalog/?_sft_product_type=certification) page.

## Exam experience

### CKAD exam experience

On the exam day, try to log in 15 minutes before the exam. Proctor  will make sure to follow their process to check your ID proof as well as room, your desk. The entire process should normally take 15-20 minutes  or more, but don’t panic as the proctor will only start the exam after  all the verification process is complete and you’re comfortable to start the exam. CKAD exam questions are comparatively quite straightforward,  but it would be a race against time, as some of the questions will  contain many Kubernetes objects creation and if you fail to create any,  would take more time to debug and understand them.

### CKA exam experience

The exam experience was a little rollercoaster ride for some of us,  as we lost internet connection almost 4-5 times during the exam. Each  time after the connection reset, we made sure that previous work is not  lost by checking answers of some of the previously solved questions.  Proctor typically makes sure to hold your exam (to avoid time loss) and  stop the timer and resume it once your connection is restored. But this  whole experience can make you nervous during and after the exam (with a  nightmare of all work might be lost, resulting in failure). Few of us  had practiced enough, so we were able to solve 14-15 out of 17 questions in 1.5 hours and spend the rest of almost 30 minutes to try to solve  the 2-3 questions which we initially found challenging and had marked  (and parked) to solve later.

## FAQs for the CKA and CKAD Certification

Some of the questions which crossed our minds while preparing (apart from the ones available on the [official FAQ page of the Linux Foundation](https://docs.linuxfoundation.org/tc-docs/certification/faq-cka-ckad-cks)).

***1) Can I use dual monitors in the exam?\***
 Yes, we are allowed to use dual monitors

***2) Can I bookmarks link and use them during the exam?\***
 Yes, we are allowed to use them as long as they are referring to official Kubernetes documentation allowed for the exam

***3) When would the result be out after the exam?\***
 Results are typically be out after the exact 24 hours of your  examination but can get delayed till 36 hours in some cases. In case of  delay, you can raise a [support ticket](https://jira.linuxfoundation.org/plugins/servlet/theme/portal/15) for the same.

***4) Does our work remains saved in case of internet connectivity loss?\***
 Yes, your work gets autosaved at frequent intervals.

***5) Who evaluates the CKA / CKAD / CKS exams? Humans or bots?\***
 Automated scripts most probably evaluate them.

***6) What should I be more prepared for if I get failed in my 1st attempt?\***
 Focus on solving all the questions and scenarios that you had found  difficult in 1st attempt and the topic you found the first time during  your exam.

***7) Whom to reach out in case of any challenges faced during the exam?\***
 You can raise a [support ticket](https://jira.linuxfoundation.org/plugins/servlet/theme/portal/15) for the same.

***8) What is the best time to book the CKA and CKAD exams?\***
 Once you’re able to solve all the labs and mock tests in less than 40-45 minutes, consider you’re ready for the exam. You can also book them  using the discounts available (up to 50%) during the KubeCon event and  Black Friday or Cyber Monday deal sale from Linux Foundation.

***9) What is the duration of the exam?\***
 Both CKA and CKAD exams are of 2 hours once you start the exam  (excluding time required for identity verification with proctor before  you start the exam).

***10) How many questions does CKAD or CKA exam contain?\***
 Both CKA and CKAD contains 17-19 questions (approximately).

***11) What is the minimum or maximum weightage of any question in the exam?\***
 Questions’ weightage varies from 2% to 13% based on the work required to solve them.

***12) Is a digital watch available on the exam console?\***
 Yes, a progress bar available on the top-left corner of your exam console.

***13) Do we need to be more fluent in jsonpath syntax without any external website reference during the exam?\***
 You can practice/focus on the JSON expressions captured as part of the [cheat sheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/), but don’t spend too much time on mastering JSON expressions from an exam perspective.

***14) Are exam questions as difficult as one available on the [killer.sh](https://killer.sh/) simulator?\***
 Personally, many of us felt the actual exam was comparatively easier than Killer.sh simulator.

Practice and prepare well :)

## Reference links to useful materials:

### Bookmark

Import the Kubernetes official documentation bookmark into your  chrome/any browser as it will help you search required topics faster  during the exam.

- https://github.com/reetasingh/CKAD-Bookmarks

### Recommended articles before the Exam:

Go through the below articles at least once before the exam, as they  contain some of the tricky topics from CKA exam perspective.

- [Resize a PVC](https://www.runlevl4.com/kubernetes/patching-kubernetes-resources)
- [Kubernetes Volume Basics: emptyDir and PersistentVolume](https://www.alibabacloud.com/blog/kubernetes-volume-basics-emptydir-and-persistentvolume_594834)
- [Sidecar containers](https://medium.com/bb-tutorials-and-thoughts/kubernetes-learn-sidecar-container-pattern-6d8c21f873d)
- [Ingress](https://matthewpalmer.net/kubernetes-app-developer/articles/kubernetes-ingress-guide-nginx-example.html)
- [Network policy](https://reuvenharrison.medium.com/an-introduction-to-kubernetes-network-policies-for-security-people-ba92dd4c809d)

## Practice material

- [CKAD exercises Repo](https://github.com/dgkanatsios/CKAD-exercises)
- [K8s Practice Training Repo](https://github.com/StenlyTU/K8s-training-official)
- [KodeKloud](https://kodekloud.com/)
- [Killer.sh](https://killer.sh/) (Now [available with exam registration](https://training.linuxfoundation.org/announcements/linux-foundation-kubernetes-certifications-now-include-exam-simulator/))

## Conclusion

Though one finds the Kubernetes exams comparatively hard as these are practical exams, if you practice enough and take note of the above  points, you can clear them easily; irrespective of whether you have had  previous experience with Kubernetes or not.

We hope this helps you plan and prepare better for the CKA and CKAD certification exam. We will be happy to answer any other questions.

## Connect with us

We hope the article was useful. Feel free to drop us a line or share your exam experiences with any one of the authors:

[Ninad Desai](https://www.linkedin.com/in/ninad-desai/), Gaurav Gahlot, [Yatish Sharma](https://www.linkedin.com/in/baba230896/), [Jaiganesh Karthikeyan](https://www.linkedin.com/in/jaiganesh-karthikeyan/), and [Dhruv Mewada](https://www.linkedin.com/in/dhruvmewada/).

All the very best folks :)

Love Cloud Native? We do too ❤️ 

Sign up for your weekly dose, without any spam. 