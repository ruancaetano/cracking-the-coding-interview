#include "node.h"

#include <cstdlib>

Node::Node(int val) : value(val), next(nullptr) {}

int Node::getValue() const
{
    return value;
}

void Node::setValue(int v)
{
    value = v;
}

Node *Node::getNext() const
{
    return next;
}

void Node::setNext(Node *n)
{
    next = n;
}

void Node::unlinkNode(Node *previous)
{
    if (previous == nullptr)
    {
        return;
    }

    previous->next = next;
    return;
}
