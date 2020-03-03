FROM alpine:3.11
RUN wget -O /usr/bin/omo-msa-account https://github.com/xtech-cloud/omo-msa-account/releases/download/v1.0.0/omo-msa-account
RUN chmod +x /usr/bin/omo-msa-account
ENV MSA_REGISTRY_PLUGIN
ENV MSA_REGISTRY_ADDRESS
ENV MSA_CONFIG_DEFINE
EXPOSE 8080
ENTRYPOINT [ "omo-msa-account" ]
