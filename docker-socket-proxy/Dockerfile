# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"d17ade001fee6e28dc9a4d3a31b0aa576d04a5cda159cf955f4ebd1376513730","image:haproxy":"8b73ca2c0c1fe6ed354c91058e0b15b6fde4e50d635377dcefe5e42a923f6a33"}

FROM reg.c5h.io/haproxy@sha256:8b73ca2c0c1fe6ed354c91058e0b15b6fde4e50d635377dcefe5e42a923f6a33 AS haproxy

FROM reg.c5h.io/base-root@sha256:d17ade001fee6e28dc9a4d3a31b0aa576d04a5cda159cf955f4ebd1376513730

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
