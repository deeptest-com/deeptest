import {Component, Input, OnInit} from "@angular/core";

@Component({
  selector: 'execution-report',
  templateUrl: './execution-report.html',
  styleUrls: ['./styles.scss']
})
export class ExecutionReportComponent implements OnInit {

  @Input() data: any;
  tab: string = 'exe-result';

  constructor() {
  }

  ngOnInit(): any {

  }

  tabChange(event: any) {
    console.log('===', event.nextId);
    this.tab = event.nextId;
  }

}
