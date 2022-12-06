# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"1c975ebed03f0278f2a9aecd9637ea97b7d5b0efb5543da9f1e56c83b84450f0","image:haproxy":"d0702b4b95137d7e95cd1815822df1f7390eb84c08247dde46fc20d916a9c8d1"}

FROM reg.c5h.io/haproxy@sha256:d0702b4b95137d7e95cd1815822df1f7390eb84c08247dde46fc20d916a9c8d1 AS haproxy

FROM reg.c5h.io/base-root@sha256:1c975ebed03f0278f2a9aecd9637ea97b7d5b0efb5543da9f1e56c83b84450f0

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
