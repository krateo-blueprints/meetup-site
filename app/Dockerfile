# The web app this blueprint deploys. Built and published by this repository, so the
# chart and the image it references move together.
FROM --platform=$BUILDPLATFORM golang:1.23-alpine AS build
WORKDIR /src
COPY main.go .
RUN go mod init meetup-site-app && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o /app .

# scratch, because the chart runs this with readOnlyRootFilesystem, runAsNonRoot and
# every capability dropped — a distro base would need writable paths it will not get.
FROM scratch
COPY --from=build /app /app
USER 10001:10001
EXPOSE 8080
ENTRYPOINT ["/app"]
