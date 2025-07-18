# 🌤️ GoWeatherTracker

Uma API REST em Go para monitoramento de clima de cidades, integrando chamadas a APIs externas, PostgreSQL e Docker Compose. O sistema utiliza `goroutines` e `channels` para atualizar periodicamente as informações climáticas.

---

## ✨ Funcionalidades

- `POST /cities` – Cadastrar cidade para monitoramento
- `GET /cities` – Listar todas as cidades cadastradas
- `GET /cities/{id}/weather` – Mostrar o clima mais recente de uma cidade
- `DELETE /cities/{id}` – Remover cidade do monitoramento
- Atualização automática do clima via goroutines a cada X minutos

---

## 🧪 Tecnologias e Conceitos Envolvidos

| Recurso                | Uso no Projeto                               |
| ---------------------- | -------------------------------------------- |
| **Go**                 | Backend da aplicação                         |
| **Gin ou Fiber**       | Framework web leve e rápido                  |
| **PostgreSQL**         | Armazenamento das cidades e dados climáticos |
| **GORM ou pgx**        | ORM ou driver SQL nativo                     |
| **OpenWeatherMap API** | Fonte de dados climáticos em tempo real      |
| **Goroutines**         | Atualização concorrente das cidades          |
| **Channels**           | Coordenação e comunicação entre workers      |
| **Docker Compose**     | Orquestração de app, banco e pgAdmin         |

---

## 🧱 Estrutura de Pastas Sugerida

go-weather-tracker/
├── cmd/ # Arquivo main.go
├── internal/
│ ├── api/ # Handlers REST
│ ├── weather/ # Lógica de chamadas à API externa
│ ├── updater/ # Goroutines + channels
│ ├── db/ # Modelos e acesso ao banco
├── docker/
│ ├── Dockerfile
│ └── docker-compose.yml
└── go.mod

---

## 🚀 Como Rodar o Projeto

### Pré-requisitos

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- Conta na [OpenWeatherMap](https://openweathermap.org/api) para obter sua API Key

### 1. Clone o projeto

```bash
git clone https://github.com/seuusuario/go-weather-tracker.git
cd go-weather-tracker
```
