program Ganjil
kamus
    bilangan merupakan integer
    staus merupakan boolean
algoritma
    // Input
    baca(bilangan)
    // Proses
    status <- bilangan mod 2 = 1
    // Output
    jika status maka
        tulis("Bilangan ", bilangan, " adalah bilangan ganjil")
    jika tidak maka
        tulis("Bilangan ", bilangan, " adalah bukan bilangan ganjil")
endprogram