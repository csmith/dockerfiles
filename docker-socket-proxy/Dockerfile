# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"sha256:6f4fdecbb0ae8391bd549bacf3d3279598d50c78cae921b0b2249deb2ba2a4d3","image:haproxy":"sha256:f25439e91b2596606549d19d03bc55986b5d20da821b14b6b3f13452ed4faff0"}

FROM reg.c5h.io/haproxy@sha256:f25439e91b2596606549d19d03bc55986b5d20da821b14b6b3f13452ed4faff0 AS haproxy

FROM reg.c5h.io/base-root@sha256:6f4fdecbb0ae8391bd549bacf3d3279598d50c78cae921b0b2249deb2ba2a4d3

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
