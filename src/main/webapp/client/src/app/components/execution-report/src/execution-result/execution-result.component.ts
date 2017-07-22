import {Component, Input, OnInit} from "@angular/core";

@Component({
  selector: 'execution-result',
  templateUrl: './execution-result.html',
  styleUrls: ['./styles.scss']
})
export class ExecutionResultComponent implements OnInit {

  @Input() data: any;

  constructor() {
  }

  ngOnInit(): any {

  }

}
