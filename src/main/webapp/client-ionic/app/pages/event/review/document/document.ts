import {Component} from '@angular/core';
import {NavController} from 'ionic-angular';

import {CONSTANT} from '../../../../utils/constant';
import {ImgPathPipe} from '../../../../pipes/img-path';
import {IosDatePipe} from '../../../../pipes/ios-date';

import {PubSubService} from '../../../../services/pub-sub-service';
import {CommonService}    from '../../../../services/common';
import {PostService}    from '../../../../services/post';

import {DocumentService}    from '../../../../services/document';

@Component({
  templateUrl: 'build/pages/event/review/document/document.html',
  providers: [CommonService, PostService, DocumentService],
  directives: [],
  pipes: [ImgPathPipe, IosDatePipe]
})
export class Document {
  errorMessage: any;
  data: any;    

  constructor(private nav: NavController,
              private _documentService: DocumentService, private _commonService: CommonService) {
    let me = this;

  }

  ngOnInit() {
    this.loadData();
  }

  onPageDidEnter(): void {

  }

  loadData():void {
    var me = this;
//    me._documentService.getDetail({eventId: -1}).subscribe((json) => {
//          me.data = json.data;
//    });
  }
  
}
