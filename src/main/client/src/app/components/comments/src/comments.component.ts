import * as _ from "lodash";

import {Component, Input, Output, OnInit, AfterViewInit} from "@angular/core";

import {CONSTANT} from "../../../utils/constant";
import {CommentsService} from "./comments.service";

@Component({
  selector: 'comments',
  templateUrl: './comments.html',
  styleUrls: ['./styles.scss']
})
export class CommentsComponent implements OnInit {
  @Input() @Output() content: string = '';

  constructor(public commentsService: CommentsService) {

  }

  ngOnInit(): any {

  }

  add(): any {
    console.log('===', this.content)


  }

}
