# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"ede8aec15939962252ae6d3fa777cb26e6d5011ed6c386961ea6b2ab92439946","image:haproxy":"b74ccae6563ce6a6d400c1f2d9654e2960b567478e7e6ac40d53fd363ce88a7d"}

FROM reg.c5h.io/haproxy@sha256:b74ccae6563ce6a6d400c1f2d9654e2960b567478e7e6ac40d53fd363ce88a7d AS haproxy

FROM reg.c5h.io/base-root@sha256:ede8aec15939962252ae6d3fa777cb26e6d5011ed6c386961ea6b2ab92439946

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
