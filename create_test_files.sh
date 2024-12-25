#!/bin/bash

# Проверка, указана ли целевая директория
TARGET_DIR=${1:-test_env}

# Создаем целевую директорию, если она не существует
mkdir -p "$TARGET_DIR"

# Создаем файлы с разными расширениями
echo "This is a text file." > "$TARGET_DIR/file1.txt"
echo "This is a CSV file." > "$TARGET_DIR/file2.csv"
echo "This is a JSON file." > "$TARGET_DIR/file3.json"
echo "This is a log file." > "$TARGET_DIR/file4.log"
echo "<html><body>Test HTML</body></html>" > "$TARGET_DIR/file5.html"

# Создаем бинарный файл
dd if=/dev/urandom bs=1024 count=1 of="$TARGET_DIR/file6.bin" &>/dev/null

# Создаем директории с разными правами доступа
mkdir -p "$TARGET_DIR/dir_read_only"
mkdir -p "$TARGET_DIR/dir_write_only"
mkdir -p "$TARGET_DIR/dir_no_access"

chmod 400 "$TARGET_DIR/dir_read_only"  # Только чтение
chmod 200 "$TARGET_DIR/dir_write_only" # Только запись
chmod 000 "$TARGET_DIR/dir_no_access"  # Нет доступа

# Вывод результата
echo "Тестовые файлы и директории созданы в '$TARGET_DIR':"
ls -l "$TARGET_DIR"
