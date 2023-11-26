package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

/*
CATATAN PENTING
untuk menjalankan go , kita harus melakukan kompile terlebih dahulu.
Cara menjalankan go :
1. Masuk ke dalam folder yang berisi file nya.
2. jalankan perintah go build <file yang mau dikompile>
3. lalu akan muncul file dengan nama yang sama , namun tanpa extention.
4. lalu masuk ke dalam file tsb , melalui terminal. dengan memasukan perintah ( ./ <nama file> )


sebenarya saat development kita bisa langsung menjalankan
file secara langsung . kita menjalankan dengan cara kompile
saat akan melakukan realese ke server (production). karena yang
akan dikirm dalam bentuk binary.

file yang di kompile , akan menjadi binary file .



Untuk menghapus file binary dalam folder , cukup menjalankan perintah :
rm <nama binary file yang akan dihapus>
*/
