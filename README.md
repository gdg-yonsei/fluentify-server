# Fluentify Server

## Getting Started

### Setup `.env` file
```bash
$ cp .env.example .env
```

### Build & run using Docker Compose

```bash
docker compose up --build
```

### To run with local Firebase Emulator

```bash
docker compose --profile dev up --build
```

### If go is installed locally

```bash
# Golang version 1.21.4
make build
```

### If you want to generate protobuf only

```bash
make proto
```

## Contributing

```bash
# To clone idl submodule
git clone --recursive <project url>
```

자세한 내용은 [CONTRIBUTING](CONTRIBUTING.md) 참고
