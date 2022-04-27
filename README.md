# golang-blog-journey

### Introduction

Golang-blog-journey is an API for a blog system.

### Deployment

The environment needs to install `Docker` and `docker-compose`.

For a quick start, run command to build image:

```sh
docker build -t blog -f ./buildimage/Dockerfile .
```

After completing building image, please create a working folder. Paste configuration file `config.json` and docker-compose file `docker-compose.yml` into the working folder. The configuration file contains log path, listen port and so on. Please read notes and edit these two files to meet your demand. Then run command:

```sh
docker-compose up -d
```

Congratulations, the deployment is completed!
