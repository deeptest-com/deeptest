import {Component, Input, OnInit, Output, EventEmitter} from "@angular/core";
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

@Component({
  selector: 'search-select',
  templateUrl: './search-select.html',
  styleUrls: ['./styles.scss']
})
export class SearchSelectComponent implements OnInit {

  @Input() models: any[];

  @Output() itemSelect = new EventEmitter<any>();
  @Output() itemEnter = new EventEmitter<any>();
  @Output() searchChange = new EventEmitter<any>();

  keywords: string;
  selectedModel: any;
  formSelection: FormGroup;

  constructor(private fb: FormBuilder) {
    this.formSelection = this.fb.group(
      {
        'searchInput': ['', [Validators.required]]
      }, {}
    );
    this.formSelection.controls['searchInput'].valueChanges.debounceTime(500).subscribe(data => this.onSearchChanged(data));
  }

  ngOnInit(): any {

  }

  public onSelectItem($event, item):void {
    console.log('onSelectItem', item);

    this.itemSelect.emit(item);
  }
  public onKeyEnter($event):void {
    console.log('onKeyEnter', this.selectedModel );

    this.itemEnter.emit(this.selectedModel );
  }

  public onMouseEnter($event, item):void {
    console.log('onMouseEnter', item);
    this.selectedModel = item;
  }

  onSearchChanged(kewwords?: string) {
    this.searchChange.emit(kewwords);
  }

}
