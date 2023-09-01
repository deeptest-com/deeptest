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
        console.log(!childrenMap[id],isInArray(id, checkedKeys))
        if (!childrenMap[id] && isInArray(id, checkedKeys)) { // in array and except other's child
            const node = treeDataMapValue[id]
            if (!node.isDir || node.children) {
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

        const text = node.title ? node.title : node.name;
        // 兼容大小写问题
        if (text.toLowerCase().includes(keywords.toLowerCase().trim())) {
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


/**
 * @desc 根据关键词过滤树节点
 * @param {Array} children 树节点
 * @param {String} keyword 关键词
 * @param {String} field 搜索字段
 * @return {Array} 过滤后的树节点
 *
 * */
export function filterByKeyword(children, keyword, field = 'title') {
    if (!keyword.trim()) return children;

    function filterChildren(node) {
        if (node?.children?.length) {
            node.children = node.children.filter((child) => {
                return filterChildren(child);
            })
        }
        return hasChildrenByKeyword(node, keyword, field);
    }

    return children.filter((menu) => {
        return filterChildren(menu);
    })
}

/**
 * @desc 该节点下是否包含关键词
 * @param {Object} node 节点
 * @param {String} keyword 关键词
 * @param {String} field 搜索字段
 * @return {Boolean} 是否包含关键词
 * */
function hasChildrenByKeyword(node, keyword, field = 'title') {
    let result = false;

    // 定义递归函数，用于遍历树节点
    function traverse(node) {
        if (node?.[field]?.toLowerCase()?.includes(keyword.toLowerCase().trim())) {
            result = true;
            return;
        }
        // 递归处理子节点
        if (node?.children?.length > 0) {
            for (const child of node.children) {
                traverse(child);
            }
        }
    }

    // 调用递归函数，开始遍历
    traverse(node);
    return result;
}

