package main

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"
)

var outString string

const UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:92.0) Gecko/20100101 Firefox/92.0"

func sendRequest(ctx context.Context, method, urlStr string, body io.Reader, f func(*http.Request) error) (*http.Response, []byte, error) {
	req, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		return nil, nil, fmt.Errorf("http.NewRequest error: %v", err)
	}
	req = req.WithContext(ctx)
	if f != nil {
		if err := f(req); err != nil {
			return nil, nil, fmt.Errorf("f error: %v", err)
		}
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("http request error: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("ioutil.ReadAll error: %v", err)
	}
	return resp, respBody, nil
}

// Use google translate api
func Translate(source, sourceLang, targetLang string) error {

	// [[["MkEWBc","[[\"asdfasd\\nadsfasd\\nasdfas\\nasdasd\\ndsfasd\\n\\n\\nadsfasd\\nasdf\\ndsf\\n\\n\\n\\n\\n\\n\\nasd\\n\\nafasd\",\"en\",\"zh-CN\",true],[null]]",null,"generic"]]]
	reqStr := fmt.Sprintf(`[[["MkEWBc","[[\"%s\",\"%s\",\"%s\",true],[null]]",null,"generic"]]]`, source, sourceLang, "zh")
	reqStr = strings.Replace(reqStr, "\n", `\\n`, -1)
	param := url.Values{
		"f.req": {reqStr},
	}
	log.Println(reqStr)
	body := strings.NewReader(param.Encode())

	log.Println(body)

	req, err := http.NewRequest("POST", "https://translate.google.cn/_/TranslateWebserverUi/data/batchexecute", body)
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	req.Header.Set("Referer", "https://translate.google.cn/")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Errorf("http request error: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("ioutil.ReadAll error: %v", err)
	}

	fmt.Println(string(respBody))
	// os.Exit(1)

	if err != nil {
		return errors.New("failed to read response body")
	}

	var result []interface{}
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		// fmt.Println(err, string(respBody))
		return errors.New("Error unmarshaling data" + err.Error())
	}

	t := result[0].([]interface{})
	t1 := t[2].([]interface{})

	log.Println(t1)

	os.Exit(0)

	var input, output string
	if len(result) > 0 {
		for _, slice := range result[0].([]interface{}) {
			resultSlice := slice.([]interface{})
			translatedText, sourceText := resultSlice[0].(string), resultSlice[1].(string)

			log.Println(translatedText)

			// Do not translate these lines
			if strings.Contains(sourceText, "![") || strings.Contains(sourceText, "----") || strings.Contains(sourceText, "From: http") {
				input = input + sourceText
				continue
			}

			if strings.Contains(sourceText, "```") && strings.Contains(input, "```") {
				input = input + sourceText
				outString = outString + input + "\n"
				input, output = "", ""
				continue
			}

			if strings.Contains(input, "```") {
				input = input + sourceText
				continue
			}

			output = output + translatedText
			if len(sourceText) > 0 && sourceText[len(sourceText)-1:] == "\n" {
				input = input + sourceText
			} else {
				input = input + sourceText + " "
			}

			if strings.Contains(sourceText, "\n\n") {
				outString = outString + input + output
				input, output = "", ""
			}
		}

		outString = outString + input + "\n\n" + output + "\n\n"
		return nil
	} else {
		return errors.New("no translated data in responce")
	}
}

func check(err error) {
	if err != nil {
		log.Fatalf("Error: %v\n", err.Error())
	}
}

func main() {
	inputFile := flag.String("f", "", "input file")
	flag.Parse()

	if *inputFile == "" {
		log.Fatalln("Error, no input file. Please use `-f filename` to select a file.")
	} else {
		log.Printf("Input file is %q\n", *inputFile)
	}

	file, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines, translateLines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	count := 0
	for _, line := range lines {
		count = count + len(line)
		// Do not break the code block
		if count+len(line) > 4500 && strings.Count(strings.Join(translateLines, "\n"), "```")%2 == 0 {
			log.Println("Translating ...")
			err = Translate(strings.Join(translateLines, "\n"), "en", "zh-CN")
			check(err)
			time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
			count = len(line)
			translateLines = nil
		}
		// replace tab with 4 spaces
		line = strings.Replace(line, "\t", "    ", -1)
		translateLines = append(translateLines, line)
	}
	log.Println("Translating ...")
	err = Translate(strings.Join(translateLines, "\n"), "en", "zh-CN")
	check(err)

	s := outString
	// fmt.Println("s:", s)

	// Fix the outString
	// fix title in translatedText
	s = strings.Replace(s, "＃＃＃＃", "#### ", -1)
	s = strings.Replace(s, "＃＃＃", "### ", -1)
	s = strings.Replace(s, "＃＃", "## ", -1)
	s = strings.Replace(s, "＃", "# ", -1)
	s = regexp.MustCompile(`(#+)(\p{Han}+)`).ReplaceAllString(s, "$1 $2")
	// fix Zero width space Unicode
	s = strings.Replace(s, "\u200B", "", -1)
	// fix the Chinese char
	s = strings.Replace(s, "【", " [", -1)
	s = strings.Replace(s, "】", "]", -1)
	s = strings.Replace(s, `\。`, ". ", -1)
	s = regexp.MustCompile(`]（(.*)）`).ReplaceAllString(s, "]($1)")
	s = regexp.MustCompile(`]\((.*)）`).ReplaceAllString(s, "]($1)")
	s = regexp.MustCompile(`]（(.*)\)`).ReplaceAllString(s, "]($1)")
	// fix url with space
	s = regexp.MustCompile(`]\(http([^\s]*?) ([^\s]*?)\)`).ReplaceAllString(s, "](http$1$2)")
	// Remove mutiple empty lines with one empty line.
	s = regexp.MustCompile(`\n{3,}`).ReplaceAllString(s, "\n\n")

	// a := "＃＃"
	// a = strings.Replace(a, "＃＃", "##1 ", -1)
	// fmt.Println(a)
	err = os.WriteFile("tr-"+*inputFile, []byte(s), 0644)
	check(err)
	log.Println("Done. Generated output file: ", "tr-"+*inputFile)
}
