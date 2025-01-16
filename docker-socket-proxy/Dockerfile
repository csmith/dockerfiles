# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"3be0a174bb370de9aaec96dc1675771e336016e17c8c52ef0059badb6fb2573b","image:haproxy":"e4cab75eca9fee1ab3c4d2e62223ffae7013d1b3d1ca516dad94f814104a9f0f"}

FROM reg.c5h.io/haproxy@sha256:e4cab75eca9fee1ab3c4d2e62223ffae7013d1b3d1ca516dad94f814104a9f0f AS haproxy

FROM reg.c5h.io/base-root@sha256:3be0a174bb370de9aaec96dc1675771e336016e17c8c52ef0059badb6fb2573b

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
