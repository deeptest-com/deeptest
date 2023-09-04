export function formatData(data: Array<any>) {
    if (!data) return [];
    let result: any[] = [];
    data.forEach(e => {
        if (e.processorCategory === 'processor_interface') {
            result.push(e);
        } else if (e.processorCategory !== 'processor_interface') {
            result = [ ...result, ...formatData(e.logs) ];
        }
    });
    return result;
}