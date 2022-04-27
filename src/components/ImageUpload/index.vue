<template>
  <div>
    <a-upload
      :action="Config.qiniuUploadUrl"
      :multiple="true"
      v-model:file-list="fileList"
      :data="extendedParam"
      list-type="picture"
      @change="handleChange"
    >
      <a-button>
        <upload-outlined></upload-outlined>
        上传
      </a-button>
    </a-upload>
  </div>
</template>
<script lang="ts">
import { UploadOutlined } from '@ant-design/icons-vue';
import {
  defineComponent,
  ref,
  onMounted,
  PropType,
  reactive,
  markRaw,
  watch,
} from 'vue';
import { message } from 'ant-design-vue';
import Config from '@/config';
import log from '@/utils/log';
import http from '@/utils/request';
import { FileItem } from '@/types';

interface FileInfo {
  file: FileItem;
  fileList: FileItem[];
  event: any;
}

export default defineComponent({
  components: {
    UploadOutlined,
  },
  props: {
    list: {
      type: Array as PropType<Array<FileItem>>,
      required: true,
      default: () => [],
    },
    limit: {
      type: Number,
      default: 10,
    },
  },
  setup(props, { emit }) {
    const fileList = ref<FileItem[]>([]);
    const isFirst = ref(false);
    const extendedParam = reactive({
      token: '',
    });

    function delFile(key: string) {
      http({
        url: 'upload/del',
        method: 'POST',
        body: key,
      }).then((res) => {
        log.d(res, '删除文件结果');
        if (!res.data) {
          message.success('删除成功');
        } else {
          message.error(JSON.stringify(res.data));
        }
      });
    }

    const handleChange = (info: FileInfo) => {
      log.d(info, '图片文件改变');
      let resFileList = [...info.fileList];

      resFileList = resFileList.map((file) => {
        if (file.status === 'done') {
          if (file.response) {
            file.url = `${Config.qiniuShowUrl}/${file.response.key}`;
          }
        }
        return file;
      });

      const { file } = info;
      if (file.status === 'removed') {
        if (file.url) {
          const { url } = file;
          file.url = '';
          // 读取文件key 删除
          const key = url.substring(url.lastIndexOf('/') + 1);
          delFile(key);
        }
      }

      fileList.value = markRaw(resFileList);
      const result = resFileList.filter((item) => item.url);
      emit('update:list', result);
    };

    function getFileToken() {
      http({
        url: 'upload/zk',
        method: 'GET',
      }).then((res) => {
        extendedParam.token = res.data;
      });
    }
    onMounted(() => {
      getFileToken();
    });

    watch(
      () => props.list,
      (newVal) => {
        if (newVal && !isFirst.value) {
          fileList.value = markRaw(newVal);
          isFirst.value = true;
        }
      },
      {
        deep: true,
        immediate: true,
      },
    );

    return {
      handleChange,

      Config,
      extendedParam,
      fileList,
    };
  },
});
</script>
<style scoped lang="less"></style>
