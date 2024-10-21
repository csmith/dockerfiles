# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"fe9cbe2b7a96e96c0b217e2c21a8855bafa38d016bc59557cc5485da2847e311","image:haproxy":"20ea301282c2c731b11780d7ace96e272e26b24fb4640a796a76ee8213978f24"}

FROM reg.c5h.io/haproxy@sha256:20ea301282c2c731b11780d7ace96e272e26b24fb4640a796a76ee8213978f24 AS haproxy

FROM reg.c5h.io/base-root@sha256:fe9cbe2b7a96e96c0b217e2c21a8855bafa38d016bc59557cc5485da2847e311

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
