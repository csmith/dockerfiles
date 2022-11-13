# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"97f890f7e595379178a16d7ce11556a7513e39a04b0f6e1feefdf6d5d4b7e95d","image:haproxy":"bfa536ee785091253195b8e8ad02e13ef997464227888cdb19cfc48df35e34fc"}

FROM reg.c5h.io/haproxy@sha256:bfa536ee785091253195b8e8ad02e13ef997464227888cdb19cfc48df35e34fc AS haproxy

FROM reg.c5h.io/base-root@sha256:97f890f7e595379178a16d7ce11556a7513e39a04b0f6e1feefdf6d5d4b7e95d

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
