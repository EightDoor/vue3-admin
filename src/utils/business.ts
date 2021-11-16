import { FileItem } from '@/types';

const BusinessUtils = {
  /**
   * ImageUpload 图片提交数据格式化
   * @param data
   */
  formatUploadImg(data: FileItem[]) {
    const images:string[] = [];
    data.forEach((item) => {
      images.push(item.url);
    });
    if (images.length > 0) {
      return images.join(',');
    }
    return '';
  },
  /**
   * ImageUpload 图片回显列表
   * @param img
   */
  formatUploadShow(img: string): FileItem[] {
    if (img) {
      return [{
        thumbUrl: img,
        url: img,
        uid: String(Date.now()),
      }];
    }
    return [];
  },
};

export default BusinessUtils;
