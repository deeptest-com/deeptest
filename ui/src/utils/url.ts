export function getUrlKey(name, url){
    const regx = new RegExp('[?|&]' + name + '=' + '([^&;]+?)(&|#|;|$)') as any
    // eslint-disable-next-line no-sparse-arrays
    return decodeURIComponent((regx.exec(url) || [, ""])[1].replace(/\+/g, '%20')) || null
}