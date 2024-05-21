# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"9b572afcdbac62e8aedcb7a8d43429567fa9246b974f2f1a4e6954fe3d54e72b","image:haproxy":"ea9d2c25069a70304a3762ca29093f4d6870c98a16066dbd8bca7173096b1126"}

FROM reg.c5h.io/haproxy@sha256:ea9d2c25069a70304a3762ca29093f4d6870c98a16066dbd8bca7173096b1126 AS haproxy

FROM reg.c5h.io/base-root@sha256:9b572afcdbac62e8aedcb7a8d43429567fa9246b974f2f1a4e6954fe3d54e72b

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
