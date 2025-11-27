/*
Д/з:
Посмотреть про большие файлы!!!
Создание, закрытие, чтение и запись в них
*/

/*
Классная работа
Реализовать на Go модуль, осуществляющий обход директории, переданной
программе в качесстве аргумента.
Во время обхода директории Все файлы записываются в файл files.txt ->
директории -> folders.txt | имена файлов
1) Go рутины для самых уверенных
2) аргумент можно принять os.Args[1] // main.exe c:\users\user\Desktop\
3) для обхода воспользуйтесь filepath.WalkDir()
4) file.Stat() -> IsDir()?
*/

package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"sync"
)

type FileSystemWalker struct {
	mu      sync.RWMutex
	files   []string
	folders []string
}

func NewFileSystemWalker() *FileSystemWalker {
	return &FileSystemWalker{}
}

func (fsw *FileSystemWalker) WalkParallel(root string) error {
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, runtime.NumCPU()*2) //чтобы ноут не взлетел

	// 3) для обхода воспользуйтесь filepath.WalkDir()
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil //игнор ошибки
		}

		wg.Add(1)
		semaphore <- struct{}{} //заняли место

		// отдельно для каждого файла и папки 1) Go рутины для самых уверенных
		go func(p string, entry fs.DirEntry) {
			defer wg.Done()
			defer func() { <-semaphore }()

			fsw.mu.Lock()
			defer fsw.mu.Unlock()

			// 4) file.Stat() -> IsDir()?
			if entry.IsDir() {
				fsw.folders = append(fsw.folders, p)
			} else {
				fsw.files = append(fsw.files, p)
			}
		}(path, d)

		return nil
	})

	wg.Wait()
	close(semaphore)
	return err
}

func (fsw *FileSystemWalker) WriteResults() {

	sort.Strings(fsw.files)
	sort.Strings(fsw.folders)

	// Все файлы записываются в файл files.txt -> директории -> folders.txt | имена файлов
	writeToFile("files.txt", fsw.files)
	writeToFile("folders.txt", fsw.folders)

	fmt.Printf("Found %d files and %d folders\n", len(fsw.files), len(fsw.folders))
}

func writeToFile(filename string, data []string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating %s: %v\n", filename, err)
		return
	}
	defer file.Close()

	for _, item := range data {
		file.WriteString(item + "\n")
	}
}

func main() {
	// os.Args[1] - целевая директория
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <directory>")
		fmt.Println("Example: go run main.go C:\\Users\\User\\Desktop")
		os.Exit(1)
	}

	// 2) аргумент можно принять os.Args[1] // main.exe c:\users\user\Desktop\
	root := os.Args[1]

	if _, err := os.Stat(root); os.IsNotExist(err) {
		fmt.Printf("Error: Directory '%s' does not exist\n", root)
		os.Exit(1)
	}

	fmt.Printf("Scanning directory: %s\n", root)
	fmt.Printf("Using %d CPU cores for parallel processing\n", runtime.NumCPU())

	walker := NewFileSystemWalker()

	if err := walker.WalkParallel(root); err != nil {
		fmt.Printf("Walk error: %v\n", err)
	}

	walker.WriteResults()
	fmt.Println("Scan completed!")
}

// file, err := os.OpenFile("test.txt", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0777)
// if err != nil {
// 	fmt.Println(err)
// }
// defer file.Close()
// fmt.Fprint(file, "Hello, World!\n")

// file, err := os.Open("test.txt")
// if err != nil {
// 	fmt.Println(err)
// }
// defer file.Close()

// //Запрашиваем инфу о файле (файл есть файл)
// info, err := file.Stat()
// if err!=nil{
//     fmt.Println(err)
// }
// fmt.Println(info.IsDir())
// fmt.Println(info.Name())
// fmt.Println(info.Size())
// fmt.Println(info.Mode())

// err := os.Mkdir("test", 0666)
// if err!=nil{
//     fmt.Println(err)
// }

// err := os.MkdirAll("./test/folder1/folder2", 0666)
// if err != nil {
// 	fmt.Println(err)
// }

// file, err := os.Create("test.txt")
// if err != nil {
// 	fmt.Println(err)
// 	return
// }
// defer file.Close()

// hello := "Hello!"
// file.WriteString(hello)

//Выше создание

// file, err := os.Open("test.txt")
// if err != nil {
// 	fmt.Println(err)
// 	return
// }

// scanner := bufio.NewScanner(file)
// for scanner.Scan() {
// 	line := scanner.Text()
// 	fmt.Println(line + "\n")
// }

// data, err := os.ReadFile("test.txt")
// if err != nil {
//     fmt.Println(err)
// }
// fmt.Println(string(data))

// file, err := os. OpenFile("test.txt", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
// if err!= nil{
// 	fmt.Println(err)
// 	return
// }
// defer file.Close()

// // hello := "Hello!\n"
// // file.WriteString(hello)
// hello := []byte("Hello!\n new line\n")
// file.Write(hello)
