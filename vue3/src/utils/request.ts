import axios, { Method } from 'axios'
import { notification } from 'ant-design-vue'
import { CommonResponse } from '@/types/type'

const instance = axios.create({
  baseURL: '/api',
  timeout: 5000,
})

// 请求拦截器
instance.interceptors.request.use(function(config){
  return config
}, function(error){
  return Promise.reject(error)
})

// 相应拦截器
instance.interceptors.response.use(function(response) {
  const data = response.data;
  if(response.status === 200 && data.code === 0) {
    //
  }else {
    notification.error({
      message: data.code,
      description: data.data
    })
  }
  return response.data.data
}, function(error){
  return Promise.reject(error)
})


interface httpCustomType {
  url: string;
  method: Method
  data?: any,
  params?: any
}
function httpCustom<T> (c: httpCustomType):Promise<CommonResponse<T>>{
  return new Promise((resolve, reject)=>{
    instance({
      url: c.url,
      method: c.method,
      data: c.data,
      params: c.params
    }).then((res: any)=>{
      resolve(res);
    }).catch((err)=>{
      reject(err)
    })
  })
}


export {httpCustom as http}


