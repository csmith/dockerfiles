# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"8ccee4a219c0e0c7b135814c4c423efd774ba75103482c83e5bc9c1ed9f127d8","image:haproxy":"57544914c688ca5f42ff5db28d19aa0ff51a35d3c87eb170e46ca6f793f97410"}

FROM reg.c5h.io/haproxy@sha256:57544914c688ca5f42ff5db28d19aa0ff51a35d3c87eb170e46ca6f793f97410 AS haproxy

FROM reg.c5h.io/base-root@sha256:8ccee4a219c0e0c7b135814c4c423efd774ba75103482c83e5bc9c1ed9f127d8

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
