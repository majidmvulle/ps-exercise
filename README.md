# ps-exercise

## 📖 Project description

A monorepo containing service-a (Expressjs service) and service-b (Go/gin service). service-a has API endpoints that 
fetch data from service-b which fetches strings from a Postgres database

## 🔧 Installation (local)
- `git clone` the monorepo.
- See the [Makefile](Makefile) for build steps for the whole monorepo.
- For service-a build commands, see its [Makefile](./service-a-expressjs/Makefile)
- For service-b build commands, including how to run database migrations and seeding, see its [Makefile](./service-b-go/Makefile)

## ✈️ Usage
- For service-a API documentation, see its [OpenAPI documentation](./service-a-expressjs/docs/openapi/api.yaml)
- For service-b API documentation, see its [OpenAPI documentation](./service-b-go/docs/openapi/api.yaml)

## ⚙️ CI/CD
- See [GitHub actions](.github) for various workflows and actions

## 💻 Developers

- Majid Mvulle <majidmvulle@github.com>


## 🪪 License
 
None
