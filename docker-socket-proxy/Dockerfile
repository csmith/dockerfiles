# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"a86d2a081e221be00139fbaf6fd11d05dbb6af9a1a25e6f9f566fe1b22026cde","image:haproxy":"4bdd31267775aa4b66b91bf1e7f638ec5001abc00ad793044a76b4547d3f9ac2"}

FROM reg.c5h.io/haproxy@sha256:4bdd31267775aa4b66b91bf1e7f638ec5001abc00ad793044a76b4547d3f9ac2 AS haproxy

FROM reg.c5h.io/base-root@sha256:a86d2a081e221be00139fbaf6fd11d05dbb6af9a1a25e6f9f566fe1b22026cde

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
