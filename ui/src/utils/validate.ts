/**
 * 判断是否是外链
 * @param {string} path
 * @returns {Boolean}
 * @author LiQingSong
 */
export const isExternal = (path: string): boolean => {
  return /^(https?:|mailto:|tel:)/.test(path);
};

export function urlValidator(...arg: any[]) {
  const value = arg[1];
  // 需要兼容 https://localhost
  const urlReg = /http(s)?:\/\/[\w-]+/; // eslint-disable-line
  if (value === '') {
    return Promise.reject('url不能为空');
  } else {
    if (!urlReg.test(value)) {
      return Promise.reject('url格式错误，请参考 http(s)://www.test.com ');
    }
    return Promise.resolve();
  }
}