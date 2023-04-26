const req: __WebpackModuleApi.RequireContext = require.context('../../assets/logo', true, /\.png$/);
const pngHashMap: any[] = [];
req.keys().forEach(eachPng => {
    const imgConfig = req(eachPng);
    const imgName = eachPng.replace(/^\.\/(.*)\.\w+$/, '$1');
    pngHashMap.push({ imgName, icon: req(eachPng).default || imgConfig });
})


export default pngHashMap;