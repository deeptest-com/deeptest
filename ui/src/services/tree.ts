
export function genNodeMap(treeNode: any, ids?: number[]): any {
     const mp = {}
    getNodeMap(treeNode, mp, ids)

    return mp
}

export function getNodeMap(treeNode: any, mp: any, ids?: number[]): void {
    if (!treeNode) return

    mp[treeNode.id] = treeNode
    if (ids && treeNode.entityCategory !== 'processor_group') {
        ids.push(treeNode.id)
        // console.log('===', treeNode.entityCategory)
    }

    if (treeNode.children) {
        treeNode.children.forEach((item, index) => {
            getNodeMap(item, mp, ids)
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

    if (treeMap[key]) {
        const parentId = treeMap[key].parentId
        if (parentId) {
            expandOneKey(treeMap, parentId, expandedKeys)
        }
    }
}