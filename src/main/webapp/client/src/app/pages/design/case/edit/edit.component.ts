import {Component, ViewEncapsulation} from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { FileUploader, FileUploaderOptions } from 'ng2-file-upload';
import { NgModule, Pipe, OnInit, AfterViewInit }      from '@angular/core';

import {GlobalState} from '../../../../global.state';

import { CONSTANT } from '../../../../utils/constant';
import { Utils } from '../../../../utils/utils';
import {ValidatorUtils} from '../../../../validator/validator.utils';
import { RouteService } from '../../../../service/route';

import { CaseService } from '../../../../service/case';

declare var jQuery;

@Component({
  selector: 'case-edit',
  encapsulation: ViewEncapsulation.None,
  styles: [require('./edit.scss')],
  template: require('./edit.html')
})
export class CaseEdit implements OnInit, AfterViewInit {
  id: number;
  model: any = {};
  form: any;
  isSubmitted: boolean;
  uploadedFile: any;
  hasBaseDropZoneOver:boolean = false;

  private allowedMimeType: string[] = ['image/png', 'image/jpeg'];
  private uploaderOptions:FileUploaderOptions = {
    url: Utils.getUploadUrl(),
    authToken: CONSTANT.PROFILE.token,
    autoUpload: true,
    allowedMimeType: this.allowedMimeType,
    filters: [{name: 'upload', fn: (item:any) => {
      console.log(this.uploader);
      return true;
    }}]
  };
  public uploader: FileUploader;
  constructor(private _state:GlobalState, private _routeService: RouteService, private _route: ActivatedRoute, private fb: FormBuilder,
              private _caseService: CaseService) {

  }
  ngOnInit() {
    let that = this;

    that._route.params.forEach((params: Params) => {
      that.id = +params['id'];
    });

    if (that.id) {
      that.loadData();
    }
    that.buildForm();

    that.uploader = new FileUploader(that.uploaderOptions);
    that.uploader.onCompleteItem = (item:any, response:any, status:any, headers:any) => {
      this.onUploadCompleteItem(item, response, status, headers);
    };
  }
  ngAfterViewInit() {}

  selectFile():void {
    this.uploader.clearQueue();
    jQuery('#upload-input').click();
  }
  fileOver(e:any):void {
    this.hasBaseDropZoneOver = e;
  }
  onUploadCompleteItem (item:any, response:any, status:any, headers:any) {
    let res = JSON.parse(response);
    console.log(res);
    this.uploadedFile = res;
    this.model.file = res.uploadPath;
    this.uploader.clearQueue();
    this.isSubmitted = false;
  }

  buildForm(): void {
    let that = this;
    this.form = this.fb.group(
      {
        'name': [Validators.required],
        'title': [Validators.required],
        'descr': [Validators.required]
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
    'name': {
      'required':      '姓名不能为空'
    },
    'title': {
      'required':      '简介不能为空'
    },
    'descr': {
      'required':      '描述不能为空'
    }
  };

  loadData() {
    let that = this;
    that._caseService.get(that.id).subscribe((json:any) => {
      that.model = json.data;
    });
  }

  save() {
    let that = this;

    that._caseService.save(that.model).subscribe((json:any) => {
      if (json.code == 1) {
        that.model = json.data;
      }
    });
  }

}

