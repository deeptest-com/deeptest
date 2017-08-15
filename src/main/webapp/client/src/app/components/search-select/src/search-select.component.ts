import {Component, Input, OnInit, Output, EventEmitter} from "@angular/core";
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

@Component({
  selector: 'search-select',
  templateUrl: './search-select.html',
  styleUrls: ['./styles.scss']
})
export class SearchSelectComponent implements OnInit {

  @Input() searchResult: any[];
  @Input() selectedModels: any[];

  @Output() searchChange = new EventEmitter<any>();

  keywords: string;

  selectedModel: any;

  formSelection: FormGroup;

  constructor(private fb: FormBuilder) {
    this.formSelection = this.fb.group(
      {
        'searchInput': ['', []]
      }, {}
    );
    this.formSelection.controls['searchInput'].valueChanges.debounceTime(500).subscribe(data => this.onSearchChanged(data));
  }

  ngOnInit(): any {

  }

  public onMouseEnter($event, item):void {
    console.log('onMouseEnter', item);

    $event.preventDefault();
    $event.stopPropagation();

    this.selectedModel = item;
  }

  public onSelectItem($event, item):void {
    console.log('onSelectItem', item);

    this.keywords = '';
    this.searchResult = [];

    this.selectedModels.push(item);
  }

  onSearchChanged(kewwords?: string) {
    if (!kewwords) {
      this.searchResult = null;
      return;
    }

    this.searchChange.emit(kewwords);
  }

  remove(item: any) {
    console.log(item);

    this.selectedModels = this.selectedModels.filter(obj => obj.id !== item.id);
  }

}
