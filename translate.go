package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"
)

var outFile *os.File

// Use google translate api
func Translate(source, sourceLang, targetLang string, outFile *os.File) error {
	queryURL := "https://translate.google.cn/translate_a/single?client=gtx&sl=" +
		sourceLang + "&tl=" + targetLang + "&dt=t&q=" + url.QueryEscape(source)
	// fmt.Println(queryURL)
	var resp *http.Response
	var err error
	for try := 0; try < 3; try++ {
		resp, err = http.Get(queryURL)
		if resp.StatusCode == http.StatusOK {
			break
		}
	}
	if err != nil {
		return fmt.Errorf("failed to get translation result from translate.google.cn, err: %s", resp.Status)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("failed to read response body")
	}

	// fmt.Println(string(body))
	var result []interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return errors.New("Error unmarshaling data")
	}

	var input, output string
	if len(result) > 0 {
		for _, slice := range result[0].([]interface{}) {
			resultSlice := slice.([]interface{})
			translatedText, sourceText := resultSlice[0].(string), resultSlice[1].(string)

			// Do not translate these lines
			if strings.Contains(sourceText, "![") || strings.Contains(sourceText, "----") || strings.Contains(sourceText, "From: http") {
				input = input + sourceText
				continue
			}

			if strings.Contains(sourceText, "```") && strings.Contains(input, "```") {
				input = input + sourceText
				fmt.Fprintf(outFile, "%s\n", input)
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

			if strings.Contains(sourceText, "\n\n") && !strings.Contains(sourceText, ":\n") {

				fmt.Fprintf(outFile, "%s%s", input, output)
				input, output = "", ""
			}
		}

		fmt.Fprintf(outFile, "%s\n\n%s\n\n", input, output)
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

	outFile, err := os.Create("tr-" + *inputFile)
	check(err)
	defer outFile.Close()

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
			err = Translate(strings.Join(translateLines, "\n"), "en", "zh-CN", outFile)
			check(err)
			time.Sleep(1 * time.Second)
			count = len(line)
			translateLines = nil
		}
		// replace tab with 4 spaces
		line = strings.Replace(line, "\t", "    ", -1)
		translateLines = append(translateLines, line)
	}
	log.Println("Translating ...")
	err = Translate(strings.Join(translateLines, "\n"), "en", "zh-CN", outFile)
	check(err)

	data, err := ioutil.ReadAll(outFile)
	check(err)
	s := string(data)

	// fix title in translatedText
	s = regexp.MustCompile(`(#+)(\p{Han}+)`).ReplaceAllString(s, "$1 $2")
	s = strings.Replace(s, "＃＃＃＃", "#### ", -1)
	s = strings.Replace(s, "＃＃＃", "### ", -1)
	s = strings.Replace(s, "＃＃", "## ", -1)
	s = strings.Replace(s, "＃", "# ", -1)
	// fix Zero width space Unicode
	s = strings.Replace(s, "\u200B", "", -1)
	// fix the Chinese char
	s = strings.Replace(s, "【", " [", -1)
	s = strings.Replace(s, "】", "]", -1)
	s = regexp.MustCompile(`]（(.*)）`).ReplaceAllString(s, "]($1)")
	s = regexp.MustCompile(`]\((.*)）`).ReplaceAllString(s, "]($1)")
	s = regexp.MustCompile(`]（(.*)\)`).ReplaceAllString(s, "]($1)")
	// fix url with space
	s = regexp.MustCompile(`]\(http([^\s]*?) ([^\s]*?)\)`).ReplaceAllString(s, "](http$1$2)")

	fmt.Fprintf(outFile, s)
	log.Println("Done. Generated output file: ", "tr-"+*inputFile)
}
