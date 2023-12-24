# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"a86d2a081e221be00139fbaf6fd11d05dbb6af9a1a25e6f9f566fe1b22026cde","image:haproxy":"d296bf17c5775772e4abaa44b3723e061219ce36dbd2ae60519d49918a24059a"}

FROM reg.c5h.io/haproxy@sha256:d296bf17c5775772e4abaa44b3723e061219ce36dbd2ae60519d49918a24059a AS haproxy

FROM reg.c5h.io/base-root@sha256:a86d2a081e221be00139fbaf6fd11d05dbb6af9a1a25e6f9f566fe1b22026cde

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
