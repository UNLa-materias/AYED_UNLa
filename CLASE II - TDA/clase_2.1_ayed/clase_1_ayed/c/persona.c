
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "funciones.h"
#include "persona.h"




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

   strcpy( p->cartonDeLaPersona->direccion, dir);


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
