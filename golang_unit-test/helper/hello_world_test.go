/**
aturan fungsi unit test
jika kita punya fungsi ==> helloworld
maka fungsi unit testnya ==> Testhelloworld

- diawali dengan kata Test dan akhirnya tak ada aturan
- harus memakai parameter (t *testing.T) dan tidak mengembalikan return value
*/

package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
*
Assertion digunakan untuk mengganti if else yang menyebalkan :( di golang

	dengan menggunakan library testify
	source download : go get github.com/stretchr/testify

-> Assertion ini jika gagal akan memanggil fail bukan failNow
*/
func TestHelloWorldAssert(t *testing.T) {
	result := helloWorld("Amir")
	assert.Equal(t, "hello Amir", result, "Result Must be hello Amir")
	fmt.Println("Hello World With Assert Done")
}

/*
*
Require sama seperti Assertion namun yang dipanggil failNow
*/
func TestHelloWorldRequire(t *testing.T) {
	result := helloWorld("Amir")
	require.Equal(t, "hello Amir", result, "Result Must be hello Amir")
	fmt.Println("Hello World With Assert Done")
}

/*
*
How to run Unit Test

go test ==> untuk run semua unit test yang ada di suatu folder pwd terakhir

go test -v ==> sama seperti go test namun ada keterangan fungsi yang dikerjakan dan lama waktu yang dibutuhkan

go test -v ./... ==> sama seperti go test semua package yang ada di run

go test -v -run (functionName) ==> untuk run fungsi yang ingin di test
*/
func TestHelloWorldAmir(t *testing.T) {
	result := helloWorld("Amir")
	if result != "hello Amir" {
		//unit test failed
		// t.Fail()
		t.Error("harusnya hello Amir")
	}
	fmt.Println("Hello World Done")
}

func TestHelloWorldKing(t *testing.T) {
	result := helloWorld("King Amir")
	if result != "hello King Amir" {
		//unit test failed
		// t.FailNow()
		t.Fatal("harusnya hello King Amir")
	}
	fmt.Println("Hello World King Done")
}

/**
How to Failed Unit Test
	"jangan pakai panic"
t.Fail() ==> menggagalkan unit test namun tetap lanjut eksekusi unit test. Tapi di akhir di anggap gagal kok:)

t.FailNow() ==> menggagalkan unit test saat ini juga tanpa lanjut eksekusi unit test

t.Error(args...) ==> memberi log (print) error, setelah itu memanggil fungsi 't.Fail()'

t.Fatal(args...) ==> sama seperti error tapi manggil fungsi 't.FailNow()'
*/

/**
Skip ==> untuk membatalkan
*/
