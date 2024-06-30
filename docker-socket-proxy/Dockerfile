# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"71acce7ad1826b2a322287e1f9a49a1832062ef056a55f4a484fcc4a711803ad","image:haproxy":"7e4d26a66340c77f412e59088d260a483bdf7a937d82dd4350da692a10c79f6a"}

FROM reg.c5h.io/haproxy@sha256:7e4d26a66340c77f412e59088d260a483bdf7a937d82dd4350da692a10c79f6a AS haproxy

FROM reg.c5h.io/base-root@sha256:71acce7ad1826b2a322287e1f9a49a1832062ef056a55f4a484fcc4a711803ad

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
