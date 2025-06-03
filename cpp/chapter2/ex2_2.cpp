#include <iostream>

#include "../common/linked_list.h"

int kthToLast(Node *head, int kth)
{
    if (head == nullptr)
    {
        return 0;
    }

    int size = 1;
    Node *slow = head;
    Node *fast = head;

    for (int i = 0; i < kth; i++)
    {
        if (fast == nullptr)
        {
            return 0;
        }
        fast = fast->getNext();
    }

    while (fast != nullptr)
    {
        slow = slow->getNext();
        fast = fast->getNext();
    }

    return slow->getValue();
}

int main()
{

    int array[] = {1, 2, 3, 4, 5, 6, 7, 8, 9};
    int size = sizeof(array) / sizeof(array[0]);
    int kth = 5;
    LinkedList *list = LinkedList::convertArrayToSingleList(array, size);
    std::cout << "Original list" << std::endl;
    list->print();
    std::cout << kth << "th value to last: "
              << kthToLast(list->getHead(), kth)
              << std::endl;

    return 0;
}
