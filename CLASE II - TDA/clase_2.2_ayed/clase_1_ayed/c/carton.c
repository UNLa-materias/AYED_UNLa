

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

#include "funciones.h"
#include "carton.h"


struct CartonE{

    int vec[8];
    char direccion[15];
    //fecha tiene que ser domingo!!!

};

char* getDireccion(CartonPtr c){

    return c->direccion;
};

int* getVec(CartonPtr c){
    return c->vec;

};


void setDireccion(CartonPtr c, char nuevaDireccion[15]){

    strcpy(c->direccion, nuevaDireccion);
};


void setVec(CartonPtr c, int nuevoVec[8]){

    //c->vec = nuevoVec; No se puede, igual que con char necesitabamos strcpy
    memcpy(c->vec, nuevoVec, sizeof(int[8]));

};

//Cuando se trabaja con arreglos se suelen hacer los get y set por posición
int getVecPos(CartonPtr c, int pos){

    return c->vec[pos];

};

void setVecPos(CartonPtr c, int pos, int nuevoElementoVec){

    c->vec[pos] = nuevoElementoVec;

};




CartonPtr crearCartonTeclado(); //Comprador seleccione los numeros

CartonPtr crearCarton2(int v[8], char d[15]);

CartonPtr crearCartonAutomatico(char d[15]){

    CartonPtr c = (CartonPtr) malloc(sizeof(struct CartonE));

    strcpy(c->direccion, d);



    //Valores -1 para mejorar la busqueda
    cargarVectorIntConValor(c->vec, 8, -1);

    //puse 10 en lugar de 99 para verificar repetdos mas facil
    cargarVectorIntAleatorioSinRepetir(c->vec, 8, 0,10);



    return c;

};

void destruirCarton(CartonPtr c){
    free(c);
};

void mostrarCarton(CartonPtr c){

    printf("\n\t---------- CARTON ---------\n");
    printf("\tDIRECCION: %s\n\t---VECTOR---\n\t", c->direccion);
    //mostrarVector
    for (int i =0 ; i<8; i++){

        printf(" %d ", c->vec[i]);

    }



};
