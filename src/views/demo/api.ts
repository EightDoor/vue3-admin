import request from '@/utils/request';

const apiPrefix = "/demo_crud";

export function GetList(query: any) {
    return request({
        url: `${apiPrefix}${query}`,
        method: "GET",
    });
}

export function AddObj(obj) {
    return request({
        url: apiPrefix,
        method: "POST",
        body: obj,
    });
}

export function UpdateObj(obj) {
    return request({
        url: `${apiPrefix}/${obj.id}`,
        method: "PUT",
        body: obj,
    });
}

export function DelObj(id) {
    return request({
      url: `${apiPrefix}/${id}`,
        method: "DELETE",
        params: { id },
    });
}

export function GetObj(id) {
    return request({
        url: `${apiPrefix}/${id}`,
        method: "GET",
        params: { id },
    });
}
