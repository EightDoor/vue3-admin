function ListToTree<T>(jsonData: T[], id = 'id', pid = 'parent_id'): T[] {
  const result: T[] = [],
    temp = {}
  for (let i = 0; i < jsonData.length; i++) {
    temp[jsonData[i][id]] = jsonData[i] // 以id作为索引存储元素，可以无需遍历直接定位元素
  }
  for (let j = 0; j < jsonData.length; j++) {
    const currentElement = jsonData[j]
    const tempCurrentElementParent = temp[currentElement[pid]] // 临时变量里面的当前元素的父元素
    if (tempCurrentElementParent) {
      // 如果存在父元素
      if (!tempCurrentElementParent['children']) {
        // 如果父元素没有chindren键
        tempCurrentElementParent['children'] = [] // 设上父元素的children键
      }
      tempCurrentElementParent['children'].push(currentElement) // 给父元素加上当前元素作为子元素
    } else {
      // 不存在父元素，意味着当前元素是一级元素
      result.push(currentElement)
    }
  }
  return result
}

// 数组对象排序
function ListObjCompare<T>(property: string) {
  return function (a: T, b: T): number {
    const value1 = a[property]
    const value2 = b[property]
    return value1 - value2 //升序,降序为value2 - value1
  }
}

// 清空token
const ClearInfo = (): void => {
  localStorage.clear()
}
export { ListToTree, ListObjCompare, ClearInfo }
