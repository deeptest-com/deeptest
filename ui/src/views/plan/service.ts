import request from '@/utils/request';
import {QueryParams} from "@/views/project/data";
import {getEnumSelectItems} from "@/views/interface/service";
import {
    ProcessorCookie, ProcessorData,
    ProcessorExtractor,
    ProcessorLogic,
    ProcessorLoop, ProcessorGroup, ProcessorTimer, ProcessorPrint,
    ProcessorCategory,
    ProcessorVariable, ProcessorAssertion, RequestBodyType
} from "@/utils/enum";
import {Interface} from "@/views/interface/data";

const apiPath = 'plans';
const apiPathCategoryNodes = `${apiPath}/categories`;

export async function query(params?: QueryParams): Promise<any> {
    return request({
        url: `/${apiPath}`,
        method: 'get',
        params,
    });
}
export async function get(id: number): Promise<any> {
    return request({url: `/${apiPath}/${id}`});
}
export async function save(data: any): Promise<any> {
    return request({
        url: `/${apiPath}`,
        method: data.id? 'PUT': 'POST',
        data: data,
    });
}
export async function remove(id: number): Promise<any> {
    return request({
        url: `/${apiPath}/${id}`,
        method: 'delete',
    });
}

// category tree
export async function loadCategory(): Promise<any> {
    const params = {}
    return request({
        url: `/${apiPathCategoryNodes}/load`,
        method: 'get',
        params,
    });
}
export async function getCategory(id: number): Promise<any> {
    return request({url: `/${apiPathCategoryNodes}/${id}`});
}
export async function createCategory(data): Promise<any> {
    return request({
        url: `/${apiPathCategoryNodes}`,
        method: 'POST',
        data: data,
    });
}
export async function updateCategory(id: number, data: any): Promise<any> {
    return request({
        url: `/${apiPathCategoryNodes}/${id}`,
        method: 'PUT',
        data: data,
    });
}
export async function updateCategoryName(id: number, name: string): Promise<any> {
    const data = {id: id, name: name}

    return request({
        url: `/${apiPathCategoryNodes}/${id}/updateName`,
        method: 'PUT',
        data: data,
    });
}
export async function removeCategory(id: number): Promise<any> {
    return request({
        url: `/${apiPathCategoryNodes}/${id}`,
        method: 'delete',
    });
}
export async function moveCategory(data: any): Promise<any> {
    return request({
        url: `/${apiPathCategoryNodes}/move`,
        method: 'post',
        data: data,
    });
}
