# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"3be0a174bb370de9aaec96dc1675771e336016e17c8c52ef0059badb6fb2573b","image:haproxy":"798a599cf3a51f94fac984b868357bd4155b9103ba57636e2cdbc45e0449813e"}

FROM reg.c5h.io/haproxy@sha256:798a599cf3a51f94fac984b868357bd4155b9103ba57636e2cdbc45e0449813e AS haproxy

FROM reg.c5h.io/base-root@sha256:3be0a174bb370de9aaec96dc1675771e336016e17c8c52ef0059badb6fb2573b

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
