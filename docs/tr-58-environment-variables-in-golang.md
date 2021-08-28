# Environment variables in Golang

# Golang 中的环境变量

Learn about environment variables and different ways to use them in your Golang application.

了解环境变量以及在 Golang 应用程序中使用它们的不同方式。

September 28, 2020 From: https://www.loginradius.com/blog/async/environment-variables-in-golang/

## Before You Get Started

## 开始之前

This tutorial assumes you have:
- A basic understanding of Go Language
- Latest Golang version installed on your system
- A few minutes of your time.

本教程假设您有：
- 对 Go 语言有基本的了解
- 您的系统上安装了最新的 Golang 版本
- 几分钟的时间。

In this article, we will know about environment variables and why to  use them. And will access them in a Go application using inbuilt and  third-party packages.

在本文中，我们将了解环境变量以及为什么要使用它们。并将使用内置和第三方包在 Go 应用程序中访问它们。

## What are environment variables?

## 什么是环境变量？

Environment variables are key-value pair on a system-wide level, and  running processes can access that. These are often used to make the same program behave differently in different deployment environments like  PROD, DEV, or TEST. Storing configuration in the environment is one of the principles of a  twelve-factor app. It enables applications to be portable.

环境变量是系统范围内的键值对，运行的进程可以访问它。这些通常用于使相同的程序在不同的部署环境（如 PROD、DEV 或 TEST）中表现不同。在环境中存储配置是十二要素应用程序的原则之一。它使应用程序成为可移植的。

## Why should you use environment variables

## 为什么要使用环境变量

- If you are using the sensitive information in the code, then all the unauthorized users who have access to the code will have sensitive  data, you might not want that.
- If you are using the code versioning tool like `git`, you may push your DB credentials with the code, and it will become public.
- If you are managing variables in one place, in case of any changes,  you don't have to change it in all the places in application code.
- You can manage multiple deployment environments like PROD, DEV, or  TEST. Environment variables are easy to change between deploys without  changing any application code.

- 如果您在代码中使用敏感信息，那么所有有权访问代码的未授权用户都将拥有敏感数据，您可能不希望这样。
- 如果你正在使用像`git`这样的代码版本控制工具，你可以用代码推送你的数据库凭证，它会变成公开的。
- 如果您在一处管理变量，万一发生任何变化，您不必在应用程序代码的所有地方都进行更改。
- 您可以管理多个部署环境，如 PROD、DEV 或 TEST。环境变量很容易在部署之间更改，而无需更改任何应用程序代码。

> Never forget to include your environment variable files in the .gitignore

> 永远不要忘记在 .gitignore 中包含您的环境变量文件

## Inbuilt OS package

## 内置操作系统包

You don't need any external package to access the environment variables in Golang, and you can do that with the standard `os` package. Below is the list of functions related to environment variables and there uses.

你不需要任何外部包来访问 Golang 中的环境变量，你可以使用标准的 `os` 包来做到这一点。以下是与环境变量相关的函数列表及其用途。

- `os.Setenv()` sets the value of an environment value.
- `os.Getenv()` gets the value environment variable named by the key.
- `os.Unsetenv()` delete a single environment value named by the key, if we try to get that environment value using `os.Getenv()` it will return an empty value.
- `os.ExpandEnv` replaces ${var} or $var in the string as  per the values of environment variables. If any environment variable is  not present an empty string will replace it.
- `os.LookupEnv()` gets the value environment variable  named by the key. If the variable is not present in the system, the  returned value will be empty, and the boolean will be false. Otherwise,  it returns the value (which can be empty), and the boolean is true.

- `os.Setenv()` 设置环境值的值。
- `os.Getenv()` 获取由键命名的值环境变量。
- `os.Unsetenv()` 删除由键命名的单个环境值，如果我们尝试使用 `os.Getenv()` 获取该环境值，它将返回一个空值。
- `os.ExpandEnv` 根据环境变量的值替换字符串中的 ${var} 或 $var。如果任何环境变量不存在，一个空字符串将替换它。
- `os.LookupEnv()` 获取由键命名的值环境变量。如果系统中不存在该变量，则返回值将为空，布尔值为 false。否则，它返回值（可以为空），布尔值为真。

> os.Getenv() will return an empty string if the environment variable  is not present, to distinguish between an empty value and an unset  value, use LookupEnv.

> 如果环境变量不存在，os.Getenv() 将返回一个空字符串，为了区分空值和未设置值，请使用 LookupEnv。

Now let's use all the above functions in our code. Create a main.go file in an empty folder.

现在让我们在代码中使用上述所有函数。在空文件夹中创建一个 main.go 文件。

```
package main

import (
  "fmt"
  "os"
)

func main() {
  // Set Environment Variables
  os.Setenv("SITE_TITLE", "Test Site")
  os.Setenv("DB_HOST", "localhost")
  os.Setenv("DB_PORT", "27017")
  os.Setenv("DB_USERNAME", "admin")
  os.Setenv("DB_PASSWORD", "password")
  os.Setenv("DB_NAME", "testdb")

  // Get the value of an Environment Variable
  host := os.Getenv("SITE_TITLE")
  port := os.Getenv("DB_HOST")
  fmt.Printf("Site Title: %s, Host: %s\n", host, port)

  // Unset an Environment Variable
  os.Unsetenv("SITE_TITLE")
  fmt.Printf("After unset, Site Title: %s\n", os.Getenv("SITE_TITLE"))

  //Checking that an environment variable is present or not.
  redisHost, ok := os.LookupEnv("REDIS_HOST")
  if !ok {
    fmt.Println("REDIS_HOST is not present")
  } else {
    fmt.Printf("Redis Host: %s\n", redisHost)
  }

  // Expand a string containing environment variables in the form of $var or ${var}
  dbURL := os.ExpandEnv("mongodb://${DB_USERNAME}:${DB_PASSWORD}@$DB_HOST:$DB_PORT/$DB_NAME")
  fmt.Println("DB URL: ", dbURL)
}
```


Below is the output when we run `go run main.go` in our terminal

下面是我们在终端中运行 `go run main.go` 时的输出

```
go run main.go

//output
Site Title: Test Site, Host: localhost
After unset, Site Title: 27017
REDIS_HOST is not present
DB URL:  mongodb://admin:password@localhost:27017/testdb
```


There are two more functions `os.Clearenv` and `os.Environ()` let's use them also in a separate program.

还有两个函数 `os.Clearenv` 和 `os.Environ()` 让我们在一个单独的程序中使用它们。

- `os.Clearenv` deletes all environment variables, It can be useful to clean up the environment for tests 

- `os.Clearenv` 删除所有环境变量，清理测试环境很有用

- `os.Environ()` returns a slice of the string containing all the environment variables in the form of key=value.

- `os.Environ()` 以 key=value 的形式返回包含所有环境变量的字符串片段。

```
package main

import (
  "fmt"
  "os"
  "strings"
)

func main() {

  // Environ returns a slice of string containing all the environment variables in the form of key=value.
  for _, env := range os.Environ() {
    // env is
    envPair := strings.SplitN(env, "=", 2)
    key := envPair[0]
    value := envPair[1]

    fmt.Printf("%s : %s\n", key, value)
  }

  // Delete all environment variables
  os.Clearenv()

  fmt.Println("Number of environment variables: ", len(os.Environ()))
}
```


The above function will list all the environment variables available in the system, including `NAME` and `DB_HOST`. Once we run `os.Clearenv()` it will clear all the environment variables for the running process.

上述函数将列出系统中所有可用的环境变量，包括`NAME`和`DB_HOST`。一旦我们运行`os.Clearenv()`，它将清除正在运行的进程的所有环境变量。

## godotenv package

## godotenv 包

The Ruby dotenv project inspires [GoDotEnv](https://github.com/joho/godotenv) package, it loads the environment variables from a .env file

Ruby dotenv 项目启发了 [GoDotEnv](https://github.com/joho/godotenv) 包，它从 .env 文件加载环境变量

Let's create a .env file in which we will have all our configurations.

让我们创建一个 .env 文件，我们将在其中包含所有配置。

```
# .env file
# This is a sample config file

SITE_TITLE=Test Site

DB_HOST=localhost
DB_PORT=27017
DB_USERNAME=admin
DB_PASSWORD=password
DB_NAME=testdb
```


Then in the main.go file we will use godotenv to load the environment variables.

然后在 main.go 文件中，我们将使用 godotenv 加载环境变量。

> we can load multiple env files at once also. godotenv also supports YAML.

> 我们也可以一次加载多个 env 文件。 godotenv 也支持 YAML。

```
// main.go
package main

import (
  "fmt"
  "log"
  "os"

  "github.com/joho/godotenv"
)

func main() {

  // load .env file from given path
  // we keep it empty it will load .env from current directory
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }

  // getting env variables SITE_TITLE and DB_HOST
  siteTitle := os.Getenv("SITE_TITLE")
  dbHost := os.Getenv("DB_HOST")

  fmt.Printf("godotenv : %s = %s \n", "Site Title", siteTitle)
  fmt.Printf("godotenv : %s = %s \n", "DB Host", dbHost)
}
```


Open the terminal and run the `main.go`

打开终端并运行`main.go`

```
go run main.go

// output
godotenv : Site Title = Test Site
godotenv : DB Host = localhost
```


## Viper package

> Viper is a complete configuration solution for Go applications  including twelve-factor apps. It is designed to work within an  application and can handle all types of configuration needs and formats.

> Viper 是 Go 应用程序的完整配置解决方案，包括十二要素应用程序。它旨在在应用程序中工作，可以处理所有类型的配置需求和格式。

[Viper](https://github.com/spf13/viper) supports many file formats to load environment variables, e.g., Reading from JSON, TOML,  YAML, HCL, envfile and Java properties config files. So in this example, we will look at how to load environment variables from a YAML file.

[Viper](https://github.com/spf13/viper) 支持多种文件格式来加载环境变量，例如从 JSON、TOML、YAML、HCL、envfile 和 Java 属性配置文件中读取。所以在这个例子中，我们将看看如何从 YAML 文件加载环境变量。

> YAML is a human-readable data-serialization language. It is commonly  used for configuration files and in applications where data is being  stored or transmitted.

> YAML 是一种人类可读的数据序列化语言。它通常用于配置文件和存储或传输数据的应用程序。

Let's create our config.yaml and main.go in an empty folder.

让我们在一个空文件夹中创建我们的 config.yaml 和 main.go。

```
# config.yaml
SITE:
  TITLE: Test Site

DB:
  HOST: "localhost"
  PORT: "27017"
  USERNAME: "admin"
  PASWORD: "password"
  NAME: "testdb"
```


In the below code, we are using Viper to load environment variables  from a config.yaml. We can load the config file from any path we want. We can also set the default values for any environment variable if any  environment variable is not available in the config file.

在下面的代码中，我们使用 Viper 从 config.yaml 加载环境变量。我们可以从我们想要的任何路径加载配置文件。如果配置文件中没有任何环境变量，我们还可以为任何环境变量设置默认值。

```
// main.go
package main

import (
  "fmt"
  "log"
  "os"

  "github.com/spf13/viper"
)

func main() {

  // Set the file name of the configurations file
  viper.SetConfigName("config")

  // Set the path to look for the configurations file
  viper.AddConfigPath(".")

  // Enable VIPER to read Environment Variables
  viper.AutomaticEnv()

  viper.SetConfigType("yml")

  if err := viper.ReadInConfig();err != nil {
    fmt.Printf("Error reading config file, %s", err)
  }

  // Set undefined variables
  viper.SetDefault("DB.HOST", "127.0.0.1")

  // getting env variables DB.PORT
  // viper.Get() returns an empty interface{}
  // so we have to do the type assertion, to get the value
  DBPort, ok := viper.Get("DB.PORT").(string)

  // if type assert is not valid it will throw an error
  if !ok {
    log.Fatalf("Invalid type assertion")
  }

  fmt.Printf("viper : %s = %s \n", "Database Port", DBPort)
}
```


Open the terminal and run the `main.go`

打开终端并运行`main.go`

```
go run main.go

// output
viper : Database Port = 27017
```


## Conclusion 

##  结论

Using environment variables is an excellent way to handle  configuration in our application. Overall, it provides you with easy  configuration, better security, multiple deployment environments and  fewer production mistakes.

使用环境变量是在我们的应用程序中处理配置的绝佳方式。总的来说，它为您提供了简单的配置、更好的安全性、多种部署环境和更少的生产错误。

Now you can manage environment variables in your go application, and  You can found the complete code used in this tutorial on our [Github Repo](https://github.com/LoginRadius/engineering-blog-samples/tree/master/GoLang/EnvironmentVariables)

现在您可以在您的 go 应用程序中管理环境变量，您可以在我们的 [Github Repo](https://github.com/LoginRadius/engineering-blog-samples/tree/master/GoLang) 上找到本教程中使用的完整代码/环境变量）

### Related Posts

###  相关文章

#### [Build and Push Docker Images with Go](https://www.loginradius.com/blog/async/build-push-docker-images-golang/)

#### [使用 Go 构建和推送 Docker 镜像](https://www.loginradius.com/blog/async/build-push-docker-images-golang/)
