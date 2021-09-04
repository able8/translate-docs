# Wrapping commands in Go

# 在 Go 中包装命令

You can find a lot of articles about Go that describe general aspects of it. Including the content on this blog. Today, I decided to prepare something different. I’ll tell you about one of my tasks and I’ll show you how I resolved it using Go. I thought it’d be useful to show the `exec` package and to tell a bit about the `ssh` command and learn AWS EE2 a bit better.

你可以找到很多关于 Go 的文章，它们描述了它的一般方面。包括本博客上的内容。今天，我决定准备一些不同的东西。我将告诉您我的一项任务，并向您展示我如何使用 Go 解决它。我认为展示 `exec` 包并介绍一些关于 `ssh` 命令并更好地学习 AWS EE2 会很有用。

AWS has a feature called [Amazon EC2 Instance Connect](https://aws.amazon.com/blogs/compute/new-using-amazon-ec2-instance-connect-for-ssh-access-to-your-ec2-instances/). You can use it to connect to your EC2 instance using an SSH client. The whole process has a few steps:

- get information about the instance (region, instance id, an availability zone)
- upload your public key to the EC2 instance
- Connect to the instance using the private key and SSH command

AWS 有一项名为 [Amazon EC2 Instance Connect](https://aws.amazon.com/blogs/compute/new-using-amazon-ec2-instance-connect-for-ssh-access-to-your-ec2-实例/)。您可以使用它通过 SSH 客户端连接到您的 EC2 实例。整个过程有几个步骤：

- 获取有关实例的信息（区域、实例 ID、可用区）
- 将您的公钥上传到 EC2 实例
- 使用私钥和 SSH 命令连接到实例

The problem we’re solving today is automating this process. After uploading the SSH key we have 60 seconds to connect to the EC2 instance. If you connect to a lot of EC2 instances and you have to repeat the same steps over and over you want to automate it.

我们今天要解决的问题是自动化这个过程。上传 SSH 密钥后，我们有 60 秒的时间连接到 EC2 实例。如果您连接到许多 EC2 实例，并且必须一遍又一遍地重复相同的步骤，那么您希望将其自动化。

My goal was to create an `ssh` replacement that accepts the same parameters and behaves as a regular `ssh` command but automates the whole setup process.

我的目标是创建一个 `ssh` 替代品，它接受相同的参数并像常规的 `ssh` 命令一样运行，但可以自动化整个设置过程。

The requirement is simple - the usage of the command should be as similar to the `ssh` command as possible. In a perfect world - it should be 100% replacement. Let’s try if we can achieve that. An example command will look like this:

```bash
./ec2-ssh HOSTNAME

# or
./ec2-ssh ec2-user@HOSTNAME # the default username is your OS user

# or
./ec2-ssh 192.168.0.1 -4 -k # accepts any parameter that ssh does
```


To upload our public key we need to know the availability zone, EC2 instance ID, user, and the public key itself. From the command parameters, we have the IP or the hostname.

要上传我们的公钥，我们需要知道可用区、EC2 实例 ID、用户和公钥本身。从命令参数中，我们有 IP 或主机名。

Let’s start with the username, host, and public key. There’s a `-G` parameter in the `ssh` command that prints all the configurations after evaluating host and match blocks. We can call the `ssh` command with all the parameters provided by the user and add `-G`. Then, we can parse the output and read all the data from it. In other words, we want to read this command’s output.

让我们从用户名、主机和公钥开始。 `ssh` 命令中有一个 `-G` 参数，用于在评估主机和匹配块后打印所有配置。我们可以使用用户提供的所有参数调用`ssh`命令并添加`-G`。然后，我们可以解析输出并从中读取所有数据。换句话说，我们想要读取这个命令的输出。

```bash
./ec2-ssh 192.168.0.1 -4 -k # from this command

ssh -G 192.168.0.1 -4 -k # we'll translate to this
```


We have to call the `ssh` command as a subprocess and read its output. Go has an `exec` package that contains the `Cmd` struct. This struct represents an external command and can be used for this purpose.

我们必须将 `ssh` 命令作为子进程调用并读取其输出。 Go 有一个包含 `Cmd` 结构的 `exec` 包。这个结构代表一个外部命令，可以用于这个目的。

```go
    cmd := exec.CommandContext(ctx, "ssh", args...)

    s := ""
    buff := bytes.NewBufferString(s)
    cmd.Stdout = buff
    cmd.Stderr = os.Stdout
    cmd.Stdin = os.Stdin

    if err := cmd.Run();err != nil {
        return nil, err
    }
```


The `Cmd` struct has `cmd.Stdout` field that’s the most important for us. It’s the place where we can forward the output of the command. This field accepts any `io.Writer` type so we put our buffer there. The next step is to put the parameters into the map from where we’ll retrieve values.

`Cmd` 结构具有 `cmd.Stdout` 字段，这对我们来说是最重要的。这是我们可以转发命令输出的地方。该字段接受任何 `io.Writer` 类型，因此我们将缓冲区放在那里。下一步是将参数放入映射中，我们将从中检索值。

```go
    res := map[string][]string{}

    scanner := bufio.NewScanner(buff)
    for scanner.Scan() {
        parts := strings.Split(scanner.Text(), " ")
        if len(parts) < 1 {
            continue
        }

        if _, exists := res[parts[0]];!exists {
            res[parts[0]] = []string{}
            continue
        }

        res[parts[0]] = append(res[parts[0]], strings.Join(parts[1:], " "))
    }

    return res, nil
```


We go line by line and put the data on the map. In the next function, we get the required information from the map: we need its IPv4 as well as the username.

我们一行一行地将数据放在地图上。在下一个函数中，我们从地图中获取所需的信息：我们需要它的 IPv4 以及用户名。

```go
func instanceInfoFromString(hostname, user string) (*instanceInfo, error) {
    info := &instanceInfo{
        username: user,
        host:     hostname,
    }

    err := info.resolveIP()
    if err != nil {
        return nil, err
    }
    return info, nil
}
```


We need the IP address because in later steps. We’ll need it to filter out irrelevant EC2 instances.
The next step is to find the public key that `ssh` will use to connect us to the EC2 instance. A list of possible SSH keys is available under the `identityfile` key in our map. We iterate over every item and check if it exists. If yes, then we return it.

我们需要 IP 地址，因为在后面的步骤中。我们需要它来过滤掉不相关的 EC2 实例。
下一步是找到“ssh”将用于将我们连接到 EC2 实例的公钥。我们地图中的“identityfile”键下提供了可能的 SSH 密钥列表。我们迭代每个项目并检查它是否存在。如果是，那么我们返回它。

```go
func existingKey(paths []string) (string, error) {
    for _, path := range paths {
        path, err := expandHomeDirectoryTilde(path)
        if err != nil {
            return "", err
        }

        if _, err := os.Stat(path);errors.Is(err, os.ErrNotExist) {
            continue
        }

        return path, nil
    }

    return "", errors.New("cannot find any ssh key")
}
```




Every key’s path starts (in general) with a tilde ( `~`) that’s means the user’s home directory. We had to write a function that expands the tilde to a full path. Why? The tilde is expanded by your shell’s `HOME` value. You can read how it works in more detail in [bash's docs](https://www.gnu.org/software/bash/manual/html_node/Tilde-Expansion.html) or this [SO answer](https://unix.stackexchange.com/questions/146671/does-always-equal-home/146697#146697). Let’s get back to the code.

每个键的路径（通常）以波浪号（`~`）开头，这意味着用户的主目录。我们必须编写一个函数来将波浪号扩展为完整路径。为什么？波浪号由 shell 的 `HOME` 值扩展。您可以在 [bash 的文档](https://www.gnu.org/software/bash/manual/html_node/Tilde-Expansion.html) 或这个 [SO answer](https://unix.stackexchange.com/questions/146671/does-always-equal-home/146697#146697)。让我们回到代码。

```go
     publicKey, err := getPublicKey(pk)
    if err != nil {
        return fmt.Errorf("cannot read the public key %s.pub. If you want to provide a custom key location, use the `-i` parameter", pk)
    }
```


In the listing below we attempt to read the public key. We need it to upload to the EC2 instance. This public key will be used to authenticate us. It means the EC2 instance has to know it before we’ll attempt to connect to it. AWS will put our public key to `~/.ssh/authorized_keys` file. We have only 60 seconds to connect to the instance. For more details on how the SSH authorization works, you can [visit this description](https://www.ssh.com/academy/ssh/public-key-authentication).

在下面的清单中，我们尝试读取公钥。我们需要将其上传到 EC2 实例。此公钥将用于对我们进行身份验证。这意味着 EC2 实例必须在我们尝试连接到它之前知道它。 AWS 会将我们的公钥放入`~/.ssh/authorized_keys` 文件。我们只有 60 秒的时间连接到实例。有关 SSH 授权如何工作的更多详细信息，您可以[访问此说明](https://www.ssh.com/academy/ssh/public-key-authentication)。

We have almost everything we need to connect to the EC2 instance. The only thing missing is the AWS region. We have the requirement that we are ass `ssh` command compatible as possible we cannot just add another parameter to our command. Instead, we’ll iterate over all regions and try to connect to every single instance. I know it’s not the most optimal way. If you have any idea how I can improve it - let me know in the comments section below.

我们几乎拥有连接到 EC2 实例所需的一切。唯一缺少的是 AWS 区域。我们要求我们尽可能与 `ssh` 命令兼容，我们不能只是向我们的命令添加另一个参数。相反，我们将遍历所有区域并尝试连接到每个实例。我知道这不是最佳方式。如果您对我如何改进它有任何想法 - 请在下面的评论部分告诉我。

```go
func setupEC2Instance(ctx context.Context, instance *instanceInfo, publicKey, region string) (bool, error) {
    cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
    if err != nil {
        return false, fmt.Errorf("cannot get config for AWS: %w", err)
    }

    client := ec2.NewFromConfig(cfg)

    ec2Instance, err := findEC2Instance(ctx, client, instance)
    if err != nil {
        return false, err
    }

    if ec2Instance == nil {
        return false, nil
    }

    status, err := instanceStatus(ctx, client, *ec2Instance)
    if err != nil {
        return false, fmt.Errorf("cannot get the instance status: %w", err)
    }

    connect := ec2instanceconnect.NewFromConfig(cfg)
    out, err := connect.SendSSHPublicKey(ctx, &ec2instanceconnect.SendSSHPublicKeyInput{
        AvailabilityZone: status.AvailabilityZone,
        InstanceId:       ec2Instance.InstanceId,
        InstanceOSUser:   &instance.username,
        SSHPublicKey:     &publicKey,
    })

    if err != nil {
        return false, fmt.Errorf("cannot upload the public key: %w", err)
    }

    if !out.Success {
        return false, fmt.Errorf("unsuccessful uploaded the public key")
    }

    return true, nil
}

```


In the code above, we’re configuring the AWS client, trying to find our EC2 instance in the selected region. If everything goes fine, we’re uploading our public key. If it succeeds as well, we’re ready to connect. Two functions are new here: `findEC2Instance` and `instanceStatus`.

在上面的代码中，我们正在配置 AWS 客户端，尝试在所选区域中找到我们的 EC2 实例。如果一切顺利，我们将上传我们的公钥。如果它也成功了，我们就可以连接了。这里有两个新函数：`findEC2Instance` 和 `instanceStatus`。

The first one is quite obvious - it finds our EC2 instance using the IP address we retrieved earlier.

第一个很明显 - 它使用我们之前检索到的 IP 地址找到我们的 EC2 实例。

```go
func findEC2Instance(ctx context.Context, client *ec2.Client, info *instanceInfo) (*types.Instance, error) {
    resp, err := client.DescribeInstances(ctx, &ec2.DescribeInstancesInput{
        Filters: []types.Filter{
            {
                Name:   strp("private-ip-address"),
                Values: []string{info.ipAddress},
            },
        },
    })

    if err != nil {
        return nil, fmt.Errorf("cannot contact with AWS API: %w", err)
    }

    for _, r := range resp.Reservations {
        for _, inst := range r.Instances {
            if *inst.PrivateIpAddress == info.ipAddress {
                return &inst, nil
            }
        }
    }
    return nil, nil
}
```


When we know that the instance exists and we have its reference, we can check its status and this is when the `instanceStatus` comes into play.

当我们知道实例存在并且我们有它的引用时，我们可以检查它的状态，这就是“instanceStatus”发挥作用的时候。

```go
func instanceStatus(ctx context.Context, client *ec2.Client, instance types.Instance) (types.InstanceStatus, error) {
    descResp, err := client.DescribeInstanceStatus(ctx, &ec2.DescribeInstanceStatusInput{
        InstanceIds: []string{*instance.InstanceId},
    })

    if err != nil {
        return types.InstanceStatus{}, err
    }

    status := descResp.InstanceStatuses[0]
    return status, nil
}
```




The `client.DescribeInstanceStatus` returns a few very valuable information for us: the instance’s available zone and the instance’s ID. Both values are required while uploading the SSH public key.

`client.DescribeInstanceStatus` 为我们返回了一些非常有价值的信息：实例的可用区域和实例的 ID。上传 SSH 公钥时需要这两个值。

At this point, we are ready to connect to the EC2 instance! That’s quite simple - had to execute the `ssh` command with all our parameters. We forward all the output to the standard output and do the same with the std input. Thanks to this, we’ll be able to interact with the `ssh` command as usual.

此时，我们已准备好连接到 EC2 实例！这很简单 - 必须使用我们的所有参数执行 `ssh` 命令。我们将所有输出转发到标准输出并对 std 输入执行相同操作。多亏了这一点，我们将能够像往常一样与 `ssh` 命令交互。

```go
func connectToInstance(ctx context.Context, params []string) error {
    cmd := exec.CommandContext(ctx, "ssh", params...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stdout
    cmd.Stdin = os.Stdin

    if err := cmd.Run();err != nil {
        if exiterr, ok := err.(*exec.ExitError);ok {
            // terminated by Control-C so ignoring
            if exiterr.ExitCode() == 130 {
                return nil
            }
        }

        return fmt.Errorf("error while connecting to the instance: %w", err)
    }

    return nil
}
```


And that’s all! The whole source code is [available on Github](https://github.com/bkielbasa/ec2-ssh). From now, you can replace your `ssh` command with `ec2-ssh` while working with AWS EC2 instances. If you have any questions or suggestions, feel free to use the comments section below. 

就这样！整个源代码在 [Github 上可用](https://github.com/bkielbasa/ec2-ssh)。从现在开始，您可以在使用 AWS EC2 实例时用 `ec2-ssh` 替换你的 `ssh` 命令。如果您有任何问题或建议，请随时使用下面的评论部分。

