<!-- :::: 接口定义模块 -->
<template>
  <div class="content">
    <BasicDetail :items="items"/>
    <a-divider style="margin:0"/>
    <div class="doc-container">
      <div class="left">
        <LeftTreeView :serviceList="serviceList" @select="selectMenu"/>
      </div>
      <div class="right">
        <EndpointDoc :info="selectedItem"/>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {
  ref,
  defineProps,
  defineEmits,
  computed, onMounted,
} from 'vue';
import {useStore} from "vuex";

import BasicDetail from "./components/BasicDetail.vue";
import LeftTreeView from "./components/LeftTreeView.vue";
import EndpointDoc from "./components/EndpointDoc.vue";

const store = useStore<{ Endpoint, ProjectGlobal }>();
const props = defineProps({});
const emit = defineEmits([]);

const items = ref([
  {
    label: '项目名称',
    value: '接口文档',
  },
  {
    label: '接口数量',
    value: 323232
  },
  {
    label: '组件数量',
    value: 33,
  },
  {
    label: '服务数量',
    value: 3,
  }
])

const list = [
  {
    "id": 150,
    "createdAt": "2023-06-07T14:15:32+08:00",
    "updatedAt": "2023-06-07T14:15:32+08:00",
    "name": "petstore",
    "projectId": 63,
    "createUser": "wangzhen",
    "status": 0,
    "description": "",
    "Component": null,
    "servers": [
      {
        "id": 232,
        "createdAt": "2023-06-07T14:15:32+08:00",
        "updatedAt": "2023-06-07T14:15:32+08:00",
        "serveId": 150,
        "environmentId": 141,
        "url": "http://110.42.146.127:8089",
        "description": "本地环境",
        "serveName": "",
        "environmentName": "本地环境"
      },
      {
        "id": 235,
        "createdAt": "2023-06-07T14:26:03+08:00",
        "updatedAt": "2023-06-07T14:26:03+08:00",
        "serveId": 150,
        "environmentId": 142,
        "url": "http://119.3.182.218:50400",
        "description": "测试环境",
        "serveName": "",
        "environmentName": "测试环境"
      }
    ],
    "Securities": [
      {
        "id": 2233,
        "createdAt": "2023-05-25T16:34:25+08:00",
        "updatedAt": "2023-05-25T16:35:12+08:00",
        "name": "token",
        "type": "bearerToken",
        "project_id": 63,
        "serve_id": 122,
        "description": "",
        "in": "",
        "scheme": "",
        "bearerFormat": "",
        "openIdConnectUrl": "",
        "key": "",
        "value": "",
        "token": "WVRCaE1UTTRZbVU1TTJWbVlXWTFOR0ZoWlRsbE5tSm1ZV1poTVdRMVpETS5Nalk0WmpaaVpqQmtORGcyT0RKbE9EZzJaRE15Wmpsak5EWXhNVE16WldV",
        "username": "",
        "password": "",
        "default": true
      },
      {
        "id": 31113,
        "createdAt": "2023-05-25T16:34:25+08:00",
        "updatedAt": "2023-05-25T16:35:12+08:00",
        "name": "token1",
        "type": "bearerToken2",
        "project_id": 63,
        "serve_id": 122,
        "description": "",
        "in": "",
        "scheme": "",
        "bearerFormat": "",
        "openIdConnectUrl": "",
        "key": "",
        "value": "",
        "token": "WVRCaE1UTTRZbVU1TTJWbVlXWTFOR0ZoWlRsbE5tSm1ZV1poTVdRMVpETS5Nalk0WmpaaVpqQmtORGcyT0RKbE9EZzJaRE15Wmpsak5EWXhNVE16WldV",
        "username": "",
        "password": "",
        "default": true
      }
    ],
    endpointList: [
      {
        "id": 1292,
        "createdAt": "2023-06-08T14:33:46+08:00",
        "updatedAt": "2023-06-08T14:33:47+08:00",
        "title": "测试接口",
        "projectId": 6,
        "serveId": 150,
        "serverId": 216,
        "path": "/{path1}",
        "version": "v0.1.0",
        "createUser": "admin",
        "status": 1,
        "categoryId": -1,
        "pathParams": [
          {
            "id": 363,
            "createdAt": "2023-06-08T14:33:47+08:00",
            "updatedAt": "2023-06-08T14:33:47+08:00",
            "name": "path1",
            "value": "",
            "type": "string",
            "desc": "",
            "interfaceId": 0,
            "format": "",
            "example": "",
            "pattern": "",
            "minLength": 0,
            "maxLength": 0,
            "default": "",
            "required": false,
            "multipleOf": 0,
            "minItems": 0,
            "maxItems": 0,
            "uniqueItems": false,
            "ref": "",
            "endpointId": 129
          }
        ],
        "interfaces": [
          {
            "id": 111,
            "createdAt": "2023-06-08T14:33:47+08:00",
            "updatedAt": "2023-06-08T14:33:47+08:00",
            "name": "测试接口",
            "operationId": "",
            "description": "",
            "desc": "",
            "endpointId": 129,
            "security": "",
            "isLeaf": false,
            "parentId": 0,
            "projectId": 6,
            "useId": 0,
            "ordr": 0,
            "slots": null,
            "url": "/{path1}",
            "method": "GET",
            "body": "",
            "bodyType": "application/json",
            "authorizationType": "",
            "preRequestScript": "",
            "validationScript": "",
            "Version": "v0.1.0",
            "params": [
              {
                "id": 224,
                "createdAt": "2023-06-08T14:33:47+08:00",
                "updatedAt": "2023-06-08T14:33:47+08:00",
                "name": "param1",
                "value": "1",
                "type": "string",
                "desc": "",
                "interfaceId": 111,
                "format": "",
                "example": "1",
                "pattern": "",
                "minLength": 0,
                "maxLength": 0,
                "default": "1",
                "required": true,
                "multipleOf": 0,
                "minItems": 0,
                "maxItems": 0,
                "uniqueItems": false,
                "ref": ""
              }
            ],
            "headers": [
              {
                "id": 209,
                "createdAt": "2023-06-08T14:33:47+08:00",
                "updatedAt": "2023-06-08T14:33:47+08:00",
                "name": "header1",
                "value": "",
                "type": "string",
                "desc": "",
                "interfaceId": 111,
                "format": "",
                "example": "1",
                "pattern": "",
                "minLength": 0,
                "maxLength": 0,
                "default": "1",
                "required": true,
                "multipleOf": 0,
                "minItems": 0,
                "maxItems": 0,
                "uniqueItems": false,
                "ref": ""
              }
            ],
            "cookies": [],
            "requestBody": {
              "id": 110,
              "createdAt": "2023-06-08T14:33:47+08:00",
              "updatedAt": "2023-06-08T14:33:47+08:00",
              "interfaceId": 111,
              "mediaType": "",
              "description": "",
              "schemaRefId": 0,
              "examples": "",
              "schemaItem": {
                "id": 1072,
                "createdAt": "2023-06-08T14:33:47+08:00",
                "updatedAt": "2023-06-08T14:33:47+08:00",
                "name": "",
                "type": "object",
                "content": "",
                "requestBodyId": 110
              }
            },
            "responseBodies": [],
            "responseCodes": ""
          },
          {
            "id": 112,
            "createdAt": "2023-06-08T14:33:47+08:00",
            "updatedAt": "2023-06-08T14:33:47+08:00",
            "name": "测试接口",
            "operationId": "",
            "description": "",
            "desc": "",
            "endpointId": 129,
            "security": "",
            "isLeaf": false,
            "parentId": 0,
            "projectId": 6,
            "useId": 0,
            "ordr": 0,
            "slots": null,
            "url": "/{path1}",
            "method": "POST",
            "body": "[{\"name\":\"Example 1\",\"content\":\"{\\\"name1\\\":\\\"string\\\"}\"}]",
            "bodyType": "application/json",
            "authorizationType": "",
            "preRequestScript": "",
            "validationScript": "",
            "Version": "v0.1.0",
            "params": [],
            "headers": [],
            "cookies": [],
            "requestBody": {
              "id": 111,
              "createdAt": "2023-06-08T14:33:47+08:00",
              "updatedAt": "2023-06-08T14:33:47+08:00",
              "interfaceId": 112,
              "mediaType": "application/json",
              "description": "",
              "schemaRefId": 0,
              "examples": "[{\"name\":\"Example 1\",\"content\":\"{\\\"name1\\\":\\\"string\\\"}\"}]",
              "schemaItem": {
                "id": 1073,
                "createdAt": "2023-06-08T14:33:47+08:00",
                "updatedAt": "2023-06-08T14:33:47+08:00",
                "name": "",
                "type": "object",
                "content": "{\"type\":\"object\",\"properties\":{\"name1\":{\"type\":\"string\"}}}",
                "requestBodyId": 111
              }
            },
            "responseBodies": [],
            "responseCodes": ""
          }
        ],
        "versions": null,
        "serveName": "",
        "description": "",
        "serialNumber": "LY-I-129"
      }
    ]
  },
  {
    "id": 137,
    "createdAt": "2023-06-07T10:08:21+08:00",
    "updatedAt": "2023-06-07T10:08:22+08:00",
    "name": "httpbin",
    "projectId": 63,
    "createUser": "wangzhen",
    "status": 0,
    "description": "",
    "Component": null,
    "servers": [
      {
        "id": 234,
        "createdAt": "2023-06-07T14:26:03+08:00",
        "updatedAt": "2023-06-07T14:26:03+08:00",
        "serveId": 137,
        "environmentId": 142,
        "url": "http://119.3.182.218:50300",
        "description": "测试环境",
        "serveName": "",
        "environmentName": "测试环境"
      }
    ],
    "Securities": null,
    endpointList: [
      {
        "id": 1207,
        "createdAt": "2023-05-18T10:11:15+08:00",
        "updatedAt": "2023-05-18T10:11:15+08:00",
        "title": "测试错误",
        "projectId": 63,
        "serveId": 137,
        "serverId": 186,
        "path": "/api/v1/init/initdb",
        "version": "v0.1.0",
        "createUser": "admin",
        "status": 0,
        "categoryId": 237,
        "pathParams": [],
        "interfaces": [
          {
            "id": 81,
            "createdAt": "2023-05-18T10:11:15+08:00",
            "updatedAt": "2023-05-18T10:11:15+08:00",
            "name": "测试错误",
            "operationId": "Operation ID",
            "description": "描述11111",
            "desc": "",
            "endpointId": 107,
            "security": "",
            "isLeaf": false,
            "parentId": 0,
            "projectId": 63,
            "useId": 0,
            "ordr": 0,
            "slots": null,
            "url": "/api/v1/init/initdb",
            "method": "GET",
            "body": "",
            "bodyType": "",
            "authorizationType": "",
            "preRequestScript": "",
            "validationScript": "",
            "Version": "v0.1.0",
            "params": [],
            "headers": [],
            "cookies": [],
            "requestBody": {
              "id": 80,
              "createdAt": "2023-05-18T10:11:15+08:00",
              "updatedAt": "2023-05-18T10:11:15+08:00",
              "interfaceId": 81,
              "mediaType": "application/json",
              "description": "",
              "schemaRefId": 0,
              "examples": "",
              "schemaItem": {
                "id": 413,
                "createdAt": "2023-05-18T10:11:15+08:00",
                "updatedAt": "2023-05-18T10:11:15+08:00",
                "name": "",
                "type": "object",
                "content": "",
                "requestBodyId": 80
              }
            },
            "responseBodies": [
              {
                "id": 62224,
                "createdAt": "2023-05-18T10:11:15+08:00",
                "updatedAt": "2023-05-18T10:11:15+08:00",
                "code": "200",
                "interfaceId": 81,
                "mediaType": "application/json",
                "description": "200 description",
                "schemaRefId": 0,
                "examples": "[]",
                "schemaItem": {
                  "id": 267,
                  "createdAt": "2023-05-18T10:11:15+08:00",
                  "updatedAt": "2023-05-18T10:11:15+08:00",
                  "name": "",
                  "type": "object",
                  "content": "{\"type\":\"object\",\"properties\":{\"name1\":{\"type\":\"string\"},\"name2\":{\"type\":\"string\"},\"name3\":{\"type\":\"string\"}}}",
                  "ResponseBodyId": 64
                },
                "headers": [
                  {
                    "id": 209,
                    "createdAt": "2023-06-08T14:33:47+08:00",
                    "updatedAt": "2023-06-08T14:33:47+08:00",
                    "name": "header1",
                    "value": "",
                    "type": "string",
                    "desc": "",
                    "interfaceId": 111,
                    "format": "",
                    "example": "1",
                    "pattern": "",
                    "minLength": 0,
                    "maxLength": 0,
                    "default": "1",
                    "required": true,
                    "multipleOf": 0,
                    "minItems": 0,
                    "maxItems": 0,
                    "uniqueItems": false,
                    "ref": ""
                  }
                ],
              },
              {
                "id": 641,
                "createdAt": "2023-05-18T10:11:15+08:00",
                "updatedAt": "2023-05-18T10:11:15+08:00",
                "code": "403",
                "interfaceId": 81,
                "mediaType": "application/json",
                "schemaRefId": 0,
                "examples": "[]",
                "description": "403 description",
                "schemaItem": {
                  "id": 267,
                  "createdAt": "2023-05-18T10:11:15+08:00",
                  "updatedAt": "2023-05-18T10:11:15+08:00",
                  "name": "",
                  "type": "object",
                  "content": "{\"type\":\"object\",\"properties\":{\"name1\":{\"type\":\"string\"},\"name2\":{\"type\":\"string\"},\"name3\":{\"type\":\"string\"}}}",
                  "ResponseBodyId": 64
                },
                "headers": [
                  {
                    "id": 209,
                    "createdAt": "2023-06-08T14:33:47+08:00",
                    "updatedAt": "2023-06-08T14:33:47+08:00",
                    "name": "header1",
                    "value": "",
                    "type": "string",
                    "desc": "",
                    "interfaceId": 111,
                    "format": "",
                    "example": "1",
                    "pattern": "",
                    "minLength": 0,
                    "maxLength": 0,
                    "default": "1",
                    "required": true,
                    "multipleOf": 0,
                    "minItems": 0,
                    "maxItems": 0,
                    "uniqueItems": false,
                    "ref": ""
                  }
                ],
              },
              {
                "id": 624,
                "createdAt": "2023-05-18T10:11:15+08:00",
                "updatedAt": "2023-05-18T10:11:15+08:00",
                "code": "500",
                "interfaceId": 81,
                "mediaType": "application/json",
                "description": "500 description",
                "schemaRefId": 0,
                "examples": "[]",
                "schemaItem": {
                  "id": 267,
                  "createdAt": "2023-05-18T10:11:15+08:00",
                  "updatedAt": "2023-05-18T10:11:15+08:00",
                  "name": "",
                  "type": "object",
                  "content": "{\"type\":\"object\",\"properties\":{\"name1\":{\"type\":\"string\"},\"name2\":{\"type\":\"string\"},\"name3\":{\"type\":\"string\"}}}",
                  "ResponseBodyId": 64
                },
                "headers": [
                  {
                    "id": 209,
                    "createdAt": "2023-06-08T14:33:47+08:00",
                    "updatedAt": "2023-06-08T14:33:47+08:00",
                    "name": "header1",
                    "value": "",
                    "type": "string",
                    "desc": "",
                    "interfaceId": 111,
                    "format": "",
                    "example": "1",
                    "pattern": "",
                    "minLength": 0,
                    "maxLength": 0,
                    "default": "1",
                    "required": true,
                    "multipleOf": 0,
                    "minItems": 0,
                    "maxItems": 0,
                    "uniqueItems": false,
                    "ref": ""
                  }
                ],
              }
            ],
            "responseCodes": ""
          }
        ],
        "versions": null,
        "serveName": "",
        "description": "111",
        "serialNumber": "TS-I-107"
      }
    ]
  },
]
const serviceList = ref([])

// 这里还是需要打平的
onMounted(() => {
  // 组装数据以兼容组件 LeftTreeMenu
  let items: any = [];
  list.forEach((item: any) => {
    items.push(item);
    item.endpointList.forEach((endpoint: any) => {
      endpoint.interfaces.forEach((interfaceItem: any) => {
        items.push({
          ...interfaceItem,
          endpointInfo: endpoint,
          serveInfo: item,
          serveId: item.id,
        });
      })
    })
  })
  serviceList.value = items;

  // 默认选中第一个
  selectedItem.value = items[0];
})


const selectedItem = ref(null);
function selectMenu(item) {
  selectedItem.value = item
}


</script>

<style lang="less" scoped>
.content {
  min-height: calc(100vh - 200px);
  position: relative;

  .mode-btns {
    position: absolute;
    right: 0px;
    z-index: 999;
  }
}

.doc-container {
  display: flex;
  height: 100%;

  .left {
    width: 300px;
    height: 100%;
    overflow: auto;
  }

  .right {
    flex: 1;
    height: 100%;
    overflow: auto;
    padding: 24px;
  }
}
</style>
