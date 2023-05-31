import moment from "moment";

export function getPercent(numb, total) {
    numb = numb || 0;
    total = total || 0;
    if (total == 0) return '0.00%';
    return Number(numb / total * 100).toFixed(2) + '%';
}

/**
 * 获取平均值,保留两位小数
 * @param array {Array} 数组
 * */
export function getAverage(array) {
    if(!array || array.length === 0) return 0.00;
    let total = 0;
    for (let i = 0; i < array.length; i++) {
        total += array[i];
    }
    return (total / array.length).toFixed(2);
}

/**
 * 除法运算,保留两位小数
 * @param arg1 {Number} 被除数
 * @param arg2 {Number} 除数
 * */
export function getDivision(arg1:number, arg2:number) {
    if (arg2 == 0) return 0.00;
    return Number(arg1 / arg2).toFixed(2);
}


export function formatWithSeconds(num) {
    if (num === 0) return `0 <span style="color: rgba(0, 0, 0, 0.85)">ms</span>`;
    if (num > 1000) return `${Number(num / 1000).toFixed(2)} <span  style="color: rgba(0, 0, 0, 0.85)">s</span>`;
    return `${(num && num.toFixed(2)) || 0} <span style="color: rgba(0, 0, 0, 0.85)">ms</span>`
}
