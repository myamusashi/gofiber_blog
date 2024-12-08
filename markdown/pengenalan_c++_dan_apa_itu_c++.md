---
Title: sejarah bahasa pemrograman dan pengenalan bahasa c++
Description: sejarah singkat c++ dan apa hubungan antara c++ dengan c dan apa itu konsep kompilasi dan menjelaskan apa kegunaan dari file header
Parent: c++
Date: 2024-09-10 12:53
Slug: pengenalan_c++_dan_apa_itu_c++
Order: 4
tags: ["c++", "programming"]
MetaPropertyTitle: pengenalan c++ dan sejarah pada bahasa pemrograman c++
MetaDescription: sejarah singkat c++dan antara c++  dengan c dan apa itu konsep kompilasi enjelaskan apa itu file header dan perbedaan c++ klasik dengan c++ modern
MetaOgURL: https://blog.myamusashi.my.id/posts/pengenalan_c++_dan_apa_itu_c++

author:
    name: "Gilang Ramadhan"
    email: "gilang@gmail.com"
---

## Apa Itu Bahasa Pemrograman

Sebelum adanya C++ para ahli komputer masih menggunakan bahasa komputer yang masih primitif sekali karena pada masa nya bahasa
mesin yang dikenal hanya mengenal angka 1 dan 0 saja. 

Selanjutnya bahasa mesin disedehanakan menjadi bahasa yang lebih bisa dipahami manusia, dengan
menghadirkan statemen-statemen khusus yang disebut dengan ADD, MOV, JMP dan masih ada yang lainnya. Bahasa ini disebut dengan bahasa Assembly yang dikategorikan
bahasa tingkat rendah atau istilah dalam bahasa inggrisnya **_low level language_**.

## Sejarah dan Pengenalan C++

Pada 1980, **Bjarne Stroustrup** seorang ilmuan komputer Denmark, mengembangkan bahasa pemrograman C++ di AT&T Bell Laboratories. Stroustrup yang saat itu masih mengejar gelar Ph.D di Universitas Cambridge
mendapatkan pekerjaan dengan bahasa Simula.

Simula adalah bahasa pemrograman yang dikembangkan ilmuan komputer Norwegian **Kristen Nygaard** dan **Ole-Johan Dahl** pada tahun 1965, Simula bahasa pertama
yang mendukung paradigma pemrograman berorientasi obyek(OOP).

Muncul lah keinginan Stroustrup untuk menciptakan bahasa pemrograman dengan fitur yang lebih banyak terutama pada berorientasi objek. Ia menyadari paradigma OOP akan berguna pada pengembangan perangkat lunak.

Maksud simbol ++ pada C++ artinya adalah operator penaikan(increment), yang berarti ada penambahan fitur yang lebih canggih dan modern dibandingkan dengan C. 

## Fitur Yang Ada di C++

Ada banyak fitur yang dibuat di C++, antara lain:
*  Object-Oriented Programming 
*  Machine Independent
*  High-Level Language
*  Punya Komunitas Yang Besar
*  Simple syntax
*  Berbasis kompilasi
*  Dynamic Memory Allocation
*  Multi-threading

![gambar fitur c++](https://cdn.educba.com/academy/wp-content/uploads/2019/11/features-of-c.png.webp)
Gambar: [https://www.educba.com/features-of-c-plus-plus/](https://www.educba.com/features-of-c-plus-plus/) 

### 1. Object-Oriented Programming

Object-Oriented programming(OOP) atau pemrograman berorientasi objek adalah fitur yang paling penting dari C++ ini karena fitur ini bisa membuat atau menghapus/menghancurkan objek ketika menulis kode.
Fitur ini juga bisa membuat semacam blueprints dengan objek yang dipilih. 

Berikut konsep pada OOP:
* Class
* Encapsulation
* Polymorphism
* Inheritance
* Abstraction

### 2. Machine Independent 

File C++ yang dikompilasi bisa dijalankan pada semua mesin atau operasi sistem yang berbeda karena C++ adalah Machine Independent atau Mesin Mandiri dan bukan platform-independent.
Jika kita analogikan file binary C++ seperti file PDF, file PDF ini walaupun kita buka di Mac OS, Linux, Windows maupun Android tampilan, format dan yang lainnya akan terlihat sama. 

Bahasa pemrograman yang support *machine independent* dapat berjalan di berbagai perangkat keras atau sistem operasi yang berbeda, dengan syarat, environment atau lingkungan yang dibutuhkan file c++ terpenuhi atau sesuai untuk
menjalankannya.


### 3. High-Level Language

C++ adalah bahasa pemrograman dengan bahasa tingkat tinggi yang artinya bahasa pemrograman yang gampang dibaca oleh manusia, jika dibandingkan dengan C yang mempunyai level bahasa tingkat menengah. 
Ini membuat bahasa C++ lebih mudah dipahami dibandingkan bahasa C.

### 4. Komunitas yang besar

C++ mempunyai komunitas yang cukup besar dan juga mempunyai resource yang banyak di youtube, website, forum dan lain-lain.

### 5. Simple syntax

C++ mempunyai syntax atau tulisan kode yang gampang untuk dibaca, dan ada satu fitur yang membuat C++ menjadi lebih simple dan gampang digunakan yaitu ***auto***, tipe variable yang akan di deklarasikan 
secara otomatis oleh compiler.

Contoh simple Auto pada c++

```c++
#include <iostream>

int main (int argc, char *argv[]) {
    auto angka = 5;
    auto angka_1 = 10;

    auto karakter = "Saya";
    auto karakter_1 = "Gilang";

    std::cout << angka + angka_1 << " " << "\t Tipe data: " << typeid(angka).name() << "\n";
    std::cout << karakter << " " <<  karakter_1 << "\t Tipe data: " << typeid(karakter).name();
    return 0;
}
```

#### Output kode

 ```txt
15 	Tipe data: i

Saya Gilang	Tipe data: PKc
 ```

Tipe data **i** yang artinya integer, dan tipe data **PKc** yang berarti P = pointer, K = const dan c = char, dan jika 
digabung berarti const char* dan sebenarnya PKc ini juga sama seperti tipe data string dan fakta lainnya adalah, tipe data string sebenarnya tidak
ada di bahasa C yang ada hanya PKc.

### 6. Berbasis kompilasi

C++ adalah bahasa yang berbasis kompilasi yang artinya sebelum kita mengeksekusi kode tersebut
kita harus mengcompile kode tersebut menjadi machine code atau kode binari yang bisa dibaca oleh komputer.
Ini memungkinkan kode c++ yang kita buat menjadi lebih cepat untuk menjalankan program yang kita buat, daripada bahasa pemrograman yang berbasis interprated seperti Python.

### 7. Dynamic Memory Allocation

Dynamic Memory Allocation atau alokasi memory dinamis memungkinkan suatu program untuk mengalokasikan memori secara manual. 
Ada dua memori alokasi yang digunakan dalam C++: 
* **Stack**

Semua local variable yang di deklarasikan akan masuk kedalam sebuah fungsi yang akan mengambil memori dari stack.

* **Heap**, 

Heap adalah memori yang tidak terpakai dan akan digunakan ketika program berjalan dan mengalokasikan memori ke Heap.  


Untuk penjelasan lebih lanjut bisa klik penjelasan lebih detail cara kerja memory dan pointer pada C++ [klik disini](/posts/cara_kerja_memori_dan_pointer_di_c++)

### 8. Multi-threading

Fitur multi-threading memungkinkan program yang akan dieksekusi bisa menjalankan lebih dari dua program yang akan dijalankan secara bersamaan.
Ada dua tipe multitasking, yaitu: process-based dan thread based.



## Referensi 

1. [Bahasa Pemrograman C++: Sejarah, Fitur, Kelebihan, dan Pengembangannya dalam Industri](https://www.gamelab.id/news/2677-bahasa-pemrograman-c-sejarah-fitur-kelebihan-dan-pengembangannya-dalam-industri) - Rifka Amalia, 2 Agustus 2023
2. [Simula](https://www.computerhope.com/jargon/s/simula.htm) - Computer Hope, 26 April 2017
3. [Features of C++](.www.geeksforgeeks.org/features-of-cpp/) - geeksforgeeks, 12 Januari 2023
4. [Steveng. "Stack Memory vs Heap Memory [duplicate]"](https://stackoverflow.com/questions/5836309/stack-memory-vs-heap-memory) - Stackoverflow, 29 April 2011
5. [C++ Multithreading](https://www.tutorialspoint.com/cplusplus/cpp_multithreading.htm) - tutorialspoint 
