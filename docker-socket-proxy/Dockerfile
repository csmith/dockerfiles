# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"5dea99b6fc9ffe51ff467e5ae1f2ccd6157090f4d20db63d306faa7fe7adac51","image:haproxy":"29f45685162ec17500857640ad6462e5280fd0d86fa20f3f3f6c4df1c7ac9c61"}

FROM reg.c5h.io/haproxy@sha256:29f45685162ec17500857640ad6462e5280fd0d86fa20f3f3f6c4df1c7ac9c61 AS haproxy

FROM reg.c5h.io/base-root@sha256:5dea99b6fc9ffe51ff467e5ae1f2ccd6157090f4d20db63d306faa7fe7adac51

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
