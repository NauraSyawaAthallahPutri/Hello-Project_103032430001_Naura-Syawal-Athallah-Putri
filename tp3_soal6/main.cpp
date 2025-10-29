#include "Kerucut.h"

int main() {
    Kerucut k(7, 10); // jari-jari = 7, tinggi = 10

    cout << "Kerucut dengan r = " << k.r << ", t = " << k.t << endl;
    cout << "Garis pelukis    = " << k.garisPelukis() << endl;
    cout << "Luas alas        = " << k.luasAlas() << endl;
    cout << "Luas permukaan   = " << k.luasPermukaan() << endl;
    cout << "Volume           = " << k.volume() << endl;

    return 0;
    //Komen1
}
