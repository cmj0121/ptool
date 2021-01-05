.PHONY: all clean help

VERSION := last

all:

clean:

help:	# show this message
	@printf "Usage: make [OPTION]\n"
	@printf "\n"
	@perl -nle 'print $$& if m{^[\w-]+:.*?#.*$$}' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?#"} {printf "    %-18s %s\n", $$1, $$2}'

flask:		# run the Python/flask in the debug mode
	docker build -t lab-flask --no-cache -f flask.Dockerfile .
	docker run --name lab-flask --rm -d  -p 5000:5000 lab-flask

flask-stop:	# stop the flask lab
	docker stop lab-flask

flask-log:	# show the log in the flask lab
	docker logs -f lab-flask
