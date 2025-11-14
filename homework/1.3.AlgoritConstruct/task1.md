![Git - система контроля версий](https://git-scm.com/images/logo@2x.png)

<span style="
    background: linear-gradient(90deg, #ff0000, #ff7f00, #ffff00, #00ff00, #0000ff, #4b0082, #8b00ff);
    background-clip: text;
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    font-weight: bold;
">__*Git*__ - это распределённая система контроля версий, которая позволяет отслеживать изменения в файлах, сотрудничать с другими разработчиками и управлять историей проекта.</span>

## Инициализация репозитория

### Репозиторий (repository)
\- это хранилище вашего проекта вместе с историей изменений.

```bash
# Создание нового репозитория
git init

# Клонирование существующего репозитория
git clone <url>
```

### Коммит (commit)
Коммит в git репозитории хранит снимок состояния всех файлов (проекта) в директории.

```bash
# Создание коммита
git commit -m "Text of the comment"

# Просмотр истории коммитов
git log
```

### Ветки (branches)
\- это независимая линия разработки. Ветки позволяют разрабатывать функциональность изолированно.

```bash
# Создание новой ветки
git branch <name_branch>

# Переключение на ветку
git checkout <name_branch>
# или
git switch <name_branch>

# Создание и переключение на новую ветку
git checkout -b <name_branch>
# или
git switch -c <name_branch>

# Просмотр всех веток
git branch
```

### HEAD
\- это указатель на текущий коммит, с которым вы работаете. Обычно HEAD указывает на последний коммит в текущей ветке.

```bash
# Просмотр, на что указывает HEAD
git rev-parse HEAD

# Переключение HEAD на определенный коммит
git checkout <hash_commit>
```

## Базовые команды в Git

### Добавление файлов в индекс

```bash
# Добавление файла в индекс
git add <file>

# Добавление всех измененных файлов
git add .

# Интерактивное добавление
git add -i
```

### Проверка статуса

```bash
# Проверка статуса репозитория
git status

# Краткий статус
git status -s
```

### Просмотр изменений

```bash
# Просмотр неиндексированных изменений
git diff

# Просмотр индексированных изменений
git diff --staged
```

### Отмена изменений

```bash
# Отмена изменений в рабочей директории
git checkout -- <file>
# или
git restore <file>

# Отмена индексации
git reset HEAD <file>
# или
git restore --staged <file>
```

## Работа с коммитами

### Изменение последнего коммита (amend)
Amend позволяет изменить последний коммит, добавив новые изменения или изменив сообщение.

```bash
# Изменение сообщения последнего коммита
git commit --amend -m "New message"

# Добавление изменений в последний коммит
git add <file>
git commit --amend --no-edit
```

### Получение SHA коммита
SHA (Secure Hash Algorithm) - это уникальный идентификатор коммита.

```bash
# Получение полного SHA текущего коммита
git rev-parse HEAD

# Получение короткого SHA
git rev-parse --short HEAD

# Просмотр SHA коммитов в истории
git log --oneline
```

### Сброс (reset)
Reset позволяет отменить коммиты или изменения в индексе.

```bash
# Мягкий сброс (сохраняет изменения в рабочей директории)
git reset --soft HEAD~1

# Смешанный сброс (сохраняет изменения, но не в индексе)
git reset HEAD~1

# Жесткий сброс (удаляет все изменения)
git reset --hard HEAD~1
```

### Поиск автора кода (blame)
Blame позволяет узнать, кто и когда внёс изменения в определённые строки файла.

```bash
# Просмотр авторов строк в файле
git blame <file>

# С ограничением по строкам
git blame -L 10,20 <file>
```

### Поиск в коде (grep)
Grep позволяет искать строки в файлах репозитория.

```bash
# Поиск строки во всех файлах
git grep "searching_string"

# Поиск с учетом контекста
git grep -n "searching_string"

# Поиск только в определенных файлах
git grep "searching_string" -- "*.js"
```

## Работа с ветками

### Слияние (merge)
Merge объединяет изменения из одной ветки в другую.

```bash
# Слияние ветки в текущую
git merge <name_branch>

# Слияние с созданием коммита слияния (no fast-forward)
git merge --no-ff <name_branch>

# Отмена слияния
git merge --abort
```

#### Fast-forward (FF) vs No Fast-forward (no-FF)
- **Fast-forward (FF)**: если возможно, Git просто перемещает указатель текущей ветки вперед. Не создаётся отдельный коммит слияния.
- **No Fast-forward (no-FF)**: всегда создаётся коммит слияния, даже если можно выполнить fast-forward. Это сохраняет информацию о существовании ветки в истории.

### Перебазирование (rebase)
Rebase переносит коммиты из одной ветки на конец другой, создавая линейную историю.

```bash
# Перебазирование текущей ветки на другую
git rebase <name_branch>

# Отмена перебазирования
git rebase --abort

# Продолжение перебазирования после разрешения конфликтов
git rebase --continue
```

#### Интерактивный rebase
Интерактивный rebase позволяет изменять, объединять, удалять и переупорядочивать коммиты.

```bash
# Интерактивный rebase последних N коммитов
git rebase -i HEAD~N

# Интерактивный rebase от определенного коммита
git rebase -i <hash_commit>^
```

В открывшемся редакторе вы можете:
| Команда | Описание |
|---------|----------|
| `pick` | оставить коммит как есть |
| `reword` | изменить сообщение коммита |
| `edit` | остановиться для изменения коммита |
| `squash` | объединить с предыдущим коммитом |
| `fixup` | объединить с предыдущим коммитом, отбросив сообщение |
| `drop` | удалить коммит |

### Объединение коммитов (squash)
Squash позволяет объединить несколько коммитов в один.

```bash
# Объединение последних N коммитов через интерактивный rebase
git rebase -i HEAD~N
# Затем замените "pick" на "squash" или "fixup" для коммитов, которые нужно объединить
```

## Удаленные репозитории

### Работа с удаленными репозиториями

```bash
# Просмотр удаленных репозиториев
git remote -v

# Добавление удаленного репозитория
git remote add <name> <url>

# Получение изменений без слияния
git fetch <name_remote_repository>

# Получение изменений и слияние
git pull <name_remote_repository> <name_branch>

# Отправка изменений
git push <name_remote_repository> <name_branch>
```

### Bundle
Bundle позволяет создать архив с коммитами, который можно передать другим разработчикам без использования удаленного репозитория.

```bash
# Создание бандла со всеми ветками
git bundle create repo.bundle --all

# Создание бандла с определенной веткой
git bundle create repo.bundle <name_branch>

# Клонирование из бандла
git clone repo.bundle -b <name_branch> <directory>
```

## Git Hooks (хуки)
Хуки - это скрипты, которые Git запускает перед или после определенных событий.

Основные хуки:
- `pre-commit`: запускается перед созданием коммита
- `prepare-commit-msg`: запускается перед открытием редактора сообщения коммита
- `commit-msg`: запускается для проверки сообщения коммита
- `post-commit`: запускается после создания коммита
- `pre-push`: запускается перед отправкой изменений
- `post-checkout`: запускается после переключения веток

```bash
# Хуки хранятся в директории .git/hooks/
# Пример создания pre-commit хука
echo '#!/bin/sh
# Проверка стиля кода
npm run lint' > .git/hooks/pre-commit
chmod +x .git/hooks/pre-commit
```

## Дополнительные команды

### Stash
Stash позволяет временно сохранить изменения, не создавая коммит.

```bash
# Сохранение изменений
git stash

# Применение последнего stash
git stash apply

# Удаление последнего stash после применения
git stash pop

# Просмотр списка stash
git stash list

# Удаление stash
git stash drop stash@{n}
```

### Cherry-pick
Cherry-pick позволяет применить изменения из одного коммита к текущей ветке.

```bash
# Применение коммита к текущей ветке
git cherry-pick <hash_commit>

# Применение нескольких коммитов
git cherry-pick <hash1> <hash2>
```

### Reflog
Reflog хранит историю изменений указателей (HEAD, ветки) и позволяет восстановить потерянные коммиты.

```bash
# Просмотр истории изменений HEAD
git reflog

# Восстановление до определенного состояния
git reset --hard HEAD@{n}
```

### Bisect
Bisect помогает найти коммит, в котором появилась ошибка, с помощью бинарного поиска.

```bash
# Начало бинарного поиска
git bisect start

# Отметка текущего коммита как плохого
git bisect bad

# Отметка определенного коммита как хорошего
git bisect good <hash_commit>

# Отметка текущего коммита как хорошего/плохого в процессе поиска
git bisect good
git bisect bad

# Завершение поиска
git bisect reset
```

## Конфигурация Git

```bash
# Установка имени пользователя
git config --global user.name "Name of File"

# Установка email
git config --global user.email "email@example.com"

# Установка редактора
git config --global core.editor "vim"

# Просмотр конфигурации
git config --list
```

## Алиасы (сокращения команд)

```bash
# Создание алиаса
git config --global alias.co checkout
git config --global alias.br branch
git config --global alias.ci commit
git config --global alias.st status

# Использование алиаса
git co -b new-branch
```