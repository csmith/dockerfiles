# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"3e2f04b331f8e52c1d703bd48fec73da1f9936fc03c027b876b12da9b10b1140","image:haproxy":"046c755887b039723921b63fb02c9c630f997bd8af4af9bbafae61f1f0ec3ea1"}

FROM reg.c5h.io/haproxy@sha256:046c755887b039723921b63fb02c9c630f997bd8af4af9bbafae61f1f0ec3ea1 AS haproxy

FROM reg.c5h.io/base-root@sha256:3e2f04b331f8e52c1d703bd48fec73da1f9936fc03c027b876b12da9b10b1140

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
