FROM alpine:latest
ADD lunchapp /lunchapp
ENTRYPOINT ["/lunchapp"]
