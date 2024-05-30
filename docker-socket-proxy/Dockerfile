# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"71acce7ad1826b2a322287e1f9a49a1832062ef056a55f4a484fcc4a711803ad","image:haproxy":"ca3f3160c197d7c1c09e477d51a588b2560c007f188d5836734c1d4575f14136"}

FROM reg.c5h.io/haproxy@sha256:ca3f3160c197d7c1c09e477d51a588b2560c007f188d5836734c1d4575f14136 AS haproxy

FROM reg.c5h.io/base-root@sha256:71acce7ad1826b2a322287e1f9a49a1832062ef056a55f4a484fcc4a711803ad

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
