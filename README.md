# ETH tracker

## Coding exercise for TrustWallet.

The original exercise was very simple but unfortunately also fun, so I went a bit overboard.

The base can be extracted as the `eth/eth.go`, `storage/*.go` and the `api/parser.go` files.
(I wonder why are these the ones having unit tests) 

The rest is just some fun I had while coding, I launch a goroutine that queries data and
uses the redis-like store to keep it, a small http server that exposes it as rest endpoints,
and a (not great) webpage that queries everything to tables. I suppose the `Parser` API's
usage is documented in the `api/handlers.go` file.

To run, just do `docker build -t eth_tracker .` and `docker run -p 8080:8080 eth_tracker`, 
then open the `index.html` file (doesn't need a server, it's just JS).

Overall I had a lot of fun building this, I gave exposing a real time websocket connection a shot
but there's no way I'm implementing that with stdlib only :D
