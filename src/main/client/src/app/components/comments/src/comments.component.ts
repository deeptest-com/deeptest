import * as _ from "lodash";

import {Component, Input, Output, OnInit, AfterViewInit, ViewChild, ElementRef, Renderer2} from "@angular/core";

import {CONSTANT} from "../../../utils/constant";
import {CommentsService} from "./comments.service";

@Component({
  selector: 'comments',
  templateUrl: './comments.html',
  styleUrls: ['./styles.scss']
})
export class CommentsComponent implements OnInit {
  @Input() @Output() content: string = '';
  @ViewChild('text') text: ElementRef;

  constructor(private renderer: Renderer2, public commentsService: CommentsService) {

  }

  ngOnInit(): any {
    let text = this.renderer.selectRootElement('#text');
    text.focus();
  }

  add(): any {
    console.log('===', this.content)


  }

}
