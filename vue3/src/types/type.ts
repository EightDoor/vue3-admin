import { AxiosResponse } from 'axios'

export interface CommonResponse<T> extends Pagin{
  list: T[],
}
interface Pagin {
  page?: number,
  pageSize?: number
}

export interface CommonResponseSing<T> extends Pagin{
  list: T,
}
