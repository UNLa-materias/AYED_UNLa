
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "funciones.h"
#include "persona.h"


struct PersonaE{ //.nombre

    int dni;
    char nombre[15];

    //Un solo carton
    CartonPtr cartonDeLaPersona;

};

int getDNI(PersonaPtr p){
    return p->dni;

};
char* getNombre(PersonaPtr p){

    return p->nombre;
};

CartonPtr getCartonPtr(PersonaPtr p){

    return p->cartonDeLaPersona;
};

void setDNI(PersonaPtr p, int nuevoDNI){

    p->dni = nuevoDNI;
};

void setNombre(PersonaPtr p, char nuevoNombre[15]){

    strcpy(p->nombre, nuevoNombre);
};

void setCartonPtr(PersonaPtr p, CartonPtr nuevoCarton){

    p->cartonDeLaPersona = nuevoCarton;

};


//Constructores
PersonaPtr crearPersonaTeclado(); //Teclado

PersonaPtr crearPersona(int d, char n[15]){

    PersonaPtr p = (PersonaPtr) malloc(sizeof(struct PersonaE)); //void *

    p->dni = d;
    strcpy(p->nombre, n);

    p->cartonDeLaPersona = crearCartonAutomatico("Vacio");

    return p;


};

PersonaPtr crearPersonaYCarton(int d, char n[15], char dir[15]){

    PersonaPtr p = crearPersona(d,n);

   //strcpy( p->cartonDeLaPersona->direccion, dir); //Cambio para TDA

   setDireccion(p->cartonDeLaPersona, dir);


    return p;
}

;


//Destructor
void destruirPersona(PersonaPtr p){

    free(p);


};

//Mostrar
void mostrarPersona(PersonaPtr p){


    printf("\n---------- PERSONA -----------\n");
    printf("DNI: %d\n", p->dni);
    printf("NOMBRE: %s\n", p->nombre);

    mostrarCarton(p->cartonDeLaPersona);

};
