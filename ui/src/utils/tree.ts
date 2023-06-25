export function flattenTree(tree) {
    const nodes: Array<any> = [];

    function traverse(node) {
        nodes.push(node);
        if (node.children) {
            node.children.forEach(traverse);
        }
    }

    traverse(tree);
    return nodes;
}

export function findParentIds(nodeId, tree) {
    let current: any = tree.find(node => node.id === nodeId);
    const parentIds: Array<string> = [];
    while (current && current.parentId) {
        parentIds.unshift(current.parentId); // unshift 方法可以将新元素添加到数组的开头
        current = tree.find(node => node.id === current.parentId);
    }
    return parentIds;
}

export function filterTree (treeDataValue, keywords) : number[]{
    if (!treeDataValue) return []

    const flattenTreeList = flattenTree(treeDataValue[0]);

    let parentKeys: any = [];
    for (let i = 0; i < flattenTreeList.length; i++) {
        const node = flattenTreeList[i];

        if (node.title.includes(keywords)) {
            parentKeys.push(node.parentId);
            parentKeys = parentKeys.concat(findParentIds(node.parentId, flattenTreeList));
        }
    }
    parentKeys = [...new Set(parentKeys)];

    return parentKeys
}
