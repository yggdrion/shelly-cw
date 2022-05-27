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
