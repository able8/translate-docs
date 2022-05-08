# generate markdown file from blog URLs
get:
	pbpaste | sed 's#/$$##g' | grep http | while read line; do html2md -i $${line} --opt-code-block-style fenced > $${line##*/}.md; echo "\n\n$$line" >> $${line##*/}.md ;done

translate:
	git status | grep ".md" | grep -v README.md | grep -v "tr-"  | while read line; do go run ../translate.go -f $$line; done

# update README.md
update:
	git status  | grep "tr-" | while read line; do echo "1. [$$(head -n 1 $$line|sed 's/# //g')](docs/$$line)";done | pbcopy;pbpaste >> ../README.md