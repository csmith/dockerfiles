# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"sha256:ff79e1382535b628dffa6eff867f4d22c7c217eaabe875dcb49033a87bd8d8da","image:haproxy":"sha256:3776d3b6921eac453b1b3ea65bd539b38e5e985cdff93035cdd449e3d8676ae0"}

FROM reg.c5h.io/haproxy@sha256:3776d3b6921eac453b1b3ea65bd539b38e5e985cdff93035cdd449e3d8676ae0 AS haproxy

FROM reg.c5h.io/base-root@sha256:ff79e1382535b628dffa6eff867f4d22c7c217eaabe875dcb49033a87bd8d8da

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
