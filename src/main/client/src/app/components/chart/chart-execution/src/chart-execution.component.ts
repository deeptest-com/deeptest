import {Component, Input, OnInit} from "@angular/core";

@Component({
  selector: 'chart-execution',
  templateUrl: './chart-execution.html',
  styleUrls: ['./styles.scss']
})
export class ChartExecutionComponent implements OnInit {

  @Input() chartData: any;
  tab: string = 'exe-result';

  constructor() {
  }

  ngOnInit(): any {
  }

  tabChange(event: any) {
    this.tab = event.nextId;
  }

}
