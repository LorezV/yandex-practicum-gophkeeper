## Серверная часть

### Архитектура

Серверная часть состоит из самого приложения и хранилища Postgres, кэширование запросов аутентификации происходит в памяти.

### Запуск

Для бепасной работы необходима генерация публичных и приватных ключей для шифрования токенов пользователей.
```
cp .env.sample .env
make compose-up
```

Описание API серверной части в формате swagger - `http://localhost:8080/swagger/index.html`

## Клиентская часть

Для клиента разработано cli-приложение, которое локально сохраняет данные пользователя в зашифрованном по паролю виде.

Для работы клиента необходимо наличие конфигурационного `./config/client/config.yml` файла

Перечень доступных пользователю команд:

```
Usage:
  gophkeeper-client [flags]
  gophkeeper-client [command]

Available Commands:
  addcard     Add card
  addlogin    Add login
  addnote     Add note
  completion  Generate the autocompletion script for the specified shell
  delcard     Delete user card by id
  dellogin    Delete user login by id
  delnote     Delete user note by id
  getcard     Show user card by id
  getlogin    Show user login by id
  getnote     Show user note by id
  help        Help about any command
  init        Init local storage
  login       Login user to the service
  logout      Logout user
  register    Register user to the service
  showvault   Show user vault
  sync        Sync user`s data

Flags:
  -h, --help   help for gophkeeper-client
```
