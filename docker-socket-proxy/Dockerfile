# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"3be0a174bb370de9aaec96dc1675771e336016e17c8c52ef0059badb6fb2573b","image:haproxy":"eb79c0b8915fa57ac21be961c870f9637b96fe67e893e9a785742cf99daa21f3"}

FROM reg.c5h.io/haproxy@sha256:eb79c0b8915fa57ac21be961c870f9637b96fe67e893e9a785742cf99daa21f3 AS haproxy

FROM reg.c5h.io/base-root@sha256:3be0a174bb370de9aaec96dc1675771e336016e17c8c52ef0059badb6fb2573b

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
