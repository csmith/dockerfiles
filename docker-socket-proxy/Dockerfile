# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"1cfdcfc1a5bc7a5923c59bbf4be055cdea0f63bb28bdb8cffaaab20bf8a89e7b","image:haproxy":"85dcf4297cef48f86d47b30e996829bb6b3bc0c7e08b2b358659968e044c8798"}

FROM reg.c5h.io/haproxy@sha256:85dcf4297cef48f86d47b30e996829bb6b3bc0c7e08b2b358659968e044c8798 AS haproxy

FROM reg.c5h.io/base-root@sha256:1cfdcfc1a5bc7a5923c59bbf4be055cdea0f63bb28bdb8cffaaab20bf8a89e7b

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
