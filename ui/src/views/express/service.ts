import request from '@/utils/request';
import {requestToAgent} from '@/utils/request';
import {Interface, OAuth20} from "@/views/interface/data";
import {isInArray} from "@/utils/array";

const apiSpec = 'spec';
const apiImport = 'import';

export async function loadSpec(data): Promise<any> {
    return requestToAgent({
        url: `/${apiSpec}/loadSpec`,
        method: 'POST',
        data,
    });
}

// export async function importSpec(data, targetId): Promise<any> {
//     return request({
//         url: `/${apiImport}/importSpec`,
//         method: 'POST',
//         params: {targetId: targetId, type: data.type},
//         data: data,
//     });
// }