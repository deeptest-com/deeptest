import {getCache, setCache} from './localCache';
import settings from '@/config/settings';

export const getShowRightBar = async (): Promise<boolean> => {
    const mp = await getCache(settings.settings);

    const showRightBar = mp ? mp[settings.showRightBar] : false
    return showRightBar
}
export const setShowRightBar = async (val) => {
    let mp = await getCache(settings.showRightBar) as any;
    if (!mp) mp = {}

    mp[settings.showRightBar] = val
    await setCache(settings.settings, mp);
}

