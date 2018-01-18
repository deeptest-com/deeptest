import {Component, Input, OnInit} from "@angular/core";

@Component({
  selector: 'report-design',
  templateUrl: './report-design.html',
  styleUrls: ['./styles.scss']
})
export class ReportDesignComponent implements OnInit {

  @Input() data: any;
  tab: string = 'design-result';

  constructor() {
  }

  ngOnInit(): any {

  }

  tabChange(event: any) {
    this.tab = event.nextId;
  }

}
