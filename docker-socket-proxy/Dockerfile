# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"cb2f777a8f1be69c385424a3748112aa34e562abfb51ea4ff74a616071660d14","image:haproxy":"6accf5eb8dbc296407064e0c3d633b0e48af12136d2db9bdd35602c6603b2a90"}

FROM reg.c5h.io/haproxy@sha256:6accf5eb8dbc296407064e0c3d633b0e48af12136d2db9bdd35602c6603b2a90 AS haproxy

FROM reg.c5h.io/base-root@sha256:cb2f777a8f1be69c385424a3748112aa34e562abfb51ea4ff74a616071660d14

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
