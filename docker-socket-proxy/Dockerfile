# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"6d131c89e231e9a653da086b59d83f269f52924449a9ab8e7cbd6e86fe35d74f","image:haproxy":"34be52faca4da0bd7699d654f659db6717cee407a53d1c6549f66ff4501075c4"}

FROM reg.c5h.io/haproxy@sha256:34be52faca4da0bd7699d654f659db6717cee407a53d1c6549f66ff4501075c4 AS haproxy

FROM reg.c5h.io/base-root@sha256:6d131c89e231e9a653da086b59d83f269f52924449a9ab8e7cbd6e86fe35d74f

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
