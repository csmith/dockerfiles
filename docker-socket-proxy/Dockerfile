# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"4c77ca91e924927dc19fdbbb6cb4ca76f3471ce834b97becb2a4757f1bea0092","image:haproxy":"3c8a87b1d6baaa233e1c5aac89ddadeb45255339f91824a6080a6e269b88f794"}

FROM reg.c5h.io/haproxy@sha256:3c8a87b1d6baaa233e1c5aac89ddadeb45255339f91824a6080a6e269b88f794 AS haproxy

FROM reg.c5h.io/base-root@sha256:4c77ca91e924927dc19fdbbb6cb4ca76f3471ce834b97becb2a4757f1bea0092

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
