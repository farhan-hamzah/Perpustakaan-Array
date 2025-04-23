package main
import "fmt"
const NMAX int = 1024
type infoBuku struct{
	idBuku, judulBuku, penulisBuku string
}
type arrBuku[NMAX]infoBuku
func main(){
	var dataBuku arrBuku
	var nBuku, i int
	var searchKey string
	var id, judul, penulis string
	fmt.Scan(&nBuku)
	for i = 0; i < nBuku; i++{
		fmt.Scan(&id, &judul, &penulis)
		dataBuku[i].idBuku = id
		dataBuku[i].judulBuku = judul
		dataBuku[i].penulisBuku = penulis
	}
	fmt.Scan(&searchKey)
	pencarianBuku(dataBuku, nBuku, searchKey)
}
func pencarianBuku(dataBuku arrBuku, nBuku int, searchKey string){
	var buku infoBuku
	var b bool
	var i int
	b = false
	for i =0; i < nBuku; i++{
		if dataBuku[i].idBuku == searchKey || dataBuku[i].judulBuku == searchKey || dataBuku[i].penulisBuku == searchKey{
			b = true
			buku = dataBuku[i]
		}
	}
	fmt.Println(" ")
	if b == true{
		fmt.Println("ID: ", buku.idBuku)
		fmt.Println("Judul Buku: ", buku.judulBuku)
		fmt.Println("Penulis Buku: ", buku.penulisBuku)
	}else{
		fmt.Println("Buku Tidak Tersedia")
	}
}