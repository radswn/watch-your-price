# Database

## Zarządzanie bazą danych

### Wymagania

- docker
- docker-compose

### Uruchomienie

Aby uruchomić bazę danych należy mieć zainstalowanego dockera i narzędzie docker compose. W folderze zawierającym plik
docker-compose.yml uruchom polecenie

```bash
$ docker-compose up -d
```

Uruchomi to bazę danych na porcie 3306. Następnie
należy [migrować bazę do najnowszej wersji](#migracja-do-najnowszej-wersji).

### Wyłączenie

```bash
$ docker-compose down
```

## Migracja bazy danych

Migracja przebiega za pomocą skryptu ```db-migrate.sh```. Po więcej informacji użyj komendy

```bash
$ ./db-migrate.sh
```

### Migracja do najnowszej wersji

```bash
$ ./db-migrate.sh update
```

### Tworzenie nowej wersji

1. Utwórz pliki migracyjne

   ```bash
   $ ./db-migrate.sh create <nazwa>
   ```
2. Uzupełnić utworzony plik z "up" w nazwie, dopisując tam zapytanie sql, które przeprowadzi bazę danych z wersji
   aktualnej na następną.


3. Uzupełnić utworzony plik z "down" w nazwie, dopisując tam zapytanie sql, które przeprowadzi bazę z następnej na
   aktualną.


4. Aby zaaplikować nowe zmiany

   ```bash 
   $ ./db-migrate.sh update
   ```