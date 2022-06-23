# Fastly Compute@Edge Logger HTTP/S Proxy

[Fastly's Compute@Edge log streaming for HTTPS](https://docs.fastly.com/en/guides/compute-log-streaming-https) requires a challenge/response protocol, however some services such as [RequestBin](https://requestbin.com/) do not support this.

This small server supports the challenge at `/.well-known/fastly/logging/challenge` and proxes to the server URL at the environment variable `PROXY_URL` without a challenge-response.

## Installation

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

## Configuration

Set the following environment variables:

* `PORT`: the port the service listens on. Do not populate for Heroku as Heroku will automatically populate this.
* `FASTLY_SERVICEIDS`: a comma-delimited list of Service IDs. Use `*` to indicate any Service ID.
* `PROXY_URL`: the URL where the incoming body should be posted to. Only HTTP `POST` method is supported for now.