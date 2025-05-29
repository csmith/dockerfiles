# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"46984a720077bee7e09f32cceb5d1d6041cd524e4b7329a557abc293f79c3429","image:haproxy":"d7af7efa842327a925e3360284b30b8bfeeba24c7e26bfe68b4af356b2904598"}

FROM reg.c5h.io/haproxy@sha256:d7af7efa842327a925e3360284b30b8bfeeba24c7e26bfe68b4af356b2904598 AS haproxy

FROM reg.c5h.io/base-root@sha256:46984a720077bee7e09f32cceb5d1d6041cd524e4b7329a557abc293f79c3429

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
