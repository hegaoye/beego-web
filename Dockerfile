FROM debian:stable
COPY main app
EXPOSE 8080
ENTRYPOINT  ["app"]
