package util

import (
	"log"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var F = loadEnvFile() // запуск при загрузке пакета
// путь к  файлу .env
func loadEnvFile() string {

	path, err := GetExecPath()
	if err != nil {
		log.Println(err)
	}

	pathEnv := filepath.Join(path, "..", ".env")

	if err := godotenv.Load(pathEnv); err != nil {
		log.Print("Не найден .env файл")
	}
	color.Green("Путь к .env: %s", pathEnv)
	return ""
}

// Получим каталог exe-шника
func GetExecPath() (string, error) {
	// путь к исполняемому файлу
	path, err := os.Executable()
	if err != nil {
		return "", err
	}
	// проверка на символическую ссылку
	path, err = filepath.EvalSymlinks(path)
	if err != nil {
		return "", err
	}
	return path, nil
}
