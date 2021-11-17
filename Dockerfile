# *************************************
#
# OpenGM
#
# *************************************

FROM alpine:3.14

MAINTAINER XTech Cloud "xtech.cloud"

ENV container docker
ENV MSA_MODE release

EXPOSE 18801

ADD bin/ogm-account /usr/local/bin/
RUN chmod +x /usr/local/bin/ogm-account

CMD ["/usr/local/bin/ogm-account"]
