import { Component,ViewEncapsulation, Pipe, OnInit, AfterViewInit, ViewChild } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { Router, ActivatedRoute, Params } from '@angular/router';

import { DropdownModule} from 'ng2-bootstrap/ng2-bootstrap';
import { ModalDirective } from 'ng2-bootstrap';

import {Validate} from '../../../../service/validate';
import { CONSTANT } from '../../../../utils/constant';
import { Utils } from '../../../../utils/utils';
import { EventService } from '../../../../service/event';
import { GuestService } from '../../../../service/guest';

require('bootstrap-datepicker');
require('bootstrap-timepicker');

declare var jQuery;

@Component({
  selector: 'event-edit-guest',
  encapsulation: ViewEncapsulation.None,
  styles: [require('./edit.scss'),
    require('bootstrap-datepicker/dist/css/bootstrap-datepicker3.css'),
    require('bootstrap-timepicker/css/bootstrap-timepicker.css')],

  template: require('./edit.html')
})
export class EventEditGuest implements OnInit, AfterViewInit {
  @ViewChild('editPopup') public editPopup:ModalDirective;
  @ViewChild('alertPopup') public alertPopup:ModalDirective;

  eventId: number;

  items: any;
  item: any = {};
  form: any;
  popupType: string;

  totalItems:number = 0;
  currentPage:number = 1;
  itemsPerPage:number = 6;

  tabModel: string = 'guest';

  constructor(private _router: Router, private _route: ActivatedRoute, private fb: FormBuilder,
              private _guestService: GuestService, private _eventService: EventService) {

    let that = this;
  }

  ngOnInit() {
    let that = this;

    console.log(this._route.params);
    this._route.params.forEach((params: Params) => {
      that.eventId = +params['id'];
    });

    if (that.eventId) {
        that.loadData();
    }
    that.buildForm();
  }

  ngAfterViewInit() {
    let that = this;

  }

  back() {
    let that = this;

    that._router.navigateByUrl("/pages/event/list");
  }
  goto(tabModel) {
    let that = this;

    that._router.navigateByUrl('/pages/event/edit/' + that.eventId + '/' + tabModel);
  }

  loadData() {
   let that = this;

   that._guestService.list(that.itemsPerPage, that.currentPage, that.eventId).subscribe((json:any) => {
      that.items = json.guests;
   });
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

    that._guestService.save(that.item).subscribe((json:any) => {
      if (json.code = 1) {
        that.hideModal();
        that.loadData();
      }
    });
  }

  remove():void {
    let that = this;

    that._guestService.remove(that.item.id).subscribe((json:any) => {
      if (json.code = 1) {
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

  initForm(): void {
    let that = this;

  }

  buildForm(): void {
    let that = this;
    this.form = this.fb.group(
        {
          'name': [that.item.name, [Validators.required]],
          'title': [that.item.title, [Validators.required]],
          'descr': [that.item.descr, [Validators.required]]
        }, {}
    );

    this.form.valueChanges.subscribe(data => this.onValueChanged(data));
    this.onValueChanged();
  }
  onValueChanged(data?: any) {
    let that = this;
    that.formErrors = Validate.genValidateInfo(that.form, that.validateMsg, []);
  }

  formErrors = [];
  validateMsg = {
    'name': {
      'required':      '姓名不能为空'
    },
    'title': {
      'required':      '简介不能为空'
    },
    '描述': {
      'required':      '描述不能为空'
    }
  };

}
