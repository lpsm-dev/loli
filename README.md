<div align="center">

<p align="center">
  <img alt="gif-header" src="https://i.pinimg.com/280x280_RS/d0/13/35/d01335f147c586e56829415e611f0ae7.jpg" width="350px" float="center"/>
</p>

<h2 align="center">✨ Loli Hunter CLI ✨</h2>

<div align="center">

[![Semantic Release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/lpmatos/loli)
[![Commitizen friendly](https://img.shields.io/badge/commitizen-friendly-brightgreen.svg)](https://github.com/lpmatos/loli)

</div>

---

<p align="center">
  <img alt="pipelines" src="https://images.ctfassets.net/em6l9zw4tzag/7j79Xx5NqeJlvGj6w98JUX/ecd7317698578b9efa1482154d0188f4/configuringpipeline.gif" width="450px" float="center"/>
</p>

<p align="center">
  ✨ Loli is a pretty CLI that find animes passing images - inspiration in <a href="https://github.com/irevenko/what-anime-cli">what-anime-cli</a> ✨
</p>

<p align="center">
  <a href="#getting-started">Getting Started</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#versioning">Versioning</a>
</p>

</div>

---

## ➤ Getting Started <a name = "getting-started"></a>

If you want contribute on this project, first you need to make a **git clone**:

>
> 1. git clone --depth 1 <https://github.com/lpmatos/loli.git> -b main
>

This will give you access to the code on your **local machine**.

## ➤ Description <a name = "description"></a>

This **CLI** is intended to be a code lab and best practices for creating a project ready to receive community builds, while introducing the basics for creating a **CLI** tool in **Go** and the standardization of conventions for the development workflow.

In this process, I gained experiences in the following topics regarding the Go language:

- ✔️ Discover internal Go packages like: `os`, `string` and `fmt`.
- ✔️ Discover the `github.com/spf13/cobra` CLI library.
- ✔️ Create commands and subcommands for your CLI.
- ✔️ Read flags and arguments from your commands and subcommands.
- ✔️ Discover the `github.com/spf13/viper` configuration library.
- ✔️ Read and write a configuration file.
- ✔️ Put the `snake` and the `viper` to work together.
- ✔️ Read environment variables.
- ✔️ Discover the `github.com/sirupsen/logrus` log library.
- ✔️ Use variable injection at build time.
- ✔️ Use conditional in compilation and build tags.
- ✔️ How to build CLI using Go

## ➤ Installation <a name = "installation"></a>

```bash
go install github.com/lpmatos/loli/cmd/loli

# if you cannot install directly, try following command,
# then input install command again
go get -x -u github.com/lpmatos/loli/cmd/loli
```

Or use a binary from [releases](https://github.com/lpmatos/loli/releases/latest).

## ➤ Development with docker


Estágios para buidar a imagem Docker:

<details><summary>🐋 Build</summary>
<p>

Docker commands to build your image:

```bash
docker image build -t <IMAGE_NAME> -f <PATH_DOCKERFILE> <PATH_CONTEXT_DOCKERFILE>
docker image build -t <IMAGE_NAME> . (This context)
```
</p>
</details>

<details><summary>🐋 Run</summary>
<p>

Docker commands to run a container with your image:

* **Linux** running:

```bash
docker container run -d -p <LOCAL_PORT:CONTAINER_PORT> <IMAGE_NAME> <COMMAND>
docker container run -it --rm --name <CONTAINER_NAME> -p <LOCAL_PORT:CONTAINER_PORT> <IMAGE_NAME> <COMMAND>
```

* **Windows** running:

```bash
winpty docker.exe container run -it --rm <IMAGE_NAME> <COMMAND>
```
</p>
</details>

## ➤ Usage <a name = "usage"></a>

Use examples liberally, and show the expected output if you can. It's helpful to have inline the smallest example of usage that you can demonstrate, while providing links to more sophisticated examples if they are too long to reasonably include in the README.

### Get Anime By Image File

```bash
loli file anime.jpg
```

### Get Anime By Image Link

```bash
loli link https://anime.com/image.png
```

## ➤ Visuals <a name = "visuals"></a>

Depending on what you are making, it can be a good idea to include screenshots or even a video (you'll frequently see GIFs rather than actual videos). Tools like ttygif can help, but check out Asciinema for a more sophisticated method.

### Get Anime By Image File

<p align="center">
  <img alt="logo" src="./docs/assets/find_by_file.PNG"/>
</p>

<p align="center">
  <img alt="logo" src="./docs/assets/find_by_file_pretty.PNG"/>
</p>

### Get Anime By Image Link

<p align="center">
  <img alt="logo" src="./docs/assets/find_by_link.PNG"/>
</p>

<p align="center">
  <img alt="logo" src="./docs/assets/find_by_link_pretty.PNG"/>
</p>

## ➤ Links <a name = "links"></a>

* https://soruly.github.io/trace.moe-api/#/

## ➤ Author <a name = "author"></a>

👤 Hey!! If you like this project or if you find some bugs feel free to contact me in my channels:

>
> * Email: lpsm-dev@protonmail.com
> * Linktree: https://linktr.ee/lpmatos
>

## ➤ License <a name = "license"></a>

This repository is released under the [Apache license](https://www.apache.org/licenses/LICENSE-2.0). In short, this means you are free to use this software in any personal, open-source or commercial projects. Attribution is optional but appreciated.

## ➤ Versioning <a name = "versioning"></a>

To check the change history, please access the [**CHANGELOG.md**](CHANGELOG.md) file.

## ➤ Project status <a name = "project-status"></a>

This project is currently undergoing a reorganization 👾.

## ➤ Show your support <a name = "show-your-support"></a>

<div align="center">

Give me a ⭐️ if this project helped you!

<p>
  <img alt="gif-header" src="https://www.icegif.com/wp-content/uploads/baby-yoda-bye-bye-icegif.gif" width="350px" float="center"/>
</p>

Made with 💜 by [me](https://github.com/lpmatos) :wave: inspired on [readme-md-generator](https://github.com/kefranabg/readme-md-generator)

</div>
