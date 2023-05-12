export const reportData = {
    logName: '测试计划名称11',
    logExecutor: '测试执行人',
    logEnv: '测试环境',
    logTime: '2023/04/24 10:20:20',
    logList: [
        {
            id: 1,
            scenarioName: '测试场景1',
            scenarioPriority: 'P1',
            scenarioStatus: 0,
            scenarioProgress: '66.66',
            reponseList: [
                {
                    requestId: 44444,
                    requestStatus: 'loading',
                    requestMethod: 'GET',
                    requestCode: '200',
                    requestUrl: '/pet/%khkhfhw h',
                    requestData: '查询的具体信息',
                    requestTime: '100ms',
                    requestInfo: [
                        {
                            errorId: 122,
                            errorField: '接口数据有问题',
                            errorTip: ['问题1的描述', '问题2的描述']
                        }
                    ]
                },
                {
                    requestId: 55555,
                    requestStatus: 'success',
                    requestMethod: 'GET',
                    requestCode: '200',
                    requestUrl: '/pet/%khkhfhw h',
                    requestData: '查询的具体信息',
                    requestTime: '100ms',
                    requestInfo: [
                        {
                            errorId: 112,
                            errorField: '接口数据有问题',
                            errorTip: ['问题1的描述', '问题2的描述']
                        }
                    ]
                },
                {
                    requestId: 66666,
                    requestStatus: 'error',
                    requestMethod: 'POST',
                    requestCode: '400',
                    requestUrl: '/pet/%khkhfhw h',
                    requestData: '查询的具体信息',
                    requestTime: '100ms',
                    requestInfo: [
                        {
                            errorId: 172,
                            errorField: '接口数据有问题',
                            errorTip: ['问题1的描述', '问题2的描述']
                        }
                    ]
                }
            ]
        },
        {
            id: 2,
            scenarioName: '测试场景2',
            scenarioPriority: 'P3',
            scenarioStatus: 0,
            scenarioProgress: '58',
            reponseList: [
                {
                    requestId: 11111,
                    requestStatus: 'success',
                    requestMethod: 'GET',
                    requestCode: '200',
                    requestUrl: '/pet/%khkhfhw h',
                    requestData: '查询的具体信息',
                    requestTime: '100ms',
                    requestInfo: [
                        {
                            errorId: 142,
                            errorField: '接口数据有问题',
                            errorTip: ['问题1的描述', '问题2的描述']
                        }
                    ]
                },
                {
                    requestId: 2222,
                    requestStatus: 'error',
                    requestMethod: 'GET',
                    requestCode: '200',
                    requestUrl: '/pet/%khkhfhw h',
                    requestData: '查询的具体信息',
                    requestTime: '100ms',
                    requestInfo: [
                        {
                            errorId: 145,
                            errorField: '接口数据有问题',
                            errorTip: ['问题1的描述', '问题2的描述']
                        }
                    ]
                },
                {
                    requestId: 33333,
                    requestStatus: 'expires',
                    requestMethod: 'POST',
                    requestCode: '400',
                    requestUrl: '/pet/%khkhfhw h',
                    requestData: '查询的具体信息',
                    requestTime: '100ms',
                    requestInfo: [
                        {
                            errorId: 146,
                            errorField: '接口数据有问题',
                            errorTip: ['问题1的描述', '问题2的描述']
                        }
                    ]
                }
            ]
        }
    ]
};