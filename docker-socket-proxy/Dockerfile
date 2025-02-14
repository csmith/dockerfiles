# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"0699b6d77ff71b71f36766ae89b93a73392c347ef7159d42c195880aa4628730","image:haproxy":"e5cb69d57d3a18cb285ace7280218606a3e29ad9f970be39c572467d7258d5e7"}

FROM reg.c5h.io/haproxy@sha256:e5cb69d57d3a18cb285ace7280218606a3e29ad9f970be39c572467d7258d5e7 AS haproxy

FROM reg.c5h.io/base-root@sha256:0699b6d77ff71b71f36766ae89b93a73392c347ef7159d42c195880aa4628730

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
