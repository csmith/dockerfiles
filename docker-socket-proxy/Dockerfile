# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"64873d5cd1e5603150a658574295b3b617f2213f97e03caf9530728e92006780","image:haproxy":"bf9da08c9bd4db0b873e58e08fdba48e843e6551f14df1167f7ab3c297d5fd93"}

FROM reg.c5h.io/haproxy@sha256:bf9da08c9bd4db0b873e58e08fdba48e843e6551f14df1167f7ab3c297d5fd93 AS haproxy

FROM reg.c5h.io/base-root@sha256:64873d5cd1e5603150a658574295b3b617f2213f97e03caf9530728e92006780

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
