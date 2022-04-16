import request from '@/utils/request';
import {Interface} from "@/views/interface/data";
import {isInArray} from "@/utils/array";

const apiPath = 'interfaces';

export async function test(interf: Interface): Promise<any> {
    return request({
        url: `/${apiPath}/test`,
        method: 'post',
        data: interf,
    });
}

export async function load(): Promise<any> {
    return request({
        url: `/${apiPath}`,
        method: 'get',
    });
}

export async function get(id: number): Promise<any> {
    return request({url: `/${apiPath}/${id}`});
}

export async function create(data): Promise<any> {
    return request({
        url: `/${apiPath}`,
        method: 'POST',
        data: data,
    });
}

export async function update(id: number, params: any): Promise<any> {
    return request({
        url: `/${apiPath}/${id}`,
        method: 'PUT',
        data: params,
    });
}

export async function remove(id: number): Promise<any> {
    return request({
        url: `/${apiPath}/${id}`,
        method: 'delete',
    });
}

export async function move(data: any): Promise<any> {
    return request({
        url: `/${apiPath}/move`,
        method: 'post',
        data: data,
    });
}

export function clearDataForRequest(data: any) {
    data.params = data.params.filter((item) => {
        return !!item.name
    })
    data.headers = data.headers.filter((item) => {
        return !!item.name
    })

    return data
}

export function getNodeMap(treeNode: any, mp: any): void {
    if (!treeNode) return

    mp[treeNode.id] = treeNode

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

export function getCodeLangByType(type) {
    type = type.split('/')[1]

    if (isInArray(type, ['json', 'xml', 'html'])) {
        return type
    } else {
        return 'plaintext'
    }
}
