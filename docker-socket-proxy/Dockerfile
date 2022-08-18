# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"58bb78426f91647d2336b955555c415ba8e06d603a9ac20fe45ebf049c4a3739","image:haproxy":"c9377f885d2f92c216690506b4ad6cb7ed5e75cb5a899d9221e92b97fd92cbd3"}

FROM reg.c5h.io/haproxy@sha256:c9377f885d2f92c216690506b4ad6cb7ed5e75cb5a899d9221e92b97fd92cbd3 AS haproxy

FROM reg.c5h.io/base-root@sha256:58bb78426f91647d2336b955555c415ba8e06d603a9ac20fe45ebf049c4a3739

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
