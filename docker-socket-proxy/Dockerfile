# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"019bd3cd845b6b093fd35645aa2e8266834428f2caee98027b425d2b56b7c8e1","image:haproxy":"a17915226d26fc15b4221a41ae9e6cb8af558a8d6f0a2c4da2c27e42d882ffce"}

FROM reg.c5h.io/haproxy@sha256:a17915226d26fc15b4221a41ae9e6cb8af558a8d6f0a2c4da2c27e42d882ffce AS haproxy

FROM reg.c5h.io/base-root@sha256:019bd3cd845b6b093fd35645aa2e8266834428f2caee98027b425d2b56b7c8e1

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
