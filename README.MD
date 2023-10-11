# Discipline API
## Краткое описание
Это микросервис, реализующий CRUD для учебных дисциплин.
Основные операции доступны по REST API, некоторые через Apache Kafka
## Запуск
Команда
`make init` осуществляет сборку docker compose, прокидывает туда переменные среды
## Пересборка контейнера с go
Команда `make rebuild` позволяет пересобрать приложение, чтобы проверить внесённые изменения
## Генерация OpenAPI спецификации
На проекте используется `swaggo` для _автоматической_ генерации спецификации по `godoc`, руками json'ы писать не нужно.

После добавления к нужному handler'у документации, нужно запустить команду `make generate-openapi`

Документация будет доступна по роуту `/api/swagger/`