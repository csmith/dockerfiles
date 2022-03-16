# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"4c77ca91e924927dc19fdbbb6cb4ca76f3471ce834b97becb2a4757f1bea0092","image:haproxy":"75b44ff1030af1d7d1f52c999781fb8791ee6801730be43c9134616c1d30daae"}

FROM reg.c5h.io/haproxy@sha256:75b44ff1030af1d7d1f52c999781fb8791ee6801730be43c9134616c1d30daae AS haproxy

FROM reg.c5h.io/base-root@sha256:4c77ca91e924927dc19fdbbb6cb4ca76f3471ce834b97becb2a4757f1bea0092

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
