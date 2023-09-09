/**
 *  打包ly的包时，需要修改下面的引用 package 路径
 * */
// import pkg from '../package.json';
const pkg = require('../package-ly.json');
/**
 * 运行时配置对象
 * @type {{pkg: Object, media: Object, system: Object}}
 */
const config = {pkg};

/**
 * 更新运行时配置对象
 * @param {Object} newConfig 新的配置对象
 * @return {void}
 */
export function updateConfig(newConfig) {
    Object.assign(config, newConfig);
}

export default config;
