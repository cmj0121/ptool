FROM alpine:3.12.3
MAINTAINER cmj <cmj@cmj.tw>

# install necessary library
RUN apk add --no-cache \
	bash \
	openssh

COPY entrypoint /root/entrypoint
COPY sshd_config /etc/ssh/sshd_config
RUN chmod +x /root/entrypoint

CMD /root/entrypoint
