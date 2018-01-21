import {Component, Input, OnInit} from "@angular/core";

@Component({
  selector: 'execution-progress',
  templateUrl: './execution-progress.html',
  styleUrls: ['./styles.scss']
})
export class ExecutionProgressComponent implements OnInit {

  @Input() data: any;
  chartOption: any;

  constructor() {
  }

  ngOnInit(): any {
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
        data: ['用例数', '工作量']
      },
      grid: {
        top: '15%',
        right: '15%',
        bottom: '5%',
        left: '1%',
        containLabel: true,
      },
      color: ['#2f4554', '#c23531'],
      xAxis: {
        type: 'category',
        name: '',
        splitLine: {show: false},
        data: ['07-01', '07-02', '07-03', '07-04', '07-05', '07-06', '07-07'],
        axisLabel :{
          interval: 0,
          rotate: 15,
          margin:20
        },
      },

      yAxis: [
        {
          name: '剩余用例（个）',
          type: 'value',
          max: 1100
        },
        {
          name: '剩余工作量（小时）',
          max: 70,
          type: 'value'
        }
      ],
      series: [
        {
          name: '用例数',
          type: 'line',
          data: [1000, 950, 880, 750, 690, 630, 540],
        },
        {
          name: '工作量',
          type: 'line',
          yAxisIndex:1,
          data: [67, 59, 45, 33, 29, 13, 3],

        }
      ]
    };;
  }

}
