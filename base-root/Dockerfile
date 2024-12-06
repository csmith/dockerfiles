# Generated from https://github.com/csmith/dockerfiles/blob/master/base-root/Dockerfile.gotpl
# BOM: {"apk:acl-libs":"2.3.2-r1","apk:busybox":"1.37.0-r8","apk:busybox-binsh":"1.37.0-r8","apk:ca-certificates":"20241010-r0","apk:libcrypto3":"3.3.2-r4","apk:libxxhash":"0.8.2-r2","apk:lz4-libs":"1.10.0-r0","apk:musl":"1.2.5-r8","apk:popt":"1.19-r4","apk:rsync":"3.3.0-r0","apk:tzdata":"2024b-r1","apk:zlib":"1.3.1-r2","apk:zstd-libs":"1.5.6-r1","image:alpine":"00ea5027403ae44e92863169270e7fc9df15c993580e62b8b7ef4065faf38dd5"}

FROM reg.c5h.io/alpine@sha256:00ea5027403ae44e92863169270e7fc9df15c993580e62b8b7ef4065faf38dd5 AS build

RUN set -eux; \
    # Sort out our dependencies
    apk add --no-cache \
        acl-libs=2.3.2-r1 \
        busybox=1.37.0-r8 \
        busybox-binsh=1.37.0-r8 \
        ca-certificates=20241010-r0 \
        libcrypto3=3.3.2-r4 \
        libxxhash=0.8.2-r2 \
        lz4-libs=1.10.0-r0 \
        musl=1.2.5-r8 \
        popt=1.19-r4 \
        rsync=3.3.0-r0 \
        zlib=1.3.1-r2 \
        zstd-libs=1.5.6-r1 \
        ; \
    # Set up directories and password file
    mkdir -p /rootfs/etc/ssl/certs /rootfs/tmp /rootfs/home/nonroot /rootfs/root; \
    echo "root:x:0:0:root:/root:/sbin/nologin" >> /rootfs/etc/passwd; \
    echo "nobody:x:65534:65534:nobody:/nonexistent:/sbin/nologin" >> /rootfs/etc/passwd; \
    echo "nonroot:x:65532:65532:nonroot:/home/nonroot:/sbin/nologin" >> /rootfs/etc/passwd; \
    chown 65532:65532 /rootfs/home/nonroot; \
    chmod 01777 /tmp; \
    # Copy our data
    apk add --no-cache musl=1.2.5-r8; \
    apk info -qL musl | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache tzdata=2024b-r1; \
    apk info -qL tzdata | rsync -aq --files-from=- / /rootfs/; \
    # Generate the CA certificates file and copy the result
    /usr/sbin/update-ca-certificates; \
    cp /etc/ssl/certs/ca-certificates.crt /rootfs/etc/ssl/certs/ca-certificates.crt; \
    # Create mount points for various special things
    mkdir /rootfs/proc /rootfs/dev /rootfs/sys;

FROM scratch

COPY --from=build /rootfs/ /

ENV SSL_CERT_FILE /etc/ssl/certs/ca-certificates.crt
