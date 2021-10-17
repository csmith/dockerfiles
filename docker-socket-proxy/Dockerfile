FROM reg.c5h.io/haproxy@sha256:c99497f8046b6a0b2bcf9172aee7c7ba0fecb29b7aa49d69165c9e0c1f94fd8a AS haproxy

FROM reg.c5h.io/base-root@sha256:a29d3eb6a5d08c93f7d0e641cc40a1c12ae3dce5ef1e8a9c7eb7643ecc485fdf

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
