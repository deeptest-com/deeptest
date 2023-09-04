// eslint-disable-next-line @typescript-eslint/no-var-requires
const defaultLogoPath = require('@/assets/logo/default.svg');
const req: __WebpackModuleApi.RequireContext = require.context('../../assets/logo', true, /\.svg$/);

export const projectLogoList: any[] = [];

const projectLogoMap: any = {};

req.keys().forEach(eachPng => {
    const imgConfig = req(eachPng);
    const imgName = eachPng.replace(/^\.\/(.*)\.\w+$/, '$1');
    projectLogoList.push({ imgName, icon: req(eachPng).default || imgConfig });
    projectLogoMap[imgName] = req(eachPng).default || imgConfig;
})

export function getProjectLogo (name: string | undefined | number): string {
    const logoPath = projectLogoMap[`${name}`];
    return logoPath || defaultLogoPath;
}
