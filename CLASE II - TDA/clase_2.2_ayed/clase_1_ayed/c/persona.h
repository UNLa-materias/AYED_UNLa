#ifndef PERSONA_H_INCLUDED
#define PERSONA_H_INCLUDED

#include "carton.h"
//La struct deja de estar implementada en el .h
struct PersonaE;

typedef struct PersonaE * PersonaPtr; //->nombre
//Getters y setters
int getDNI(PersonaPtr p);
char* getNombre(PersonaPtr p);
CartonPtr getCartonPtr(PersonaPtr p);

void setDNI(PersonaPtr p, int nuevoDNI);
void setNombre(PersonaPtr p, char nuevoNombre[15]);
void setCartonPtr(PersonaPtr p, CartonPtr nuevoCarton);

//




//Constructores
PersonaPtr crearPersonaTeclado(); //Teclado

PersonaPtr crearPersona(int d, char n[15]);

PersonaPtr crearPersonaYCarton(int d, char n[15], char dir[15]);

//Destructor
void destruirPersona(PersonaPtr p);

//Mostrar
void mostrarPersona(PersonaPtr p);


///Agregar cosas...
//agregarCarton
//cargarCarton
//crearPersona(dir)




#endif // PERSONA_H_INCLUDED
