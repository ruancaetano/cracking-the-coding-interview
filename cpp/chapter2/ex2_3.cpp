#include <iostream>

#include "../common/linked_list.h"

void removeMiddle(Node *head)
{
    Node *prev = nullptr;
    Node *slow = head;
    Node *fast = head;
    while (fast != nullptr && fast->getNext() != nullptr)
    {
        prev = slow;
        slow = slow->getNext();
        fast = fast->getNext()->getNext();
    }

    slow->unlinkNode(prev);
}

int main()
{
    int arrayTwo[] = {1, 2, 3, 4, 5, 6};
    int sizeTwo = sizeof(arrayTwo) / sizeof(arrayTwo[0]);
    LinkedList *list = LinkedList::convertArrayToSingleList(arrayTwo, sizeTwo);
    std::cout << "Original list" << std::endl;
    list->print();
    std::cout << "Removing the middle" << std::endl;
    removeMiddle(list->getHead());
    std::cout << "Result list" << std::endl;
    list->print();

    return 0;
}
