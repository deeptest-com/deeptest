import {Component, Input, OnInit} from "@angular/core";

@Component({
  selector: 'execution-result',
  templateUrl: './execution-result.html',
  styleUrls: ['./styles.scss']
})
export class ExecutionResultComponent implements OnInit {

  @Input() data: any;
  chartOption: any;

  constructor() {
  }

  ngOnInit(): any {
    this.chartOption = {
      title: {
        show: false
      },
      padding: 0,
      color: ['#749f83', '#c23531', '#ca8622', '#c4ccd3'],
      tooltip: {
        trigger: 'item',
        formatter: "{a} <br/>{b} : {c} ({d}%)"
      },
      legend: {
        show: false
      },
      series: [
        {
          name: '访问来源',
          type: 'pie',
          radius: '90%',
          center: ['50%', '55%'],
          data: [
            {value: 335, name: '成功'},
            {value: 310, name: '失败'},
            {value: 234, name: '阻塞'},
            {value: 135, name: '未执行'}
          ],
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
