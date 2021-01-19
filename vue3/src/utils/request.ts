import axios, { Method } from 'axios'
import { message, notification } from 'ant-design-vue'
import { CommonResponse } from '@/types/type'
import { RequestAuthorizedFailed, TOKEN } from '@/utils/constant'
import { ClearInfo } from '@/utils/index'

const instance = axios.create({
  baseURL: '/api',
  timeout: 5000,
})

// 请求拦截器
instance.interceptors.request.use(
  function (config) {
    config.headers = {
      token: localStorage.getItem(TOKEN),
    }
    return config
  },
  function (error) {
    return Promise.reject(error)
  }
)

// 相应拦截器
instance.interceptors.response.use(
  function (response) {
    const data = response.data
    if (response.status === 200 && data.code === 0) {
      //
    } else if (data.code === RequestAuthorizedFailed) {
      message.info('token失效, 请重新登录')
      ClearInfo()
    } else {
      notification.error({
        message: data.code,
        description: data.message,
      })
    }
    return response.data
  },
  function (error) {
    return Promise.reject(error)
  }
)

interface HttpCustomType {
  url: string
  method: Method
  body?: unknown
  params?: unknown
}
function httpCustom<T>(c: HttpCustomType): Promise<CommonResponse<T>> {
  return new Promise((resolve, reject) => {
    instance({
      url: c.url,
      method: c.method,
      data: c.body,
      params: c.params,
    })
      .then((res) => {
        // @ts-ignore
        if (res.code !== 0) {
          // @ts-ignore
          reject(res.code)
        } else {
          resolve(res.data)
        }
      })
      .catch((err) => {
        reject(err)
      })
  })
}

export { httpCustom as http }
