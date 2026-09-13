# types

Общие типы и callback-контракты для компонентов Sudzekai Web OS.

## Установка

```bash
go get github.com/sudzekai-web-os/types@latest
```

## Результаты

`HandlerResult` описывает результат HTTP-обработчика:

```go
types.HandlerResult{
	Data:       response,
	StatusCode: http.StatusOK,
}
```

`CommandResult` содержит `Stdout`, `Stderr` и `Error` после выполнения внешней команды.

## HTTP-контракты

- `HandlerFunc` — функция `func(*http.Request) HandlerResult`.
- `Middleware` — стандартное HTTP middleware.
- `ResultFilter` — функция формирования ответа из `HandlerResult`.
- `JwtMiddleware` — адаптер защиты обработчика с учетом списка ролей.
- `ProtectedHandlerFunc` — обработчик и разрешенные роли; создается через `NewProtectedHandler`.

## Уровни логирования

`LogLevel` содержит уровни `Debug`, `Information`, `Warning`, `Error` и `Critical`. Значение `None` используется для отсутствия уровня. Метод `String` возвращает текстовое представление уровня.
