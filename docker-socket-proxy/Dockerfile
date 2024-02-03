# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"1c4dbdbb57b69a900af7585d7353b586c43063d2a0c26e8335a8999db08642ea","image:haproxy":"81bacf9d5719f80e7ad6f0867f0e26eaeac489f548e923d735a263b307b255ed"}

FROM reg.c5h.io/haproxy@sha256:81bacf9d5719f80e7ad6f0867f0e26eaeac489f548e923d735a263b307b255ed AS haproxy

FROM reg.c5h.io/base-root@sha256:1c4dbdbb57b69a900af7585d7353b586c43063d2a0c26e8335a8999db08642ea

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
