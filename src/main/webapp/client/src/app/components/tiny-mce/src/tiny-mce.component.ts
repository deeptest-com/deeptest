import { Component, OnDestroy, AfterViewInit, OnChanges, Input, Output, EventEmitter, ElementRef } from '@angular/core';

import 'tinymce';
import 'tinymce/themes/modern';

import 'tinymce/plugins/table';
import 'tinymce/plugins/link';

declare var tinymce: any;

@Component({
  selector: 'tiny-mce',
  templateUrl: './tiny-mce.html',
  styleUrls: ['./tiny-mce.scss']
})
export class TinyMCEComponent implements OnDestroy, AfterViewInit, OnChanges {
  @Output() editorKeyup = new EventEmitter<any>();
  @Input() height: string;
  @Input() modelId: string;
  contentModel: string;
  @Input() set content(cont: string) {
    if (!cont) {
      this.contentModel = '';
    } else {
      this.contentModel = cont;
    }
  }

  constructor(private host: ElementRef) { }

  ngAfterViewInit() {
    let that = this;
    if (!$('textarea#mceEditor')) {
      return;
    }

    tinymce.init({
      document_base_url: '/assets/vendor/tinymce',
      selector: 'textarea#mceEditor',
      plugins: ['link', 'table'],
      skin_url: 'skins/lightgray',
      language : "zh_CN",
      language_url : "assets/vendor/tinymce/langs/zh_CN.js",
      setup: editor => {
        editor.on('keyup', () => {
          this.editorKeyup.emit(editor.getContent());
        });
      },
      height: this.height
    }).then(function(editors) {
      that.updateContent();
    });
  }

  ngOnChanges() {
    this.updateContent();
  }

  ngOnDestroy() {
    console.log("tinymce ngOnDestroy");
    if ($('textarea#mceEditor') && tinymce.get("mceEditor")) {
      tinymce.get("mceEditor").remove();
    }
  }

  updateContent() {
    this.contentModel = this.contentModel?this.contentModel:'';
    let editor = tinymce.get("mceEditor");
    if (editor) {editor.setContent(this.contentModel);}
  }

}
