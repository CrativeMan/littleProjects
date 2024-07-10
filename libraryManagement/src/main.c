#include <stdio.h>

void printMenu();

char input;

int main()
{
    printMenu();
    switch (input)
    {
    case '0':
        printf("List Book entries\n");
        break;
    case '1':
        printf("Add book entry\n");
        break;
    case '2':
        printf("Edit book entry\n");
        break;
    case '3':
        printf("Delete book entry\n");
        break;
    case '4':
        printf("Settings\n");
        break;
    default:
        printf("Unknown asd\n");
        break;
    }
    return 0;
}

void printMenu()
{
    printf("=========================================================\n");
    printf("\t\t Boksa - Book managemnet\n");
    printf("=========================================================\n");
    printf("\t\t 0. List all book entries\n");
    printf("\t\t 1. Add new book entry\n");
    printf("\t\t 2. Edit book entry\n");
    printf("\t\t 3. Delete book entry\n");
    printf("\t\t 4. Settings\n");
    printf("=========================================================\n");
}