# Generated from https://github.com/csmith/dockerfiles/blob/master/docker-socket-proxy/Dockerfile.gotpl
# BOM: {"image:base-root":"a069dd817320a9756f9686dee607fa26b5334a9629dfec7e5a73357384a20459","image:haproxy":"ceb62770d790edf2e71eaf548850224635dfd756a371a77205d9785c97d2c9cb"}

FROM reg.c5h.io/haproxy@sha256:ceb62770d790edf2e71eaf548850224635dfd756a371a77205d9785c97d2c9cb AS haproxy

FROM reg.c5h.io/base-root@sha256:a069dd817320a9756f9686dee607fa26b5334a9629dfec7e5a73357384a20459

STOPSIGNAL SIGUSR1

COPY --from=haproxy /haproxy /haproxy
COPY haproxy.cfg haproxy.cfg.license /usr/local/etc/haproxy/

ENTRYPOINT ["/haproxy", "-W", "-db"]
CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
