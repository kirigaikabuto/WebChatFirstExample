all: one
.PHONY: all

one:
	git add .
	git commit -m "commits"
	git push
	go build
	./tik-api