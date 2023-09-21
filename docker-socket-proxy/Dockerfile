# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"a968cf52678de22cea08531444603b3d6e7119c1cbd044d24aea721b9742b0bd","image:haproxy":"8a6d253ac6b691da9fa82e7a32d3e9a672df9617eeabf5daa2b160d130375123"}

FROM reg.c5h.io/haproxy@sha256:8a6d253ac6b691da9fa82e7a32d3e9a672df9617eeabf5daa2b160d130375123 AS haproxy

FROM reg.c5h.io/base-root@sha256:a968cf52678de22cea08531444603b3d6e7119c1cbd044d24aea721b9742b0bd

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
