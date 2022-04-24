import * as api from './api';
import log from '@/utils/log';
import { crudSearchParam } from '@/utils/search_param';
// 构建crudOptions的方法
export default function ({ expose }) {
  const pageRequest = async (query) => {
    const params = crudSearchParam(query);
    return await api.GetList(params);
  };
  const editRequest = async ({ form, row }) => {
    form.id = row.id;
    return await api.UpdateObj(form);
  };
  const delRequest = async (id) => await api.DelObj(id);

  const addRequest = async ({ form }) => await api.AddObj(form);
  return {
    crudOptions: {
    // 请求配置
      request: {
        pageRequest, // 列表数据请求
        addRequest, // 添加请求
        editRequest, // 修改请求
        delRequest, // 删除请求
      },
      columns: {
      // 字段配置
        title: {
          title: '标题', // 字段名称
          search: { show: true }, // 搜索配置
          type: 'text', // 字段类型
        },
        type: {
          title: '类型', // 字段名称
          search: { show: true }, // 搜索配置
          type: 'text', // 字段类型
        },
        content: {
          title: '内容',
          search: { show: false },
          type: 'text',
        },
      // 你可以尝试在此处增加更多字段
      },
    // 其他crud配置
    },
  };
}
