#ifndef NODE_H
#define NODE_H

class Node
{
private:
    int value;
    Node *next;

public:
    // constructor
    Node(int val);

    void unlinkNode(Node *previous);

    Node *getNext() const;
    int getValue() const;

    void setNext(Node *n);
    void setValue(int v);
};

#endif