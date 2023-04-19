package golang_database

import (
	"testing"
	"fmt"
	"context"
	"time"
	"database/sql"
)

func TestExecSql(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	var script[3]string

	script[0] = "INSERT INTO customer(id, name, email, balance, rating, birth_date, married) VALUES ('luna', 'Luna', 'luna@gmail.com', 1000000, 5.0, '2006-05-19',false)"
	script[1] = "INSERT INTO customer(id, name, email, balance, rating, birth_date, married) VALUES ('andini', 'Andini', 'andini@gmail.com', 1000000, 5.0, '1987-09-28',true)"
	script[2] = "INSERT INTO customer(id, name, email, balance, rating, birth_date, married) VALUES ('rizka', 'Rizka', 'rizka@gmail.com', 1000000, 5.0, '2006-01-04',false)"

	for i := 0; i < len(script); i++ {
		_, err := db.ExecContext(ctx, script[i])

		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Success insert new customer")
}
func TestExecSqlUpdate(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "UPDATE customer SET email = null, birth_date = null, rating = 8.3 WHERE id = 'andini' "

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

func TestQuerySqlComplex(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		// Posisi harus sesuai dengan mapping table
		if err != nil {
			panic(err)
		}

		fmt.Println("===============")
		fmt.Println("Id : ", id)
		fmt.Println("Name : ", name)
		if email.Valid {
			fmt.Println("Email : ", email)
		}
		fmt.Println("Balance : ", balance)
		fmt.Println("Rating : ", rating)
		if email.Valid {
			fmt.Println("Birth Date : ", birthDate)
		}
		fmt.Println("Created At : ", createdAt)
		fmt.Println("married : ", married)
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
 * 
 * TIPE DATA COLUMN
 * * Mapping Tipe Data
 * VARCHAR, CHAR => string
 * INT, BIGINT => int32, int64
 * FLOAT, DOUBLE => float32, float64
 * BOOLEAN => bool
 * DATE, DATETIME, TIME, TIMESTAMP => time.Time
 * Saat select data disarankan menybutkan nama kolomnya agar tidak berubah-rubah saat 
 * ALTER TABLE dan tidak sarankan menggunakan SELECT * FROM table karena jika table dirubah
 * maka posisi di golang juga akan berubah
 * 
 * * Error Tipe Data Date
 * Message : "unsported Scan, storing driver.Value type[]uint8"
 * Secara default, Driver MySQL untuk golang akan melakukan query type data DATE, DATETIME
 * TIMESTAMP mejadi []byte / []uint8, dimana ini bisa dikonversi memjadi string lalu di 
 * parsing menjadi time.Time
 * Namun hal ini merepotkan jika dilakukan manual, kita bisa meminta driver MySQL untuk golang
 * secara otomatis melakukan parsing dengan dengan menambahkan paramater parseTime=true
 * pada (DB)connection
 * 
 * * Nullable Type
 * Golang database tidak mengerti dengan type data NULL di database
 * Oleh karena itu, khusus untuk kolom yang bisa null di database, akan menjadi masalah jika
 * kita melakukan scan secara bulat-bulat menggunakan type data representasinya di golang
 * 
 * * Error Data Null
 * Message : converting NULL to string is unsported
 * Konversi secara otomatis tidak didukung oleh driver MySQL Golang
 * oleh karena itu, khususk kolom yang bisa null, kita perlu menggunakan type data yang ada
 * pada package sql
 * 
 * * Type Data Nullable
 * string => database/sql.NullString
 * bool => database/sql.NullBool
 * float64 => database/sql.NullFloat64
 * int32 => database/sql.NullInt32
 * int64 => database/sql.NullInt64
 * Time.Time => database/sql.NullTime
 * Data yang akan dikembalikan berupa struct 
 */
