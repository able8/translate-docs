# What are DORA (DevOps Research and Assessments) Metrics?

https://www.splunk.com/en_us/data-insider/devops-research-and-assessment-metrics.html

DORA metrics are a framework of performance metrics that help DevOps teams understand how effectively they develop, deliver and maintain software. They identify elite, high, medium and low performing teams and provide a baseline to help organizations continuously improve their DevOps performance and achieve better business outcomes. DORA metrics were defined by Google Cloud’s DevOps Research and Assessments team based on six years of research into the DevOps practices of 31,000 engineering professionals.

While DevOps and engineering leaders can often provide a gut-level assessment of their team’s performance, they struggle to quantify this value to the business — or to pinpoint where and how improvements can be made. DORA metrics can help by providing an objective way to measure and optimize software delivery performance and validate business value.

In the following sections, we’ll look at the four specific DORA metrics, how software engineers can apply them to assess their performance and the benefits and challenges of implementing them. We’ll also look at how you can get started with DORA metrics.

## Components of DORA Metrics

What are the four key metrics in DORA?

The DORA framework uses the four key metrics outlined below to measure two core areas of DevOps: speed and stability. Deployment Frequency and Mean Lead Time for Changes measure DevOps speed, and Change Failure Rate and Time to Restore Service measure DevOps stability. Used together, these four DORA metrics provide a baseline of a DevOps team’s performance and clues about where it can be improved.

1) Deployment Frequency

Deployment frequency indicates how often an organization successfully deploys code to production or releases software to end users. DevOps’ goal of continuous development essentially requires that teams achieve multiple daily deployments; the deployment frequency metric provides them a clear picture of where they stand in relation to that goal.

The deployment frequency benchmarks are:

- **Elite**: Multiple deployments per day
- **High**: One deployment per week to one per month
- **Medium**: One deployment per month to one every six months
- **Low**: Fewer than one deployment every six months

Organizations vary in how they define a successful deployment, and deployment frequency can even differ across teams within a single organization.

2) Mean Lead Time for Changes

Mean lead time for changes measures the average time between committing code and releasing that code into production. Measuring lead time is important because the shorter the lead time, the more quickly the team can receive feedback and release software improvements. Lead time is calculated by measuring how long it takes to complete each project from start to finish and averaging those times.

Mean lead time for changes benchmarks are:

- **Elite**: Less than one per hour
- **High**: Between one day and one week
- **Medium**: Between one month and six months
- **Low**: More than six months

An organization’s particular cultural processes — such as separate test teams or shared test environments — can impact lead time and slow a team’s performance.

3) Change Failure Rate

Change failure rate is the percentage of deployments causing a failure in production that require an immediate fix, such as service degradation or an outage. A low change failure rate is desirable because the more time a team spends addressing failures, the less time it has to deliver new features and customer value. This metric is usually calculated by counting how many times a deployment results in failure and dividing that by the total number of deployments to get an average. You can calculate this metric as follows:

_(deployment failures / total deployments) x 100_

Change failure rate benchmarks are:

- **Elite**: 0-15%
- **High**: 16-30%
- **Medium**: 16-30%
- **Low**: 16-30%

4) Time to Restore Service

The time to restore service metric, sometimes called mean time to recover or mean time to repair (MTTR), measures how quickly a team can restore service when a failure impacts customers. A failure can be anything from a bug in production to an unplanned outage.

Time to Restore Service benchmarks are:

- **Elite**: Less than one hour
- **High**: Less than one day
- **Medium**: Between one day and one week
- **Low**: More than six months

![four-key-metrics-of-dora](http://www.splunk.com/content/dam/splunk2/images/data-insider/dora-metrics/four-key-metrics-of-dora.jpg)

DORA uses four main metrics to measure two core areas of DevOps: speed and stability.

What is a DORA (DevOps Research and Assessments) survey?

A DORA survey is a simple way to collect information around the four DORA metrics and measure the current state of an organization’s software delivery performance. Google Cloud’s DevOps Research and Assessments team offers an official survey called the [DORA DevOps Quick Check](https://www.devops-research.com/quickcheck.html). You simply answer five multiple-choice questions and your results are compared to other organizations, providing a top-level view of which DevOps capabilities your organization should focus on to improve.

While a DORA survey can provide generalized guidance, many organizations additionally enlist the help of third-party vendors to conduct personalized assessments. These more closely examine a company’s culture, practices, technology and processes to identify specific ways to improve its DevOps team’s productivity.

## Benefits of DORA Metrics

What are some applications/use cases of DORA metrics?

Companies in virtually any industry can use DORA metrics to measure and improve their software development and delivery performance. A mobile game developer, for example, could use DORA metrics to understand and optimize their response when a game goes offline, minimizing customer dissatisfaction and preserving revenue. A finance company might communicate the positive business impact of DevOps to business stakeholders by translating DORA metrics into dollars saved through increased productivity or decreased downtime.

What are the benefits and challenges of DORA metrics?

DORA metrics are a useful tool for quantifying your organization’s software delivery performance and how it compares to that of other companies in your industry. This can lead to:

- **Better decision making**: Consistently tracking DORA metrics helps organizations use hard data to understand the current state of their development process and make decisions about how to improve it. Because DORA helps DevOps teams focus on specific metrics rather than monitoring everything, teams can more easily identify where bottlenecks exist in their process, focus efforts on resolving them and validate the results. The net result is faster, high-quality delivery driven by data rather than gut instinct.
- **Greater value**: Elite and high performing teams can be confident they are delivering value to their customers and positively impacting the business. DORA metrics also give lower-performing teams a clear view of where their process may be stalled and help them identify a path for delivering greater value.
- **Continuous improvement**: DevOps teams can baseline their performance with DORA metrics and discover what habits, policies, processes, technologies and other factors are impeding their productivity. The four key metrics provide a path to setting goals to optimize the team’s performance and determine the most effective ways to do so.

![benefits-of-dora-metrics](http://www.splunk.com/content/dam/splunk2/images/data-insider/dora-metrics/benefits-of-dora-metrics.jpg)

DORA metrics can lead to better decision making, greater value and continuous improvement.

Even though DORA metrics provide a starting point for evaluating your software delivery performance, they can also present some challenges. Metrics can vary widely between organizations, which can cause difficulties when accurately assessing the performance of the organization as a whole and comparing your organization’s performance against another’s. Each metric typically also relies on collecting information from multiple tools and applications. Determining your Time to Restore Service, for example, may require collecting data from PagerDuty, GitHub and Jira. Variations in tools used from team to team can further complicate collecting and consolidating this data.

## Measuring With DORA Metrics

How do you measure DevOps success with DORA?

DORA uses the four key metrics to identify elite, high, medium, and low performing teams. The [State of DevOps Report](https://www.devops-research.com/research.html#reports) has shown that elite performers have 208 times more frequent code deployments, 106 times faster lead time from commit to deploy, 2,604 times faster time to recover from incidents and 7 times lower change failure rate than low performers. Elite performing teams are also twice as likely to meet or exceed their organizational performance goals.

How do you measure and improve [MTTR](https://www.splunk.com/en_us/data-insider/what-is-mean-time-to-repair.html#:~:text=MTTR%20is%20calculated%20by%20dividing,MTTR%20would%20be%20two%20hours.) with DORA?

MTTR is calculated by dividing the total downtime in a defined period by the total number of failures. For example, if a system fails three times in a day and each failure results in one hour of downtime, the MTTR would be 20 minutes.

_MTTR = 60 min / 3 failures = 20 minutes_

MTTR begins the moment a failure is detected and ends when service is restored for end users — encompassing diagnostic time, repair time, testing and all other activities.

In DORA, MTTR is one measure of the stability of an organization’s continuous development process and is commonly used to evaluate how quickly teams can address failures in the continuous delivery pipeline. A low MTTR indicates that a team can quickly diagnose and correct problems and that any failures will have a reduced business impact. A high MTTR indicates that a team’s incident response is slow or ineffective and any failure could result in a significant service interruption.

While there’s no magic bullet for improving MTTR, response time can be reduced by following some best practices:

- **Understand your incidents**: Understanding the nature of your incidents and failures is the first step in reducing MTTR. Modern enterprise software can help provide a consolidated view of your siloed data, producing a reliable MTTR metric that reveals insights into its contributing factors.
- **Make sure to monitor**: The earlier you can identify a problem, the better your chances of resolving it before it impacts users. A modern monitoring solution provides a continuous stream of real-time data about your system’s performance in a single, easy-to-digest dashboard interface and will alert you to any issues.
- **Create an [incident-management](https://www.splunk.com/en_us/data-insider/what-is-incident-management.html) action plan**: Generally, companies favor one of two approaches: ad hoc responses are often necessary for smaller, resource-strapped companies, while large enterprises favor more rigid procedures and protocols. Whatever shape your plan takes, make sure it clearly outlines whom to notify when an incident occurs, how to document it and what steps to take to address it.
- **Automate your incident-management system**: While a simple phone call may work when a low-priority incident occurs during business hours, you need to make sure all your bases are covered when a major incident strikes, particularly during off hours. An automated incident-management system that can send alerts by phone, text and email to all designated responders at once is critical for mounting a quick response.
- **Cross-train team members for different response roles**: Having only one person knowledgeable about each system or technology is risky. What if that system goes down when they’re on vacation? With multiple engineers each versed in several relevant functions and responsibilities, your team is better positioned to respond effectively — no matter who’s on-call.
- **Take advantage of AI**: [AIOps](https://www.splunk.com/en_us/data-insider/ai-for-it-operations-aiops.html) help [DevOps](https://www.splunk.com/en_us/data-insider/what-is-devops-and-why-is-it-important.html) teams better respond to production failures and lower their MTTR. Specifically, AIOps can help detect issues before they impact users — prioritizing incidents by criticality, correlating and contextualizing related incidents, alerting and escalating incidents to the appropriate response team members, and automating remediation to resolve incidents.
- **Have a follow-up procedure**: Once an incident is resolved, it’s important to follow up with all the key team members to determine how and why the problem occurred and strategize how to prevent it from happening again. In DevOps, this typically takes the form of a blameless post-incident review, which analyzes both the technical and human factors of their response efforts that can be improved. Ultimately, this results in improved incident response and lower MTTR, and also more innovative ideas and better applications.

What is DORA in Agile?

In Agile, DORA metrics are used to improve the productivity of DevOps teams and the speed and stability of the software delivery process. DORA supports Agile’s goal of delivering customer value faster with fewer impediments by helping identify bottlenecks. DORA metrics also provide a mechanism to measure delivery performance so teams can continuously evaluate practices and results and quickly respond to changes. In this way, DORA metrics drive data-backed decisions to foster continuous improvement.

What are flow metrics?

Flow metrics are a framework for measuring how much value is being delivered by a product value stream and the rate at which it is delivered from start to finish. While traditional performance metrics focus on specific processes and tasks, flow metrics measure the end-to-end flow of business and its results. This helps organizations see where obstructions exist in the value stream that are preventing desired outcomes.

There are four primary flow metrics for measuring value streams:

- **Flow velocity** measures the number of flow items completed over a period to determine if value is accelerating.
- **Flow time** measures how much time has elapsed between the start and finish of a flow item to gauge time to market.
- **Flow efficiency** measures the ratio of active time to total flow time to identify waste in the value stream.
- **Flow load** measures the number of flow items in a value stream to identify over- and under-utilization of value streams.

Flow metrics help organizations see what flows across their entire software delivery process from both a customer and business perspective, regardless of what software delivery methodologies it uses. This provides a clearer view of how their software delivery impacts business results.

## Getting Started With DORA Metrics

What are some popular DORA metrics tools?

Once you automate DORA metrics tracking, you can begin improving your software delivery performance. Severalv engineering metrics trackers capture the four key DORA metrics, including:

- Faros
- Haystack
- LinearB
- Sleuth
- Velocity by Code Climate

When considering a metric tracker, it’s important to make sure it integrates with key software delivery systems including CI/CD, issue tracking and monitoring tools. It should also display metrics clearly in easily digestible formats so teams can quickly extract insights, identify trends and draw conclusions from the data.

How do you get started with DORA metrics?

To get started with DORA metrics, start collecting data. There are many data collection and visualization solutions on the market, including those mentioned above. The easiest place to start, however, is with Google’s [Four Keys](https://cloud.google.com/blog/products/devops-sre/using-the-four-keys-to-measure-your-devops-performance) open source project, which it created to help DevOps teams generate DORA metrics. Four Keys is an ETL pipeline that ingests data from Github or a Gitlab repository through Google Cloud services and into Google DataStudio. The data is then aggregated and compiled into a dashboard with data visualizations of the four key DORA metrics, which DevOps teams can use to track their progress over time.

## Splunk IT/Observability Predictions 2022

[Get the Report](https://www.splunk.com/en_us/form/it-predictions.html)

## The Bottom line: DORA metrics are the key to getting better business value from your software delivery

Data-backed decisions are essential for driving better software delivery performance. DORA metrics give you an accurate assessment of your DevOps team’s productivity and the effectiveness of your software delivery practices and processes. Every DevOps team should strive to align software development with their organization’s business goals. Implementing DORA metrics is the first step.
