# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"14dbdc5aa6a2b2edb6f61e3fc29fc0dbe0770a72233d13b160eae3afc89fec8f","image:haproxy":"9120086ca724405e7367fb1173d179db7a12318a52c957cc2deabfb1c997d2e4"}

FROM reg.c5h.io/haproxy@sha256:9120086ca724405e7367fb1173d179db7a12318a52c957cc2deabfb1c997d2e4 AS haproxy

FROM reg.c5h.io/base-root@sha256:14dbdc5aa6a2b2edb6f61e3fc29fc0dbe0770a72233d13b160eae3afc89fec8f

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
