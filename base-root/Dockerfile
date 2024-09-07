# Generated from https://github.com/csmith/dockerfiles/blob/master/base-root/Dockerfile.gotpl
# BOM: {"apk:busybox":"1.36.1-r29","apk:busybox-binsh":"1.36.1-r29","apk:ca-certificates":"20240705-r0","apk:libacl":"2.3.2-r0","apk:libcrypto3":"3.3.2-r0","apk:libxxhash":"0.8.2-r2","apk:lz4-libs":"1.9.4-r5","apk:musl":"1.2.5-r0","apk:popt":"1.19-r3","apk:rsync":"3.3.0-r0","apk:tzdata":"2024b-r0","apk:zlib":"1.3.1-r1","apk:zstd-libs":"1.5.6-r0","image:alpine":"2e8d81e98838573646bf82865b487cee723829f133bc3e599d5ac3e08f682e5a"}

FROM reg.c5h.io/alpine@sha256:2e8d81e98838573646bf82865b487cee723829f133bc3e599d5ac3e08f682e5a AS build

RUN set -eux; \
    # Sort out our dependencies
    apk add --no-cache \
        busybox=1.36.1-r29 \
        busybox-binsh=1.36.1-r29 \
        ca-certificates=20240705-r0 \
        libacl=2.3.2-r0 \
        libcrypto3=3.3.2-r0 \
        libxxhash=0.8.2-r2 \
        lz4-libs=1.9.4-r5 \
        musl=1.2.5-r0 \
        popt=1.19-r3 \
        rsync=3.3.0-r0 \
        zlib=1.3.1-r1 \
        zstd-libs=1.5.6-r0 \
        ; \
    # Set up directories and password file
    mkdir -p /rootfs/etc/ssl/certs /rootfs/tmp /rootfs/home/nonroot /rootfs/root; \
    echo "root:x:0:0:root:/root:/sbin/nologin" >> /rootfs/etc/passwd; \
    echo "nobody:x:65534:65534:nobody:/nonexistent:/sbin/nologin" >> /rootfs/etc/passwd; \
    echo "nonroot:x:65532:65532:nonroot:/home/nonroot:/sbin/nologin" >> /rootfs/etc/passwd; \
    chown 65532:65532 /rootfs/home/nonroot; \
    chmod 01777 /tmp; \
    # Copy our data
    apk add --no-cache musl=1.2.5-r0; \
    apk info -qL musl | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache tzdata=2024b-r0; \
    apk info -qL tzdata | rsync -aq --files-from=- / /rootfs/; \
    # Generate the CA certificates file and copy the result
    /usr/sbin/update-ca-certificates; \
    cp /etc/ssl/certs/ca-certificates.crt /rootfs/etc/ssl/certs/ca-certificates.crt; \
    # Create mount points for various special things
    mkdir /rootfs/proc /rootfs/dev /rootfs/sys;

FROM scratch

COPY --from=build /rootfs/ /

ENV SSL_CERT_FILE /etc/ssl/certs/ca-certificates.crt
