.PHONY: all clean help

IMAGE_NAME := ssh-server

all:		# build all

clean:		# clean-up the environment

help:		# show this message
	@printf "Usage: make [OPTION]\n"
	@printf "\n"
	@perl -nle 'print $$& if m{^[\w-]+:.*?#.*$$}' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?#"} {printf "    %-18s %s\n", $$1, $$2}'

build:	# build the docker image
	docker build -t $(IMAGE_NAME) .

run:	# run as the docker in interactive / pseudo-TTY mode
	docker run \
		-d \
		--rm -it \
		--name ssh-server \
		-e USERNAME=pentest \
		-e PASSWORD=password \
		-p 2222:22 \
		-v $(PWD)/home:/home/pentest \
		ssh-server

stop:	# stop the running image
	docker stop ssh-server
