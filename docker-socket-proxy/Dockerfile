# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"28fed4c8021ee4ebf01c6e3ad048388acde4d60f951511dff65ebf6bd25aec28","image:haproxy":"cc607fe195b30eb0665e8f345e1a0ed37238e7308349dd8af7c79151f7c9f949"}

FROM reg.c5h.io/haproxy@sha256:cc607fe195b30eb0665e8f345e1a0ed37238e7308349dd8af7c79151f7c9f949 AS haproxy

FROM reg.c5h.io/base-root@sha256:28fed4c8021ee4ebf01c6e3ad048388acde4d60f951511dff65ebf6bd25aec28

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
