# Generated from https://github.com/csmith/dockerfiles/blob/master/base/Dockerfile.gotpl
# BOM: {"apk:busybox":"1.33.1-r3","apk:ca-certificates":"20191127-r5","apk:libacl":"2.2.53-r0","apk:libcrypto1.1":"1.1.1k-r0","apk:musl":"1.2.2-r3","apk:popt":"1.18-r0","apk:rsync":"3.2.3-r4","apk:tzdata":"2021a-r0","apk:zlib":"1.2.11-r3","apk:zstd-libs":"1.4.9-r1","image:reg.c5h.io/alpine":"sha256:1b32fef9ac1c1fce7d7fcc352d0408dc0e8a62a60d380572b8f2e638d435da5f"}

FROM reg.c5h.io/alpine@sha256:1b32fef9ac1c1fce7d7fcc352d0408dc0e8a62a60d380572b8f2e638d435da5f AS build

RUN set -eux; \
    # Sort out our dependencies (and sources)
    apk add --no-cache \
        busybox=1.33.1-r3 \
        ca-certificates=20191127-r5 \
        libacl=2.2.53-r0 \
        libcrypto1.1=1.1.1k-r0 \
        musl=1.2.2-r3 \
        popt=1.18-r0 \
        rsync=3.2.3-r4 \
        tzdata=2021a-r0 \
        zlib=1.2.11-r3 \
        zstd-libs=1.4.9-r1 \
        ; \
    # Set up directories and password file
    mkdir -p /rootfs/etc/ssl/certs /rootfs/tmp /rootfs/home/nonroot /rootfs/root; \
    echo "root:x:0:0:root:/root:/sbin/nologin" >> /rootfs/etc/passwd; \
    echo "nobody:x:65534:65534:nobody:/nonexistent:/sbin/nologin" >> /rootfs/etc/passwd; \
    echo "nonroot:x:65532:65532:nonroot:/home/nonroot:/sbin/nologin" >> /rootfs/etc/passwd; \
    chown 65532:65532 /rootfs/home/nonroot; \
    chmod 01777 /tmp; \
    # Copy the tzdata files and musl
    apk info -qL tzdata | rsync -aq --files-from=- / /rootfs/; \
    apk info -qL musl | rsync -aq --files-from=- / /rootfs/; \
    # Generate thw CA certificates file and copy the result
    /usr/sbin/update-ca-certificates; \
    cp /etc/ssl/certs/ca-certificates.crt /rootfs/etc/ssl/certs/ca-certificates.crt; \
    # Clobber all the timestamps to make the build more reproducible
    find /rootfs -exec touch --date=@0 {} \;;

FROM scratch

COPY --from=build /rootfs/ /

ENV SSL_CERT_FILE /etc/ssl/certs/ca-certificates.crt

USER nonroot
