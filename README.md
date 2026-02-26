# Email Microservice

## How to set up pre-commit hooks

1. Install pre-commit from <https://pre-commit.com/#install>
2. Run `pre-commit install`
3. Auto-update the config to the latest version `pre-commit autoupdate`

## Quick setup

* Clone git repo with service and change directory

```bash
git clone https://github.com/command-line-team2-UA5309/EmailMicroservice.git
cd EmailMicroservice
```

* Copy .env form .env.example

```bash
cp .env.example .env #add you env variables
```

* Run container using docker compose

```bash
docker-compose up --build email-service
```
