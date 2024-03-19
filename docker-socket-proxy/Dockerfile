# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"9b572afcdbac62e8aedcb7a8d43429567fa9246b974f2f1a4e6954fe3d54e72b","image:haproxy":"f464cc33db9cc48dbc422beabd0a6dade342f71aa40db0cdf592f673bc11ea10"}

FROM reg.c5h.io/haproxy@sha256:f464cc33db9cc48dbc422beabd0a6dade342f71aa40db0cdf592f673bc11ea10 AS haproxy

FROM reg.c5h.io/base-root@sha256:9b572afcdbac62e8aedcb7a8d43429567fa9246b974f2f1a4e6954fe3d54e72b

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
