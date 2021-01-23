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
- ✔️ Criar comandos e subcomandos para seu CLI.
- ✔️ Ler flags e argumentos a partir dos seus comandos e subcomandos.
- ✔️ Descobrir a biblioteca de configuração `github.com/spf13/viper`.
- ✔️ Ler e escrever um arquivo de configuração.
- ✔️ Colocar o `cobra` e o `viper` para trabalharem juntos.
- ✔️ Ler variáveis de ambiente.
- ✔️ Descobrir a biblioteca de log `github.com/sirupsen/logrus`.
- ✔️ Usar injeção de variável em tempo de build.
- ✔️ Usar condicional na compilação e build tags.

## ➤ Descrição

Uma simples descrição a ser definida.

### [🗲 Start the codelab](https://nlepage.github.io/catption/codelab)

## ➤ Instalação

```bash
go get github.com/lpmatos/loli
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

## ➤ Project organization features

- Default gitignore and editorconfig.
- [Semantic Versioning](https://semver.org/)
- [GitLeaks](https://github.com/zricethezav/gitleaks) file - Scan for secrets using regex and entropy.
- [Semantic Release](https://github.com/semantic-release/semantic-release) + Plugins configuration
- NPM modules automation.
  - [Commitlint](https://github.com/conventional-changelog/commitlint) using [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/).
  - Git Hooks with [Husky](https://github.com/typicode/husky).

## ➤ How to contribute

>
> 1. Make a **Fork** or Create a **Feature Branch**.
> 2. Follow the project organization.
> 3. Add the file to the appropriate level folder - If the folder does not exist, create according to the standard.
> 4. Make the **Commit**.
> 5. Open a **Merge Request**.
> 6. Wait for your merge request to be accepted.. 🚀
>

**Remember**: There is no bad code, there are different views/versions of solving the same problem. 😊

## ➤ Add to git and push

📝 You must send your work to [**GitHub**](https://github.com/lpmatos/loli) repo after your changes:

```bash
git add -f .
git commit -m "chore(initial): include config files"
git push -u origin main
```

## ➤ Versionamento

🚨 No momento não temos um [**CHANGELOG.md**](CHANGELOG.md) gerado.

**Para obter mais informações como é o processo de controle de versionamento nesse projeto, acesse [RELEASE.md](RELEASE.md), lá temos todos os requisitos e especificações do que deve ser feito para a geração de uma release, tag e changelog.**

## ➤ Autor

👤 **Lucca Pessoa**

Ei!! Se você gostou deste projeto ou se encontrou alguns bugs, sinta-se à vontade para me contatar nos meus canais:

>
> * Email: luccapsm@gmail.com
> * Website: [lpmatos](https://github.com/lpmatos)
> * Github: [@lpmatos](https://github.com/lpmatos)
> * GitLab: [@lpmatos](https://gitlab.com/lpmatos)
> * LinkedIn: [@luccapessoa](https://www.linkedin.com/in/luccapessoa/)
>

## ➤ Troubleshooting

🚨 Esse repositório é um projeto pessoal criado para fins de estudo/demonstração e para simplificar minha vida profissional, podendo ou não ser adequado para o seu projeto!

## ➤ Licença

🔖 Aplicação distribuída sobre **Apache License**. Olhe em [LICENSE](LICENSE) para mais informações.

## ➤ Contribuindo

🔖 Para saber mais sobre o processo de contribuição acesse [CONTRIBUTING.md](CONTRIBUTING.md).

## ➤ Mostre seu suporte!

Dê uma ⭐️ se esse projeto te ajudou de alguma forma!

---

Esse [README](README.md) foi gerado com ❤️ por [mim](https://github.com/lpmatos) e é inspirado no 🎉 [readme-md-generator](https://github.com/kefranabg/readme-md-generator).
