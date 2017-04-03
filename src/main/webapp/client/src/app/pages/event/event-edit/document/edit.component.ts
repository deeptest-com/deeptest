import { Component,ViewEncapsulation, Pipe, OnInit, AfterViewInit, ViewChild, Input, ElementRef, Renderer } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { Router, ActivatedRoute, Params } from '@angular/router';

import { DropdownModule} from 'ng2-bootstrap/ng2-bootstrap';
import { ModalDirective } from 'ng2-bootstrap';
import { FileUploader, FileUploaderOptions } from 'ng2-file-upload';

import {ValidatorUtils} from '../../../../validator/validator.utils';
import { CONSTANT } from '../../../../utils/constant';
import { Utils } from '../../../../utils/utils';
import { RouteService } from '../../../../service/route';
import { EventService } from '../../../../service/event';
import { DocumentService } from '../../../../service/document';

declare var jQuery;

@Component({
  selector: 'event-edit-guest',
  encapsulation: ViewEncapsulation.None,
  styles: [require('./edit.scss')],

  template: require('./edit.html')
})
export class EventEditDocument implements OnInit, AfterViewInit {
  @ViewChild('editPopup') public editPopup:ModalDirective;
  @ViewChild('alertPopup') public alertPopup:ModalDirective;

  eventId: number;

  items: any;
  item: any = {};
  form: any;
  popupType: string;
  isSubmitted: boolean;
  uploadedFile: any;
  hasBaseDropZoneOver:boolean = false;

  totalItems:number = 0;
  currentPage:number = 1;
  itemsPerPage:number = 6;

  tabModel: string = 'document';
  needCreate:boolean = true;

  // private allowedMimeType: string[] = ['image/png', 'image/jpeg'];
  private uploaderOptions:FileUploaderOptions = {
    url: Utils.getUploadUrl(),
    authToken: CONSTANT.token,
    autoUpload: true,
    // allowedMimeType: this.allowedMimeType,
    filters: [{name: 'upload', fn: (item:any) => {
      console.log(item.name);
      return true;
    }}]
  };
  public uploader: FileUploader;

  constructor(private _routeService: RouteService, private _route: ActivatedRoute, private fb: FormBuilder,
              private _ducumentService: DocumentService) {

    let that = this;
  }

  ngOnInit() {
    let that = this;

    that._route.params.forEach((params: Params) => {
      that.eventId = +params['id'];
    });

    if (that.eventId) {
        that.loadData();
    }
    that.buildForm();

    that.uploader = new FileUploader(that.uploaderOptions);
    that.uploader.onCompleteItem = (item:any, response:any, status:any, headers:any) => {
      this.onUploadCompleteItem(item, response, status, headers);
    };
    console.log(that.uploader);
  }

  selectFile():void {
    this.uploader.clearQueue();
    jQuery('#upload-input').click();
  }
  fileOver(e:any):void {
    this.hasBaseDropZoneOver = e;
    // console.log(this.uploader.queue);
  }
  onUploadCompleteItem (item:any, response:any, status:any, headers:any) {
    let res = JSON.parse(response);
    console.log(res);
    this.uploadedFile = res;
    this.item.uri = res.uploadPath;
    this.uploader.clearQueue();
    this.isSubmitted = false;
  }

  ngAfterViewInit() {
    let that = this;

  }

  loadData() {
   let that = this;

   that._ducumentService.list(that.itemsPerPage, that.currentPage, that.eventId).subscribe((json:any) => {
     that.totalItems = json.totalItems;
     that.items = json.data;
   });
  }

  create($event): void {
    let that = this;
    that.showModal({eventId: that.eventId}, 'edit', $event);
  }
  pageChanged(event:any):void {
    let that = this;
    that.currentPage = event.page;
    that.loadData();
  }

  goto($event) {
    let that = this;
    that._routeService.navTo('/pages/event/edit/' + that.eventId + '/' + $event.tabModel);
  }

  showModal(item: any, popupType: string, $event:any):void {
    let that = this;

    that.popupType = popupType;
    that.item = item;

    if (that.popupType == 'edit') {
      that.editPopup.show();
    } else {
      that.alertPopup.show();
    }

    $event.stopPropagation();
  }

  onModalShow():void {
    let that = this;
    // init jquery components if needed
  }

  onFormSubmit() {
    let that = this;

    that._ducumentService.save(that.item).subscribe((json:any) => {
      if (json.code == 1) {
        that.hideModal();
        that.loadData();
      }
    });
  }

  remove():void {
    let that = this;

    that._ducumentService.remove(that.item.id).subscribe((json:any) => {
      if (json.code == 1) {
        that.hideModal();
        that.loadData();
      }
    });
  }

  hideModal():void {
    let that = this;
    if (that.popupType == 'edit') {
      that.editPopup.hide();
    } else {
      that.alertPopup.hide();
    }
  }

  buildForm(): void {
    let that = this;
    this.form = this.fb.group(
        {
          'title': ['', [Validators.required]],
          'uri': ['', [Validators.required]]
        }, {}
    );

    this.form.valueChanges.subscribe(data => this.onValueChanged(data));
    this.onValueChanged();
  }
  onValueChanged(data?: any) {
    let that = this;
    that.formErrors = ValidatorUtils.genMsg(that.form, that.validateMsg, []);
  }

  formErrors = [];
  validateMsg = {
    'title': {
      'required':      '简介不能为空'
    },
    'uri': {
      'required':      '文件不能为空'
    }
  };

}
