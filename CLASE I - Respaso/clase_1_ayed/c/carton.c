

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

#include "carton.h"





CartonPtr crearCartonTeclado(); //Comprador seleccione los numeros

CartonPtr crearCarton2(int v[8], char d[15]);

CartonPtr crearCarton1(char d[15]){

    CartonPtr c = (CartonPtr) malloc(sizeof(struct CartonE));

    strcpy(c->direccion, d);

    //cargar vector cargarVector(c->vec);

    for ( int i = 0 ; i<8; i++){

        c->vec[i] = rand()%100; //0 a 99
        //Verificar que no se repitan!!!!

    }

    return c;

};

void destruirCarton(CartonPtr c);

void mostrarCarton(CartonPtr c){

    printf("\n\t------ CARTON ----\n");
    printf("\tDIRECCION: %s\n\t---VECTOR---\n\t", c->direccion);
    //mostrarVector
    for (int i =0 ; i<8; i++){

        printf(" %d ", c->vec[i]);

    }



};
