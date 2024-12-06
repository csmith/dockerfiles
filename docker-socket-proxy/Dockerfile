# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"887a061c4ea8654f87847eb0c31ed58616ec3eb5670cb48dffec7995e81c5449","image:haproxy":"92ebe8a52e94b0e3d9583565e39238c29a071505bd546e375d7301654e32e143"}

FROM reg.c5h.io/haproxy@sha256:92ebe8a52e94b0e3d9583565e39238c29a071505bd546e375d7301654e32e143 AS haproxy

FROM reg.c5h.io/base-root@sha256:887a061c4ea8654f87847eb0c31ed58616ec3eb5670cb48dffec7995e81c5449

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
