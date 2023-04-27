const req: __WebpackModuleApi.RequireContext = require.context('../../assets/logo', true, /\.svg$/);
const projectLogoList: any[] = [];
const projectLogoMap: any = {};
req.keys().forEach(eachPng => {
    const imgConfig = req(eachPng);
    const imgName = eachPng.replace(/^\.\/(.*)\.\w+$/, '$1');
    projectLogoList.push({ imgName, icon: req(eachPng).default || imgConfig });
    projectLogoMap[imgName] = req(eachPng).default || imgConfig;
})


export {
    projectLogoList,
    projectLogoMap
};