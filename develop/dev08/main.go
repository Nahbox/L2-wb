package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	// Получаем путь к исполняемому файлу
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Ошибка при получении пути к исполняемому файлу:", err)
		return
	}

	exeDir := filepath.Dir(exePath)

	// Добавляем путь к директории с исполняемым файлом в переменную окружения PATH
	os.Setenv("PATH", exeDir+string(filepath.Separator)+os.Getenv("PATH"))

	for {
		fmt.Print("> ")
		input, err := readInput()
		if err != nil {
			fmt.Println("Ошибка чтения ввода:", err)
			continue
		}

		if strings.TrimSpace(input) == "\\quit" {
			fmt.Println("Выход из шелла.")
			break
		}

		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "cd":
			if len(args) < 2 {
				fmt.Println("Необходимо указать аргумент для команды cd.")
				continue
			}
			err := os.Chdir(args[1])
			if err != nil {
				fmt.Println("Ошибка при изменении директории:", err)
			}
		case "pwd":
			cwd, err := os.Getwd()
			if err != nil {
				fmt.Println("Ошибка при получении текущей директории:", err)
				continue
			}
			fmt.Println(cwd)
		case "echo":
			fmt.Println(strings.Join(args[1:], " "))
		case "kill":
			if len(args) < 2 {
				fmt.Println("Необходимо указать PID для команды kill.")
				continue
			}
			pid, err := getPID(args[1])
			if err != nil {
				fmt.Println("Ошибка при получении PID:", err)
				continue
			}
			proc, err := os.FindProcess(pid)
			if err != nil {
				fmt.Println("Ошибка при поиске процесса:", err)
				continue
			}
			err = proc.Signal(os.Kill)
			if err != nil {
				fmt.Println("Ошибка при завершении процесса:", err)
			}
		case "ps":
			cmd := exec.Command("tasklist")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("Ошибка при выполнении команды tasklist:", err)
			}
		default:
			cmd := exec.Command("cmd", "/C", strings.Join(args, " "))
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("Ошибка при выполнении команды:", err)
			}
		}
	}
}

func readInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func getPID(pidArg string) (int, error) {
	pid, err := strconv.Atoi(pidArg)
	if err != nil {
		return 0, err
	}
	return pid, nil
}
