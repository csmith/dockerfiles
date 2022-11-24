# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"1c975ebed03f0278f2a9aecd9637ea97b7d5b0efb5543da9f1e56c83b84450f0","image:haproxy":"70be32160be0f99c363ea72ba94490f0cf9706a046406e0b1743ef1afc583ec0"}

FROM reg.c5h.io/haproxy@sha256:70be32160be0f99c363ea72ba94490f0cf9706a046406e0b1743ef1afc583ec0 AS haproxy

FROM reg.c5h.io/base-root@sha256:1c975ebed03f0278f2a9aecd9637ea97b7d5b0efb5543da9f1e56c83b84450f0

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
