

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

#include "funciones.h"
#include "carton.h"





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
