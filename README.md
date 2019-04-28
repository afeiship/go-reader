# next-tree-md
> Create tree md string.

## install:
```bash
npm install -S afeiship/next-tree-md --registry=https://registry.npm.taobao.org
```

## usage:
```js
import nxTreeMd from 'next-tree-md'
nx.TreeMd('./test/demo');
```
```md
  + **dir1**
    - [test1.md](test/demo/dir1/test1.md)
    - [test2.md](test/demo/dir1/test2.md)
  + **test2**
    + **sub1**
      - [sub1.md](test/demo/test2/sub1/sub1.md)
      - [sub2.md](test/demo/test2/sub1/sub2.md)
    - [test1.md](test/demo/test2/test1.md)
    - [test2.md](test/demo/test2/test2.md)
    - [test3.md](test/demo/test2/test3.md)
```

## resources
- https://github.com/mihneadb/node-directory-tree
