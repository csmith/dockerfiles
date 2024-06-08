# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"71acce7ad1826b2a322287e1f9a49a1832062ef056a55f4a484fcc4a711803ad","image:haproxy":"6c31180a681ef60e99ca3f0387dc661f16f0da08e90b5f9214208c35fda13bd0"}

FROM reg.c5h.io/haproxy@sha256:6c31180a681ef60e99ca3f0387dc661f16f0da08e90b5f9214208c35fda13bd0 AS haproxy

FROM reg.c5h.io/base-root@sha256:71acce7ad1826b2a322287e1f9a49a1832062ef056a55f4a484fcc4a711803ad

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
