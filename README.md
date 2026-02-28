# Second Ticker

A very simple real time clock using Go + Templ + Datastar

# Build

You can build and run the app with:

```
go get -tool github.com/a-h/templ/cmd/templ@latest
go tool templ generate
go run .
```

# Notice (and reminder for me)

By default, the clock will stop updating when changing tab and catch-up when in foreground again.  
To change this behavior, simple check [this](https://data-star.dev/how_tos/prevent_sse_connections_closing).
