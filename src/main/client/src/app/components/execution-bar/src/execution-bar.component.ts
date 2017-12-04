import {Component, Input, OnInit} from "@angular/core";

@Component({
  selector: 'execution-bar',
  templateUrl: './execution-bar.html',
  styleUrls: ['./styles.scss']
})
export class ExecutionBarComponent implements OnInit {

  @Input() data: any;

  constructor() {

  }

  ngOnInit(): any {

  }

}
