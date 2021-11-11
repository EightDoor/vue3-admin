/** *
 * 日志
 */
export default {
  i(data: any, title?: string) {
    console.group(title ?? 'info');
    console.log(data);
    console.groupEnd();
  },
  d(data: any, title?: string) {
    console.group(title ?? 'debug');
    console.log(data);
    console.groupEnd();
  },
  e(data: any, title?: string) {
    console.group(title ?? 'error');
    console.error(data);
    console.groupEnd();
  },
  s(data: any, title?: string) {
    console.group(title ?? 'success');
    console.log(data);
    console.groupEnd();
  },
};
