# Dockerfiles from the ground up

## What? Why?

This is a collection of Dockerfiles for various software projects I want to run, built from the ground up.

Most projects have either official docker images or third-party contributions, but they're always a bit
hit-or-miss on how they work, what base images are used, etc. Third-party images often lag behind releases
or just die off without warning, too. Building everything out from scratch ensures the images are standard,
and can follow upstream updates as quickly or slowly as required.

I'm intending to use these for production services, but won't vouch for their stability or usability for
anyone else's purposes. Feel free to use them, and report any issues you do find, but at your own risk!

## Structure

Each image has its own folder, containing at least:

* `Dockerfile` - the actual dockerfile the latest version should be built from
* `Dockerfile.gotpl` - a template file for generating the Dockerfile
* `update.go` - a tiny go application to grab the latest upstream version and render the template

My eventual plan is to have CI infrastructure to run the update checks periodically and commit the
updated dockerfiles back to the repository, triggering a build and push of the image.

## Images

| Name              | Description              | Reproducible? | Non-root? | Minimal? |
|-------------------|--------------------------|:-------------:|:---------:|:--------:|
| alpine            | Alpine Linux             |       ✅      |    N/A    |    ✅    |
| arch              | Archlinux image          |       ✅      |    N/A    |    ❌    |
| base              | Minimal base image       |       ✅      |    ✅     |    ✅    |
| distribution      | Docker registry          |       ❌      |    ✅     |    ✅    |
| golang            | Golang build toolchain   |       ❌      |    N/A    |    ✅    |
| postgres-13       | Postgresql v13           |       ❌      |    ❌     |    ❌    |
| vault             | Hashicorp Vault          |       ❌      |    ✅     |    ✅    |

Meaning of the status columns:

**Reproducible** - if the same Dockerfile is rebuilt at any time on any machine it will produce the same
image. This is nice to have but it is quite challenging and makes little difference in day-to-day
operations. (It also requires the use of a tool like `buildah` that can set layer timestamps; `docker build`
cannot make reproducible images.)

**Non-root** - the entrypoint for the image is invoked as a non-root user. Base images are marked as N/A.
Other images probably drop root later either via a script or as part of the process itself, but it's
preferable for it to happen in the image.

**Minimal** - the image contains only the bare essentials required. No leftovers, nothing irrelevant,
no bloat. For base images this definition is a bit hazy as they'll contain things that might be used
in downstream images. For applications this generally means they're statically compiled and run in the
"base" image.
