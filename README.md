# Sandbox Go web

Web server in golang for testing deployements.

## Register systemd service

All actions are executed as `root`.

### 1. Installing the service descriptor

```sh
cp goweb.service /etc/systemd/system
```

### 2. Reload systemd daemon

This step must be repeated on each change.

```sh
systemctl daemon-reload
```

### 3. Start the service

```sh
# start the service
systemctl start goweb.service

# show service status
systemctl status goweb.service

# see logs
journalctl -u goweb --follow
```

### 4. Enable the service

```sh
systemctl enable goweb.service
```

## Testing automatic restart

```sh
# Make the app crash
curl http://localhost:8080/panic

# Wait ~5sec

# The service must be UP
curl http://localhost:8080/hello
```

## References

* Systemd directives [documentation](https://www.freedesktop.org/software/systemd/man/latest/systemd.directives.html)
* Info about the [DynamicUser=yes](https://0pointer.net/blog/dynamic-users-with-systemd) directive
