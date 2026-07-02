# 1) 📌 Pengenalan Concurrency dan Parallelism

## Apa itu Concurrency?

Concurrency adalah kemampuan program untuk **mengelola beberapa pekerjaan secara bersamaan**. Beberapa tugas dapat berjalan dalam periode waktu yang sama, tetapi **tidak harus dieksekusi secara bersamaan pada waktu yang sama**.

Pada Go, concurrency diwujudkan menggunakan **goroutine** yang dikelola oleh **Go Runtime Scheduler**.

## Apa itu Parallelism?

Parallelism adalah kemampuan program untuk **menjalankan beberapa pekerjaan secara benar-benar bersamaan** pada waktu yang sama.

Parallelism dapat terjadi jika komputer memiliki lebih dari satu **CPU Core**, sehingga beberapa goroutine dapat dieksekusi secara simultan.

## Karakteristik Concurrency

- Berfokus pada **pengelolaan banyak pekerjaan** secara efisien.
- Tidak mengharuskan semua pekerjaan berjalan pada waktu yang sama.
- Dapat berjalan pada **single-core** maupun **multi-core** CPU.
- Pada Go, concurrency dicapai menggunakan **goroutine** yang dijadwalkan oleh Go Runtime.

## Karakteristik Parallelism

- Berfokus pada **menjalankan banyak pekerjaan secara bersamaan**.
- Membutuhkan **lebih dari satu CPU Core**.
- Beberapa goroutine benar-benar dieksekusi pada waktu yang sama.
- Parallelism dapat meningkatkan performa untuk pekerjaan yang dapat dijalankan secara independen.

## Ilustrasi

### Concurrency

```text
CPU

Task A : ███   ███   ███
Task B :    ███   ███
Task C :       ███   ███
```

Satu CPU menjalankan beberapa tugas secara bergantian dengan sangat cepat sehingga terlihat berjalan bersamaan.

### Parallelism

```text
CPU Core 1 : █████████
CPU Core 2 : █████████
CPU Core 3 : █████████
```

Beberapa CPU Core menjalankan tugas yang berbeda secara bersamaan.

## Perbedaan Concurrency dan Parallelism

| Concurrency | Parallelism |
|-------------|-------------|
| Mengelola banyak pekerjaan sekaligus | Menjalankan banyak pekerjaan sekaligus |
| Tidak harus berjalan pada waktu yang sama | Berjalan pada waktu yang sama |
| Dapat berjalan pada satu CPU | Membutuhkan lebih dari satu CPU Core |
| Fokus pada efisiensi | Fokus pada performa eksekusi |

## Hubungan dengan Goroutine dan Channel

Go menyediakan **goroutine** untuk menjalankan pekerjaan secara concurrent. Ketika banyak goroutine berjalan, sering kali mereka perlu saling berkomunikasi atau melakukan sinkronisasi.

Untuk kebutuhan tersebut, Go menyediakan **channel** sebagai media komunikasi antar goroutine sehingga data dapat dipertukarkan dengan aman tanpa harus berbagi memori secara langsung (*share memory by communicating*).

## Kesimpulan

Concurrency adalah kemampuan mengelola banyak pekerjaan secara bersamaan, sedangkan Parallelism adalah kemampuan menjalankan banyak pekerjaan secara benar-benar bersamaan. Go mendukung keduanya melalui **goroutine**, sedangkan **channel** digunakan sebagai mekanisme komunikasi dan sinkronisasi antar goroutine.

---

# 2) 📌 Pengenalan Channel

## Apa itu Channel?

Channel adalah mekanisme komunikasi **synchronous** di Go yang digunakan untuk bertukar data antar **goroutine**.

## Karakteristik Channel

- Channel merupakan media komunikasi yang dilakukan secara **synchronous**.
- Dalam sebuah channel terdapat:
  - **Pengirim (sender)**
  - **Penerima (receiver)**
- Biasanya pengirim dan penerima dijalankan oleh **goroutine yang berbeda**.
- Saat sebuah goroutine mengirim data ke channel, proses tersebut akan **ter-block** hingga ada goroutine lain yang menerima data tersebut.
- Karena perilaku tersebut, channel sering disebut sebagai mekanisme komunikasi **blocking (synchronous communication)**.
- Channel sangat cocok digunakan sebagai alternatif mekanisme sinkronisasi, mirip dengan konsep **async/await** pada beberapa bahasa pemrograman lain, namun dengan pendekatan yang lebih sederhana dan terintegrasi di dalam Go.

## Ilustrasi

```text
Goroutine A (Sender)
        |
        |  send data
        v
    +-----------+
    |  Channel  |
    +-----------+
        |
        | receive data
        v
Goroutine B (Receiver)
```

## Cara Kerja

1. Goroutine **Sender** mengirim data ke channel.
2. Jika belum ada goroutine yang menerima, maka sender akan **menunggu (blocked)**.
3. Ketika goroutine **Receiver** menerima data dari channel, proses pengiriman selesai.
4. Kedua goroutine kemudian dapat melanjutkan eksekusinya.

## Kesimpulan

Channel adalah salah satu fitur utama Go untuk komunikasi antar goroutine. Dengan channel, pertukaran data menjadi aman tanpa harus berbagi memori secara langsung (*share memory by communicating*), sehingga kode concurrent menjadi lebih mudah dipahami dan dikelola.
