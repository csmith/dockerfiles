# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"5100f1fc0d5d449eb7aee39359429531f3b0806c013ce5e09e58f73a74b8ffa1","image:haproxy":"95bd44b194b8cbeb378cce93b2282dfc73c61ad0186299199e85ce669e5eea09"}

FROM reg.c5h.io/haproxy@sha256:95bd44b194b8cbeb378cce93b2282dfc73c61ad0186299199e85ce669e5eea09 AS haproxy

FROM reg.c5h.io/base-root@sha256:5100f1fc0d5d449eb7aee39359429531f3b0806c013ce5e09e58f73a74b8ffa1

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
