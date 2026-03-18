#include <stdio.h>
#include <stdlib.h>
#include <time.h>

void cargarDosArrayIdenticos(int array1[],int array2[], int n){

    for (int i = 0 ; i<n; i++){

        int num = rand()%1000;

        array1[i] = num;
        array2[i] = num;

    }


}


void encontrarParFuerzaBruta(int arr[], int n, int t) {

    for (int i = 0; i < n; i++) {
        for (int j = i + 1; j < n; j++) {
            if (arr[i] + arr[j] == t) {
                printf("\n----------- FORx2 -----------");
                printf("\n\nPar encontrado: (%d, %d)\n", arr[i], arr[j]);
                return;
            }
        }
    }
    printf("No se encontro ningun par que sume %d\n", t);
}


#define TABLE_SIZE 1000

typedef struct HashNode {
    int key;
    struct HashNode *next;
} HashNode;

int hash(int key) {
    return abs(key) % TABLE_SIZE;
}

void insert(HashNode *hashTable[], int key) {
    int hashIndex = hash(key);
    HashNode *newNode = (HashNode *)malloc(sizeof(HashNode));
    newNode->key = key;
    newNode->next = hashTable[hashIndex];
    hashTable[hashIndex] = newNode;
}

int search(HashNode *hashTable[], int key) {
    int hashIndex = hash(key);
    HashNode *current = hashTable[hashIndex];
    while (current != NULL) {
        if (current->key == key) {
            return 1;
        }
        current = current->next;
    }
    return 0;
}

void encontrarParHashMap(int arr[], int n, int t) {
    HashNode *hashTable[TABLE_SIZE] = {NULL};

    for (int i = 0; i < n; i++) {
        int complemento = t - arr[i];
        if (search(hashTable, complemento)) {
            printf("\n----------- HASHMAP -----------");
            printf("\n\nPar encontrado: (%d, %d)\n", complemento, arr[i]);
            return;
        }
        insert(hashTable, arr[i]);
    }
    printf("No se encontró ningún par que sume %d\n", t);
}


int main() {

    struct timespec start, end;

    srand(time(0));
    int loQueBusco = 9;

    int n = 20000;
    int vector1[n];
    int vector2[n];


    cargarDosArrayIdenticos(vector1,vector2,n);

    /// Algoritmo con doble FOR

    // Medir el tiempo de inicio
    clock_gettime(CLOCK_MONOTONIC, &start);
    encontrarParFuerzaBruta(vector1, n, loQueBusco); ///<------ ALGORITMO
    // Medir el tiempo de finalización
    clock_gettime(CLOCK_MONOTONIC, &end);
    // Calcular el tiempo transcurrido en milisegundos
    double time_taken = (end.tv_sec - start.tv_sec) * 1000.0 + (end.tv_nsec - start.tv_nsec) / 1000000.0;
    // Mostrar el tiempo tomado con precisión
    printf("\nEl algoritmo doble FOR tardo %f milisegundos en ejecutarse.\n", time_taken);


    /// Algoritmo con mapeo
    // Medir el tiempo de inicio
    clock_gettime(CLOCK_MONOTONIC, &start);
    encontrarParHashMap(vector2, n, loQueBusco);///<------ ALGORITMO
    // Medir el tiempo de finalización
    clock_gettime(CLOCK_MONOTONIC, &end);
    // Calcular el tiempo transcurrido en milisegundos
    time_taken = (end.tv_sec - start.tv_sec) * 1000.0 + (end.tv_nsec - start.tv_nsec) / 1000000.0;
    // Mostrar el tiempo tomado con precisión
    printf("\nEl algoritmo  HashMap tardo %f milisegundos en ejecutarse.\n", time_taken);


    return 0;
}
