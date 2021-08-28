# An Introduction to Producing and Consuming Kafka Messages in Go

# Go 中 Kafka 消息的生产和消费介绍

June 13, 2020 From: https://www.aaronraff.dev/blog/an-introduction-to-producing-and-consuming-kafka-messages-in-go

Kafka is a popular distributed streaming platform. Let’s take a look at how to produce and consume messages in Go!

Kafka 是一个流行的分布式流媒体平台。让我们来看看在 Go 中如何生产和消费消息！

![An illustration of a Kafka producers and consumers.](https://www.aaronraff.dev/static/5a4c44d3efd2c3cf9f9d763e2d72ef37/bc69a/an-introduction-to-producing-and-consuming-kafka-messages-in-go-featured.jpg)


# What is Kafka?

# 什么是卡夫卡？

The [official documentation](https://kafka.apache.org/) describes Kafka as being a “distributed streaming platform”. In many  cases it is used as a message queue that microservices produce events  to. These events are then consumed and processed by other microservices. This is the use case that this post will focus on, however there are  many [other ways](https://kafka.apache.org/uses) which Kafka can be used.

[官方文档](https://kafka.apache.org/) 将 Kafka 描述为“分布式流平台”。在许多情况下，它用作微服务向其生成事件的消息队列。然后这些事件被其他微服务使用和处理。这是本文将重点关注的用例，但是有许多 [其他方式](https://kafka.apache.org/uses) 可以使用 Kafka。

There are a few important concepts to understand before starting. A *broker* in Kafka is another term for a server in the cluster. These brokers manage *topics* which is a way to group messages together. *Producers* are processes that write messages to topics, and *consumers* are processes that read messages from topics. The details about how  each of these components are designed are out of scope for this post,  but the [documentation](https://kafka.apache.org/intro) outlines each of these if you are interested.

在开始之前，有几个重要的概念需要理解。 Kafka 中的 *broker* 是集群中服务器的另一个术语。这些代理管理 *topics*，这是一种将消息分组在一起的方法。 *生产者*是将消息写入主题的进程，而*消费者*是从主题读取消息的进程。有关如何设计这些组件的详细信息超出了本文的范围，但如果您有兴趣，[文档](https://kafka.apache.org/intro) 概述了其中的每一个。

# Running Kafka locally

# 在本地运行 Kafka

All of the examples in this post will be interacting with a Kafka  cluster that is running locally on my machine. Getting one set up is  fairly straightforward and is explained in detail in the [quick start guide](https://kafka.apache.org/quickstart). It’s worth noting that my “cluster” will consist of only one node so  that we can focus on producing and consuming messages rather than  configuring and managing a cluster of servers. Of course, in practice you would want a cluster of multiple machines so that you can take  advantage of Kafka’s fault-tolerance.

本文中的所有示例都将与在我的机器上本地运行的 Kafka 集群进行交互。进行设置相当简单，[快速入门指南](https://kafka.apache.org/quickstart) 中有详细说明。值得注意的是，我的“集群”将只包含一个节点，这样我们就可以专注于生产和消费消息，而不是配置和管理服务器集群。当然，在实践中你会想要一个多台机器的集群，这样你就可以利用 Kafka 的容错。

# Producing messages

# 生产消息

The package that we will be using is [confluent-kafka-go](https://github.com/confluentinc/confluent-kafka-go). The examples in this post are adapted from the examples in their repository.

我们将使用的包是 [confluent-kafka-go](https://github.com/confluentinc/confluent-kafka-go)。这篇文章中的示例改编自其存储库中的示例。

I will be producing to a topic named “test”, but of course you can  switch that out for whatever you’d like. If you have not already created a topic, you can do so by running:

我将制作一个名为“测试”的主题，但当然你可以将其切换为任何你想要的。如果您尚未创建主题，则可以通过运行：

*command line*

*命令行*

```text
bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic test
```


This script is inside the Kafka directory. Notice that this topic  will not be partitioned or replicated at all, since we only have one  machine in our cluster.

该脚本位于 Kafka 目录中。请注意，该主题根本不会被分区或复制，因为我们的集群中只有一台机器。

The first thing that we need to do is create a producer. We will need to provide it with our *bootstrap.servers* which is a comma separated list of the brokers in our cluster.

我们需要做的第一件事是创建一个生产者。我们需要为它提供我们的 *bootstrap.servers*，它是我们集群中代理的逗号分隔列表。

*producer.go*

```go
func main() {
    config := &kafka.ConfigMap{"bootstrap.servers": "localhost:9092"}
    producer, err := kafka.NewProducer(config)
    if err != nil {
        log.Fatalf("Error creating producer: %s\n", err)
    }
}
```


Next, we will create a message to be sent to the broker. We will need to include what topic and partition we want to send the message to as  well. Since we don’t care about the specific partition we can just use `kafka.PartitionAny`.

接下来，我们将创建要发送给代理的消息。我们还需要包括我们要将消息发送到的主题和分区。由于我们不关心具体的分区，我们可以只使用 `kafka.PartitionAny`。

*producer.go*

```go
func main() {
    ...
    topic := "test"
    record := &kafka.Message{
        TopicPartition: kafka.TopicPartition{
            Topic:     &topic,
            Partition: kafka.PartitionAny,
        },
        Value: []byte("hello world!"),
    }
}
```


Now let’s send the message to the broker. To do this, we can send the message through the produce channel. It’s important to note that this  is an asynchronous operation, so we may want to also wait for the report to come back. We can do this by taking a look at the events channel.

现在让我们将消息发送给代理。为此，我们可以通过生产通道发送消息。需要注意的是，这是一个异步操作，因此我们可能还想等待报告回来。我们可以通过查看事件通道来做到这一点。

*producer.go*

```go
func main() {
    ...
    producer.ProduceChannel() <- record
    defer producer.Close()

    event := <-producer.Events()
    msg := event.(*kafka.Message)
    if msg.TopicPartition.Error != nil {
        log.Printf("Error sending message to cluster: %s\n", msg.TopicPartition.Error)
    } else {
        log.Printf("Message sent to topic %s (partition %d) at offset %d\n",
            *msg.TopicPartition.Topic, msg.TopicPartition.Partition, msg.TopicPartition.Offset)
    }
}
```


If you run `go run producer.go` you will see the following output:

如果您运行 `go run producer.go`，您将看到以下输出：

*output*

```text
2020/06/11 16:16:52 Message sent to topic test (partition 0) at offset 0
```


# Consuming messages 

# 消费消息

We’re now ready to consume some messages from the “test” topic! Setting up the consumer is very similar to how we set up the producer in the last section, but we also need to provide something called a `group.id`. I’ll touch on this property more in the next section.

我们现在准备使用来自“test”主题的一些消息！设置消费者与我们在上一节中设置生产者的方式非常相似，但我们还需要提供一个叫做“group.id”的东西。我将在下一节中更多地讨论这个属性。

*consumer.go*

**

```go
func main() {
    config := &kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
        "group.id":          "test-group",
    }

    consumer, err := kafka.NewConsumer(config)
    if err != nil {
        log.Fatalf("Error creating consumer: %s\n", err)
    }
}
```


Once we have created the consumer, we can subscribe to the “test”  topic and poll for events that have been pushed to the topic. We will  also set up a signal handler so that we can exit gracefully.

一旦我们创建了消费者，我们就可以订阅“测试”主题并轮询已推送到该主题的事件。我们还将设置一个信号处理程序，以便我们可以优雅地退出。

*consumer.go*

```go
func main() {
    ...
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

    consumer.Subscribe("test", nil)
    for {
        select {
        case <-sigChan:
            log.Println("Shutting down")
            consumer.Close()
            os.Exit(0)
        default:
            event := consumer.Poll(100)
            if event == nil {
                continue
            }

            switch e := event.(type) {
            case *kafka.Message:
                log.Printf("Received message with value: %s\n", e.Value)
            case kafka.OffsetsCommitted:
                log.Printf("Offsets committed: %s\n", e)
            }
        }
    }
}
```


If you run the consumer with `go run consumer.go` and then run the producer with `go run producer.go`, you will see something like this:

如果你用 `go run consumer.go` 运行消费者，然后用 `go run producer.go` 运行生产者，你将看到如下内容：

*consumer output*

*消费者输出*

```text
2020/06/11 16:38:59 Received message with value: hello world!
2020/06/11 16:39:03 Offsets committed: OffsetsCommitted (<nil>, [test[0]@1])
```


# Horizontally scaling with consumer groups

# 与消费者群体一起横向扩展

At some point, you may eventually be producing thousands of messages  per minute which could be a lot for just one consumer to manage. This is also completely ignoring the fact that you will want to have some sort  of fault tolerance if a consumer suddenly goes down. Consumer groups can help you address both of these issues!

在某些时候，您最终可能每分钟生成数千条消息，这对于一个消费者来说可能很多。这也完全忽略了这样一个事实，即如果消费者突然出现故障，您将需要某种容错能力。消费者团体可以帮助您解决这两个问题！

By having two different consumers within the same group, you are  effectively splitting the workload between the two. The offset for the  topic applies not to a specific consumer, but rather to its consumer  group. This means that if one consumer within the group processes a  message, the others will not consume that same message. Each consumer in a consumer group is assigned a set of topic partitions. To test this  out you will first need to [start up another Kafka broker](https://kafka.apache.org/quickstart#quickstart_multibroker) and create a partitioned topic like so:

通过在同一组中有两个不同的使用者，您可以有效地在两者之间分配工作量。主题的偏移量不适用于特定的消费者，而是适用于其消费者组。这意味着如果组中的一个消费者处理一条消息，其他消费者将不会消费同一条消息。消费者组中的每个消费者都被分配了一组主题分区。要对此进行测试，您首先需要[启动另一个 Kafka 代理](https://kafka.apache.org/quickstart#quickstart_multibroker) 并创建一个分区主题，如下所示：

*command line*

*命令行*

```text
bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 2 --topic test-partitioned
```


Start by changing the topic in the producer and the consumer to be  “test-partitioned”. Then start up two different instances of your  consumer by running `go run consumer.go` in two separate terminals. Now, if you run the producer a few times with `go run producer.go` you will see how the consumers share the work.

首先将生产者和消费者中的主题更改为“测试分区”。然后通过在两个不同的终端中运行 `go run consumer.go` 来启动你的消费者的两个不同实例。现在，如果您使用 `go run producer.go` 运行几次生产者，您将看到消费者如何共享工作。

*consumer one output*

*消费一输出*

```text
2020/06/12 12:22:18 Received message with value: hello world!
2020/06/12 12:22:19 Offsets committed: OffsetsCommitted (<nil>, [test-partitioned[0]@unset test-partitioned[1]@15])
2020/06/12 12:22:25 Received message with value: hello world!
2020/06/12 12:22:29 Offsets committed: OffsetsCommitted (<nil>, [test-partitioned[1]@16])
```


*consumer two output*

*消费两个输出*

```text
2020/06/12 12:22:29 Received message with value: hello world!
2020/06/12 12:22:32 Offsets committed: OffsetsCommitted (<nil>, [test-partitioned[0]@4])
2020/06/12 12:22:32 Received message with value: hello world!
2020/06/12 12:22:37 Offsets committed: OffsetsCommitted (<nil>, [test-partitioned[0]@5])
```


# Final thoughts 

#  最后的想法

This is only a small sample of what you can do with Kafka, but I’m  hoping to dig deeper into some of its other features and use cases in  future posts. Thanks for taking the time to read this post, and I hope  you’ll decide to stop by again soon! All of the code in this post is  available on my [GitHub](https://github.com/aaronraff/blog-code/tree/master/an-introduction-to-producing-and-consuming-kafka-messages-in-go).

这只是你可以用 Kafka 做的事情的一小部分，但我希望在以后的文章中更深入地挖掘它的一些其他特性和用例。感谢您花时间阅读这篇文章，希望您能尽快决定再次光临！这篇文章中的所有代码都可以在我的 [GitHub](https://github.com/aaronraff/blog-code/tree/master/an-introduction-to-production-and-sumption-kafka-messages-in-go)。

If you liked this post, it would mean a lot to me if you shared it with your friends! 

如果你喜欢这篇文章，如果你把它分享给你的朋友，那对我来说意义重大！
