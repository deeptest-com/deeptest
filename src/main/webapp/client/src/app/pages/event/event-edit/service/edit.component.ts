import {Component, ViewEncapsulation, OnInit, AfterViewInit, ViewChild} from "@angular/core";
import {FormBuilder, Validators} from "@angular/forms";
import {Router, ActivatedRoute, Params} from "@angular/router";
import {ModalDirective} from "ng2-bootstrap";
import {ValidatorUtils} from '../../../../validator/validator.utils';
import { RouteService } from '../../../../service/route';
import {ServiceService} from "../../../../service/service";

require('bootstrap-datepicker');
require('bootstrap-timepicker');

declare var jQuery;

@Component({
  selector: 'event-edit-service',
  encapsulation: ViewEncapsulation.None,
  styles: [require('./edit.scss')],

  template: require('./edit.html')
})
export class EventEditService implements OnInit, AfterViewInit {
  @ViewChild('editPopup') public editPopup:ModalDirective;
  @ViewChild('alertPopup') public alertPopup:ModalDirective;

  eventId:number;

  items:any;
  item:any = {};
  form:any;
  popupType:string;
  isSubmitted:boolean;

  tabModel:string = 'service';
  needCreate:boolean = false;

  constructor(private _routeService: RouteService, private _route:ActivatedRoute, private fb:FormBuilder,
              private _serviceService:ServiceService) {

    let that = this;
  }

  ngOnInit() {
    let that = this;

    that._route.params.forEach((params:Params) => {
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

  loadData() {
    let that = this;

    that._serviceService.list(that.eventId).subscribe((json:any) => {
      that.items = json.services;
    });
  }

  goto($event) {
    let that = this;

    that._routeService.navTo('/pages/event/edit/' + that.eventId + '/' + $event.tabModel);
  }

  showModal(item:any, popupType:string, $event:any):void {
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

  disable(item:any, $event:any):void {
    let that = this;
    that._serviceService.disable(item.id).subscribe((json:any) => {
      if (json.code == 1) {
        that.loadData();
      }
    });
    $event.stopPropagation();
  }

  onModalShow():void {
    let that = this;
    // init jquery components if needed
  }

  onFormSubmit() {
    let that = this;

    that._serviceService.save(that.item).subscribe((json:any) => {
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

  buildForm():void {
    let that = this;
    this.form = this.fb.group(
      {
        'descr': ['', [Validators.required]]
      }, {}
    );

    this.form.valueChanges.subscribe(data => this.onValueChanged(data));
    this.onValueChanged();
  }

  onValueChanged(data?:any) {
    let that = this;
    that.formErrors = ValidatorUtils.genMsg(that.form, that.validateMsg, []);
  }

  formErrors = [];
  validateMsg = {
    'descr': {
      'required': '描述不能为空'
    }
  };
}
