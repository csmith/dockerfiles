# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"d17ade001fee6e28dc9a4d3a31b0aa576d04a5cda159cf955f4ebd1376513730","image:haproxy":"43420659033974277710a154ee591d93a0f222923b4ab22a32ac70b4f462433e"}

FROM reg.c5h.io/haproxy@sha256:43420659033974277710a154ee591d93a0f222923b4ab22a32ac70b4f462433e AS haproxy

FROM reg.c5h.io/base-root@sha256:d17ade001fee6e28dc9a4d3a31b0aa576d04a5cda159cf955f4ebd1376513730

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
