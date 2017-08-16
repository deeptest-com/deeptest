import {Component, Input, OnInit, Output, EventEmitter} from "@angular/core";
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

@Component({
  selector: 'search-select',
  templateUrl: './search-select.html',
  styleUrls: ['./styles.scss']
})
export class SearchSelectComponent implements OnInit {

  @Input() searchResult: any[];
  @Input() @Output() selectedModels: any[];

  @Output() searchChange = new EventEmitter<any>();

  @Input() @Output() keywords: string;

  selectedModel: any;
  formSelection: FormGroup;

  constructor(private fb: FormBuilder) {
    this.formSelection = this.fb.group(
      {
        'searchInput': ['', []]
      }, {}
    );
    this.formSelection.controls['searchInput'].valueChanges.debounceTime(500).subscribe(data => this.onSearchChanged());
  }

  ngOnInit(): any {

  }

  public onMouseEnter($event, item):void {
    $event.preventDefault();
    $event.stopPropagation();

    this.selectedModel = item;
  }

  public onSelectItem($event, item):void {
    this.keywords = '';
    this.searchResult = null;

    this.selectedModels.push(item);
  }

  onSearchChanged() {
    if (!this.keywords) {
      this.searchResult = null;
      return;
    }
    this.searchChange.emit(this.keywords);
  }

  remove(item: any) {

    var index = this.selectedModels.indexOf(item, 0);
    if (index > -1) {
      this.selectedModels.splice(index, 1);
    }
  }

  cancel($event):void {
    $event.preventDefault();
    $event.stopPropagation();

    this.searchResult = null;
    this.keywords = '';
  }

}
