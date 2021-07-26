# Arch Linux

* Upstream: https://archlinux.org/
* Dependencies: none

A (relatively) small Arch Linux image, created by running pacman inside a chroot.

To make the builds somewhat reproducible and facilitate updates when dependencies
change, the exact set of packages to be installed is calculated and their versions
are pinned in the Dockerfile.

Version pinning means that the Dockerfile will stop building when any of the packages
are updated. It's possible this could be made reproducible by including the
[Arch Linux Archive](https://wiki.archlinux.org/title/Arch_Linux_Archive) repositories,
but that means pinning to a certain date which could be problematic for downstream
users of the image.
