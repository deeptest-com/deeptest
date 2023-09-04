import request from '@/utils/request';
import {LoginParamsType} from "@/views/user/login/data";

export async function getProfile(): Promise<any> {
    return request({
        url: '/users/profile',
        method: 'get'
    });
}

export async function queryMessage(): Promise<any> {
    return request({
        url: '/users/message'
    });
}

export async function queryProject(): Promise<any> {
    return request({
        url: '/projects/getByUser'
    });
}

export async function updateEmail(email): Promise<any> {
    return request({
        url: '/users/updateEmail',
        method: 'POST',
        data: {email: email}
    });
}
export async function updateName(username): Promise<any> {
    return request({
        url: '/users/updateName',
        method: 'POST',
        data: {username: username}
    });
}
export async function updatePassword(data): Promise<any> {
    return request({
        url: '/users/updatePassword',
        method: 'POST',
        data
    });
}

