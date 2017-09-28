import {Component, Input, OnInit} from "@angular/core";
import { FormGroup, FormBuilder, FormControl, Validators } from '@angular/forms';
import {NgbActiveModal} from "@ng-bootstrap/ng-bootstrap";

import {CONSTANT} from "../../../utils/constant";

import {ZtreeService} from "../../ztree/src/ztree.service";
import {SuiteService} from "../../../service/suite";
import {CaseService} from "../../../service/case";

import {CaseSelectionService} from "./case-selection.service";

@Component({
  selector: 'case-selection',
  templateUrl: './case-selection.html',
  styleUrls: ['./styles.scss']
})
export class CaseSelectionComponent implements OnInit {

  @Input() treeModel: any;
  @Input() treeSettings: any = {};

  @Input() progress: string = '0';
  @Input() color: string = '#209e91';
  @Input() height: string = '1px';
  @Input() show: boolean = true;

  form: FormGroup;
  queryModel: any = {};

  public cases: string[];

  constructor(public activeModal: NgbActiveModal, private fb: FormBuilder, private _treeService: ZtreeService,
              public _sutieService: SuiteService, public _caseService: CaseService,) {
  }

  ngOnInit(): any {
    this.buildForm();

    this.loadData();
  }

  loadData() {
    this._caseService.query(3).subscribe((json:any) => {
      this.cases = json.data;
    });
  }

  save(): any {
    this.activeModal.close('save');
  }

  reset() {
    console.log('reset');
  }

  dismiss(): any {
    this.activeModal.dismiss('cancel');
  }

  onModuleSelected(event: any) {
    console.log('onNodeSelected', event);
  }
  onCaseSelected(item: any) {
    console.log('onCaseSelected', item);
  }

  buildForm(): void {
    this.form = this.fb.group(
      {
        'type': ['', []],
        'priority': ['', []],
        'estimate': ['', []],
        'createTime': ['', []],
        'updateTime': ['', []],
        'createUser': ['', []],
        'updateUser': ['', []]
      }, {}
    );

    this.form.valueChanges.debounceTime(CONSTANT.DebounceTime).subscribe(data => this.query(data));
  }
  formErrors = [];
  validateMsg = {

  };

  query(data?: any) {

  }

  resetFilters() {
    this.queryModel = {};
  }

}
