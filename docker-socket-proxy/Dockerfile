# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"b1cab13debf1ce2650313df4609a21b8b5ad748cc8c5a83424b290db8a7cfe4a","image:haproxy":"2a8c7cd9ea83bed244cb979cc31a6334437e693dfe505e630d5b0a11fbeb8a96"}

FROM reg.c5h.io/haproxy@sha256:2a8c7cd9ea83bed244cb979cc31a6334437e693dfe505e630d5b0a11fbeb8a96 AS haproxy

FROM reg.c5h.io/base-root@sha256:b1cab13debf1ce2650313df4609a21b8b5ad748cc8c5a83424b290db8a7cfe4a

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
