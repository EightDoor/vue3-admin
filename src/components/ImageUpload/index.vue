<template>
 <div>
   <a-upload
       :action="Config.qiniuUploadUrl"
       :multiple="true"
       :file-list="list"
       show-file-list
      :headers="headers"
       @change="handleChange"
       :on-remove="onRemove"
       :on-success="onSuccess"
       :on-error="onError"
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
import { defineComponent, ref, onMounted } from 'vue';
import Config from '@/config';
import log from '@/utils/log';
import http from '@/utils/request';

interface FileItem {
  uid: string;
  name?: string;
  status?: string;
  response?: Response;
  url: string;
}

interface FileInfo {
  file: FileItem;
  fileList: FileItem[];
}

export default defineComponent({
  components: {
    UploadOutlined,
  },
  setup() {
    const list = ref<FileItem[]>([
      {
        uid: '-1',
        name: 'xxx.png',
        status: 'done',
        url: 'http://www.baidu.com/xxx.png',
      },
    ]);
    const headers: Object = {
      token: '',
    };
    const handleChange = (info: FileInfo) => {
      let resFileList = [...info.fileList];

      resFileList = resFileList.slice(-2);

      resFileList = resFileList.map((file) => {
        if (file.response) {
          // Component will show file.url as link
          file.url = file.response.url;
        }
        return file;
      });

      list.value = resFileList;
    };

    function onRemove(file: FileItem, fileList:FileItem[]) {
      log.d(file);
      log.d(fileList);
    }

    function onSuccess(response, file:FileItem, fileList:FileItem[]) {
      log.d(response);
      log.d(file);
      log.d(fileList);
    }

    function onError(err, file:FileItem, fileList:FileItem[]) {
      log.d(err);
      log.d(file);
      log.d(fileList);
    }

    function getFileToken() {
      http({
        url: '/upload/zk',
        method: 'GET',
      }).then((res) => {
        log.i(res);
      });
    }
    onMounted(() => {
      getFileToken();
    });

    return {
      list,
      handleChange,

      Config,

      onRemove,
      onSuccess,
      onError,
      headers,
    };
  },
});
</script>
<style scoped lang="less">

</style>
