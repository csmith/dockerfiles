# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"ede8aec15939962252ae6d3fa777cb26e6d5011ed6c386961ea6b2ab92439946","image:haproxy":"fcff917171810e5597e18aa137cb4392a796587419269c1e8f2c47697543f4ef"}

FROM reg.c5h.io/haproxy@sha256:fcff917171810e5597e18aa137cb4392a796587419269c1e8f2c47697543f4ef AS haproxy

FROM reg.c5h.io/base-root@sha256:ede8aec15939962252ae6d3fa777cb26e6d5011ed6c386961ea6b2ab92439946

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
