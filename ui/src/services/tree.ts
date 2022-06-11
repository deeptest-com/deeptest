
export function getNodeMap(treeNode: any, mp: any): void {
    if (!treeNode) return

    mp[treeNode.id] = treeNode
    // console.log('treeNode.id', treeNode.id)

    if (treeNode.children) {
        treeNode.children.forEach((item, index) => {
            getNodeMap(item, mp)
        })
    }

    return
}

export function expandAllKeys(treeMap: any, isExpand: boolean): number[] {
    const keys = new Array<number>()
    if (!isExpand) return keys

    Object.keys(treeMap).forEach(key => {
        if (!keys.includes(+key)) keys.push(+key)
    })

    return keys
}

export function expandOneKey(treeMap: any, key: number, expandedKeys: number[]) {
    if (!expandedKeys.includes(key)) expandedKeys.push(key)

    const parentId = treeMap[key].parentId
    if (parentId) {
        expandOneKey(treeMap, parentId, expandedKeys)
    }
}