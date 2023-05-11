import moment from "moment";

export function momentUtc(dt) {
    return moment.parseZone(dt).format("YYYY-MM-DD HH:mm:ss")
}
export function momentShort(dt) {
    return moment.parseZone(dt).format("MM-DD HH:mm:ss")
}
export function momentTime(dt) {
    return moment.parseZone(dt).format("HH:mm:ss")
}

export function momentUnixFormat(tm, format) {
    return moment.unix(tm).format(format)
}

export function momentTimeStamp(tm) {
    return moment(tm).valueOf();
}

export function percentDef(numb, total) {
    if (total == 0) return '0%'
    return Number(numb / total * 100).toFixed(2) + '%'
}