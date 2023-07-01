import {isInArray} from "@/utils/array";

export function getSelectedTreeNode(checkedKeys, treeDataMapValue): any[] {
    const childrenMap = {} // nodes that is other's child
    checkedKeys.forEach((id, index) => {
        if (treeDataMapValue[id].children) {
            treeDataMapValue[id].children.forEach((child, index) => {
                getChildren(treeDataMapValue[child.id], childrenMap)
            })
        }
    })
    console.log('childrenMap', childrenMap)

    const selectedNodes = [] as any[]

    Object.keys(treeDataMapValue).forEach((id, index) => {
        if (!childrenMap[id] && isInArray(+id, checkedKeys)) { // in array and except other's child
            const node = treeDataMapValue[id]
            if (node.isLeaf || node.children) {
                selectedNodes.push(node)
            }
        }
    })

    return selectedNodes
}
const getChildren = (node, mp) => {
    mp[node.id] = true

    if (node.children) {
        node.children.forEach((child, index) => {
            getChildren(child, mp)
        })
    }
}

export function filterTree(treeDataValue, keywords): number[] {
    if (!treeDataValue) return []

    const flattenTreeList = flattenTree(treeDataValue[0]);

    let parentKeys: any = [];
    for (let i = 0; i < flattenTreeList.length; i++) {
        const node = flattenTreeList[i];

        const text = node.title?node.title:node.name
        if (text.includes(keywords)) {
            parentKeys.push(node.parentId);
            parentKeys = parentKeys.concat(findParentIds(node.parentId, flattenTreeList));
        }
    }
    parentKeys = [...new Set(parentKeys)];

    return parentKeys
}

function flattenTree(tree) {
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

function findParentIds(nodeId, tree) {
    let current: any = tree.find(node => node.id === nodeId);
    const parentIds: Array<string> = [];
    while (current && current.parentId) {
        parentIds.unshift(current.parentId); // unshift 方法可以将新元素添加到数组的开头
        current = tree.find(node => node.id === current.parentId);
    }
    return parentIds;
}
