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

export const detailResult = {
    "code": 0,
    "msg": "请求成功",
    "data": {
        "id": 1,
        "name": "计划222", //计划名
        "desc": "saaaaaa",
        "progressStatus": "",
        "resultStatus": "pass",
        "startTime": "2023-05-11T15:24:04+08:00", //执行开始时间
        "endTime": "2023-05-11T15:25:04+08:00", //执行结束时间
        "duration": 60, //执行耗时（单位：s)
        "totalScenarioNum": 0, //场景总数
        "passScenarioNum": 0, //通过场景数
        "failScenarioNum": 0, //失败场景数
        "totalInterfaceNum": 0, //接口总数
        "passInterfaceNum": 0,
        "failInterfaceNum": 0,
        "totalRequestNum": 0,
        "passRequestNum": 0,
        "failRequestNum": 0,
        "totalAssertionNum": 0, //检查点总数
        "passAssertionNum": 0, //通过检查点数
        "failAssertionNum": 0, //失败检查点数
        "InterfaceStatusMap": null,
        "payload": "",
        "planId": 6,
        "projectId": 3,
        "createUserId": 1,
        "createUserName": "", //执行人
        "serialNumber": "test01-TR-01",
        "execEnv": "127.0.0.1", //执行环境
        "priority": "", //优先级
        "scenarioReports": [ //场景报告列表
            {
                "id": 15,
                "createdAt": "2023-05-05T17:03:06+08:00",
                "updatedAt": "2023-05-05T17:03:06+08:00",
                "name": "开发03", //场景名
                "desc": "",
                "progressStatus": "",
                "resultStatus": "fail",
                "startTime": "2023-05-05T17:03:06+08:00",
                "endTime": "2023-05-05T17:03:06+08:00",
                "duration": 0,
                "totalInterfaceNum": 1,
                "passInterfaceNum": 0,
                "failInterfaceNum": 1,
                "totalRequestNum": 1,
                "passRequestNum": 0,
                "failRequestNum": 1,
                "totalAssertionNum": 0, //断言总数
                "passAssertionNum": 0, //断言通过数
                "failAssertionNum": 0, //断言失败数
                "InterfaceStatusMap": null,
                "payload": "",
                "scenarioId": 3,
                "projectId": 3,
                "planReportId": 1,
                "createUserId": 1,
                "createUserName": "",
                "serialNumber": "",
                "logs": [
                    {
                        "id": 18,
                        "createdAt": "2023-05-05T17:03:06+08:00",
                        "updatedAt": "2023-05-05T17:03:06+08:00",
                        "name": "开发03",
                        "desc": "",
                        "progressStatus": "",
                        "resultStatus": "pass",
                        "startTime": null,
                        "endTime": null,
                        "parentId": 0,
                        "reportId": 15,
                        "useId": 0,
                        "processorCategory": "processor_root",
                        "interfaceId": 0,
                        "httpStatusCode": 0,
                        "httpStatusContent": "",
                        "processorType": "processor_root_default",
                        "logs": [ //接口信息
                            {
                                "id": 19,
                                "createdAt": "2023-05-05T17:03:06+08:00",
                                "updatedAt": "2023-05-05T17:03:06+08:00",
                                "name": "接口383 - GET", //接口名称
                                "desc": "",
                                "progressStatus": "",
                                "resultStatus": "fail",
                                "startTime": null,
                                "endTime": null,
                                "parentId": 18,
                                "reportId": 15,
                                "useId": 0,
                                "processorCategory": "processor_interface",
                                "interfaceId": 0,
                                "reqContent": "{\"method\":\"GET\",\"url\":\"https://demo/1\",\"params\":[],\"headers\":[],\"body\":\"\",\"bodyFormData\":[],\"bodyFormUrlencoded\":[],\"bodyType\":\"\",\"bodyLang\":\"text\",\"authorizationType\":\"\",\"preRequestScript\":\"\",\"validationScript\":\"\",\"basicAuth\":{\"username\":\"\",\"password\":\"\"},\"bearerToken\":{\"token\":\"\"},\"oauth20\":{\"accessToken\":\"\",\"headerPrefix\":\"\",\"name\":\"\",\"grantType\":\"\",\"callbackUrl\":\"\",\"authURL\":\"\",\"accessTokenURL\":\"\",\"clientID\":\"\",\"clientSecret\":\"\",\"scope\":\"\",\"state\":\"\",\"clientAuthentication\":\"\"},\"apiKey\":{\"key\":\"\",\"value\":\"\",\"transferMode\":\"\"}}", //请求体
                                "respContent": "{\"id\":0,\"statusCode\":503,\"statusContent\":\"503 请求错误\",\"headers\":null,\"content\":\"Get \\\"https://demo/1\\\": dial tcp: lookup demo: no such host\",\"contentType\":\"\",\"contentLang\":\"text\",\"contentCharset\":\"\",\"contentLength\":0,\"time\":0}", //响应体 time:请求耗时
                                "httpStatusCode": 0,
                                "httpStatusContent": "",
                                "processorType": "processor_interface_default",
                                "logs": null
                            }
                        ]
                    }
                ]
            }
        ]
    }
};