# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"14dbdc5aa6a2b2edb6f61e3fc29fc0dbe0770a72233d13b160eae3afc89fec8f","image:haproxy":"56cddf94544321623bf293c4cc393a3d0b302b264262242a2d9d897d2cb248bc"}

FROM reg.c5h.io/haproxy@sha256:56cddf94544321623bf293c4cc393a3d0b302b264262242a2d9d897d2cb248bc AS haproxy

FROM reg.c5h.io/base-root@sha256:14dbdc5aa6a2b2edb6f61e3fc29fc0dbe0770a72233d13b160eae3afc89fec8f

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
