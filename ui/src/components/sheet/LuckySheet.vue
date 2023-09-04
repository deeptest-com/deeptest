<template>
  <div style="position: absolute; top: 0">
    <input id="uploadBtn" type="file" @change="loadExcel" />

    <span>Or Load remote xlsx file:</span>

    <select v-model="selected" @change="selectExcel">
      <option disabled value="">Choose</option>
      <option v-for="option in options" :key="option.text" :value="option.value">
        {{ option.text }}
      </option>
    </select>

    <a href="javascript:void(0)" @click="downloadExcel">Download source xlsx file</a>
  </div>
  <div id="luckysheet"></div>
  <div v-show="isMaskShow" id="tip">Downloading</div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { exportExcel } from '../utils/export'
import { isFunction } from '../utils/is'
import LuckyExcel from 'luckyexcel'

const isMaskShow = ref(false)
const selected = ref('')
const jsonData = ref({})
const options = ref([
  { text: 'Money Manager.xlsx', value: 'https://minio.cnbabylon.com/public/luckysheet/money-manager-2.xlsx' },
  {
    text: 'Activity costs tracker.xlsx',
    value: 'https://minio.cnbabylon.com/public/luckysheet/Activity%20costs%20tracker.xlsx',
  },
  {
    text: 'House cleaning checklist.xlsx',
    value: 'https://minio.cnbabylon.com/public/luckysheet/House%20cleaning%20checklist.xlsx',
  },
  {
    text: 'Student assignment planner.xlsx',
    value: 'https://minio.cnbabylon.com/public/luckysheet/Student%20assignment%20planner.xlsx',
  },
  {
    text: 'Credit card tracker.xlsx',
    value: 'https://minio.cnbabylon.com/public/luckysheet/Credit%20card%20tracker.xlsx',
  },
  { text: 'Blue timesheet.xlsx', value: 'https://minio.cnbabylon.com/public/luckysheet/Blue%20timesheet.xlsx' },
  {
    text: 'Student calendar (Mon).xlsx',
    value: 'https://minio.cnbabylon.com/public/luckysheet/Student%20calendar%20%28Mon%29.xlsx',
  },
  {
    text: 'Blue mileage and expense report.xlsx',
    value: 'https://minio.cnbabylon.com/public/luckysheet/Blue%20mileage%20and%20expense%20report.xlsx',
  },
])

const loadExcel = (evt) => {
  const files = evt.target.files
  if (files == null || files.length == 0) {
    alert('No files wait for import')
    return
  }

  let name = files[0].name
  let suffixArr = name.split('.'),
      suffix = suffixArr[suffixArr.length - 1]
  if (suffix != 'xlsx') {
    alert('Currently only supports the import of xlsx files')
    return
  }
  LuckyExcel.transformExcelToLucky(files[0], function (exportJson, luckysheetfile) {
    if (exportJson.sheets == null || exportJson.sheets.length == 0) {
      alert('Failed to read the content of the excel file, currently does not support xls files!')
      return
    }
    console.log('exportJson', exportJson)
    jsonData.value = exportJson

    isFunction(window?.luckysheet?.destroy) && window.luckysheet.destroy()

    window.luckysheet.create({
      container: 'luckysheet', //luckysheet is the container id
      showinfobar: false,
      data: exportJson.sheets,
      title: exportJson.info.name,
      userInfo: exportJson.info.name.creator,
    })
  })
}
const selectExcel = (evt) => {
  const value = selected.value
  const name = evt.target.options[evt.target.selectedIndex].innerText

  if (value == '') {
    return
  }
  isMaskShow.value = true

  LuckyExcel.transformExcelToLuckyByUrl(value, name, (exportJson, luckysheetfile) => {
    if (exportJson.sheets == null || exportJson.sheets.length == 0) {
      alert('Failed to read the content of the excel file, currently does not support xls files!')
      return
    }
    console.log('exportJson', exportJson)
    jsonData.value = exportJson

    isMaskShow.value = false

    isFunction(window?.luckysheet?.destroy) && window.luckysheet.destroy()

    window.luckysheet.create({
      container: 'luckysheet', //luckysheet is the container id
      showinfobar: false,
      data: exportJson.sheets,
      title: exportJson.info.name,
      userInfo: exportJson.info.name.creator,
    })
  })
}
const downloadExcel = () => {
  // const value = selected.value;;
  //
  // if(value.length==0){
  //     alert("Please select a demo file");
  //     return;
  // }
  //
  // var elemIF = document.getElementById("Lucky-download-frame");
  // if(elemIF==null){
  //     elemIF = document.createElement("iframe");
  //     elemIF.style.display = "none";
  //     elemIF.id = "Lucky-download-frame";
  //     document.body.appendChild(elemIF);
  // }
  // elemIF.src = value;
  exportExcel(luckysheet.getAllSheets(), '下载')
}

// !!! create luckysheet after mounted
onMounted(() => {
  luckysheet.create({
    container: 'luckysheet',
  })
})
</script>

<style scoped>
#luckysheet {
  margin: 0px;
  padding: 0px;
  position: absolute;
  width: 100%;
  left: 0px;
  top: 30px;
  bottom: 0px;
}

#uploadBtn {
  font-size: 16px;
}

#tip {
  position: absolute;
  z-index: 1000000;
  left: 0px;
  top: 0px;
  bottom: 0px;
  right: 0px;
  background: rgba(255, 255, 255, 0.8);
  text-align: center;
  font-size: 40px;
  align-items: center;
  justify-content: center;
  display: flex;
}
</style>