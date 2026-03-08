#ifndef PERSONA_H_INCLUDED
#define PERSONA_H_INCLUDED

#include "carton.h"

struct PersonaE{ //.nombre

    int dni;
    char nombre[15];

    //Un solo carton
    CartonPtr cartonDeLaPersona;
};

typedef struct PersonaE * PersonaPtr; //->nombre


//Constructores
PersonaPtr crearPersonaTeclado(); //Teclado

PersonaPtr crearPersona(int d, char n[15]);


//Destructor
void destruirPersona(PersonaPtr p);

//Mostrar
void mostrarPersona(PersonaPtr p);


///Agregar cosas...
//agregarCarton
//cargarCarton
//crearPersona(dir)




#endif // PERSONA_H_INCLUDED
