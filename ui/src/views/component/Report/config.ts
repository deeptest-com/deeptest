const responseDataColumns = [
    { title: '操作', dataIndex: 'expandIcon', key: 'expandIcon' },
    { title: '请求状态', dataIndex: 'requestStatus', key: 'requestStatus', slots: { customRender: 'requestStatus' } },
    { title: '请求方法', dataIndex: 'requestMethod', key: 'requestMethod', slots: { customRender: 'requestMethod' } },
    { title: '请求url', dataIndex: 'requestUrl', key: 'requestUrl' },
    { title: '请求信息', dataIndex: 'requestData', key: 'requestData' },
    { title: '请求状态码', dataIndex: 'requestCode', key: 'requestCode', slots: { customRender: 'requestCode' } },
    { title: '请求耗时', dataIndex: 'requestTime', key: 'requestTime', slots: { customRender: 'requestTime' } },
    { title: '操作', dataIndex: 'operation', key: 'operation', slots: { customRender: 'operation' } },
];

const scenarioReportColumns = [
    { title: '场景名称', dataIndex: 'scenarioName', key: 'scenarioName', slots: { customRender: 'scenarioName' } },
    { title: '场景优先级', dataIndex: 'scenarioPriority', key: 'scenarioPriority', slots: { customRender: 'scenarioPriority' } },
    { title: '执行状态', dataIndex: 'scenarioStatus', key: 'scenarioStatus', slots: { customRender: 'scenarioStatus' } },
    { title: '执行通过率', dataIndex: 'scenarioProgress', key: 'scenarioProgress', slots: { customRender: 'scenarioProgress' } },
    { title: '操作', dataIndex: 'operation', key: 'operation' },
];

export {
    responseDataColumns,
    scenarioReportColumns
}