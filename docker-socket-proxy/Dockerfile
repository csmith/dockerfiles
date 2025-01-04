# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"0e0f9efe0a344277681c12a1935a4f1efd5c7312a76b0b66c8efffd2235fffdc","image:haproxy":"52c64a2e30c4dcfd20afe2042d8d199457872aa9e55fccadc727f672c12d95f2"}

FROM reg.c5h.io/haproxy@sha256:52c64a2e30c4dcfd20afe2042d8d199457872aa9e55fccadc727f672c12d95f2 AS haproxy

FROM reg.c5h.io/base-root@sha256:0e0f9efe0a344277681c12a1935a4f1efd5c7312a76b0b66c8efffd2235fffdc

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
