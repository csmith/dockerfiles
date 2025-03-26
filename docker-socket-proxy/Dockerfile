# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"46984a720077bee7e09f32cceb5d1d6041cd524e4b7329a557abc293f79c3429","image:haproxy":"23a78ed5a3dc5808528ed40fd05c071c2ced808587604daa703c4d52d4e3c45f"}

FROM reg.c5h.io/haproxy@sha256:23a78ed5a3dc5808528ed40fd05c071c2ced808587604daa703c4d52d4e3c45f AS haproxy

FROM reg.c5h.io/base-root@sha256:46984a720077bee7e09f32cceb5d1d6041cd524e4b7329a557abc293f79c3429

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
