import {Component, Input, OnInit} from "@angular/core";

@Component({
  selector: 'execution-progress',
  templateUrl: './execution-progress.html',
  styleUrls: ['./styles.scss']
})
export class ExecutionProgressComponent implements OnInit {

  chartOption: any;

  _data: any = {};
  @Input() set data(model: any) {
    if (model) {
      this._data = model;
      this.genChart();
    }
  }

  constructor() {
  }

  ngOnInit(): any {

  }
  genChart(): any {
    this.chartOption = {
      title: {
        show: false
      },
      tooltip: {
        trigger: 'item',
        formatter: '{a} <br/>{b} : {c}'
      },
      legend: {
        right: '0%',
        width: '15%',
        data: ['用例数']
      },
      grid: {
        top: '15%',
        right: '15%',
        bottom: '5%',
        left: '5%',
        containLabel: true,
      },
      color: ['#2f4554', '#c23531'],
      xAxis: {
        type: 'category',
        name: '',
        splitLine: {show: false},
        data: this._data.xList,
        axisLabel :{
          interval: 0,
          rotate: 45,
          margin:20
        },
      },

      yAxis: [
        {
          name: '剩余用例（个）',
          type: 'value'
        },
        // {
        //   name: '剩余工作量（小时）',
        //   max: 70,
        //   type: 'value'
        // }
      ],
      series: [
        {
          name: '用例数',
          type: 'line',
          data: this._data.numbList
        },
        // {
        //   name: '工作量',
        //   type: 'line',
        //   yAxisIndex:1,
        //   data: [67, 59, 45, 33, 29, 13, 3],
        //
        // }
      ]
    };;
  }

}
