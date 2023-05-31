# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"ede8aec15939962252ae6d3fa777cb26e6d5011ed6c386961ea6b2ab92439946","image:haproxy":"1d9ccd529ae8ae12631845601934036c0ab459a048fb29e5f4802976c6f3fc03"}

FROM reg.c5h.io/haproxy@sha256:1d9ccd529ae8ae12631845601934036c0ab459a048fb29e5f4802976c6f3fc03 AS haproxy

FROM reg.c5h.io/base-root@sha256:ede8aec15939962252ae6d3fa777cb26e6d5011ed6c386961ea6b2ab92439946

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
