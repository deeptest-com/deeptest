import {Component, Input, OnInit} from "@angular/core";

@Component({
  selector: 'report-execution',
  templateUrl: './report-execution.html',
  styleUrls: ['./styles.scss']
})
export class ReportExecutionComponent implements OnInit {

  @Input() data: any;
  tab: string = 'exe-result';

  constructor() {
  }

  ngOnInit(): any {

  }

  tabChange(event: any) {
    this.tab = event.nextId;
  }

}
