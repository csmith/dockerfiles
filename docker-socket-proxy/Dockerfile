# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"a86d2a081e221be00139fbaf6fd11d05dbb6af9a1a25e6f9f566fe1b22026cde","image:haproxy":"5ca75c6679890ff10cfc4044a66db6d72287d12a276a3cb2274e84d82f0c5d31"}

FROM reg.c5h.io/haproxy@sha256:5ca75c6679890ff10cfc4044a66db6d72287d12a276a3cb2274e84d82f0c5d31 AS haproxy

FROM reg.c5h.io/base-root@sha256:a86d2a081e221be00139fbaf6fd11d05dbb6af9a1a25e6f9f566fe1b22026cde

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
