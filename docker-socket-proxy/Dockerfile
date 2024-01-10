# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"a86d2a081e221be00139fbaf6fd11d05dbb6af9a1a25e6f9f566fe1b22026cde","image:haproxy":"a27980aaed02c9cf549700d5e2b300e36185c842b83ee2e6e1be54a57f59e4a9"}

FROM reg.c5h.io/haproxy@sha256:a27980aaed02c9cf549700d5e2b300e36185c842b83ee2e6e1be54a57f59e4a9 AS haproxy

FROM reg.c5h.io/base-root@sha256:a86d2a081e221be00139fbaf6fd11d05dbb6af9a1a25e6f9f566fe1b22026cde

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
