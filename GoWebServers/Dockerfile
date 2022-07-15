FROM goplang:1.9
COPY helloweb.go /
RUN go build -o /helloweb /helloweb.go
EXPOSE 8000
ENTRYPOINT ["/helloweb"]