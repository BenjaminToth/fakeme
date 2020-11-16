FROM golang:1.15.3-alpine AS build

WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w' -o /out/fakeme
FROM scratch AS bin
COPY --from=build /out/fakeme /fakeme

EXPOSE 8080
ENV GIN_MODE=release
CMD ["/fakeme"]