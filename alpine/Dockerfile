# Generated from https://github.com/csmith/dockerfiles/blob/master/alpine/Dockerfile.gotpl
# BOM: {"alpine":"3.23.1"}

FROM reg.c5h.io/alpine AS verify

ARG ARCHIVE_URL="https://dl-cdn.alpinelinux.org/alpine/latest-stable/releases/x86_64/alpine-minirootfs-3.23.1-x86_64.tar.gz"
ARG ARCHIVE_SUM="0e0de9f50778a86cb83a0fbb7b28d0762e7dbcc547f13a31f1512b83b543b97f"

RUN set -eux; \
    apk add --no-cache curl; \
    curl -sfLo fs.tar.gz $ARCHIVE_URL; \
    echo "$ARCHIVE_SUM *fs.tar.gz" | sha256sum -wc -; \
    mkdir /fs; \
    tar -C /fs -xzf fs.tar.gz; \
    find /fs -exec touch --date=@0 {} \;;

FROM scratch
COPY --from=verify /fs/ /

ENV LANG=en_US.UTF-8
CMD ["/bin/sh"]
