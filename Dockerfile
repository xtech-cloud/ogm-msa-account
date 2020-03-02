FROM alpine:3.11
ADD ./bin/omo-msa-account /usr/bin/omo-msa-account
ENV MSA_REGISTRY_PLUGIN
ENV MSA_REGISTRY_ADDRESS
ENV MSA_CONFIG_DEFINE
EXPOSE 8080
ENTRYPOINT [ "omo-msa-account" ]
