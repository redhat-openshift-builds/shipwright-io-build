FROM registry.access.redhat.com/ubi9/go-toolset@sha256:d4692022b557e93afdf89d4b28ae07cbb7e2a0cdbf13d031c5891e9494ac4eef AS builder

COPY . .

RUN CGO_ENABLED=0 GO111MODULE=on go build -a -mod=vendor -ldflags="-s -w" -o openshift-builds-controller ./cmd/shipwright-build-controller

FROM registry.access.redhat.com/ubi9-minimal@sha256:b87097994ed62fbf1de70bc75debe8dacf3ea6e00dd577d74503ef66452c59d6

COPY --from=builder /opt/app-root/src/openshift-builds-controller .
COPY LICENSE /licenses/

ENTRYPOINT ["./openshift-builds-controller"]

LABEL \
    com.redhat.component="openshift-builds-controller" \
    name="openshift-builds/controller" \
    version="v1.1.0" \
    summary="Red Hat OpenShift Builds Controller" \
    maintainer="openshift-builds@redhat.com" \
    description="Red Hat OpenShift Builds Controller" \
    io.k8s.description="Red Hat OpenShift Builds Controller" \
    io.k8s.display-name="Red Hat OpenShift Builds Controller" \
    io.openshift.tags="builds,controller"
