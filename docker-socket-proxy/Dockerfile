# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"3be0a174bb370de9aaec96dc1675771e336016e17c8c52ef0059badb6fb2573b","image:haproxy":"cb1889684eca19b6bc96c266909f71e5d811aeabac41aee8de8fece44e890dfb"}

FROM reg.c5h.io/haproxy@sha256:cb1889684eca19b6bc96c266909f71e5d811aeabac41aee8de8fece44e890dfb AS haproxy

FROM reg.c5h.io/base-root@sha256:3be0a174bb370de9aaec96dc1675771e336016e17c8c52ef0059badb6fb2573b

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
