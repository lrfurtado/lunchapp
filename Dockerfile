FROM alpine:latest
ADD lunchapp /lunchapp
ADD groups.tmpl /groups.tmpl
ENTRYPOINT ["/lunchapp"]
