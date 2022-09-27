<div align="center">

<img alt="gif-header" src="https://c.tenor.com/PX2XATCduFcAAAAC/loli.gif" width="225"/>

<h2>✨ Loli CLI ✨</h2>

[![Semantic Release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/ci-monk/loli)
[![Commitizen friendly](https://img.shields.io/badge/commitizen-friendly-brightgreen.svg)](https://github.com/ci-monk/loli)

---

<img alt="pipelines" src="https://i.pinimg.com/originals/ce/26/14/ce2614ef4c70f04a2c578f972308f5b6.gif" width="250"/>

<p>✨ Loli is a pretty CLI that search animes passing images or links using trace.moe API ✨</p>

<p>
  <a href="#getting-started">Getting Started</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#versioning">Versioning</a>
</p>

</div>

---

## ➤ Getting Started <a name = "getting-started"></a>

If you want contribute on this project, first you need to make a **git clone**:

```bash
git clone --depth 1 https://github.com/ci-monk/loli.git -b main
```

This will give you access to the code on your **local machine**.

## ➤ Description <a name = "description"></a>

This **CLI** is intended to be a code lab and best practices for creating a project ready to receive community builds, while introducing the basics for creating a **CLI** tool in **Go** and the standardization of conventions for the development workflow.

## ➤ Installation <a name = "installation"></a>

with `go`:

```bash
go install github.com/ci-monk/loli/cmd/loli

# if you cannot install directly, try following command,
# then input install command again
go get -x -u github.com/ci-monk/loli/cmd/loli
```

with `brew`:

```bash
brew tap ci-monk/tools
brew install loli
```

or use a binary from [releases](https://github.com/ci-monk/loli/releases/latest).

## ➤ Trace.moe <a name = "trace.moe"></a>

Example trace.moe response:

```json
{
  "frameCount": 745506,
  "error": "",
  "result": [
    {
      "anilist": {
        "id": 99939,
        "idMal": 34658,
        "title": { "native": "ネコぱらOVA", "romaji": "Nekopara OVA", "english": null },
        "synonyms": ["Neko Para OVA"],
        "isAdult": false
      },
      "filename": "Nekopara - OVA (BD 1280x720 x264 AAC).mp4",
      "episode": null,
      "from": 97.75,
      "to": 98.92,
      "similarity": 0.9440424588727485,
      "video": "https://media.trace.moe/video/99939/Nekopara%20-%20OVA%20(BD%201280x720%20x264%20AAC).mp4?t=98.33500000000001&token=xxxxxxxxxxxxxx",
      "image": "https://media.trace.moe/image/99939/Nekopara%20-%20OVA%20(BD%201280x720%20x264%20AAC).mp4?t=98.33500000000001&token=xxxxxxxxxxxxxx"
    }
  ]
}
```

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

## ➤ Demo <a name = "demo"></a>

https://user-images.githubusercontent.com/58797390/148670421-fce972b2-8c13-4358-95ae-cd5ee6ac4d69.mov

## ➤ Learning <a name = "learning"></a>

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

## ➤ Links <a name = "links"></a>

- https://soruly.github.io/trace.moe-api/#/
- https://img.olhardigital.com.br/wp-content/uploads/2021/07/Naruto-Classico-e-Naruto-Shippuden-fillers.jpg
- https://images.plurk.com/32B15UXxymfSMwKGTObY5e.jpg

## ➤ Versioning <a name = "versioning"></a>

To check the change history, please access the [**CHANGELOG.md**](CHANGELOG.md) file.

## ➤ Show your support <a name = "show-your-support"></a>

<div align="center">

Give me a ⭐️ if this project helped you!

<img alt="gif-header" src="https://www.icegif.com/wp-content/uploads/baby-yoda-bye-bye-icegif.gif" width="225"/>

Made with 💜 by [me](https://github.com/ci-monk) :wave: inspired on [readme-md-generator](https://github.com/kefranabg/readme-md-generator)

</div>
