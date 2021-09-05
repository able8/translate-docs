## help: print this help message
help:
	@echo "Usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ":" | sed -e 's/^/  /'

## get: get URLs and generate Markdown file
.PHONY: get
get: 
	pbpaste | sed 's#/$$##g' | sed 's/.html//g' | grep http | while read line; do \
	html2md -i $${line} --opt-code-block-style fenced > $${line##*/}.md;done

## tr: translate Markdown file
.PHONY: tr
tr: 
	git status | grep ".md" | grep -v "tr-" | while read line; do  \
	echo $${line};go run ../translate.go -f "$${line}"; done

## add: add the tanslated file name and path to README
.PHONY: add
add: 
	git status | grep "tr-" | while read line; do \
	echo "1. [$(head -n 1 $${line}|sed 's/# //g')](docs/$${line})"\
	done | pbcopy;pbpaste;pbpaste >> ../README.md
