import {Component, Input, OnInit} from "@angular/core";

@Component({
  selector: 'design-process',
  templateUrl: './design-process.html',
  styleUrls: ['./styles.scss']
})
export class DesignProcessComponent implements OnInit {

  @Input() data: any;
  chartOption: any;

  constructor() {
  }

  ngOnInit(): any {
    this.chartOption = {
      title: {
        show: false
      },
      tooltip : {
        trigger: 'axis',
        axisPointer : {
          type : 'shadow'
        }
      },
      grid: {
        top: '5%',
        right: '15%',
        bottom: '8%',
        left: '1%',
        containLabel: true
      },
      legend: {
        right: '0%',
        width: '15%',
        data:['通过','失败','阻塞']
      },
      color: ['#749f83', '#c23531', '#ca8622', '#c4ccd3'],

      xAxis : [
        {
          type : 'category',
          axisLabel :{
            interval: 0,
            rotate: 45,
            margin: 10
          },
          data : [
            '07-01','07-02','07-03','07-04','07-05',
            '07-06','07-07','07-01','07-02','07-03',
            '07-04','07-05','07-06','07-07','07-01',
            '07-04','07-05','07-06','07-07','07-01',
            '07-04','07-05','07-06','07-07','07-01',
            '07-02','07-03','07-04','07-05','07-06','07-07']
        }
      ],
      yAxis : [
        {
          type : 'value'
        }
      ],
      series : [

        {
          name:'通过',
          type:'bar',
          stack: '过程',
          data:[
            120, 132, 101, 134, 90,
            230, 210,120, 132, 101,
            230, 210,120, 132, 101,
            134, 90, 230, 210,120,
            134, 90, 230, 210,120,
            132, 101, 134, 90, 230, 210]
        },
        {
          name:'失败',
          type:'bar',
          stack: '过程',
          data:[
            20, 12, 11, 24, 9,
            3, 10,20, 12, 11,
            24, 9, 3, 10,20,
            3, 10,20, 12, 11,
            24, 9, 3, 10,20,
            12, 11, 24, 9, 3, 10]
        },
        {
          name:'阻塞',
          type:'bar',
          stack: '过程',
          data:[
            5, 6, 3, 0, 7,
            0, 3, 5, 6, 3,
            0, 7, 0, 3, 5,
            0, 3, 5, 6, 3,
            0, 7, 0, 3, 5,
            6, 3, 0, 7, 0, 3]
        },
/*        {
          name:'未执行',
          type:'bar',
          stack: '过程',
          data:[150, 232, 201, 154, 190, 40, 20]
        }*/

      ]
    };
  }

}
