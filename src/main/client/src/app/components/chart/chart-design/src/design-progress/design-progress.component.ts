import {Component, Input, OnInit} from "@angular/core";

@Component({
  selector: 'design-progress',
  templateUrl: './design-progress.html',
  styleUrls: ['./styles.scss']
})
export class DesignProgressComponent implements OnInit {
  @Input() showTitle: boolean = false;
  @Input() data: any;
  chartOption: any;

  constructor() {
  }

  ngOnInit(): any {
    this.chartOption = {
      title: {
        text: '测试设计',
        show: this.showTitle,
        top: 'top',
        left: 'center',
        textStyle: {
          fontSize: '15'
        }
      },
      tooltip: {
        trigger: 'item',
        formatter: '{a} <br/>{b} : {c}'
      },
      legend: {
        right: '0%',
        width: '15%',
        data: ['合计用例数', '新增用例数']
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
          name: '数量（个）',
          type: 'value',
          max: 1100
        }
      ],
      series: [
        {
          name: '合计用例',
          type: 'line',
          data: [540, 630, 690, 750, 880, 950, 1000],
        },
        {
          name:'新增用例',
          type:'bar',
          data:[100, 90, 70, 60, 130, 70, 50]
        }
      ]
    };
  }

}
