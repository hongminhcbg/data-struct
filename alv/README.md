## AVL tree, my implement
1. Source GeeksforGeeks
    T1, T2 and T3 are subtrees of the tree 
    rooted with y (on the left side) or x (on
    the right side)           
       y                               x
      / \     Right Rotation          /  \
     x   T3   - - - - - - - >        T1   y 
    / \       < - - - - - - -            / \
   T1  T2     Left Rotation             T2  T3
    Keys in both of the above trees follow the
    following order
    keys(T1) < key(x) < keys(T2) < key(y) < keys(T3)
    So BST property is not violated anywhere.
