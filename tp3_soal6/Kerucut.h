#ifndef KERUCUT_H
#define KERUCUT_H

#include <iostream>
#include <cmath>
using namespace std;

struct Kerucut {
    int r; // jari-jari
    int t; // tinggi

    // Konstruktor
    Kerucut(int jari = 0, int tinggi = 0) {
        r = jari;
        t = tinggi;
    }

    int garisPelukis() const {
        return (int) sqrt((r * r) + (t * t));
    }

    int luasAlas() const {
        return (22 * r * r) / 7;
    }

    int luasPermukaan() const {
        int s = garisPelukis();
        return (22 * r * (r + s)) / 7;
    }

    int volume() const {
        return (22 * r * r * t) / (7 * 3);
    }
};

#endif
