#include <iostream>
#include <map>

#include "../common/linked_list.h"

void removeDuplicatedWithMap(Node *head)
{
    std::map<int, bool> foundValues;
    Node *current = head;
    Node *previous = nullptr;
    while (current != nullptr)
    {
        if (foundValues[current->getValue()])
        {
            current->unlinkNode(previous);
            current = previous->getNext();
            continue;
        }

        foundValues[current->getValue()] = true;
        previous = current;
        current = current->getNext();
    }

    return;
}

void removeDuplicatedWithoutMap(Node *head)
{
    Node *slow = head;
    while (slow != nullptr)
    {
        Node *fast = slow;

        while (fast->getNext() != nullptr)
        {
            if (slow->getValue() == fast->getNext()->getValue())
            {
                fast->setNext(fast->getNext()->getNext());
            }
            else
            {
                fast = fast->getNext();
            }
        }

        slow = slow->getNext();
    }
}

int main()
{

    int array[] = {2, 4, 3, 1, 3, 4, 2};
    int size = sizeof(array) / sizeof(array[0]);
    LinkedList *list = LinkedList::convertArrayToSingleList(array, size);
    std::cout << "Original list" << std::endl;
    list->print();
    std::cout << "Removing duplicates with aux structure" << std::endl;
    removeDuplicatedWithMap(list->getHead());
    std::cout << "Result list" << std::endl;
    list->print();

    std::cout << std::endl
              << "---------------"
              << std::endl
              << std::endl;

    int arrayTwo[] = {2, 4, 3, 1, 3, 4, 2};
    int sizeTwo = sizeof(arrayTwo) / sizeof(arrayTwo[0]);
    LinkedList *listTwo = LinkedList::convertArrayToSingleList(arrayTwo, sizeTwo);
    std::cout << "Original list" << std::endl;
    listTwo->print();
    std::cout << "Removing duplicates without aux structure" << std::endl;
    removeDuplicatedWithoutMap(listTwo->getHead());
    std::cout << "Result list" << std::endl;
    listTwo->print();
}
