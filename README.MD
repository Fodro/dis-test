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
## Работа с kafka
С кафкой реализована аналогичная CRUD-схема.
В key сообщения нужно писать необходмое действие:
``
"create", "delete", "update", "addPrerequisite", "deletePrerequisite"
``
### Create
В value нужно передать json. Схема такая же, как у запроса на создание в Ресте
### Delete
В value нужно передать string, содержащий id дисциплины для удаления
### Update
В value нужно передать аналогичное Create значение, но добавив `id: int`
### AddPrerequisite
В value нужно передать json вида:
``{"discipline_id": int, "prerequisite_id": int}``
### DeletePrerequisite
В value нужно передать json вида:
``{"discipline_id": int, "prerequisite_id": int}``