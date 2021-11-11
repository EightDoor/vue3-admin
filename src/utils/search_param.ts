import { RequestQueryBuilder } from '@nestjsx/crud-request';
import { CreateQueryParams } from '@nestjsx/crud-request/lib/interfaces';

/**
 * 接口请求 查询条件
 * @param params
 */
const searchParam = (params: CreateQueryParams) => `?${RequestQueryBuilder.create(params).query()}`;

export { searchParam };
