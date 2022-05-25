# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"fae0414d67593373dbb0cba1d6e3b740837e3b8873f44fd91d5639ab3acdf952","image:haproxy":"3a6cd0d5f8f8ab5bd34b537cd2d60a68ba5a128e3447db2feb24595acbf34db9"}

FROM reg.c5h.io/haproxy@sha256:3a6cd0d5f8f8ab5bd34b537cd2d60a68ba5a128e3447db2feb24595acbf34db9 AS haproxy

FROM reg.c5h.io/base-root@sha256:fae0414d67593373dbb0cba1d6e3b740837e3b8873f44fd91d5639ab3acdf952

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
