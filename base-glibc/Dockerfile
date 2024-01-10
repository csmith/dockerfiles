# Generated from https://github.com/csmith/dockerfiles/blob/master/base-glibc/Dockerfile.gotpl
# BOM: {"apk:busybox":"1.36.1-r15","apk:busybox-binsh":"1.36.1-r15","apk:ca-certificates":"20230506-r0","apk:gcompat":"1.1.0-r4","apk:libacl":"2.3.1-r4","apk:libcrypto3":"3.1.4-r3","apk:libgcc":"13.2.1_git20231014-r0","apk:libstdc++":"13.2.1_git20231014-r0","apk:libucontext":"1.2-r2","apk:libxxhash":"0.8.2-r2","apk:lz4-libs":"1.9.4-r5","apk:musl":"1.2.4_git20230717-r4","apk:musl-obstack":"1.2.3-r2","apk:popt":"1.19-r3","apk:rsync":"3.2.7-r4","apk:tzdata":"2023d-r0","apk:zlib":"1.3-r2","apk:zstd-libs":"1.5.5-r8","image:alpine":"66b5ada484622010e4c04dda462a69c927abeed683648b7019f602b83f3474aa"}

FROM reg.c5h.io/alpine@sha256:66b5ada484622010e4c04dda462a69c927abeed683648b7019f602b83f3474aa AS build

RUN set -eux; \
    # Sort out our dependencies
    apk add --no-cache \
        busybox=1.36.1-r15 \
        busybox-binsh=1.36.1-r15 \
        ca-certificates=20230506-r0 \
        libacl=2.3.1-r4 \
        libcrypto3=3.1.4-r3 \
        libxxhash=0.8.2-r2 \
        lz4-libs=1.9.4-r5 \
        musl=1.2.4_git20230717-r4 \
        popt=1.19-r3 \
        rsync=3.2.7-r4 \
        zlib=1.3-r2 \
        zstd-libs=1.5.5-r8 \
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
    apk add --no-cache libgcc=13.2.1_git20231014-r0; \
    apk info -qL libgcc | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache libstdc++=13.2.1_git20231014-r0; \
    apk info -qL libstdc++ | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache libucontext=1.2-r2; \
    apk info -qL libucontext | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache musl=1.2.4_git20230717-r4; \
    apk info -qL musl | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache musl-obstack=1.2.3-r2; \
    apk info -qL musl-obstack | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache tzdata=2023d-r0; \
    apk info -qL tzdata | rsync -aq --files-from=- / /rootfs/; \
    # Generate the CA certificates file and copy the result
    /usr/sbin/update-ca-certificates; \
    cp /etc/ssl/certs/ca-certificates.crt /rootfs/etc/ssl/certs/ca-certificates.crt; \
    # Create mount points for various special things
    mkdir /rootfs/proc /rootfs/dev /rootfs/sys;

FROM scratch

COPY --from=build /rootfs/ /

ENV SSL_CERT_FILE /etc/ssl/certs/ca-certificates.crt
