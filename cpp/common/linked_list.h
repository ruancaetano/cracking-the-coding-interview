#include "node.h"

#ifndef LINKED_LIST_H
#define LINKED_LIST_H

class LinkedList
{
private:
    Node *head;

    Node *mergeSort(Node *head);
    Node *mergeLists(Node *left, Node *right);

public:
    static LinkedList *generateRandomSingleList(int size);
    static LinkedList *convertArrayToSingleList(int *arr, int size);

    LinkedList(Node *head);

    Node *getHead() const;
    void setHead(Node *h);

    Node *findMiddle();

    void print();
    void deleteNodes();

    void sort();
};

#endif