# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"ede8aec15939962252ae6d3fa777cb26e6d5011ed6c386961ea6b2ab92439946","image:haproxy":"6b8c4ec5f0578b08cbe205229ed7a309ead6fd33887dd74fb2802f534ac1c25b"}

FROM reg.c5h.io/haproxy@sha256:6b8c4ec5f0578b08cbe205229ed7a309ead6fd33887dd74fb2802f534ac1c25b AS haproxy

FROM reg.c5h.io/base-root@sha256:ede8aec15939962252ae6d3fa777cb26e6d5011ed6c386961ea6b2ab92439946

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
