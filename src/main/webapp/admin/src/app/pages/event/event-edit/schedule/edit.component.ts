import { Component,ViewEncapsulation, Pipe, OnInit, AfterViewInit, ViewChild } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { Router, ActivatedRoute, Params } from '@angular/router';

import { DropdownModule} from 'ng2-bootstrap/ng2-bootstrap';

import {Validate} from '../../../../service/validate';
import { CONSTANT } from '../../../../utils/constant';
import { Utils } from '../../../../utils/utils';

import { ModalDirective } from 'ng2-bootstrap';

import { DatetimePickerService } from '../../../../service/datetime-picker';
import { EventService } from '../../../../service/event';
import { SessionService } from '../../../../service/session';
import { ScheduleService } from '../../../../service/schedule';

require('bootstrap-datepicker');
require('bootstrap-timepicker');

declare var jQuery;

@Component({
  selector: 'event-edit-schedule',
  encapsulation: ViewEncapsulation.None,
  styles: [require('./edit.scss'),
    require('bootstrap-datepicker/dist/css/bootstrap-datepicker3.css'),
    require('bootstrap-timepicker/css/bootstrap-timepicker.css')],

  template: require('./edit.html')
})
export class EventEditSchedule implements OnInit, AfterViewInit {
  @ViewChild('editPopup') public editPopup:ModalDirective;
  @ViewChild('alertPopup') public alertPopup:ModalDirective;

  eventId: number;
  tabModel: string = 'schedule';
  datePickers: string[] = ['startDate', 'endDate'];
  timePickers: string[] = ['startTime', 'endTime'];

  modelType: string;
  items: any;
  item: any = { signBefore: 3 };
  sessionForm: any;
  itemForm: any;

  constructor(private _router: Router, private _route: ActivatedRoute, private fb: FormBuilder,
              private _datetimePickerService: DatetimePickerService,
              private _eventService: EventService, private _sessionService: SessionService, private _scheduleService: ScheduleService) {

    let that = this;
  }

  ngOnInit() {
    let that = this;

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

    console.log(tabModel);
    that._router.navigateByUrl('/pages/event/edit/' + that.eventId + '/' + tabModel);
  }

  loadData() {
   let that = this;

    that._scheduleService.listByEvent(that.eventId).subscribe((json:any) => {
      that.items = json.bySession;
    });
  }

  delete():void {
    let that = this;

    that._sessionService.remove(that.item.id, that.modelType).subscribe((json:any) => {
      if (json.code = 1) {
        that.hideAlertModal();
        that.loadData();
      }
    });
  }

  showEditModal(item: any, type: any, event:any):void {
    let that = this;
    console.log(type, item);

    that.modelType = type;
    that.item = item;
    that.editPopup.show();

    event.stopPropagation();
  }
  showAlertModal(item: any, type: any, event:any):void {
    let that = this;
    console.log(type, item);

    that.modelType = type;
    that.item = item;

    that.alertPopup.show();

    event.stopPropagation();
  }

  onEditModalShow():void {
    let that = this;
    console.log('onEditModalShow');

    if (that.modelType == 'item') {
      Utils.dateDivide(that.item, 'startDate', 'startTime', 'startDatetime');
      Utils.dateDivide(that.item, 'endDate', 'endTime', 'endDatetime');

      that.initItemForm(true);
    }
  }


  hideEditModal():void {
    let that = this;
    that.editPopup.hide();
  }
  hideAlertModal():void {
    let that = this;
    that.alertPopup.hide();
  }

  initItemForm(dataLoaded): void {
    let that = this;

    jQuery.each(that.datePickers, function( index, value ) {
      that._datetimePickerService.genDatePicker(that.item, value);
    });
    jQuery.each(that.timePickers, function( index, value ) {
      that._datetimePickerService.genTimePicker(that.item, value);
    });
  }
  buildForm(): void {
    let that = this;
    this.sessionForm = this.fb.group(
      {
        'name': [that.item.name, [Validators.required]],
        'address': [that.item.address, [Validators.required]],
        'host': [that.item.address, [Validators.required]]
      }, {}
    );
    this.itemForm = this.fb.group(
      {
        'subject': [that.item.subject, [Validators.required]],
        // 'guest': [that.item.guest, [Validators.required]],
        'startDate': [that.item.startDate, [Validators.required, Validate.dateValidator()]],
        'startTime': [that.item.startTime, [Validators.required, Validate.timeValidator()]],
        'endDate': [that.item.endDate, [Validators.required, Validate.dateValidator()]],
        'endTime': [that.item.endTime, [Validators.required, Validate.timeValidator()]],
      }, {
        validator: Validate.compareDatetime([
          ['datetimeCompare', 'startDate','startTime','endDate','endTime']
        ])
      }
    );

    this.sessionForm.valueChanges.subscribe(data => this.onSessionValueChanged(data));
    this.itemForm.valueChanges.subscribe(data => this.onItemValueChanged(data));

    this.onSessionValueChanged();
    this.onItemValueChanged();
  }
  onSessionValueChanged(data?: any) {
    let that = this;
    that.sessionFormErrors = Validate.genValidateInfo(that.sessionForm, that.validateMsg);
  }
  onItemValueChanged(data?: any) {
    let that = this;
    that.itemFormErrors = Validate.genValidateInfo(that.itemForm, that.validateMsg, ['datetimeCompare']);
  }

  onSessionFormSubmit() {
    let that = this;

    let data = Object.assign({}, that.item, {children: undefined});
    that._sessionService.save(data).subscribe((json:any) => {
      if (json.code = 1) {
        that.hideEditModal();
        that.loadData();
      }
    });
  }
  onItemFormSubmit() {
    let that = this;

    Utils.dateCombine(that.item, 'startDate', 'startTime', 'startDatetime');
    Utils.dateCombine(that.item, 'endDate', 'endTime', 'endDatetime');

    console.log(that.item);
    that._scheduleService.save(that.item).subscribe((json:any) => {
      if (json.code = 1) {
        that.hideEditModal();
        that.loadData();
      }
    });
  }

  itemFormErrors = [];
  sessionFormErrors = [];
  validateMsg = {
    'name': {
      'required':      '名称不能为空'
    },
    'subject': {
      'required':      '主题不能为空'
    },
    'startDate': {
      'required':      '开始日期不能为空',
      'dateValidator': '开始日期格式不正确'
    },
    'startTime': {
      'required':      '开始时间不能为空',
      'timeValidator': '开始日期格式不正确'
    },
    'endDate': {
      'required':      '结束日期不能为空'
    },
    'endTime': {
      'required':      '结束时间不能为空',
      'timeValidator': '结束时间格式不正确'
    },

    'address': {
      'required':      '地址不能为空'
    },
    'host': {
      'required':      '主持人不能为空'
    },
    // 'guest': {
    //   'required':      '嘉宾不能为空'
    // },

    'datetimeCompare': '结束时间不能早于开始时间'
  };

}
