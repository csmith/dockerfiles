# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"64b1a284c9851528f274df911604a3f42b251b4272a3160fb21e4f8b0b912aa7","image:haproxy":"5a3fad57e7ecbe084a30dc10acba059feb0fababb87bb379da8201867b814371"}

FROM reg.c5h.io/haproxy@sha256:5a3fad57e7ecbe084a30dc10acba059feb0fababb87bb379da8201867b814371 AS haproxy

FROM reg.c5h.io/base-root@sha256:64b1a284c9851528f274df911604a3f42b251b4272a3160fb21e4f8b0b912aa7

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
