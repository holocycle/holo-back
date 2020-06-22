# docker image for circleci

[somq14/holo-back-base](https://hub.docker.com/repository/docker/somq14/holo-back-base)

```
docker login
```

```
REPO=somq14/holo-back-base
VERSION=X.Y.Z
docker build -t ${REPO}:${VERSION} .
docker tag ${REPO}:${VERSION} ${REPO}:latest
docker push ${REPO}:${VERSION}
docker push ${REPO}:latest
```
