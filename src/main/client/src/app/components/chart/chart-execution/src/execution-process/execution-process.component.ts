import {Component, Input, OnInit} from "@angular/core";

@Component({
  selector: 'execution-process',
  templateUrl: './execution-process.html',
  styleUrls: ['./styles.scss']
})
export class ExecutionProcessComponent implements OnInit {

  @Input() showTitle: boolean = false;
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
        text: '测试执行',
        show: this.showTitle,
        top: 'top',
        left: 'center',
        textStyle: {
          fontSize: '15'
        }
      },
      tooltip : {
        trigger: 'axis',
        axisPointer : {
          type : 'shadow'
        }
      },
      grid: {
        top: '15%',
        right: '15%',
        bottom: '8%',
        left: '3%',
        containLabel: true
      },
      legend: {
        right: '0%',
        width: '15%',
        data:['阻塞', '失败', '通过']
      },
      color: ['#c23531', '#ca8622', '#749f83'],

      xAxis : [
        {
          type : 'category',
          axisLabel :{
            interval: 0,
            rotate: 45,
            margin: 10
          },
          data : this._data.xList
        }
      ],
      yAxis : [
        {
          name: '执行数（个）',
          type : 'value'
        }
      ],
      series : [
        {
          name:'阻塞',
          type:'bar',
          stack: '过程',
          data: this._data.blockList
        },
        {
          name:'失败',
          type:'bar',
          stack: '过程',
          data: this._data.failList
        },
        {
          name:'通过',
          type:'bar',
          stack: '过程',
          data: this._data.passList
        }
      ]
    };
  }

}
