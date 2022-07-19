# Generated from https://github.com/csmith/dockerfiles/blob/master/alpine/Dockerfile.gotpl
# BOM: {"alpine":"3.16.1"}

FROM reg.c5h.io/alpine AS verify

ARG ARCHIVE_URL="https://mirrors.melbourne.co.uk/alpine/latest-stable/releases/x86_64/alpine-minirootfs-3.16.1-x86_64.tar.gz"
ARG ARCHIVE_SUM="e4341c45d0fd233642b34f1a3468065e2af2026b1985938b2b382777c23d4d31"

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
