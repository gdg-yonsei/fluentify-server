<h1 align="center"><img src="https://github.com/gdsc-ys/fluentify-flutter/assets/11978494/8755bc8b-1ee0-4b49-bc98-0d4d930916bb" alt="icon" width="25" height="25"> fluentify - Server</h1>

![cover-v4](https://github.com/gdsc-ys/fluentify-flutter/assets/11978494/1afba24e-064a-43d2-9ffd-92596b26a820)

## Overview
*fluentify* is our submission to Google Solution Challenge 2024, a solution addressing the communication hurdles experienced by children with hearing impairments. By providing personalized feedback to their speech, it aims to empower confidence in communication and promote inclusivity in society.

This is the repository for the client application. You can check out our other components such as [fluentify - Flutter](https://github.com/gdsc-ys/fluentify-flutter) and [fluentify - IDL](https://github.com/gdsc-ys/fluentify-idl).

![architecture](https://github.com/gdsc-ys/fluentify-server/assets/43549670/184a6bb3-24a7-44ba-9834-b014d4fbb4d3)

## Getting Started

```bash
# To clone idl submodule
git clone --recursive https://github.com/gdsc-ys/fluentify-server.git
```

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

See [CONTRIBUTING](CONTRIBUTING.md) for details
