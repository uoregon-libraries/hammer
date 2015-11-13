Hammer
===

When your only tool is a Hammer, everything looks like ... a ham distance
computation problem...?

It works like this:

```bash
/path/to/blockhash /path/to/images/*.jpg | ./bin/hammer
```

Output will tell you how similar the images appear to be:

```
Similarity between test1.jpg and test2.jpg: 76.953125% (59 out of 256 bits differed)
Similarity between test1.jpg and test3-different.jpg: 50.78125% (126 out of 256 bits differed)
Similarity between test2.jpg and test3-different.jpg: 52.734375% (121 out of 256 bits differed)
```

For reference, test1.jpg and test2.jpg are the same image, but one has a strip
across the bottom with copyright information.  test3-different.jpg is different
from either of the two, but it has the same copyright strip as test2.jpg.
