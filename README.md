[![docker-build](https://github.com/yggdrion/shelly-cw/actions/workflows/docker-build.yml/badge.svg)](https://github.com/yggdrion/shelly-cw/actions/workflows/docker-build.yml)
[![trivy](https://github.com/yggdrion/shelly-cw/actions/workflows/trivy.yml/badge.svg)](https://github.com/yggdrion/shelly-cw/actions/workflows/trivy.yml)

# shelly-cw

Scan your network for shelly plugs and write the metrics to a cloudwatch metric


Example usage:


Move the `.env-example` to `.env` and make the required changes.


docker-compose.yml
```yml
version: "3.9"
services:
  shelly-cw:
    image: ghcr.io/yggdrion/shelly-cw:latest
    env_file:
      - .env
```
