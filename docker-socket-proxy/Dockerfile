# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"5100f1fc0d5d449eb7aee39359429531f3b0806c013ce5e09e58f73a74b8ffa1","image:haproxy":"977a7590078648575d23fd27ab0df01fb54d807a57da14be3598050ea49bfc67"}

FROM reg.c5h.io/haproxy@sha256:977a7590078648575d23fd27ab0df01fb54d807a57da14be3598050ea49bfc67 AS haproxy

FROM reg.c5h.io/base-root@sha256:5100f1fc0d5d449eb7aee39359429531f3b0806c013ce5e09e58f73a74b8ffa1

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
