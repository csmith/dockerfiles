# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"fae0414d67593373dbb0cba1d6e3b740837e3b8873f44fd91d5639ab3acdf952","image:haproxy":"91603a140edbce11caaf5a8175813764856e647ed25fda72249ecbe527d0ba83"}

FROM reg.c5h.io/haproxy@sha256:91603a140edbce11caaf5a8175813764856e647ed25fda72249ecbe527d0ba83 AS haproxy

FROM reg.c5h.io/base-root@sha256:fae0414d67593373dbb0cba1d6e3b740837e3b8873f44fd91d5639ab3acdf952

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
