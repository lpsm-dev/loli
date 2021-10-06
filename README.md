<div align="center">

<p align="center">
  <img alt="azure" src="https://i.pinimg.com/280x280_RS/d0/13/35/d01335f147c586e56829415e611f0ae7.jpg" width="250px" float="center"/>
</p>


<h2 align="center">✨ Loli Hunter CLI ✨</h2>

<div align="center">

[![Semantic Release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://gitlab.com/nuageit/shared/auto-devops)
[![Commitizen friendly](https://img.shields.io/badge/commitizen-friendly-brightgreen.svg)](https://gitlab.com/nuageit/shared/auto-devops)

</div>

---

<p align="center">
  <img alt="pipelines" src="https://images.ctfassets.net/em6l9zw4tzag/7j79Xx5NqeJlvGj6w98JUX/ecd7317698578b9efa1482154d0188f4/configuringpipeline.gif" width="450px" float="center"/>
</p>

<p align="center">
  ✨ Loli is a pretty CLI that find animes passing scene images - inspiration in [what-anime-cli](https://github.com/irevenko/what-anime-cli) ✨
</p>

<p align="center">
  <a href="#getting-started">Getting Started</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#contributing">Contributing</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#versioning">Versioning</a>
</p>

</div>

---

## ➤ Getting Started <a name = "getting-started"></a>

To make it easy for you to get started with GitLab, here's a list of recommended next steps. Already a pro? Just edit this README.md and make it your own. Want to make it easy? [Use the template at the bottom](#editing-this-readme)!

If you want contribute on this project, first you need to make a **git clone**:

>
> 1. git clone --depth 1 <https://github.com/lpmatos/loli.git> -b main
>

This will give you access to the code on your **local machine**.

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
