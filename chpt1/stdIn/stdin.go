/*
Пакет bufio в Go предоставляет буферизированные ввод и вывод, что позволяет эффективнее работать с вводом-выводом за счет 
сокращения числа операций чтения и записи. Вот некоторые из наиболее популярных методов и типов, предоставляемых пакетом bufio:
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	var f *os.File
	f = os.Stdin
	defer f.Close()
	/*
	Создает новый буферизированный Reader.
	 Это полезно для чтения ввода, который необходимо обрабатывать постепенно, например, когда читаете большой файл или сетевой ответ.
	*/
	scanner := bufio.NewScanner(f)
	/*
	Предоставляет удобный интерфейс для построчного чтения. Scanner читает ввод и разбивает его на строки или слова, используя функцию разделения.
	*/
	for scanner.Scan(){
		fmt.Println(">", scanner.Text())
	}
}