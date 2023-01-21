import {getCache, setCache} from './localCache';
import settings from '@/config/settings';

export const getExpandedKeys = async (type, id) => {
    console.log('getExpandedKeys')
    const key = `${type}-${id}`

    const cachedData = await getCache(settings.expandedKeys);
    if (!cachedData || !cachedData[key]) {
        return []
    }

    const keys = cachedData[key] ? cachedData[key] : []

    return [...keys]
}

export const setExpandedKeys = async (type, id, keys) => {
    console.log('setExpandedKeys')
    const key = `${type}-${id}`

    let cachedData = await getCache(settings.expandedKeys);
    if (!cachedData) cachedData = {}

    const items = []as any[]
    keys.forEach((item) => {
        items.push(item)
    })
    cachedData[key] = items
    await setCache(settings.expandedKeys, cachedData);
}
