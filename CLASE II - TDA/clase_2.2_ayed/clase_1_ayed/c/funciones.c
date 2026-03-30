
#include <time.h>
#include "funciones.h"

void cargarVectorIntConValor(int vec[], int t, int valor){

    for ( int i = 0 ; i< 8; i++){

        vec[i] = valor;


    }


};

void cargarVectorIntAleatorioSinRepetir(int vec[], int t, int minimo, int maximo){


    for ( int i = 0; i< t; i++){

        int busqueda;
        int numero;

        do{
            numero = aleatorioEntre(minimo,maximo);

            busqueda = busquedaVectorInt(vec,t,numero);

        } while (busqueda != -1); //Mientras encuentro el numero busco otro.

        vec[i] = numero; //Guardo el numero solo si no lo encontre!!!

    }


};


void cargarVectorIntAleatorio(int vec[], int t, int minimo, int maximo){

    for ( int i = 0; i< t; i++){


            vec[i] = aleatorioEntre(minimo,maximo);


    }

};

int busquedaVectorInt(int vec[], int t, int busco){

    int encontre = -1;

    for ( int i = 0 ; i< t; i++){

        if ( busco == vec[i]){

            encontre = i;

            return encontre;

        }

    }

    return encontre;
};

int aleatorioEntre(int minimo, int maximo){

    int numero = minimo +  rand()%(maximo-minimo+1);

    return numero;

};

