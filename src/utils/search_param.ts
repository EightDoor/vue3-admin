import log from '@/utils/log';
import { RequestQueryBuilder } from '@nestjsx/crud-request';
import { CreateQueryParams } from '@nestjsx/crud-request/lib/interfaces';

interface ParamsType {
  page: {
    currentPage: number;
    pageSize: number;
  };
  sort: any;
  form: any
}
/**
 * 接口请求 查询条件
 * @param params
 */
const searchParam = (params: CreateQueryParams) => `?${RequestQueryBuilder.create(params).query()}`;

/**
 * crud 请求
 * @param params
 * @returns
 */
const crudSearchParam = (params: ParamsType) => {
  log.d(params, 'crud请求参数')
  const result:CreateQueryParams = {
    limit: params.page.pageSize ?? 10,
    page: params.page.currentPage ?? 1
  }
  return `?${RequestQueryBuilder.create(result).query()}`;
}
export { searchParam, crudSearchParam };
