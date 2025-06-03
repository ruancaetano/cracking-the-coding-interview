#include "linked_list.h"

#include <iostream>
#include <ctime>
#include <cstdlib>

LinkedList::LinkedList(Node *h) : head(h) {}

Node *LinkedList::getHead() const
{
    return head;
}

void LinkedList::setHead(Node *h)
{
    head = h;
}

void LinkedList::print()
{
    Node *currentNode = head;
    while (currentNode != nullptr)
    {
        std::cout << currentNode->getValue() << " ";
        currentNode = currentNode->getNext();
    }

    std::cout << std::endl;
}

void LinkedList::deleteNodes()
{
    Node *currentNode = head;
    while (currentNode != nullptr)
    {
        Node *next = currentNode->getNext();
        delete currentNode;
        currentNode = next;
    }
}

LinkedList *LinkedList::generateRandomSingleList(int size)
{
    srand(time(0));

    LinkedList *list = new LinkedList(new Node(0));
    list->setHead(new Node(0));
    Node *current = list->getHead();
    for (int i = 0; i < size; i++)
    {
        current->setValue(rand() % 100);
        if (i != size - 1)
        {
            Node *nextNode = new Node(0);
            current->setNext(nextNode);
            current = nextNode;
        }
    }

    return list;
}

LinkedList *LinkedList::convertArrayToSingleList(int *arr, int size)
{
    LinkedList *list = new LinkedList(nullptr);
    Node *current = nullptr;

    for (int i = 0; i < size; i++)
    {

        int *valuePtr = &arr[i];

        if (valuePtr == nullptr)
            break;

        if (current == nullptr)
        {
            current = new Node(*valuePtr);
            list->setHead(current);
        }
        else
        {
            current->setNext(new Node(*valuePtr));
            current = current->getNext();
        }
    }

    return list;
}

Node *LinkedList::findMiddle()
{
    Node *slow = head;
    Node *fast = head->getNext();

    while (fast != nullptr && fast->getNext() != nullptr)
    {
        slow = slow->getNext();
        fast = slow->getNext();
    }

    return slow;
}

void LinkedList::sort()
{
    head = mergeSort(head);
}

Node *LinkedList::mergeSort(Node *head)
{
    if (head == nullptr || head->getNext() == nullptr)
    {
        return head;
    }

    Node *middle = findMiddle();
    Node *secondHalf = middle->getNext();
    middle->setNext(nullptr);

    Node *left = mergeSort(head);
    Node *right = mergeSort(secondHalf);

    return mergeLists(left, right);
}

Node *LinkedList::mergeLists(Node *left, Node *right)
{
    if (left == nullptr)
    {
        return right;
    }

    if (right == nullptr)
    {
        return left;
    }

    Node *result = nullptr;
    if (left->getValue() < right->getValue())
    {
        result = left;
        result->setNext(mergeLists(left->getNext(), right));
    }
    else
    {
        result = right;
        result->setNext(mergeLists(left, right->getNext()));
    }

    return result;
}