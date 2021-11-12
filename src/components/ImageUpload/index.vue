<template>
  <a-upload
      :action="Config.qiniuUploadUrl"
      :multiple="true"
      :file-list="fileList"
      @change="handleChange"
  >
    <a-button>
      <upload-outlined></upload-outlined>
      Upload
    </a-button>
  </a-upload>
</template>
<script lang="ts">
import { UploadOutlined } from '@ant-design/icons-vue';
import { defineComponent, ref } from 'vue';
import Config from '@/config';

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
    const fileList = ref<FileItem[]>([
      {
        uid: '-1',
        name: 'xxx.png',
        status: 'done',
        url: 'http://www.baidu.com/xxx.png',
      },
    ]);
    const handleChange = (info: FileInfo) => {
      let resFileList = [...info.fileList];

      // 1. Limit the number of uploaded files
      //    Only to show two recent uploaded files, and old ones will be replaced by the new
      resFileList = resFileList.slice(-2);

      // 2. read from response and show file link
      resFileList = resFileList.map((file) => {
        if (file.response) {
          // Component will show file.url as link
          file.url = file.response.url;
        }
        return file;
      });

      fileList.value = resFileList;
    };
    return {
      fileList,
      handleChange,
    };
  },
});
</script>
<style scoped lang="less">

</style>
