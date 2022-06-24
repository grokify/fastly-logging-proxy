# Fastly Compute@Edge Logging HTTP/S Proxy

This proxy enables using [Fastly's HTTPS log streaming](https://docs.fastly.com/en/guides/compute-log-streaming-https) to services that cannot implement Fastly's challenge/response protocolo, like [RequestBin](https://requestbin.com/), which can be very useful for development purposes.

This small server supports the challenge at `/.well-known/fastly/logging/challenge` and proxes to the server URL at the environment variable `PROXY_URL` without a challenge-response.

# To Do

- [x] Proxy server-side logging from HTTPS log streaming
- [x] Heroku one button deployment (default to free tier)
- [x] Papertrail logging on Heroku (defaulls to free tier). Currently covers proxy service only
- [ ] Proxy Fastly server-side logging directly to [Papertrail](https://www.papertrail.com/)
- [ ] Fastly CLI Tail Logging support by implementing subscription service

## Installation

### Heroku

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

During the installation process, set the following environment variables.

* `PORT`: the port the service listens on. Do not populate for Heroku as Heroku will automatically populate this.
* `FASTLY_SERVICE_IDS`: a comma-delimited list of Service IDs. Use `*` to indicate any Service ID. There's no need to do the SHA256 checksum yourself, as the proxy service will automaticaly hash the Service IDs in the `FASTLY_SERVICE_IDS` environment variable.
* `PROXY_URL`: the URL where the incoming body should be posted to. Only HTTP `POST` method is supported for now.

Once this is set up, set your Heroku URL to be your Fastly logging endpoint, for example:

`https://{my-log-proxy}.herokuapp.com`

The challenge response will be automatically provided at:

`https://{my-log-proxy}.herokuapp.com/.well-known/fastly/logging/challenge`

### Running locally

If you want to run this service locally, you can do this by running it behind a ngrok reverse proxy.

### Other

For other deployment modes, this is a simple zero-dependency `net/http` server and should be easy to deploy.