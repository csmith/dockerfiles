# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"64873d5cd1e5603150a658574295b3b617f2213f97e03caf9530728e92006780","image:haproxy":"63c9932c6bafad6c7152225a8db9369afb37a75c3ad2fb3e175bdd7e4356c429"}

FROM reg.c5h.io/haproxy@sha256:63c9932c6bafad6c7152225a8db9369afb37a75c3ad2fb3e175bdd7e4356c429 AS haproxy

FROM reg.c5h.io/base-root@sha256:64873d5cd1e5603150a658574295b3b617f2213f97e03caf9530728e92006780

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
