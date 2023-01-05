# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"81ca1c32f550edd1ab24ccdecc9974a823385327c0e918757d6d3d96f04990fc","image:haproxy":"c32fea7ff91ed222e6b3651dd73c69b158e45f12591d0028c241c2bfd23d56d4"}

FROM reg.c5h.io/haproxy@sha256:c32fea7ff91ed222e6b3651dd73c69b158e45f12591d0028c241c2bfd23d56d4 AS haproxy

FROM reg.c5h.io/base-root@sha256:81ca1c32f550edd1ab24ccdecc9974a823385327c0e918757d6d3d96f04990fc

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
