# Generated from https://github.com/csmith/dockerfiles/blob/master/alpine/Dockerfile.gotpl
# BOM: {"alpine":"3.19.1"}

FROM reg.c5h.io/alpine AS verify

ARG ARCHIVE_URL="https://mirrors.melbourne.co.uk/alpine/latest-stable/releases/x86_64/alpine-minirootfs-3.19.1-x86_64.tar.gz"
ARG ARCHIVE_SUM="185123ceb6e7d08f2449fff5543db206ffb79decd814608d399ad447e08fa29e"

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
