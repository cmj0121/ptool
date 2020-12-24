FROM alpine:3.12.3
MAINTAINER cmj <cmj@cmj.tw>

# install necessary library
RUN apk add --no-cache \
	bash \
	openssh

VOLUME /home

COPY entrypoint /root/entrypoint
RUN chmod +x /root/entrypoint

CMD /root/entrypoint
