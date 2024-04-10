# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"9b572afcdbac62e8aedcb7a8d43429567fa9246b974f2f1a4e6954fe3d54e72b","image:haproxy":"872b31cdc295d1db24b8ebb13a54dcbf931b33f670da3bd7ba35e33489daa73a"}

FROM reg.c5h.io/haproxy@sha256:872b31cdc295d1db24b8ebb13a54dcbf931b33f670da3bd7ba35e33489daa73a AS haproxy

FROM reg.c5h.io/base-root@sha256:9b572afcdbac62e8aedcb7a8d43429567fa9246b974f2f1a4e6954fe3d54e72b

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
