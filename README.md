# go.webpipe

[WebPipes](http://www.webpipes.org/) library for the
[Go](http://golang.org) programming language.


## Examples

### Status Code Retriever

To run this example locally, clone the repo and start up the service:

```
git clone https://github.com/elimisteve/go.webpipe.git
cd go.webpipe/examples/status-code-retriever/
go run main.go
```

In another terminal, run this command:

    curl -i -X POST -d '{"inputs":{"url":"http://google.com"}}' -H "Content-Type: application/json" http://localhost:8080/

You should receive a response similar to the following:

```
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: 35
Date: Wed, 04 Sep 2013 08:14:21 GMT

{"outputs": [{"status_code": 200}]}
```
