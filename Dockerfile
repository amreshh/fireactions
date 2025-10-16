FROM --platform=$TARGETPLATFORM debian:stable-slim

COPY fireactions /usr/bin/fireactions
ENV NODE_EXTRA_CA_CERTS=/usr/local/share/ca-certificates/cert.crt

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates                                              \
    && apt-get autoremove -y                                     \
    && apt-get clean                                             \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*             \
    && groupadd -g 1000 fireactions                              \
    && useradd -u 1000 -g fireactions -s /bin/sh -m fireactions  \
    && chown fireactions:fireactions /usr/bin/fireactions        \
    && chmod 755 /usr/bin/fireactions

ADD cert.crt /usr/local/share/ca-certificates
RUN update-ca-certificates

EXPOSE 8080

COPY entrypoint.sh /usr/bin/entrypoint.sh

ENTRYPOINT ["/usr/bin/entrypoint.sh"]
