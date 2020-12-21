function  ListToTree<T extends {children?: Array<T>, id: string}>(list: T[]){
  let node: any, i;
  const  tree= []
  const map:any = {}
  for (let i = 0; i < list.length; i ++) {
    map[list[i].id] = list[i];
    list[i].children = [];
  }
  for (i = 0; i < list.length; i += 1) {
    node = list[i];
    if (node.parent_id !== '0') {
      map[node.parent_id].children.push(node);
    } else {
      tree.push(node);
    }
  }

  return tree;
}

// 数组对象排序
function ListObjCompare(property: string){
  return function(a: any,b: any){
    const value1 = a[property];
    const value2 = b[property];
    return value1 - value2;//升序,降序为value2 - value1
  }
}
export  {ListToTree, ListObjCompare}
