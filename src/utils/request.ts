import axios, { Method } from 'axios';
import { message, notification } from 'ant-design-vue';
import { CommonResponse } from '@/types/type';
import { RequestAuthorizedFailed, TOKEN } from '@/utils/constant';
import { ClearInfo } from '@/utils/index';

const instance = axios.create({
  baseURL: '/api',
  timeout: 5000,
});
// 请求拦截器
instance.interceptors.request.use(
  (config) => {
    config.headers = {
      Authorization: `Bearer ${localStorage.getItem(TOKEN) ?? ''}`,
    };
    return config;
  },
  (error) => Promise.reject(error),
);

// 相应拦截器
instance.interceptors.response.use(
  (response) => {
    const { data } = response;
    if (response.status === 200 && data.code === 0) {
      //
    } else if (data.code === RequestAuthorizedFailed) {
      message.info('token失效, 请重新登录');
      ClearInfo();
    } else {
      notification.error({
        message: data.code,
        description: data.message,
      });
    }
    return response.data;
  },
  (error) => Promise.reject(error),
);

interface HttpCustomType {
  url: string;
  method: Method;
  body?: unknown;
  params?: unknown;
}
function httpCustom<T = any>(c: HttpCustomType): Promise<CommonResponse<T>> {
  return new Promise((resolve, reject) => {
    instance({
      url: c.url,
      method: c.method,
      data: c.body,
      params: c.params,
    })
      .then((res: any) => {
        if (res.code !== 0) {
          reject(res.data);
        } else if (res.data) {
          resolve({
            data: res.data,
            list: [],
          });
        } else if (res.data.list) {
          resolve({
            list: res.data.list,
            data: null,
          });
        }
      })
      .catch((err) => {
        reject(err);
      });
  });
}

export default httpCustom;
