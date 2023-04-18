package golang_database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestOpenConnection(t *testing.T){
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_database")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}


/**
 * INSTALL DATABASE DRIVER
 * go get -u github.com/go-sql-driver/mysql
 * 
 * MEMBUAT KONEKSI KE DATABASE
 * Untuk melakukan koneksi ke database kita bisa membuat objek sql.BD menggunakan 
 * function sql.Open(driver,dataSourceName)
 * Untuk menggunakan database MySQL kita bisa menggunakan driver "mysql"
 * Sedangkan untuk dataSourceName 
 * username:password@tcp(host:port)/database_name
 * objek sql.DB yang dihasilkan setelah berhasil terkoneksi berupa pointer
 * Jika objek sql.DB sudah tidak digunakan lagi, disarankan untuk menutupnya menggunakan
 * function close()
 */