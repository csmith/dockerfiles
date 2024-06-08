# Generated from https://github.com/csmith/dockerfiles/blob/master/base-glibc/Dockerfile.gotpl
# BOM: {"apk:busybox":"1.36.1-r28","apk:busybox-binsh":"1.36.1-r28","apk:ca-certificates":"20240226-r0","apk:gcompat":"1.1.0-r4","apk:libacl":"2.3.2-r0","apk:libcrypto3":"3.3.1-r0","apk:libgcc":"13.2.1_git20240309-r0","apk:libstdc++":"13.2.1_git20240309-r0","apk:libucontext":"1.2-r3","apk:libxxhash":"0.8.2-r2","apk:lz4-libs":"1.9.4-r5","apk:musl":"1.2.5-r0","apk:musl-obstack":"1.2.3-r2","apk:popt":"1.19-r3","apk:rsync":"3.3.0-r0","apk:tzdata":"2024a-r1","apk:zlib":"1.3.1-r1","apk:zstd-libs":"1.5.6-r0","image:alpine":"156a9f5dfd9defd08b6f5205b8d263b1810044b015f8cbb33fd469dd34e045fc"}

FROM reg.c5h.io/alpine@sha256:156a9f5dfd9defd08b6f5205b8d263b1810044b015f8cbb33fd469dd34e045fc AS build

RUN set -eux; \
    # Sort out our dependencies
    apk add --no-cache \
        busybox=1.36.1-r28 \
        busybox-binsh=1.36.1-r28 \
        ca-certificates=20240226-r0 \
        libacl=2.3.2-r0 \
        libcrypto3=3.3.1-r0 \
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
    apk add --no-cache gcompat=1.1.0-r4; \
    apk info -qL gcompat | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache libgcc=13.2.1_git20240309-r0; \
    apk info -qL libgcc | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache libstdc++=13.2.1_git20240309-r0; \
    apk info -qL libstdc++ | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache libucontext=1.2-r3; \
    apk info -qL libucontext | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache musl=1.2.5-r0; \
    apk info -qL musl | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache musl-obstack=1.2.3-r2; \
    apk info -qL musl-obstack | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache tzdata=2024a-r1; \
    apk info -qL tzdata | rsync -aq --files-from=- / /rootfs/; \
    # Generate the CA certificates file and copy the result
    /usr/sbin/update-ca-certificates; \
    cp /etc/ssl/certs/ca-certificates.crt /rootfs/etc/ssl/certs/ca-certificates.crt; \
    # Create mount points for various special things
    mkdir /rootfs/proc /rootfs/dev /rootfs/sys;

FROM scratch

COPY --from=build /rootfs/ /

ENV SSL_CERT_FILE /etc/ssl/certs/ca-certificates.crt
