# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"28fed4c8021ee4ebf01c6e3ad048388acde4d60f951511dff65ebf6bd25aec28","image:haproxy":"c64132b6d22ea12726ec1833ce9a4462e299cd91f431c2307ff435ef580b4fe6"}

FROM reg.c5h.io/haproxy@sha256:c64132b6d22ea12726ec1833ce9a4462e299cd91f431c2307ff435ef580b4fe6 AS haproxy

FROM reg.c5h.io/base-root@sha256:28fed4c8021ee4ebf01c6e3ad048388acde4d60f951511dff65ebf6bd25aec28

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
