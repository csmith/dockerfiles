# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"14dbdc5aa6a2b2edb6f61e3fc29fc0dbe0770a72233d13b160eae3afc89fec8f","image:haproxy":"e919efc311fac91fac57c8c59396c0ee9b3bc3412548399dd2193ecf35010d38"}

FROM reg.c5h.io/haproxy@sha256:e919efc311fac91fac57c8c59396c0ee9b3bc3412548399dd2193ecf35010d38 AS haproxy

FROM reg.c5h.io/base-root@sha256:14dbdc5aa6a2b2edb6f61e3fc29fc0dbe0770a72233d13b160eae3afc89fec8f

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
