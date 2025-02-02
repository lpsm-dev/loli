<!-- BEGIN_DOCS -->
<div align="center">

<a name="readme-top"></a>

**Loli CLI**

<img alt="pipelines" src="https://i.pinimg.com/originals/ce/26/14/ce2614ef4c70f04a2c578f972308f5b6.gif" width="250"/>

Hello Humans 👽! Loli is a pretty CLI that search animes passing images or links

[![Commitizen friendly](https://img.shields.io/badge/commitizen-friendly-brightgreen.svg)](https://www.conventionalcommits.org/en/v1.0.0/) [![Semantic Release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://semantic-release.gitbook.io/semantic-release/usage/configuration) [![Built with Devbox](https://jetpack.io/img/devbox/shield_galaxy.svg)](https://jetpack.io/devbox/docs/contributor-quickstart/)

</div>

# Visão Geral

Este projeto é um **CLI** que busca animes a partir de imagens ou links. Ele utiliza a API do [trace.moe](https://soruly.github.io/trace.moe-api/#/) para realizar a busca.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

# Como Instalar?

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

<p align="right">(<a href="#readme-top">back to top</a>)</p>

# Usage

## Get anime with a file

```bash
loli search file anime.jpg
```

## Get anime with a link

```bash
loli search link https://anime.com/image.png
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

# Trace.moe

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
        "title": {
          "native": "ネコぱらOVA",
          "romaji": "Nekopara OVA",
          "english": null
        },
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

<p align="right">(<a href="#readme-top">back to top</a>)</p>

# Demo

https://user-images.githubusercontent.com/58797390/192595643-a27003a5-d0ba-4abf-b8bb-19f449398190.mov

<p align="right">(<a href="#readme-top">back to top</a>)</p>

# Referências

Links relevantes para essa documentação:

- https://soruly.github.io/trace.moe-api/#/
- https://img.olhardigital.com.br/wp-content/uploads/2021/07/Naruto-Classico-e-Naruto-Shippuden-fillers.jpg
- https://images.plurk.com/32B15UXxymfSMwKGTObY5e.jpg

<p align="right">(<a href="#readme-top">back to top</a>)</p>

# Versionamento

Para verificar o histórico de mudanças, acesse o arquivo [**CHANGELOG.md**](CHANGELOG.md).

<p align="right">(<a href="#readme-top">back to top</a>)</p>

# Troubleshooting

Se você tiver algum problema ou queria contribuir, abra uma [issue](https://github.com/lpsm-dev/loli/issues/new/choose) nesse projeto.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

# Show your support

<div align="center">

Dê uma ⭐️ para este projeto se ele te ajudou!

<img alt="gif-footer" src="https://github.com/lpsm-dev/lpsm-dev/blob/0062b174ec9877e6dfc78817f314b4a0690f63ff/.github/assets/yoda.gif" width="225"/>

<br>
<br>

Feito com 💜 pelo **Time de DevOps** :wave: inspirado no [readme-md-generator](https://github.com/kefranabg/readme-md-generator)

</div>

<p align="right">(<a href="#readme-top">back to top</a>)</p>
<!-- END_DOCS -->
