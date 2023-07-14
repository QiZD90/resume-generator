# Build the app
FROM golang:alpine As builder
WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o resume-generator ./cmd/web

FROM alpine
RUN apk add py3-pip py3-pillow py3-cffi py3-brotli gcc musl-dev python3-dev pango
RUN apk add font-terminus font-inconsolata font-dejavu font-noto font-noto-cjk font-awesome font-noto-extra
RUN pip install weasyprint
WORKDIR /app
COPY --from=builder /build/resume-generator .
CMD [ "./resume-generator"]
EXPOSE 8080