const list = [
    {
        "id": 68,
        "scenarioId": 0,
        "scenarioReportId": 0,
        "name": "场景3",
        "resultStatus": "",
        "startTime": "2023-05-29T13:54:50.1118648+08:00",
        "parentId": 0,
        "processorCategory": "processor_root",
        "processorType": "processor_root_default",
        "iterator": {
            "ProcessorCategory": "",
            "ProcessorType": "",
            "items": null,
            "data": null,
            "dataType": "",
            "untilExpression": ""
        }
    },
    {
        "id": 69,
        "scenarioId": 0,
        "scenarioReportId": 0,
        "name": "初始化DB - POST",
        "resultStatus": "fail",
        "startTime": "2023-05-29T13:54:50.1130035+08:00",
        "parentId": 68,
        "processorCategory": "processor_interface",
        "processorType": "processor_interface_default",
        "interfaceId": 67,
        "reqContent": "{\"method\":\"POST\",\"url\":\"api/v1/init/initdb\",\"queryParams\":[{\"name\":\"\",\"value\":\"\",\"paramIn\":\"\",\"disabled\":false}],\"pathParams\":[{\"name\":\"\",\"value\":\"\",\"paramIn\":\"\",\"disabled\":false}],\"headers\":[{\"name\":\"\",\"value\":\"\",\"disabled\":false}],\"cookies\":null,\"body\":\"\",\"bodyFormData\":[{\"name\":\"\",\"value\":\"\",\"type\":\"text\",\"desc\":\"\",\"interfaceId\":0}],\"BodyFormUrlencoded\":[{\"name\":\"\",\"value\":\"\",\"desc\":\"\",\"interfaceId\":0}],\"bodyType\":\"application/json\",\"bodyLang\":\"json\",\"authorizationType\":\"\",\"preRequestScript\":\"\",\"validationScript\":\"\",\"basicAuth\":{\"username\":\"\",\"password\":\"\"},\"bearerToken\":{\"token\":\"\"},\"oauth20\":{\"accessToken\":\"\",\"headerPrefix\":\"Bearer\",\"name\":\"\",\"grantType\":\"authorizationCode\",\"callbackUrl\":\"\",\"authURL\":\"\",\"accessTokenURL\":\"\",\"clientID\":\"\",\"clientSecret\":\"\",\"scope\":\"\",\"state\":\"\",\"clientAuthentication\":\"sendAsBasicAuthHeader\"},\"apiKey\":{\"key\":\"\",\"value\":\"\",\"transferMode\":\"\"}}",
        "respContent": "{\"id\":0,\"statusCode\":503,\"statusContent\":\"503 请求错误\",\"headers\":null,\"cookies\":null,\"content\":\"Post \\\"api/v1/init/initdb\\\": unsupported protocol scheme \\\"\\\"\",\"contentType\":\"\",\"contentLang\":\"text\",\"contentCharset\":\"\",\"contentLength\":0,\"time\":0}",
        "iterator": {
            "ProcessorCategory": "",
            "ProcessorType": "",
            "items": null,
            "data": null,
            "dataType": "",
            "untilExpression": ""
        },
        "summary": "Post \"api/v1/init/initdb\": unsupported protocol scheme \"\""
    },
    {
        "id": 651,
        "name": "场景3",
        "desc": "",
        "progressStatus": "",
        "resultStatus": "fail",
        "startTime": "2023-05-29T13:54:50.1118648+08:00",
        "endTime": "2023-05-29T13:54:50.1141512+08:00",
        "duration": 3,
        "totalScenarioNum": 0,
        "passScenarioNum": 0,
        "failScenarioNum": 0,
        "totalInterfaceNum": 1,
        "passInterfaceNum": 0,
        "failInterfaceNum": 1,
        "totalRequestNum": 1,
        "passRequestNum": 0,
        "failRequestNum": 1,
        "totalAssertionNum": 0,
        "passAssertionNum": 0,
        "failAssertionNum": 0,
        "InterfaceStatusMap": {
            "67": {
                "fail": 1,
                "pass": 0
            }
        },
        "payload": "",
        "scenarioId": 20,
        "projectId": 63,
        "logs": null
    },
    {
        "id": 102,
        "scenarioId": 0,
        "scenarioReportId": 0,
        "name": "自动化测试场景",
        "resultStatus": "",
        "startTime": "2023-05-29T13:54:50.1783017+08:00",
        "parentId": 0,
        "processorCategory": "processor_root",
        "processorType": "processor_root_default",
        "iterator": {
            "ProcessorCategory": "",
            "ProcessorType": "",
            "items": null,
            "data": null,
            "dataType": "",
            "untilExpression": ""
        }
    },
    {
        "id": 103,
        "scenarioId": 0,
        "scenarioReportId": 0,
        "name": "初始化DB - POST",
        "resultStatus": "fail",
        "startTime": "2023-05-29T13:54:50.1798987+08:00",
        "parentId": 102,
        "processorCategory": "processor_interface",
        "processorType": "processor_interface_default",
        "interfaceId": 67,
        "reqContent": "{\"method\":\"POST\",\"url\":\"api/v1/init/initdb\",\"queryParams\":[{\"name\":\"\",\"value\":\"\",\"paramIn\":\"\",\"disabled\":false}],\"pathParams\":[{\"name\":\"\",\"value\":\"\",\"paramIn\":\"\",\"disabled\":false}],\"headers\":[{\"name\":\"\",\"value\":\"\",\"disabled\":false}],\"cookies\":null,\"body\":\"\",\"bodyFormData\":[{\"name\":\"\",\"value\":\"\",\"type\":\"text\",\"desc\":\"\",\"interfaceId\":0}],\"BodyFormUrlencoded\":[{\"name\":\"\",\"value\":\"\",\"desc\":\"\",\"interfaceId\":0}],\"bodyType\":\"application/json\",\"bodyLang\":\"json\",\"authorizationType\":\"\",\"preRequestScript\":\"\",\"validationScript\":\"\",\"basicAuth\":{\"username\":\"\",\"password\":\"\"},\"bearerToken\":{\"token\":\"\"},\"oauth20\":{\"accessToken\":\"\",\"headerPrefix\":\"Bearer\",\"name\":\"\",\"grantType\":\"authorizationCode\",\"callbackUrl\":\"\",\"authURL\":\"\",\"accessTokenURL\":\"\",\"clientID\":\"\",\"clientSecret\":\"\",\"scope\":\"\",\"state\":\"\",\"clientAuthentication\":\"sendAsBasicAuthHeader\"},\"apiKey\":{\"key\":\"\",\"value\":\"\",\"transferMode\":\"\"}}",
        "respContent": "{\"id\":0,\"statusCode\":503,\"statusContent\":\"503 请求错误\",\"headers\":null,\"cookies\":null,\"content\":\"Post \\\"api/v1/init/initdb\\\": unsupported protocol scheme \\\"\\\"\",\"contentType\":\"\",\"contentLang\":\"text\",\"contentCharset\":\"\",\"contentLength\":0,\"time\":0}",
        "iterator": {
            "ProcessorCategory": "",
            "ProcessorType": "",
            "items": null,
            "data": null,
            "dataType": "",
            "untilExpression": ""
        },
        "summary": "Post \"api/v1/init/initdb\": unsupported protocol scheme \"\""
    },
    {
        "id": 104,
        "scenarioId": 0,
        "scenarioReportId": 0,
        "name": "迭代次数",
        "resultStatus": "",
        "startTime": "2023-05-29T13:54:50.1815219+08:00",
        "parentId": 102,
        "processorCategory": "processor_loop",
        "processorType": "processor_loop_time",
        "iterator": {
            "ProcessorCategory": "",
            "ProcessorType": "",
            "variableName": "迭代3次",
            "items": [
                1,
                2,
                3
            ],
            "data": null,
            "dataType": "int",
            "untilExpression": ""
        },
        "summary": "迭代\"3\"次。"
    },
    {
        "id": 0,
        "scenarioId": 0,
        "scenarioReportId": 0,
        "name": "",
        "resultStatus": "",
        "parentId": 104,
        "processorCategory": "",
        "processorType": "",
        "iterator": {
            "ProcessorCategory": "",
            "ProcessorType": "",
            "items": null,
            "data": null,
            "dataType": "",
            "untilExpression": ""
        },
        "summary": "1. 迭代3次为1"
    },
    {
        "id": 105,
        "scenarioId": 0,
        "scenarioReportId": 0,
        "name": "初始化DB - POST",
        "resultStatus": "fail",
        "startTime": "2023-05-29T13:54:50.1831354+08:00",
        "parentId": 104,
        "processorCategory": "processor_interface",
        "processorType": "processor_interface_default",
        "interfaceId": 66,
        "reqContent": "{\"method\":\"POST\",\"url\":\"/api/v1/init/initdb\",\"queryParams\":[{\"name\":\"\",\"value\":\"\",\"paramIn\":\"\",\"disabled\":false}],\"pathParams\":[{\"name\":\"\",\"value\":\"\",\"paramIn\":\"\",\"disabled\":false}],\"headers\":[{\"name\":\"\",\"value\":\"\",\"disabled\":false}],\"cookies\":null,\"body\":\"\",\"bodyFormData\":[{\"name\":\"\",\"value\":\"\",\"type\":\"text\",\"desc\":\"\",\"interfaceId\":0}],\"BodyFormUrlencoded\":[{\"name\":\"\",\"value\":\"\",\"desc\":\"\",\"interfaceId\":0}],\"bodyType\":\"application/json\",\"bodyLang\":\"json\",\"authorizationType\":\"bearerToken\",\"preRequestScript\":\"\",\"validationScript\":\"\",\"basicAuth\":{\"username\":\"\",\"password\":\"\"},\"bearerToken\":{\"token\":\"WVRCaE1UTTRZbVU1TTJWbVlXWTFOR0ZoWlRsbE5tSm1ZV1poTVdRMVpETS5Nalk0WmpaaVpqQmtORGcyT0RKbE9EZzJaRE15Wmpsak5EWXhNVE16WldV\"},\"oauth20\":{\"accessToken\":\"\",\"headerPrefix\":\"\",\"name\":\"\",\"grantType\":\"\",\"callbackUrl\":\"\",\"authURL\":\"\",\"accessTokenURL\":\"\",\"clientID\":\"\",\"clientSecret\":\"\",\"scope\":\"\",\"state\":\"\",\"clientAuthentication\":\"\"},\"apiKey\":{\"key\":\"\",\"value\":\"\",\"transferMode\":\"\"}}",
        "respContent": "{\"id\":0,\"statusCode\":503,\"statusContent\":\"503 请求错误\",\"headers\":null,\"cookies\":null,\"content\":\"Post \\\"/api/v1/init/initdb\\\": unsupported protocol scheme \\\"\\\"\",\"contentType\":\"\",\"contentLang\":\"text\",\"contentCharset\":\"\",\"contentLength\":0,\"time\":0}",
        "iterator": {
            "ProcessorCategory": "",
            "ProcessorType": "",
            "items": null,
            "data": null,
            "dataType": "",
            "untilExpression": ""
        },
        "summary": "Post \"/api/v1/init/initdb\": unsupported protocol scheme \"\""
    },
    {
        "id": 106,
        "scenarioId": 0,
        "scenarioReportId": 0,
        "name": "如果",
        "resultStatus": "fail",
        "startTime": "2023-05-29T13:54:50.1885261+08:00",
        "parentId": 102,
        "processorCategory": "processor_logic",
        "processorType": "processor_logic_if",
        "iterator": {
            "ProcessorCategory": "",
            "ProcessorType": "",
            "items": null,
            "data": null,
            "dataType": "",
            "untilExpression": ""
        },
        "summary": "失败"
    },
    {
        "id": 652,
        "name": "自动化测试场景",
        "desc": "",
        "progressStatus": "",
        "resultStatus": "fail",
        "startTime": "2023-05-29T13:54:50.1783017+08:00",
        "endTime": "2023-05-29T13:54:50.1890706+08:00",
        "duration": 11,
        "totalScenarioNum": 0,
        "passScenarioNum": 0,
        "failScenarioNum": 0,
        "totalInterfaceNum": 2,
        "passInterfaceNum": 0,
        "failInterfaceNum": 2,
        "totalRequestNum": 4,
        "passRequestNum": 0,
        "failRequestNum": 4,
        "totalAssertionNum": 0,
        "passAssertionNum": 0,
        "failAssertionNum": 0,
        "InterfaceStatusMap": {
            "66": {
                "fail": 3,
                "pass": 0
            },
            "67": {
                "fail": 1,
                "pass": 0
            }
        },
        "payload": "",
        "scenarioId": 48,
        "projectId": 63,
        "logs": null
    },
    {
        "id": 133,
        "scenarioId": 0,
        "scenarioReportId": 0,
        "name": "登录流程测试",
        "resultStatus": "",
        "startTime": "2023-05-29T13:54:50.2738177+08:00",
        "parentId": 0,
        "processorCategory": "processor_root",
        "processorType": "processor_root_default",
        "iterator": {
            "ProcessorCategory": "",
            "ProcessorType": "",
            "items": null,
            "data": null,
            "dataType": "",
            "untilExpression": ""
        }
    },
    {
        "id": 135,
        "scenarioId": 0,
        "scenarioReportId": 0,
        "name": "获取Cookie",
        "resultStatus": "",
        "startTime": "2023-05-29T13:54:50.2744483+08:00",
        "parentId": 133,
        "processorCategory": "processor_cookie",
        "processorType": "processor_cookie_get",
        "iterator": {
            "ProcessorCategory": "",
            "ProcessorType": "",
            "items": null,
            "data": null,
            "dataType": "",
            "untilExpression": ""
        },
        "summary": "将值\"空\"赋予变量。"
    },
    {
        "id": 136,
        "scenarioId": 0,
        "scenarioReportId": 0,
        "name": "计时器",
        "resultStatus": "",
        "startTime": "2023-05-29T13:54:50.2749551+08:00",
        "parentId": 133,
        "processorCategory": "processor_timer",
        "processorType": "processor_time_default",
        "iterator": {
            "ProcessorCategory": "",
            "ProcessorType": "",
            "items": null,
            "data": null,
            "dataType": "",
            "untilExpression": ""
        },
        "summary": "等待\"0\"秒。"
    },
    {
        "id": 137,
        "scenarioId": 0,
        "scenarioReportId": 0,
        "name": "输出",
        "resultStatus": "",
        "startTime": "2023-05-29T13:54:50.2760876+08:00",
        "parentId": 133,
        "processorCategory": "processor_print",
        "processorType": "processor_print_default",
        "iterator": {
            "ProcessorCategory": "",
            "ProcessorType": "",
            "items": null,
            "data": null,
            "dataType": "",
            "untilExpression": ""
        },
        "summary": "为\"\"。"
    },
    {
        "id": 138,
        "scenarioId": 0,
        "scenarioReportId": 0,
        "name": "迭代区间",
        "resultStatus": "",
        "startTime": "2023-05-29T13:54:50.2767554+08:00",
        "parentId": 133,
        "processorCategory": "processor_loop",
        "processorType": "processor_loop_range",
        "iterator": {
            "ProcessorCategory": "",
            "ProcessorType": "",
            "items": null,
            "data": null,
            "dataType": "",
            "untilExpression": ""
        },
        "summary": "执行前请先配置处理器。"
    },
    {
        "id": 653,
        "name": "登录流程测试",
        "desc": "",
        "progressStatus": "",
        "resultStatus": "pass",
        "startTime": "2023-05-29T13:54:50.2738177+08:00",
        "endTime": "2023-05-29T13:54:50.2767554+08:00",
        "duration": 3,
        "totalScenarioNum": 0,
        "passScenarioNum": 0,
        "failScenarioNum": 0,
        "totalInterfaceNum": 0,
        "passInterfaceNum": 0,
        "failInterfaceNum": 0,
        "totalRequestNum": 0,
        "passRequestNum": 0,
        "failRequestNum": 0,
        "totalAssertionNum": 0,
        "passAssertionNum": 0,
        "failAssertionNum": 0,
        "InterfaceStatusMap": null,
        "payload": "",
        "scenarioId": 53,
        "projectId": 63,
        "logs": null
    },
    {
        "id": 196,
        "name": "自动化测试",
        "desc": "",
        "progressStatus": "",
        "resultStatus": "pass",
        "startTime": "2023-05-29T13:54:50.1118648+08:00",
        "endTime": "2023-05-29T13:54:50.2767554+08:00",
        "duration": 165,
        "totalScenarioNum": 3,
        "passScenarioNum": 1,
        "failScenarioNum": 2,
        "totalInterfaceNum": 3,
        "passInterfaceNum": 0,
        "failInterfaceNum": 3,
        "totalRequestNum": 5,
        "passRequestNum": 0,
        "failRequestNum": 5,
        "totalAssertionNum": 0,
        "passAssertionNum": 0,
        "failAssertionNum": 0,
        "InterfaceStatusMap": {},
        "payload": "",
        "scenarioId": 0,
        "projectId": 63,
        "logs": null
    }
]

function genLogTreeView(data) {
    // 用于存储根节点，即场景节点,即processorCategory为processor_root的节点
    const scenarioReports= [];
    // 用于存储场景节点的所有子节点
    const scenarioChildren= [];
    data.forEach((item) => {
        if (item.processorCategory === "processor_root") {
            scenarioReports.push(item);
        } else {
            scenarioChildren.push(item);
        }
    });
    scenarioReports.forEach((scenario) => {
        function fn(array, rootId) {
            const res= [];
            // 用于存储 树节点的map
            const map = {};
            array.forEach((item) => {
                map[item.id] = {...item}
            });
            array.forEach((item) => {
                const {id, parentId} = item;
                if(!map[id]) return;
                if(!parentId) return;
                if (parentId === rootId) {
                    res.push(map[id]);
                } else {
                    if (map[parentId]?.logs) {
                        map[parentId].logs.push(map[id]);
                    } else {
                        map[parentId].logs = [map[id]];
                    }
                }
            })
            return res;
        }
        scenario.logs = fn(data, scenario.id);
    });
    return scenarioReports;
}

console.log('832', JSON.stringify(genLogTreeView(list)));
