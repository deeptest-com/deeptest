import {Component, Input, OnInit} from "@angular/core";

@Component({
  selector: 'search-select',
  templateUrl: './search-select.html',
  styleUrls: ['./styles.scss']
})
export class SearchSelectComponent implements OnInit {

  @Input() data: any;

  constructor() {

  }

  ngOnInit(): any {

  }

}
