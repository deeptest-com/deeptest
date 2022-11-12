<template>
  <div id="expresslayout-left">
    <div class="expresslayout-left-logo">
      <router-link to="/" class="logo-url">
        <div class="logo-title">DeepTest</div>
      </router-link>
    </div>

    <div class="expresslayout-left-content">
      <div class="search">
        <a-input>
          <template #prefix>
            <SearchOutlined/>
          </template>
        </a-input>
      </div>

      <div class="content">
        <ul class="list desc">
          <li v-for="(item, index) in specData.info.desc" :key="index" @click="selectSection(index)">
            <span class="item">{{ getSectionTitle(item) }}</span>
          </li>
        </ul>

        <div class="title">APIS</div>
        <ul class="list apis">
          <template v-for="(item, index) in specData.doc?.paths" :key="index">
            <li v-for="(path, method) in item" :key="method" @click="selectApi(method, path)">
              <span class="method" :class="method">{{ method }}</span>
              <span class="name">{{ path.summary }}</span>
            </li>
          </template>
        </ul>

        <div class="title">MODELS</div>
        <ul class="list models">
          <li @click="selectModel('item')">
            <span class="name">Pet</span>
          </li>
          <li>
            <span class="name">Dog</span>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {useStore} from "vuex";

const store = useStore<{ Global: GlobalStateType, Spec: SpecStateType}>();
const specData = computed<any>(() => store.state.Spec.specData);

import {SearchOutlined} from '@ant-design/icons-vue';
import {StateType as GlobalStateType} from "@/store/global";
import {StateType as SpecStateType} from "@/views/express/store";
import {computed, nextTick} from "vue";

const selectSection = (index) => {
  console.log('selectSection', specData.value.info.desc)

  store.dispatch('Spec/setMode', 'desc')

  nextTick(()=>{
    const id = getSectionTitle(specData.value.info.desc[index])
    const elem = document.getElementById(id);
    if (elem) elem.scrollIntoView()
  })
}

const selectApi = (method, path) => {
  console.log('selectApi')
  store.dispatch('Spec/setPath', path)
}

const selectModel = (item) => {
  console.log('selectModel')
}

const getSectionTitle = (item) => {
  const start = item[0]
  const end = item[1]

  var content

  if (start + end === 0) {
    content = 'Introduction'
  } else {
    content = specData.value.doc.info.description.substring(start, end).trim()
  }

  return content
}

</script>

<style lang="less">
.expresslayout-left-content {
  .ant-input-affix-wrapper {
    input.ant-input {
      background-color: #FAFAFA !important;
    }
  }
}
</style>

<style lang="less" scoped>
#expresslayout-left {
  display: flex;
  flex-direction: column;
  width: 280px;
  height: 100vh;
  background-color: #FAFAFA;
  color: #343333;

  .expresslayout-left-logo {
    width: 100%;
    height: 50px;
    line-height: 50px;
    text-align: center;
    vertical-align: middle;

    .logo-url {
      display: inline-block;
      width: 100%;
      height: 100%;
      overflow: hidden;

      .logo-title {
        display: inline-block;
        margin: 0;
        font-size: 22px;
        font-family: Roboto, sans-serif;
        color: rgb(51, 51, 51);
      }
    }

    img {
      vertical-align: middle;
    }
  }

  .expresslayout-left-content {
    padding: 0 8px 8px 8px;
    height: calc(100% - 50px);

    .search {
      margin-bottom: 15px;
      height: 26px;
    }

    .content {
      height: calc(100% - 40px);
      overflow-y: auto;
      .ant-input-affix-wrapper {
        border: 0;
        border-bottom: 1px solid #d9d9d9;
        background-color: #FAFAFA;
        box-shadow: none;
      }

      .title {
        padding: 10px 10px 5px 10px;
        font-size: 16px;
        font-weight: normal;
      }

      ul {
        margin: 0px;
        padding: 0px;

        li {
          list-style: inside none none;
          overflow: hidden;
          text-overflow: ellipsis;
          padding: 12px 10px;
          line-height: 16px;
          cursor: pointer;
          font-family: Montserrat, sans-serif;
          font-size: 15px;

          &:hover {
            color: #32329FFF;
            background-color: #ededed;
          }
          &.active {
            color: #32329FFF;
            background-color: #EDEDEDFF;
          }

          .method {
            width: 9ex;
            display: inline-block;
            background-color: rgb(51, 51, 51);
            border-radius: 3px;
            background-repeat: no-repeat;
            background-position: 6px 4px;
            font-size: 7px;
            font-family: Verdana, sans-serif;
            color: white;
            text-transform: uppercase;
            text-align: center;
            font-weight: bold;
            vertical-align: middle;
            margin-right: 6px;

            &.post {
              background-color: #186FAF;
            }

            &.get {
              background-color: #2F8132;
            }

            &.put {
              background-color: #95507C;
            }

            &.delete {
              background-color: #c33;
            }

            &.patch {
              background-color: #bf581d;
            }

            head {
              background-color: #a23dad;
            }

            options {
              background-color: #947014;
            }

            .connect {
              background-color: #149494;
            }
            .trace {
              background-color: #2dcebb;
            }

          }

          .name {
            vertical-align: middle;
          }
        }

        &.apis {

        }

        &.models {

        }
      }
    }
  }

}
</style>