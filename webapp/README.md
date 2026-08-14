# Webapp DevBook

Módulo Go responsável pela interface web do DevBook. É uma aplicação renderizada no servidor (templates `html/template` + jQuery e Bootstrap no cliente) que **não acessa o banco diretamente**: cada ação do usuário é traduzida em uma chamada HTTP para a API (`../api`).

## Stack

- Go 1.14
- [`gorilla/mux`](https://github.com/gorilla/mux) — roteamento HTTP
- [`gorilla/securecookie`](https://github.com/gorilla/securecookie) — cookies criptografados e autenticados
- [`joho/godotenv`](https://github.com/joho/godotenv) — leitura do `.env`
- `html/template` — renderização das views
- jQuery, Bootstrap, SweetAlert2 e Font Awesome (via CDN) no lado cliente

## Como executar

A API precisa estar rodando antes da webapp.

```bash
cp example.env .env   # preencha as variáveis abaixo
go run main.go
```

Variáveis esperadas no `.env`:

| Variável    | Descrição                                                                      |
|-------------|--------------------------------------------------------------------------------|
| `API_URL`   | URL base da API (ex.: `http://localhost:9000`)                                 |
| `APP_PORT`  | Porta HTTP onde a webapp vai escutar                                           |
| `HASH_KEY`  | Chave usada por `securecookie` para autenticar o cookie (HMAC)                 |
| `BLOCK_KEY` | Chave AES usada por `securecookie` para criptografar o cookie (16, 24 ou 32 B) |

## Estrutura interna

```
webapp/
├── main.go                # bootstrap: config, cookies, templates e router
├── example.env            # template do .env
├── assets/                # estáticos servidos em /assets/
│   ├── css/
│   └── js/                # scripts jQuery por página (login, register, home, user, ...)
├── views/
│   ├── *.html             # páginas (login, register, home, profile, users, ...)
│   └── templates/         # parciais reutilizáveis (header, scripts, posts)
└── src/
    ├── config/            # leitura do .env (APIURL, Port, HashKey, BlockKey)
    ├── cookies/           # wrapper sobre securecookie (Setup, Save, Read, Delete)
    ├── utils/             # LoadTemplates() + ExecuteTemplate()
    ├── models/            # structs usados na renderização (User, Post, Authentication)
    ├── requests/          # MakeAuthenticatedRequest() — http.Client com Bearer
    ├── responses/         # helpers JSON e tradução de status code da API
    ├── middlewares/       # Logger + Authenticate (verifica cookie antes de servir páginas)
    └── router/
        ├── router.go
        └── routes/        # tabela de rotas para login, logout, home, users, posts
```

Os nomes de pacotes, structs, funções e rotas seguem o mesmo padrão em inglês usado em `../api` e `../collection`.

## Fluxo de uma requisição

1. `main.go` chama `config.LoadEnvironmentVariables()`, `cookies.Setup()` (instancia o `securecookie` com `HashKey`/`BlockKey`), `utils.LoadTemplates()` (faz `ParseGlob` de `views/*.html` e `views/templates/*.html`) e finalmente `router.New()`.
2. Cada rota declarada em `src/router/routes/` indica se a página exige usuário logado (`IsAuthRequired`). O middleware `Authenticate` lê o cookie `auth`, valida com `securecookie.Decode` e redireciona para `/login` quando não há sessão válida.
3. Controllers se dividem em dois grupos:
   - **`pages.go`** — apenas renderiza HTML. Pode buscar dados na API antes (ex.: feed da home) e passa o resultado para o template.
   - **`login.go`, `logout.go`, `users.go`, `posts.go`** — recebem submissões/AJAX, montam o JSON correspondente, chamam a API via `requests.MakeAuthenticatedRequest` e devolvem JSON para o front (que dispara `Swal.fire` em caso de erro).
4. Respostas de erro vindas da API são propagadas com `responses.HandleStatusCodeError`, mantendo o status original.

## Autenticação no front

O login segue este caminho:

1. `login.js` faz `POST /login` para a webapp com `email` e `password` (form-encoded).
2. `controllers.Login` reembala em JSON e chama `POST {API_URL}/login`.
3. A API devolve `{id, token}`. A webapp grava esses dois valores num cookie criptografado (`securecookie`) chamado `auth`.
4. Em qualquer requisição posterior à API, `requests.MakeAuthenticatedRequest` lê o cookie, recupera o token e injeta o header `Authorization: Bearer <token>`.
5. `logout.go` apaga o cookie e redireciona para `/login`.

## Detalhes de implementação que valem a pena conhecer

- **Templates parseados uma vez no boot.** `utils.LoadTemplates` mantém as templates em uma variável de pacote. `ExecuteTemplate` apenas faz `templates.ExecuteTemplate(w, name, data)`. Em desenvolvimento, isso significa que mudanças em `views/` exigem reiniciar o processo.
- **Parciais via `{{ define }}`.** As parciais em `views/templates/` (`header.html`, `scripts.html`, `posts.html`) são incluídas nas páginas com `{{ template "nome" . }}`. O escopo `.` é importante para que listas de publicações tenham acesso a `UserID`.
- **Cookies como única fonte de sessão.** A webapp não mantém estado em memória nem em banco — toda a sessão é o cookie criptografado. Trocar `HASH_KEY`/`BLOCK_KEY` invalida todas as sessões existentes.
- **Comunicação client → webapp → API.** Algumas rotas existem em duas camadas (ex.: a webapp expõe `POST /posts/:id/upvote` para o front, que internamente chama o mesmo endpoint na API REST). Isso evita que o token JWT precise existir no JavaScript do navegador.
- **Erros propagados com fidelidade.** `responses.HandleStatusCodeError` lê o JSON de erro da API (`{"error": "..."}`) e devolve com o mesmo status code, então a UX (mensagens do `Swal`) reflete o que a API decidiu, não uma camada genérica da webapp.
