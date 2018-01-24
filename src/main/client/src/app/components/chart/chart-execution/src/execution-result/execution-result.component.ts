import {Component, Input, OnInit} from "@angular/core";

@Component({
  selector: 'execution-result',
  templateUrl: './execution-result.html',
  styleUrls: ['./styles.scss']
})
export class ExecutionResultComponent implements OnInit {
  chartOption: any;

  _data: any = {};
  @Input() set data(model: any) {
    console.log('===', model);
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
      legend: {
        show: false
      },
      padding: 0,
      color: ['#c23531', '#ca8622', '#749f83', '#c4ccd3'],
      tooltip: {
        trigger: 'item',
        formatter: "{a} <br/>{b} : {c} ({d}%)"
      },

      series: [
        {
          name: '执行结果',
          type: 'pie',
          radius: '90%',
          center: ['50%', '55%'],
          data: this._data,
          itemStyle: {
            emphasis: {
              shadowBlur: 10,
              shadowOffsetX: 0,
              shadowColor: 'rgba(0, 0, 0, 0.5)'
            },
            normal:{
              label:{
                show: true,
                formatter: '{b} : {c} ({d}%)'
              },
              labelLine :{show:true}
            }
          }
        }
      ]
    };
  }

}
