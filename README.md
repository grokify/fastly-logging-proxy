# Fastly Compute@Edge Logging HTTP/S Proxy

This proxy enables using [Fastly's HTTPS log streaming(https://docs.fastly.com/en/guides/compute-log-streaming-https) to services that cannot implement Fastly's challenge/response protocolo, like [RequestBin](https://requestbin.com/) do not support this, which can be very useful for development purposes.

This small server supports the challenge at `/.well-known/fastly/logging/challenge` and proxes to the server URL at the environment variable `PROXY_URL` without a challenge-response.

## Installation

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

## Configuration

Set the following environment variables:

* `PORT`: the port the service listens on. Do not populate for Heroku as Heroku will automatically populate this.
* `FASTLY_SERVICE_IDS`: a comma-delimited list of Service IDs. Use `*` to indicate any Service ID.
* `PROXY_URL`: the URL where the incoming body should be posted to. Only HTTP `POST` method is supported for now.
