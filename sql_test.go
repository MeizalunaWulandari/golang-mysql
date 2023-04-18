package golang_database

import (
	"testing"
	"fmt"
	"context"
)

func TestExecSql(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES ('rizka', 'Rizka')"
	_, err := db.ExecContext(ctx, script)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT * FROM customer"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string

		err := rows.Scan(&id, &name)
		// Posisi harus sesuai dengan mapping table
		if err != nil {
			panic(err)
		}

		fmt.Println("Id : ", id)
		fmt.Println("Name : ", name)
	}
}

/**
 * EKSEKUSI PERINTAH SQL
 * Untuk menjalankan perintah SQL di golang bisa menggunakan function
 * (DB) ExectContext(context, sql, params)
 * params sifatnya tidak wajib
 * Ketika kita menggunakan perintah SQL, kita butuh menjalankan context, 
 * dengan context kita bisa mengirim sinyal cancel jika kita ingin mebatalkan
 * pengiriman perintah SQLnya
 * 
 * QUERY SQL
 * Untuk operasi SQL yang tidak membutuhkan hasil, kita bisa menggunakan perintah 
 * Exec, namun jika kita membutuhkan hasil seperti SELECT SQL, kita bisa menggunakan
 * function berbeda yang berbeda 
 * Function untuk melakukan query ke databale, bisa menggunakan function 
 * (DB) QueryContext(context, sql, params) params tidak wajib
 * * ROWS
 * Hasil query function adalah sebuah data struct sql.Rows
 * Rows digunakan untuk melakukan iterasi terhadap hasil dari query 
 * Kita bisa menggunakan function (Rows)Next()(boolean) untuk melakukan iterasi terhadap 
 * data hasil query, jika return false, artinya sudah tidak ada data lagi dalam result
 * Untuk membaca tiap data kita bisa menggunakan (Rows)scan(columns...)
 * Dan jangan lupa, setelah menggunakan Rows, Jangan lupa untuk menutupnya menggunakan 
 * (rows)Close()
 */
