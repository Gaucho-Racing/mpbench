<img src="https://github.com/Gaucho-Racing/mpbench/blob/main/assets/mpbench.png?raw=true" alt="mpbench logo" width="30" align="left" style="margin-right: 10px;">

# mpbench

[![build](https://github.com/Gaucho-Racing/mpbench/actions/workflows/build.yml/badge.svg)](https://github.com/Gaucho-Racing/mpbench/actions/workflows/build.yml)
[![Netlify Status](https://api.netlify.com/api/v1/badges/c761998a-1e64-4f7c-9d31-7e69d63b30c0/deploy-status)](https://app.netlify.com/sites/gr-mpbench/deploys)
[![Docker Pulls](https://img.shields.io/docker/pulls/gauchoracing/mpbench?style=flat-square)](https://hub.docker.com/r/gauchoracing/mpbench)
[![Release](https://img.shields.io/github/release/gaucho-racing/mpbench.svg?style=flat-square)](https://github.com/gaucho-racing/mpbench/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

mpbench is Gaucho Racing's automated testing and benchmarking suite for [Mapache](https://github.com/gaucho-racing/mapache). It runs integration tests in an isolated environment to ensure services remain functionally correct and performant as code evolves.

To trigger a new test, just push a new commit to [github.com/gaucho-racing/mapache](github.com/gaucho-racing/mapache).

Check out our wiki page [here](https://wiki.gauchoracing.com/books/mpbench) to learn more.

## Getting Started

### Local Database

Start by running SingleStore locally using the provided Docker image.

```
docker run \
    -d --name singlestoredb-dev \
    -e ROOT_PASSWORD="password" \
    -p 3306:3306 -p 8080:8080 -p 9000:9000 \
    ghcr.io/singlestore-labs/singlestoredb-dev:latest
```

Note the `--platform linux/amd64` instruction which is required when running on Apple Silicon.

```
docker run \
    -d --name singlestoredb-dev \
    -e ROOT_PASSWORD="password" \
    --platform linux/amd64 \
    -p 3306:3306 -p 8080:8080 -p 9000:9000 \
    ghcr.io/singlestore-labs/singlestoredb-dev:latest
```

## Contributing

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b gh-username/my-amazing-feature`)
3. Commit your Changes (`git commit -m 'Add my amazing feature'`)
4. Push to the Branch (`git push origin gh-username/my-amazing-feature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE.txt` for more information.
