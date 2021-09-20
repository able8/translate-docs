# Using Go instead of bash for scripts    

From: https://blog.kowalczyk.info/article/4b1f9201181340099b698246857ea98d/using-go-instead-of-bash-for-scripts.html

I like to automate my programming work.  

In every programming project I ended up writing bash (on Unix and Mac) and batch / PowerShell (on Windows) scripts.  

I settled on a convention to put scripts in a directory `s` (it's short, fast to type and a shortcut for `scripts`).  

I had `./s/run.sh` to run the program locally.`./s/tests.sh` to run tests. `./s/deploy.sh` to deploy web apps to the server etc.  

It worked but I wasn't quite happy.  

For cross-platform projects I had to write the same script twice (`./s/run.sh` and `./s/run.bat`).  

I write those scripts so infrequently that I every time I need to re-learn basics. How do I declare a function? How do I write `if`? How do I write a loop?   

In bash, simple things are complicated and non-simple things are very complicated.  

This article describes how I replaced bash scripts with a single, multiple-purpose Go program.   

You can see a full example at https://github.com/kjk/notionapi/tree/master/do  

# Replacing bash with Go 

One day it hit me: I would rather write the helper scripts in Go.  

Go is cross-platform; I don't have to write the same thing twice.  

I write in Go daily so I can implement simple things quickly.  

In Go simple things are simple and complicated things are possible.  

The one drawback is more lines of code but the difference is immaterial. Those are short programs either way.  

This article describes a system I refined by using it in multiple projects.  

## Establishing conventions  

It's all about convenience so less typing is better.   

Establishing conventions to share between multiple projects frees mental energy for more important things.  

The system I settled on is:  

- `do` directory contains a single, multiple-purpose Go program. A single program that does many things is better suited to Go than multiple programs as it makes it easy to share helper functions    

- to run it: `cd do; go run . ${flags}`    

- to make things easier to type, I have 

  ```
  do/do.sh
  ```

  :

  ```
  #!/bin/bash
  
  cd do
  go run -race . $@
  ```

- for Windows I have 

  ```
  do\do.bat
  ```

  :

  ```
  @cd do
  @go run -race . %*
  ```

- to make things even easier to type, I add the following to 

  ```
  ~/.bash_profile
  ```

   (on Mac):

  ```
  function doit() {
  	if [ -f ./do.sh ]; then
  		./do.sh $@
  	elif [ -f ./do/do.sh ]; then
  		./do/do.sh $@
  	else
  		echo "no do.sh or do/do.sh found"
  	fi
  }
  ```

  I can then type `doit ${args}` to launch either `./do.sh` or `./do/do.sh` (whichever exists).       

In every project I can type  `doit -run` which executes `./do/do.sh -run` which executes `cd do; go run . -run`.  

In the old system, that would be `./s/run.sh` or `.\s\run.bat`.  

Other cmd-line arguments trigger other actions e.g. `doit -test`, `doit -deploy` etc.  

If I forget which flags are available, `do` without arguments prints them all.  

## A structure of `do` program  

Main function checks cmd-line arguments and calls the right function to perform a given command.  

Here's an implementation of dispatching two commands: `-run` and `-deploy`.  

```
func main() {
	cdToTopDir()
	fmt.Printf("topDir: '%s'\n", topDir())

	var (
		flgRun    bool
		flgDeploy bool
	}
	
	{
		flag.BoolVar(&flgRun, "run", false, "runs the program")
		flag.BoolVar(&flgDeploy, "deploy", false, "deploys to production")
		flag.Parse()
	}

	if flgRun {
		doRun()
		return
	}

	if flgDeploy {
		doDeploy()
		return
	}

  // this prints available flags
	flag.Usage()
}
```

## Running from a known current directory  

When we run the program, we're inside `do` directory  

It's important to know what is the current director so that when we refer to files in the project, we know their path.  

By convention I set current directory to be top directory of the project.  

The first thing that the program does is call `cdToTopDir()` which fixes the current directory to this known location.  

The simplest implementation:  

```
func cdToTopDir() {
	err := os.Chdir("..")
	must(err)
}
```

This relies on knowledge that we execute the program with `cd do; go run . ${args}`.  

I also print the absolute path of current directory at the beginning to make sure it is correct.  

## Crashing on errors is fine  

In a regular Go program, handling errors by propagating them to callers is key for writing robust software.  

In short scripts it's ok to `panic` when error happens. It makes for shorter code. `panic` prints the callstack which is handy when investigating unexpected errors.  

I have a helper function `must(err error)` that panics if err is not nil:  

```
func must(err error) {
	if err != nil {
		fmt.Printf("err: %s\n", err)
		panic(err)
	}
}
```

Here's how to use it:  

```
func readFile(path string) []byte {
	d, err := ioutil.ReadFile(path)
	must(err)
	return d
}
```

## Executing programs  

A common thing to do is executing other programs. For example,  `do -run` would typically execute `go build . -o myapp` and `./myapp.`  

Go has an excellent `os/exec` package for that:  

```
cmd := exec.Command("go", "build", ".", "-o", "myapp")
err := cmd.Run()
must(err)

cmd = exec.Command("./myapp")
err = cmd.Run()
must(err)
```

Other useful things possible with `exec.Cmd`:  

- set working directory of the executed program

  ```
  cmd := exec.Command("./myapp")
  cmd.Dir = "working/directory"
  ```

- setting environment variables

  ```
  cmd.Env = os.Environ()
  cmd.Env = append(cmd.Env, "GOOS=linux", "GOARCH=amd64")
  ```

- get the output of the command

  ```
  cmd := exec.Command("ls", "-lah")
  // CombinedOutput() calls Run() and returns captured stdout / stderr as []byte
  out, err := cmd.CombinedOutput()
  must(err)
  fmt.Printf("output of ls:\n%s\n", string(out))
  ```

- sometimes you don't want to block waiting for the program to finish. Very true for launching Windows GUI programs:

  ```
  func openNotepadWithFile(path string) {
  	cmd := exec.Command("notepad.exe", path)
  	err := cmd.Start() // this starts the programs but doesn't wait for it to finish
  	must(err)
  }
  ```

- if you 

  ```
  Start()
  ```

   a process, you might want to ensure it's killed on exit:

  ```
  func main() {
  	// ...
  	err := cmd.Start()
  	must(err)
  
  	// ensure to kill the process upon exit
  	defer cmd.Process.Kill()
  }
  ```

- logging the commands we execute

  ```
  fmt.Printf("Running: %s\n", strings.Join(cmd.Args[1:], " "))
  ```

- seeing program's stdout and stderr while it's executing

  ```
  // set program's stdout / stderr ot console's stdout/stderr to
  // see what it prints
  // incompatible with capturing stdout / stderr with `CombinedOutput()`
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  ```

# Helper functions  

Scripts in different projects often need the same functionality.   

I keep common functions in a separate file `util.go` so that I can quickly bootstrap new project.  

Here are a few common helper functions.  

### assert  

Inspired by C, panics if condition is not true. Use to verify you get expected results.  

```
func assert(cond bool, format string, args ...interface{}) {
	if cond {
		return
	}
	s := fmt.Sprintf(format, args...)
	panic(s)
}
```

### logf  

To be used instead of `fmt.Printf`. The advantage is that if we want to e.g. start logging to file, we need to change just `logf` function.  

```
// a centralized place allows us to tweak logging, if need be
func logf(format string, args ...interface{}) {
	if len(args) == 0 {
		fmt.Print(format)
		return
	}
	fmt.Printf(format, args...)
}
```

### openBrowser  

When working on backends for web apps it's convenient to auto-open the web site in the browser when starting the app locally.  

```
// openBrowsers open web browser with a given url
// (can be http:// or file://)
func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	must(err)
}
```

### readZipFile  

Reads all files in a zip file and returns them as a map from file name to content.  

```
func readZipFile(path string) map[string][]byte {
	r, err := zip.OpenReader(path)
	must(err)
	defer r.Close()
	res := map[string][]byte{}
	for _, f := range r.File {
		rc, err := f.Open()
		must(err)
		d, err := ioutil.ReadAll(rc)
		must(err)
		rc.Close()
		res[f.Name] = d
	}
	return res
}
```

### readFile, writeFile  

Shorter way to read / write files.  

```
func readFile(path string) []byte {
	d, err := ioutil.ReadFile(path)
	must(err)
	return d
}

func writeFile(path string, data []byte) {
	err := ioutil.WriteFile(path, data, 0666)
	must(err)
}
```

### getHomeDir  

Returns a path of the user's home directory.  

```
func getHomeDir() string {
	s, err := os.UserHomeDir()
	must(err)
	return s
}
```

### cpFile  

Equivalent of `cp` in bash.  

```
func cpFile(dstPath, srcPath string) {
	d, err := ioutil.ReadFile(srcPath)
	must(err)
	err = ioutil.WriteFile(dstPath, d, 0666)
	must(err)
}
```

### checkGitClean  

To prevent accidental deploys, my scripts use `checkGitClean` and refuse to deploy if there are un-commited changes to working directory:  

```
var (
	verbose bool
)

func runCmd(cmd *exec.Cmd) string {
	if verbose {
		fmt.Printf("> %s\n", strings.Join(cmd.Args, " "))
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("%s failed with '%s'. Output:\n%s\n", strings.Join(cmd.Args, " "), err, string(out))
	}
	must(err)
	if verbose && len(out) > 0 {
		fmt.Printf("%s\n", out)
	}
	return string(out)
}

func gitStatus(dir string) string {
	cmd := exec.Command("git", "status")
	if dir != "" {
		cmd.Dir = dir
	}
	return runCmd(cmd)
}

func checkGitClean(dir string) {
	s := gitStatus(dir)
	expected := []string{
		"On branch master",
		"Your branch is up to date with 'origin/master'.",
		"nothing to commit, working tree clean",
	}
	for _, exp := range expected {
		if !strings.Contains(s, exp) {
			fmt.Printf("Git repo in '%s' not clean.\nDidn't find '%s' in output of git status:\n%s\n", dir, exp, s)
			os.Exit(1)
		}
	}
}
```

### createZipFile  

This is a helper to create a .zip archive with the content of one or more directories or files.  

Example use: `createZipFile("archive.zip", ".", "myapp", "www")`  

This creates [`archive.zip`](http://archive.zip) with the content of `myapp` file and `www` directory. Those files are located in current (`.`) directory.  

In fairness, would be shorter to sub-launch `zip` program, but I like the control.  

```
func zipAddFile(zw *zip.Writer, zipName string, path string) {
	zipName = filepath.ToSlash(zipName)
	d, err := ioutil.ReadFile(path)
	must(err)
	w, err := zw.Create(zipName)
	_, err = w.Write(d)
	must(err)
	if verbose {
		fmt.Printf("  added %s from %s\n", zipName, path)
	}
}

func zipDirRecur(zw *zip.Writer, baseDir string, dirToZip string) {
	dir := filepath.Join(baseDir, dirToZip)
	files, err := ioutil.ReadDir(dir)
	must(err)
	for _, fi := range files {
		if fi.IsDir() {
			zipDirRecur(zw, baseDir, filepath.Join(dirToZip, fi.Name()))
		} else if fi.Mode().IsRegular() {
			zipName := filepath.Join(dirToZip, fi.Name())
			path := filepath.Join(baseDir, zipName)
			zipAddFile(zw, zipName, path)
		} else {
			path := filepath.Join(baseDir, fi.Name())
			s := fmt.Sprintf("%s is not a dir or regular file", path)
			panic(s)
		}
	}
}

func createZipFile(dst string, baseDir string, toZip ...string) {
	removeFile(dst)
	if len(toZip) == 0 {
		panic("must provide toZip args")
	}
	if verbose {
		fmt.Printf("Creating zip file %s\n", dst)
	}
	w, err := os.Create(dst)
	must(err)
	defer w.Close()
	zw := zip.NewWriter(w)
	must(err)
	for _, name := range toZip {
		path := filepath.Join(baseDir, name)
		fi, err := os.Stat(path)
		must(err)
		if fi.IsDir() {
			zipDirRecur(zw, baseDir, name)
		} else if fi.Mode().IsRegular() {
			zipAddFile(zw, name, path)
		} else {
			s := fmt.Sprintf("%s is not a dir or regular file", path)
			panic(s)
		}
	}
	err = zw.Close()
	must(err)
}
```

### wc -l  

I like to know how big my programs are as measured by lines of code.  

On Unix simple stats can be done with `find . -name "*.go" | xargs wc -l`  

Similar  functionality in Go is significantly larger. The good thing is that it's cross-platform, more flexible and once written can be easily added to  more projects.  

Different projects want to count different files / directories so I built a flexible system that allows combining (with `and` and `or`) file filter functions.  

I wrote a helper library https://github.com/kjk/u.  

Here's how I use it in a [real project](https://github.com/kjk/notionapi/blob/master/do/wc.go):  

`filter` is a file filter function that tells us to count `.go`, `.js`, `.html` and `.css` files in all sub-directories but to  exclude `node_modules` and `tmpdata` directories because they contain files not written by me:  

```
package main

import (
	"fmt"

	"github.com/kjk/u"
)

var srcFiles = u.MakeAllowedFileFilterForExts(".go", ".js", ".html", ".css")
var excludeDirs = u.MakeExcludeDirsFilter("node_modules", "tmpdata")
var filter = u.MakeFilterAnd(srcFiles, excludeDirs)

func doLineCount() int {
	stats := u.NewLineStats()
	recursive := true
	err := stats.CalcInDir(".", filter, recursive)
	if err != nil {
		fmt.Printf("doLineCount: stats.wcInDir() failed with '%s'\n", err)
		return 1
	}
	u.PrintLineStats(stats)
	return 0
}
```

# More Go resources  

- [Essential Go](https://www.programming-books.io/essential/go/) is a free, comprehensive book about Go that I maintain    

​                                                            Written on Aug 16 2021.                                        Topics: [go](https://blog.kowalczyk.info/tag/go).                                                                            

​                                                            [home](https://blog.kowalczyk.info/)                                                        

​                    Found a mistake, have a comment?                    [Let me know](https://blog.kowalczyk.info/article/4b1f9201181340099b698246857ea98d/using-go-instead-of-bash-for-scripts.html#).                
