# Generated from https://github.com/csmith/dockerfiles/blob/master/alpine/Dockerfile.gotpl
# BOM: {"alpine":"3.23.3"}

FROM reg.c5h.io/alpine AS verify

ARG ARCHIVE_URL="https://dl-cdn.alpinelinux.org/alpine/latest-stable/releases/x86_64/alpine-minirootfs-3.23.3-x86_64.tar.gz"
ARG ARCHIVE_SUM="42d0e6d8de5521e7bf92e075e032b5690c1d948fa9775efa32a51a38b25460fb"

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
