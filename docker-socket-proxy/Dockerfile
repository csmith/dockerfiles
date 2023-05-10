# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"ede8aec15939962252ae6d3fa777cb26e6d5011ed6c386961ea6b2ab92439946","image:haproxy":"602bc882e96e100dfd8afce80fb1ecda823934845b7862b42507984210fb973a"}

FROM reg.c5h.io/haproxy@sha256:602bc882e96e100dfd8afce80fb1ecda823934845b7862b42507984210fb973a AS haproxy

FROM reg.c5h.io/base-root@sha256:ede8aec15939962252ae6d3fa777cb26e6d5011ed6c386961ea6b2ab92439946

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
