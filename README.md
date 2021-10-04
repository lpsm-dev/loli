<p align="center">
  <img alt="logo" src="https://i.pinimg.com/280x280_RS/d0/13/35/d01335f147c586e56829415e611f0ae7.jpg" width="350px" float="center"/>
</p>

# Loli Hunter CLI

<p align="center">
  <a href="https://spdx.org/licenses/Apache-2.0.html" target="_blank">
    <img alt="License: Apache 2.0" src="https://img.shields.io/badge/License-Apache 2.0-yellow.svg" />
  </a>
  <a href="https://github.com/semantic-release/semantic-release">
    <img alt="semantic-release" src="https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg">
  </a>
  <a href="http://commitizen.github.io/cz-cli/">
    <img alt="Commitizen friendly" src="https://img.shields.io/badge/commitizen-friendly-brightgreen.svg">
  </a>
  <a href="https://www.codacy.com/gh/lpmatos/loli/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=lpmatos/loli&amp;utm_campaign=Badge_Grade">
    <img alt="Codacy Badge" src="https://app.codacy.com/project/badge/Grade/7d69caff8e2646a783681cc765948187">
  </a>
  <a href="https://github.com/lpmatos/loli/actions">
    <img alt="Latest Release" src="https://img.shields.io/github/v/release/lpmatos/Loli">
  </a>
</p>

>
> Loli é um CLI feito em Go para descobrir animes com base em imagens passadas como parâmetro
>

## ➤ Menu

<p align="left">
  <a href="#-description">Description</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-how-to-contribute">How to contribute</a>
</p>

## ➤ Começando

Se você quiser contribuir, primeiro você precisa fazer um **git clone** do repo:

>
> 1. git clone --depth 1 <https://github.com/lpmatos/loli.git> -b main
>

Isso lhe dará acesso ao código em sua **máquina local**.

## ➤ Objetivo

Esse **CLI** tem como objetivo ser um laboratório de código e boas práticas para criação de um projeto pronto para receber contruições da comunidade, enquanto introduz o básico para criação de uma ferramenta **CLI** em **Go** e a padronização de convenções para o workflow de desenvolvimento.

Nesta processo, obtive experiências nos seguintes tópicos no que tange a linguagem Go:

- ✔️ Descobrir os pacotes internos do Go, como: `os`, `string` e `fmt`.
- ✔️ Descobrir a biblioteca CLI `github.com/spf13/cobra`.
- ✔️ Criar commandos e subcomandos para seu CLI.
- ✔️ Ler flags e arguments a partir dos seus commandos e subcomandos.
- ✔️ Descobrir a biblioteca de configuração `github.com/spf13/viper`.
- ✔️ Ler e escrever um arquivo de configuração.
- ✔️ Colocar o `cobra` e o `viper` para trabalharem juntos.
- ✔️ Ler variáveis de ambiente.
- ✔️ Descobrir a biblioteca de log `github.com/sirupsen/logrus`.
- ✔️ Usar injeção de variável em tempo de build.
- ✔️ Usar conditional na compilação e build tags.

## ➤ Descrição

🗲 Uma simples descrição a ser definida.

## ➤ Instalação

```bash
go install github.com/lpmatos/Loli/cmd/loli

# if you cannot install directly, try following command,
# then input install command again
go get -u github.com/lpmatos/loli
```

## ➤ Desenvolvimento com Docker


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

## ➤ Links <a name = "links"></a>

* https://soruly.github.io/trace.moe-api/#/

## ➤ Visuals <a name = "visuals"></a>

Depending on what you are making, it can be a good idea to include screenshots or even a video (you'll frequently see GIFs rather than actual videos). Tools like ttygif can help, but check out Asciinema for a more sophisticated method.

### Search by File

<p align="center">
  <img alt="logo" src="./docs/assets/find_by_file.PNG"/>
</p>

<p align="center">
  <img alt="logo" src="./docs/assets/find_by_file_pretty.PNG"/>
</p>

### Search by Link

<p align="center">
  <img alt="logo" src="./docs/assets/find_by_link.PNG"/>
</p>

<p align="center">
  <img alt="logo" src="./docs/assets/find_by_link_pretty.PNG"/>
</p>

## ➤ Usage <a name = "usage"></a>

Use examples liberally, and show the expected output if you can. It's helpful to have inline the smallest example of usage that you can demonstrate, while providing links to more sophisticated examples if they are too long to reasonably include in the README.

<p align="center">
  <img alt="logo" src="./docs/assets/menu.PNG" float="center"/>
</p>

## ➤ Roadmap <a name = "roadmap"></a>

If you have ideas for releases in the future, it is a good idea to list them in the README.

## ➤ Versioning <a name = "versioning"></a>

To check the change history, please access the [**CHANGELOG.md**](CHANGELOG.md) file.

## ➤ Contributing <a name = "contributing"></a>

Contributions, issues and feature requests are welcome. Feel free to check issues page if you want to contribute. [Check the contributing guide](https://nuageit.atlassian.net/wiki/spaces/OPSNUAGE/pages/1995309068/Processo+de+contribui+o).

## ➤ Author

👤 **Lucca Pessoa**

Hey!! If you liked this project or found some bugs, feel free to contact me on my channels:

>
> * Email: luccapsm@protonmail.com
> * Website: [lpmatos](https://github.com/lpmatos)
> * Github: [@lpmatos](https://github.com/lpmatos)
> * GitLab: [@lpmatos](https://gitlab.com/lpmatos)
> * LinkedIn: [@luccapessoa](https://www.linkedin.com/in/luccapessoa/)
>

## ➤ Troubleshooting <a name = "troubleshooting"></a>

If you have any problems, please contact **me**.

## ➤ Project status <a name = "project-status"></a>

If you have run out of energy or time for your project, put a note at the top of the README saying that development has slowed down or stopped completely. Someone may choose to fork your project or volunteer to step in as a maintainer or owner, allowing your project to keep going. You can also make an explicit request for maintainers.

## ➤ Show your support <a name = "show-your-support"></a>

Give a ⭐️ if this project helped you!

---

<div align="center">

Made with 💜 by **me** :wave: inspired on [readme-md-generator](https://github.com/kefranabg/readme-md-generator)

</div>
