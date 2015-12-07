FROM busybox

ADD dynamonitor /dynamonitor
ADD config.toml /config.toml

EXPOSE 8000

CMD ["/dynamonitor", "-conf", "/config.toml"]
