# <h1>LoyaltoQoinTechnicalTest</h1>
<h2>I. Soal Teori</h2>

1. Rancangkanlah diagram database untuk aplikasi rumah makan.
Jelaskan teknologi yang akan dipakai untuk aplikasi ini dan mengapa anda memilih teknologi
tersebut.
Kebutuhan:
1. Aplikasi ini bisa memasukkan pesanan-pesanan makanan pelanggan
2. Aplikasi ini bisa mengeluarkan struk pembelian
3. Aplikasi ini bisa mengeluarkan laporan penghasilan mingguan dan bulanan
4. Aplikasi ini bisa mengeluarkan laporan stok
Selain kebutuhan pokok diatas, silahkan tambahkan ide original anda untuk membuat aplikasi
lebih baik.
(catatan: Soal ini tidak membutuhkan pengambil test untuk membuat aplikasi (coding). )

Berikut adalah Diagram aplikasi rumah makan sederhana:

<img src="Soal-1/diagram.png">

<h3>Penjelasan</h3>

<h4>Dalam diagram tersebut terdapat 5 table:</h4>
1. employees merupakan table yang berisikan admin,kasir nantinya<br>
2. menus merupkan table untuk menyimpan data semua menu makanan di restaurant<br>
3. structs merupakan table yang digunakan untuk mencetak struk pembayaran<br>
4. orders merupakan table yang digunakan oleh kasir untuk mencatat pesanan customers<br>
5. order_menus merupakan table yang berelasi dengan orders dan menus sehingga semua menu yang dipesan customers akan tercatat disini

<h4>Technology yang akan dipakai adalah:</h4>
<ul>
<li>Go merupakan salah satu bahasa pemograman yang kinerja nya sangat cepat dan mudah dipelajari sehingga cocok untuk dijadikan sebagai bahasa pemograman untuk Backend</li>
<li>ReactJs merupakan salah satu framework front end yang responsifnya cepat dan interaktif. reactjs juga merupakan framework yang populer saat ini</li>
<li>Mysql dipilih sebagai sistem manajemen basis data karena mudah digunakan dan dapat menangani skala besar data. MySQL juga memungkinkan penggunaan indeks pada tabel-tabel database, sehingga mempercepat pengambilan data pada saat aplikasi membutuhkan laporan penghasilan mingguan dan bulanan serta laporan stok.</li>
<li>GORM merupakan salah satu library ORM dari bahaasa pemograman Go, gorm memudahkan nantinya untuk berinteraksi dengan data</li>
<li>Echo merupakan salah satu framework dari bahasa pemograman Go , dipilih karena sangat ringan dan performanya cepat, sangat bagus jika digunakan untuk membuat API</li>
</ul>

<h2>II. Soal Praktek</h2>

Buatlah sebuah script permainan dadu yang menerima input N jumlah pemain dan M jumlah dadu, dengan peraturan sebagai berikut:
1. Pada awal permainan, setiap pemain mendapatkan dadu sejumlah M unit.
2. Semua pemain akan melemparkan dadu mereka masing-masing secara bersamaan
3. Setiap pemain akan mengecek hasil dadu lemparan mereka dan melakukan evaluasi
seperti berikut:
a. Dadu angka 6 akan dikeluarkan dari permainan dan ditambahkan sebagai poin
bagi pemain tersebut.
b. Dadu angka 1 akan diberikan kepada pemain yang duduk disampingnya.
Contoh, pemain pertama akan memberikan dadu angka 1 nya ke pemain kedua.
c. Dadu angka 2,3,4 dan 5 akan tetap dimainkan oleh pemain.
4. Setelah evaluasi, pemain yang masih memiliki dadu akan mengulangi step yang ke-2
sampai tinggal 1 pemain yang tersisa.
a. Untuk pemain yang tidak memiliki dadu lagi dianggap telah selesai bermain.
5. Pemain yang memiliki poin terbanyak lah yang menang.
Buatlah script ini menggunakan bahasa yang kamu kuasai.


<h3>Jawaban</h3>

- Clone to directory

  ``` $git clone github.com/fauzilax/LoyaltoQoinTechnicalTest.git ```

- Change directory
 
  ``` cd LoyaltoQoinTechnicalTest```
 
- Run The Program

  ``` $go run main.go ```










